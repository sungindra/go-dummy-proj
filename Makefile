.PHONY: vendor
vendor:
	go mod vendor && go mod tidy && go mod vendor

.PHONY: test
test:
	go test ./...

.PHONY: server/start
server/start:
	docker-compose -f docker-compose.yaml up -d

.PHONY: server/stop
server/stop:
	docker-compose -f docker-compose.yaml down
