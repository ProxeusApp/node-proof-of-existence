FROM golang:1.13

WORKDIR /app
COPY ./main .

CMD ["./main"]