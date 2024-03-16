package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Username      string             `json:"username" bson:"username"`
	PasswordHash  string             `json:"passwordHash" bson:"passwordHash"`
	PublicKey     string             `json:"publicKey" bson:"publicKey"`
	SignedPreKey  SignedPreKey       `json:"signedPreKey" bson:"signedPreKey"`
	OneTimePreKey OneTimePreKey      `json:"oneTimePreKey" bson:"oneTimePreKey"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type SignedPreKey struct {
	Key       string `json:"key" bson:"key"`
	Signature string `json:"signature" bson:"signature"`
}

type OneTimePreKey struct {
	Key string `json:"key" bson:"key"`
}
