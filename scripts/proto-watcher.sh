#!/bin/sh
(./scripts/protoc-gen.sh && echo "Generated!" ) || true
reflex -r '\.proto$' -- sh -c './scripts/protoc-gen.sh && echo "Generated!"'