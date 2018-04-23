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

// Event 事件基础
type Event struct {
	eventID EventID
	context EventContext
}

// NewEvent 创建 Event 对象
func NewEvent(eventID EventID, context EventContext) *Event {
	return &Event{
		eventID: eventID,
		context: context,
	}
}

// ID 返回事件ID
func (e *Event) ID() EventID {
	return e.eventID
}

// Context 返回事件现场
func (e *Event) Context() EventContext {
	return e.context
}
