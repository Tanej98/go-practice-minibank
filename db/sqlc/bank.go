package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Bank struct {
	*Queries
	db *sql.DB
}

func NewBank(db *sql.DB) *Bank {
	return &Bank{
		db:      db,
		Queries: New(db),
	}
}

func (bank *Bank) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := bank.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err : %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxResponse struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (bank *Bank) TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResponse, error) {
	var result TransferTxResponse

	err := bank.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, arg)
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
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

		if arg.FromAccountID < arg.ToAccountID{
			result.FromAccount, result.ToAccount, err =  addMoney(ctx, q, arg.FromAccountID, arg.ToAccountID, -arg.Amount, arg.Amount)
			if err != nil {
				return err
			}
		}else{
			result.ToAccount, result.FromAccount, err =  addMoney(ctx, q, arg.ToAccountID, arg.FromAccountID, arg.Amount, -arg.Amount)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}


func addMoney(ctx context.Context, q *Queries, fromAccount , toAccount, debitAmount, creditAmount int64) (account1 , account2 Account, err error){
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     fromAccount,
		Amount: debitAmount,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     toAccount,
		Amount: creditAmount,
	})
	if err != nil {
		return
	}
	return
}