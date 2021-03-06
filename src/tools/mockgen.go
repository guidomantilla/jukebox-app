package tools

//go:generate mockgen -package datasource -destination ../tests/mocks/misc/datasource/MockDBDataSource.go 				jukebox-app/src/misc/datasource DBDataSource

//go:generate mockgen -package environment -destination ../tests/mocks/misc/environment/MockEnvironment.go 				jukebox-app/src/misc/environment Environment

//go:generate mockgen -package properties -destination ../tests/mocks/misc/properties/MockPropertySource.go 			jukebox-app/src/misc/properties PropertySource
//go:generate mockgen -package properties -destination ../tests/mocks/misc/properties/MockProperties.go 				jukebox-app/src/misc/properties Properties

//go:generate mockgen -package transaction -destination ../tests/mocks/misc/transaction/MockDBTransactionHandler.go 	jukebox-app/src/misc/transaction DBTransactionHandler
