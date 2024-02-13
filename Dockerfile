FROM golang:1.21.5

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .
ENV KAFKA_URL=""
ENV KAFKA_TOPIC="notification"

# Build the Go app
RUN go build -o notification-ms .

# Command to run the executable
CMD ["./notification-ms"]