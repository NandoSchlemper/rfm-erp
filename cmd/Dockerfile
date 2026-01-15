FROM golang:1.25.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN apk add --no-cache ca-certificates && \
    update-ca-certificates
COPY . .
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o main ./api/

FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main .
CMD ["./main"]