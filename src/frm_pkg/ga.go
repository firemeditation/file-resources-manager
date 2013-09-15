package frm_pkg

import (
	"os"
	"path/filepath"
	"github.com/msbranco/goconfig"
	"encoding/binary"
)

//func GetConfig(sorc string) *goconfig.ConfigFile
//获取配置文件信息
func GetConfig(sorc string) *goconfig.ConfigFile {
	cfg_file := filepath.Dir(os.Args[0])
	cfg_file = cfg_file + "/" + sorc + ".cfg"
	c, _ := goconfig.ReadConfigFile(cfg_file)
	return c
}

//Uint64转[]byte
func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

//[]byte转uint64
func BytesToUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}

//Uint32转[]byte
func Uint32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, i)
	return buf
}

//[]byte转uint32
func BytesToUint32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

//Uint16转[]byte
func Uint16ToBytes(i uint16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, i)
	return buf
}

//[]byte转uint16
func BytesToUint16(buf []byte) uint16 {
	return binary.BigEndian.Uint16(buf)
}
