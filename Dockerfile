#
# Build
#

FROM golang:alpine as builder

WORKDIR /go/src/app
COPY . .

RUN apk add --update-cache make && \
    make install

#
# Run
#

FROM alpine:latest

COPY --from=builder /go/bin/waiter /usr/local/bin/waiter

USER 1000

ENTRYPOINT ["waiter"]
