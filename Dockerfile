FROM golang:1.18.0-buster as build

RUN apt-get update

COPY . /discord_logger
RUN cd /discord_logger && go build -o service main.go

FROM debian:buster-slim

RUN apt-get update && apt-get install -y ca-certificates openssl
ARG cert_location=/usr/local/share/ca-certificates
RUN update-ca-certificates

RUN mkdir -p /opt/discord_logger
COPY --from=build /discord_logger/service /opt/discord_logger/service
COPY --from=build /discord_logger/config/config.json /opt/discord_logger/config/config.json

WORKDIR /opt/discord_logger
CMD ["./service"]