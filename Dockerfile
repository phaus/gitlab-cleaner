FROM golang:latest

WORKDIR /scripts

COPY scripts/ /scripts

CMD /scripts/get.sh

