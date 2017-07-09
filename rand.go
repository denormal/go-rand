package rand

import (
	"crypto/rand"
	"encoding/binary"
	"sync"
	"sync/atomic"
)

var (
	endian atomic.Value
	length = 8
	pool   = sync.Pool{
		New: func() interface{} { return make([]byte, 0, length) },
	}
)

func BigEndian()           { endian.Store(binary.BigEndian) }
func LittleEndian()        { endian.Store(binary.LittleEndian) }
func IsBigEndian() bool    { return binary.BigEndian == byteorder() }
func IsLittleEndian() bool { return binary.LittleEndian == byteorder() }

func Bytes(bytes int) ([]byte, error) {
	if bytes <= 0 {
		return []byte{}, nil
	}

	_bytes := pool.Get().([]byte)
	if cap(_bytes) < bytes {
		_bytes = make([]byte, 0, bytes)
	}

	_, _err := rand.Read(_bytes[:bytes])
	if _err != nil {
		Put(_bytes)
		return nil, _err
	}

	return _bytes[:bytes], nil
}

func Put(bytes []byte) {
	if cap(bytes) != 0 {
		pool.Put(bytes[:0])
	}
}

//
// private methods
//

func byteorder() binary.ByteOrder {
	return endian.Load().(binary.ByteOrder)
}

func init() {
	LittleEndian()
}
