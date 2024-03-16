package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mongoUser struct {
	ID            primitive.ObjectID   `bson:"_id"`
	Username      string               `bson:"username"`
	PasswordHash  string               `bson:"passwordHash"`
	PublicKey     string               `bson:"publicKey"`
	SignedPreKey  MongoSignedPreKey    `bson:"signedPreKey"`
	OneTimePreKey []MongoOneTimePreKey `bson:"oneTimePreKeys"`
	CreatedAt     time.Time            `bson:"createdAt"`
	UpdatedAt     time.Time            `bson:"updatedAt"`
}

type MongoSignedPreKey struct {
	Key       string `bson:"key"`
	Signature string `bson:"signature"`
}

type MongoOneTimePreKey struct {
	Key string `bson:"key"`
}
