// 全局资源锁，兼顾多线程功能

package main

import (
	"sync"
	. "frmPkg"
	"fmt"
	"time"
)


type GlobalResourceLockUser struct {
	UserId string //所属用户的登录hashid
	Time int64  //锁更新时间
}

type GlobalResourceLockStruct struct {
	ReadProcess map[string]*GlobalResourceLockUser  //string为进程hashid
	WriteProcess string //写锁的进程hashid
	WriteUser *GlobalResourceLockUser
	LockType uint8  // 加锁类型：1读，2写
}

type GlobalResourceLock struct {
	lock *sync.RWMutex
	grls map[string]*GlobalResourceLockStruct  // 这里的[string]为资源的hashid
	timeout int64
}

func NewGlobalResourceLock() *GlobalResourceLock {
	timeout , _ := serverConfig.GetInt64("lock","timeout")
	return &GlobalResourceLock{new(sync.RWMutex),make(map[string]*GlobalResourceLockStruct),timeout}
}

// Lock 添加一个锁
// 1. 查看资源是否已经有锁
// 2. 如果有锁则看是读锁还是写锁，如果是写锁，则看是否已经超时，如果不超时则退回，如果超时则修改添加
// 3. 如果是读锁，而自己也是读锁，则把自己加到读锁序列
// 4. 如果是读锁，而自己是写锁，则遍历读锁看是否全部超时，如果全部超时就删除读锁新建写锁
// 5. 如果资源没有锁，则添加锁
// 6. 最终返回进程hashid
func (grl *GlobalResourceLock) Lock (userid string, resourceid string, locktype uint8) (processid string, err error){
	grl.lock.Lock()
	defer grl.lock.Unlock()
	processid = grl.getProcessid(userid, resourceid)
	// 看是否有锁
	one_grls , found := grl.grls[resourceid]
	// 如果没有锁
	if found == false {
		if locktype == 1 {
			grl.grls[resourceid] = &GlobalResourceLockStruct{WriteProcess: processid, WriteUser: &GlobalResourceLockUser{UserId: userid, Time: time.Now().Unix()}, LockType: 1}
		}else{
			grlsr := GlobalResourceLockStruct{ReadProcess: make(map[string]*GlobalResourceLockUser),LockType: 2}
			grlsr.ReadProcess[processid] = &GlobalResourceLockUser{UserId: userid, Time: time.Now().Unix()}
			grl.grls[resourceid] = &grlsr
		}
	} else {
		// 如果有锁
		// 如果有写锁，自己是写锁
		if locktype == 1 && one_grls.LockType == 1 {
			if one_grls.WriteUser.Time + grl.timeout >= time.Now().Unix(){
				err = fmt.Errorf("无法加锁：%s", resourceid)
			}else{
				grl.grls[resourceid] = &GlobalResourceLockStruct{WriteProcess: processid, WriteUser: &GlobalResourceLockUser{UserId: userid, Time: time.Now().Unix()}, LockType: 1}
			}
		}else if locktype == 2 && one_grls.LockType == 1 {
			// 如果是写锁，自己是读锁
			if one_grls.WriteUser.Time + grl.timeout >= time.Now().Unix(){
				err = fmt.Errorf("无法加锁：%s", resourceid)
			}else{
				grlsr := GlobalResourceLockStruct{ReadProcess: make(map[string]*GlobalResourceLockUser),LockType: 2}
				grlsr.ReadProcess[processid] = &GlobalResourceLockUser{UserId: userid, Time: time.Now().Unix()}
				grl.grls[resourceid] = &grlsr
			}
		}else if locktype == 2 && one_grls.LockType == 2{
			// 如果是读锁，自己是读锁
			grlsr := GlobalResourceLockStruct{ReadProcess: make(map[string]*GlobalResourceLockUser),LockType: 2}
			grlsr.ReadProcess[processid] = &GlobalResourceLockUser{UserId: userid, Time: time.Now().Unix()}
			grl.grls[resourceid] = &grlsr
		}else if locktype == 1 && one_grls.LockType == 2 {
			// 如果是读锁，自己是写锁
			allout := 1
			for _, one_grlu := range grl.grls[resourceid].ReadProcess {
				if one_grlu.Time + grl.timeout >= time.Now().Unix() {
					allout = 2
					break
				}
			}
			if allout == 2 {
				grl.grls[resourceid] = &GlobalResourceLockStruct{WriteProcess: processid, WriteUser: &GlobalResourceLockUser{UserId: userid, Time: time.Now().Unix()}, LockType: 1}
			}else{
				err = fmt.Errorf("无法加锁：%s", resourceid)
			}
		}
	}
	return
}

// TryLock 尝试加锁10次，每次停顿1秒钟，如果10次都不成功则返回错误
func (grl *GlobalResourceLock) TryLock (userid string, resourceid string, locktype uint8) (processid string, err error){
	err = fmt.Errorf("无法加锁：%s", resourceid)
	for i := 0; i < 10; i++ {
		processid, err = grl.Lock(userid, resourceid, locktype)
		if err == nil {
			break
		}else{
			time.Sleep(time.Second)
		}
	}
	return
}

// Unlock 解锁
func (grl *GlobalResourceLock) Unlock (resourceid , processid string) (err error) {
	grl.lock.Lock()
	defer grl.lock.Unlock()
	one_grls , found := grl.grls[resourceid]
	if found == true {
		if one_grls.LockType == 1 {
			delete(grl.grls, resourceid)
		}else{
			_, found := grl.grls[resourceid].ReadProcess[processid]
			if found == true {
				delete(grl.grls[resourceid].ReadProcess, processid)
				if len(grl.grls[resourceid].ReadProcess) == 0 {
					delete(grl.grls, resourceid)
				}
			}else{
				err = fmt.Errorf("键找不到：%s", processid)
			}
		}
	}else{
		err = fmt.Errorf("键找不到：%s", resourceid)
	}
	return 
}

// Uptime 更新时间
func (grl *GlobalResourceLock) Uptime (resourceid , processid string) (err error) {
	one_grls , found := grl.grls[resourceid]
	if found == true {
		if one_grls.LockType == 1 {
			if one_grls.WriteUser.Time + grl.timeout >= time.Now().Unix() {
				grl.grls[resourceid].WriteUser.Time = time.Now().Unix()
			}
		}else{
			one_grlu, found := grl.grls[resourceid].ReadProcess[processid]
			if found == true {
				if one_grlu.Time + grl.timeout >= time.Now().Unix() {
					grl.grls[resourceid].ReadProcess[processid].Time = time.Now().Unix()
				}
			}else{
				err = fmt.Errorf("键找不到：%s", processid)
			}
		}
	}else{
		err = fmt.Errorf("键找不到：%s", resourceid)
	}
	return 
}

// GlobalResourceLockClearup 定时清理函数，用go异步执行，每30秒清理一次
func GlobalResourceLockClearup() {
	for {
		time.Sleep(30 * time.Second)
	}
}

// getProcessid 获取进程id
func (grl *GlobalResourceLock) getProcessid (a, b string) string {
	thes := a + b + time.Now().String()
	return GetSha1(thes)
}
