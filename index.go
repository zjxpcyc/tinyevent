package tinyevent

import (
	"strconv"
	"strings"
)

// DefaultBus 默认事件BUS
type DefaultBus struct {
	actions map[string][]Action
}

// Emit 执行事件
func (t *DefaultBus) Emit(evt Event) {
	if t.actions == nil {
		return
	}

	for evtName, acts := range t.actions {
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
func (t *DefaultBus) On(evtName string, act Action) string {
	if t.actions == nil {
		t.actions = make(map[string][]Action)
	}

	actions, ok := t.actions[evtName]
	if !ok {
		actions = make([]Action, 0)
	}

	actions = append(actions, act)
	t.actions[evtName] = actions

	pos := len(actions) - 1
	return evtName + "-" + strconv.Itoa(pos)
}

// Off 取消监听
func (t *DefaultBus) Off(id string) {
	ids := strings.Split(id, "-")
	if len(ids) != 2 {
		return
	}

	evtName := ids[0]
	actID, err := strconv.Atoi(ids[1])
	if err != nil {
		return
	}

	actions, ok := t.actions[evtName]
	if !ok {
		return
	}

	actions[actID] = nil
	t.actions[evtName] = actions
}

func (t *DefaultBus) execAction(act Action, evt Event) {
	if act == nil {
		return
	}

	go act(evt)
}

var bus EventBus = &DefaultBus{}
