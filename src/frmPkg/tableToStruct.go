package frmPkg


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
	Id uint32 `PK`
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
	Btime int64
	Derivative string
	UnitsId uint16
	PowerLevel uint8
	UsersId uint32
	Expand uint16
	MetaData string 
}
//MetaData
type ResourceGroupTable_MD struct{
	Author string
	Editor string
	ISBN string
}
//Related search results
type ResourceGroupTable_RSR struct{
	RtName string
	UnintsName string
	UsersName string
}
type ResourceItemTable struct {
	HashId string
	Name string
	RiType uint8
	LastTime int64
	Version uint16
	RgHashId string
	Derivative string
	UnitsId uint16
	PowerLevel uint32
	UsersId uint32
	Expand uint16
	MetaData string
}
type ResourceFileTable struct {
	ResourceItemTable
	Fname string
	ExtName string
	Opath string
	Fpath string
	Fsite string
	Fsize int64
} 
type ResourceTextTable struct {
	ResourceItemTable
	ContentType string
	Conent string
}
type ResourceRelationTable struct {
	QuoteSide string
	BeQuote string
	RrType uint8
}
type ResourceStatusTable struct {
	HashId string
	Status1 uint8
	Status2 uint8
	Status4 uint8
	Status5 uint8
	Status6 uint8
	Status7 uint8
	Status8 uint8
	Status9 uint8
}
//数据表对应结构体结束
