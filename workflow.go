package workflow

import (
	"context"
	"fmt"
	"strings"
)

var (
	ActionMap = make(map[string]Action, 0)
	empty     struct{}
)

//	dag := &entity.Dag{
//			BaseInfo: entity.BaseInfo{
//				ID: "test-dag",
//			},
//			Name: "test",
//			Tasks: []entity.Task{
//				{ID: "task1", ActionName: "PrintAction"},
//				{ID: "task2", ActionName: "PrintAction", DependOn: []string{"task1"}},
//				{ID: "task3", ActionName: "PrintAction", DependOn: []string{"task2"}},
//			},
//			Status: entity.DagStatusNormal,
//		}
//
//	type Dag struct {
//		BaseInfo `yaml:",inline" json:",inline" bson:"inline"`
//		Name     string    `yaml:"name,omitempty" json:"name,omitempty" bson:"name,omitempty"`
//		Desc     string    `yaml:"desc,omitempty" json:"desc,omitempty" bson:"desc,omitempty"`
//		Cron     string    `yaml:"cron,omitempty" json:"cron,omitempty" bson:"cron,omitempty"`
//		Vars     DagVars   `yaml:"vars,omitempty" json:"vars,omitempty" bson:"vars,omitempty"`
//		Status   DagStatus `yaml:"status,omitempty" json:"status,omitempty" bson:"status,omitempty"`
//		Tasks    []Task    `yaml:"tasks,omitempty" json:"tasks,omitempty" bson:"tasks,omitempty"`
//	}
//type Task struct {
//	Name     string
//	DependOn []string
//	Fn       Action
//}

type InitialOption struct {
	Ins []struct {
		Name     string
		DependOn string
	}
}

// RegisterAction 注册节点action
func RegisterAction(acts []Action) {
	for i := range acts {
		ActionMap[acts[i].Name()] = acts[i]
	}
}

type TaskFn = func(ctx context.Context, params interface{}, results Results) (interface{}, error)

// A Task is a unit of work along with a name and set of dependencies.
type Task struct {
	name string
	fn   Action
	deps map[string]struct{}
}

func NewTask(name string, deps []string, fn Action) Task {
	depSet := make(map[string]struct{}, len(deps))

	for _, dep := range deps {
		depSet[dep] = empty
	}

	return Task{
		name: name,
		fn:   fn,
		deps: depSet,
	}
}

func withTasks(options InitialOption) ([]Task, error) {
	var tasks []Task
	for _, v := range options.Ins {
		if _, ok := ActionMap[v.Name]; !ok {
			return nil, fmt.Errorf("action name:%s nɒt faʊnd", v.Name)
		}
		var deps []string
		if v.DependOn != "" {
			deps = strings.Split(v.DependOn, ",")
		}
		tasks = append(tasks, NewTask(v.Name, deps, ActionMap[v.Name]))
	}
	return tasks, nil
}

type Graph struct {
	tasks map[string]Task
	// Map of task name to set of tasks that depend on it.
	// Map<String, Set<String>>
	taskToDependants map[string]map[string]struct{}
}

func NewGraph(options InitialOption) (Graph, error) {
	var (
		tasks []Task
		err   error
	)
	tasks, err = withTasks(options)
	if err != nil {
		return Graph{}, err
	}
	taskMap := make(map[string]Task, len(tasks))
	taskToDependants := make(map[string]map[string]struct{}, len(tasks))

	for _, task := range tasks {
		name := task.name
		taskMap[name] = task
		for depName := range task.deps {
			depSet, ok := taskToDependants[depName]
			if !ok {
				depSet = make(map[string]struct{})
			}

			depSet[name] = empty
			taskToDependants[depName] = depSet
		}
	}

	g := Graph{
		tasks:            taskMap,
		taskToDependants: taskToDependants,
	}

	err = g.isWellFormed()

	return g, err
}

func (g Graph) isWellFormed() error {
	tasksWithNoDeps := make([]string, 0, len(g.tasks))
	taskToNumDeps := make(map[string]int32, len(g.tasks))

	// Count the number of dependencies and mark as ready to execute if no
	// deps are present.
	for name, task := range g.tasks {
		numDeps := len(task.deps)

		if numDeps == 0 {
			tasksWithNoDeps = append(tasksWithNoDeps, name)
		}

		taskToNumDeps[name] = int32(numDeps)
	}

	visitedJobs := 0

	for i := 0; i < len(tasksWithNoDeps); i++ {
		name := tasksWithNoDeps[i]
		visitedJobs++

		for dep := range g.taskToDependants[name] {
			taskToNumDeps[dep]--

			// If a job has no dependencies unfulfilled, we can visit it.
			if taskToNumDeps[dep] == 0 {
				tasksWithNoDeps = append(tasksWithNoDeps, dep)
			}
		}
	}

	// If we have failed to visit all jobs, or we have visited some more than
	// once somehow, we either are missing a dependency or we have a cycle.
	// Either way, we can't execute the graph.
	if visitedJobs != len(g.tasks) {
		return InvalidGraphError{}
	}

	return nil
}

func (g Graph) Run(ctx context.Context, params interface{}) error {
	return newExecutionCtx(ctx, g).run(params)
}
