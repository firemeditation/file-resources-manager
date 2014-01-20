package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func iDeleteResourceFile(theUser *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	
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
	
	if _, found := r.Form["dtype"]; found == false{
		fmt.Fprint(w,"{\"err\":\"没有提供正确的删除请求\"}")
		return
	}
	
	del_type := r.Form["dtype"][0]
	
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
	
	if del_type == "all" {
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
		
		//更新文件数量
		up_count, _ := DbConn.Prepare("update resourceGroupStatus set status1 = 0 where hashid = $1")
		up_count.Exec(book_hashid)
		fmt.Fprint(w,"{\"ok\":\"文件清空完毕\"}")
	}else if del_type == "one"{
		if _, found := r.Form["file"]; found == false{
			fmt.Fprint(w,"{\"err\":\"没有提供正确的删除请求\"}")
			return
		}
		var fpath, fsite string
		DbConn.QueryRow("select fpath, fsite from resourceFile where hashid = $1", r.Form["file"][0]).Scan(&fpath, &fsite);
		fsite = strings.TrimSpace(fsite)
		fpath = strings.TrimSpace(fpath)
		file_true_path := fsite + fpath
		os.Remove(file_true_path)
		del_all, _ := DbConn.Prepare("delete from resourceFile where hashid = $1")
		del_all.Exec(r.Form["file"][0]);
		
		//更新文件数量
		up_count, _ := DbConn.Prepare("update resourceGroupStatus set status1 = status1 - 1 where hashid = $1")
		up_count.Exec(book_hashid)
		fmt.Fprint(w,"{\"ok\":\"文件删除完毕\"}")
	}else if del_type == "part" {
		if _, found := r.Form["file"]; found == false{
			fmt.Fprint(w,"{\"err\":\"没有提供正确的删除请求\"}")
			return
		}
		file_array := strings.Split(r.Form["file"][0],",")
		file_array = file_array[:len(file_array)-1]
		dbBegin, _ := DbConn.Begin()
		del_file, _ := dbBegin.Prepare("delete from resourceFile where hashid = $1")
		for _, onefile := range file_array {
			var fpath, fsite string
			DbConn.QueryRow("select fpath, fsite from resourceFile where hashid = $1", onefile).Scan(&fpath, &fsite);
			fsite = strings.TrimSpace(fsite)
			fpath = strings.TrimSpace(fpath)
			file_true_path := fsite + fpath
			os.Remove(file_true_path)
			del_file.Exec(onefile)
		}
		dbBegin.Commit()
		
		var now_count uint64
		DbConn.QueryRow("select COUNT(*) from resourceFile where rg_hashid = $1", book_hashid).Scan(&now_count)
		//更新文件数量
		up_count, _ := DbConn.Prepare("update resourceGroupStatus set status1 = $1 where hashid = $2")
		up_count.Exec(now_count, book_hashid)
		fmt.Fprint(w,"{\"ok\":\"文件删除完毕\"}")
	}
}
