package main

import (
	"bytes"
	"testing"
)

var str10 = "1234567890"
var byt10 = []byte(str10)

var str100 string
var byt100 []byte

var str1000 string
var byt1000 []byte

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

func BenchmarkTestStringByteLen10(b *testing.B) {
	benchmarkStringByte(b, byt10)
}

func BenchmarkTestStringByteLen100(b *testing.B) {
	benchmarkStringByte(b, byt100)
}

func BenchmarkTestStringByteLen1000(b *testing.B) {
	benchmarkStringByte(b, byt1000)
}

func BenchmarkTestByteStringLen10(b *testing.B) {
	benchmarkByteString(b, str10)
}
func BenchmarkTestByteStringLen100(b *testing.B) {
	benchmarkByteString(b, str100)
}
func BenchmarkTestByteStringLen1000(b *testing.B) {
	benchmarkByteString(b, str1000)
}
