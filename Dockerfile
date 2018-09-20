FROM alpine:3.4

LABEL maintainer=philipp@haussleiter.de

USER root

RUN apk add --update ca-certificates openssl

COPY bin/cleaner /cleaner

RUN chmod 0755 /cleaner

RUN ["chmod", "+x", "/cleaner"]

WORKDIR /

CMD ["/cleaner"] 