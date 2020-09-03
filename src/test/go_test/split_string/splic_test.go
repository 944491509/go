package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("abcdbcdesbcdr", "b")
	want := []string{"a", "cd", "cdes", "cdr"}
	// 切换比较
	if !reflect.DeepEqual(got, want) {
		// 测试失败
		t.Errorf("want:%v but got:%v\n", want, got)
	}
}

func Test2Split(t *testing.T) {
	got := Split("a:b:c:d:e", ":")
	want := []string{"a", "b", "c", "d", "e"}
	if !reflect.DeepEqual(got, want) {
		// 测试失败
		t.Errorf("want:%v but got:%v\n", want, got)
	}
}
