FROM golang AS build

WORKDIR /go/src/postastix-api

ADD . .

RUN go get -v ./...
RUN go build *.go

FROM alpine

WORKDIR /usr/bin

COPY --from=build /go/src/postastix-api/main postastix
RUN chmod a+x postastix

CMD ["postastix", "serve"]
