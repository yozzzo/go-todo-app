# Use the official Golang image as the base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Run the binary
CMD ["./main"]
