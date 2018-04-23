package fsm

import (
	"fmt"
	"testing"
)

const (
	Colding StateID = 1
	Hoting  StateID = 2
	Idle    StateID = 3
)

type Airconditioner struct {
	FSM
}

type NotifyEnterColdStateAction struct{}

func (a *NotifyEnterColdStateAction) Do(f FSMInterface, e EventInterface) {
	fmt.Println("colding enter")
}

type NotifyLeaveColdStateAction struct{}

func (a *NotifyLeaveColdStateAction) Do(f FSMInterface, e EventInterface) {
	fmt.Println("condling leave")
}

type NotifyEnterHotStateAction struct{}

func (a *NotifyEnterHotStateAction) Do(f FSMInterface, e EventInterface) {
	fmt.Println("hoting enter")
}

type NotifyLeaveHotStateAction struct{}

func (a *NotifyLeaveHotStateAction) Do(f FSMInterface, e EventInterface) {
	fmt.Println("hoting leave")
}

type NotifyIdleHotStateAction struct{}

func (a *NotifyIdleHotStateAction) Do(f FSMInterface, e EventInterface) {
	fmt.Println("idle enter")
}

type NotifyIdleHotStateAction struct{}

func (a *NotifyIdleHotStateAction) Do(f FSMInterface, e EventInterface) {
	fmt.Println("hoting leave")
}

func constructAC() {
	ac := Airconditioner{}

	coldState := State{
		SID: Colding,
	}
	coldState.AddEntryActions(&NotifyEnterColdStateAction{})
	coldState.AddExitActions(&NotifyLeaveColdStateAction{})

	hotState := State{
		SID: Hoting,
	}
	hotState.AddEntryActions(&NotifyEnterHotStateAction{})
	hotState.AddExitActions(&NotifyLeaveHotStateAction{})

}

func Test_fsm1(t *testing.T) {

}
