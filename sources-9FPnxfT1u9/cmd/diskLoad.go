package main

import (
	"fmt"
	"os"
	"runtime"
)

/*
./group
./config

 */

type diskLoader struct{
	baseUrl string
	osType string
	osSplit string
}

func (req *diskLoader)checkAndCreateDir(){
	groupDir:=req.baseUrl+req.osSplit+"group.txt"
	configDir:=req.baseUrl+req.osSplit+"config.txt"

	if _, err := os.Stat(groupDir); err != nil{
		if os.IsNotExist(err) {
			_, _=os.Create(groupDir)
		}
	}

	if _, err := os.Stat(configDir); err != nil{
		if os.IsNotExist(err) {
			_, _=os.Create(configDir)
		}
	}

}

func (req *diskLoader)init(){
	req.osType=runtime.GOOS
	if req.osType=="windows"{
		req.osSplit="\\"
	}else{
		req.osSplit="/"
	}
	req.baseUrl,_=os.Getwd()
	fmt.Println("当前的工作路径"+req.baseUrl)
	req.checkAndCreateDir()
}

func (req *diskLoader)writeWork(data *[]byte){

}