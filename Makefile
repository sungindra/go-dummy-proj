.PHONY: vendor
vendor:
	go mod vendor && go mod tidy && go mod vendor

.PHONY: test
test:
	go test ./...