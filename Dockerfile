FROM golang:1.16-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o bookstore ./cmd/main.go

EXPOSE 8080

CMD [ "./bookstore" ]
