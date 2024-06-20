GOPATH:=$(shell go env GOPATH)

.PHONY: generate-gqlgen
generate-gqlgen:
	go run github.com/99designs/gqlgen generate --verbose

.PHONY: gqlgen
gqlgen:
	rm -rf ./graph/model/model.go
	rm -rf ./graph/*.generated.go
	rm -rf ./graph/federation.go
	go run github.com/99designs/gqlgen generate --verbose

.PHONY: run
run:
	go run main.go