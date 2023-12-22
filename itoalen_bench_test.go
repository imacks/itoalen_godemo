package itoalentest

import (
	"fmt"
	"testing"
)

func BenchmarkItoalenWithItoa(b *testing.B) {
	samples := genRandInts(19)
	for i, v := range samples {
		b.Run(fmt.Sprintf("intlen_%d", i+1), func(b2 *testing.B) {
			for j := 0; j < b2.N; j++ {
				if n := ItoalenWithItoa(v); n != i+1 {
					b2.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
				}
			}
		})
	}
}

func BenchmarkItoalenWithLog(b *testing.B) {
	samples := genRandInts(19)
	for i, v := range samples {
		b.Run(fmt.Sprintf("intlen_%d", i+1), func(b2 *testing.B) {
			for j := 0; j < b2.N; j++ {
				if n := ItoalenWithLog(v); n != i+1 {
					b2.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
				}
			}
		})
	}
}

func BenchmarkItoalenWithMul(b *testing.B) {
	samples := genRandInts(19)
	for i, v := range samples {
		b.Run(fmt.Sprintf("intlen_%d", i+1), func(b2 *testing.B) {
			for j := 0; j < b2.N; j++ {
				if n := ItoalenWithMul(v); n != i+1 {
					b2.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
				}
			}
		})
	}
}

func BenchmarkItoalenWithMul2(b *testing.B) {
	samples := genRandInts(19)
	for i, v := range samples {
		b.Run(fmt.Sprintf("intlen_%d", i+1), func(b2 *testing.B) {
			for j := 0; j < b2.N; j++ {
				if n := ItoalenWithMul2(v); n != i+1 {
					b2.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
				}
			}
		})
	}
}

func BenchmarkItoalenWithDiv2(b *testing.B) {
	samples := genRandInts(19)
	for i, v := range samples {
		b.Run(fmt.Sprintf("intlen_%d", i+1), func(b2 *testing.B) {
			for j := 0; j < b2.N; j++ {
				if n := ItoalenWithDiv2(v); n != i+1 {
					b2.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
				}
			}
		})
	}
}

func BenchmarkItoalenWithBranch(b *testing.B) {
	samples := genRandInts(19)
	for i, v := range samples {
		b.Run(fmt.Sprintf("intlen_%d", i+1), func(b2 *testing.B) {
			for j := 0; j < b2.N; j++ {
				if n := ItoalenWithBranch(v); n != i+1 {
					b2.Fatalf("expect len(itoa(%d)) == %d but got %d", v, i+1, n)
				}
			}
		})
	}
}