package transaction

import (
	"database/sql"
	"jukebox-app/pkg/datasource"

	"go.uber.org/zap"
)

type DBTransactionContext struct{}

type DBTransactionHandlerFunction func(tx *sql.Tx) error

type DBTransactionHandler interface {
	HandleTransaction(fn DBTransactionHandlerFunction) error
}

type DefaultDBTransactionHandler struct {
	datasource.DBDataSource
}

func NewDefaultDBTransactionHandler(dbDatasource datasource.DBDataSource) *DefaultDBTransactionHandler {
	return &DefaultDBTransactionHandler{
		DBDataSource: dbDatasource,
	}
}

func (handler *DefaultDBTransactionHandler) HandleTransaction(fn DBTransactionHandlerFunction) error {

	db, err := handler.GetDatabase()
	if err != nil {
		handleError(err)
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		handleError(err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			handleError(tx.Rollback())
			zap.L().Error(err.Error())
		} else if err != nil {
			// something went wrong, rollback
			handleError(tx.Rollback())
		} else {
			// all good, commit
			err = tx.Commit()
			handleError(err)
		}
	}()

	err = fn(tx)
	return err
}

func handleError(err error) {
	if err != nil {
		zap.L().Error(err.Error())
	}
}
