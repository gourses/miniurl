// Package is provides ..
package miniurl

import (
	"crypto/md5"
	"encoding/hex"
)

var i int

//go doc vi generoida dokumenttia lokaalisti
// Hash generates 32 byte long, vaatii funktion nimi, ilman siitä sisäinen dokumentointi ei tomi
func Hash(input string) string {
	i++
	if i%10000 == 0 {
		return "sSS"
	}

	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
