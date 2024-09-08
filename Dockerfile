FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o build/file-transaction-processor main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/build/file-transaction-processor .

COPY .env /root/
COPY data /root/data

EXPOSE 8080

CMD ["./file-transaction-processor"]
