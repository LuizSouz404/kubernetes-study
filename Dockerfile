FROM golang:1.18

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get ./...

RUN go build -o main .

EXPOSE 3000

CMD ["/app/main"]