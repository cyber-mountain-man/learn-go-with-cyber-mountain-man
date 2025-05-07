# ----------------------
# 🛠️ Build Stage
# ----------------------
    FROM golang:1.24 AS builder

    WORKDIR /app
    
    # Copy Go module files and download dependencies
    COPY go.mod go.sum ./
    RUN go mod tidy
    
    # Copy the rest of the source code including static assets
    COPY . .
    # Explicitly copy .env into builder stage
    COPY .env .env    
    # Build the binary (targeting Linux for distroless image)
    RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app ./cmd/api
    
    # ----------------------
    # 🧼 Final Stage (Distroless)
    # ----------------------
    FROM gcr.io/distroless/static:nonroot
    
    WORKDIR /app
    
    # Copy the built binary
    COPY --from=builder /app/app .
    
    # ✅ Copy the static folder to serve frontend files
    COPY --from=builder /app/static ./static
    
    # ✅ Copy the .env file into the container
    COPY --from=builder /app/.env .env
    
    # Run as non-root user (recommended for security)
    USER nonroot:nonroot
    
    # Entry point for the application
    CMD ["/app/app"]
    