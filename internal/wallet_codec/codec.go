package walletcodec

import (
	"encoding/json"
	"sbit-processor/domain/model"
)

type WalletCodec struct{}

func (c *WalletCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *WalletCodec) Decode(data []byte) (interface{}, error) {
	var m model.Wallet
	return &m, json.Unmarshal(data, &m)
}

type WalletListCodec struct{}

func (c *WalletListCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *WalletListCodec) Decode(data []byte) (interface{}, error) {
	var m []model.Wallet
	err := json.Unmarshal(data, &m)
	return m, err
}
