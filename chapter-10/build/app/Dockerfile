FROM ubuntu:22.04

ENV TZ=Asia/Tokyo

RUN apt update && \
    apt install -y wget

RUN wget https://go.dev/dl/go1.17.5.linux-amd64.tar.gz && \
    tar -C /usr/local/ -xvf go1.17.5.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

WORKDIR /usr/src/ws-s

COPY go.mod .
COPY go.sum .

COPY . .

RUN go build -o bin/ws-s

ENTRYPOINT [ "bin/ws-s" ]

EXPOSE 8080
