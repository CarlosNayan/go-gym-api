package services

import (
	"api-gym-on-go/src/modules/auth/repository"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthRepository *repository.UserRepository
	JWTKey         []byte
}

func NewAuthService(authRepo *repository.UserRepository, jwtKey []byte) *AuthService {
	return &AuthService{AuthRepository: authRepo, JWTKey: jwtKey}
}

const (
	errInvalidCredentials = "invalid credentials"
	errTokenGeneration    = "failed to generate token"
)

func generateToken(userID string, role string, duration time.Duration, jwtKey []byte) (string, error) {
	expirationTime := time.Now().Add(duration)

	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func (s *AuthService) Auth(email string, password string) (map[string]string, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.AuthRepository.FindByEmail(email)
	if err != nil {
		log.Printf("Failed to fetch user %s: %v", email, err)
		return nil, err
	}

	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return nil, errors.New(errInvalidCredentials)
	}

	tokenString, err := generateToken(user.ID, string(user.Role), 15*time.Minute, s.JWTKey)
	if err != nil {
		log.Printf("Failed to generate access token for user %s: %v", email, err)
		return nil, errors.New(errTokenGeneration)
	}

	refreshTokenString, err := generateToken(user.ID, string(user.Role), 72*time.Hour, s.JWTKey)
	if err != nil {
		log.Printf("Failed to generate refresh token for user %s: %v", email, err)
		return nil, errors.New(errTokenGeneration)
	}

	return map[string]string{
		"token":         tokenString,
		"refresh_token": refreshTokenString,
	}, nil
}
