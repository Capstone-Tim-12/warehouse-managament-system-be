package generate

import (
	"os"
	"strings"
)

func GenerateEmailOTP(name, otp string) (emailBody string, err error) {
	configPath := "./template/otp.html"
	tempByte, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	emailBody = string(tempByte)
	emailBody = strings.ReplaceAll(emailBody, "[Nama Pengguna]", name)
    emailBody = strings.ReplaceAll(emailBody, "[Kode OTP]", otp)
	return
}