FROM golang:1.22.6

COPY . .

RUN go build .

CMD ./evert-http