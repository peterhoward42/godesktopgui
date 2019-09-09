.PHONY: generate
generate: 
	cd generate/cmd; \
	go generate; \
	cd ../..

.PHONY: build
build: generate
	go build ./...

.PHONY: style
style: build
	golint ./...

.PHONY: run
run: style
	go run ./cmd/main.go
