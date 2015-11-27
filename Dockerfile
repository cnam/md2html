FROM scratch

COPY ./md2html ./

ENTRYPOINT ["./md2html"]
