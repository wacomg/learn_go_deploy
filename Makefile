PROJECT?=git.orglot.office/opanteleev/learn_go_deploy
APP?=advent
PORT?=8000

RELEASE?=0.0.1

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')



clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	-ldflags "-s -w -X ./version.Release=${RELEASE} \
        -X ./version.Commit=${COMMIT} -X ${PROJECT}version.BuildTime=${BUILD_TIME}" \
	-o ${APP}

run: build
	PORT=${PORT} ./${APP}

test:
	go test -v -race ./...
