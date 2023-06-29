ARG TARGETARCH

FROM --platform=$TARGETARCH golang:1.20-alpine AS Build

WORKDIR /usr/local/go/src/github.com/aws-ses-mock

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./assets ./assets

RUN apk add tree
RUN tree

RUN go build -o ./main ./cmd/main.go
RUN mkdir -p /build/bin && \
    mv ./main /build/bin/main && \
    mv ./assets /build/assets

FROM alpine:2.6

# Set our workdir to /bin since the binary expects the assets to be located under ../assets
WORKDIR /app/bin

COPY --from=Build /build/bin/main /app/bin/main
COPY --from=Build /build/assets /app/assets

ENV GIN_MODE=release

EXPOSE 8081
ENTRYPOINT [ "/app/bin/main" ]
