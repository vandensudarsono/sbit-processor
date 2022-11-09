package usecase

import (
	"context"
	"fmt"
	"sbit-processor/domain/model"
)

type WalletInteractor struct{}

func NewWalletInteractor() *WalletInteractor {
	return &WalletInteractor{}
}

// Sum: sum current wallet with new deposit
func (w *WalletInteractor) Sum(ctx context.Context, wallet, deposit *model.Wallet) (model.Wallet, error) {
	// validate wallet id
	var result model.Wallet
	if wallet != nil && deposit != nil {
		if wallet.WalletID != deposit.WalletID {
			return result, fmt.Errorf("wallet id and deposit id is invalid")
		}

		result = model.Wallet{
			WalletID: wallet.WalletID,
			Amount:   wallet.Amount + deposit.Amount,
		}
	} else if wallet == nil && deposit != nil {
		result = model.Wallet{
			WalletID: result.WalletID,
			Amount:   deposit.Amount,
		}
	}

	return result, nil
}
