FROM golang:latest

RUN apk --no-cache add \
    ttf-dejavu=2.37-r0

WORKDIR /scripts

COPY scripts/ /scripts

CMD bash get.sh

