FROM golang:latest

WORKDIR /app

COPY cleaner/ /app

CMD /app/cleaner

