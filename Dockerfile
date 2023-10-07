# Use the official Golang image as the base image
FROM golang:1.20

# Copy the necessary files into the container
COPY . /app

# Set the working directory inside the container
WORKDIR /app

RUN go mod tidy

# Build the Go application
RUN go build -o app .

# Database configuration
ENV DB_USERNAME=root
ENV DB_PASSWORD=
ENV DB_PORT=3306
ENV DB_HOST=localhost
ENV DB_NAME=go_simple_clean_rest_api

# JWT secret
ENV JWT_SECRET=kayokoonikata

# Expose the port that your API will run on
EXPOSE 8080

# Run the application
CMD ["/app/app"]
