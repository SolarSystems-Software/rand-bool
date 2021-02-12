package randbool

import (
	"crypto/rand"
	"encoding/binary"
)

// Generates a cryptographically random unsigned 64-bit integer.
func CryptoRand() (uint64, error) {
	var bytes [8]byte

	if _, err := rand.Read(bytes[:]); err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(bytes[:]), nil
}

// Same as CryptoRand but panics if it encounters an error.
func CryptoRandOrPanic() uint64 {
	i, err := CryptoRand()
	if err != nil {
		panic(err)
	}
	return i
}
