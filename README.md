# Movie-Microservices

A few services as examples to show how to write & build microservices in Go that use gRPC, REST & Messaging Queues.

To run:

````
protoc --go_out=./server --go_opt=paths=source_relative \
--go-grpc_out=./server --go-grpc_opt=paths=source_relative \
proto/service.proto
