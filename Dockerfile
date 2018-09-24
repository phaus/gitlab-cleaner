FROM phaus/gobelt:latest

RUN mkdir -p $GOPATH/src/github.com/phaus/gitlab-cleaner /dist

WORKDIR $GOPATH/src/github.com/phaus/gitlab-cleaner
COPY . $GOPATH/src/github.com/phaus/gitlab-cleaner

RUN cd $GOPATH/src/github.com/phaus/gitlab-cleaner && \
    glide install && \
    go get ./... && \
    go fmt $(go list ./... | grep -v /vendor/) && \
    go vet $(go list ./... | grep -v /vendor/) && \
    go test ./...

RUN go build -o /dist/cleaner && upx /dist/cleaner 

FROM golang:1.11.0-stretch

USER root

LABEL maintainer=philipp@haussleiter.de

RUN mkdir -p /app && \
    apt-get update -y  && apt-get upgrade -y && \
    apt-get autoremove -y && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /var/cache/*

COPY --from=0 /dist/cleaner /app/cleaner

ENV PATH=/app:$PATH

WORKDIR /app

CMD ["/app/cleaner"]