# TinyEvent
极简的事件系统


## Install
```bash
go get github.com/zjxpcyc/tinyevent
```

## Useage

```golang

// 初始化一个中控中心
var evtBus := &tinyevent.DefaultBus{}

// 声明一个事件
const (
  EvtStart = "event_start"
)

// 注册事件
evtID := evtBus.On(
  // 事件的名称
  EvtStart,
  // 事件需要执行的动作
  func(evt tinyevent.Event) error {
    // TODO
    return nil
  }
)

// 触发事件执行 - 此处是 go 线程执行
evt := tinyevent.Event {
  Name: EvtStart,
  Payload: []byte(`这里是需要传入的数据`),
}
evtBus.Emit(evt)

// 取消事件注册
evtBus.Off(evtID)

// 更多信息请查阅 event.go 文件

```
