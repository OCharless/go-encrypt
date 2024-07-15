package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"go-encrypt/utils"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	envFile := true
	check_map := map[bool]string{true: "✅", false: "❌"}
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		envFile = false
	}

	// Get DATA and SALT from environment
	data, salt := os.Getenv("DATA"), os.Getenv("SALT")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----- go-encrypt (AES-256) -----")
	fmt.Print("Encode Y/n: ")
	action, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading action input: %v", err)
	}
	action = strings.TrimSpace(action)

	fmt.Printf("\033[1A\033[K")
	fmt.Println("Encode:", check_map[action == "Y"])

	if data != "" && salt != "" && envFile {

		fmt.Print("We noticed you have a .env file with DATA and SALT variables loaded, use them? Y/n :")
		use, err := reader.ReadString('\n')
		fmt.Printf("\033[1A\033[K")
		fmt.Println("We noticed you have a .env file with DATA and SALT variables loaded, use them?", check_map[use == "Y"])

		if err != nil {
			log.Fatal("Error reading user input")
		}
		use = strings.TrimSpace(use)
		if use != "Y" {
			data, salt, err = utils.GetArgs(reader, action)
			if err != nil {
				log.Fatal(err)
			}
		}

	} else {
		data, salt, err = utils.GetArgs(reader, action)
		if err != nil {
			log.Fatal(err)
		}

	}

	key := []byte(salt)
	switch action {
	case "Y":
		// Encryption path
		encryptedText, err := utils.EncryptString(data, key)
		if err != nil {
			log.Fatalf("Error encrypting data: %v", err)
		}
		fmt.Printf("Encrypted data to bytes: %v\n", encryptedText)
		encoded := hex.EncodeToString(encryptedText)
		fmt.Printf("Encoded data from bytes to hex: \033[32m%s\033[0m\n", encoded)
	default:
		// Decryption path
		hexDecoded, err := hex.DecodeString(data)
		if err != nil {
			log.Fatalf("Can't decode hex string to bytes: %v", err)
		}
		decryptedData, err := utils.DecryptString(hexDecoded, key)
		if err != nil {
			log.Fatalf("Error decrypting data: %v", err)
		}
		fmt.Printf("Decrypted data: \033[32m%s\033[0m\n", decryptedData)
	}
}
