package web

import (
	//. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	//"regexp"
	//"os"
)


func doMain(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	cookiename, _ := ServerConfig.GetString("web","cookie")
	_, found := r.Cookie(cookiename)
	if found != nil {
		fmt.Fprint(w,"<script language='javascript'>window.location.href='login'</script>")
		return
	}
	t_file := GlobalRelativePath + staticPath + "main.htm"
	http.ServeFile(w, r, t_file)
}
