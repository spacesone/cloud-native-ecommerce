# Stage 1: Build stage
FROM golang:1.24-alpine AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o inventory-service cmd/main.go

# Stage 2: Final stage
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/inventory-service .
COPY --from=build /app/cmd/dist cmd/dist
# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENTRYPOINT ["/app/inventory-service"]
