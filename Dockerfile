# build stage
FROM golang:1.16beta1-alpine3.12 AS build-env

WORKDIR /app

COPY ./src/ ./

RUN go build -ldflags="-s -w" -o demoapp main.go

# runtime stage
FROM alpine:3.12
WORKDIR /app
COPY --from=build-env /app/demoapp .
ENTRYPOINT ["./demoapp"]
