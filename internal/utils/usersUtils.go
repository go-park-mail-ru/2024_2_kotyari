package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

const (
	timeCost    uint32 = 1         // Время обработки (количество итераций)
	memoryCost  uint32 = 64 * 1024 // Память, используемая Argon2 (в KB)
	parallelism uint8  = 4         // Количество параллельных потоков
	keyLength   uint32 = 32        // Длина генерируемого ключа
	saltLength  int    = 16        // Длина соли в байтах
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

func HashPassword(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, parallelism, keyLength)
	fullHash := append(salt, hash...)

	return base64.RawStdEncoding.EncodeToString(fullHash)
}

// Разделение соли и хеша
func splitSaltAndHash(saltHashBase64 string) ([]byte, []byte, error) {
	saltHash, err := base64.RawStdEncoding.DecodeString(saltHashBase64)
	if err != nil {
		return nil, nil, err
	}

	salt := saltHash[:saltLength] // Первые saltLength байт — это соль
	hash := saltHash[saltLength:] // Остальное — это хеш

	return salt, hash, nil
}

func VerifyPassword(storedSaltHashBase64, password string) bool {
	salt, storedHash, err := splitSaltAndHash(storedSaltHashBase64)
	if err != nil {
		return false
	}

	computedHash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, parallelism, keyLength)

	return string(storedHash) == string(computedHash)
}
