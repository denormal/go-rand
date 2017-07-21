package rand

import (
	"encoding/base64"
)

// Base64 returns a base 64 encoded string of random bytes, as generated
// by crypto/rand, or an error if sufficient random bits cannot be obtained.
// The string encoding is performed using encoding/base64.RawStdEncoding.
func Base64(bytes int) (string, error) {
	_bytes, _err := Bytes(bytes)
	if _err != nil {
		return "", _err
	}
	defer ReturnBytes(_bytes)

	return base64.RawStdEncoding.EncodeToString(_bytes), nil
}
