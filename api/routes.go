package api

import (
	"goawesome/model"
)

/*
Define all the API routes/versions here.
*/
func AllRoutes() model.Versions {
	return model.Versions{
		{Routes: v1(), Version: "v1"},
		{Routes: v2(), Version: "v2"},
	}
}
