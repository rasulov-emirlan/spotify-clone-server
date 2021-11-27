FROM golang:latest
COPY . /app
WORKDIR /app
RUN go mod tidy
RUN make build
FROM alpine
COPY --from=0 /app/apiserver .
CMD ["./apiserver"]