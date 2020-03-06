.PHONY: gen
gen:
	./scripts/protoc-gen.sh

.PHONY: clean
clean:
	rm -r pkg/api
	rm -r api/swagger