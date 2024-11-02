# Build stage
FROM golang:1.23.2 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bot .

# Final stage
FROM alpine:latest

WORKDIR /

COPY --from=builder /bot /bot

ENTRYPOINT ["/bot"]
