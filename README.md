# TinyEvent
极简的事件系统

![](https://img.shields.io/badge/golang-v0.0.3-blue.svg)
![](https://img.shields.io/github/license/zjxpcyc/tinyevent.svg)

## Install
```bash
// x.y.z 是版本号
go get github.com/zjxpcyc/tinyevent@vx.y.z
```

## Useage

```golang

// 初始化一个中控中心
// 或者自己实现一个 满足接口 EventBus 即可
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
  func(data interface{}) error {

    // 下面是事件动作的具体内容
    someStr, _ := data.(string)
    fmt.Println(someStr)
    
    return nil
  }
)

// 触发事件执行 - 此处是 go 线程执行
evt := tinyevent.Event {
  Name: EvtStart,
  Payload: "这个字符串将会被打印-在本示例中",
}
evtBus.Emit(evt)

// 取消事件注册
evtBus.Off(evtID)

// 更多信息请查阅 event.go 文件

```
