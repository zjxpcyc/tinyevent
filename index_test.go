package tinyevent_test

import (
	"testing"
	"time"

	"github.com/zjxpcyc/tinyevent"
)

var tbus = &tinyevent.DefaultBus{}
var gStr = "foo"

func on(name string) string {
	return tbus.On(name, func(dt interface{}) error {
		gStr = "bar"
		return nil
	})
}

func TestOn(t *testing.T) {
	id := on("click")

	if id != "click-0" {
		t.Fatalf("Test on function fail")
	}

	if gStr != "foo" {
		t.Fatalf("Test On function fail, unexpected calc tNum")
	}
}

func TestEmit(t *testing.T) {
	on("click")

	evt := tinyevent.Event{
		Name:    "click",
		Payload: nil,
	}

	tbus.Emit(evt)

	time.Sleep(1 * time.Second)

	if gStr != "bar" {
		t.Fatalf("Test Emit function fail")
	}
}

func TestOff(t *testing.T) {
	tbus.Off(on("click"))

	evt := tinyevent.Event{
		Name:    "click",
		Payload: nil,
	}
	tbus.Emit(evt)

	time.Sleep(1 * time.Second)

	if gStr == "bar" {
		t.Fatalf("Test Off function fail")
	}
}
