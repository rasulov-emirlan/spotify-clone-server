FROM golang:latest as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -v ./cmd/apiserver

FROM alpine
COPY --from=build /app/apiserver /app/.dev.env ./
CMD ["./apiserver"]
