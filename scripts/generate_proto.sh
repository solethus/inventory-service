#!/bin/bash

# generate_proto.sh

# Exit immediately if a command exits with a non-zero status.
set -e

# Define variables
PROTO_DIR="./proto/inventory"
GO_OUT_DIR="./proto/inventory"
PROTO_FILE="inventory.proto"

# Check if protoc is installed
if ! command -v protoc &> /dev/null
then
    echo "protoc could not be found. Please install Protocol Buffers compiler."
    exit 1
fi

# Check if required Go plugins are installed
if ! command -v protoc-gen-go &> /dev/null || ! command -v protoc-gen-go-grpc &> /dev/null
then
    echo "Required protoc plugins are missing. Installing..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi

# Create output directory if it doesn't exist
mkdir -p ${GO_OUT_DIR}

# Generate Go code
protoc --proto_path=${PROTO_DIR} \
       --go_out=${GO_OUT_DIR} --go_opt=paths=source_relative \
       --go-grpc_out=${GO_OUT_DIR} --go-grpc_opt=paths=source_relative \
       ${PROTO_DIR}/${PROTO_FILE}

echo "Go code generated successfully from ${PROTO_FILE}"

# Optionally, run go fmt on generated files
go fmt ${GO_OUT_DIR}/*.pb.go

echo "Generated files formatted"

# Optionally, add a reminder about regenerating mocks if you're using them
echo "Remember to regenerate any mocks if you've changed the service interface"