.PHONY: build
build:
	CGO_ENABLED=1 GOARCH=arm64 GOHOSTARCH=arm64 go build -o ../lib/goprocess/goprocess.a -buildmode=c-archive main.go
