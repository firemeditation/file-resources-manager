package main


// 一组数据库表对应结构体
type UnitsTable struct {
	Id uint16 `PK`
	Name string
	Expand uint16
	PowerLevel string
	Info string
}
type GroupsTable struct {
	Id uint16 `PK`
	Name string
	Expend uint16
	PowerLevel string
	Info string
}
type UsersTable struct {
	Id uint32 `pk`
	Name string
	Passwd string
	UnitsId uint16
	GroupsId uint16
	Expend uint16
	PowerLevel string
}
type ResourceTypeTable struct {
	id uint32 `PK`
	Name string
	PowerLevel uint8
	Expend uint16
	Info string
}
type ResourceGroupTable struct {
	HashId string
	Name string
	RtId uint32
	Info string
	Btime uint64
	Derivative string
	UnitsId uint16
	PowerLevel uint8
	UsersId uint32
	Expand uint16
}
type ResourceItemTable struct {
	HashId string
	Name string
	LastTime uint64
	Version uint16
	RgHashId string
	Derivative string
	UnitsId uint16
	PowerLevel uint32
	UsersId uint32
	Expand uint16
}
type ResourceFileTable struct {
	ResourceItemTable
	Fname string
	ExtName string
	Opath string
	Fpath string
	Fsite uint16
	Fsize uint64
	MetaData string
} 
type ResourceTextTable struct {
	ResourceItemTable
	Conent string
	MetaData string
}
type ResourceRelationTable struct {
	QuoteSide string
	BeQuote string
	RrType uint8
}
//数据表对应结构体结束
