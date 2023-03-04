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
	taskGraph, _ := workflow.NewGraph(workflow.InitialOption{
		Ins: []struct {
			Name     string
			DependOn string
		}{},
	})
	var params = make(map[string]string)
	params["c"] = "12"
	taskGraph.Run(context.Background(), params)
}

type Task1Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task1Action) Name() string {
	return "task1"
}
func (a *Task1Action) Run(ctx context.Context, params interface{}, res workflow.Results) (interface{}, error) {
	fmt.Println("Task1Action start: ", time.Now(), params)
	return nil, nil
}

type Task2Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task2Action) Name() string {
	return "task2"
}
func (a *Task2Action) Run(ctx context.Context, params interface{}, res workflow.Results) (interface{}, error) {
	fmt.Println("Task2Action start: ", time.Now(), params)
	return nil, nil
}

type Task3Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task3Action) Name() string {
	return "task3"
}
func (a *Task3Action) Run(ctx context.Context, params interface{}, res workflow.Results) (interface{}, error) {
	fmt.Println("Task3Action start: ", time.Now(), params)
	return nil, nil
}

type Task4Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task4Action) Name() string {
	return "task4"
}
func (a *Task4Action) Run(ctx context.Context, params interface{}, res workflow.Results) (interface{}, error) {
	fmt.Println("Task4Action start: ", time.Now(), params)
	return nil, nil
}

type Task5Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task5Action) Name() string {
	return "task5"
}
func (a *Task5Action) Run(ctx context.Context, params interface{}, res workflow.Results) (interface{}, error) {
	fmt.Println("Task5Action start: ", time.Now(), params)
	return nil, nil
}

type Task6Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task6Action) Name() string {
	return "task6"
}
func (a *Task6Action) Run(ctx context.Context, params interface{}, res workflow.Results) (interface{}, error) {
	fmt.Println("Task6Action start: ", time.Now(), params)
	return nil, nil
}

type Task7Action struct {
}

// Name define the unique action identity, it will be used by Task
func (a *Task7Action) Name() string {
	return "task7"
}
func (a *Task7Action) Run(ctx context.Context, params interface{}, res workflow.Results) (interface{}, error) {
	fmt.Println("Task7Action start: ", time.Now(), params)
	return nil, nil
}
