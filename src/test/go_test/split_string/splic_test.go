package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	// 测试分组
	testGroup := []testCase{
		{"abcdbcdesbcdr", "b", []string{"a", "cd", "cdes", "cdr"}},
		{"a:b:c:d:e", ":", []string{"a", "b", "c", "d", "e"}},
		{"abcedsbcasd", "bc", []string{"a", "eds", "asd"}},
		{"小洋爱小鑫", "爱", []string{"小洋", "小鑫"}},
	}

	testGroup1 := map[string]testCase{
		"test1": {"abcdbcdesbcdr", "b", []string{"a", "cd", "cdes", "cdr"}},
		"test2": {"a:b:c:d:e", ":", []string{"a", "b", "c", "d", "e"}},
		"test3": {"abcedsbcasd", "bc", []string{"a", "eds", "asd"}},
		"test4": {"小洋爱小鑫", "爱", []string{"小洋", "小鑫"}},
	}

	// go test -run=TestSplit/test4  单独跑一个测试的命令
	// go test cover -html=cover.out 生成文件
	// go tool cover -html=cover.out 查看文件
	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			// 测试失败
			t.Errorf("want:%v but got:%v\n", tc.want, got)
		}
	}

	for name, tc := range testGroup1 {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want:%v but got:%v\n", tc.want, got)
			}
		})
	}
}

// BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
