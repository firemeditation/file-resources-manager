package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/msbranco/goconfig"
)

func main() {
	c := getConfig("server")
	fmt.Println(c)
}

func getConfig(sorc string) *goconfig.ConfigFile {
	cfg_check1 := filepath.Dir(os.Args[0]) 
    cfg_check2 := filepath.Dir(cfg_check1)
    var cfg_file string
    if cfg_check1 == "." && cfg_check2 == "." {
		cfg_file = "../config/"
	}else{
		cfg_file = cfg_check2 + "/config/"
	}
	if sorc == "server"{
		cfg_file = cfg_file + "server.cfg"
	}else{
		cfg_file = cfg_file + "client.cfg"
	}
	fmt.Println(cfg_file)
	c, _ := goconfig.ReadConfigFile(cfg_file)
	return c
}
