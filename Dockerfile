FROM golang:1.16 as intermediate

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN GO_ENABLED=0 GOOS=linux go build -o rest-api cmd/server/main.go

FROM alpine:latest AS production
COPY --from=intermediate /app/rest-api .
CMD [./rest-api]
