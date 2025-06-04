package auth_test

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "file::memory:?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&models.User{}, &models.Token{}))
	return db
}

func TestCreateAndGetUser(t *testing.T) {
	db := setupTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	authRepo := repositories.NewAuthRepository(db)

	user := &models.User{
		ID:       uuid.New(),
		Email:    "test@example.com",
		Password: "hashedpw",
		Fullname: "Test User",
	}

	err := userRepo.CreateUser(user)
	require.NoError(t, err)

	fetched, err := authRepo.GetUserByEmail("test@example.com")
	require.NoError(t, err)
	assert.Equal(t, user.Email, fetched.Email)
}

func TestStoreAndFindToken(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewAuthRepository(db)

	token := &models.Token{
		ID:        uuid.New(),
		UserID:    uuid.New(),
		Token:     "refresh-token",
		ExpiredAt: time.Now().Add(time.Hour),
	}

	err := repo.StoreRefreshToken(token)
	require.NoError(t, err)

	found, err := repo.FindRefreshToken("refresh-token")
	require.NoError(t, err)
	assert.Equal(t, token.Token, found.Token)
}
