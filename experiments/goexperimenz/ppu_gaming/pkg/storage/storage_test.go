package storage

import (
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := testifyAssert.New(t)
	storage, err := NewInMemStorage()
	defer storage.Close()
	assert.NoError(err)
	user, err := storage.User.Create("test-name", "test@test.de", "qqqq")
	assert.NoError(err)
	assert.NotNil(user)
	assert.NotNil(user.ID)
	assert.NotEqual(user.PasswordHash, "qqqq")
}

func TestCreateUserNonUniqueEmail(t *testing.T) {
	assert := testifyAssert.New(t)
	storage, err := NewInMemStorage()
	defer storage.Close()
	assert.NoError(err)
	user, err := storage.User.Create("test-name", "test@test.de", "qqqq")
	assert.NoError(err)
	assert.NotNil(user)
	assert.NotNil(user.ID)
	assert.NotEqual(user.PasswordHash, "qqqq")
	user2, err := storage.User.Create("test-name", "test@test.de", "qqqq")
	assert.Error(err)
	assert.Nil(user2)
}
