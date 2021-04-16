regenerate:
	go install github.com/ivansukach/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/ivansukach/protobuf/protobuf/ \
	--gogo_out=. events.proto 
