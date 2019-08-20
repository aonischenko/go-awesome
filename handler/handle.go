package handler

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"goawesome/model"
	"io"
	"io/ioutil"
	"net/http"
)

const MaxBodySize = 1048576

type requestReader func(r *http.Request, model interface{}) error

/*
Writes handler response with payload as json format
*/
func WriteOk(w http.ResponseWriter, payload interface{}) {
	writeJson(w, http.StatusOK, payload)
}

/*
Writes an error response with payload as json format
*/
func WriteError(w http.ResponseWriter, apiError model.ApiError) {
	writeJson(w, apiError.Status, apiError)
}

/*
Reads model from request URL parameters
e.g. /test?a=100&b=something
*/
func ReadUrlParams(r *http.Request, model interface{}) error {
	query := r.URL.Query()
	//todo replace with decoder hook
	params := make(map[string]string, len(query))
	for key, val := range query {
		params[key] = val[0]
	}
	return mapstructure.WeakDecode(params, model)
}

/*
Reads model from request body as json format
*/
func ReadBody(r *http.Request, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MaxBodySize))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}

func writeJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponseBody(w, []byte(err.Error()))
		return
	}
	writeResponseBody(w, response)
}

func writeResponseBody(w http.ResponseWriter, payload []byte) {
	//todo checkTestCase how can it happen & if it's a good practice to call Fatal() in this case
	if _, err := w.Write(payload); err != nil {
		logrus.Fatalf("Error writing response body: %s", err.Error())
	}
}
