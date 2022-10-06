package control

import (
	"log"
	"testing"
)

// 测试代码是否正常运行
func TestRandStr(t *testing.T) {
	log.Println(RandStr(10))
}

// 测试代码运行速度
func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr(10)
	}
}
