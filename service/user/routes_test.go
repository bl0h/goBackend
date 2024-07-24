package user

import (
	"testing"

	"github.com/bl0h/goBackend/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
