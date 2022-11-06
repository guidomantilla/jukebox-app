package tools

//go:generate mockgen -package cachemanager -destination ../pkg/cachemanager/mocks.go	jukebox-app/pkg/cachemanager CacheManager
//go:generate mockgen -package datasource 	-destination ../pkg/datasource/mocks.go		jukebox-app/pkg/datasource RelationalDataSource
//go:generate mockgen -package environment 	-destination ../pkg/environment/mocks.go	jukebox-app/pkg/environment Environment
//go:generate mockgen -package properties 	-destination ../pkg/properties/mocks.go 	jukebox-app/pkg/properties Properties,PropertySource

//go:generate mockgen -package repository -destination ../internal/repository/mocks.go	jukebox-app/internal/repository ArtistRepository,SongRepository,UserRepository
