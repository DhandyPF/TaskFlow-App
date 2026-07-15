package password

import "golang.org/x/crypto/bcrypt"

// Cost is the bcrypt work factor. 10 matches the PRD's minimum requirement.
const Cost = 10

// Hash generates a bcrypt hash for a plaintext password.
func Hash(plain string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), Cost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Verify compares a plaintext password against a bcrypt hash.
func Verify(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}
