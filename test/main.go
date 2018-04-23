package main

import (
	fsm "FSM"
	"fmt"
	"math/rand"
	"time"
)

const (
	Colding = fsm.StateID(iota) // 冷却中
	Hoting                      // 加热中
	Idle                        // 空闲
)

const (
	TooHot  = fsm.EventID(iota) // 太热了
	TooCold                     // 太冷了
	AtTemp                      // 温度合适
)

// Airconditioner 温度调节器
type Airconditioner struct {
	fsm.FSM
}

type NotifyEnterColdStateAction struct{}

func (a *NotifyEnterColdStateAction) Do(f fsm.FSMInterface, e fsm.EventInterface) {
	fmt.Println("colding enter")
}

type NotifyLeaveColdStateAction struct{}

func (a *NotifyLeaveColdStateAction) Do(f fsm.FSMInterface, e fsm.EventInterface) {
	fmt.Println("condling leave")
}

type NotifyEnterHotStateAction struct{}

func (a *NotifyEnterHotStateAction) Do(f fsm.FSMInterface, e fsm.EventInterface) {
	fmt.Println("hoting enter")
}

type NotifyLeaveHotStateAction struct{}

func (a *NotifyLeaveHotStateAction) Do(f fsm.FSMInterface, e fsm.EventInterface) {
	fmt.Println("hoting leave")
}

type NotifyEnterIdleStateAction struct{}

func (a *NotifyEnterIdleStateAction) Do(f fsm.FSMInterface, e fsm.EventInterface) {
	fmt.Println("idle enter")
}

type NotifyLeaveIdleStateAction struct{}

func (a *NotifyLeaveIdleStateAction) Do(f fsm.FSMInterface, e fsm.EventInterface) {
	fmt.Println("idle leave")
}

type NotifyIdle2HotingAction struct{}

func constructAC() *Airconditioner {
	ac := Airconditioner{
		FSM: fsm.FSM{
			States:          make(map[fsm.StateID]fsm.StateInterface),
			TransitionTable: make(map[fsm.StateID](map[fsm.EventID][]fsm.TransitionInterface)),
		},
	}

	// 冷却中状态
	coldState := fsm.NewState(Colding)
	coldState.AddEntryActions(&NotifyEnterColdStateAction{})
	coldState.AddExitActions(&NotifyLeaveColdStateAction{})

	// 加热中状态
	hotState := fsm.NewState(Hoting)
	hotState.AddEntryActions(&NotifyEnterHotStateAction{})
	hotState.AddExitActions(&NotifyLeaveHotStateAction{})

	// 空闲状态
	idleState := fsm.NewState(Idle)
	idleState.AddEntryActions(&NotifyEnterIdleStateAction{})
	idleState.AddExitActions(&NotifyLeaveIdleStateAction{})

	// 将所有状态添加到空气调节器中
	ac.AddStates(coldState, idleState, hotState)

	// 空闲状态到加热状态的转换
	idle2hoting := fsm.NewTransition(Hoting)
	ac.AddTransition(Idle, TooCold, idle2hoting)

	// 加热中到空闲
	hoting2idle := fsm.NewTransition(Idle)
	ac.AddTransition(Hoting, AtTemp, hoting2idle)

	// 空闲状态到冷却状态的转换
	idle2colding := fsm.NewTransition(Colding)
	ac.AddTransition(Idle, TooHot, idle2colding)

	// 冷却中到空闲
	colding2idle := fsm.NewTransition(Idle)
	ac.AddTransition(Colding, AtTemp, colding2idle)

	// 加热中到冷却中的转换
	hoting2colding := fsm.NewTransition(Colding)
	ac.AddTransition(Hoting, TooHot, hoting2colding)

	// 冷却中到加热中
	colding2hoting := fsm.NewTransition(Hoting)
	ac.AddTransition(Colding, TooCold, colding2hoting)

	ac.SetState(Idle)
	return &ac
}

func main() {
	ac := constructAC()
	events := []fsm.EventID{TooCold, TooHot, AtTemp}

	for i := 1; i < 100; i++ {
		index := rand.Intn(len(events))
		event := fsm.NewEvent(events[index], nil)
		fmt.Println("执行事件：", event)
		if err := ac.ProcessEvent(event); err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
