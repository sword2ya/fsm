package fsm

import "fmt"

// FSMInterface 状态机接口
type FSMInterface interface {
	ProcessEvent(e EventInterface) error
	GetCurState() StateInterface
}

// FSM 状态机基础
type FSM struct {
	CurState        StateInterface                                  // 当前状态
	States          map[StateID]StateInterface                      // 状态表
	TransitionTable map[StateID](map[EventID][]TransitionInterface) // 转换表
}

// ProcessEvent 处理外部事件
func (f *FSM) ProcessEvent(e EventInterface) error {
	CurStateID := f.CurState.ID()
	m, ok := f.TransitionTable[CurStateID]
	if !ok {
		return nil
	}
	transitions, ok := m[e.ID()]
	if !ok {
		return nil
	}

	validTransitions := []TransitionInterface{}
	for _, transition := range transitions {
		if transition.Valid(f, e) {
			validTransitions = append(validTransitions, transition)
			break
		}
	}

	if len(validTransitions) == 0 {
		return nil
	}
	if len(validTransitions) > 1 {
		return fmt.Errorf("有效的转换数大于1. StateID:%v 事件ID:%v 事件现场:%v 可用转换:%v ", CurStateID, e.ID(), e.Context(), validTransitions)
	}
	transition := validTransitions[0]
	destination := transition.Destination()
	dState, ok := f.States[destination]
	if !ok {
		return fmt.Errorf("目标状态未找到. StateID:%v", destination)
	}

	f.CurState.Exit(f, e)
	transition.Process(f, e)
	dState.Entry(f, e)
	return nil
}

// GetCurState 获取当前状态
func (f *FSM) GetCurState() StateInterface {
	return f.CurState
}

// AddStates 添加状态
func (f *FSM) AddStates(states ...StateInterface) {
	for _, s := range states {
		f.States[s.ID()] = s
	}
}
