FROM golang:1.14.4-alpine3.12 as builder
RUN mkdir -p $GOPATH/src/github.com/juicypy/todo_list_service
WORKDIR $GOPATH/src/github.com/juicypy/todo_list_service
COPY . .
RUN go mod download
RUN go build -o $GOPATH/src/github.com/juicypy/todo_list_service/app cmd/server/main.go
RUN go build -o $GOPATH/src/github.com/juicypy/todo_list_service/migrate cmd/migrate/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/juicypy/todo_list_service/app .
COPY --from=builder /go/src/github.com/juicypy/todo_list_service/migrate .
COPY --from=builder /go/src/github.com/juicypy/todo_list_service/migrations /migrations
CMD ["./app"]