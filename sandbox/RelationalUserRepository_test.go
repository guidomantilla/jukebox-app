package main

import (
	"context"
	"database/sql"
	"jukebox-app/internal/repository"
	"log"
	"testing"

	"jukebox-app/internal/config"
	"jukebox-app/pkg/transaction"
)

func Test_Create(t *testing.T) {

	var args []string
	env := config.InitConfig(&args)
	defer func() {
		_ = config.StopConfig()
	}()

	dataSource := config.InitDB(env)
	defer func() {
		_ = config.StopDB()
	}()

	txHandler := transaction.NewRelationalTransactionHandler(dataSource)
	repository := repository.NewRelationalUserRepository()

	err := txHandler.HandleTransaction(func(tx *sql.Tx) error {

		txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

		user, err := repository.FindById(txCtx, 15)
		if err != nil {
			return err
		}
		log.Println(user)
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}
