package znet

import (
	"TankServer/zinx/utils"
	"TankServer/zinx/ziface"
	"fmt"
	"strconv"
)

/*
	消息处理模块的实现
 */
type MsgHandle struct {
	// 存放每一个MsgID所对应的处理方法
	Apis map[uint32]ziface.IRouter
	// 负责Worker区任务的消息队列
	TaskQueue []chan ziface.IRequest
	// 业务工作Worker池的worker数量
	WorkerPoolSize uint32
}

// 初始化 创建MsgHandle方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
		TaskQueue: make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize, // 从全局配置中获取
	}
}

// 调度、执行对应Router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	// 1 从Request中找到msgID
	handle, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), "is not found! need register")
		return
	}
	// 2 根据MsgID 调度对应router业务即可
	handle.PreHandle(request)
	handle.Handle(request)
	handle.PostHandle(request)
}

// 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	// 判断 当前msg绑定的api处理方法是否已经存在
	if _, ok := mh.Apis[msgID]; ok {
		// id已经存在
		panic("request api, msgID = " + strconv.Itoa(int(msgID)))
	}
	// 添加msg与API的绑定关系
	mh.Apis[msgID] = router
	fmt.Println("Add api MsgID = ", msgID, " success!")
}

// 启动一个Worker工作池(开启工作池的动作只能发生一次 一个zinx框架只能有一个worker工作池)
func (mh *MsgHandle)StartWorkerPool()  {
	// 根据workerPoolSize 分别开启Worker 每个Worker用一个go来承载
	for i := 0; i< int(mh.WorkerPoolSize); i++ {
		// 一个worker被启动
		// 1 当前的worker对应的channel消息队列 开辟空间 第0个worker就用第0个channel
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// 2 启动当前的Worker 阻塞等待消息从channel传递过来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// 启动一个Worker工作流程
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID = ", workerID, " is started ...")

	// 不断阻塞等待对应消息队列的消息
	for {
		select {
		// 如果有消息过来 出列的就是一个客户端的Request 执行当前Request所绑定的业务
		case request := <- taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

// 将消息交给TaskQueue 由Worker进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest)  {
	// 1 将消息平均分配给不同的worker
	// 根据客户端建立的ConnID来进行分配
	// 平均轮询法则
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID = ", request.GetConnection().GetConnID(), " request MsgID = ", request.GetMsgID(), " to WorkerID = ", workerID)
	// 2 将消息发送给对应的worker的TaskQueue即可
	mh.TaskQueue[workerID] <- request
}