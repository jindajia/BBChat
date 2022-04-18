package handlers

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NUmStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func generatePasswd() string {
	//初始化密码切片
	var passwd []byte = make([]byte, 10, 10)
	//源字符串
	var sourceStr string
	//判断字符类型,如果是数字

	sourceStr = fmt.Sprintf("%s%s%s", NUmStr, CharStr, SpecStr)

	for i := 0; i < 10; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return string(passwd)
}
func generateroomno() string {
	//初始化密码切片
	var roomno []byte = make([]byte, 8, 8)
	//源字符串
	var sourceStr string
	//判断字符类型,如果是数字

	sourceStr = fmt.Sprintf("%s", NUmStr)

	for i := 0; i < 8; i++ {
		index := rand.Intn(len(sourceStr))
		roomno[i] = sourceStr[index]
	}

	return string(roomno)
}

func generatepassword() string {
	//随机种子
	rand.Seed(time.Now().UnixNano())
	//test1()
	passwd := generatePasswd()
	return passwd
}

func generateroomnoo() string {
	rand.Seed(time.Now().UnixNano())
	//test1()
	roomno := generateroomno()

	return roomno
}
