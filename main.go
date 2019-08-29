package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"goawesome/api"
	. "goawesome/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	cfg := Config{}
	if env.Parse(&cfg) != nil {
		logrus.Fatal("application configuration failed")
		os.Exit(1)
	}
	ConfigureLogger(cfg)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%v", cfg.Host, cfg.Port),
		Handler:      api.AppHandler(cfg),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	defer gracefulShutdown(srv)

	go func() {
		Log.Fatal(srv.ListenAndServe())
	}()
}

func gracefulShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)

	Log.Info("Shutting down")
	os.Exit(0)
}
