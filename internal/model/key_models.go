package model

type SignedPreKeyModel struct {
	Key       string `json:"key"`
	Signature string `json:"signature"`
}

type OneTimePreKeyModel struct {
	Key string `json:"key"`
}
