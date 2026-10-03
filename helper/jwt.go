package helper

import (
	"time"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, role string, cfg config.Config) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.JWTSecret))
}