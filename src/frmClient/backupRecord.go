package main

import (
	"time"
	"fmt"
)

type backupRecordStruct struct{
	Record map[string]map[string]map[string]string
	Num map[string]int
 }

func newBackupRecordStuct () (brt *backupRecordStruct) {
	return &backupRecordStruct{map[string]map[string]map[string]string{}, map[string]int{}}
}

func (brt *backupRecordStruct) AddRecord (userid, content string) {
	thetime := fmt.Sprint(time.Now().Unix())
	thenano := fmt.Sprint(time.Now().UnixNano())
	if _, found := brt.Record[userid] ; found == false {
		brt.Record[userid] =make(map[string]map[string]string)
	}
	if _, found := brt.Record[userid][thetime] ; found == false {
		brt.Record[userid][thetime] = make(map[string]string)
	}
	brt.Record[userid][thetime][thenano] = content
}

func (brt *backupRecordStruct) AddNum (userid string) {
	if _, found := brt.Num[userid]; found == false {
		brt.Num[userid] = 1
	}else{
		brt.Num[userid]++
	}
}

func (brt *backupRecordStruct) DoneNum (userid string) {
	if _, found := brt.Num[userid]; found == true {
		if brt.Num[userid] >= 1 {
			brt.Num[userid]--
		}
	}
}

func (brt *backupRecordStruct) GetNum (userid string) (num int) {
	if _, found := brt.Num[userid]; found == true {
		return brt.Num[userid]
	}
	return 
}
