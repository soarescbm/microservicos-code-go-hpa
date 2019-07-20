FROM golang:1.12.6-stretch as buildando

WORKDIR /go/src/app

COPY src/app/app.go .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux  go install -a -installsuffix cgo -v ./...

FROM scratch

WORKDIR /bin

COPY --from=buildando /go/bin/app /bin

EXPOSE 80

CMD ["/bin/app"]