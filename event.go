package tinyevent

// Action 动作
type Action func(Event) error

// Event 事件
type Event struct {
	Name    string
	Payload interface{}
}

// EventBus 中控
type EventBus interface {
	Emit(Event)
	On(string, Action) string
	Off(string)
}
