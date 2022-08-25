package main

import (
	"context"
	"database/sql"
	"jukebox-app/internal/config"
	"jukebox-app/internal/core/repository"
	"jukebox-app/pkg/transaction"
	"log"
	"testing"
)

func Test_Create(t *testing.T) {

	var args []string
	environment := config.InitConfig(&args)
	defer config.StopConfig()

	dataSource := config.InitDB(environment)
	defer config.StopDB()

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
