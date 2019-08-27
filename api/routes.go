package api

import (
	"github.com/julienschmidt/httprouter"
)

/*
Define all the API routes/versions here.
*/
type API interface {
	RegisterHandlers(router *httprouter.Router)
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
