.PHONY: run
run: style
	go run ./cmd/main.go

.PHONY: style
style: build
	golint ./...

.PHONY: build
build: generate
	go build ./...

.PHONY: generate
generate: dep
	cd generated/cmd; \
	go generate; \
	cd ../..

.PHONY: dep
dep:
	go get golang.org/x/lint/golint
