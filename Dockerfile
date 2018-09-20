FROM alpine:3.4

LABEL maintainer=philipp@haussleiter.de

USER root

RUN apk add --update ca-certificates openssl

COPY bin/cleaner /app/cleaner

RUN chmod 0755 /app/cleaner

RUN ["chmod", "-R", "+x", "/app"]

ENV PATH=/app:$PATH

WORKDIR /app

CMD ["/app/cleaner"] 