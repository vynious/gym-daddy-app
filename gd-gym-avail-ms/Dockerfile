FROM golang:1.19

RUN apt-get update && apt-get install -y telnet


WORKDIR /app

COPY . .

COPY .env .

RUN go build -o main main.go

EXPOSE 3005

COPY scripts/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
