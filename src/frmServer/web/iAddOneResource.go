package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	"time"
	"strconv"
	//"strings"
	//"regexp"
	//"os"
)


func iAddOneResource(theUser *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	
	// start 查看用户是否有建立资源的权力
	if theUser.UPower["resource"]["origin"] < 2 {
		fmt.Fprint(w,"{\"err\":\"无添加权限\"}")
		return
	}
	
	var rgt ResourceGroupTable
	
	n_rgt, err := DbConn.Prepare("insert into resourceGroup (hashid, name, rt_id, info, btime, units_id, users_id, metadata) values ($1, $2, $3, $4, $5, $6, $7, $8)")
	
	rgt.Name = r.PostFormValue("bookname")
	booktype, err := strconv.Atoi(r.PostFormValue("booktype"))
	if err != nil {
		fmt.Fprint(w,"{\"err\":\"数据错误，无法添加\"}")
		return
	}
	rgt.RtId = uint32(booktype)
	rgt.Info = r.PostFormValue("bookinfo")
	rgt.MetaData = r.PostFormValue("json")
	
	sha1string := fmt.Sprint(time.Now(), rgt.Name, theUser.Name, rgt.RtId)
	rgt.HashId = GetSha1(sha1string)
	rgt.Btime = time.Now().Unix()
	rgt.UnitsId = theUser.UnitId
	rgt.UsersId = theUser.Id
	_, err = n_rgt.Exec(rgt.HashId, rgt.Name, rgt.RtId, rgt.Info, rgt.Btime, rgt.UnitsId, rgt.UsersId, rgt.MetaData)
	if err != nil {
		fmt.Fprint(w,"{\"err\":\"数据库错误，无法添加\"}")
		return
	}
	
	n_rgstatus, err := DbConn.Prepare("insert into resourceGroupStatus (hashid) values ($1)")
	
	_, err = n_rgstatus.Exec(rgt.HashId)
	
	thereturn := "{\"hashid\":\"" + rgt.HashId + "\"}"
	fmt.Fprint(w, thereturn)
}
