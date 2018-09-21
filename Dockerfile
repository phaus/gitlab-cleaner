FROM alpine:3.8

LABEL maintainer=philipp@haussleiter.de

USER root

RUN apk add --update --no-cache ca-certificates

RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

COPY bin/cleaner /usr/local/bin/cleaner

RUN chmod 0755 /usr/local/bin/cleaner

RUN ["chmod", "-R", "+x", "/usr/local/bin/"]

COPY cleaner-entrypoint.sh /usr/local/bin/

ENTRYPOINT ["cleaner"]

CMD ["--help"]