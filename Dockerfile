FROM golang:alpine as builder

WORKDIR $GOPATH/app

COPY . .

RUN apk update && apk add --no-cache git

RUN go get -d -v ./cmd/
RUN go build -o /go/bin/app cmd/main.go



# FROM scratch
FROM golang:alpine

COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]
