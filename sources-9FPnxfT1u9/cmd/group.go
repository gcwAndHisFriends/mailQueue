package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
)

type person struct{
	PersonId string //
	GroupId string //personId与groupId唯一确定一个对象

	PersonIp string //临时的Ip,会更新
	PersonPort string //临时port,可能更新
}

//群组锁,大粗锁
var groupLock []sync.RWMutex
type group struct{
	GroupId string //唯一标识
	Leader []person
	Worker []person
	Talk []talk //由groupId workId SendId 确定TalkId的自增
}

var GroupListLock=sync.RWMutex{} //对groups的结构进行修改时写锁,修改其中信息时读锁
type groupList struct{
	groups []group
}

func (g *groupList)init()(e error){
	exists, err := RedisManager.exists("mailQueueGroupInfo")
	if err != nil {
		return err
	}
	if exists==false{
		_, err := RedisManager.set("mailQueueGroupInfo","")
		if err != nil {
			return err
		}
	}

	tmp,err:=RedisManager.get("mailQueueGroupInfo")
	if err!=nil{
		return err
	}

	err = json.Unmarshal([]byte(tmp), &g)
	if err!=nil{
		return err
	}

	go g.writeLoop()
	return nil
}

func (g *groupList)write()(e error){
	GroupListLock.Lock()
	defer GroupListLock.Unlock()
	bs, err := json.Marshal(&g)
	if err!=nil{
		return err
	}
	_, err = RedisManager.set("mailQueueGroupInfo", string(bs))
	if err!=nil{
		fmt.Println("redis写入错误")
		return err
	}
	return nil
}

func (g *groupList)writeLoop(){
	for true {
		_ = g.write()
		time.Sleep(10*time.Second)
	}
}

//需要使用到groupList的写锁
func (g *groupList)addGroup(group2 group){ //加入一个新群组
	GroupListLock.Lock()
	defer GroupListLock.Unlock()
	g.groups=append(g.groups, group2)
}

//直接将整个组发过来
func (g *groupList)updateLeader(leader []person,gId string)(e error){
	GroupListLock.RLock()
	defer GroupListLock.RUnlock()
	for i:=0;i<len(g.groups);i++{
		if g.groups[i].GroupId==gId {
			groupLock[i].Lock()
			g.groups[i].Leader=leader
			return nil
		}
	}
	return errors.New("更新群组leader错误:找不到该群组")
}

//直接将整个组发过来
func (g *groupList)updateWorker(worker []person,gId string)(e error){
	GroupListLock.RLock()
	defer GroupListLock.RUnlock()
	for i:=0;i<len(g.groups);i++{
		if g.groups[i].GroupId==gId{
			groupLock[i].Lock()
			g.groups[i].Worker=worker
			return nil
		}
	}
	return errors.New("更新群组成员时错误:找不到该群组")
}

func (g *groupList)updateTalk(t talk,groupId string,workId string,sendId string,talkId int)(e error){
	GroupListLock.RLock()
	defer GroupListLock.RUnlock()
	for i:=0;i<len(g.groups);i++{
		if g.groups[i].GroupId==groupId{
			for j:=0;j<len(g.groups[i].Talk);j++{
				if g.groups[i].Talk[j].equal(workId,sendId,talkId){
					g.groups[i].Talk[j]=t
					return nil
				}
			}
		}
	}
	return errors.New("找不到该聊天信息")
}

//判断是否存在Id
func (g *groupList)liveId(str string)bool{
	for i:=0;i<len(g.groups);i++{
		if g.groups[i].GroupId==str{
			return true
		}
	}
	return false
}

func (g *groupList)createGroup(str string)(e error){
	if g.liveId(str){
		return errors.New("错误,存在相同Id群组")
	}
	g.addGroup(group{str,[]person{},[]person{},[]talk{}})
	return nil
}
