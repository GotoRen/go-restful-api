FROM golang:alpine as builder

WORKDIR /app
ADD . /app/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -installsuffix cgo -o hello main.go


FROM alpine:latest

COPY --from=builder /app/hello /bin/hello

CMD ["/bin/hello"]
