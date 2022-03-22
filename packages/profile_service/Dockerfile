<<<<<<< HEAD
FROM golang:1.18.0

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get
=======
FROM golang:1.18.0-alpine as builder

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN apk --no-cache add git alpine-sdk build-base gcc

RUN go get
     # && go get golang.org/x/tools/cmd/cover \
    # && go get github.com/mattn/goveralls

RUN go build -o build ./main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/build .
COPY --from=builder /app/config/*.yaml ./config/
CMD ["./build"]
>>>>>>> 18c1201 (new commit)
