#first stage - builder
FROM golang:1.17-stretch as builder
COPY . /app
WORKDIR /app
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o apiserver /app/cmd/apiserver/main.go
#second stage
FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache tzdata
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app .
CMD ["./apiserver"]
