package token

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/ahmetsabri/go-auth/pkg/helpers"
)

type Token struct {
	ID        uint
	Token     string
	UserId    uint
	ExpiredAt time.Time
}

func init() {
	var token Token = Token{}
	helpers.DB.AutoMigrate(&token)
}

func Create(t *Token) {
	helpers.DB.Create(t)
}

func GenerateToken(email string) string {

	now := strconv.Itoa(int(time.Now().Unix()))
	token := sha256.Sum256([]byte(now + email))
	return hex.EncodeToString(token[:])
}

func Verify(t string) Token {
	var token Token
	now := time.Now().String()
	helpers.DB.First(&token, "token=? AND expired_at > ?", t, now)

	return token
}
