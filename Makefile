BROKERS=localhost:9092
TOPIC=product.ingest

build:
	go build -o bin/producer cmd/consumer/main.go
	go build -o bin/consumer cmd/producer/main.go
	go build -o bin/dlq_consumer cmd/dlq_consumer/main.go

clean:
	rm -rf bin/
	
create-topics:
	bash create-topic.sh $(TOPIC)
	bash create-topic.sh $(DLQ_TOPIC)

docker up:
	docker compose up -d

docker down:
	docker compose down

produce:
	go run cmd/producer/main.go

consume:
	go run cmd/consumer/main.go

consume-dlq:
	go run cmd/dlq_consumer/main.go

test:
	go test ./tests/... -v
	