package fsm

// Validator 验证器
type Validator interface {
	Valid(f FSMInterface, e EventInterface) bool
}

// TransitionInterface 转换接口
type TransitionInterface interface {
	Validator
	Destination() StateID // 目标状态 ID
	Process(f FSMInterface, e EventInterface)
}

// Transition 转换基础
type Transition struct {
	destination StateID
	validators  []Validator
	actions     []ActionInterface
}

// NewTransition 创建 Transition 对象
func NewTransition(destination StateID) *Transition {
	return &Transition{
		destination: destination,
		validators:  []Validator{},
		actions:     []ActionInterface{},
	}
}

var _ TransitionInterface = new(Transition)

// Valid 验证转换是否能进行
func (t *Transition) Valid(f FSMInterface, e EventInterface) bool {
	for _, v := range t.validators {
		if !v.Valid(f, e) {
			return false
		}
	}
	return true
}

// Process 执行转换 actions
func (t *Transition) Process(f FSMInterface, e EventInterface) {
	for _, action := range t.actions {
		action.Do(f, e)
	}
}

// Destination 目标状态 ID
func (t *Transition) Destination() StateID {
	return t.destination
}

// AddValidators 添加验证器
func (t *Transition) AddValidators(validators ...Validator) {
	t.validators = append(t.validators, validators...)
}

// AddActions 添加行为
func (t *Transition) AddActions(actions ...ActionInterface) {
	t.actions = append(t.actions, actions...)
}
