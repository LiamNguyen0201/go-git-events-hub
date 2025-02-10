package utils

import (
	"git_events_hub/configs"
	"log"
	"net"
	"os"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/lokirus"
)

var logrusLogger = initLogger()

// Setup Logrus with Loki
func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Connect to Logstash via TCP
	if configs.LogWithLogstash {
		conn, err := net.Dial("tcp", configs.LogstashAddress) // Replace with your Logstash host/port
		if err != nil {
			log.Println("Failed to connect to Logstash:", err)
		} else {
			// Attach Logstash Hook
			hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"app": configs.ApplicationName}))
			logger.AddHook(hook)
		}
	}

	// Configure the Loki hook
	if configs.LogWithLoki {
		opts := lokirus.NewLokiHookOptions().
			// Grafana doesn't have a "panic" level, but it does have a "critical" level
			// https://grafana.com/docs/grafana/latest/explore/logs-integration/
			WithLevelMap(lokirus.LevelMap{logrus.PanicLevel: "critical"}).
			WithFormatter(&logrus.JSONFormatter{}).
			WithStaticLabels(lokirus.Labels{
				"app":         configs.ApplicationName,
				"environment": configs.Environment,
			}).
			WithBasicAuth(configs.LokiUsername, configs.LokiPassword) // Optional

		lokiHook := lokirus.NewLokiHookWithOpts(configs.LokiAddress, opts, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel)
		logger.AddHook(lokiHook)
	}

	return logger
}

func GetLogger() *logrus.Logger {
	return logrusLogger
}

func LogDebug(args ...interface{}) {
	logrusLogger.Debug(args...)
}

func LogDebugf(input string, args ...interface{}) {
	logrusLogger.Debugf(input, args...)
}

func LogInfo(args ...interface{}) {
	logrusLogger.Info(args...)
}

func LogInfof(input string, args ...interface{}) {
	logrusLogger.Infof(input, args...)
}

func LogFatal(args ...interface{}) {
	logrusLogger.Fatal(args...)
}
