run:
	docker-compose up -d
	go run main.go

dev:
	docker-compose up -d
	docker-compose exec -T redis redis-cli flushall
	REDIS_ADDR=localhost:6379 go run main.go

test:
	docker-compose up -d
	docker-compose exec -T redis redis-cli flushall
	REDIS_ADDR=localhost:6379 go test ./...

stop:
	docker-compose down
