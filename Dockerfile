FROM golang:1.23.4-alpine3.20 as build
WORKDIR /build
COPY . .

RUN /usr/local/go/bin/go build -ldflags "-s -w" ./cmd/forwarder

### TARGET Container
FROM alpine:latest
LABEL authors="thanhnt169"

COPY --from=build /build/forwarder /main/forwarder

# RUN command
CMD /main/forwarder

