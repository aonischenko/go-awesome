package main

import (
	"github.com/sirupsen/logrus"
	"goawesome/api"
	"goawesome/config"
	"net/http"
)

const addr = ":8080"

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
	router := config.AppRouter(api.AllRoutes())
	logrus.Fatal(http.ListenAndServe(addr, router))
}
