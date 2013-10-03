//File Resources Manager

package main

import (
	"fmt"
	"net"
	"os"
	. "frmPkg"
	"github.com/msbranco/goconfig"
	"runtime"
	"log"
	"database/sql"
)

var serverConfig  *goconfig.ConfigFile
var userLoginStatus UserIsLogin
var dbConn *sql.DB
var logInfo *log.Logger

func init() {
	serverConfig = GetConfig("server")
	userLoginStatus = NewUserIsLogin()
	runtime.GOMAXPROCS(runtime.NumCPU())
	dbConn = connDB()
	prepareLog()
}

func main() {
	service, _ := serverConfig.GetString("server","port")
	service = ":" + service
	IPAdrr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "出错了，错误是：", err)
		os.Exit(1)
	}
	listens, err := net.ListenTCP("tcp", IPAdrr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "出错了，错误是：", err)
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

func doAccept (conn *net.TCPConn) {
	defer conn.Close()
	_, vtype := getFirstRequest(conn)
	switch vtype {
		case 1 :
			processLogin(conn)
		case 2 :
			processAddNewResource(conn)
	}
}

func getFirstRequest(conn *net.TCPConn) (ver, vtype uint8) {
	ver_b , _ := ReadSocketBytes(conn, 1)
	ver = BytesToUint8(ver_b)
	vtype_b , _ := ReadSocketBytes(conn, 1)
	vtype = BytesToUint8(vtype_b)
	return ver , vtype
}

func prepareLog() {
	logFile, _ := serverConfig.GetString("server","log")
	logw, _ := os.OpenFile(logFile, os.O_WRONLY | os.O_APPEND | os.O_CREATE , 0664)
	logInfo = log.New(logw, "frm_server : ", log.Ldate | log.Ltime)
}
