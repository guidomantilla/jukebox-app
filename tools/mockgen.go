package tools

//go:generate mockgen -package cachemanager -destination ../pkg/cache-manager/MockCacheManager.go 				jukebox-app/pkg/cache-manager CacheManager
//go:generate mockgen -package datasource   -destination ../pkg/datasource/MockRelationalDataSource.go 			jukebox-app/pkg/datasource RelationalDataSource
//go:generate mockgen -package environment  -destination ../pkg/environment/MockEnvironment.go 					jukebox-app/pkg/environment Environment
//go:generate mockgen -package properties   -destination ../pkg/properties/MockPropertySource.go 				jukebox-app/pkg/properties PropertySource
//go:generate mockgen -package properties   -destination ../pkg/properties/MockProperties.go 					jukebox-app/pkg/properties Properties
//go:generate mockgen -package transaction  -destination ../pkg/transaction/MockRelationalTransactionHandler.go	jukebox-app/pkg/transaction RelationalTransactionHandler

//go:generate mockgen -package repository -destination ../internal/core/repository/MockArtistRepository.go	jukebox-app/internal/core/repository ArtistRepository
//go:generate mockgen -package repository -destination ../internal/core/repository/MockSongRepository.go	jukebox-app/internal/core/repository SongRepository
//go:generate mockgen -package repository -destination ../internal/core/repository/MockUserRepository.go	jukebox-app/internal/core/repository UserRepository
