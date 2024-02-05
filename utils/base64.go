package utils

import (
	"encoding/base64"
)

func Base64Decode(encodedString string) (string, error) {
	// Decode the Base64 string
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return "", err
	}

	// Convert the decoded bytes to a string
	decodedString := string(decodedBytes)

	return decodedString, nil
}

func Base64Encode[T []byte | string](s T) (string, error) {
	encodedBytes := []byte(s)
	encodedString := base64.StdEncoding.EncodeToString(encodedBytes)

	return encodedString, nil
}

func Base64DecodeFast(s string) string {
	s, err := Base64Decode(s)
	if err != nil {
		panic(err)
	}
	return s
}
