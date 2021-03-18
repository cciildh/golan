package test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestAdd(t *testing.T) {
	ret := Add(4, 5)
	fmt.Println(ret)

	fmt.Println(runtime.NumCPU())
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < 5; i++ {
		Add(5, i)
	}

}
