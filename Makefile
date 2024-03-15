HOSTNAME=hsulimann.com
NAMESPACE=dev
NAME=structure-deploy
BINARY=terraform-provider-${NAME}
VERSION=0.0.4-223
OS_ARCH=darwin_amd64

default: install

build:
	go clean -cache
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	rm -f ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${BINARY}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	rm -f ./terraform/.terraform.lock.hcl
	terraform init
	
