FROM golang:1.14-alpine AS builder

### add make gcc g++ python if you will use native dependencies
RUN mkdir -p $GOPATH/src/github.com/CristianHCol/go-deviget-minesweeper-apis && \
apk update && apk add --no-cache git openssh-client

WORKDIR $GOPATH/src/github.com/CristianHCol/go-deviget-minesweeper-apis

ENV GO111MODULE=on 
ENV GOPRIVATE=github.com

COPY . .

RUN cp .netrc /root/

RUN go build ./...
RUN go build -o ./ms-api cmd/ms/main.go

FROM alpine

RUN apk --no-cache add ca-certificates bash

WORKDIR /

COPY --from=builder /go/src/github.com/CristianHCol/go-deviget-minesweeper-apis/ms-api /

ARG API_NAME
RUN chmod +x ./ms-api && ln -s "${API_NAME}" /main

CMD [ "/main" ]
