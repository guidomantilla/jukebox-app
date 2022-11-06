package repository

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"jukebox-app/internal/model"
)

func Test_RelationalArtistRepository_Create_Ok(t *testing.T) {

	artist := &model.Artist{
		Id:   -1,
		Code: 1,
		Name: "2",
	}

	repository := NewRelationalArtistRepository()
	err := CallRelationalArtistRepositorySaveFunction(t, repository.statementCreate, artist, repository.Create, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), artist.Id)
}

func Test_RelationalArtistRepository_Create_Err(t *testing.T) {

	artist := &model.Artist{
		Id:   -1,
		Code: 1,
		Name: "2",
	}

	repository := NewRelationalArtistRepository()
	err := CallRelationalArtistRepositorySaveFunction(t, repository.statementCreate, artist, repository.Create, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(-1), artist.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalArtistRepository_Update_Ok(t *testing.T) {

	artist := &model.Artist{
		Id:   1,
		Code: 2,
		Name: "3",
	}

	repository := NewRelationalArtistRepository()
	err := CallRelationalArtistRepositorySaveFunction(t, repository.statementUpdate, artist, repository.Update, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), artist.Id)
}

func Test_RelationalArtistRepository_Update_Err(t *testing.T) {

	artist := &model.Artist{
		Id:   1,
		Code: 2,
		Name: "3",
	}

	repository := NewRelationalArtistRepository()
	err := CallRelationalArtistRepositorySaveFunction(t, repository.statementUpdate, artist, repository.Update, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(1), artist.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalArtistRepository_DeleteById_Ok(t *testing.T) {

	Artist := &model.Artist{
		Id: 1,
	}

	repository := NewRelationalArtistRepository()
	err := CallRelationalArtistRepositoryDeleteFunction(t, repository.statementDelete, Artist.Id, repository.DeleteById, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), Artist.Id)
}

func Test_RelationalArtistRepository_DeleteById_Err(t *testing.T) {

	Artist := &model.Artist{
		Id: 1,
	}

	repository := NewRelationalArtistRepository()
	err := CallRelationalArtistRepositoryDeleteFunction(t, repository.statementDelete, Artist.Id, repository.DeleteById, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(1), Artist.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalArtistRepository_FindAll_Ok(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artists, err := CallRelationalArtistRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, Artists)
}

func Test_RelationalArtistRepository_FindAll_Query_Err(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artists, err := CallRelationalArtistRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, Artists)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalArtistRepository_FindAll_Scan_Err(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artists, err := CallRelationalArtistRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, Artists)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalArtistRepository_FindById_Ok(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artist, err := CallRelationalArtistRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, Artist)
}

func Test_RelationalArtistRepository_FindById_Query_Err(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artist, err := CallRelationalArtistRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, Artist)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalArtistRepository_FindById_Scan_Err(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artist, err := CallRelationalArtistRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, Artist)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalArtistRepository_FindByCode_Ok(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artist, err := CallRelationalArtistRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, Artist)
}

func Test_RelationalArtistRepository_FindByCode_Query_Err(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artist, err := CallRelationalArtistRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, Artist)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalArtistRepository_FindByCode_Scan_Err(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artist, err := CallRelationalArtistRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, Artist)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalArtistRepository_FindByName_Ok(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artist, err := CallRelationalArtistRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, Artist)
}

func Test_RelationalArtistRepository_FindByName_Query_Err(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artist, err := CallRelationalArtistRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, Artist)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalArtistRepository_FindByName_Scan_Err(t *testing.T) {

	repository := NewRelationalArtistRepository()
	Artist, err := CallRelationalArtistRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, Artist)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}
