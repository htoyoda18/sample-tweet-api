package response

import (
	"encoding/json"
	"net/http/httptest"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
)

func TweetResponse(req *httptest.ResponseRecorder) *model.Tweet {
	var body *model.Tweet
	json.Unmarshal([]byte(req.Body.Bytes()), &body)

	return body
}
