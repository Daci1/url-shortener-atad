package security

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	AccessTokenSecret  = []byte("b527697922c00cb7f55f6d417107edd65b7c5fc4849af8e49cf6bc6fe9eea6b5")
	RefreshTokenSecret = []byte("f87fac9f44e3b4a5a77181a7d375e9aadcfc2826b4bca0d08a3f5e823429f1d7")
)

type TokenPair struct {
	Token        string
	RefreshToken string
}

func GenerateTokens(userId, username string) (*TokenPair, error) {
	// --- Access Token ---
	accessTokenClaims := jwt.MapClaims{
		"sub":      userId,
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(), // expires in 15 mins
		"iat":      time.Now().Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	signedAccessToken, err := accessToken.SignedString(AccessTokenSecret)
	if err != nil {
		return nil, err
	}

	// --- Refresh Token ---
	refreshTokenClaims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(), // expires in 7 days
		"iat":      time.Now().Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	signedRefreshToken, err := refreshToken.SignedString(RefreshTokenSecret)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}
func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Verify that the signing method is what you expect (HS256)
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return AccessTokenSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("token is invalid or expired")
	}

	return token, nil
}
