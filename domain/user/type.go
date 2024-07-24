package user

import "time"

type (
	UserEntity struct {
		ID         int       `json:"id" db:"id"`
		Username   string    `json:"username" db:"username"`
		Password   string    `json:"password" db:"password"`
		Role       int       `json:"role" db:"role"`
		PublicKey  string    `json:"public_key" db:"public_key"`
		PrivateKey string    `json:"private_key" db:"private_key"`
		CreatedAt  time.Time `json:"created_at" db:"created_at"`
	}

	UserKey struct {
		ID         int    `json:"id" db:"id"`
		PublicKey  string `json:"public_key" db:"public_key"`
		PrivateKey string `json:"private_key" db:"private_key"`
	}
)
