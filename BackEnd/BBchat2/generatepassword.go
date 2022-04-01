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

//检测字符串中的空格
func test1() {
	for i := 0; i < len(CharStr); i++ {
		if CharStr[i] != ' ' {
			fmt.Printf("%c", CharStr[i])
		}
	}
}

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

func generatepassword() string {
	//随机种子
	rand.Seed(time.Now().UnixNano())
	//test1()
	passwd := generatePasswd()
	return passwd
}
