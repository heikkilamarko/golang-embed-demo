FROM node:20-alpine AS build-ui
WORKDIR /app
COPY ui/package*.json ./
RUN npm ci
COPY ui/ .
RUN npm run build

FROM golang:alpine AS build
COPY ./src/ ./
COPY --from=build-ui /app/dist/ ./ui/
ENV GOPATH=""
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -trimpath -a -ldflags="-w -s" -o golang-embed-demo

FROM gcr.io/distroless/static
COPY --from=build /go/golang-embed-demo /golang-embed-demo
ENTRYPOINT ["/golang-embed-demo"]
