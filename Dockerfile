FROM golang:1.18 as build

RUN mkdir /go/src/hg
COPY . /go/src/hg

WORKDIR /go/src/hg

RUN go get -d -v ./...

RUN go build -o /go/bin/hg ./cmd/api/main.go

FROM gcr.io/distroless/base
COPY --from=build /go/bin/hg /
CMD ["/hg"]