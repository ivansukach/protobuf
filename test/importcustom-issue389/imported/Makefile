regenerate:
	go install github.com/ivansukach/protobuf/protoc-gen-gogo
	go install github.com/ivansukach/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --proto_path=../../../../../../:../../../protobuf/:. --gogo_out=. a.proto