.PHONY: all weatherllo tests docker release clean gqlgen

all: gqlgen tests weatherllo

gqlgen:
	@gqlgen

weatherllo: gqlgen
	@go build ./cmd/weatherllo

tests: gqlgen
	go test -v ./owm

clean:
	rm -fr model
	rm -f graphql/generated.go
