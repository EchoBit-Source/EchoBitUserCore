package dto

import (
	"time"

	"github.com/EchoBit-Source/EchoBitUserCore/internal/model"
)

type UserDto struct {
	ID             string             `json:"id"`
	Username       string             `json:"username"`
	PasswordHash   string             `json:"passwordHash"`
	PublicKey      string             `json:"publicKey"`
	SignedPreKey   SignedPreKeyDto    `json:"signedPreKey"`
	OneTimePreKeys []OneTimePreKeyDto `json:"oneTimePreKeys"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
}

type PublicUserDto struct {
	Username  string `json:"username"`
	PublicKey string `json:"publicKey"`
}

type CreateUserDto struct {
	Username       string             `json:"username"`
	Password       string             `json:"password"`
	PublicKey      string             `json:"publicKey"`
	SignedPreKey   SignedPreKeyDto    `json:"signedPreKey"`
	OneTimePreKeys []OneTimePreKeyDto `json:"oneTimePreKeys"`
}

func (c *CreateUserDto) ToUser(hashedPassword string) *model.UserModel {
	return &model.UserModel{
		Username:     c.Username,
		PasswordHash: hashedPassword,
		PublicKey:    c.PublicKey,
		SignedPreKey: model.SignedPreKeyModel(c.SignedPreKey),
		OneTimePreKeys: func() (otks []model.OneTimePreKeyModel) {
			for _, otk := range c.OneTimePreKeys {
				otks = append(otks, model.OneTimePreKeyModel(otk))
			}
			return otks
		}(),
	}
}
