package web

import (
	. "frmPkg"
	//. "frmServer/public"
	"fmt"
	"net/http"
	//"strings"
	//"regexp"
	//"os"
)


func iAddOneResource(ili *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, r.PostFormValue("type"))
	fmt.Fprint(w, r.PostFormValue("json"))
	fmt.Fprint(w, r.Form )
}
