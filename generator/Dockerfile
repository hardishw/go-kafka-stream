FROM golang:1.10.2

RUN apt-get update && git clone https://github.com/edenhill/librdkafka.git && \
	cd librdkafka && ./configure --prefix /usr && make && make install && \
	go get -u github.com/confluentinc/confluent-kafka-go/kafka

ADD . /opt/generator-app/

CMD go run /opt/generator-app/kafka-producer.go
