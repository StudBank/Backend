package etc

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func initConfig() {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	logger.Debug().Msg("Reading configuration file at ./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Err(err).Msg("Error reading the config file")
		return
	}

}
