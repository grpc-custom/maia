package codes

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	maia "github.com/grpc-custom/maia/proto"
)

type Component struct {
	Type   maia.ComponentType
	Input  *descriptor.DescriptorProto
	Output *descriptor.DescriptorProto
}
