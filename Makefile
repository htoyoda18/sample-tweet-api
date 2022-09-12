include $(PWD)/golang/src/controller/handler/Makefile
include $(PWD)/golang/src/Makefile

ARG = table

.PHONY: setup
setup:
	docker-compose up -d --build
	make swag

.PHONY: ps
ps:
	docker-compose ps

.PHONY: start
start:
	docker-compose up -d