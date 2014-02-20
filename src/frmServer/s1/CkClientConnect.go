package s1

import (
	"net"
	. "frmPkg"
)

func CkClientConnect(conn *net.TCPConn){
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
}
