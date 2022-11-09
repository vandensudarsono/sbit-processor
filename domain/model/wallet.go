package model

type Wallet struct {
	WalletID int64   `json:"wallet_id"`
	Amount   float32 `json:"amount"`
}
