package cipherer

import (
	"encoding/base64"
	"errors"
)

func Cipher(rawString, secret string) (string, error) {
	if len(secret) == 0 {
		return "", errors.New("Secret key cannot be empty!")
	}

	encryptedBytes, err := process([]byte(rawString), []byte(secret))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

func Decipher(cipheredText, secret string) (string, error) {
	if len(secret) == 0 {
		return "", errors.New("Secret key cannot be empty!")
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(cipheredText)
	if err != nil {
		return "", errors.New("Failed to decode Base64 data!")

	}

	decryptedBytes, err := process(cipherBytes, []byte(secret))
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil

}
func process(input, secret []byte) ([]byte, error) {
	// if len(secret) == 0 {
	// 	fmt.Println("Secret key cannot be empty. Exiting now...")
	// 	os.Exit(1)
	// }

	for i, b := range input {
		input[i] = b ^ secret[i%len(secret)] // 0...len(secret)
	}
	return input, nil

}
