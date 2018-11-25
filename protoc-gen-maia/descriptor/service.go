package descriptor

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	maia "github.com/grpc-custom/maia/proto"
)

func extractComponentOptions(md *descriptor.MethodDescriptorProto) (*Component, error) {
	if md.Options == nil {
		return nil, nil
	}
	if !proto.HasExtension(md.Options, maia.E_Component) {
		return nil, nil
	}
	ext, err := proto.GetExtension(md.Options, maia.E_Component)
	if err != nil {
		return nil, err
	}
	opts, ok := ext.(*maia.Component)
	if !ok {
		return nil, fmt.Errorf("extension is %T; want an maia.Component", ext)
	}
	comp := &Component{
		Type: opts.Type,
	}
	return comp, nil
}
