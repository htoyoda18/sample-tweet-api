package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
	"github.com/htoyoda18/sample-tweet-api/golang/src/injector"
	"github.com/htoyoda18/sample-tweet-api/golang/src/router"
)

type Result struct {
	Url    string      `json:"url"`
	Method string      `json:"method"`
	Body   interface{} `json:"body"`
}

func Request(t *testing.T, pass string, index ...int) *httptest.ResponseRecorder {
	read, err := ioutil.ReadFile(pass)
	if err != nil {
		t.Errorf("request fail to file error %s", err.Error())
		return nil
	}

	var result Result
	json.Unmarshal([]byte(read), &result)

	jsonData, _ := json.Marshal(result.Body)

	body := buffer(string(jsonData))

	r := router.SetupRouter()
	w := httptest.NewRecorder()

	req, errReq := http.NewRequest(result.Method, result.Url, body)
	if errReq != nil {
		log.Fatal(errReq)
		return nil
	}

	token, errToken := bearerToken()

	if errToken != nil {
		t.Errorf("request fail to token error %s", errToken.Error())
		return nil
	}

	req.Header.Add("Authorization", token)

	r.ServeHTTP(w, req)

	return w
}

// token取得
func bearerToken() (string, error) {
	service := injector.NewService()

	jwt := service.AuthCore.NewJwt(&model.User{
		ID: 1,
	})

	return "Bearer " + jwt, nil
}

func buffer(str string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(str))
}
