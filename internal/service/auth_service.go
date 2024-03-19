package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GenerateHash(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
	GenerateTokens(userID string) (string, string, error)
	ValidateToken(token string) (*TokenClaims, error)
}

type TokenClaims struct {
	UserID string `json:"userID"`
	Exp    int64  `json:"exp"`
}

type AuthServiceImpl struct {
	secretKey []byte
}

func NewAuthServiceImpl(secretKey []byte) *AuthServiceImpl {
	return &AuthServiceImpl{secretKey: secretKey}
}

func (s *AuthServiceImpl) GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *AuthServiceImpl) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *AuthServiceImpl) ValidateToken(token string) (*TokenClaims, error) {
	tkn, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tkn.Claims.(*jwt.MapClaims); ok && tkn.Valid {
		userID, ok := (*claims)["userID"].(string)
		if !ok {
			return nil, nil
		}

		exp, ok := (*claims)["exp"].(float64)
		if !ok {
			return nil, nil
		}

		return &TokenClaims{
			UserID: userID,
			Exp:    int64(exp),
		}, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func (s *AuthServiceImpl) GenerateTokens(userID string) (string, string, error) {
	accessToken, err := generateToken(userID, s.secretKey, 15*time.Minute)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateToken(userID, s.secretKey, 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateToken(userID string, secretKey []byte, expiry time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiry)

	jwtClaims := jwt.MapClaims{
		"userID": userID,
		"exp":    expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}
