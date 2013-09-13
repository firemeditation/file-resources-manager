package frm_pkg

import (
	"os"
	"path/filepath"
	"github.com/msbranco/goconfig"
)

func GetConfig(sorc string) *goconfig.ConfigFile {
	cfg_file := filepath.Dir(os.Args[0])
	cfg_file = cfg_file + "/" + sorc + ".cfg"
	c, _ := goconfig.ReadConfigFile(cfg_file)
	return c
}
