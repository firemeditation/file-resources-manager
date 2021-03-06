//File Resources Manager

package main

import (
	"fmt"
	"net"
	"os"
	. "frmPkg"
	"log"
	"frmServer/s1"
	. "frmServer/public"
	"frmServer/web"
	"sync"
)

func main(){
	StartSystem()
	
	var wg sync.WaitGroup 
	
	wg.Add(1)
	go startServer()
	go startWeb()
	wg.Wait()
}

func startWeb(){
	web.StartWeb()
}

func startServer() {
	service, _ := ServerConfig.GetString("server","port")
	service = ":" + service
	IPAdrr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		fmt.Fprintln(os.Stderr, "出错了，错误是：", err)
		os.Exit(1)
	}
	listens, err := net.ListenTCP("tcp", IPAdrr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "出错了，错误是：", err)
		os.Exit(1)
	}
	for {
		Connecter, err := listens.AcceptTCP()

		if err != nil {
			log.Println("错误：", err)
			continue
		}
		go doAccept(Connecter)
	}
}

// doAccept 进行客户端连接
func doAccept (conn *net.TCPConn) {
	defer conn.Close()
	_, vtype := getFirstRequest(conn)
	switch vtype {
		case 1 :
			s1.ProcessLogin(conn)
		case 2 :
			s1.ProcessAddNewResource(conn)
		case 3 :
			s1.GeneralLock(conn)
		case 4 :
			s1.ProcessUploadProcess(conn)
		case 5 :
			s1.CkClientConnect(conn)
		case 6 :
			s1.ProcessGetFileList(conn)
		case 7 :
			s1.ProcessDownloadFile(conn)
	}
}

// getFirstRequest 获取客户端最初的操作请求：版本号，操作代码
func getFirstRequest(conn *net.TCPConn) (ver, vtype uint8) {
	ver_b , _ := ReadSocketBytes(conn, 1)
	ver = BytesToUint8(ver_b)
	vtype_b , _ := ReadSocketBytes(conn, 1)
	vtype = BytesToUint8(vtype_b)
	return ver , vtype
}

