//File Resources Manager

package main

import (
	"fmt"
	. "frm_pkg"
	"github.com/msbranco/goconfig"
)

var serverConfig  *goconfig.ConfigFile
var myLogin SelfLoginInfo

func init() {
	serverConfig = GetConfig("client")
}

func main() {
	c := GetConfig("client")
	fmt.Println(c)
}

