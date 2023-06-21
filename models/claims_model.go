package models

import (
	"Application/services"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"strconv"
	"time"
)

var (
	jwtSigningKey      = []byte("secret_key")
	expireTimeDuration = time.Duration(9999)
)

type Claims struct {
	jwt.RegisteredClaims
	ClaimID  int    `json:"claim_id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Contact  int64  `json:"contact,omitempty"`
	Type     string `json:"type,omitempty"`
}

func (claim *Claims) GenerateToken() (string, error) {
	claim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Hour * expireTimeDuration),
		},
	}
	claim.ClaimID = services.GetRandomNumber(16)
	for len(strconv.Itoa(claim.ClaimID)) < 16 {
		claim.ClaimID = services.GetRandomNumber(16)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	return token.SignedString(jwtSigningKey)
}

func (claim *Claims) Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func (claim *Claims) Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
