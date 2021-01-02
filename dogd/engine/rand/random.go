package rand

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"math/rand"
)

const keySize = 32

// NewRandom creates a rand.Rand from a cryptographically-secure pseudo-random-generator (CSPRNG)
func NewRandom(seed []byte) (*rand.Rand, error) {
	source, err := NewRandomSource(seed)
	if err != nil {
		return nil, err
	}

	return rand.New(source), nil
}

// NewRandomSource creates a rand.Source from a cryptographically-secure pseudo-random-generator (CSPRNG)
func NewRandomSource(seed []byte) (rand.Source, error) {

	if len(seed) == 0 {
		return nil, fmt.Errorf("seed length %d is too short", len(seed))
	}

	key, iv := hashSeed(seed)

	stream, err := initializeAesCtr(key, iv)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize AES256-CTR for CSPRNG: %w", err)
	}

	return newStreamCipherSource64(stream), nil
}

// hashSeed generates uniformly distributed entropy from potentially non-uniform entropy and returns a key and
// IV appropriate for AES256-CTR
func hashSeed(seed []byte) ([]byte, []byte) {

	hashed := sha512.Sum512(seed)

	key := hashed[:keySize]
	iv := hashed[keySize : keySize+aes.BlockSize]
	return key, iv
}

func initializeAesCtr(key []byte, iv []byte) (cipher.Stream, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return cipher.NewCTR(block, iv), nil
}

func newStreamCipherSource64(stream cipher.Stream) rand.Source64 {
	return &streamCipherSource64{
		stream: stream,
	}
}

type streamCipherSource64 struct {
	stream cipher.Stream
}

func (s *streamCipherSource64) Int63() int64 {
	for {
		r := int64(s.Uint64())
		if r >= 0 {
			return r
		}
	}
}

func (s *streamCipherSource64) Seed(seed int64) {
	panic("re-seeding CSPRNG not supported")
}

func (s *streamCipherSource64) Uint64() uint64 {

	buf := make([]byte, 8)
	zeroBytes := make([]byte, 8)

	// 8 bytes of 0 plaintext will yield the raw CSPRNG stream
	s.stream.XORKeyStream(buf, zeroBytes)

	return binary.BigEndian.Uint64(buf)
}
