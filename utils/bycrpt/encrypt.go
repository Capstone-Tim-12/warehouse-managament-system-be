package bycrpt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(key, data string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	hmacResult := h.Sum(nil)
	return hex.EncodeToString(hmacResult)
}
