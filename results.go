package workflow

import (
	"encoding/json"
	"sync"
)

// Results holds a reference to the intermediate results of a workflow
// execution. It used to task the result of one function into its
// dependencies.
type Results struct {
	resMap *sync.Map
}

// Load retrieves the result for a particular task.
func (r Results) Load(taskName string) (val Result, ok bool) {
	var value interface{}
	value, ok = r.resMap.Load(taskName)
	d, _ := json.Marshal(value)
	json.Unmarshal(d, &val)
	return val, ok
}

func (r Results) Store(taskName string, result Result) {
	r.resMap.Store(taskName, result)
}

type Result struct {
	Status     Status
	NextStatus Status
	Data       interface{}
}
