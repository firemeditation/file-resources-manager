package main

import (
	"time"
)

type backupRecordType map[string]map[int64]string

func newBackupRecord() (brt backupRecordType) {
	return brt
}

func (brt backupRecordType) Add (userid, content string) {
	thetime := time.Now().Unix()
	if _, found := brt[userid] ; found == true {
		brt[userid][thetime] = content
	}else{
		brt[userid] = map[int64]string{}
		brt[userid][thetime] = content
	}
}
