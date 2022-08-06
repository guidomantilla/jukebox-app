package tools

//go:generate mockgen -package datasource -destination ../tests/mocks/datasource/MockDBDataSource.go 				jukebox-app/pkg/datasource DBDataSource

//go:generate mockgen -package environment -destination ../tests/mocks/environment/MockEnvironment.go 				jukebox-app/pkg/environment Environment

//go:generate mockgen -package properties -destination ../tests/mocks/properties/MockPropertySource.go 				jukebox-app/pkg/properties PropertySource
//go:generate mockgen -package properties -destination ../tests/mocks/properties/MockProperties.go 					jukebox-app/pkg/properties Properties

//go:generate mockgen -package transaction -destination ../tests/mocks/transaction/MockDBTransactionHandler.go 		jukebox-app/pkg/transaction DBTransactionHandler
