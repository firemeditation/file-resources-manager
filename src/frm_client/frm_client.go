//File Resources Manager

package main

import (
	"fmt"
	frm "frm_pkg"
)

func main() {
	c := frm.GetConfig("client")
	fmt.Println(c)
}

