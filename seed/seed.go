package seed

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
)

type Seed []byte

const CHARSET = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const SEED_LEN = 8

func GenerateSeed() Seed {
	now := time.Now().UnixNano()
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(now))
	sha := sha256.Sum256(buf)
	return toCharset(sha[:SEED_LEN])
}

func toCharset(chars []byte) []byte {
	for i, char := range chars {
		chars[i] = CHARSET[int(char)%len(CHARSET)]
	}
	return chars
}

func (seed Seed) To32byte() [32]byte {
	result := [32]byte{}
	copy(result[:], seed)
	return result
}

func (seed Seed) Isvalid() bool {
	if len(seed) != SEED_LEN {
		return false
	}
	for _, char := range seed {
		if !isInCharset(char) {
			return false
		}
	}
	return true
}

func isInCharset(char byte) bool {
	for _, c := range CHARSET {
		if char == byte(c) {
			return true
		}
	}
	return false
}
