PKG=github.com/smblab/candidate
GO_PLUGIN=bin/protoc-gen-go
GO_PROTOBUF_REPO=github.com/golang/protobuf
GO_PLUGIN_PKG=$(GO_PROTOBUF_REPO)/protoc-gen-go
GO_PTYPES_ANY_PKG=$(GO_PROTOBUF_REPO)/ptypes/any
GOOGLEAPIS_DIR=third_party/googleapis
PKGMAP=Mgoogle/protobuf/descriptor.proto=$(GO_PLUGIN_PKG)/descriptor,Mapi/transactions.proto=$(PKG)/pkg/transactions/pb,Mapi/balance.proto=$(PKG)/pkg/balance/pb

PROTOC_INC_PATH=$(dir $(shell which protoc))../include

.SUFFIXES: .go .proto

$(GO_PLUGIN):
	go get $(GO_PLUGIN_PKG)
	go build -o $@ $(GO_PLUGIN_PKG)

generate: $(GO_PLUGIN)
	protoc -I $(PROTOC_INC_PATH) -I$(GOOGLEAPIS_DIR) --plugin=$(GO_PLUGIN) -I. --go_out=$(PKGMAP),plugins=grpc:. api/transactions.proto
	protoc -I $(PROTOC_INC_PATH) -I$(GOOGLEAPIS_DIR) --plugin=$(GO_PLUGIN) -I. --go_out=$(PKGMAP),plugins=grpc:. api/balance.proto

build: cmd/candidate
	go build -o bin/candidate ./cmd/candidate
