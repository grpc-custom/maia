package descriptor

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	maia "github.com/grpc-custom/maia/proto"
)

type file struct {
	*descriptor.FileDescriptorProto
}

func newFile(f *descriptor.FileDescriptorProto) *file {
	return &file{
		FileDescriptorProto: f,
	}
}

type message struct {
	*descriptor.DescriptorProto
	*file
}

func (m *message) FQMN() string {
	components := []string{""}
	if m.file.Package != nil {
		components = append(components, m.file.GetPackage())
	}
	components = append(components, m.GetName())
	return strings.Join(components, ".")
}

type Component struct {
	Type   maia.ComponentType
	Input  *message
	Output *message
}

type Registry struct {
	msgs  map[string]*message
	files map[string]*file

	components map[maia.ComponentType]*Component
}

func NewRegistry() *Registry {
	return &Registry{
		msgs:       make(map[string]*message),
		files:      make(map[string]*file),
		components: make(map[maia.ComponentType]*Component),
	}
}

func (r *Registry) Load(req *plugin.CodeGeneratorRequest) error {
	for _, file := range req.GetProtoFile() {
		r.loadFile(file)
	}

	for _, name := range req.FileToGenerate {
		target, ok := r.files[name]
		if !ok {
			return fmt.Errorf("descripter: no such file: %s", name)
		}
		if err := r.loadService(target); err != nil {
			return err
		}
	}

	return nil
}

func (r *Registry) loadFile(f *descriptor.FileDescriptorProto) {
	file := newFile(f)
	r.files[file.GetName()] = file

	r.registerMsg(file, f.GetMessageType())
}

func (r *Registry) registerMsg(file *file, msgs []*descriptor.DescriptorProto) {
	for _, msg := range msgs {
		m := &message{
			DescriptorProto: msg,
			file:            file,
		}
		r.msgs[m.FQMN()] = m
	}
}

func (r *Registry) loadService(file *file) error {
	for _, svc := range file.GetService() {
		for _, md := range svc.GetMethod() {

			input, err := r.LookupMsg(file.GetPackage(), md.GetInputType())
			if err != nil {
				return err
			}
			// log.Println("[input]", input)
			output, err := r.LookupMsg(file.GetPackage(), md.GetOutputType())
			if err != nil {
				return err
			}
			// log.Println("[output]", input)

			comp, err := extractComponentOptions(md)
			if err != nil {
				return err
			}
			if comp != nil {
				comp.Input = input
				comp.Output = output
				r.components[comp.Type] = comp
			}
		}
	}
	return nil
}

func (r *Registry) LookupMsg(location, name string) (*message, error) {
	if strings.HasPrefix(name, ".") {
		m, ok := r.msgs[name]
		if !ok {
			return nil, fmt.Errorf("no message found: %s", name)
		}
		return m, nil
	}
	if !strings.HasPrefix(location, ".") {
		location = fmt.Sprintf(".%s", location)
	}
	components := strings.Split(location, ".")
	for len(components) > 0 {
		fqmn := strings.Join(append(components, name), ".")
		if m, ok := r.msgs[fqmn]; ok {
			return m, nil
		}
		components = components[:len(components)-1]
	}
	return nil, fmt.Errorf("no message found: %s", name)
}
