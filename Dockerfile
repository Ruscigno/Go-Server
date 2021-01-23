FROM golang:1-alpine as build

WORKDIR /app
COPY server server
RUN go build server/go-server.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/go-server /app/go-server

EXPOSE 8081
ENTRYPOINT [".go-server"]