FROM golang:1.11 as build
WORKDIR /app
COPY . /app
RUN go install && CGO_ENABLED=1 GOOS=linux  go build

FROM alpine:latest  
RUN apk --update upgrade \
    && apk add sqlite \
    && mkdir /lib64 \
    && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 \
    && apk --no-cache add ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=build /app/go-playground .
ENTRYPOINT /app/go-playground
EXPOSE 8080
