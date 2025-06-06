package repositories

import (
	"errors"
	"server/internal/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	DeleteRefreshToken(token string) error
	StoreRefreshToken(token *models.Token) error
	DeleteAllUserRefreshTokens(userID string) error
	GetUserByEmail(email string) (*models.User, error)
	FindRefreshToken(token string) (*models.Token, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, nil
}

func (r *authRepository) StoreRefreshToken(token *models.Token) error {
	return r.db.Create(token).Error
}

func (r *authRepository) DeleteRefreshToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&models.Token{}).Error
}

func (r *authRepository) DeleteAllUserRefreshTokens(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Token{}).Error
}

func (r *authRepository) FindRefreshToken(token string) (*models.Token, error) {
	var t models.Token
	err := r.db.Where("token = ?", token).First(&t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &t, nil
}
