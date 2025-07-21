package transaction

import (
	"context"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	transactionRepository repository.TransactionRepository
	userService           user.Service
}

type Service interface {
	CreateTransaction(ctx context.Context, transactionInfo Transaction) (Transaction, error)
	GetTransactionByContributionId(ctx context.Context, contributionId int) (Transaction, error)
	CreateTransactionForContribution(ctx context.Context, contribution Contribution) (Transaction, error)
	HandleTransactionCreation(ctx context.Context, contribution Contribution) (Transaction, error)
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

func (s *service) GetTransactionByContributionId(ctx context.Context, contributionId int) (Transaction, error) {
	transaction, err := s.transactionRepository.GetTransactionByContributionId(ctx, nil, contributionId)
	if err != nil {
		slog.Error("error fetching transaction using contribution id", "error", err)
		return Transaction{}, err
	}

	return Transaction(transaction), nil
}

func (s *service) CreateTransactionForContribution(ctx context.Context, contribution Contribution) (Transaction, error) {
	transactionInfo := Transaction{
		UserId:            contribution.UserId,
		ContributionId:    contribution.Id,
		IsRedeemed:        false,
		IsGained:          true,
		TransactedBalance: contribution.BalanceChange,
		TransactedAt:      contribution.ContributedAt,
	}
	transaction, err := s.CreateTransaction(ctx, transactionInfo)
	if err != nil {
		slog.Error("error creating transaction for current contribution", "error", err)
		return Transaction{}, err
	}

	return transaction, nil
}

func (s *service) HandleTransactionCreation(ctx context.Context, contribution Contribution) (Transaction, error) {
	var transaction Transaction

	transaction, err := s.GetTransactionByContributionId(ctx, contribution.Id)
	if err != nil {
		if err == apperrors.ErrTransactionNotFound {
			transaction, err = s.CreateTransactionForContribution(ctx, contribution)
			if err != nil {
				slog.Error("error creating transaction for exisiting contribution", "error", err)
				return Transaction{}, err
			}
		} else {
			slog.Error("error fetching transaction", "error", err)
			return Transaction{}, err
		}
	}

	return transaction, nil
}
