package services

import (
	"api-gym-on-go/src/modules/auth/repository"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthRepository *repository.UserRepository
}

func NewAuthService(authRepo *repository.UserRepository) *AuthService {
	return &AuthService{AuthRepository: authRepo}
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func (s *AuthService) Auth(email string, password string) (map[string]string, error) {
	user, err := s.AuthRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if error != nil {
		return nil, errors.New("invalid credentials")
	}

	// access_token
	expirationTokenTime := time.Now().Add(15 * time.Minute)

	claims := &jwt.MapClaims{
		"sub":  user.ID,
		"role": string(user.Role),
		"exp":  expirationTokenTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// refresh_token
	expirationRefreshTokenTime := time.Now().Add(72 * time.Hour)

	refresh_claims := &jwt.MapClaims{
		"sub":  user.ID,
		"role": string(user.Role),
		"exp":  expirationRefreshTokenTime.Unix(),
	}

	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)

	refreshTokenString, err := refresh_token.SignedString(jwtKey)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return map[string]string{
		"token":         tokenString,
		"refresh_token": refreshTokenString,
	}, nil
}
