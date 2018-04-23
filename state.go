package fsm

// StateID 状态ID
type StateID int

// StateInterface 状态接口
type StateInterface interface {
	// ID 返回状态 ID
	ID() StateID

	// Entry 进入状态
	Entry(f FSMInterface, e EventInterface)

	// Exit 退出状态
	Exit(f FSMInterface, e EventInterface)
}

// State 状态基础
type State struct {
	SID          StateID
	EntryActions []ActionInterface
	ExitActions  []ActionInterface
}

var _ StateInterface = new(State)

// NewState 创建 State 对象
func NewState(stateID StateID) *State {
	return &State{
		SID:          stateID,
		EntryActions: []ActionInterface{},
		ExitActions:  []ActionInterface{},
	}
}

// ID 返回状态 ID
func (s *State) ID() StateID {
	return s.SID
}

// Entry 进入状态， 执行所有 EntryActions
func (s *State) Entry(f FSMInterface, e EventInterface) {
	for _, action := range s.EntryActions {
		action.Do(f, e)
	}
}

// Exit 退出状态， 执行所有 ExitActions
func (s *State) Exit(f FSMInterface, e EventInterface) {
	for _, action := range s.ExitActions {
		action.Do(f, e)
	}
}

// AddEntryActions 添加进入动作
func (s *State) AddEntryActions(actions ...ActionInterface) {
	s.EntryActions = append(s.EntryActions, actions...)
}

// AddExitActions 添加退出动作
func (s *State) AddExitActions(actions ...ActionInterface) {
	s.ExitActions = append(s.ExitActions, actions...)
}
