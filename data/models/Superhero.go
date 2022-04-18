package models

import "gorm.io/gorm"

type Atributos struct {
	Debilidad string `protobuf:"bytes,1,opt,name=debilidad,proto3" json:"debilidad,omitempty"`
	Ataque    int32  `protobuf:"varint,2,opt,name=ataque,proto3" json:"ataque,omitempty"`
	Vida      int32  `protobuf:"varint,3,opt,name=vida,proto3" json:"vida,omitempty"`
	Genero    string `protobuf:"varint,4,opt,name=genero,proto3,enum=superhero.v1.Genero" json:"genero,omitempty"`
}

// gorm:"embedded" to put array of Childrens inside Parent
type Superhero struct {
	gorm.Model
	Nombre    string    `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Habilidad string    `protobuf:"bytes,2,opt,name=habilidad,proto3" json:"habilidad,omitempty"`
	Atributos Atributos `gorm:"embedded"`
	Casa      string    `protobuf:"varint,4,opt,name=casa,proto3,enum=superhero.v1.Casa" json:"casa,omitempty"`
}
