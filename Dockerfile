FROM golang:1.24.3-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE 8082
CMD ["./main"]