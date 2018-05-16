FROM golang:1.10.2

RUN apt-get update && git clone https://github.com/edenhill/librdkafka.git && \
	cd librdkafka && ./configure --prefix /usr && make && make install && \
	go get -u github.com/confluentinc/confluent-kafka-go/kafka

ADD . /opt/counter-app/

WORKDIR /opt/counter-app/

RUN go build counter.go

ENTRYPOINT ["/opt/counter-app/counter"]
