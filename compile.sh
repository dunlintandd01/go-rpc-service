#!/usr/bin/env sh

protoc --proto_path=proto \
  --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
  --go_out ./gen --go_opt paths=source_relative \
  hero.proto