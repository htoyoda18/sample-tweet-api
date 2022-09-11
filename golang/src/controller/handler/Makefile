NAME=all
FN=all

.PHONY: test
test:
	docker exec -it go go test -v controller/handler/${NAME}/main_test.go -run $(FN) -count=1