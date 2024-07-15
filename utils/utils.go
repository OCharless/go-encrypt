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
	if action == "Y" {
		action = "encode"
	} else {
		action = "decode"
	}
	fmt.Printf("Enter text to %s: ", action)

	data, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("error reading text input for data: %w", err)
	}
	data = strings.TrimSpace(data)
	fmt.Printf("\033[1A\033[K")
	fmt.Printf("Enter text to %s: ", action)
	fmt.Printf("\033[36m%s\033[0m\n", data)

	fmt.Print("Enter salt: ")
	salt, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("error reading text input for salt: %w", err)
	}
	salt = strings.TrimSpace(salt)

	fmt.Printf("\033[1A\033[K")
	fmt.Print("Enter salt: ")
	fmt.Printf("\033[36m%s\033[0m\n", strings.Repeat("*", len(salt)))

	if len(salt) != 32 {
		return "", "", errors.New("salt must be 32 characters long")
	}

	return data, salt, nil
}
