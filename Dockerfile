#
# Build
#

FROM golang:latest as builder

WORKDIR /go/src/app
COPY . .

RUN apt-get install -y make && \
    make install

#
# Run
#

FROM debian:10

COPY --from=builder /go/bin/waiter /usr/local/bin/waiter

ENTRYPOINT ["waiter"]
