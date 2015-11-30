FROM scratch

COPY ./md2html /usr/local/bin

VOLUME ["/data", "/cache"]

WORKDIR /data

ENTRYPOINT ["md2html"]
