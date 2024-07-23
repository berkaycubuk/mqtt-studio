package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/berkaycubuk/mqtt-studio/types"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		e := echo.New()

		payload := types.RegisterUserPayload{
			FullName: "Test User",
			Email: "invalid",
			Password: "123",
		}

		marshalled, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewBuffer(marshalled))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rr := httptest.NewRecorder()
		c := e.NewContext(req, rr)

		if assert.Error(t, handler.handleRegister(c)) {
			assert.Equal(t, http.StatusBadRequest, rr.Code)
		}

		if assert.NoError(t, handler.handleRegister(c)) {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
			//assert.Equal(t, http.StatusCreated, rr.Code)
		}
	})

	t.Run("should register the user", func(t *testing.T) {
		e := echo.New()

		payload := types.RegisterUserPayload{
			FullName: "Test User",
			Email: "test@test.test",
			Password: "password",
		}

		marshalled, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewBuffer(marshalled))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rr := httptest.NewRecorder()
		c := e.NewContext(req, rr)

		if assert.NoError(t, handler.handleRegister(c)) {
			assert.Equal(t, http.StatusCreated, rr.Code)
		}
	})
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
