FROM golang:1.21-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

# ADD .ENV file to /app
ADD .env /app

RUN go build -o main cmd/main.go

CMD ["/app/main"]
