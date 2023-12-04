package generate

import (
	"math/rand"
	"strconv"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func FormatRupiah(nominal int) string {
	// Mengonversi nilai nominal ke string
	nominalStr := strconv.Itoa(nominal)

	// Menambahkan separator ribuan
	length := len(nominalStr)
	separatorIndex := length % 3
	if separatorIndex == 0 {
		separatorIndex = 3
	}
	formattedNominal := nominalStr[:separatorIndex]

	for i := separatorIndex; i < length; i += 3 {
		formattedNominal += "." + nominalStr[i:i+3]
	}

	return formattedNominal
}
