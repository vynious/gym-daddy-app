gen:
	protoc --proto_path=proto --go_out=pb --go-grpc_out=pb ./proto/queue/*.proto
	protoc --proto_path=proto --go_out=pb --go-grpc_out=pb ./proto/notification/*.proto



clean:
	rm pb/proto_files/queue/*.go
	rm pb/proto_files/notification/*.go