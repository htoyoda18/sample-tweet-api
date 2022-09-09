ARG = table

.PHONY: setup
setup:
	docker-compose up -d --build

.PHONY: migrate-create
migrate-create: 
	goose create ${ARG} sql