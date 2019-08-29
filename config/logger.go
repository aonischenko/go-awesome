package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

// setting default logger
var Log = *logrus.NewEntry(logrus.New())

func ConfigureLogger(cfg Config) {
	var formatter logrus.Formatter
	formatter = &logrus.TextFormatter{}
	if strings.EqualFold("json", cfg.LogFormat) {
		// Log as JSON instead of the default ASCII formatter.
		formatter = &logrus.JSONFormatter{}
	}

	var level logrus.Level
	// Setting INFO level as default
	switch strings.ToUpper(cfg.LogLevel) {
	case "DEBUG":
		level = logrus.DebugLevel
		break
	case "TRACE":
		level = logrus.TraceLevel
		break
	default:
		level = logrus.InfoLevel
	}

	Log = logrus.Entry{
		Logger: &logrus.Logger{
			Out:          os.Stdout,
			Formatter:    formatter,
			Level:        level,
			ReportCaller: cfg.LogCaller,
		},
	}

	Log.Trace("Logger initialized successfully")
}
