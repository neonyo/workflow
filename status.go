package workflow

type Status string

func (n Status) String() string {
	return string(n)
}

const (
	Init     Status = "init"     //初始化
	Wait     Status = "wait"     //等待执行
	Running  Status = "running"  //执行中
	Ending   Status = "ending"   //结束
	Retrying Status = "retrying" //重试
	Failed   Status = "failed"   //执行失败
	Success  Status = "success"  //已执行成功
	Blocked  Status = "blocked"  //任务已阻塞，需要人工启动
	Skipped  Status = "skipped"  //跳过
)
