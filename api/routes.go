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

/*
Define all the API routes/versions here.
*/
type API interface {
	ListRoutes() Routes
}

const (
	Version1 = "v1"
	Version2 = "v2"
)

func Versions() []API {
	return []API{
		NewV1(),
		NewV2(),
	}
}
