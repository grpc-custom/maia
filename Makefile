proto-maia:
	protoc --go_out=${GOPATH}/src ./*.proto

build-debug-maia: proto-maia
	go build -o protoc-gen-debug-maia ./protoc-gen-maia/main.go

proto-debug-maia: build-debug-maia
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-debug-maia \
		--debug-maia_out=hoge=true:. \
		_example/*.proto

proto-js:
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--js_out=import_style=commonjs:. \
		_example/*.proto \
		maia.proto

proto-grpc-web:
	protoc \
		-I=${GOPATH}/src:. \
		--grpc-web_out=import_style=commonjs,mode=grpcweb:. \
		_example/*.proto
