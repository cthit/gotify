.PHONY: setup
setup: gen-setup gen
	go mod download

.PHONY: gen-setup
gen-setup:
	docker build --target protocGenerator -t gotify-protoc-generator .

.PHONY: gen
gen:
	docker run -v `pwd`:/app  gotify-protoc-generator ./scripts/protoc-gen.sh

.PHONY: build
build: gen
	go build -o gotify-bin ./cmd/gotify

.PHONY: run
run: gen
	go run ./cmd/gotify

.PHONY: dev
dev:
	docker-compose up

.PHONY: clean
clean:
	rm -rf pkg/api
	git restore api/swagger

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

.PHONY: lint-docker
lint-docker:
	docker run --rm -v `pwd`:/app -w /app golangci/golangci-lint:v1.31.0-alpine golangci-lint run

.PHONY: test
test:
	go test ./...