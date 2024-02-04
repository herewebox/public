PROJECT=box::public
PREFIX=$(shell pwd)
VERSION=$(shell git describe --match 'v[0-9]*'  --always)
DEFAULT_BRANCH=$(shell git symbolic-ref --short -q HEAD)

ifndef OS
	OS=linux
	unameOut=$(shell uname -s)
	ifeq ($(unameOut),Darwin)
		OS=darwin
	endif

	ifeq ($(OSTYPE),win32)
		OS=windows
	endif
endif

ifndef ARCH
	ARCH=amd64
	unameOut=$(shell uname -m)
	ifeq ($(unameOut),i386)
		ARCH=386
	endif

	ifeq ($(unameOut),i686)
		ARCH=386
	endif
endif

ifndef BRANCH
	BRANCH=$(DEFAULT_BRANCH)
endif

ifdef CI_COMMIT_REF_SLUG
	BRANCH=$(CI_COMMIT_REF_SLUG)
endif

ifndef DEPLOY_REPLICA
	DEPLOY_REPLICA=1
endif

ifndef GO
	GO=go
endif

ifndef GOFMT
	GOFMT=gofmt
endif

ifndef PROTOC
	PROTOC=protoc
endif

ifndef RUBY_PROTOC
	RUBY_PROTOC=grpc_tools_ruby_protoc
endif

ifndef PYTHON_PROTOC
	PYTHON_PROTOC=python -m grpc_tools.protoc
endif

ifndef GIT
	GIT=git
endif

ifndef SWAG
	SWAG=swag
endif

ifndef DOCKER
	DOCKER=docker
endif

PROTO_DIR=$(PREFIX)/proto
PROTO_GO_DIR=$(PREFIX)/go
PROTO_RUBY_DIR=$(PREFIX)/ruby
PROTO_PYTHON_DIR=$(PREFIX)/python
DOC_DIR=$(PREFIX)/docs

.PHONY: all summary proto_go proto_ruby proto_python proto_doc
.DEFAULT: all

# Targets
all: summary proto_go proto_ruby proto_python proto_doc

summary:
	@printf "\033[1;37m  == \033[1;32m$(PROJECT) \033[1;33m$(VERSION) \033[1;37m==\033[0m\n"
	@printf "    Platform    : \033[1;37m$(shell uname -sr)\033[0m\n"
	@printf "    Target OS   : \033[1;37m$(OS)\033[0m\n"
	@printf "    Target Arch : \033[1;37m$(ARCH)\033[0m\n"
	@printf "    Go          : \033[1;37m`$(GO) version`\033[0m\n"
	@printf "    Protoc      : \033[1;37m`$(PROTOC) --version`\033[0m\n"
	@printf "    Git         : \033[1;37m$(shell git version)\033[0m\n"
	@printf "    Branch      : \033[1;37m$(BRANCH)\033[0m\n"
	@echo

proto_go:
	@printf "\033[1;36m  Compiling protobuf definitions to go code ...\033[0m\n"
	@mkdir -p $(PROTO_GO_DIR)
	@$(PROTOC) -I $(PROTO_DIR) --go_out=. --go-grpc_out=. general.proto
	@$(PROTOC) -I $(PROTO_DIR)/account --go_out=. --go-grpc_out=. account.proto
	@$(PROTOC) -I $(PROTO_DIR)/tenant --go_out=. --go-grpc_out=. tenant.proto
	@$(PROTOC) -I $(PROTO_DIR)/world --go_out=. --go-grpc_out=. map.proto
	@$(PROTOC) -I $(PROTO_DIR)/world --go_out=. --go-grpc_out=. region.proto
	@GOOS=$(OS) GOARCH=$(ARCH) $(GO) mod tidy
	@echo

proto_ruby:
	@printf "\033[1;36m  Compiling protobuf definitions to ruby code ...\033[0m\n"
	@mkdir -p $(PROTO_RUBY_DIR)/box/account
	@mkdir -p $(PROTO_RUBY_DIR)/box/tenant
	@mkdir -p $(PROTO_RUBY_DIR)/box/world
	@$(RUBY_PROTOC) -I $(PROTO_DIR) --ruby_out=$(PROTO_RUBY_DIR)/box --grpc_out=$(PROTO_RUBY_DIR)/box general.proto
	@$(RUBY_PROTOC) -I $(PROTO_DIR)/account --ruby_out=$(PROTO_RUBY_DIR)/box/account --grpc_out=$(PROTO_RUBY_DIR)/box/account account.proto
	@$(RUBY_PROTOC) -I $(PROTO_DIR)/tenant --ruby_out=$(PROTO_RUBY_DIR)/box/tenant --grpc_out=$(PROTO_RUBY_DIR)/box/tenant tenant.proto
	@$(RUBY_PROTOC) -I $(PROTO_DIR)/world --ruby_out=$(PROTO_RUBY_DIR)/box/world --grpc_out=$(PROTO_RUBY_DIR)/box/world map.proto
	@$(RUBY_PROTOC) -I $(PROTO_DIR)/world --ruby_out=$(PROTO_RUBY_DIR)/box/world --grpc_out=$(PROTO_RUBY_DIR)/box/world region.proto
	@echo

proto_python:
	@printf "\033[1;36m  Compiling protobuf definitions to python code ...\033[0m\n"
	@mkdir -p $(PROTO_PYTHON_DIR)/box/account
	@mkdir -p $(PROTO_PYTHON_DIR)/box/tenant
	@mkdir -p $(PROTO_PYTHON_DIR)/box/world
	@$(PYTHON_PROTOC) -I $(PROTO_DIR) --python_out=$(PROTO_PYTHON_DIR)/box --pyi_out=$(PROTO_PYTHON_DIR)/box --grpc_python_out=$(PROTO_PYTHON_DIR)/box general.proto
	@$(PYTHON_PROTOC) -I $(PROTO_DIR)/account --python_out=$(PROTO_PYTHON_DIR)/box/account --pyi_out=$(PROTO_PYTHON_DIR)/box/account --grpc_python_out=$(PROTO_PYTHON_DIR)/box/account account.proto
	@$(PYTHON_PROTOC) -I $(PROTO_DIR)/tenant --python_out=$(PROTO_PYTHON_DIR)/box/tenant --pyi_out=$(PROTO_PYTHON_DIR)/box/tenant --grpc_python_out=$(PROTO_PYTHON_DIR)/box/tenant tenant.proto
	@$(PYTHON_PROTOC) -I $(PROTO_DIR)/world --python_out=$(PROTO_PYTHON_DIR)/box/world --pyi_out=$(PROTO_PYTHON_DIR)/box/world --grpc_python_out=$(PROTO_PYTHON_DIR)/box/world map.proto
	@$(PYTHON_PROTOC) -I $(PROTO_DIR)/world --python_out=$(PROTO_PYTHON_DIR)/box/world --pyi_out=$(PROTO_PYTHON_DIR)/box/world --grpc_python_out=$(PROTO_PYTHON_DIR)/box/world region.proto
	@echo

proto_doc:
