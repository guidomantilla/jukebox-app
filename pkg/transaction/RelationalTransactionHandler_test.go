package transaction

import (
	"jukebox-app/tests/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NewRelationalTransactionHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	relationalDataSource := mocks.NewMockRelationalDataSource(ctrl)
	handler := NewRelationalTransactionHandler(relationalDataSource)

	assert.NotNil(t, handler)
}
