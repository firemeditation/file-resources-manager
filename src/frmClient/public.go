package main

import (
	"net"
	"os"
	"fmt"
	. "frmPkg"
	"os/exec"
)

//sendTheFirstRequest 发送版本号和请求操作的类型
func sendTheFirstRequest (version , retype uint8, conn *net.TCPConn) error {
	ver := Uint8ToBytes(version)
	n, err := conn.Write(ver)  //发送版本号1
	if n != 1 || err != nil {
		return err;
	}
	doType := Uint8ToBytes(retype)
	n, err = conn.Write(doType)
	if n != 1 || err != nil {
		return err;
	}
	return err
}

//connectServer 根据配置文件地址连接服务器
func connectServer() *net.TCPConn {
	serviceAddr, _ := clientConfig.GetString("server","addr")
	IPAdrr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "连接服务器出错：", err)
		os.Exit(1)
	}
	Connecter, err := net.DialTCP("tcp", nil, IPAdrr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "连接服务器出错：", err)
		os.Exit(1)
	}
	return Connecter
}

//clearScreen 清空屏幕内容
func clearScreen() {
	c := exec.Command("clear")  //清空屏幕
    c.Stdout = os.Stdout
    c.Run()
}

