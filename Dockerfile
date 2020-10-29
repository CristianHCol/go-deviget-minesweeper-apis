FROM golang:1.14-alpine AS builder

### add make gcc g++ python if you will use native dependencies
RUN mkdir -p $GOPATH/src/gitlab.com/fixit/go-apis && \
apk update && apk add --no-cache git openssh-client

WORKDIR $GOPATH/src/gitlab.com/fixit/go-apis

ENV GO111MODULE=on 
ENV GOPRIVATE=gitlab.com

COPY . .

RUN cp .netrc /root/

RUN go build ./...
RUN go build -o ./user-api cmd/user/main.go

FROM alpine

RUN apk --no-cache add ca-certificates bash

WORKDIR /

COPY --from=builder /go/src/gitlab.com/fixit/go-apis/user-api /

ARG API_NAME
RUN chmod +x ./user-api && ln -s "${API_NAME}" /main

CMD [ "/main" ]
