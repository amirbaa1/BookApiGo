package Utils

import (
	"Auth/Config"
	"Auth/Model"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"regexp"
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

	minLength, maxLength := 4, 16

	if len(password) < minLength {
		return "", errors.New("password must be at least 8 characters")
	}
	if len(password) > maxLength {
		return "", errors.New("password must be less than 16 characters")
	}
	passUp, _ := regexp.MatchString(`[A-Z]`, password)
	passLow, _ := regexp.MatchString(`[a-z]`, password)
	hasNumber, _ := regexp.MatchString(`[0-9]`, password)

	if !passUp {
		//return fmt.Errorf("password must contain at least one uppercase letter")
		return "", errors.New("password must contain at least one uppercase letter")
	}
	if !passLow {
		return "", errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return "", errors.New("password must contain at least one number")
	}

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
