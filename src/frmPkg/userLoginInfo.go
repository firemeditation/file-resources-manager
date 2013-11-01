// userIsLogin 内含 UserPower，IsLoginInfo


package frmPkg

import (
	"time"
	"fmt"
	"sync"
)


//UserPower 用户具体权限为[权限大类][权限小类]权限级别
type UserPower map[string]map[string]uint8


//AddPower 增加一个Power值
func (up UserPower) Add (key1, key2 string, value uint8) {
	if _, f := up[key1] ; f == false {
		up[key1] = make(map[string]uint8)
	}
	up[key1][key2] = value
}

// CheckPowerLevel 检查UPower的权限
func (up UserPower) CheckPowerLevel (topp string, secp string, asklevel uint8) bool {
	_, found := up[topp][secp]
	if found == true {
		if up[topp][secp] >= asklevel {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// UpdatePowerLevel 更新UserPower的值
func (up UserPower) UpdatePowerLevel (topp string, secp string, asklevel uint8) {
	up[topp][secp] = asklevel
}


// IsLoginInfo 是一个记录单个已经登录的人员信息的表
type IsLoginInfoBasic struct {
	Id uint32  //用户ID
	Name string //用户名
	GroupId uint16 //所在组ID
	UnitId uint16 //所在机构ID
	UnitName string //所在机构名称
	UPower UserPower
}

type IsLoginInfo struct {
	IsLoginInfoBasic
	LastTime time.Time
}

// NewIsLoginInfo 是初始化一个人员信息，必须给定name, level, lastTime
func NewIsLoginInfo(id uint32, name string, groupid uint16, unitid uint16, unitname string, lastTime time.Time) *IsLoginInfo{
	return &IsLoginInfo{IsLoginInfoBasic {id, name, groupid, unitid, unitname, UserPower{} }, lastTime}
}

// NotTimeOut 根据给定的int类型的秒数，判断登录是否已经超时，没超时返回true，超时返回false
func (ili *IsLoginInfo) NotTimeOut (timeout int64) bool {
	oldtime := ili.LastTime.Unix()
	nowtime := time.Now().Unix()
	if oldtime + timeout < nowtime {
		return false
	}else{
		return true
	}
}

// UpdateLastTime 已当前时间写入LastTime中进行更新，并返回写入的时间
func (ili *IsLoginInfo) UpdateLastTime () time.Time {
	ili.LastTime = time.Now()
	return ili.LastTime
}




// UserIsLogin 是一个map，记录所有已经登录的人员信息
type UserIsLogin struct{
	lock *sync.RWMutex
	islogin map[string]*IsLoginInfo
}

// NewUserIsLogin 初始化UserIsLogin的map
func NewUserIsLogin () *UserIsLogin {
	return &UserIsLogin{lock: new(sync.RWMutex),islogin: make(map[string]*IsLoginInfo)}
}

// Add 增加一条用户信息，返回响应的IsLoginInfo，如果ckcode重复，则返回错误
func (uil *UserIsLogin) Add (ckcode string, id uint32, name string, groupid uint16, unitid uint16, unitname string, lastTime time.Time) (ili *IsLoginInfo, err error) {
	uil.lock.Lock()
	defer uil.lock.Unlock()
	_, found := uil.islogin[ckcode]
	if  found == false {
		uil.islogin[ckcode] = NewIsLoginInfo(id, name, groupid, unitid, unitname, lastTime)
		return uil.islogin[ckcode], err
	}else{
		err = fmt.Errorf("键值 %x 已经存在，不能新建", ckcode)
		return ili, err
	}
}

// Get 获得ckode的用户登录信息，如果err不为nil则为找不到
func (uil *UserIsLogin) Get (ckcode string) (ili *IsLoginInfo, err error) {
	uil.lock.RLock()
	defer uil.lock.RUnlock()
	if ili , found := uil.islogin[ckcode] ; found == true {
		return ili, nil
	}else{
		err = fmt.Errorf("键值 %x 不存在", ckcode)
		return ili, err
	}
}

// Del 删除一条用户信息，通常是在其过期之后
func (uil *UserIsLogin) Del (ckcode string) {
	uil.lock.Lock()
	defer uil.lock.Unlock()
	delete(uil.islogin, ckcode)
}


type SelfLoginInfo struct {
	IsLoginInfoBasic
	SID string
}

func NewSelfLoginInfo (id uint32, name string, groupid uint16 , unitid uint16, unitname string, sid string) *SelfLoginInfo{
	return &SelfLoginInfo{IsLoginInfoBasic {id, name, groupid, unitid, unitname, UserPower{}}, sid}
}
