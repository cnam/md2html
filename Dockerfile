FROM alpine:3.2

RUN apk add --update curl && \
    curl -L https://github.com/cnam/md2html/releases/download/0.1/linux_md2html > md2html && \
    chmod  +x md2html && \
    apk del curl && \
    rm -rf /var/cache/apk/*

WORKDIR /data

ENTRYPOINT ["/md2html"]
