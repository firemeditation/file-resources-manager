// 全局资源锁，兼顾多线程功能

package main

import (
	"sync"
)


type GlobalResourceLockStruct struct {
	ProcessId string  // 进程编号，一个hashid
	ResourceType string  //资源类型
	User string  // 所属用户的登录hashid
	UseType string // 使用类型，是下载、删除、上传还是什么之类的。
	Lock uint8  // 加锁类型：1读，2写
	Time uint64  // 锁更新时间
}

type GlobalResourceLock struct {
	lock *sync.RWMutex
	grls map[string]*GlobalResourceLock  // 这里的[string]为资源的hashid
}

// NewGlobalResourceLock 新建全局资源锁
func NewGlobalResourceLock() *GlobalResourceLock {
	return &GlobalResourceLock{new(sync.RWMutex), make(map[string]*GlobalResourceLock)}
}
