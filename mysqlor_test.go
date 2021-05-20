package hotelaah

import "testing"

func TestSetProvince(t *testing.T) {
	m := NewMysqlor("root:ohmysql@tcp(127.0.0.1)/mysql")
	m.Open()

	m.QueryTest()

}
