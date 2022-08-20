package encryption

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type hashingConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	hashLen uint32
}

// GetPasswordHash is used to generate a new password hash for storing and comparing at a later date.
//
// Note: When using this hash fetcher, keep in mind that for the same password it will differ based on the time
// it was initiated at, due to the nature of argon2.IDKey func.
func GetPasswordHash(password string) string {
	c := &hashingConfig{
		time:    1,
		memory:  64 * 1024, // nolint:gomnd // allowed magic number
		threads: 4,         // nolint:gomnd // allowed magic number
		hashLen: 32,        // nolint:gomnd // allowed magic numberd
	}

	// Generate a Salt
	salt := make([]byte, 16) // nolint:gomnd // allowed magic number

	if _, err := rand.Read(salt); err != nil {
		// TODO: Return an error instead of panic
		panic(fmt.Errorf("failed to generate salt seed: %w", err))
	}

	hash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.hashLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)
}

// CheckPassword is used to compare a user-inputted password to a hash to see
// weather the password matches.
func CheckPassword(hash, password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}

	if hash == "" {
		return errors.New("hash cannot be empty")
	}

	parts := strings.Split(hash, "$")

	if len(parts) != 6 { // nolint:gomnd // allowed magic number; should match with number of arguments from GetPasswordHash() result
		return errors.New("hash has invalid format")
	}

	var config hashingConfig

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &config.memory, &config.time, &config.threads)
	if err != nil {
		return fmt.Errorf("failed to scan: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return fmt.Errorf("failed to decode parts: %w", err)
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return fmt.Errorf("failed to decode parts: %w", err)
	}

	config.hashLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, config.time, config.memory, config.threads, config.hashLen)

	const compareSuccess = 1

	if subtle.ConstantTimeCompare(decodedHash, comparisonHash) != compareSuccess {
		return errors.New("hash does not match password")
	}

	return nil
}
