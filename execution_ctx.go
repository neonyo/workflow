package workflow

import (
	"context"
	"sync"
)

// executionCtx holds everything needed to execute a Graph.
// The Graph type can be thought of as stateless, whereas the
// executionCtx type can be thought of as mutable. This allows the Graph
// to be executed multiple times without needing any external cleanup, so long
// as the caller's tasks are idempotent.
type executionCtx struct {
	wg            sync.WaitGroup
	g             Graph
	taskToNumdeps map[string]*Int32
	err           error
	ctx           context.Context
	cancel        context.CancelFunc
	errCounter    *Int32 // There are no atomic errors, sadly
	params        interface{}
	results       Results
}

func newExecutionCtx(ctx context.Context, g Graph) *executionCtx {
	iCtx, cancel := context.WithCancel(ctx)
	return &executionCtx{
		taskToNumdeps: make(map[string]*Int32, len(g.tasks)),
		g:             g,
		ctx:           iCtx,
		cancel:        cancel,
		errCounter:    NewInt32(0),
		results:       Results{resMap: &sync.Map{}},
	}
}

func (ec *executionCtx) run(params interface{}) error {
	ec.params = params
	for _, t := range ec.g.tasks {
		ec.taskToNumdeps[t.name] = NewInt32(int32(len(t.deps)))
	}

	for _, t := range ec.g.tasks {
		// When a task has no dependencies, it is free to be run.
		if ec.taskToNumdeps[t.name].Load() == 0 {
			ec.enqueueTask(t)
		}
	}

	ec.wg.Wait()
	ec.cancel()

	return ec.err
}

func (ec *executionCtx) hasEncounteredErr() bool {
	return ec.errCounter.Load() != 0
}

func (ec *executionCtx) withActionError() {

}

func (ec *executionCtx) markFailure(err error) {
	// Return only the first error encountered
	if !ec.hasEncounteredErr() {
		ec.err = err
	}

	ec.cancel()
}

func (ec *executionCtx) enqueueTask(t Task) {
	ec.wg.Add(1)
	go ec.runTask(t)
}

func (ec *executionCtx) runTask(t Task) {
	defer ec.wg.Done()
	// Do not execute if we have encountered an error.
	if ec.hasEncounteredErr() {
		return
	}
	var (
		res Result
		err error
	)
	if ec.ctx.Value(t.name) != Success && ec.ctx.Value(t.name) != Skipped {
		res, err = t.fn.Run(ec.ctx, ec.params, ec.results)
	}
	if err != nil {
		ec.markFailure(err)
		// Do not queue up additional tasks after encountering an error
		return
	}

	ec.results.Store(t.name, res)
	if res.NextStatus == Blocked {
		return
	}

	for dep := range ec.g.taskToDependants[t.name] {
		if ec.taskToNumdeps[dep].Add(-1) == int32(0) {
			ec.enqueueTask(ec.g.tasks[dep])
		}
	}
}
