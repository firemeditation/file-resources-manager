package frmPkg

import (
	"os"
	"path/filepath"
	"github.com/msbranco/goconfig"
	"encoding/binary"
	"crypto/sha1"
	"fmt"
	"io"
	"net"
	"bufio"
)

const bytelen = 1024

// GetConfig 为获取配置文件信息
func GetConfig(sorc string) *goconfig.ConfigFile {
	cfg_file := filepath.Dir(os.Args[0])
	cfg_file = cfg_file + "/" + sorc + ".cfg"
	c, _ := goconfig.ReadConfigFile(cfg_file)
	return c
}

// Uint64ToBytes 为Uint64转[]byte
func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

// BytesToUint64 []byte转uint64
func BytesToUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}

// Uint32ToBytes Uint32转[]byte
func Uint32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, i)
	return buf
}

// BytesToUint32 []byte转uint32
func BytesToUint32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

// Uint16ToBytes Uint16转[]byte
func Uint16ToBytes(i uint16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, i)
	return buf
}

// BytesToUint16 []byte转uint16
func BytesToUint16(buf []byte) uint16 {
	return binary.BigEndian.Uint16(buf)
}

// Uint8ToBytes Uint8转[]byte
func Uint8ToBytes(i uint8) []byte {
	var buf = []byte{i}
	return buf
}

// BytesToUint8 []byte转uint8
func BytesToUint8(buf []byte) uint8 {
	return uint8(buf[0])
}

// GetSha1 生成SHA1值
func GetSha1(data string) string {
    t := sha1.New();
    io.WriteString(t,data);
    return fmt.Sprintf("%x",t.Sum(nil));
}

// ReadSocketToFile 从soket读出文件，用带缓冲的方式写入文件里
func ReadSocketToFile(conn net.Conn, len uint64, file *os.File) (err error) {
	write := bufio.NewWriter(file)
	for {
		tempdata := []byte{}
		if len < uint64(bytelen) {
			tempdata = make([]byte, len)
		} else {
			tempdata = make([]byte, bytelen)
		}
		r, err := conn.Read(tempdata)
		if err != nil {
			return err
		}
		if r != 0 {
			write.Write(tempdata[0:r])
			len = len - uint64(r)
		}

		if len == 0 {
			break
		}
	}
	write.Flush()
	return err
}

// ReadSocketBytes 从socket读出一定长度的数据，放入[]byte中，保证完整读出
func ReadSocketBytes(conn net.Conn, len uint64) (returnByte []byte, err error) {
	returnByte = make([]byte, 0, len)
	for {
		tempdata := []byte{}
		if len < uint64(bytelen) {
			tempdata = make([]byte, len)
		} else {
			tempdata = make([]byte, bytelen)
		}
		r, err := conn.Read(tempdata)
		if err != nil {
			return returnByte, err
		}
		returnByte = append(returnByte, tempdata[:r]...)

		len = len - uint64(r)

		if len == 0 {
			break
		}
	}
	return returnByte, err
}

func SendSocketBytes (conn *net.TCPConn, bytes []byte, len uint64) error {
	n, err := conn.Write(bytes)
	if uint64(n) != len {
		err = fmt.Errorf("不能完整发送信息")
	}
	return err
}
