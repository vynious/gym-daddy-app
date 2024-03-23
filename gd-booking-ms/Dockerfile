FROM golang:1.21.6
WORKDIR /app
COPY . .
COPY .env .
RUN go build -o main main.go
EXPOSE 6000
ENTRYPOINT ["/app/main"]

