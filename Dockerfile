FROM golang:1.13

WORKDIR /app
COPY ./artifacts/node-proof-of-existence .

CMD ["./node-proof-of-existence"]