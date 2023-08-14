FROM golang:1.21

COPY . /app

WORKDIR /app/cmd/finance-tracker

RUN go build -o finance-tracker

CMD ["./finance-tracker"]