package generate

import "math/rand"

func GenerateExternalId(length int) string {
	const charset = "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return "DP-" + string(result)
}
