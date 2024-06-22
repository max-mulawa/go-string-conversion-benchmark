package main

import (
	"bytes"
	"testing"
	"unsafe"
)

var str10 = "1234567890"
var byt10 = []byte(str10)

var str100 string
var byt100 []byte

var str1000 string
var byt1000 []byte

var str10000 string
var byt10000 []byte

func init() {
	buf := &bytes.Buffer{}

	for i := 0; i < 100; i++ {
		buf.WriteByte('a')
	}

	str100, byt100 = buf.String(), []byte(buf.String())

	buf.Reset()

	for i := 0; i < 1000; i++ {
		buf.WriteByte('a')
	}
	str1000, byt1000 = buf.String(), []byte(buf.String())

	buf.Reset()

	for i := 0; i < 10000; i++ {
		buf.WriteByte('a')
	}
	str10000, byt10000 = buf.String(), []byte(buf.String())

	buf.Reset()

}

func benchmarkStringByte(b *testing.B, byt []byte) {
	var str string //Some code to make sure our benchmark isn't optimized away
	for i := 0; i < b.N; i++ {
		str = string(byt)
	}
	_ = str
}

func benchmarkByteString(b *testing.B, str string) {
	var by []byte
	for i := 0; i < b.N; i++ {
		by = []byte(str)
	}
	_ = by
}

func benchmarkUnsafeByteString(b *testing.B, str string) {
	var by []byte
	for i := 0; i < b.N; i++ {
		by = unsafe.Slice(unsafe.StringData(str), len(str))
	}
	_ = by
}

func benchmarkUnsafeStringByte(b *testing.B, byt []byte) {
	var str string
	for i := 0; i < b.N; i++ {
		str = unsafe.String(unsafe.SliceData(byt), len(byt))
	}
	_ = str
}

func benchmarkCopyByteString(b *testing.B, str string) {
	var by []byte
	for i := 0; i < b.N; i++ {
		by = make([]byte, len(str))
		copy(by, str)
	}
	_ = by
}

func BenchmarkTestStringByteLen10(b *testing.B) {
	benchmarkStringByte(b, byt10)
}

func BenchmarkTestUnsafeStringByteLen10(b *testing.B) {
	benchmarkUnsafeStringByte(b, byt10)
}

func BenchmarkTestStringByteLen100(b *testing.B) {
	benchmarkStringByte(b, byt100)
}

func BenchmarkTestUnsafeStringByteLen100(b *testing.B) {
	benchmarkUnsafeStringByte(b, byt100)
}

func BenchmarkTestStringByteLen1000(b *testing.B) {
	benchmarkStringByte(b, byt1000)
}

func BenchmarkTestStringByteLen10000(b *testing.B) {
	benchmarkStringByte(b, byt10000)
}

func BenchmarkTestUnsafeStringByteLen10000(b *testing.B) {
	benchmarkUnsafeStringByte(b, byt10000)
}

func BenchmarkTestByteStringLen10(b *testing.B) {
	benchmarkByteString(b, str10)
}

func BenchmarkTestUnsafeByteStringLen10(b *testing.B) {
	benchmarkUnsafeByteString(b, str10)
}

func BenchmarkTestCopyByteStringLen10(b *testing.B) {
	benchmarkCopyByteString(b, str10)
}

func BenchmarkTestByteStringLen100(b *testing.B) {
	benchmarkByteString(b, str100)
}

func BenchmarkTestUnsafeByteStringLen100(b *testing.B) {
	benchmarkUnsafeByteString(b, str100)
}

func BenchmarkTestCopyByteStringLen100(b *testing.B) {
	benchmarkCopyByteString(b, str100)
}

func BenchmarkTestByteStringLen1000(b *testing.B) {
	benchmarkByteString(b, str1000)
}

func BenchmarkTestByteStringLen10000(b *testing.B) {
	benchmarkByteString(b, str10000)
}

func BenchmarkTestUnsafeByteStringLen10000(b *testing.B) {
	benchmarkUnsafeByteString(b, str10000)
}

func BenchmarkTestCopyByteStringLen10000(b *testing.B) {
	benchmarkCopyByteString(b, str10000)
}
