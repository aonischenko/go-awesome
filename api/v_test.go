package api

import (
	"goawesome/config"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

type testRequestCase struct {
	Test     *testing.T
	Router   http.Handler
	Method   string
	Url      string
	Body     string
	Code     int
	Expected string
}

func checkTestCase(tc testRequestCase) {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(tc.Method, tc.Url, strings.NewReader(tc.Body))
	if err != nil {
		tc.Test.Fatal(err)
	}
	tc.Router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != tc.Code {
		tc.Test.Errorf("Wrong status: %v; expected: %v", status, tc.Code)
	}
	res := recorder.Body.String()
	matched, err := regexp.MatchString(tc.Expected, res)
	if err != nil {
		tc.Test.Fatal(err)
	}
	if !matched {
		tc.Test.Errorf("Response body doesn't match `%v`; expected pattern: `%v`", res, tc.Expected)
	}
}

func prepareRouter(api API) http.Handler {
	return AppHandler(config.Config{})
}
