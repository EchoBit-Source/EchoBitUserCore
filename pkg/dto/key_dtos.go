package dto

type SignedPreKeyDto struct {
	Key       string `json:"key"`
	Signature string `json:"signature"`
}

type OneTimePreKeyDto struct {
	Key string `json:"key"`
}
