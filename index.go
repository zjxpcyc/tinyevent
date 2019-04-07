package tinyevent

import (
	"strconv"
	"strings"
)

// Actions 动作集合
type Actions = []Action

// DefaultBus 默认事件BUS
type DefaultBus struct {
	actMap map[EventName]Actions
}

// Emit 执行事件
func (t *DefaultBus) Emit(evt Event) {
	if t.actMap == nil {
		return
	}

	for evtName, acts := range t.actMap {
		if acts == nil || len(acts) == 0 {
			continue
		}

		if evtName == evt.Name {
			for _, act := range acts {
				t.execAction(act, evt)
			}
		}
	}
}

// On 监听
func (t *DefaultBus) On(evtName EventName, act Action) EventID {
	if t.actMap == nil {
		t.actMap = make(map[EventName]Actions)
	}

	actions, ok := t.actMap[evtName]
	if !ok {
		actions = make(Actions, 0)
	}

	actions = append(actions, act)
	t.actMap[evtName] = actions

	pos := len(actions) - 1
	return evtName + "-" + strconv.Itoa(pos)
}

// Off 取消监听
func (t *DefaultBus) Off(id EventID) {
	ids := strings.Split(id, "-")
	if len(ids) != 2 {
		return
	}

	evtName := ids[0]
	actID, err := strconv.Atoi(ids[1])
	if err != nil {
		return
	}

	actions, ok := t.actMap[evtName]
	if !ok {
		return
	}

	actions[actID] = nil
	t.actMap[evtName] = actions
}

func (t *DefaultBus) execAction(exec Action, evt Event) {
	if exec == nil {
		return
	}

	go exec(evt.Payload)
}

var bus EventBus = &DefaultBus{}
