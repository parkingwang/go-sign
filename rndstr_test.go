package sign

import (
	"testing"
	"fmt"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestRandString(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(RandString(6))
	}
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(128)
	}
}