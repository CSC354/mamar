FROM golang:1.19

RUN mkdir -p /usr/src/mamar
COPY . /usr/src/mamar
COPY testfile /opt/mamar_ports
WORKDIR /usr/src/mamar
RUN go mod tidy
ENTRYPOINT go run cmd/mamar.go /usr/src/mamar
