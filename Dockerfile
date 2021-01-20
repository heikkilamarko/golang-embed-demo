# build stage

FROM golang:1.16-rc-alpine3.13 AS build-env

WORKDIR /app

COPY ./src/ ./

RUN go build -ldflags="-s -w" -o demoapp main.go

# runtime stage

FROM alpine:3.13

WORKDIR /app

COPY --from=build-env /app/demoapp .

ENTRYPOINT ["./demoapp"]
