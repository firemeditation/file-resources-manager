package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func iDeleteResourceGroup(theUser *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	
	// start 查看用户是否有查看资源的权力
	if theUser.UPower["resource"]["origin"] < 2 {
		fmt.Fprint(w,"{\"err\":\"无删除权限\"}")
		return
	}
	
	if _, found := r.Form["hashid"]; found == false {
		fmt.Fprint(w,"{\"err\":\"没有提供正确的图书信息\"}")
		return
	}
	book_hashid := r.Form["hashid"][0]
	
	//查看这个资源自己有权利删除吗
	var b_t_unit uint16
	DbConn.QueryRow("select units_id from resourceGroup where hashid = $1", book_hashid).Scan(&b_t_unit)
	if b_t_unit != theUser.UnitId{
		fmt.Fprint(w,"{\"err\":\"无删除权限\"}")
		return
	}
	
	// 请求加锁
	processid, err := GlobalLock.TryLock(theUser.HashId, book_hashid, 1)
	if err != nil {
		fmt.Fprint(w,"{\"err\":\"加锁失败，请稍后尝试\"}")
		return
	}
	defer GlobalLock.Unlock(book_hashid, processid)
	
	// 删除所有的文件
	all_file, _ := DbConn.Query("select fpath, fsite from resourceFile where rg_hashid = $1", book_hashid)
	for all_file.Next(){
		var fpath, fsite string
		all_file.Scan(&fpath, &fsite)
		fsite = strings.TrimSpace(fsite)
		fpath = strings.TrimSpace(fpath)
		file_true_path := fsite + fpath
		os.Remove(file_true_path)
	}
	del_all, _ := DbConn.Prepare("delete from resourceFile where rg_hashid = $1")
	del_all.Exec(book_hashid);
		
	//删除资源条目
	del_group, _ := DbConn.Prepare("delete from resourceGroup where hashid = $1")
	del_group.Exec(book_hashid)
	
	// 搜索缓存的删除
	SearchCache.Insert("del", book_hashid, "rg")
	
	fmt.Fprint(w,"{\"ok\":\"资源删除完毕\"}")
}
