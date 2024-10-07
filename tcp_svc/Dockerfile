FROM golang:latest AS build
WORKDIR /tcp_svc
COPY . .
RUN go mod download
RUN go build -v -o /bin/tcp_svc

FROM ubuntu:latest

WORKDIR /
RUN apt-get update
RUN apt-get install -y ca-certificates

COPY --from=build /bin/tcp_svc /bin

CMD ["/bin/tcp_svc"]