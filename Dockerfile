FROM alpine:3.4

RUN apk add --update ca-certificates openssl sh

COPY bin/cleaner /usr/local/bin/cleaner

WORKDIR /

ENTRYPOINT ["cleaner"]

CMD ["version"]
