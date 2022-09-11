NAME=all
FN=all

.PHONY: test
test:
	docker exec -it go go test -v controller/handler/${NAME}/main_test.go -run $(FN) -count=1

.PHONY: test-ci
test-ci:
	docker exec -it go go test -v github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/login/... -count=1
	docker exec -it go go test -v github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/tweet/... -count=1
	docker exec -it go go test -v github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/user/... -count=1