package user

import (
	"log"
	"time"

	"github.com/ahmetsabri/go-auth/models/token"
	"github.com/ahmetsabri/go-auth/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `gorm:"type:varchar(100);NOT NULL"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;NOT NULL"`
	Password  string    `gorm:"type:varchar(100);NOT NULL"`
	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoCreateTime:true"`
}

func init() {
	var u User = User{}
	helpers.Migrate(&u)
}

func Create(u *User) string {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 2)

	if err != nil {
		log.Fatal(err)
	}

	u.Password = string(hashedPass)
	res := helpers.DB.Create(u)

	if res.Error != nil {
		return ""
	}

	t := &token.Token{
		UserId:    u.ID,
		Token:     token.GenerateToken(u.Email),
		ExpiredAt: time.Now().Add(1 * time.Hour),
	}

	token.Create(t)

	return t.Token
}

func CheckPassword(hashed string, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}
