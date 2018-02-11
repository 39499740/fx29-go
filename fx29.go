package fx29

import (
	"encoding/base64"
	"net/url"
	"strings"
)

// Encode returns a URL-safe string from bytes using Base64 without padding and url.QueryEscape, with a optional xor parameter used for obfuscation.
func Encode(bytes []byte, xor []byte) string {
	dest := bytes
	if xor != nil {
		dest = xorBytes(dest, xor)
	}
	return url.QueryEscape(base64EncodeStripped(dest))
}

// Decode returns a byte array from a string using url.QueryUnescape and Base64 without padding, with a optional xor parameter used for obfuscation.
func Decode(s string, xor []byte) ([]byte, error) {
	decodedQuery, err := url.QueryUnescape(s)
	if err != nil {
		return nil, err
	}

	decodedBase64, err := base64DecodeStripped(decodedQuery)
	if err != nil {
		return nil, err
	}

	dest := decodedBase64
	if xor != nil {
		dest = xorBytes(dest, xor)
	}
	return dest, nil
}

func xorBytes(src []byte, xor []byte) []byte {
	j := 0
	for i := range src {
		src[i] = src[i] ^ xor[j]
		j++

		if j >= len(xor) {
			j = 0
		}
	}
	return src
}

/// https://stackoverflow.com/questions/31971614/base64-encode-decode-without-padding-on-golang-appengine
func base64EncodeStripped(bytes []byte) string {
	encoded := base64.StdEncoding.EncodeToString(bytes)
	return strings.TrimRight(encoded, "=")
}

func base64DecodeStripped(s string) ([]byte, error) {
	if i := len(s) % 4; i != 0 {
		s += strings.Repeat("=", 4-i)
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	return decoded, err
}
