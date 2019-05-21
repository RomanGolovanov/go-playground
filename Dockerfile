FROM golang:1.11 as build

WORKDIR /app
COPY . /app
RUN go install /app

ENTRYPOINT /app/go-playground
EXPOSE 8080
