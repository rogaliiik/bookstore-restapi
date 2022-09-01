FROM golang:1.16-alpine

WORKDIR /app

COPY . ./
COPY /cmd/.env ./

RUN go mod download


RUN go build -o bookstore /app/cmd/main.go

EXPOSE 8080

CMD [ "./bookstore" ]
