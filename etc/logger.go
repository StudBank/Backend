package etc

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yukitsune/lokirus"
)

type customLogger struct {
	defaultField string
	formatter    logrus.Formatter
}

func (l customLogger) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Data["src"] = l.defaultField
	return l.formatter.Format(entry)
}

func GetPlainLogger(name string, logLevel logrus.Level) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(customLogger{
		defaultField: name,
		formatter:    logrus.StandardLogger().Formatter,
	})
	log.SetLevel(logLevel)

	return log
}

func GetLogger(name string, logLevel logrus.Level) *logrus.Logger {
	log := GetPlainLogger(name, logLevel)
	opts := lokirus.NewLokiHookOptions().
		WithLevelMap(lokirus.LevelMap{logrus.PanicLevel: "critical"}).
		WithFormatter(&logrus.JSONFormatter{}).
		WithStaticLabels(lokirus.Labels{
			"app":         viper.GetString("name"),
			"environment": viper.GetString("env"),
			"version":     viper.GetString("version"),
		})
	if viper.GetString("log.lokiAuth.login") != "" {
		opts = opts.WithBasicAuth(
			viper.GetString("log.lokiAuth.login"),
			viper.GetString("log.lokiAuth.password"),
		)
	}

	hook := lokirus.NewLokiHookWithOpts(
		viper.GetString("log.lokiAddress"),
		opts,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel)

	// Configure the logger
	log.AddHook(hook)

	log.Trace("Service logger set up")
	return log
}
