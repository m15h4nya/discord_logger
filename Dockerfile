# syntax=docker/dockerfile:1
FROM golang:1.17.1-buster AS build
WORKDIR /opt
ENV CGO_ENABLED=0
COPY go.mod /opt
RUN go mod download
COPY . /opt

RUN go build -o /discord_logger

FROM alpine:3.14

WORKDIR /

COPY --from=build /opt/config/config.json /config/config.json
COPY --from=build /discord_logger /discord_logger

CMD ["/discord_logger"]