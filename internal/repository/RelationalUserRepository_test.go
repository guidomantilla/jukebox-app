package repository

import (
	"jukebox-app/internal/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RelationalUserRepository_Create_Ok(t *testing.T) {

	user := &model.User{
		Id:    -1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositorySaveFunction(t, repository.statementCreate, user, repository.Create, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_RelationalUserRepository_Create_Err(t *testing.T) {

	user := &model.User{
		Id:    -1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositorySaveFunction(t, repository.statementCreate, user, repository.Create, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(-1), user.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalUserRepository_Update_Ok(t *testing.T) {

	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositorySaveFunction(t, repository.statementUpdate, user, repository.Update, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_RelationalUserRepository_Update_Err(t *testing.T) {

	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositorySaveFunction(t, repository.statementUpdate, user, repository.Update, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(1), user.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalUserRepository_DeleteById_Ok(t *testing.T) {

	user := &model.User{
		Id: 1,
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositoryDeleteFunction(t, repository.statementDelete, user.Id, repository.DeleteById, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_RelationalUserRepository_DeleteById_Err(t *testing.T) {

	user := &model.User{
		Id: 1,
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositoryDeleteFunction(t, repository.statementDelete, user.Id, repository.DeleteById, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(1), user.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalUserRepository_FindAll_Ok(t *testing.T) {

	repository := NewRelationalUserRepository()
	users, err := CallRelationalUserRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, users)
}

func Test_RelationalUserRepository_FindAll_Query_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	users, err := CallRelationalUserRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, users)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalUserRepository_FindAll_Scan_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	users, err := CallRelationalUserRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, users)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalUserRepository_FindById_Ok(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalUserRepository_FindById_Query_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalUserRepository_FindById_Scan_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalUserRepository_FindByCode_Ok(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalUserRepository_FindByCode_Query_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalUserRepository_FindByCode_Scan_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalUserRepository_FindByName_Ok(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalUserRepository_FindByName_Query_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalUserRepository_FindByName_Scan_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalUserRepository_FindByEmail_Ok(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByEmail, "some_email", repository.FindByEmail, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalUserRepository_FindByEmail_Query_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByEmail, "some_email", repository.FindByEmail, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalUserRepository_FindByEmail_Scan_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByEmail, "some_email", repository.FindByEmail, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}
