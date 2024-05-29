.PHONY: build_go
build_go:
	CGO_ENABLED=1 GOARCH=arm64 GOHOSTARCH=arm64 go build -o gohello.so -buildmode=c-shared main.go

.PHONY: build
build: build_go
	g++ -o main main.cpp ./gohello.so

.PHONY: run
run: build
	./main
