package rand

import (
	"crypto/rand"
	"encoding/binary"
	"sync"
)

var (
	mu     sync.RWMutex
	endian binary.ByteOrder
	length = 8
	pool   = sync.Pool{
		New: func() interface{} { return make([]byte, 0, length) },
	}
)

// BigEndian causes a sequence of bytes to be interpreted in big endian byte
// order when retrieving integers from random bytes.
func BigEndian() { setorder(binary.BigEndian) }

// LittleEndian causes a sequence of bytes to be interpreted in little endian
// byte order when retrieving integers from random bytes. This is the default
// byte order.
func LittleEndian() { setorder(binary.LittleEndian) }

// IsBigEndian() returns true if byte sequences will be interpreted in
// big endian byte order when retrieving integers from random bytes.
func IsBigEndian() bool { return binary.BigEndian == byteorder() }

// IsLitteEndian() returns true if byte sequences will be interpreted in
// big little byte order when retrieving integers from random bytes.
func IsLittleEndian() bool { return binary.LittleEndian == byteorder() }

// Bytes returns a slice of random bytes, or an error if crypto/rand is
// unable to read the requested number of bytes. Bytes returns memory from
// a pool of slices, which may be returned to the pool via ReturnBytes().
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
		ReturnBytes(_bytes)
		return nil, _err
	}

	return _bytes[:bytes], nil
}

// ReturnBytes returns a slice of bytes to the memory pool used by Bytes().
func ReturnBytes(bytes []byte) {
	if cap(bytes) != 0 {
		pool.Put(bytes[:0])
	}
}

//
// private methods
//

func setorder(b binary.ByteOrder) {
	mu.Lock()
	endian = b
	mu.Unlock()
}

func byteorder() binary.ByteOrder {
	mu.RLock()
	_endian := endian
	mu.RUnlock()

	return _endian
}

func init() {
	LittleEndian()
}
