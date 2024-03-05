run:
	go run main.go

build:


gen:
	protoc --proto_path=proto --go_out=pb --go-grpc_out=pb ./proto/notification/*.proto
	protoc --proto_path=proto --go_out=pb --go-grpc_out=pb ./proto/queue/*.proto



clean:
	rm pb/proto_files/notification/*.go
	rm pb/proto_files/queue/*.go