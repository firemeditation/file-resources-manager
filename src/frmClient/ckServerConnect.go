package main

import (
	. "frmPkg"
	"time"
)

func ckServerConnectLoop() {
	for{	
		ckServerConnect()
		
		time.Sleep(900*time.Second)
	}
}

func ckServerConnect() (status uint8){
	conn, err := connectServer()
	if err != nil {
		serverConnectStatus = 2
		status = 2
		return
	}
	err = sendTheFirstRequest (1, 5, conn)
	if err != nil {
		serverConnectStatus = 2
		status = 2
		return
	}
	_ , err = ReadSocketBytes(conn, 1)
	if err != nil {
		serverConnectStatus = 2
		status = 2
		return
	}
	serverConnectStatus = 1
	status = 1
	return
}
