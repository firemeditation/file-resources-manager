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
	userhash_cookie, err := r.Cookie(cookiename)
	if err != nil {
		fmt.Fprint(w,"<script language='javascript'>window.location.href='login'</script>")
		return
	}
	
	// begin 查看用户是否存在或超时
	_, found  := CkLogedUser (userhash_cookie.Value)
	if found == false {
		fmt.Fprint(w,"<script language='javascript'>window.location.href='login'</script>")
		return
	}
	// end
	
	t_file := GlobalRelativePath + staticPath + "main.htm"
	http.ServeFile(w, r, t_file)
}
