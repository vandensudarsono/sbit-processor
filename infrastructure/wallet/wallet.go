package wallet

import (
	"context"
	"fmt"
	logging "sbit-processor/infrastructure/log"
	"sbit-processor/internal/topicinit"
	wc "sbit-processor/internal/wallet_codec"

	"github.com/lovoo/goka"
	"github.com/spf13/viper"
)

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	cd "sbit-processor/adapter/balance/codec"
// 	bi "sbit-processor/adapter/balance/input"
// 	logging "sbit-processor/infrastructure/log"

// 	"github.com/lovoo/goka"
// 	"github.com/lovoo/goka/codec"
// 	"github.com/spf13/viper"
// )

// var (
// 	balance *goka.Processor
// 	tmc     *goka.TopicManagerConfig
// )

// // func init() {
// // 	tmc = goka.NewTopicManagerConfig()

// // 	tmc.Table.Replication = 1
// // 	tmc.Stream.Replication = 1
// // }

// // func InitBalance() {
// // 	var (
// // 		err error
// // 	)

// // 	//define group
// // 	group := goka.Group(viper.GetString("processor.balance.group"))
// // 	topic := goka.Stream(viper.GetString("brocker.topic"))

// // 	b := bi.NewBalanceInput()
// // 	g := goka.DefineGroup(
// // 		group,
// // 		goka.Input(topic, new(codec.String), b.BalanceInputCB),
// // 		goka.Persist(new(cd.DepositCodec)),
// // 	)
// // 	balance, err = goka.NewProcessor(
// // 		viper.GetStringSlice("brocker.url"),
// // 		g,
// // 		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
// // 		goka.WithConsumerGroupBuilder(goka.DefaultConsumerGroupBuilder),
// // 	)

// // 	if err != nil {
// // 		log.Fatalf("error creating processor: %v", err)
// // 	}
// // }

// func GetBalanceProcessor() *goka.Processor {
// 	return balance
// }

// func RunBallanceProcessor() {
// 	var (
// 		err error
// 	)

// 	tmc = goka.NewTopicManagerConfig()

// 	tmc.Table.Replication = 1
// 	tmc.Stream.Replication = 1

// 	//define group
// 	group := goka.Group(viper.GetString("processor.balance.group"))
// 	topic := goka.Stream(viper.GetString("brocker.topic"))

// 	b := bi.NewBalanceInput()
// 	g := goka.DefineGroup(
// 		group,
// 		goka.Input(topic, new(codec.String), b.BalanceInputCB),
// 		goka.Persist(new(cd.DepositCodec)),
// 	)
// 	balance, err = goka.NewProcessor(
// 		viper.GetStringSlice("brocker.url"),
// 		g,
// 		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
// 		goka.WithConsumerGroupBuilder(goka.DefaultConsumerGroupBuilder),
// 	)

// 	if err != nil {
// 		log.Fatalf("error creating processor: %v", err)
// 	}

// 	err = balance.Run(context.Background())
// 	if err != nil {
// 		logging.Errorf("error running balance processor: %v", err)
// 	}

// 	fmt.Println("Runing Balance Processor.. ")
// }

var (
	group        goka.Group
	Table        goka.Table
	WalletStream goka.Stream
	broker       []string
)

func PrepareTopics() {
	stream := viper.GetString("processor.wallet.topic")
	group = goka.Group(viper.GetString("processor.wallet.group"))
	Table = goka.GroupTable(group)

	if stream == "" {
		logging.WithFields(logging.Fields{"component": "prepare topic", "action": "create wallet topic"}).
			Infof("wallet processor stream undfined. stream = %v", stream)
		//useing default topc
		stream = "wallet"
	}

	WalletStream = goka.Stream(stream)
	broker = []string{fmt.Sprintf("%s:%s", viper.GetString("processor.broker.host"), viper.GetString("processor.broker.port"))}

	topicinit.EnsureStreamExists(stream, broker)
}

func Run(ctx context.Context, cb goka.ProcessCallback) func() error {
	return func() error {
		g := goka.DefineGroup(group,
			goka.Input(WalletStream, new(wc.WalletCodec), cb),
			goka.Persist(new(wc.WalletListCodec)),
		)
		p, err := goka.NewProcessor(broker, g)
		if err != nil {
			return err
		}
		return p.Run(ctx)
	}
}
