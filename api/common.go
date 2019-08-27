package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Common struct{}

/*
API V1 routes
*/
func (c *Common) Routes() Routes {
	return Routes{
		{Method: "GET", Path: "/panic", Handle: c.startPanic},
	}
}

func (c *Common) startPanic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// just for router recovery testing purposes
	panic("It's me and I just PANIC!")
}
