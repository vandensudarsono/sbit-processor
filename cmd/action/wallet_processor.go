package action

import (
	"context"
	"log"
	"os"
	"os/signal"
	wCB "sbit-processor/adapter/wallet_callback"
	logging "sbit-processor/infrastructure/log"
	"sbit-processor/infrastructure/wallet"
	"sbit-processor/usecase"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func RunWalletProcessor() {

	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)

	walletUC := usecase.NewWalletInteractor()
	walletHanlder := wCB.NewWalletCB(walletUC)

	//prepare topic first
	wallet.PrepareTopics()

	grp.Go(wallet.Run(ctx, walletHanlder.DepositRequest))

	logging.WithFields(logging.Fields{"component": "procesor", "action": "wallet processor"}).
		Infof("Running wallet processor")

	// Wait for SIGINT/SIGTERM
	waiter := make(chan os.Signal, 1)
	signal.Notify(waiter, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-waiter:
	case <-ctx.Done():
	}
	cancel()
	if err := grp.Wait(); err != nil {
		log.Println(err)
	}
	log.Println("done")

}
