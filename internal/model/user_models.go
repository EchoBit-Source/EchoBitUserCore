package model

import (
	"time"
)

type UserModel struct {
	ID             string               `json:"id"`
	Username       string               `json:"username"`
	PasswordHash   string               `json:"passwordHash"`
	PublicKey      string               `json:"publicKey"`
	SignedPreKey   SignedPreKeyModel    `json:"signedPreKey"`
	OneTimePreKeys []OneTimePreKeyModel `json:"oneTimePreKeys"`
	CreatedAt      time.Time            `json:"createdAt"`
	UpdatedAt      time.Time            `json:"updatedAt"`
}
