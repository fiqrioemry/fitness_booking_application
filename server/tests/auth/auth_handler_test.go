package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"server/internal/dto"
	"server/internal/handlers"
	"server/tests/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(req *dto.RegisterRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func TestRegisterHandler(t *testing.T) {
	mockSvc := new(mocks.MockAuthService)
	handler := handlers.NewAuthHandler(mockSvc)

	r := gin.Default()
	r.POST("/register", handler.Register)

	body := dto.RegisterRequest{
		Email:    "test@example.com",
		Fullname: "Test User",
		Password: "password",
	}
	jsonValue, _ := json.Marshal(body)

	mockSvc.On("Register", mock.Anything).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockSvc.AssertExpectations(t)
}
