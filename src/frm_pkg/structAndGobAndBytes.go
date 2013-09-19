package frm_pkg

import (
	"bytes"
	"encoding/gob"
)

func StructGobBytes(e interface{}) []byte {
	var gob_buff bytes.Buffer  //建立缓冲
	gob_en := gob.NewEncoder(&gob_buff)  //gob开始编码
	gob_en.Encode(e)  //gob编码
	gob_b := gob_buff.Bytes()  //bytes.buffer转[]byte
	return gob_b
}

func BytesGobStruct(f_b []byte, stur interface{}) {
	b_buf := bytes.NewBuffer(f_b)  //将SelfLoginInfo的[]byte放入bytes的buffer中
	b_go := gob.NewDecoder(b_buf)  //将SelfLoginInfo的buffer放入gob的decoder中
	b_go.Decode(stur)  //将gob解码放入myLogin
}
