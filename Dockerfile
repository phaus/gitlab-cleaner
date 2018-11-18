FROM golang:alpine3.8
RUN apk --update add git openssh upx glide build-base && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

RUN mkdir -p $GOPATH/src/github.com/phaus/gitlab-cleaner /dist

WORKDIR $GOPATH/src/github.com/phaus/gitlab-cleaner
COPY . $GOPATH/src/github.com/phaus/gitlab-cleaner

RUN glide install
RUN go fmt $(go list ./... | grep -v /vendor/)
RUN go vet $(go list ./... | grep -v /vendor/)
RUN go test ./...

RUN go build -o /dist/cleaner && upx /dist/cleaner 

FROM alpine:3.8

USER root

LABEL maintainer=philipp@haussleiter.de

RUN mkdir -p /app && \
    apk update && \
    apk upgrade && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

COPY --from=0 /dist/cleaner /app/cleaner

ENV PATH=/app:$PATH

WORKDIR /app

CMD ["/app/cleaner"]