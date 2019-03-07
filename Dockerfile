FROM golang:alpine3.8
RUN apk --update add git openssh upx build-base && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

WORKDIR /app
COPY . /app

RUN mkdir -p /dist

RUN go fmt $(go list ./... | grep -v /vendor/)
RUN go vet $(go list ./... | grep -v /vendor/)
RUN go test ./...

RUN go build -o /dist/cleaner && upx /dist/cleaner 

FROM alpine:3.8

USER root

LABEL maintainer=philipp@haussleiter.de

ADD cert/* /tmp/cert/

RUN apk add --update ca-certificates && \
    cp -R /tmp/cert/* /usr/share/ca-certificates/ && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

COPY --from=0 /dist/cleaner /app/cleaner

ENV PATH=/app:$PATH

WORKDIR /app

CMD ["/app/cleaner"]