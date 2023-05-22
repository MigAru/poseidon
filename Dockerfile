FROM golang:1.18

WORKDIR /app

COPY . .

EXPOSE 8000

RUN go build -o srv ./cmd/main.go
ENTRYPOINT [ "./srv" ]