package encryption

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptData(data string) string {
	e := sha256.New()
	e.Write([]byte(data))
	return hex.EncodeToString(e.Sum(nil))
}
