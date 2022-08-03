package main

import (
	"fmt"
)

/*
Состояние — это поведенческий паттерн проектирования, который позволяет
объектам менять поведение в зависимости от своего состояния. Извне создаётся
впечатление, что изменился класс объекта.
Основная идея в том, что программа может находиться в одном из нескольких
состояний, которые всё время сменяют друг друга. Набор этих состояний,
а также переходов между ними, предопределён и конечен. Находясь в разных
состояниях, программа может по-разному реагировать на одни и те же события,
которые происходят с ней.
Плюсы:
- Избавляет от множества больших условных операторов машины состояний.
- Концентрирует в одном месте код, связанный с определённым состоянием.
Минусы:
- Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

func main(){
	StateDemo()
}

type State interface {
	executeState(c *Context)
}

type Context struct {
	StepIndex int
	StepName  string
	Current   State
}

type StartState struct{}

func (s *StartState) executeState(c *Context) {
	c.StepIndex = 1
	c.StepName = "start"
	c.Current = &StartState{}
}

type InprogressState struct{}

func (s *InprogressState) executeState(c *Context) {
	c.StepIndex = 2
	c.StepName = "inprogress"
	c.Current = &InprogressState{}
}

type StopState struct{}

func (s *StopState) executeState(c *Context) {
	c.StepIndex = 3
	c.StepName = "stop"
	c.Current = &StopState{}
}

func StateDemo() {
	context := &Context{}
	var state State

	state = &StartState{}
	state.executeState(context)
	fmt.Println("state: ", *context)

	state = &InprogressState{}
	state.executeState(context)
	fmt.Println("state: ", *context)

	state = &StopState{}
	state.executeState(context)
	fmt.Println("state: ", *context)
}