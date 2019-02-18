FROM golang AS build

WORKDIR /go/src/postastix-api

ADD . .

RUN go get -v ./...
RUN go build

FROM alpine

WORKDIR /usr/bin

COPY --from=build /go/src/postastix-api/postastix-api postastix-api
RUN chmod a+x postastix-api

CMD ["postastix-api", "serve"]
