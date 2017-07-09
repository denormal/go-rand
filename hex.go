package rand

import (
	"encoding/hex"
)

func Hex(bytes int) (string, error) {
	_bytes, _err := Bytes(bytes)
	if _err != nil {
		return "", _err
	}
	defer Put(_bytes)

	return hex.EncodeToString(_bytes), nil
}
