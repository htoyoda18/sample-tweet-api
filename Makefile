include $(PWD)/golang/src/controller/handler/Makefile

ARG = table

.PHONY: setup
setup:
	docker-compose up -d --build

.PHONY: ps
ps:
	docker-compose ps