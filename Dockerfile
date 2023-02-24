FROM golang:1.20-alpine AS Build

WORKDIR /usr/local/go/src/github.com/aws-ses-mock

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY ./cmd ./main

RUN apk add tree
RUN tree

RUN go build -o ./main ./main/main.go

FROM alpine:2.6

USER nonroo:nonroot
WORKDIR /app

COPY --from=Build /build/main /app/main

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
