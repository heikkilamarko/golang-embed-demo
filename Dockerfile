FROM node:lts AS build-ui
WORKDIR /app
COPY ui/package*.json ./
RUN npm ci
COPY ui/ .
RUN npm run build

FROM golang AS build
COPY ./src/ ./
COPY --from=build-ui /app/build/ ./ui/
ENV GOPATH=""
ENV CGO_ENABLED=0
RUN go build -trimpath -a -ldflags="-w -s" -o golang-embed-demo

FROM gcr.io/distroless/static
COPY --from=build /go/golang-embed-demo /golang-embed-demo
ENTRYPOINT ["/golang-embed-demo"]
