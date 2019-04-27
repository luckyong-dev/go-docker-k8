# Builder step
FROM golang:1.10-alpine AS builder
WORKDIR /go/src/github.com/luckyong-dev/learn-k8s
ADD . .
RUN apk update \
    && apk add git
RUN go get
RUN go build -o hellobin .

# Deployment
FROM alpine
WORKDIR /app
COPY --from=builder /go/src/github.com/luckyong-dev/learn-k8s/hellobin .
CMD ./hellobin
