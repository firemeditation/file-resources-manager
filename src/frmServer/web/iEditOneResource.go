package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	"strconv"
)


func iEditOneResource(theUser *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	
	// start 查看用户是否有建立资源的权力
	if theUser.UPower["resource"]["origin"] < 2 {
		fmt.Fprint(w,"{\"err\":\"无修改权限\"}")
		return
	}
	
	book_hashid := r.PostFormValue("hashid")
	var book_units_id uint16
	DbConn.QueryRow("select units_id from resourceGroup where hashid = $1", book_hashid).Scan(&book_units_id)
	if book_units_id != theUser.UnitId {
		fmt.Fprint(w,"{\"err\":\"无修改权限\"}")
		return
	}
	
	var rgt ResourceGroupTable
	rgt.Name = r.PostFormValue("bookname")
	booktype, err := strconv.Atoi(r.PostFormValue("booktype"))
	if err != nil {
		fmt.Fprint(w,"{\"err\":\"数据错误，无法修改\"}")
		return
	}
	rgt.RtId = uint32(booktype)
	rgt.Info = r.PostFormValue("bookinfo")
	rgt.MetaData = r.PostFormValue("json")
	rgt.UsersId = theUser.Id
	
	edit_rgt, err := DbConn.Prepare("update resourceGroup set name = $1, rt_id = $2, info = $3, users_id = $4, metadata = $5 where hashid = $6")
	if err != nil {
		fmt.Fprint(w,"{\"err\":\"数据库错误，无法修改\"}")
		return
	}
	
	_, err = edit_rgt.Exec(rgt.Name, rgt.RtId, rgt.Info, rgt.UsersId, rgt.MetaData, book_hashid)
	if err != nil {
		fmt.Fprint(w,"{\"err\":\"数据库错误，无法修改\"}")
		return
	}
	SearchCache.Insert("up", book_hashid, "rg")
	
	thereturn := "{\"hashid\":\"" + book_hashid + "\"}"
	fmt.Fprint(w, thereturn)
}
