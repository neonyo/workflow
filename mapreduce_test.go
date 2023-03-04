package workflow

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestWorkflowNoJobs(f *testing.T) {
	//f.Add("")

	//f.Add(1)
	//f.Fuzz(func(t *testing.T, i int) {
	RegisterAction([]Action{
		&Task1Action{},
		&Task2Action{},
		&Task3Action{},
		&Task4Action{},
		&Task5Action{},
		&Task6Action{},
		&Task7Action{},
	})
	//NewTask("task1", nil, ActionMap["task1"]),
	//		NewTask("task2", []string{"task1"}, ActionMap["task2"]),
	//		NewTask("task3", []string{"task1"}, ActionMap["task3"]),
	//		NewTask("task4", []string{"task2"}, ActionMap["task4"]),
	//		NewTask("task5", []string{"task3"}, ActionMap["task5"]),
	//		NewTask("task6", nil, ActionMap["task6"]),
	taskGraph, _ := NewGraph(InitialOption{
		Ins: []struct {
			Name     string
			DependOn string
		}{},
	})
	var params = make(map[string]int)
	params["id"] = 1
	err := taskGraph.Run(context.Background(), params)
	if err != nil {
		f.Fatal(err)
	}
	//})
}

type Task1Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task1Action) Name() string {
	return "task1"
}
func (a *Task1Action) Run(ctx context.Context, params interface{}, res Results) (interface{}, error) {
	fmt.Println("Task1Action start: ", time.Now(), params)
	return "task1 成功了哈哈", nil
}

type Task2Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task2Action) Name() string {
	return "task2"
}
func (a *Task2Action) Run(ctx context.Context, params interface{}, res Results) (interface{}, error) {
	fmt.Println("Task2Action start: ", time.Now(), params)
	return nil, nil
}

type Task3Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task3Action) Name() string {
	return "task3"
}
func (a *Task3Action) Run(ctx context.Context, params interface{}, res Results) (interface{}, error) {
	fmt.Println("Task3Action start: ", time.Now(), params)
	return nil, nil
}

type Task4Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task4Action) Name() string {
	return "task4"
}
func (a *Task4Action) Run(ctx context.Context, params interface{}, res Results) (interface{}, error) {
	fmt.Println("Task4Action start: ", time.Now(), params)
	return nil, nil
}

type Task5Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task5Action) Name() string {
	return "task5"
}
func (a *Task5Action) Run(ctx context.Context, params interface{}, res Results) (interface{}, error) {
	ress, _ := res.Load("task1")
	fmt.Println("Task5Action start: ", time.Now(), params, ress)
	return nil, nil
}

type Task6Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task6Action) Name() string {
	return "task6"
}
func (a *Task6Action) Run(ctx context.Context, params interface{}, res Results) (interface{}, error) {
	fmt.Println("Task6Action start: ", time.Now(), params)
	return nil, errors.New("发生错误了")
}

type Task7Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task7Action) Name() string {
	return "task7"
}
func (a *Task7Action) Run(ctx context.Context, params interface{}, res Results) (interface{}, error) {
	fmt.Println("Task7Action start: ", time.Now(), params)
	return nil, nil
}
