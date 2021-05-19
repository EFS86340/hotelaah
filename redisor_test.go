package hotelaah

import (
	"testing"
)

func TestSet(t *testing.T) {

	r := NewRedisor("localhost:6379", "", 0)
	err := r.Init()
	if err != nil {
		t.Fatal("Redis not online")
	}

	r.SetCity("wuhan", "hubei")

	_, err = r.GetCity("wuhan")
	if err != nil {
		t.Fatal("value not right")
	}

}
