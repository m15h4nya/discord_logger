# syntax=docker/dockerfile:1
FROM golang:1.17.1-buster AS project_base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /discord_logger

EXPOSE 8080

CMD [ "/discord_logger" ]