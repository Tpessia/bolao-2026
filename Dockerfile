FROM golang:1.25.3-alpine AS builder

WORKDIR /app

RUN apk update
RUN apk add make git
RUN rm -rf /var/cache/apk/*

COPY go.mod go.sum Makefile .
RUN make init

COPY . .
RUN make test
RUN make build

# FROM scratch
# COPY --from=builder /app/myapp /
# # Copy any other necessary files, like SSL certs or config files
# # COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# CMD ["/myapp"]

FROM alpine:3.23.2
COPY --from=builder /app/bolao-2026 /app/
CMD ["/app/bolao-2026"]