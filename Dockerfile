# FROM golang:alpine as builder
# RUN mkdir /build
# ADD . /build/
# WORKDIR /build
# RUN apk add git
# RUN go get -d -v ./...
# RUN go build -o main .
# FROM alpine
# RUN adduser -S -D -H -h /app appuser
# USER appuser
# COPY --from=builder /build/main /app/
# WORKDIR /app
# CMD ["./main"]

FROM golang:alpine
WORKDIR /go/src/app
COPY . .
RUN apk add git
RUN go get -d -v ./...
CMD ["go","run","main.go"]

EXPOSE 9000
