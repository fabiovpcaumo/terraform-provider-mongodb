ifndef VERBOSE
	MAKEFLAGS += --no-print-directory
endif

default: install

.PHONY: install lint unit

OS_ARCH=linux_amd64
#
# Set correct OS_ARCH on Mac
UNAME := $(shell uname -s)
ifeq ($(UNAME),Darwin)
	HW := $(shell uname -m)
	ifeq ($(HW),arm64)
		ARCH=$(HW)
	else
		ARCH=amd64
	endif
	OS_ARCH=darwin_$(ARCH)
endif

HOSTNAME=registry.terraform.io
NAMESPACE=fabiovpcaumo
NAME=mongodb
VERSION=9.9.9
## on linux base os
TERRAFORM_PLUGINS_DIRECTORY=~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}


install:
	mkdir -p ${TERRAFORM_PLUGINS_DIRECTORY}
	go build -o ${TERRAFORM_PLUGINS_DIRECTORY}/terraform-provider-${NAME}
	cd examples && rm -rf .terraform*
	cd examples && make init
re-install:
	rm -f ${TERRAFORM_PLUGINS_DIRECTORY}/terraform-provider-${NAME}
	go build -o ${TERRAFORM_PLUGINS_DIRECTORY}/terraform-provider-${NAME}
	cd examples && rm -rf .terraform*
	cd examples && make init
lint:
	 golangci-lint run
