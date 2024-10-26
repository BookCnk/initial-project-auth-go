package repository

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/entity"
)

type EncryptorRepository interface {
	GetPassphrase() (entity.Encryptor, error)
	GeneratePassphrase(length int) (string, error)
	Encrypt(plaintext string) []byte
	HashPassword(salt string, password string) string
	ComparePassword(storedHashedPassword string, salt string, inputPassword string) bool
	Decrypt(ciphertext []byte) string
}

type encryptorRepository struct {
	db     *gorm.DB
	config config.AppConfig
}

func (r *encryptorRepository) HashPassword(salt string, password string) string {
	combined := salt + password

	hashes := sha256.New()
	hashes.Write([]byte(combined))

	hashesBytes := hashes.Sum(nil)
	hashedPassword := hex.EncodeToString(hashesBytes)

	return hashedPassword
}

func (r *encryptorRepository) DecryptField(ciphertext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, encryptedMessage := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, encryptedMessage, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func (r *encryptorRepository) ComparePassword(storedHashedPassword string, salt string, inputPassword string) bool {
	inputHashPassword := r.HashPassword(salt, inputPassword)
	return storedHashedPassword == inputHashPassword
}

func (r *encryptorRepository) storePassphrase() (entity.Encryptor, error) {
	var encryptor entity.Encryptor
	passphrase, err := r.GeneratePassphrase(32)
	encryptor.Hash = []byte(passphrase)

	if err != nil {
		return entity.Encryptor{}, err
	}

	result := r.db.Create(&encryptor)

	if result.Error != nil {
		return entity.Encryptor{}, result.Error
	}

	return encryptor, nil
}

func (r *encryptorRepository) GetPassphrase() (entity.Encryptor, error) {
	var encryptor entity.Encryptor
	result := r.db.Last(&encryptor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			encryptor, err := r.storePassphrase()
			if err != nil {
				return entity.Encryptor{}, err
			}
			return encryptor, nil
		}
		return entity.Encryptor{}, result.Error
	}
	return encryptor, nil
}

func (r *encryptorRepository) GeneratePassphrase(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytePassphrase := make([]byte, length)

	_, err := rand.Read(bytePassphrase)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		bytePassphrase[i] = charset[bytePassphrase[i]%byte(len(charset))]
	}

	return string(bytePassphrase), nil
}

func (r *encryptorRepository) Encrypt(plaintext string) []byte {
	secretKey, err := r.GetPassphrase()
	if err != nil {
		panic(err)
	}

	newCipher, err := aes.NewCipher(secretKey.Hash)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(newCipher)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return ciphertext
}

func (r *encryptorRepository) Decrypt(ciphertext []byte) string {
	secretKey, err := r.GetPassphrase()
	if err != nil {
		panic(err)
	}

	aes, err := aes.NewCipher(secretKey.Hash)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}

func ProvideEncryptorRepository(db *gorm.DB, config config.AppConfig) EncryptorRepository {
	return &encryptorRepository{
		db:     db,
		config: config,
	}
}
