package api

import (
	"net/http"
	"testing"
)

func TestDivGet(t *testing.T) {
	router := prepareRouter(NewV1())

	t.Log("Normal flow scenario for GET /v1/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "GET",
		Url:      "/v1/div?x=15&y=10",
		Code:     http.StatusOK,
		Expected: `^{"operation":{"name":"division","x":15,"y":10},"success":true,"result":"1\(1/2\)"}$`,
	})

	t.Log("Division by zero scenario for GET /v1/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "GET",
		Url:      "/v1/div?x=15&y=0",
		Code:     http.StatusBadRequest,
		Expected: `^{"status":400,"message":"Operation error","details":"division by zero","ts":"[0-9T:\-\.]+Z"}$`,
	})

	t.Log("Unprocessable entry scenario for GET /v1/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "GET",
		Url:      "/v1/div?x=15&y=a",
		Code:     http.StatusUnprocessableEntity,
		Expected: `^{"status":422,"message":"Can't read input entity","details":"1 error.*","ts":"[0-9T:\-\.]+Z"}$`,
	})
}

func TestDivPut(t *testing.T) {
	router := prepareRouter(NewV1())

	t.Log("Normal flow scenario for PUT /v1/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "PUT",
		Url:      "/v1/div",
		Body:     `{"x":15,"y":10}`,
		Code:     http.StatusOK,
		Expected: `^{"operation":{"name":"division","x":15,"y":10},"success":true,"result":"1\(1/2\)"}$`,
	})

	t.Log("Division by zero scenario for GET /v1/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "PUT",
		Url:      "/v1/div",
		Body:     `{"x":15,"y":0}`,
		Code:     http.StatusBadRequest,
		Expected: `^{"status":400,"message":"Operation error","details":"division by zero","ts":"[0-9T:\-\.]+Z"}$`,
	})

	t.Log("Unprocessable entry scenario for GET /v1/div")
	checkTestCase(testRequestCase{
		Test:     t,
		Router:   router,
		Method:   "PUT",
		Url:      "/v1/div",
		Body:     `{"x":15,"y":"a"}`,
		Code:     http.StatusUnprocessableEntity,
		Expected: `^{"status":422,"message":"Can't read input entity","details":"json: cannot unmarshal string.*","ts":"[0-9T:\-\.]+Z"}$`,
	})
}
