# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that your API will run on
EXPOSE 8080

# Run the application
CMD ["./main"]