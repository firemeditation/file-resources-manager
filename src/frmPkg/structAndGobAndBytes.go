package frmPkg

import (
	"bytes"
	"encoding/gob"
)

// StructGobBytes 将结构体数据转成Gob再转成[]Byte
func StructGobBytes(e interface{}) []byte {
	var gob_buff bytes.Buffer  //建立缓冲
	gob_en := gob.NewEncoder(&gob_buff)  //gob开始编码
	gob_en.Encode(e)  //gob编码
	gob_b := gob_buff.Bytes()  //bytes.buffer转[]byte
	return gob_b
}

// BytesGobStruct 将[]byte转成Gob再转成结构体
func BytesGobStruct(f_b []byte, stur interface{}) {
	b_buf := bytes.NewBuffer(f_b)  //将[]byte放入bytes的buffer中
	b_go := gob.NewDecoder(b_buf)  //将buffer放入gob的decoder中
	b_go.Decode(stur)  //将gob解码放入stur
}
