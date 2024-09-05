FROM golang:1.22.6

COPY . .

RUN go build ./cmd/evert-http

CMD ./evert-http