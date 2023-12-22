package itoalentest

import (
	"math"
	"math/rand"
	"testing"
)

func TestItoalen(t *testing.T) {
	samples := genRandInts(19)
	for i, v := range samples {
		if n := ItoalenWithItoa(v); n != i+1 {
			t.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
		}
		if n := ItoalenWithLog(v); n != i+1 {
			t.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
		}
		if n := ItoalenWithMul(v); n != i+1 {
			t.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
		}
		if n := ItoalenWithMul2(v); n != i+1 {
			t.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
		}
		if n := ItoalenWithDiv2(v); n != i+1 {
			t.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
		}
		if n := ItoalenWithBranch(v); n != i+1 {
			t.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
		}
	}
}

// helpers

func genRandInts(maxlen int) []int {
	nums := make([]int, maxlen)
	for i := 1; i < maxlen; i++ {
		if i == 0 {
			nums[i] = rand.Intn(10)
			continue
		}
		v := int(math.Pow10(i)) // 10, 100, ...

		if i == 18 {
			nums[i] = rand.Intn(math.MaxInt - v) + v
			continue
		}
		nums[i] = rand.Intn(v*10 - v) + v // 10-99, 100-999, etc...
	}
	return nums
}
