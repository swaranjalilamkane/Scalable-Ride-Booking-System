FROM --platform=linux/amd64 golang:1.25.1-alpine AS builder

WORKDIR /app

# Set Go proxy for networking issues
ENV GOPROXY=https://proxy.golang.org,direct

# Copy module files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code and drivers.json
COPY . .
COPY drivers.json drivers.json

# Build binary
RUN go build -o main .

FROM --platform=linux/amd64 alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/drivers.json .   

EXPOSE 8080
CMD ["./main"]
