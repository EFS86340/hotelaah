package hotelaah

import "testing"

func TestPublish(t *testing.T) {
	k := NewKafkee("test:1:1", "127.0.0.1:9092")
	sp := StringPair{
		First:  "wang",
		Second: "xu",
	}
	k.Init()
	k.Publish(&sp)
}
