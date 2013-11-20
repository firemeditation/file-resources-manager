package public

import (
	"math/rand"
	"time"
	"strconv"
)


type StorageInfo struct {
	Name string  //存储位置的名称
	Path string  //位置（绝对）
	CanUse bool  //是否可用
	Min int  //剩余的最小空间
	SmallPath string  //存储内的序列文件夹位置
}


func StorageChanSequence(){
	for {
		for _, ones := range StorageArray {
			if ones.CanUse == false {
				continue
			}else{
				spr := rand.New(rand.NewSource(time.Now().UnixNano()))
				ones.SmallPath = strconv.Itoa(spr.Intn(StorageSequenceNum)) + "/"
				StorageChan <- ones
			}
		}
	}
}
