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
	"strconv"
)

const StorageSequenceNum = 999  //存储内序列目录的最大值

var serverConfig  *goconfig.ConfigFile  //配置文件
var userLoginStatus *UserIsLogin  //登录用户表
var dbConn *sql.DB   //数据库连接
var storageArray []StorageInfo  //存储盘位置
var storageChan = make(chan StorageInfo,5)
var logInfo *log.Logger  //日志
var errLog *log.Logger  //错误日志

func init() {
	serverConfig = GetConfig("server")  //初始化配置文件
	prepareStorage()  //准备存储
	userLoginStatus = NewUserIsLogin()  //初始化用户登录表
	runtime.GOMAXPROCS(runtime.NumCPU())
	dbConn = connDB()  //初始化数据库连接
	prepareLog()  //准备日志文件
}

func main() {
	service, _ := serverConfig.GetString("server","port")
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
	/*
	for i := 0; i < 10; i++ {
		aa := <- storageChan
		path := aa.Path + aa.SmallPath
		fmt.Println(path)
	}
	*/
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
			processLogin(conn)
		case 2 :
			processAddNewResource(conn)
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

// propareLog 准备日志文件
func prepareLog() {
	logFile, _ := serverConfig.GetString("server","log")
	logw, _ := os.OpenFile(logFile, os.O_WRONLY | os.O_APPEND | os.O_CREATE , 0660)
	logInfo = log.New(logw, "frm_server : ", log.Ldate | log.Ltime)
	
	errFile, _ := serverConfig.GetString("server","err")
	errw, _ := os.OpenFile(errFile, os.O_WRONLY | os.O_APPEND | os.O_CREATE , 0660)
	errLog = log.New(errw, "frm_server : ", log.Ldate | log.Ltime)
}

// prepareStorage 准备存储
func prepareStorage() {
	theS, _ := serverConfig.GetOptions("storage")
	for _, oneS := range theS {
		oneSt, _ := serverConfig.GetString("storage",oneS)
		oneSt = DirMustEnd(oneSt)
		oneInfo := StorageInfo{Name: oneS, Path: oneSt, CanUse: true}
		storageArray = append(storageArray,oneInfo)
	}
	for _, oneStorage := range storageArray {
		dirinfo , err := os.Stat(oneStorage.Path)
		if err != nil {
			fmt.Fprintln(os.Stderr, "存储位置无法打开：", oneStorage)
			os.Exit(1)
		}
		if dirinfo.IsDir() == false {
			fmt.Fprintln(os.Stderr, "存储位置需要为一个路径：", oneStorage)
			os.Exit(1)
		}
		
		//开始准备存储内序列目录
		for n := 0; n <= StorageSequenceNum; n++ {
			dirName := strconv.Itoa(n)
			dirName = oneStorage.Path + dirName
			os.Mkdir(dirName, 0600)
		}
		//准备完毕
		go StorageChanSequence()
	}
}
