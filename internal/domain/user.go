package domain

import (
	"time"
)

type User struct {
	ID            string          `json:"id"`
	Username      string          `json:"username"`
	PasswordHash  string          `json:"passwordHash"`
	PublicKey     string          `json:"publicKey"`
	SignedPreKey  SignedPreKey    `json:"signedPreKey"`
	OneTimePreKey []OneTimePreKey `json:"oneTimePreKeys"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
}
