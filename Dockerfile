FROM golang:1.10

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go install

EXPOSE 1323

CMD ["app"]
