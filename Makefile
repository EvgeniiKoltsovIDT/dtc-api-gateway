.PHONY: build
build:
	go build -o ./build/gateway ./cmd/gateway/main.go

.PHONY: test
test:
	go test -v -coverprofile coverage.out ./...

.PHONY: build-prof
build-prof: build
	go tool pprof -http:9001 ./build/gateway ./app.pprof

.PHONY: build-run
build-run: build
	./build/gateway

.PHONY: dc
dc:
	docker-compose up --remove-orphans --build

.PHONY: gen
gen:
	go generate ./...

.PHONY: test-html
test-html: test
	go tool cover -html=coverage.out

.PHONY: run
run:
	go run -race ./cmd/gateway/main.go

lint:
	golangci-lint run