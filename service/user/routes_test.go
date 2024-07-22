package user

import (
	"testing"

	"github.com/berkaycubuk/mqtt-studio/types"
)

func TestUserServiceHandlers(t *testing.T) {
	/*
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	*/
}

type mockUserStore struct {}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
