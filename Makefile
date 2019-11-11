protopath = pkg/api/proto
bindir = ./bin

all: proto build

build:
	go build -i -v -o $(bindir)/server ./cmd/

proto:
	protoc --go_out=plugins=grpc:$(protopath) --proto_path=$(protopath) charbalancer.proto

clean:
	rm $(bindir)/server
