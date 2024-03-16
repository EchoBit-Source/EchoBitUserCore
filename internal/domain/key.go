package domain

type SignedPreKey struct {
	Key       string `json:"key"`
	Signature string `json:"signature"`
}

type OneTimePreKey struct {
	Key string `json:"key"`
}
