FROM golang:1.20-alpine AS Build

WORKDIR /usr/local/go/src/github.com/aws-ses-mock

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./templates ./templates

RUN apk add tree
RUN tree

RUN go build -o ./main ./cmd/main.go
RUN mkdir /build && \
    mv main /build/main

FROM alpine:2.6

WORKDIR /app
COPY --from=Build /build/main /app/main

ENV GIN_MODE=release

EXPOSE 8080
ENTRYPOINT [ "/app/main" ]
