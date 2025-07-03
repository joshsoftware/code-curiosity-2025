package repository

import (
	"context"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
)

type transactionRepository struct {
	BaseRepository
}

type TransactionRepository interface {
	RepositoryTransaction
	CreateTransaction(ctx context.Context, tx *sqlx.Tx, transactionInfo Transaction) (Transaction, error)
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		BaseRepository: BaseRepository{db},
	}
}

const (
	createTransactionQuery = `INSERT INTO transactions (
	user_id,
	contribution_id,
	is_redeemed,
	is_gained,
	transacted_balance,
	transacted_at
	)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING *`
)

func (tr *transactionRepository) CreateTransaction(ctx context.Context, tx *sqlx.Tx, transactionInfo Transaction) (Transaction, error) {
	executer := tr.BaseRepository.initiateQueryExecuter(tx)

	var transaction Transaction
	err := executer.GetContext(ctx, &transaction, createTransactionQuery,
		transactionInfo.UserId,
		transactionInfo.ContributionId,
		transactionInfo.IsRedeemed,
		transactionInfo.IsGained,
		transactionInfo.TransactedBalance,
		transactionInfo.TransactedAt,
	)
	if err != nil {
		slog.Error("error occured while creating transaction", "error", err)
		return Transaction{}, apperrors.ErrTransactionCreationFailed
	}

	return transaction, nil
}
