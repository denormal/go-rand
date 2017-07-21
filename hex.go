package rand

import (
	"encoding/hex"
)

// Hex returns a hex encoded string of random bytes, as generated
// by crypto/rand, or an error if sufficient random bits cannot be obtained.
// The string encoding is performed using encoding/hex.
func Hex(bytes int) (string, error) {
	_bytes, _err := Bytes(bytes)
	if _err != nil {
		return "", _err
	}
	defer ReturnBytes(_bytes)

	return hex.EncodeToString(_bytes), nil
}
