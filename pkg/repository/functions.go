package repository

import (
	"context"
	"database/sql"
	"fmt"
	"jukebox-app/pkg/transaction"
	"strings"

	"go.uber.org/zap"
)

const (
	Error_Closing_Statement = "Error closing the statement"
	Error_Closing_ResultSet = "Error closing the result set"
)

type RelationalFunction func(statement *sql.Stmt) error
type RelationalQueryFunction func(rows *sql.Rows) error

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
			zap.L().Error(Error_Closing_Statement)
		}
	}(statement)

	if err = fn(statement); err != nil {
		return err
	}
	return nil
}

func RelationalQueryContext(ctx context.Context, sqlStatement string, fn RelationalQueryFunction) error {

	var err error
	var statement *sql.Stmt
	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)
	if statement, err = tx.Prepare(sqlStatement); err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error(Error_Closing_Statement)
		}
	}(statement)

	var rows *sql.Rows
	if rows, err = statement.Query(); err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			zap.L().Error(Error_Closing_ResultSet)
		}
	}(rows)

	if err = fn(rows); err != nil {
		return err
	}
	return nil
}

//

func RelationalWriteContext(ctx context.Context, sqlStatement string, args ...any) (*int64, error) {

	var err error
	var serial int64
	err = RelationalContext(ctx, sqlStatement, func(statement *sql.Stmt) error {

		var result sql.Result
		if result, err = statement.Exec(args...); err != nil {
			return err
		}

		if strings.Index(strings.ToLower(sqlStatement), "insert") == 0 {
			if serial, err = result.LastInsertId(); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &serial, nil
}

func RelationalQueryRowContext(ctx context.Context, sqlStatement string, key any, dest ...any) error {

	var err error
	var statement *sql.Stmt
	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)
	if statement, err = tx.Prepare(sqlStatement); err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error(Error_Closing_Statement)
		}
	}(statement)

	row := statement.QueryRow(key)
	if err = row.Scan(dest...); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return fmt.Errorf("row with key %d not found", key)
		}
		return err
	}

	return nil
}
