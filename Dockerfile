# Builder step
FROM golang:1.12-alpine AS builder
WORKDIR /app
ADD . .
RUN apk update \
    && apk add git
RUN go mod vendor
RUN go build -o hellobin .

# Deployment
FROM alpine
WORKDIR /app
COPY --from=builder /app .
CMD ./hellobin
