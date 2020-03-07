#!/bin/sh
./scripts/protoc-gen.sh || true
reflex -r '\.proto$' -- ./scripts/protoc-gen.sh