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
	rm -rf api/swagger
