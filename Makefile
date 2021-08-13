GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init
# init env
init:
	go get entgo.io/ent/cmd/ent

.PHONY: example
# example 初始化一个user(默认生成目录在项目根目录ent 下 ./ent/schema)， 并生成user的 entity ; 使用go run 是保证ent命令使用的库版本一致
example:
	@echo "go run -mod=vendor entgo.io/ent/cmd/ent@v0.9.0 init User"
#	@echo "go generate ./ent"
#	@echo "go run -mod=vendor entgo.io/ent/cmd/ent generate --target {./internal/data/ent} ./ent/schema/"

.PHONY: schema
# generate ent proto
schema:
	@echo "protoc -I=$GOPATH -I=./third-party/ent/ --ent_out=./internal/ent/ --ent_opt=schemadir=./ent/schema/ api/hello/hello.proto"
	protoc --proto_path=. \
		   --proto_path=$GOPATH \
	       --proto_path=./third_party/ent \
 	       --ent_out=./internal/ent \
 	       --ent_opt=schemadir=./ent/schema/ \
	       $(API_PROTO_FILES)

.PHONY: gen-schema
# generate ent schema to sql
gen-schema:
	go generate ./ent

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
