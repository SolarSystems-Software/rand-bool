package randbool

import (
	"crypto/rand"
	"encoding/binary"
)

// CryptoRand generates a cryptographically secure unsigned 64-bit integer.
// The error returned is non-nil if an error occurred.
func CryptoRand() (uint64, error) {
	var bytes [8]byte

	if _, err := rand.Read(bytes[:]); err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(bytes[:]), nil
}

// CryptoRandOrPanic calls CryptoRand, but panics if err != nil.
func CryptoRandOrPanic() uint64 {
	i, err := CryptoRand()
	if err != nil {
		panic(err)
	}
	return i
}
