package repository

import (
	"context"
	"database/sql"
	"errors"
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
	GetTransactionByContributionId(ctx context.Context, tx *sqlx.Tx, contributionId int) (Transaction, error)
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

	getTransactionByContributionIdQuery = `SELECT * from transactions where contribution_id=$1`
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

func (tr *transactionRepository) GetTransactionByContributionId(ctx context.Context, tx *sqlx.Tx, contributionId int) (Transaction, error) {
	executer := tr.BaseRepository.initiateQueryExecuter(tx)

	var transaction Transaction
	err := executer.GetContext(ctx, &transaction, getTransactionByContributionIdQuery, contributionId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("transaction for the contribution id does not exist", "error", err)
			return Transaction{}, apperrors.ErrTransactionNotFound
		}
		slog.Error("error fetching transaction using contributionid", "error", err)
		return Transaction{}, err
	}

	return transaction, nil
}
