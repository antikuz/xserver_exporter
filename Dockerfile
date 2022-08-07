FROM golang:1.18.3-alpine3.16 AS builder

WORKDIR /build

COPY . ./
RUN go mod download \
 && go build -o xserver_exporter


FROM alpine:3.16

WORKDIR /app

COPY --from=builder /build/xserver_exporter ./xserver_exporter

ENTRYPOINT ["/app/xserver_exporter"]