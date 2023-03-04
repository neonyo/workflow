package workflow

import (
	"context"
)

// Action is the interface that user action must implement
type Action interface {
	Name() string
	Run(ctx context.Context, params interface{}, res Results) (interface{}, error)
}

// BeforeAction run before run action
//type BeforeAction interface {
//	RunBefore(ctx ExecuteContext, params interface{}) error
//}
//
//// AfterAction run after run action
//type AfterAction interface {
//	RunAfter(ctx ExecuteContext, params interface{}) error
//}
