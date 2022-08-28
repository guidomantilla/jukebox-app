package repository

import (
	"jukebox-app/internal/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RelationalSongRepository_Create_Ok(t *testing.T) {

	song := &model.Song{
		Id:       -1,
		Code:     1,
		Name:     "2",
		ArtistId: 3,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositorySaveFunction(t, repository.statementCreate, song, repository.Create, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), song.Id)
}

func Test_RelationalSongRepository_Create_Err(t *testing.T) {

	song := &model.Song{
		Id:       -1,
		Code:     1,
		Name:     "2",
		ArtistId: 3,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositorySaveFunction(t, repository.statementCreate, song, repository.Create, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(-1), song.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalSongRepository_Update_Ok(t *testing.T) {

	song := &model.Song{
		Id:       1,
		Code:     2,
		Name:     "3",
		ArtistId: 4,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositorySaveFunction(t, repository.statementUpdate, song, repository.Update, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), song.Id)
}

func Test_RelationalSongRepository_Update_Err(t *testing.T) {

	song := &model.Song{
		Id:       1,
		Code:     2,
		Name:     "3",
		ArtistId: 4,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositorySaveFunction(t, repository.statementUpdate, song, repository.Update, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(1), song.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalSongRepository_DeleteById_Ok(t *testing.T) {

	user := &model.User{
		Id: 1,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositoryDeleteFunction(t, repository.statementDelete, user.Id, repository.DeleteById, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_RelationalSongRepository_DeleteById_Err(t *testing.T) {

	user := &model.User{
		Id: 1,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositoryDeleteFunction(t, repository.statementDelete, user.Id, repository.DeleteById, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(1), user.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalSongRepository_FindAll_Ok(t *testing.T) {

	repository := NewRelationalSongRepository()
	users, err := CallRelationalSongRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, users)
}

func Test_RelationalSongRepository_FindAll_Query_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	users, err := CallRelationalSongRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, users)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalSongRepository_FindAll_Scan_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	users, err := CallRelationalSongRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, users)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalSongRepository_FindById_Ok(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalSongRepository_FindById_Query_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalSongRepository_FindById_Scan_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalSongRepository_FindByCode_Ok(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalSongRepository_FindByCode_Query_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalSongRepository_FindByCode_Scan_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalSongRepository_FindByName_Ok(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalSongRepository_FindByName_Query_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalSongRepository_FindByName_Scan_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}
