FROM golang

ADD . /go/src/github.com/dewadg/postastix-api

WORKDIR /go/src/github.com/dewadg/postastix-api

RUN go get ./...
RUN go install
RUN /go/bin/postastix-api migrate

ENTRYPOINT /go/bin/postastix-api serve