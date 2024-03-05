FROM golang:1.21.5

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Copy the .env file into the container
COPY .env .


# Build the Go app
RUN go build -o main main.go
EXPOSE 3001

# Command to run the executable
ENTRYPOINT ["/app/main"]
