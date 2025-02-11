FROM golang:1.23.1-alpine

WORKDIR /usr/src/app

COPY . .

RUN go build -o editor-commacount

CMD ["./editor-commacount"]
