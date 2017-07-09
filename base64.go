package rand

import (
	"encoding/base64"
)

func Base64(bytes int) (string, error) {
	_bytes, _err := Bytes(bytes)
	if _err != nil {
		return "", _err
	}
	defer Put(_bytes)

	return base64.RawStdEncoding.EncodeToString(_bytes), nil
}
