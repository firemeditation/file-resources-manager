// userIsLogin 内含 UserPower，IsLoginInfo


package frmPkg

import (
	"time"
	"fmt"
)


//UserPower 用户具体权限为[权限大类][权限小类]权限级别
type UserPower map[string]map[string]uint8

// IsLoginInfo 是一个记录单个已经登录的人员信息的表
type IsLoginInfoBasic struct {
	Id uint32  //用户ID
	Name string //用户名
	Group uint16 //所在组
	UPower UserPower
}

type IsLoginInfo struct {
	IsLoginInfoBasic
	LastTime time.Time
}

// NewIsLoginInfo 是初始化一个人员信息，必须给定name, level, lastTime
func NewIsLoginInfo(id uint32, name string, group uint16, lastTime time.Time) *IsLoginInfo{
	return &IsLoginInfo{IsLoginInfoBasic {id, name, group, UserPower{"user":{"user":0}} }, lastTime}
}

// CheckPowerLevel 检查UPower的权限
func (ili IsLoginInfoBasic) CheckPowerLevel (topp string, secp string, asklevel uint8) bool {
	_, found := ili.UPower[topp][secp]
	if found == true {
		if ili.UPower[topp][secp] >= asklevel {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// UpdatePowerLevel 更新UserPower的值
func (ili IsLoginInfoBasic) UpdatePowerLevel (topp string, secp string, asklevel uint8) {
	ili.UPower[topp][secp] = asklevel
}

// NotTimeOut 根据给定的int类型的秒数，判断登录是否已经超时，没超时返回true，超时返回false
func (ili *IsLoginInfo) NotTimeOut (timeout int) bool {
	oldtime := ili.LastTime.Unix()
	nowtime := time.Now().Unix()
	if oldtime + int64(timeout) < nowtime {
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
type UserIsLogin map[string]*IsLoginInfo

// NewUserIsLogin 初始化UserIsLogin的map
func NewUserIsLogin () UserIsLogin {
	return UserIsLogin{}
}

// Add 增加一条用户信息，返回响应的IsLoginInfo，如果ckcode重复，则返回错误
func (uil UserIsLogin) Add (ckcode string, id uint32, name string, group uint16, lastTime time.Time) (ili *IsLoginInfo, err error) {
	_, found := uil[ckcode]
	if  found == false {
		uil[ckcode] = NewIsLoginInfo(id, name, group, lastTime)
		return uil[ckcode], err
	}else{
		err = fmt.Errorf("键值 %x 已经存在，不能新建", ckcode)
		return ili, err
	}
}

// Get 获得ckode的用户登录信息，如果err不为nil则为找不到
func (uil UserIsLogin) Get (ckcode string) (ili *IsLoginInfo, err error) {
	if ili , found := uil[ckcode] ; found == true {
		return ili, nil
	}else{
		err = fmt.Errorf("键值 %x 不存在", ckcode)
		return ili, err
	}
}

// Del 删除一条用户信息，通常是在其过期之后
func (uil UserIsLogin) Del (ckcode string) {
	delete(uil, ckcode)
}


type SelfLoginInfo struct {
	IsLoginInfoBasic
	SID string
}

func NewSelfLoginInfo (id uint32, name string, group uint16 ,sid string) *SelfLoginInfo{
	return &SelfLoginInfo{IsLoginInfoBasic {id, name, group, UserPower{"user":{"user":0}} }, sid}
}