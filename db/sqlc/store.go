package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTX execute a function  within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb err : %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other
// It creates a transfer racord, add account entries, and update accounts balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	var fct = func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// move money out of account1
		if arg.FromAccountID < arg.ToAccountID {
			acc1 := AddMoneyAccountParams{Amount: -arg.Amount, Id: arg.FromAccountID}
			acc2 := AddMoneyAccountParams{Amount: arg.Amount, Id: arg.ToAccountID}
			result.FromAccount, result.ToAccount, _ = addMoney(ctx, q, acc1, acc2)
		} else {
			acc1 := AddMoneyAccountParams{Amount: arg.Amount, Id: arg.ToAccountID}
			acc2 := AddMoneyAccountParams{Amount: -arg.Amount, Id: arg.FromAccountID}
			result.ToAccount, result.FromAccount, _ = addMoney(ctx, q, acc1, acc2)
		}

		return nil
	}

	err := store.execTx(ctx, fct)

	return result, err
}

type AddMoneyAccountParams struct {
	Id     int64
	Amount int64
}

func addMoney(ctx context.Context, q *Queries, firstAccount AddMoneyAccountParams, secondAccount AddMoneyAccountParams) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     firstAccount.Id,
		Amount: firstAccount.Amount,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     secondAccount.Id,
		Amount: secondAccount.Amount,
	})
	return
}
