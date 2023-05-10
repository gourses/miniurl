package miniurl

import (
	"crypto/md5"
	"encoding/hex"
)

// Returns Hashed output of long input
func Hash(input string) string {
	hash := md5.Sum([]byte(input))

	return hex.EncodeToString(hash[:])
}