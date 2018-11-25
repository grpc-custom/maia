package generator

import (
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	maia "github.com/grpc-custom/maia/proto"
)

type Form struct {
	Name        string
	Title       string
	Description string
	Schema      *Schema
}

func NewForm() *Form {
	return &Form{
		Schema: newSchema(),
	}
}

func (f *Form) AddProperty(field *descriptor.FieldDescriptorProto) error {
	opts, err := f.formOptions(field)
	if err != nil {
		return err
	}
	required := f.requiredOption(field)

	key := field.GetName()
	label := opts.GetLabel()
	if label == "" {
		label = key
	}

	attrType, attrStyle := f.convAttrs(opts)
	s := &Schema{
		Key:         key,
		Title:       label,
		Description: opts.GetDescription(),
		Index:       field.GetNumber(),
		Type:        f.convTypes(field.GetType(), opts.GetNullable()),
		Required:    required,
	}
	if attrType != "" || attrStyle != "" {
		s.Attrs = &FormAttrs{}
		if attrType != "" {
			s.Attrs.Type = attrType
		}
		if attrStyle != "" {
			s.Attrs.Style = attrStyle
		}
	}
	f.Schema.Properties = append(f.Schema.Properties, s)

	return nil
}

func (f *Form) convTypes(typ descriptor.FieldDescriptorProto_Type, nullable bool) []string {
	types := make([]string, 1)
	switch typ {
	case
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		types[0] = Integer
	case
		descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:
		types[0] = Number
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		types[0] = Boolean
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		types[0] = String
	}
	if nullable {
		types = append(types, Null)
	}
	return types
}

func (f *Form) convAttrs(form *maia.Form) (string, string) {
	typ, style := "", ""
	if form == nil {
		return typ, style
	}

	switch form.Type {
	case maia.FormType_HIDDEN:
		typ = "hidden"
	case maia.FormType_TEXT:
		typ = "text"
	case maia.FormType_TEXTAREA:
		style = "textarea"
	case maia.FormType_PASSWORD:
		typ = "password"
	}

	return typ, style
}

func (f *Form) formOptions(field *descriptor.FieldDescriptorProto) (*maia.Form, error) {
	if !proto.HasExtension(field.Options, maia.E_Form) {
		return nil, nil
	}
	ext, err := proto.GetExtension(field.Options, maia.E_Form)
	if err != nil {
		return nil, err
	}
	formOpt, ok := ext.(*maia.Form)
	if !ok {
		return nil, nil
	}
	return formOpt, nil
}

func (f *Form) requiredOption(field *descriptor.FieldDescriptorProto) bool {
	return proto.HasExtension(field.Options, maia.E_Required)
}

func (f *Form) GenCode() string {
	required := make([]string, 0)
	buf := new(strings.Builder)
	buf.Grow(300)
	buf.WriteString("{schema:{")
	buf.WriteString("title:'" + f.Title + "',")
	buf.WriteString("description:'" + f.Description + "',")
	buf.WriteString("type:'object',")
	buf.WriteString("properties:{")
	for i, property := range f.Schema.Properties {
		if property.Required {
			required = append(required, property.Key)
		}
		// type
		buf.WriteString(property.Key + ":{")
		if len(property.Type) == 1 {
			buf.WriteString("type:'" + property.Type[0] + "',")
		} else {
			buf.WriteString("type:[")
			for i, typ := range property.Type {
				buf.WriteString("'" + typ + "'")
				if len(property.Type)-1 > i {
					buf.WriteString(",")
				}
			}
			buf.WriteString("],")
		}
		// title
		buf.WriteString("title:'" + property.Title + "'")
		// attrs
		if property.Attrs != nil {
			buf.WriteString(",attrs:{")
			if property.Attrs.Type != "" {
				buf.WriteString("type:'" + property.Attrs.Type + "'")
			}
			if property.Attrs.Type != "" && property.Attrs.Style != "" {
				buf.WriteString(",")
			}
			if property.Attrs.Style != "" {
				buf.WriteString("style:'" + property.Attrs.Style + "'")
			}
			buf.WriteString("}")
		}
		buf.WriteString("}")
		if len(f.Schema.Properties)-1 > i {
			buf.WriteString(",")
		}
	}
	buf.WriteString("},")
	buf.WriteString("required:[")
	for i, req := range required {
		buf.WriteString("'" + req + "'")
		if len(required)-1 > i {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	buf.WriteString("}}")

	return buf.String()
}

const (
	Null    = "null"
	Integer = "integer"
	Number  = "number"
	String  = "string"
	Boolean = "boolean"
	Object  = "object"
	Array   = "array"
)

func newSchema() *Schema {
	return &Schema{
		Properties: make([]*Schema, 0),
	}
}

type Schema struct {
	Key         string
	Title       string
	Description string
	Index       int32
	Type        []string
	Required    bool
	Properties  []*Schema
	Attrs       *FormAttrs
}

type FormAttrs struct {
	Type  string
	Style string
}
