FROM alpine:3.4

RUN apk add --update ca-certificates openssl

COPY --from=0 bin/cleaner /usr/local/bin/cleaner

WORKDIR /

ENTRYPOINT ["cleaner"]

CMD ["version"]
