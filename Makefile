.PHONY: all weatherllo tests docker release clean

all: tests weatherllo

weatherllo:
	@go build ./cmd/weatherllo

tests:
	go test -v ./owm

tests:
	go test -v ./graphql
