// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bq_field.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var E_Require = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1022,
	Name:          "gen_bq_schema.require",
	Tag:           "varint,1022,opt,name=require",
	Filename:      "bq_field.proto",
}

var E_TypeOverride = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1023,
	Name:          "gen_bq_schema.type_override",
	Tag:           "bytes,1023,opt,name=type_override,json=typeOverride",
	Filename:      "bq_field.proto",
}

var E_Ignore = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1024,
	Name:          "gen_bq_schema.ignore",
	Tag:           "varint,1024,opt,name=ignore",
	Filename:      "bq_field.proto",
}

func init() {
	proto.RegisterExtension(E_Require)
	proto.RegisterExtension(E_TypeOverride)
	proto.RegisterExtension(E_Ignore)
}

func init() { proto.RegisterFile("bq_field.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2a, 0x8c, 0x4f,
	0xcb, 0x4c, 0xcd, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0x4f, 0xcd, 0x8b,
	0x4f, 0x2a, 0x8c, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x94, 0x52, 0x48, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x07, 0x4b, 0x26, 0x95, 0xa6, 0xe9, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16, 0x94,
	0xe4, 0x17, 0x41, 0x34, 0x58, 0x59, 0x70, 0xb1, 0x17, 0xa5, 0x16, 0x96, 0x66, 0x16, 0xa5, 0x0a,
	0xc9, 0xea, 0x41, 0x54, 0xeb, 0xc1, 0x54, 0xeb, 0xb9, 0x81, 0x4c, 0xf6, 0x2f, 0x28, 0xc9, 0xcc,
	0xcf, 0x2b, 0x96, 0xf8, 0xc7, 0xae, 0xc0, 0xa8, 0xc1, 0x11, 0x04, 0x53, 0x6e, 0xe5, 0xcc, 0xc5,
	0x5b, 0x52, 0x59, 0x90, 0x1a, 0x9f, 0x5f, 0x96, 0x5a, 0x54, 0x94, 0x99, 0x42, 0x50, 0xff, 0x7f,
	0x90, 0x7e, 0xce, 0x20, 0x1e, 0x90, 0x26, 0x7f, 0xa8, 0x1e, 0x2b, 0x33, 0x2e, 0xb6, 0xcc, 0xf4,
	0xbc, 0x7c, 0xc2, 0xb6, 0x37, 0x70, 0x80, 0x6d, 0x87, 0xaa, 0x76, 0xe2, 0x88, 0x62, 0x03, 0xab,
	0x2b, 0x4e, 0x82, 0xd0, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x93, 0x0d, 0x9c, 0x0a,
	0x01, 0x00, 0x00,
}
