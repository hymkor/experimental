package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func stringsBuilder(b *testing.B, n int, grow bool) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		b.StopTimer()
		if grow {
			builder.Grow(n * 4)
		}
		b.StartTimer()
		for j := 0; j < n; j++ {
			fmt.Fprint(&builder, "TEST")
		}
		_ = builder.String()
	}
}

func bytesBuffer(b *testing.B, n int, grow bool) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		b.StopTimer()
		if grow {
			buf.Grow(n * 4)
		}
		b.StartTimer()
		for j := 0; j < n; j++ {
			fmt.Fprint(&buf, "TEST")
		}
		_ = buf.String()
	}
}

func bytesAppend(b *testing.B, n int, grow bool) {
	for i := 0; i < b.N; i++ {
		var buf []byte
		b.StopTimer()
		if grow {
			buf = make([]byte, 0, n*4)
		} else {
			buf = make([]byte, 0)
		}
		b.StartTimer()
		for j := 0; j < n; j++ {
			buf = append(buf, "TEST"...)
		}
		_ = string(buf)
	}
}

func BenchmarkStringsBuilderWithoutGrow(b *testing.B) {
	for _, n := range []int{10, 50, 100, 200, 500} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			stringsBuilder(b, n, false)
		})
	}
}

func BenchmarkStringsBuilderWithGrow(b *testing.B) {
	for _, n := range []int{10, 50, 100, 200, 500} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			stringsBuilder(b, n, true)
		})
	}
}

func BenchmarkBytesBufferWithoutGrow(b *testing.B) {
	for _, n := range []int{10, 50, 100, 200, 500} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			bytesBuffer(b, n, false)
		})
	}
}

func BenchmarkBytesBufferWithGrow(b *testing.B) {
	for _, n := range []int{10, 50, 100, 200, 500} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			bytesBuffer(b, n, true)
		})
	}
}

func BenchmarkBytesAppendWithoutGrow(b *testing.B) {
	for _, n := range []int{10, 50, 100, 200, 500} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			bytesAppend(b, n, false)
		})
	}
}

func BenchmarkBytesAppendWithGrow(b *testing.B) {
	for _, n := range []int{10, 50, 100, 200, 500} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			bytesAppend(b, n, true)
		})
	}
}
