BINARY=Auto-YunZhanYi-Go
VERSION=1.0.0
DATE=`date +%FT%T%z`
.PHONY: build build_mac build_arm

default:
	@echo ${BINARY}
	@echo ${VERSION}
	@echo ${DATE}

build:
	@GOOS=linux GOARCH=amd64 go build -o ${BINARY}
	@echo "[ok] build ${BINARY}"

build_mac:
	@go build -o ${BINARY}
	@echo "[ok] build_mac ${BINARY}"

build_arm:
	@GOOS=linux GOARCH=arm go build -o ${BINARY}
	@echo "[ok] build_arm ${BINARY}"