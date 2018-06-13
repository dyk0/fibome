FROM golang:1.8.3 AS build
MAINTAINER ddyke@dyk0.tech

ADD . /usr/src/
RUN cd /usr/src && make build-linux && make test

FROM alpine
WORKDIR /usr/src/
COPY --from=build /usr/src/bin/fibome /usr/src/

ENTRYPOINT ["/usr/src/fibome"]
