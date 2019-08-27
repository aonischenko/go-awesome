package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

//todo some use variable to identify environment
func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	logrus.SetOutput(os.Stdout)

	// Setting DEBUG level as default
	// Should be configurable depending on env
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("Logger initialized successfully")
}
