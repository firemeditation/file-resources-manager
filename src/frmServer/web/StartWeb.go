package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	"regexp"
	"os"
)


type MyMux struct {}

var staticPath string


func (p *MyMux) ServeHTTP (w http.ResponseWriter, r *http.Request){
	theStaticMatch := "^/" + staticPath + "(.*)"
	
	if match , _ := regexp.MatchString(theStaticMatch, r.URL.Path); match == true {
		file := GlobalRelativePath + r.URL.Path
		http.ServeFile(w, r, file)
		return
	}
	if r.URL.Path == "/login" {
		doLogin(w, r)
		return
	}
	if r.URL.Path == "/" {
		doMain(w,r)
		return
	}
	http.NotFound(w,r)
	return
}

func StartWeb(){
	staticPath, _ = ServerConfig.GetString("web","static")
	staticPath = DirMustEnd(staticPath)
	
	mux := &MyMux{}
	thePort, _ := ServerConfig.GetString("web","port")
	thePort = ":" + thePort
	err := http.ListenAndServe(thePort, mux)
	if err != nil {
		fmt.Fprintln(os.Stderr, "出错了，错误是：", err)
		os.Exit(1)
	}
	
}
