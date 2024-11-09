FROM golang:1.19 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
WORKDIR /app
COPY exmpl_app/main.go .
RUN go mod init helloapp
RUN go mod tidy
RUN go build -a -installsuffix cgo -o helloapp

FROM alpine:latest
COPY --from=builder /app/helloapp /
CMD ["/helloapp"]
