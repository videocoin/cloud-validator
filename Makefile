GOOS?=linux
GOARCH?=amd64

NAME=validator
VERSION?=$$(git rev-parse HEAD)

REGISTRY_SERVER?=registry.videocoin.net
REGISTRY_PROJECT?=cloud

.PHONY: deploy

default: build

version:
	@echo ${VERSION}

build:
	GOOS=${GOOS} GOARCH=${GOARCH} \
		go build \
			-mod vendor \
			-ldflags="-w -s -X main.Version=${VERSION}" \
			-o bin/${NAME} \
			./cmd/main.go

deps:
	GO111MODULE=on go mod vendor

lint: docker-lint

release: docker-build docker-push

docker-build:
	docker build -t ${REGISTRY_SERVER}/${REGISTRY_PROJECT}/${NAME}:${VERSION} -f Dockerfile .

docker-push:
	docker push ${REGISTRY_SERVER}/${REGISTRY_PROJECT}/${NAME}:${VERSION}

deploy:
	helm upgrade -i --wait --set image.tag="${VERSION}" -n console validator ./deploy/helm