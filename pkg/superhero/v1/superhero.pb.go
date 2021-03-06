// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: superhero/v1/superhero.proto

package superherov1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Genero int32

const (
	Genero_mujer  Genero = 0
	Genero_hombre Genero = 1
)

// Enum value maps for Genero.
var (
	Genero_name = map[int32]string{
		0: "mujer",
		1: "hombre",
	}
	Genero_value = map[string]int32{
		"mujer":  0,
		"hombre": 1,
	}
)

func (x Genero) Enum() *Genero {
	p := new(Genero)
	*p = x
	return p
}

func (x Genero) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Genero) Descriptor() protoreflect.EnumDescriptor {
	return file_superhero_v1_superhero_proto_enumTypes[0].Descriptor()
}

func (Genero) Type() protoreflect.EnumType {
	return &file_superhero_v1_superhero_proto_enumTypes[0]
}

func (x Genero) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Genero.Descriptor instead.
func (Genero) EnumDescriptor() ([]byte, []int) {
	return file_superhero_v1_superhero_proto_rawDescGZIP(), []int{0}
}

type Casa int32

const (
	Casa_marvel Casa = 0
	Casa_dc     Casa = 1
)

// Enum value maps for Casa.
var (
	Casa_name = map[int32]string{
		0: "marvel",
		1: "dc",
	}
	Casa_value = map[string]int32{
		"marvel": 0,
		"dc":     1,
	}
)

func (x Casa) Enum() *Casa {
	p := new(Casa)
	*p = x
	return p
}

func (x Casa) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Casa) Descriptor() protoreflect.EnumDescriptor {
	return file_superhero_v1_superhero_proto_enumTypes[1].Descriptor()
}

func (Casa) Type() protoreflect.EnumType {
	return &file_superhero_v1_superhero_proto_enumTypes[1]
}

func (x Casa) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Casa.Descriptor instead.
func (Casa) EnumDescriptor() ([]byte, []int) {
	return file_superhero_v1_superhero_proto_rawDescGZIP(), []int{1}
}

type Atributos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Debilidad string `protobuf:"bytes,1,opt,name=debilidad,proto3" json:"debilidad,omitempty"`
	Ataque    int32  `protobuf:"varint,2,opt,name=ataque,proto3" json:"ataque,omitempty"`
	Vida      int32  `protobuf:"varint,3,opt,name=vida,proto3" json:"vida,omitempty"`
	Genero    Genero `protobuf:"varint,4,opt,name=genero,proto3,enum=superhero.v1.Genero" json:"genero,omitempty"`
}

func (x *Atributos) Reset() {
	*x = Atributos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superhero_v1_superhero_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Atributos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Atributos) ProtoMessage() {}

