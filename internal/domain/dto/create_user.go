package dto

type CreateUser struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	PublicKey     string `json:"public_key"`
	SignedPreKey  string `json:"signed_pre_key"`
	OneTimePreKey string `json:"one_time_pre_key"`
}
