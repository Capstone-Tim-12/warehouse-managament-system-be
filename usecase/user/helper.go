package user

import (
	"time"

	utilsModel "github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(userId int, userRole string) string {
	var payloadParser utilsModel.JwtCustomClaims

	payloadParser.UserId = userId
	payloadParser.UserRole = userRole
	payloadParser.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 60))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payloadParser)
	t, _ := token.SignedString([]byte("1234"))
	return t
}

func HashPassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(result)
}

func ComparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
