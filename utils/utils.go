package utils

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"strings"
)

func EncryptString(text string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func DecryptString(ciphertext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func GetArgs(reader *bufio.Reader, action string) (string, string, error) {
	switch action {
	case "Y":
		fmt.Print("Enter text to encode: ")
	default:
		fmt.Print("Enter text to decode: ")
	}

	data, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("error reading text input for data: %w", err)
	}
	data = strings.TrimSpace(data)

	fmt.Print("Enter salt: ")
	salt, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("error reading text input for salt: %w", err)
	}
	salt = strings.TrimSpace(salt)
	if len(salt) != 32 {
		return "", "", errors.New("salt must be 32 characters long")
	}

	return data, salt, nil
}
