package test

import (
	"context"
	"database/sql"
	"log"

	"github.com/eko/gocache/v2/store"
	"github.com/spf13/cobra"

	"jukebox-app/internal/config"
	"jukebox-app/internal/model"
	"jukebox-app/internal/repository"
	cachemanager "jukebox-app/pkg/cache-manager"
	"jukebox-app/pkg/transaction"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	environment := config.InitConfig(&args)
	defer config.StopConfig()

	dataSource := config.InitDB(environment)
	defer config.StopDB()

	txHandler := transaction.NewRelationalTransactionHandler(dataSource)
	userRepository := repository.NewRelationalUserRepository()

	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "Guido",
		Email: "guido.mantilla@yahoo.com",
	}
	cache, _ := cachemanager.NewCache(store.GoCacheType, environment)
	cacheManager := cachemanager.NewDefaultCacheManager(cache)
	cacheUserRepository := repository.NewCachedUserRepository(userRepository, cacheManager)

	var err error
	err = txHandler.HandleTransaction(func(tx *sql.Tx) error {

		txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)
		if err = cacheUserRepository.Create(txCtx, user); err != nil {
			return err
		}

		log.Println("UserRepository.Create Done: ", *user, err)

		localUser, err := cacheUserRepository.FindById(txCtx, int64(1))
		if err != nil {
			return err
		}

		log.Println("UserRepository.FindById Done: ", *localUser, err)

		return nil
	})
	log.Println("CacheUserRepository. Done: ")

}
