FROM golang:alpine AS build

COPY ./src/ ./

ENV GOPATH=""
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -trimpath -a -ldflags="-w -s" -o golang-embed-demo

FROM scratch

COPY --from=build /go/golang-embed-demo /golang-embed-demo

ENTRYPOINT ["/golang-embed-demo"]