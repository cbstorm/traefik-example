FROM golang:latest AS build
WORKDIR /udp_svc
COPY . .
RUN go mod download
RUN go build -v -o /bin/udp_svc

FROM ubuntu:latest

WORKDIR /
RUN apt-get update
RUN apt-get install -y ca-certificates

COPY --from=build /bin/udp_svc /bin

CMD ["/bin/udp_svc"]