package walletcallback

import (
	"sbit-processor/domain/model"
	logging "sbit-processor/infrastructure/log"
	"sbit-processor/usecase"

	"github.com/lovoo/goka"
)

type WalletCB struct {
	uc usecase.UsecaseInput
}

func NewWalletCB(uc usecase.UsecaseInput) *WalletCB {
	return &WalletCB{uc: uc}
}

func (wcb *WalletCB) DepositRequest(ctx goka.Context, msg interface{}) {
	var (
		wallet  *model.Wallet
		deposit *model.Wallet
	)

	if v := ctx.Value(); v != nil {
		wallet = v.(*model.Wallet)
	}

	deposit = msg.(*model.Wallet)

	if deposit.WalletID != 0 {
		//sum add deposit amount to current wallet
		updateWallet, err := wcb.uc.Sum(ctx.Context(), wallet, deposit)
		if err != nil {
			logging.WithFields(logging.Fields{"component": "wallet callback", "action": "deposit request"}).
				Errorf("error at sum wallet with deposit: %v", err)
		}

		ctx.SetValue(updateWallet)
	}

}
