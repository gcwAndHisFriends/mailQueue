想要在团队里展示自己的任务列表(自己和领导才能看到详细的,他人
只能看到标题)，想要像处理邮件那样有序地完成工作。

加入一个工作群,自己创建/领导指派/同事委托。
所有的任务将呈现在任务区中，可以调整任务的状态，
>未开始 正在进行 已结束

有三类中断状态
>任务暂停(自己) 即将结束(收尾) 等待其他任务结束

亦可在一件任务下创建一系列任务会话。在一件任务里点击新建任务,即将两件任务放入同一任务会话中

领导指派/同事委托任务。除了他人通过该软件发出，也可以自己创建。这类任务在完成后需要添加"回复"功能。双方均可点击结束以关闭会话

意见事物,有三个等级

```
紧急任务:立即跳出提醒(可设置),并列在正在进行的任务之后的第一个

正常任务:新出列在任务队列的最后一个

不紧急任务:不会自动添加到任务栏中,主动点击添加后才会出现
```

暂时做一个命令行



```
start 开始运行

stop 结束运行

create group <group_name>

bind <IP> <port=53012> <group_name> 绑定IP,端口(默认53012),群组名,这个三元组是必须unique的 如果group_name唯一,那么切换group_name时直接切换,否则要选择IP和port

cd <群组名> 进入一个群组
cd .. 离开当前群组

ls -group 展示群组内人员 
ls -work 展示任务
ls -other <群组内人员id>展示某人在该群组内任务

cat -work <任务ID> 展示任务内容(输入ID,数字)
cat -other <别人ID> 展示其他人的(公开任务列表)

active <任务ID> 开始任务
stop <任务ID> 结束任务

```





