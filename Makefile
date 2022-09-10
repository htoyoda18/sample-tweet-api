include $(PWD)/golang/src/controller/handler/Makefile
include $(PWD)/golang/src/Makefile

ARG = table

.PHONY: setup
setup:
	docker-compose up -d --build