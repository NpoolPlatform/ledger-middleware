FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY ledger-middleware /app

ENTRYPOINT ["/app/ledger-middleware"]

EXPOSE 50420 50421 50422 50423

CMD ["run"]
