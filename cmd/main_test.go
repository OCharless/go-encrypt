package main_test

import (
	"encoding/hex"
	"go-encrypt/utils"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMainEncod(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name           string
		Text           string
		IntendedResult string
	}{
		{
			Name:           "test1",
			Text:           "collar boat plane",
			IntendedResult: "0000000000000000000000000000000024f3069ed01cfd1e078de92b23a26116c8",
		},
		{
			Name:           "test2",
			Text:           "this algorithm allows you to encrypt a string with the AES-256 encryption protocol.",
			IntendedResult: "0000000000000000000000000000000033f40381910fb11b079ef47f3ba32019f5d371a77c58e90541e85dc437440be21d4e47a976cda191d42448adb4adf16701d41d8e3d8bf82462782d90794003bdb47c95ebfeb6ce927806e228095f187ac7c1d5",
		},
	}
	err := godotenv.Load("../.env.example")
	if err != nil {
		t.Error(err)
	}
	SALT := os.Getenv("SALT")
	if err != nil {
		t.Error(err)
	}
	key := []byte(SALT) // AES-256 requires a 32-byte key

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			encryptedText, err := utils.EncryptString(tc.Text, key)
			if err != nil {
				t.Error(SALT)
			}
			if hex.EncodeToString(encryptedText) != tc.IntendedResult {
				t.Error("Error in encrypting")
			}
		})
	}
}

func TestMainDecod(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name           string
		Text           string
		IntendedResult string
	}{
		{
			Name:           "test1",
			Text:           "0000000000000000000000000000000029f90999dd0fbe19",
			IntendedResult: "necklace",
		},
		{
			Name:           "test2",
			Text:           "0000000000000000000000000000000033f40381910fb11b079ef47f3ba32019f5d371a77c58e90541e85dc4374500e2826739921826c67a8028c415faf75a729aeb1b293d75df96c9d1d4bfd301eb9090cac0c670b10a085623b737bb95",
			IntendedResult: "this algorithm allows you to decrypt an hex string if you have the key for it.",
		},
	}
	err := godotenv.Load("../.env.example")
	if err != nil {
		t.Error(err)
	}
	SALT := os.Getenv("SALT")
	key := []byte(SALT) // AES-256 requires a 32-byte key

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			hexDecoded, err := hex.DecodeString(tc.Text)
			if err != nil {
				t.Error("Can't encode hex string")
			}
			decryptedText, err := utils.DecryptString(hexDecoded, key)
			if err != nil {
				t.Error(err)
			}
			if decryptedText != tc.IntendedResult {
				t.Error("Error in decrypting")
			}
		})
	}
}
