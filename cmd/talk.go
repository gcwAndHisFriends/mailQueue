package main

import (
	"sync"
)

const ( //status
	statusNotViewed=iota //尚未被查看
	statusViewed //已经被查看,未开始
	statusStart //已经开始
	statusEnd //结束
	statusPaused //暂停
)
type massage struct {
	Id int //work有独一无二标签
	Head string
	Info string
	Status int
	Tag string //标注信息
}

type work struct {
	massage
	mut sync.RWMutex //读写锁
}


const ( //创建的类型 mode
	modeSelf=iota //值从0开始递增
	modeOther //他人发起的请求
)

type talk struct { //talk应该有一个唯一标识
	TalkId int
	Mode int //自己创建/别人发起
	WorkId string //接受
	SendId string //如果是自己发起,那么为空
	WorkList []work
	//func getPreTalkId(GroupId,WorkId,SendId) redis
}

func (t *talk)equal(workId string,sendId string,talkId int)bool{
	return t.TalkId==talkId&&t.SendId==sendId&&t.WorkId==workId
}


