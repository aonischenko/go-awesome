package handler

import (
	"github.com/julienschmidt/httprouter"
	"goawesome/config"
	"goawesome/model"
	"net/http"
	"testing"
)

func TestDivGetV2(t *testing.T) {
	router := prepareRouterV2()

	t.Log("Normal flow scenario for GET /v2/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "GET",
		Url:      "/v2/div?x=15&y=10",
		Code:     http.StatusOK,
		Expected: `^{"operation":{"name":"division","x":15,"y":10},"success":true,"result":"1\.5"}$`,
	})

	t.Log("Division by zero scenario for GET /v2/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "GET",
		Url:      "/v2/div?x=15&y=0",
		Code:     http.StatusOK,
		Expected: `^{"operation":{"name":"division","x":15,"y":0},"success":false,"result":"division by zero"}$`,
	})

	t.Log("Unprocessable entry scenario for GET /v2/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "GET",
		Url:      "/v2/div?x=15&y=a",
		Code:     http.StatusBadRequest,
		Expected: `^{"status":400,"message":"Can't read input entity","details":"1 error.*","ts":"[0-9T:\-\.]+Z"}$`,
	})
}

func TestDivPutV2(t *testing.T) {
	router := prepareRouterV2()

	t.Log("Normal flow scenario for PUT /v2/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "PUT",
		Url:      "/v2/div",
		Body:     `{"x":15,"y":10}`,
		Code:     http.StatusOK,
		Expected: `^{"operation":{"name":"division","x":15,"y":10},"success":true,"result":"1\.5"}$`,
	})

	t.Log("Division by zero scenario for GET /v2/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "PUT",
		Url:      "/v2/div",
		Body:     `{"x":15,"y":0}`,
		Code:     http.StatusOK,
		Expected: `^{"operation":{"name":"division","x":15,"y":0},"success":false,"result":"division by zero"}$`,
	})

	t.Log("Unprocessable entry scenario for GET /v2/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "PUT",
		Url:      "/v2/div",
		Body:     `{"x":15,"y":"a"}`,
		Code:     http.StatusBadRequest,
		Expected: `^{"status":400,"message":"Can't read input entity","details":"json: cannot unmarshal string.*","ts":"[0-9T:\-\.]+Z"}$`,
	})
}

func prepareRouterV2() *httprouter.Router {
	//todo checkTestCase how can we use "api" package w/o getting into circle dependencies
	//seems we just have to get routes & handlers within same package
	routes := model.Routes{
		{
			Name:        "Division using request url params",
			Method:      "GET",
			Path:        "/div",
			HandlerFunc: DivGetV2,
		},
		{
			Name:        "Division using request body",
			Method:      "PUT",
			Path:        "/div",
			HandlerFunc: DivPutV2,
		},
	}
	version := model.Version{Routes: routes, Version: "v2"}
	return config.AppRouter(model.Versions{version})
}
