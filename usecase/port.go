package usecase

import (
	"context"
	"sbit-processor/domain/model"
)

type UsecaseInput interface {
	Sum(ctx context.Context, wallet, deposit *model.Wallet) (model.Wallet, error)
}
