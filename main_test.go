package main

import "testing"

func BenchmarkWriteInt(b *testing.B) {
	writeInt(500)
}

func BenchmarkReadInt(b *testing.B) {
	readInt(500)
}

func BenchmarkWriteStr(b *testing.B) {
	writeStr(500)
}
func BenchmarkReadStr(b *testing.B) {
	readStr(500)
}
