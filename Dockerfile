FROM golang:latest

WORKDIR /app

COPY cleaner/ /app

ENTRYPOINT ["/app/cleaner"]

CMD ["help"]
