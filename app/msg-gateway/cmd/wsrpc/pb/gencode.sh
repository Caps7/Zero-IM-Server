#!/bin/bash
protoc *.proto --proto_path ./ --go_out .. --go-grpc_out ..
