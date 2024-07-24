package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// RSA encryption
func Encrypt(pub *rsa.PublicKey, msg []byte) (string, error) {
	label := []byte("")
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, label)
	if err != nil {
		return "", err
	}

	base64String := base64.StdEncoding.EncodeToString(ciphertext)

	return base64String, nil
}

func Decrypt(priv *rsa.PrivateKey, encryptedReq string) ([]byte, error) {
	label := []byte("")
	hash := sha256.New()

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedReq)
	if err != nil {
		fmt.Println("Error decoding Base64 string:", err)
		return nil, err
	}

	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, label)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
