package response

import (
	"encoding/base64"
	"log"

	"github.com/azizyogo/bank-app/common/encryption"
	"github.com/azizyogo/bank-app/domain/user"

	jsoniter "github.com/json-iterator/go"
)

func EncryptResponse(userKey user.UserEntity, v interface{}) (string, error) {

	// Convert pem to rsa public key
	pubKey, err := encryption.DecodePEMToPublicKey(userKey.PublicKey)
	if err != nil {
		return "", err
	}

	// Encode struct to JSON
	byt, err := jsoniter.MarshalToString(v)
	if err != nil {
		log.Fatalf("Error marshalling struct to JSON: %v", err)
		return "", err
	}

	// Encode JSON string to Base64
	encodedBytes := base64.StdEncoding.EncodeToString([]byte(byt))

	// Decode from Base64
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedBytes)
	if err != nil {
		log.Fatalf("Error decoding Base64 string: %v", err)
		return "", err
	}

	// Encrypt
	encryptRes, err := encryption.Encrypt(pubKey, decodedBytes)
	if err != nil {
		return "", err
	}

	return encryptRes, nil
}
