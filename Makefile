default: build

clearbin:
	rm -rf ./bin

build: main.go clearbin
	go build -v -o ./bin/mytool ./main.go

run: build
	./bin/mytool

install: main.go
	go list -f '{{.Target}}'
	go install

.PHONY: default run clearbin
