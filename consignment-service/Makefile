build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/maurofran/shipper/consignment-service \
		proto/consignment/consignment.proto
	docker build -t consignment-service .

run:
	docker run \
		-p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		consignment-service