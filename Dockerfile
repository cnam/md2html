FROM scratch

COPY ./md2html ./

VOLUME ["/data", "/cache"]

WORKDIR /data

ENTRYPOINT ["./md2html"]
