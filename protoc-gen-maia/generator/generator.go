package generator

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	maia "github.com/grpc-custom/maia/proto"
	maiaDesc "github.com/grpc-custom/maia/protoc-gen-maia/descriptor"
)

type Generator struct {
	register *maiaDesc.Registry
}

func NewGenerator(reg *maiaDesc.Registry) *Generator {
	return &Generator{
		register: reg,
	}
}

func (g *Generator) Form() (*plugin.CodeGeneratorResponse_File, error) {
	comp, err := g.register.Component(maia.ComponentType_CREATE)
	if err != nil {
		return nil, err
	}

	form := NewForm()

	// read input
	for _, inputField := range comp.Input.GetField() {
		if inputField.GetType() != descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			continue // todo message型以外をチェックする
		}

		msg, err := g.register.LookupMsg(comp.Input.GetPackage(), inputField.GetTypeName())
		if err != nil {
			return nil, err
		}

		form.Name = msg.GetName()

		if proto.HasExtension(msg.Options, maia.E_FormOption) {
			ext, err := proto.GetExtension(msg.Options, maia.E_FormOption)
			if err != nil {
				return nil, err
			}
			formOpt, ok := ext.(*maia.FormOption)
			if !ok {
				return nil, fmt.Errorf("invalid form option")
			}
			form.Title = formOpt.GetTitle()
			form.Description = formOpt.GetDescription()
		}

		for _, field := range msg.GetField() {
			if err := form.AddProperty(field); err != nil {
				return nil, err
			}
		}
	}

	file := &plugin.CodeGeneratorResponse_File{
		Name:    proto.String(strings.ToLower(form.Name) + ".pb.maia.js"),
		Content: proto.String(form.GenCode()),
	}
	return file, nil
}