func (x *Atributos) ProtoReflect() protoreflect.Message {
	mi := &file_superhero_v1_superhero_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Atributos.ProtoReflect.Descriptor instead.
func (*Atributos) Descriptor() ([]byte, []int) {
	return file_superhero_v1_superhero_proto_rawDescGZIP(), []int{0}
}

func (x *Atributos) GetDebilidad() string {
	if x != nil {
		return x.Debilidad
	}
	return ""
}

func (x *Atributos) GetAtaque() int32 {
	if x != nil {
		return x.Ataque
	}
	return 0
}

func (x *Atributos) GetVida() int32 {
	if x != nil {
		return x.Vida
	}
	return 0
}

func (x *Atributos) GetGenero() Genero {
	if x != nil {
		return x.Genero
	}
	return Genero_mujer
}

type Superhero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre    string     `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Habilidad string     `protobuf:"bytes,2,opt,name=habilidad,proto3" json:"habilidad,omitempty"`
	Atributos *Atributos `protobuf:"bytes,3,opt,name=atributos,proto3" json:"atributos,omitempty"`
	Casa      Casa       `protobuf:"varint,4,opt,name=casa,proto3,enum=superhero.v1.Casa" json:"casa,omitempty"`
}

func (x *Superhero) Reset() {
	*x = Superhero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superhero_v1_superhero_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Superhero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Superhero) ProtoMessage() {}

func (x *Superhero) ProtoReflect() protoreflect.Message {
	mi := &file_superhero_v1_superhero_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Superhero.ProtoReflect.Descriptor instead.
func (*Superhero) Descriptor() ([]byte, []int) {
	return file_superhero_v1_superhero_proto_rawDescGZIP(), []int{1}
}

func (x *Superhero) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Superhero) GetHabilidad() string {
	if x != nil {
		return x.Habilidad
	}
	return ""
}

func (x *Superhero) GetAtributos() *Atributos {
	if x != nil {
		return x.Atributos
	}
	return nil
}

func (x *Superhero) GetCasa() Casa {
	if x != nil {
		return x.Casa
	}
	return Casa_marvel
}

var File_superhero_v1_superhero_proto protoreflect.FileDescriptor

var file_superhero_v1_superhero_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x83, 0x01, 0x0a,
	0x09, 0x41, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65,
	0x62, 0x69, 0x6c, 0x69, 0x64, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x65, 0x62, 0x69, 0x6c, 0x69, 0x64, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x74, 0x61, 0x71,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x74, 0x61, 0x71, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x64, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x76, 0x69, 0x64, 0x61, 0x12, 0x2c, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x6f, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x09, 0x53, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x64, 0x61, 0x64, 0x12, 0x35, 0x0a, 0x09, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x6f, 0x73, 0x52, 0x09, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x73, 0x12, 0x26, 0x0a,
	0x04, 0x63, 0x61, 0x73, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x73, 0x61, 0x52,
	0x04, 0x63, 0x61, 0x73, 0x61, 0x2a, 0x1f, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x6f, 0x12,
	0x09, 0x0a, 0x05, 0x6d, 0x75, 0x6a, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x68, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x10, 0x01, 0x2a, 0x1a, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x61, 0x12, 0x0a,
	0x0a, 0x06, 0x6d, 0x61, 0x72, 0x76, 0x65, 0x6c, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x64, 0x63,
	0x10, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x73, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x75, 0x70, 0x65, 0x72, 0x68, 0x65, 0x72, 0x6f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_superhero_v1_superhero_proto_rawDescOnce sync.Once
	file_superhero_v1_superhero_proto_rawDescData = file_superhero_v1_superhero_proto_rawDesc
)

func file_superhero_v1_superhero_proto_rawDescGZIP() []byte {
	file_superhero_v1_superhero_proto_rawDescOnce.Do(func() {
		file_superhero_v1_superhero_proto_rawDescData = protoimpl.X.CompressGZIP(file_superhero_v1_superhero_proto_rawDescData)
	})
	return file_superhero_v1_superhero_proto_rawDescData
}

var file_superhero_v1_superhero_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_superhero_v1_superhero_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_superhero_v1_superhero_proto_goTypes = []interface{}{
	(Genero)(0),       // 0: superhero.v1.Genero
	(Casa)(0),         // 1: superhero.v1.Casa
	(*Atributos)(nil), // 2: superhero.v1.Atributos
	(*Superhero)(nil), // 3: superhero.v1.Superhero
}
var file_superhero_v1_superhero_proto_depIdxs = []int32{
	0, // 0: superhero.v1.Atributos.genero:type_name -> superhero.v1.Genero
	2, // 1: superhero.v1.Superhero.atributos:type_name -> superhero.v1.Atributos
	1, // 2: superhero.v1.Superhero.casa:type_name -> superhero.v1.Casa
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_superhero_v1_superhero_proto_init() }
func file_superhero_v1_superhero_proto_init() {
	if File_superhero_v1_superhero_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_superhero_v1_superhero_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Atributos); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_superhero_v1_superhero_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Superhero); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_superhero_v1_superhero_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_superhero_v1_superhero_proto_goTypes,
		DependencyIndexes: file_superhero_v1_superhero_proto_depIdxs,
		EnumInfos:         file_superhero_v1_superhero_proto_enumTypes,
		MessageInfos:      file_superhero_v1_superhero_proto_msgTypes,
	}.Build()
	File_superhero_v1_superhero_proto = out.File
	file_superhero_v1_superhero_proto_rawDesc = nil
	file_superhero_v1_superhero_proto_goTypes = nil
	file_superhero_v1_superhero_proto_depIdxs = nil
}
