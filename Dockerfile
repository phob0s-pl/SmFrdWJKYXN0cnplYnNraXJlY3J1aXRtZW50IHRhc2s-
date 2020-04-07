############################
# Build
############################
FROM golang:buster AS builder

RUN apt-get update && apt-get install git make
RUN go get -u github.com/99designs/gqlgen
WORKDIR $GOPATH/src/github.com/phob0s-pl/weatherllo
COPY . .
RUN make && mv weatherllo /weatherllo

############################
# Actual image
############################
FROM debian:buster

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /weatherllo /weatherllo

EXPOSE 8080/tcp

ENTRYPOINT ["/weatherllo", "listen", "--debug"]