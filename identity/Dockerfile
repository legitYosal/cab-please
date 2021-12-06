FROM golang:1.17-alpine

WORKDIR /usr/src/app

RUN apk update \ 
    && apk add gcc g++ make build-base linux-headers

COPY ./go.mod ./go.sum /usr/src/app/
RUN go mod download

COPY . .

ENV CGO_CFLAGS="-g -O2 -Wno-return-local-addr"

RUN go build -o serverBuild .

CMD ["./serverBuild"]