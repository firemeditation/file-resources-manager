package main

import (
	"time"
)

type backupRecordStuct struct{
	Record map[string]map[int64]string
	Num map[string]int
 }

func newBackupRecord() (brt *backupRecordStuct) {
	return &backupRecordStuct{make(map[string]map[int64]string), make(map[string]int)}
}

func (brt *backupRecordStuct) AddRecord (userid, content string) {
	thetime := time.Now().Unix()
	if _, found := brt.Record[userid] ; found == true {
		brt.Record[userid][thetime] = content
	}else{
		brt.Record[userid] =make(map[int64]string)
		brt.Record[userid][thetime] = content
	}
}
