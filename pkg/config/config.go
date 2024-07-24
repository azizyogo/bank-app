package config

import (
	"crypto/rsa"
	"github.com/azizyogo/bank-app/common/encryption"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser        string
	DBPassword    string
	DBName        string
	DBHost        string
	DBPort        string
	JWTSecret     string
	RSAKeyPublic  *rsa.PublicKey
	RSAKeyPrivate *rsa.PrivateKey
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	publicKeyPEM := viper.GetString("rsa-key.public")
	privateKeyPEM := viper.GetString("rsa-key.secret")

	publicKey, err := encryption.DecodePEMToPublicKey(publicKeyPEM)
	if err != nil {
		log.Fatalf("Failed to decode public key: %v", err)
	}

	privateKey, err := encryption.DecodePEMToPrivateKey(privateKeyPEM)
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
	}

	return Config{
		DBUser:        viper.GetString("db.user"),
		DBPassword:    viper.GetString("db.password"),
		DBName:        viper.GetString("db.name"),
		DBHost:        viper.GetString("db.host"),
		DBPort:        viper.GetString("db.port"),
		JWTSecret:     viper.GetString("jwt.secret"),
		RSAKeyPublic:  publicKey,
		RSAKeyPrivate: privateKey,
	}
}
