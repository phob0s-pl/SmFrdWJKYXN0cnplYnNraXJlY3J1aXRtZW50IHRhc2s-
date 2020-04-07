.PHONY: all weatherllo tests docker release clean

all: tests weatherllo

weatherllo:
	@gqlgen
	@go build ./cmd/weatherllo

tests:
	go test -v ./owm

clean:
	rm -fr model
	rm -f graphql/generated.go
