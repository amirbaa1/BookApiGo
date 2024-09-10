package Utils

import (
	"Auth/Config"
	"Auth/Model"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GeneratorToken(user Model.Auth) (string, error) {
	//claims := jwt.MapClaims{}
	//claims["user_id"] = user.Id
	//claims["user_name"] = user.UserName
	//claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	claims := jwt.MapClaims{
		"user-id":   user.Id,
		"user_name": user.UserName,
		"exp":       time.Now().Add(time.Minute * 15).Unix(),
		"Role":      "User",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Config.Secret)
}

func GeneratorPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func ValidatePassword(userPassword, hashPassword string) error {
	vPass := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(userPassword))
	if vPass != nil {
		return errors.New("invalid password")
	}
	return nil
}
