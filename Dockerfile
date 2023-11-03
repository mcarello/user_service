# Use the official Go 1.21.3 image as the base image
FROM golang:1.21.3

# Set the working directory in the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . /app

# Build the Go application inside the container
RUN go get -d -v ./...
RUN go install -v ./...

# Expose port 1323
EXPOSE 1323

# Command to run the Go application
CMD ["go", "run", "./cmd/user_service/user_service.go"]
