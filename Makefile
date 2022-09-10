ARG = table

.PHONY: setup
setup:
	docker-compose up -d --build

.PHONY: migrate-create
migrate-create: 
	goose create ${ARG} sql

.PHONY: test
test:
	docker exec -it go go test golang/src/controller/handler/user/main_test.go