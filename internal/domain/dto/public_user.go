package dto

type PublicUser struct {
	Username      string `json:"username"`
	PublicKey     string `json:"publicKey"`
	SignedPreKey  string `json:"signedPreKey"`
	OneTimePreKey string `json:"oneTimePreKey"`
}
