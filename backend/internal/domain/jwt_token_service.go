package domain

import (
	"crypto/ecdsa"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtTokenService struct {
	privateKey *ecdsa.PrivateKey
	clock      Clock
	tokenTTL   time.Duration
}

func NewJwtTokenService(privateKey *ecdsa.PrivateKey, clock Clock, tokenTTL time.Duration) JwtTokenService {
	return JwtTokenService{
		privateKey: privateKey,
		clock:      clock,
		tokenTTL:   tokenTTL,
	}
}

func (s JwtTokenService) IssueToken(playerId PlayerId) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   playerId.String(),
		IssuedAt:  s.clock.Now().Unix(),
		ExpiresAt: s.clock.Now().Add(s.tokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedString, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}
