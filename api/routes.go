package api

import (
	"goawesome/model"
)

/*
Define all the API routes/versions here.
*/
//todo refactor according to https://stackoverflow.com/questions/42606535/how-to-create-a-prefix-route-in-swagger
func AllRoutes() model.Versions {
	return model.Versions{
		{Routes: v1(), Version: "v1"},
		{Routes: v2(), Version: "v2"},
	}
}
