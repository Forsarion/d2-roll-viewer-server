FROM golang:1.11

WORKDIR $GOPATH/github.com/Berk0ld/d2-roll-viewer-server
COPY . .

RUN go get -d -v ./...
RUN go build -v -x

EXPOSE 8080

ENTRYPOINT [ "./d2-roll-viewer-server" ]
