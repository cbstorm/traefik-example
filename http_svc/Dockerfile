FROM golang:latest AS build
WORKDIR /http_svc
COPY . .
RUN go mod download
RUN go build -v -o /bin/http_svc

FROM ubuntu:latest

WORKDIR /
RUN apt-get update
RUN apt-get install -y ca-certificates

COPY --from=build /bin/http_svc /bin

CMD ["/bin/http_svc"]