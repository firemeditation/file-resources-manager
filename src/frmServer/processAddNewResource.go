package main

import (
	"net"
	. "frmPkg"
)

func processAddNewResource(conn *net.TCPConn) {
	theSIDb, _ := ReadSocketBytes(conn, 40)
	
	// begin 查看用户是否存在或超时
	theUser, found  := ckLogedUser (string(theSIDb))
	if found == false {
		SendSocketBytes (conn , Uint8ToBytes(3), 1)
		return
	}
	// end
	
	// start 查看用户是否有建立资源的权力
	if theUser.UPower["resource"]["origin"] < 10 {
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}else{
		SendSocketBytes (conn , Uint8ToBytes(1), 1)
	}
	// end
}
