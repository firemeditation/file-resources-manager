//File Resources Manager

package main

import (
	"fmt"
	frm "frm_pkg"
)

func main() {
	c := frm.GetConfig("server")
	fmt.Println(c)
	testIsLoginInfo()
}

