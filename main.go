package main

import (
	"fmt"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"goawesome/api"
	"goawesome/config"
	"net/http"
)

//todo add configurable environments
//todo check if swagger info can be moved into 'config' package

// @title Swagger Go Awesome API
// @version 1.0
// @description This is a sample Go Awesome server.
// @termsOfService http://awesome.go/terms/
// @contact.name API Support
// @contact.url http://www.awesome.go/support
// @contact.email support@awesome.go
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host awesome.go
// @BasePath /
func main() {
	cfg := config.Config{}
	if env.Parse(&cfg) != nil {
		log.Fatal("application configuration failed")
	}
	config.ConfigureLogger(cfg)
	handler := api.AppHandler(cfg)
	addr := fmt.Sprintf("%s:%v", cfg.Host, cfg.Port)
	log.Fatal(http.ListenAndServe(addr, handler))
}
