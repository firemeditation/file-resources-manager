package main

import (
	"net"
	. "frmPkg"
)

func processAddNewResource(conn *net.TCPConn) {
	theSIDb, _ := ReadSocketBytes(conn, 40)
	if _, found := userLoginStatus[string(theSIDb)] ; found == false {
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	
	theUser := userLoginStatus[string(theSIDb)]
	
	// start 查看用户是否有建立资源的权力
	if theUser.UPower["resource"]["origin"] < 10 {
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}else{
		SendSocketBytes (conn , Uint8ToBytes(1), 1)
	}
	// end
}
