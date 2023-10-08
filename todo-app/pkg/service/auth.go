package service

import (
	"crypto/sha1"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"lincoln.boris/todo"
	"lincoln.boris/todo/pkg/repository"
)

const (
	salt = "3478thw3iufu395g3qb48qbfh34u3bgibugi3422i9433qjhfb"
	signingKey = "8rjng29ngw4ung103gn9g5kj4wgn4295n2485n425g4jw51-g]4giw3g329"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%s", hash.Sum([]byte(salt)))
}
