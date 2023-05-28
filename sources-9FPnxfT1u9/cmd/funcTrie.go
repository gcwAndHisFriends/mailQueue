package main

import (
	"errors"
	"fmt"
)

func outA(str string){
	fmt.Println(str)
}
func outB(str string){
	fmt.Println(str)
	fmt.Println(str)
}
func createGroup(str string){
	err := UserManager.groupList_.createGroup(str)
	if err != nil {
		println(err.Error())
		return
	}
	println("创建成功")
}
func (ft *funcTrie)initFuncTrie(){
	ft.Add([]string{"create","group"},createGroup)
}

type node struct {
	NodeName string
	Function func(str string)
	Next []*node
}

type funcTrie struct{
	Head node
}


func GetFuncIndex(url string,ne []*node)int{
	for i:=0;i<len(ne);i++{
		if ne[i].NodeName==url{
			return i
		}
	}
	return -1
}
func nullFunction (str string){
	fmt.Println("调用了一个错误的函数"+str)
}

func (ft *funcTrie) Add(url []string,Function func(str string))bool {
	target:=&ft.Head
	ok:=false //是否找到了位置
	for i:=0;i<len(url);i++{
		idx:=GetFuncIndex(url[i],target.Next)
		if idx!=-1 { //存在节点
			target = target.Next[idx]
		}else { //不存在节点
			ok=true
			var tmp1 []*node
			tmp:=node{url[i],nullFunction,tmp1}
			target.Next=append(target.Next, &tmp)
			target=&tmp
		}
	}
	if ok==false{
		fmt.Println("两个函数的命令相同")
		return false
	}
	target.Function=Function
	return true
}
func (ft *funcTrie) GetAndUse(url []string,order string)(e error){
	target:=&ft.Head
	for i:=0;i<len(url);i++{
		println(url[i])
	}
	for i:=0;i<len(url);i++{
		idx:=GetFuncIndex(url[i],target.Next)
		if idx!=-1 { //存在节点
			target = target.Next[idx]
		}else { //不存在节点
			return errors.New("调用了不存在的函数")
		}
	}
	target.Function(order)
	return nil
}

