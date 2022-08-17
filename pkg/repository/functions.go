package repository

import (
	"context"
	"database/sql"
	"jukebox-app/pkg/transaction"

	"go.uber.org/zap"
)

type RelationalFunction func(statement *sql.Stmt) error

func RelationalContext(ctx context.Context, sqlStatement string, fn RelationalFunction) error {

	var err error
	var statement *sql.Stmt
	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)
	if statement, err = tx.Prepare(sqlStatement); err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error("Error closing the statement")
		}
	}(statement)

	if err = fn(statement); err != nil {
		return err
	}
	return nil
}
