package config

import (
	"log"
	logger "sbit-processor/infrastructure/log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func LoadConfig() {
	readConfig()        // read config file
	logger.InitLogger() // init logger

}

func readConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("reading config error: ", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config updated, automatically reloading services....")
	})
}
