package main

import (
	"errors"
	"fmt"
)

type User struct{
	Uid string
	groupList_ groupList
}

func (u *User)register()(e error){
	fmt.Println("第一次登录,请输入你的id")
	fmt.Println("一个群组内的id应是唯一的")
	_, err := fmt.Scanln(&u.Uid)
	if err != nil {
		return errors.New("错误")
	}
	return nil
}

func (u *User)init()(e error){
	get, err := RedisManager.get("Uid")
	if err != nil {
		return u.register()
	}
	u.Uid=get
	err = u.groupList_.init()
	if err != nil {
		return err
	}
	return nil
}