#---Build stage---
FROM golang:1.20 AS builder
COPY . /go/src/one-way-bot
WORKDIR /go/src/one-way-bot/cmd

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/service

#---Final stage---
FROM alpine:latest
COPY --from=builder /go/bin/service /go/bin/service

EXPOSE 7785
ENTRYPOINT ["go/bin/service"]