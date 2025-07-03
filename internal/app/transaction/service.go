package transaction

import (
	"context"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	transactionRepository repository.TransactionRepository
}

type Service interface {
	CreateTransaction(ctx context.Context, transactionInfo Transaction) (Transaction, error)
}

func NewService(transactionRepository repository.TransactionRepository) Service {
	return &service{
		transactionRepository: transactionRepository,
	}
}

func (s *service) CreateTransaction(ctx context.Context, transactionInfo Transaction) (Transaction, error) {
	transaction, err := s.transactionRepository.CreateTransaction(ctx, nil, repository.Transaction(transactionInfo))
	if err != nil {
		slog.Error("error occured while creating transaction", "error", err)
		return Transaction{}, err
	}

	return Transaction(transaction), nil
}
