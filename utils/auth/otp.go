package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/pquerna/otp/totp"
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT. Makmur Subur Jaya"
const CONFIG_AUTH_EMAIL = "ihyamars@students.amikom.ac.id"
const CONFIG_AUTH_PASSWORD = "striker123546"

func GenerateOTP(email string) (otpCode string, err error) {

	now := time.Now()
	expirationTime := now.Add(3 * time.Minute)

	key, err := totp.Generate(totp.GenerateOpts{
		AccountName: email,
	})
	if err != nil {
		return "", err
	}

	otpCode, err = totp.GenerateCode(key.Secret(), expirationTime)
	if err != nil {
		return "", err
	}

	return otpCode, nil
}

func SendEmail(email, otp string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetAddressHeader("CC", email, CONFIG_SENDER_NAME)
	mailer.SetHeader("Subject", "OTP CODE")
	mailer.SetBody("text/html", fmt.Sprintf("OTP Code: <b>%s</b>", otp))

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")

	return nil
}
