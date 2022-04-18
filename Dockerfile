FROM golang:1.13-alpine as dev

RUN apk add --no-cache make git curl build-base

COPY cmd/api /app/cmd/api
COPY internal/app /app/internal/app
COPY internal/pkg /app/internal/pkg
COPY go.mod /app
COPY go.sum /app

WORKDIR /app

RUN go build -o ./build/api developer-orientenergy-golang/cmd/api

EXPOSE 8080

ENTRYPOINT ./build/api
