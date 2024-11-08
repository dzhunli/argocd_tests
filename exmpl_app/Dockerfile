FROM golang:1.19 AS builder
WORKDIR /app
COPY main.go .
RUN go mod init helloapp
RUN go mod tidy
RUN go build -o helloapp

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/helloapp .
CMD ["./helloapp"]

