# Build flarectl binary

FROM golang:buster AS flarectl

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /src

ADD . /src/

WORKDIR /src/cmd/flarectl

RUN GOOS=linux go build -o flarectl

# Build discord bot binary

FROM golang:buster AS discord

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /src

ADD . /src/

RUN GOOS=linux go build -o discord

# Run time

FROM debian:buster

RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates

RUN update-ca-certificates

WORKDIR /src

COPY --from=flarectl /src/cmd/flarectl/flarectl /usr/bin/flarectl

RUN chmod +x /usr/bin/flarectl

COPY --from=discord /src/discord discord

CMD ["./discord"]