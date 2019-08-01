# build
FROM            golang:1.12-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/cryptoguesser
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimalist runtime
FROM            alpine:3.10
COPY            --from=builder /go/bin/cryptoguesser /bin/
ENTRYPOINT      ["/bin/cryptoguesser"]
CMD             []
