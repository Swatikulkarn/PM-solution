# Step 1: Build the Go app
FROM golang:1.20-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module and Go sum files
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o myapp

# Step 2: Create a minimal image with the app
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built application from the build stage
COPY --from=build /app/myapp .

# Expose the application port (adjust if needed)
EXPOSE 8080

# Run the Go app
CMD ["./myapp"]
