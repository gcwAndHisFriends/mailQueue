package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var RedisManager=redisManager{}
var UserManager =User{}
var FuncTrie=funcTrie{}

func Init()(e error){
	fmt.Println("正在初始化redis")
	err := RedisManager.init()
	if err != nil {
		return err
	}
	fmt.Println("正在初始化群组数据")

	err = UserManager.init()
	if err != nil {
		return err
	}

	return nil
}

func userListen(){
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		if scanner.Scan(){
			arr:=strings.Fields(scanner.Text())
			orderLen:=len(arr)

			if orderLen==0{
				println("错误,请输入命令")
				continue
			}else if orderLen==1{
				err := FuncTrie.GetAndUse(arr[:1],"")
				if err != nil {
					println(err.Error())
				}
			}else{
				err := FuncTrie.GetAndUse(arr[:orderLen-1],arr[orderLen-1])
				if err != nil {
					println(err.Error())
				}
			}

		}
	}
}

func max(i int, i2 int) int {
	if i<i2{
		return i2
	}else{
		return i
	}
}

//ip port
func main() {
	/*
	err := Init()
	if err != nil {
		return
	}
	fmt.Println("正在启动监听服务器")
	listen := ListenServer{"9090"}
	go listen.init()

	fmt.Println("正在同步信息")
	req:=requestServer{}
	go req.init()

	fmt.Println("启动完成")

	*/
	userListen()
}



