bildproto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protocolPosition/serverPosition.proto
bild:
	docker build -t position_server:latest .
