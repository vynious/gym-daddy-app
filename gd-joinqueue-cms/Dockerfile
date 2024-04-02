FROM golang:1.21.5

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Copy the .gitignore file into the container
# COPY .gitignore .


# Build the Go app
RUN go build -o main main.go
EXPOSE 3003

# Command to run the executable
ENTRYPOINT ["/app/main"]
