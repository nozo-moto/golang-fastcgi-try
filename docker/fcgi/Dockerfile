FROM golang:1.13.7-alpine3.11 as builder

WORKDIR /go/src/github.com/nozo-moto/golang-fastcgi-try
COPY . .
RUN apk update && apk add alpine-sdk
RUN make build

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/nozo-moto/golang-fastcgi-try/main /main

EXPOSE 8080
ENTRYPOINT ["/main"]

