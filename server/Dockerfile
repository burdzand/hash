FROM golang:1.9 as builder

WORKDIR /go/src/github.com/burdzand/hash/server/
COPY . /go/src/github.com/burdzand/hash/server/

RUN ./install-deps.sh
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/burdzand/hash/server/main .
COPY --from=builder /go/src/github.com/burdzand/hash/server/config .
COPY --from=builder /go/src/github.com/burdzand/hash/server/wait-for .
EXPOSE 3000
CMD ["./main"]