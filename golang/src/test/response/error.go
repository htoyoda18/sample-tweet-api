package test

import (
	"net/http/httptest"
)

func ErrorResponseCode(req *httptest.ResponseRecorder) int {
	return req.Code
}
