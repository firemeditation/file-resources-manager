package main

import (
	"fmt"
	"github.com/msbranco/goconfig"
)

func main() {
	c, _ := goconfig.ReadConfigFile("config/client.cfg")
	fmt.Println(c)
}
