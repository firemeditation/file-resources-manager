package main

import (
	"fmt"
	"time"
	. "frm_pkg"
)

func testIsLoginInfo(){
	loginIsIn := NewUserIsLogin()
	oneUser,_ := loginIsIn.Add("adfadf", "User1sName", 100, time.Now(), 10)
	fmt.Println(oneUser.Name)
	oneUser2 ,err := loginIsIn.Add("adfadwf", "User2sName", 100, time.Now() ,20)
	fmt.Println(err)
	oneUser.Name ="User1sName-改"
	fmt.Println(loginIsIn["adfadf"].Name)
	fmt.Println(loginIsIn["adfadwf"].Name)
	fmt.Println(oneUser2.LastTime)
	loginIsIn["adfadwf"].UpdateLastTime()
	fmt.Println(oneUser2.LastTime)
	fmt.Println(loginIsIn["adfadwf"].LastTime)
	oneUser2.UpdateLastTime()
	fmt.Println(loginIsIn["adfadwf"].LastTime)
	
	loginIsIn.Del("adfadf")
	//fmt.Println(loginIsIn["adfadf"].Name)
	
	//loginIsIn := make(map[string]*IsLoginInfo)
	//loginIsIn["user1"] = NewIsLoginInfo("User1sName",100,time.Now())
	/*
	loginIsIn["user1"] = &IsLoginInfo{"User1sName", 100, time.Now()}
	fmt.Println(loginIsIn["user1"].Name)
	fmt.Println(loginIsIn["user1"].LastTime)
	time.Sleep(10 * time.Second)
	fmt.Println(loginIsIn["user1"].CheckLevel(200))
	fmt.Println(loginIsIn["user1"].NotTimeOut(20))
	fmt.Println(loginIsIn["user1"].UpdateLastTime())
	*/
	
}
