package tools

//go:generate mockgen -package mocks -destination ../tests/mocks/MockRelationalDataSource.go 			jukebox-app/pkg/datasource RelationalDataSource
//go:generate mockgen -package mocks -destination ../tests/mocks/MockEnvironment.go 					jukebox-app/pkg/environment Environment
//go:generate mockgen -package mocks -destination ../tests/mocks/MockPropertySource.go 					jukebox-app/pkg/properties PropertySource
//go:generate mockgen -package mocks -destination ../tests/mocks/MockProperties.go 						jukebox-app/pkg/properties Properties
//go:generate mockgen -package mocks -destination ../tests/mocks/MockRelationalTransactionHandler.go	jukebox-app/pkg/transaction RelationalTransactionHandler

//go:generate mockgen -package mocks -destination ../tests/mocks/MockArtistRepository.go	jukebox-app/internal/core/repository ArtistRepository
//go:generate mockgen -package mocks -destination ../tests/mocks/MockSongRepository.go		jukebox-app/internal/core/repository SongRepository
//go:generate mockgen -package mocks -destination ../tests/mocks/MockUserRepository.go		jukebox-app/internal/core/repository UserRepository
