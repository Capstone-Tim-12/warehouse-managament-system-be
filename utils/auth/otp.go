package auth

import (
	"time"

	"github.com/pquerna/otp/totp"
)

func GenerateOTP(email string) (otpCode string, err error) {

	now := time.Now()

	key, err := totp.Generate(totp.GenerateOpts{
		AccountName: email,
		Digits:      4,
		Period:      180,
	})
	if err != nil {
		return "", err
	}

	otpCode, err = totp.GenerateCode(key.Secret(), now)
	if err != nil {
		return "", err
	}

	return otpCode, nil
}
