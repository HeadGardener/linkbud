package authservice

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/HeadGardener/linkbud/internal/app/repository"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	salt      = "iewimecqiweobweobv"
	secretKey = "nwignbebvnoiwnvibweoivinu"
)

type AuthService struct {
	repos *repository.Repository
}

func NewAuthService(repos *repository.Repository) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) Create(user models.User) (int, error) {
	user.Password = getPasswordHash(user.Password)

	return s.repos.Authorization.Create(user)
}

func (s *AuthService) GenerateToken(userInput models.UserInput) (string, error) {
	userInput.Password = getPasswordHash(userInput.Password)

	id, err := s.repos.Authorization.IfUserExist(userInput)
	if err != nil {
		return "", errors.New("user not exists")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(15 * time.Minute).Unix()
	claims["authorized"] = true
	claims["userID"] = id

	return token.SignedString([]byte(secretKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, _ := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(secretKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && claims.VerifyExpiresAt(time.Now().Unix(), false) {
		userID := claims["userID"]
		return int(userID.(float64)), nil
	}

	return 0, errors.New("unable to extract claims")
}

func (s *AuthService) CheckUsername(username string) (int, error) {
	return s.repos.Authorization.IfUserExistByUN(username)
}

func getPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
