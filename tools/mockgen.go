package tools

//go:generate mockgen -package mocks -destination ../tests/mocks/MockDBDataSource.go 				jukebox-app/pkg/datasource DBDataSource

//go:generate mockgen -package mocks -destination ../tests/mocks/MockEnvironment.go 				jukebox-app/pkg/environment Environment

//go:generate mockgen -package mocks -destination ../tests/mocks/MockPropertySource.go 				jukebox-app/pkg/properties PropertySource
//go:generate mockgen -package mocks -destination ../tests/mocks/MockProperties.go 					jukebox-app/pkg/properties Properties

//go:generate mockgen -package mocks -destination ../tests/mocks/MockDBTransactionHandler.go 		jukebox-app/pkg/transaction DBTransactionHandler
