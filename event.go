package fsm

// EventID 事件 ID
type EventID int

// EventContext 事件现场
type EventContext interface{}

// EventInterface Event 接口
type EventInterface interface {
	ID() EventID
	Context() EventContext
}
