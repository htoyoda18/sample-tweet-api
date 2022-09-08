package user

import (
	"encoding/json"
	"net/http/httptest"

	"github.com/htoyoda18/sample-tweet-api/domain/model"
)

func AddUsersResponse(req *httptest.ResponseRecorder) *model.User {
	var body *model.User
	json.Unmarshal([]byte(req.Body.Bytes()), &body)

	return body
}
