package response

import (
	"encoding/json"
	"net/http/httptest"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
)

func UserResponse(req *httptest.ResponseRecorder) *model.User {
	var body *model.User
	json.Unmarshal([]byte(req.Body.Bytes()), &body)

	return body
}
