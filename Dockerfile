# ---- Build stage ----
FROM golang:1.25.1-alpine AS builder

WORKDIR /src

# Cache module downloads
COPY go.mod go.sum ./
RUN go mod download

# Build a static binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/server .

# ---- Run stage ----
FROM alpine:3.20

# certs for any outbound TLS, and tzdata if you need timezones
RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/server /usr/local/bin/server

EXPOSE 8080
CMD ["server"]