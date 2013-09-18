//File Resources Manager

package main

import (
	//"fmt"
	. "frm_pkg"
	"github.com/msbranco/goconfig"
)

var serverConfig  *goconfig.ConfigFile
var userLoginStatus UserIsLogin

func init() {
	serverConfig = GetConfig("server")
	userLoginStatus = NewUserIsLogin()
}

func main() {
	testIsLoginInfo()
}

