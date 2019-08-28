package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func ConfigureLogger(cfg Config) {
	if strings.EqualFold("json", cfg.LogFormat) {
		// Log as JSON instead of the default ASCII formatter.
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	// Output to stdout instead of the default stderr
	logrus.SetOutput(os.Stdout)

	// Setting INFO level as default
	switch strings.ToUpper(cfg.LogLevel) {
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	if cfg.LogCaller {
		logrus.SetReportCaller(true)
	}

	logrus.Trace("Logger initialized successfully")
}
