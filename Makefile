PROJECT?=github.com/wacomg/learn_go_deploy
APP?=advent
PORT?=8000

RELEASE?=0.1.0

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')



clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	-ldflags "-s -w -X main.Release=${RELEASE} \
        -X main.Commit=${COMMIT} -X main.BuildTime=${BUILD_TIME}" \
	-o ${APP}

run: build
	PORT=${PORT} ./${APP}

test:
	go test -v -race ./...
