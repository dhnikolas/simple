FROM golang:latest

WORKDIR /app

COPY one .

CMD ./one
