package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
)

// This function encrypts the Registry of the document
// Parameters:
// plainStr (string): this is the registry of the document that will be encrypt
// keyStr (string): this is a secret string that serve for encrypt the registry
// Returns:
// A string with the registry encrypted and an error.
func EncryptRegistry(plainStr, keyStr string) (string, error) {
	plainText := []byte(plainStr)
	plainKey := []byte(keyStr)

	// Create the AES cipher
	block, err := aes.NewCipher(plainKey)
	if err != nil {
		log.Fatalf("It was a problem when the AES cipher was created: %v", err)
		return "", nil
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	// Slice of first 16 bytes
	iv := cipherText[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatalf("There is a problem writting the 16 rand bytes in the fill iv: %v", err)
		return "", err
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return string(cipherText), nil
}

// This function decrypt the registry of the document
// Parameters:
// encryptRegistry (string): string of the registry encrypted
// keyStr (string): this is a secret string that serve for desencrypt the registry
// Returns:
// the string decrypt and error
func DecryptRegistry(encryptString, keyString string) (string, error) {
	ciphertext := []byte(encryptString)
	key := []byte(keyString)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error to created the AES cipher: %v", err)
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		log.Fatalf("Text is too short")
		return "", errors.New("ext is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}
