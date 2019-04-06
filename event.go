package tinyevent

// Action 动作
type Action func(Event) error

// EventName 事件名称
type EventName = string

// EventID 事件ID
type EventID = string

// Event 事件
type Event struct {
	Name    EventName
	Payload interface{}
}

// EventBus 中控
type EventBus interface {
	Emit(Event)
	On(EventName, Action) EventID
	Off(EventID)
}
