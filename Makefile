GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init
# init env
init:
	go get entgo.io/ent/cmd/ent

.PHONY: example-add
# 输出样板命令：初始化一个user(默认生成目录在项目根目录ent 下 ./ent/schema)， 并生成user的 entity ; 使用go run 是保证ent命令使用的库版本一致
example-add:
	@echo "go run entgo.io/ent/cmd/ent@v0.9.0 init User"
#	@echo "go generate ./ent"
#	@echo "go run -mod=vendor entgo.io/ent/cmd/ent generate --target {./internal/data/ent} ./ent/schema/"

.PHONY: schema
# generate ent proto
schema:
	protoc --proto_path=. \
	       --proto_path=./third-party/ \
		   --proto_path=$$GOPATH \
 	       --ent_out=./internal/ent \
 	       --ent_opt=schemadir=./ent/schema/ \
	       $(API_PROTO_FILES)

.PHONY: gen-schema
# generate ent schema to sql
gen-schema:
#	go run -mod=mod entgo.io/ent/cmd/ent generate --target ./internal/ent ./ent/schema
	go run -mod=mod entgo.io/ent/cmd/ent generate --target ./internal/ ./ent/schema/


.PHONY: run
# run main.go
run:
	go run -mod=vendor -v main.go



.PHONY: swagger
# gen swagger
swagger:
	protoc --proto_path=. \
    --proto_path=./third-party/ \
    --go_out=paths=source_relative:. \
    --openapiv2_out . \
    api/user/user.proto

.PHONY: help
# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)


.DEFAULT_GOAL := help
