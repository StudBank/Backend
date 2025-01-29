package etc

import (
	"encoding/hex"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitConfig !! Can exit program
func InitConfig() {
	log := GetPlainLogger("config", logrus.DebugLevel)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	configDefaults()
	log.Debug("Reading configuration file at ./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.WithError(err).Fatal("Error reading the config file")
		return
	}

	err = rewrite()
	if err != nil {
		log.WithError(err).Fatal("Error while rewriting")
	}
}

func configDefaults() {
	viper.SetDefault("name", "bak-default")
	viper.SetDefault("version", "v0.0.0-def")
	viper.SetDefault("env", "dev")

	viper.SetDefault("log.level", "ERROR")
	viper.SetDefault("log.lokiAddress", "https://localhost:3100")
	viper.SetDefault("log.lokiAuth.login", "")
	viper.SetDefault("log.lokiAuth.password", "")

	viper.SetDefault("api.port", 8080)
	viper.SetDefault("api.ssl.on", false)
	viper.SetDefault("api.ssl.cert", "")
	viper.SetDefault("api.ssl.key", "")

	viper.SetDefault("api.middlewares.cors.allowOrigins", []string{"*"})
	viper.SetDefault("api.middlewares.cors.allowMethods", []string{"GET", "PUT", "POST"})
	viper.SetDefault("api.middlewares.cors.allowHeaders", []string{"*"})
	viper.SetDefault("api.middlewares.cors.allowCredentials", true)
	viper.SetDefault("api.middlewares.cors.exposeHeaders", []string{"*"})
	viper.SetDefault("api.middlewares.cors.maxAge", 0)

	viper.SetDefault("api.middlewares.csrf.tokenLength", 32)
	viper.SetDefault("api.middlewares.csrf.tokenLookup", "header:X-CSRF-Token")

	viper.SetDefault("api.middlewares.csrf.secure.xss", "1; mode=block")
	viper.SetDefault("api.middlewares.csrf.secure.ctNosniff", "nosniff")
	viper.SetDefault("api.middlewares.csrf.secure.xFrame", "")
	viper.SetDefault("api.middlewares.csrf.secure.hstsMaxAge", 0)
	viper.SetDefault("api.middlewares.csrf.secure.hstsExclude", false)
	viper.SetDefault("api.middlewares.csrf.secure.csPolicy", "")

	viper.SetDefault("api.metricsPort", 8081)
	viper.SetDefault("api.swaggerPort", 8082)

	viper.SetDefault("api.auth.secret", "envs/dev/sec/jwt.key")
	viper.SetDefault("api.auth.accessExpiary", 600)
	viper.SetDefault("api.auth.refreshExpiary", 2592000)

	viper.SetDefault("db.host", "")
	viper.SetDefault("db.port", "")
	viper.SetDefault("db.dbName", "")
	viper.SetDefault("db.user", "")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.ssl", false)

	viper.SetDefault("tinkoff.terminalId", "")
	viper.SetDefault("tinkoff.token", "envs/dev/sec/tinkoff.token") // TODO:
}

func rewrite() error {
	if viper.GetString("log.lokiAuth.login") != "" {
		lokkiPasswordPath := viper.GetString("log.lokiAuth.password")
		lokiPasswordBytes, err := os.ReadFile(lokkiPasswordPath)
		if err != nil {
			return NewErr(0, err, "failed loading loki password", "config")
		}
		viper.Set("log.lokiAuth.password", string(lokiPasswordBytes))
	}

	jwtSecretPath := viper.GetString("api.auth.secret")
	jwtSecret, err := os.ReadFile(jwtSecretPath)
	if err != nil {
		return NewErr(0, err, "failed loading jwt secret", "config")
	}
	viper.Set("api.auth.secret", hex.EncodeToString(jwtSecret))

	dbPasswordPath := viper.GetString("db.password")
	dbPasswordBytes, err := os.ReadFile(dbPasswordPath)
	if err != nil {
		return NewErr(0, err, "failed loading db password", "config")
	}
	viper.Set("db.password", string(dbPasswordBytes))

	// TODO: When add tinkoff uncomment
	//tinkoffTokenPath := viper.GetString("tinkoff.token")
	//tinkoffTokenBytes, err := os.ReadFile(tinkoffTokenPath)
	//if err != nil {
	//	return NewErr(0, err, "failed loading tonkoff token", "config")
	//}
	//viper.Set("tinkoff.token", string(tinkoffTokenBytes))

	return nil
}
