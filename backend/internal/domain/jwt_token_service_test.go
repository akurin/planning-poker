package domain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"github.com/stretchr/testify/assert"
	"strings"
	"time"

	//"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Issue_token(t *testing.T) {
	// Arrange
	privateKeyBytes := []byte(`-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgevZzL1gdAFr88hb2
OF/2NxApJCzGCEDdfSp6VQO30hyhRANCAAQRWz+jn65BtOMvdyHKcvjBeBSDZH2r
1RTwjmYSi9R/zpBnuQ4EiMnCqfMPWiZqB4QdbAd0E7oH50VpuZ1P087G
-----END PRIVATE KEY-----`)
	privateKey := parsePrivateKey(t, privateKeyBytes)

	playerId := NewFakePlayerId("some-player-id")
	clock := NewFakeClock(WithNow(time.Unix(1634482913, 0)))
	sut := NewJwtTokenService(privateKey.(*ecdsa.PrivateKey), clock, time.Second)

	// Act
	token, _ := sut.IssueToken(playerId)

	tokenWithoutSignature := dropSignature(token)

	// Assert

	// Expected:
	// {
	//   "exp": 1634482913,
	//   "iat": 1634482912,
	//   "sub": "some-player-id"
	// }
	assert.Equal(t,
		"eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzQ0ODI5MTQsImlhdCI6MTYzNDQ4MjkxMywic3ViIjoic29tZS1wbGF5ZXItaWQifQ",
		tokenWithoutSignature)
}

func parsePrivateKey(t *testing.T, privateKeyBytes []byte) interface{} {
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	require.NoError(t, err)
	return privateKey
}

func dropSignature(token string) string {
	lastDotIndex := strings.LastIndex(token, ".")
	return token[:lastDotIndex]
}

func Test_Invalid_key(t *testing.T) {
	// Arrange
	privateKey, _ := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	playerId := NewFakePlayerId("some-player-id")
	clock := NewFakeClock()
	sut := NewJwtTokenService(privateKey, clock, time.Second)

	// Act
	token, err := sut.IssueToken(playerId)

	// Assert
	assert.Equal(t, "", token)
	assert.NotNil(t, err)
}
