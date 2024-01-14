package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func md5Encrypt(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
