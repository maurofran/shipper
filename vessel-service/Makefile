build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/maurofran/shipper/vessel-service \
		proto/vessel/vessel.proto
	docker build -t vessel-service .

run:
	docker run -t 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service