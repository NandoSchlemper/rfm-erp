FROM golang:1.25.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o main .api/

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]