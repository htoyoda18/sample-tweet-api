package test

import (
	"encoding/json"
	"net/http/httptest"
)

func NilResponse(req *httptest.ResponseRecorder) interface{} {
	var body interface{}
	json.Unmarshal([]byte(req.Body.Bytes()), &body)

	return body
}
