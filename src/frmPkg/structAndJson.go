package frmPkg

import (
	"bytes"
	"encoding/json"
)

// StructToJson 将结构体转成Json的字符串
func StructToJson(e interface{}) string {
	var j_buff bytes.Buffer  //建立缓冲
	j_en := json.NewEncoder(&j_buff)  //json开始编码
	j_en.Encode(e)  //json编码
	j_b := j_buff.Bytes()  //bytes.buffer转[]byte
	return string(j_b)
}

// JsonToStruct 将json的字符串转成结构体
func JsonToStruct(f_b string, stur interface{}) {
	j_buf := bytes.NewBuffer([]byte(f_b))  //将[]byte放入bytes的buffer中
	j_go := json.NewDecoder(j_buf)  //将buffer放入json的decoder中
	j_go.Decode(stur)  //将json解码放入stur
}
