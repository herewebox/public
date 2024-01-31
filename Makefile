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
PROTO_GO_DIR=$(PREFIX)/pkg
DOC_DIR=$(PREFIX)/docs

.PHONY: all summary proto_go doc
.DEFAULT: all

# Targets
all: summary proto_go doc

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

doc:
