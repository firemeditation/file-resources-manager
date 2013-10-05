package main

import (
	. "frmPkg"
)


// mergePower 根据Unit、Group、User中的权限合并出最大值
func mergePower(p1, p2, p3 UserPower) UserPower {
	merge := mergePowerAss(p2, p3)
	merge = mergePowerAss(merge, p1)
	return merge
}

// mergePowerAss 为mergePower的辅助函数
func mergePowerAss(p1, p2 UserPower) UserPower {
	tp := UserPower{}
	for k1, _ := range p1 {
		tp[k1] = make(map[string]uint8)
		for k2, v2 := range p1[k1] {
			tp[k1][k2] = v2
		}
	}
	for key1, _ := range p2 {
		if _, f := tp[key1] ; f == false {
			tp[key1] = make(map[string]uint8)
		}
		for key2, value2 := range p2[key1]{
			if v3, found := tp[key1][key2] ; found == true {
				if value2 > v3 {
					tp[key1][key2] = value2
				}
			}else{
				tp[key1][key2] = value2
			}
		}
	}
	return tp
}

// ckLogedUser 检查已经登录的用户是否存在，或者是否登录超时
func ckLogedUser (ckcode string) (ili *IsLoginInfo, ok bool) {
	
	// begin 查看用户是否存在
	ili, found := userLoginStatus.Get(ckcode)
	if found != nil {
		userLoginStatus.Del(ckcode)
		ok = false
		return ili, ok
	}
	// end
	
	// begin 查看用户是否超时
	timeout_time, _ := serverConfig.GetInt64("user","timeout")
	ck_timeout := ili.NotTimeOut(timeout_time)
	if ck_timeout == false {
		userLoginStatus.Del(ckcode)
		ok = false
		return ili, ok
	}
	// end
	
	ok = true
	
	return  ili, ok
}
