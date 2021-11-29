gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative priceProtocol/priceServer.proto
bild:
	docker build -t price_server:latest .
