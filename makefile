SHELL := /bin/bash

## Golang Stuff
GOCMD=go
GORUN=$(GOCMD) run
ENV=local

SERVICE=perqara

init:
	$(GOCMD) mod init $(SERVICE)

tidy:
	ENV=local GOPRIVATE=$(GOPRIVATE) $(GOCMD) mod tidy

run:
	swag init && ENV=$(ENV) $(GORUN) main.go

# Swagger API docs
SWAGGER_PORT=51234
