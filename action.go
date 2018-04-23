package fsm

// ActionInterface Action 接口
type ActionInterface interface {
	Do(f FSMInterface, e EventInterface)
}
