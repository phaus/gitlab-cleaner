FROM alpine:3.4

RUN apk add --update ca-certificates openssl

COPY bin/cleaner /bin/cleaner

RUN chmod +x /bin/cleaner

ENTRYPOINT ["/bin"]

CMD []