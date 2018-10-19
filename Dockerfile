FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN apk add git

RUN go get -d -v ./...

CMD ["go", "run main.go"]