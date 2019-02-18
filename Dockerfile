FROM golang AS build

WORKDIR /go/src/github.com/dewadg/postastix-api

ADD . .

RUN go get -v ./...
RUN go build *.go

FROM alpine

WORKDIR /usr/bin

COPY --from=build /go/src/github.com/dewadg/postastix-api/app .
RUN chmod a+x app

CMD ["app", "serve"]
