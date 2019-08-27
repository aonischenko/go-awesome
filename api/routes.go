package api

import (
	"github.com/julienschmidt/httprouter"
)

/*
Define all the API routes/versions here.
*/
type Route struct {
	Method string
	Path   string
	Handle httprouter.Handle
}

type Routes []Route

type API interface {
	Routes() Routes
}

const (
	Version1 = "v1"
	Version2 = "v2"
)

func Apis() []API {
	return []API{
		NewV1(),
		NewV2(),
		&Common{},
	}
}
