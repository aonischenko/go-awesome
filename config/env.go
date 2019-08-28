package config

type Config struct {
	Host      string `env:"HOST"`
	Port      int    `env:"PORT" envDefault:"8080"`
	LogLevel  string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	LogFormat string `env:"LOG_FORMAT"`
	LogCaller bool   `env:"LOG_CALLER"`
}
