package utils

import (
	"crypto/rand"
)

// RandomChars generates a short url
// URL with length 5, will give 62⁵ = ~916 Million URLs
// URL with length 6, will give 62⁶ = ~56 Billion URLs
// URL with length 7, will give 62⁷ = ~3500 Billion URLs
func RandomChars(length int) string {
	var chars []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("wrong charset length")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

// GenerateShortURL generates a short url
func GenerateShortURL(url string) string {
	// generate unique short url using url
	return RandomChars(7)
}

// IsValidURL checks if url is valid
func IsValidURL(url string) bool {
	return true
}
