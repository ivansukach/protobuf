regenerate:
	go install github.com/ivansukach/protobuf/protoc-min-version
	go install github.com/ivansukach/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/ivansukach/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/ivansukach/protobuf/types:. timefail.proto
