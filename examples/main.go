package main

import (
	"context"
	"fmt"
	"github.com/neonyo/workflow"
	"time"
)

func main() {
	workflow.RegisterAction([]workflow.Action{
		&Task1Action{},
		&Task2Action{},
		&Task3Action{},
		&Task4Action{},
		&Task5Action{},
		&Task6Action{},
		&Task7Action{},
	})
	var tasks []workflow.Task
	tasks = append(tasks, workflow.NewTask("task1", nil, workflow.ActionMap["task1"]))
	tasks = append(tasks, workflow.NewTask("task2", nil, workflow.ActionMap["task2"]))
	tasks = append(tasks, workflow.NewTask("task3", nil, workflow.ActionMap["task3"]))
	tasks = append(tasks, workflow.NewTask("task4", nil, workflow.ActionMap["task4"]))
	tasks = append(tasks, workflow.NewTask("task5", nil, workflow.ActionMap["task5"]))
	var task workflow.InitialTaskOption
	task.TaskOption = append(task.TaskOption, workflow.TaskOption{Name: "task1", DependOn: ""})
	task.TaskOption = append(task.TaskOption, workflow.TaskOption{Name: "task2", DependOn: "task1"})
	task.TaskOption = append(task.TaskOption, workflow.TaskOption{Name: "task3", DependOn: "task2"})
	task.TaskOption = append(task.TaskOption, workflow.TaskOption{Name: "task4", DependOn: "task3"})
	task.TaskOption = append(task.TaskOption, workflow.TaskOption{Name: "task5", DependOn: "task3"})
	taskGraph, _ := workflow.NewGraph(task)
	var params = make(map[string]string)
	params["c"] = "12"
	ctx := context.Background()
	//ctx, _ := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "task1", workflow.Success)
	ctx = context.WithValue(ctx, "task2", workflow.Success)
	ctx = context.WithValue(ctx, "task3", workflow.Success)
	ctx = context.WithValue(ctx, "task4", workflow.Init)
	ctx = context.WithValue(ctx, "task5", workflow.Skipped)
	//context.WithValue(ctx, "task1", "122")
	taskGraph.Run(ctx, params)
}

type Task1Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task1Action) Name() string {
	return "task1"
}

func (a *Task1Action) Run(ctx context.Context, params interface{}, res workflow.Results) (workflow.Result, error) {
	fmt.Println("系统处理 start: ", time.Now(), params)
	return workflow.Result{}, nil
}

type Task2Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task2Action) Name() string {
	return "task2"
}
func (a *Task2Action) Run(ctx context.Context, params interface{}, res workflow.Results) (workflow.Result, error) {
	fmt.Println("组长 start: ", time.Now(), params)
	return workflow.Result{}, nil
}

type Task3Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task3Action) Name() string {
	return "task3"
}
func (a *Task3Action) Run(ctx context.Context, params interface{}, res workflow.Results) (workflow.Result, error) {
	fmt.Println("是否运维捞单 start: ", time.Now(), params)
	//context.WithValue(ctx, "status", "blocked")
	return workflow.Result{
		Status:     workflow.Success,
		NextStatus: workflow.Blocked,
	}, nil
}

type Task4Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task4Action) Name() string {
	return "task4"
}
func (a *Task4Action) Run(ctx context.Context, params interface{}, res workflow.Results) (workflow.Result, error) {
	fmt.Println("捕头 start: ", time.Now(), params)
	return workflow.Result{}, nil
}

type Task5Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task5Action) Name() string {
	return "task5"
}
func (a *Task5Action) Run(ctx context.Context, params interface{}, res workflow.Results) (workflow.Result, error) {

	//is, _ := res.Load("task3")
	//if is.(bool) {
	//	fmt.Println("运维捞单 start: ", time.Now(), params)
	//} else {
	//	fmt.Println("运维捞单 不运行: ", time.Now(), params)
	//}
	fmt.Println("运维捞单 start: ", time.Now(), params)
	return workflow.Result{}, nil
}

type Task6Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task6Action) Name() string {
	return "task6"
}
func (a *Task6Action) Run(ctx context.Context, params interface{}, res workflow.Results) (workflow.Result, error) {
	fmt.Println("Task6Action start: ", time.Now(), params)
	return workflow.Result{}, nil
}

type Task7Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task7Action) Name() string {
	return "task7"
}
func (a *Task7Action) Run(ctx context.Context, params interface{}, res workflow.Results) (workflow.Result, error) {
	fmt.Println("Task7Action start: ", time.Now(), params)
	return workflow.Result{}, nil
}
