package generate

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOTP() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	otp := r.Intn(10000) % 10000
	return fmt.Sprintf("%04d", otp)
}