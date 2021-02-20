# build stage
FROM golang:1.16 AS build

COPY ./src/ ./

ENV GOPATH=""
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -trimpath -a -ldflags="-w -s" -o api

RUN useradd -u 12345 apiuser

# runtime stage
FROM scratch

COPY --from=build /go/api /api
COPY --from=build /etc/passwd /etc/passwd

USER apiuser

ENTRYPOINT ["/api"]
