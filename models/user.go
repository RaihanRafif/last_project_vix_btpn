package models

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model

	Username  string `gorm:"varchar(255);not null"`
	Email     string `gorm:"column:email;unique_index"`
	Password  string `gorm:"column:password;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password should not be empty")
	}
	bytePassword := []byte(password)

	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

func (u *User) IsValidPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (user *User) GenerateJwtToken() string {
	jwt_token := jwt.New(jwt.SigningMethodHS512)

	jwt_token.Claims = jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 90).Unix(),
	}
	token, _ := jwt_token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token
}
