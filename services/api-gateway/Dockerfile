FROM golang:1.26-alpine

WORKDIR /app

COPY . .

RUN go build -o gateway

EXPOSE 8080

CMD ["./gateway"]