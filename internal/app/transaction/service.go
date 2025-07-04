package transaction

import (
	"context"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	transactionRepository repository.TransactionRepository
	userService           user.Service
}

type Service interface {
	CreateTransaction(ctx context.Context, transactionInfo Transaction) (Transaction, error)
}

func NewService(transactionRepository repository.TransactionRepository, userService user.Service) Service {
	return &service{
		transactionRepository: transactionRepository,
		userService:           userService,
	}
}

func (s *service) CreateTransaction(ctx context.Context, transactionInfo Transaction) (Transaction, error) {
	tx, err := s.transactionRepository.BeginTx(ctx)
	if err != nil {
		slog.Error("failed to start transaction creation")
		return Transaction{}, err
	}

	ctx = middleware.EmbedTxInContext(ctx, tx)

	defer func() {
		if txErr := s.transactionRepository.HandleTransaction(ctx, tx, err); txErr != nil {
			slog.Error("failed to handle transaction", "error", txErr)
			err = txErr
		}
	}()

	transaction, err := s.transactionRepository.CreateTransaction(ctx, tx, repository.Transaction(transactionInfo))
	if err != nil {
		slog.Error("error occured while creating transaction", "error", err)
		return Transaction{}, err
	}

	err = s.userService.UpdateUserCurrentBalance(ctx, user.Transaction(transaction))
	if err != nil {
		slog.Error("error occured while updating user current balance", "error", err)
		return Transaction{}, err
	}

	return Transaction(transaction), nil
}
