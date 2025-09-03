# ---- Build stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# ---- Run stage ----
FROM scratch

WORKDIR /app

# Copy binary
COPY --from=builder /app/server .

# Copy static files
COPY static ./static

EXPOSE 8080

CMD ["./server"]


# docker build -t go-web-app .
# docker run -p 8080:8080 go-web-app