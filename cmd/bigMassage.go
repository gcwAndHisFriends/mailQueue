//用于存储所有相关信息的massage
package main

type bigMassage struct {
	groupId string //群组ID 这条信息属于哪个群组
	senderId string //发送者的ID
	receiverId string //接收者的ID
	massageId string //

	Head string
	Info string

	Status int
	Tag string //标注信息
}
