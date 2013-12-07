package web

import (
	. "frmPkg"
	//. "frmServer/public"
	"fmt"
	"net/http"
	//"regexp"
	//"os"
)


func iGetBasicUserInfo(ili *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	theJSON := StructToJson(ili)
	fmt.Fprint(w, theJSON )
}
