bildproto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative positionProtocol/positionServer.proto
bild:
	docker build -t position_server:latest .
