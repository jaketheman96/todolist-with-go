FROM golang:1.23
 
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /go/bin/godocker

EXPOSE 8080

CMD ["/go/bin/godocker"]