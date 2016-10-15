default: build

ORG          = jasonrichardsmith
REPO         = pipesandfilters
BUILD_REPO   = ${ORG}/${REPO}-build
DOCKER_REPO   = ${ORG}/${REPO}
VERSION     ?= latest
DOCKER_RM   ?= --rm
SERVER_PORT ?= 8080

# Building and testing

pull:
	@docker pull golang:1.7-alpine

build:
	mkdir -p app
	docker build -t ${BUILD_REPO} -f ./docker/Dockerfile-build .
	@docker create --name build ${BUILD_REPO}
	docker cp build:/http app/
	docker build -t ${DOCKER_REPO}:${VERSION} -f ./docker/Dockerfile .

run:
	@docker run ${DOCKER_RM} -e SERVER_PORT ${DOCKER_REPO}:latest


