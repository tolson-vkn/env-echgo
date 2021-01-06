FROM golang:1.15 AS builder

WORKDIR /opt/build
ADD main.go ./
ADD main_test.go ./
ADD go.mod ./

RUN go test

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w'

# ---

FROM alpine:latest

EXPOSE 8080

WORKDIR /root/
COPY --from=builder /opt/build/env_echgo /root/env_echgo
CMD ["./env_echgo"]
