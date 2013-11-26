package public

import (
	"fmt"
	"os"
	. "frmPkg"
	"github.com/msbranco/goconfig"
	"runtime"
	"log"
	"database/sql"
	"strconv"
)

const StorageSequenceNum = 999  //存储内序列目录的最大值

var ServerConfig  *goconfig.ConfigFile  //配置文件
var UserLoginStatus *UserIsLogin  //登录用户表
var DbConn *sql.DB   //数据库连接
var StorageArray []StorageInfo  //存储盘位置
var StorageChan = make(chan StorageInfo,5)
var LogInfo *log.Logger  //日志
var ErrLog *log.Logger  //错误日志
var GlobalLock *GlobalResourceLock  //全局资源锁

func StartSystem() {
	ServerConfig = GetConfig("server")  //初始化配置文件
	prepareStorage()  //准备存储
	UserLoginStatus = NewUserIsLogin()  //初始化用户登录表
	runtime.GOMAXPROCS(runtime.NumCPU())
	DbConn = connDB()  //初始化数据库连接
	prepareLog()  //准备日志文件
	GlobalLock = NewGlobalResourceLock()  //启动全局资源锁
}


// propareLog 准备日志文件
func prepareLog() {
	logFile, _ := ServerConfig.GetString("server","log")
	logw, _ := os.OpenFile(logFile, os.O_WRONLY | os.O_APPEND | os.O_CREATE , 0660)
	LogInfo = log.New(logw, "frm_server : ", log.Ldate | log.Ltime)
	
	errFile, _ := ServerConfig.GetString("server","err")
	errw, _ := os.OpenFile(errFile, os.O_WRONLY | os.O_APPEND | os.O_CREATE , 0660)
	ErrLog = log.New(errw, "frm_server : ", log.Ldate | log.Ltime)
}

// prepareStorage 准备存储
func prepareStorage() {
	theS, _ := ServerConfig.GetOptions("storage")
	for _, oneS := range theS {
		oneSt, _ := ServerConfig.GetString("storage",oneS)
		oneSt = DirMustEnd(oneSt)
		oneInfo := StorageInfo{Name: oneS, Path: oneSt, CanUse: true}
		StorageArray = append(StorageArray,oneInfo)
	}
	for _, oneStorage := range StorageArray {
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
			os.Mkdir(dirName, 0700)
		}
		//准备完毕
		go StorageChanSequence()
	}
}



// mergePower 根据Unit、Group、User中的权限合并出最大值
func MergePower(p1, p2, p3 UserPower) UserPower {
	merge := mergePowerAss(p2, p3)
	merge = mergePowerAss(merge, p1)
	return merge
}

// mergePowerAss 为mergePower的辅助函数
func mergePowerAss(p1, p2 UserPower) UserPower {
	tp := UserPower{}
	for k1, _ := range p1 {
		tp[k1] = make(map[string]uint8)
		for k2, v2 := range p1[k1] {
			tp[k1][k2] = v2
		}
	}
	for key1, _ := range p2 {
		if _, f := tp[key1] ; f == false {
			tp[key1] = make(map[string]uint8)
		}
		for key2, value2 := range p2[key1]{
			if v3, found := tp[key1][key2] ; found == true {
				if value2 > v3 {
					tp[key1][key2] = value2
				}
			}else{
				tp[key1][key2] = value2
			}
		}
	}
	return tp
}

// ckLogedUser 检查已经登录的用户是否存在，或者是否登录超时
func CkLogedUser (ckcode string) (ili *IsLoginInfo, ok bool) {
	
	// begin 查看用户是否存在
	ili, found := UserLoginStatus.Get(ckcode)
	if found != nil {
		UserLoginStatus.Del(ckcode)
		ok = false
		return ili, ok
	}
	// end
	
	// begin 查看用户是否超时
	timeout_time, _ := ServerConfig.GetInt64("user","timeout")
	ck_timeout := ili.NotTimeOut(timeout_time)
	if ck_timeout == false {
		UserLoginStatus.Del(ckcode)
		ok = false
		return ili, ok
	}
	// end
	
	ok = true
	
	return  ili, ok
}
