#!/bin/bash
goctl rpc protoc -I=./ msgcallback.proto -v --go_out=.. --go-grpc_out=..  --zrpc_out=.. --style=goZero --home ../../../../../goctl/home
