package api

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Method  string
	Path    string
	Handler httprouter.Handle
}

type Routes []Route

type API interface {
	ListRoutes() Routes
}

type APIs []API

const (
	Version1 = "v1"
	Version2 = "v2"
)

/*
Define all the API routes/versions here.
*/
func ListAPIs() APIs {
	return APIs{
		NewV1(),
		NewV2(),
		&Common{},
	}
}
