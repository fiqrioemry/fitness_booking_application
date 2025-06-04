package auth_test

import (
	"errors"
	"testing"

	"server/internal/dto"
	"server/internal/models"
	"server/internal/services"
	"server/tests/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthRepo struct {
	mock.Mock
}

func (m *MockAuthRepo) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockAuthRepo) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func TestRegisterSuccess(t *testing.T) {
	mockRepo := new(mocks.MockAuthRepository)
	service := services.NewAuthService(mockRepo)

	mockRepo.On("GetUserByEmail", "test@example.com").Return(nil, errors.New("not found"))

	req := &dto.RegisterRequest{
		Email:    "test@example.com",
		Fullname: "Test User",
		Password: "password",
	}

	err := service.Register(req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
