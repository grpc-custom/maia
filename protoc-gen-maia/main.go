package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	maia "github.com/grpc-custom/maia/proto"
	desc "github.com/grpc-custom/maia/protoc-gen-maia/descriptor"
	"github.com/grpc-custom/maia/protoc-gen-maia/generator"
)

func parseReq(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	req := new(plugin.CodeGeneratorRequest)
	if err := proto.Unmarshal(buf, req); err != nil {
		return nil, err
	}
	return req, nil
}

func processReq(req *plugin.CodeGeneratorRequest) {
	files := make(map[string]*descriptor.FileDescriptorProto)
	for _, file := range req.ProtoFile {
		files[file.GetName()] = file
	}
	for _, name := range req.FileToGenerate {
		target := files[name]

		pkg := target.GetPackage()
		log.Println("[pkg]", pkg)
		if proto.HasExtension(target.GetOptions(), maia.E_Version) {
			ext, err := proto.GetExtension(target.Options, maia.E_Version)
			if err != nil {
				log.Fatal(err)
			}
			version, ok := ext.(*string)
			log.Println("[version]", *version, ok)
		}

		for _, srv := range target.GetService() {
			log.Println("[service]", srv.GetName())
			for _, method := range srv.GetMethod() {
				log.Println("[InputType]", method.GetInputType())
				log.Println("[OutputType]", method.GetOutputType())
				if proto.HasExtension(method.GetOptions(), maia.E_Component) {
					comp, err := proto.GetExtension(method.GetOptions(), maia.E_Component)
					if err != nil {
						log.Fatal(err)
					}
					c, ok := comp.(*maia.Component)
					log.Println("[method]", c, ok)
				}
			}
		}

		for _, msg := range target.GetMessageType() {
			log.Println("[message]", "."+pkg+"."+msg.GetName())
			for _, field := range msg.GetField() {
				opt := field.Options

				// required
				if proto.HasExtension(opt, maia.E_Required) {

				}

				// form
				if proto.HasExtension(opt, maia.E_Form) {
					ext, err := proto.GetExtension(opt, maia.E_Form)
					if err != nil {
						log.Fatal(err)
					}
					if form, ok := ext.(*maia.Form); ok {
						log.Println("[maia.form.type]", form.Type)
						log.Println("[maia.form.description]", form.Description)
						log.Println("[maia.form.disabled]", form.Disabled)
						log.Println("[maia.form.pattern]", form.Pattern)
						log.Println("[maia.form.format]", form.Format)
					}
				}

				// autocomplete
				if proto.HasExtension(opt, maia.E_Autocomplete) {
					ext, err := proto.GetExtension(opt, maia.E_Autocomplete)
					if err != nil {
						log.Fatal(err)
					}
					if autocomplete, ok := ext.(*maia.Autocomplete); ok {
						log.Println("[maia.autocomplete.rpc]", autocomplete.Rpc)
						log.Println("[maia.autocomplete.service]", autocomplete.Service)
					}
				}

				// table
				if proto.HasExtension(opt, maia.E_Table) {
					ext, err := proto.GetExtension(opt, maia.E_Table)
					if err != nil {
						log.Fatal(err)
					}
					if table, ok := ext.(*maia.Table); ok {
						for _, col := range table.Columns {
							log.Println("[maia.table.columns]", col.Label, col.Field)
						}
					}
				}
			}
		}
	}
}

func emitFiles(out []*plugin.CodeGeneratorResponse_File) {
	emitResp(&plugin.CodeGeneratorResponse{
		File: out,
	})
}

func emitResp(resp *plugin.CodeGeneratorResponse) {
	buf, err := proto.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stdout.Write(buf); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	req, err := parseReq(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	reg := desc.NewRegistry()
	if err := reg.Load(req); err != nil {
		log.Fatal(err)
	}

	gen := generator.NewGenerator(reg)
	formFile, err := gen.Form()
	if err != nil {
		log.Fatal(err)
	}

	emitFiles([]*plugin.CodeGeneratorResponse_File{
		formFile,
	})
}
