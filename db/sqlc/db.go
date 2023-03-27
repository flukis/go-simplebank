// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addBalanceAccountStmt, err = db.PrepareContext(ctx, addBalanceAccount); err != nil {
		return nil, fmt.Errorf("error preparing query AddBalanceAccount: %w", err)
	}
	if q.createAccountStmt, err = db.PrepareContext(ctx, createAccount); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccount: %w", err)
	}
	if q.createEntryStmt, err = db.PrepareContext(ctx, createEntry); err != nil {
		return nil, fmt.Errorf("error preparing query CreateEntry: %w", err)
	}
	if q.createTransferStmt, err = db.PrepareContext(ctx, createTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTransfer: %w", err)
	}
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.fetchAccountsStmt, err = db.PrepareContext(ctx, fetchAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query FetchAccounts: %w", err)
	}
	if q.fetchEntriesStmt, err = db.PrepareContext(ctx, fetchEntries); err != nil {
		return nil, fmt.Errorf("error preparing query FetchEntries: %w", err)
	}
	if q.fetchTransferStmt, err = db.PrepareContext(ctx, fetchTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query FetchTransfer: %w", err)
	}
	if q.getAccountStmt, err = db.PrepareContext(ctx, getAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccount: %w", err)
	}
	if q.getAccountForUpdateStmt, err = db.PrepareContext(ctx, getAccountForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccountForUpdate: %w", err)
	}
	if q.getEntryStmt, err = db.PrepareContext(ctx, getEntry); err != nil {
		return nil, fmt.Errorf("error preparing query GetEntry: %w", err)
	}
	if q.getTransferStmt, err = db.PrepareContext(ctx, getTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransfer: %w", err)
	}
	if q.updateBalanceAccountStmt, err = db.PrepareContext(ctx, updateBalanceAccount); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateBalanceAccount: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addBalanceAccountStmt != nil {
		if cerr := q.addBalanceAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addBalanceAccountStmt: %w", cerr)
		}
	}
	if q.createAccountStmt != nil {
		if cerr := q.createAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccountStmt: %w", cerr)
		}
	}
	if q.createEntryStmt != nil {
		if cerr := q.createEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createEntryStmt: %w", cerr)
		}
	}
	if q.createTransferStmt != nil {
		if cerr := q.createTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTransferStmt: %w", cerr)
		}
	}
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.fetchAccountsStmt != nil {
		if cerr := q.fetchAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing fetchAccountsStmt: %w", cerr)
		}
	}
	if q.fetchEntriesStmt != nil {
		if cerr := q.fetchEntriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing fetchEntriesStmt: %w", cerr)
		}
	}
	if q.fetchTransferStmt != nil {
		if cerr := q.fetchTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing fetchTransferStmt: %w", cerr)
		}
	}
	if q.getAccountStmt != nil {
		if cerr := q.getAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountStmt: %w", cerr)
		}
	}
	if q.getAccountForUpdateStmt != nil {
		if cerr := q.getAccountForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountForUpdateStmt: %w", cerr)
		}
	}
	if q.getEntryStmt != nil {
		if cerr := q.getEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEntryStmt: %w", cerr)
		}
	}
	if q.getTransferStmt != nil {
		if cerr := q.getTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransferStmt: %w", cerr)
		}
	}
	if q.updateBalanceAccountStmt != nil {
		if cerr := q.updateBalanceAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateBalanceAccountStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                       DBTX
	tx                       *sql.Tx
	addBalanceAccountStmt    *sql.Stmt
	createAccountStmt        *sql.Stmt
	createEntryStmt          *sql.Stmt
	createTransferStmt       *sql.Stmt
	deleteAccountStmt        *sql.Stmt
	fetchAccountsStmt        *sql.Stmt
	fetchEntriesStmt         *sql.Stmt
	fetchTransferStmt        *sql.Stmt
	getAccountStmt           *sql.Stmt
	getAccountForUpdateStmt  *sql.Stmt
	getEntryStmt             *sql.Stmt
	getTransferStmt          *sql.Stmt
	updateBalanceAccountStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		addBalanceAccountStmt:    q.addBalanceAccountStmt,
		createAccountStmt:        q.createAccountStmt,
		createEntryStmt:          q.createEntryStmt,
		createTransferStmt:       q.createTransferStmt,
		deleteAccountStmt:        q.deleteAccountStmt,
		fetchAccountsStmt:        q.fetchAccountsStmt,
		fetchEntriesStmt:         q.fetchEntriesStmt,
		fetchTransferStmt:        q.fetchTransferStmt,
		getAccountStmt:           q.getAccountStmt,
		getAccountForUpdateStmt:  q.getAccountForUpdateStmt,
		getEntryStmt:             q.getEntryStmt,
		getTransferStmt:          q.getTransferStmt,
		updateBalanceAccountStmt: q.updateBalanceAccountStmt,
	}
}