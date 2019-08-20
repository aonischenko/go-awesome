package main

import (
	"github.com/sirupsen/logrus"
	"goawesome/api"
	"goawesome/config"
	"net/http"
)

const addr = ":8080"

//todo add configurable environments
//todo add swagger
func main() {
	router := config.AppRouter(api.AllRoutes())
	logrus.Fatal(http.ListenAndServe(addr, router))
}
