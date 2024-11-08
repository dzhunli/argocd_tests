FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o helloapp

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/helloapp .
CMD ["./helloapp"]

