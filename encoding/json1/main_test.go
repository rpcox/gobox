package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ( 
	s string = `{"key1":"Value1","key2":4294967296,"key3":3.14,"key4":{"length":5,"width":6}}`
	_b []byte = []byte(s)
	js JsonStruct = JsonStruct{
		Key1: "Value1", 
		Key2: 1 << 32, 
		Key3: 3.14, 
		Key4: JsonSub1{
			Length: 5, 
			Width: 6,
		},
	}
)

// go test 
func TestMarshalStd(t *testing.T) {
	data, err := MarshalStd(&js)
	assert.NoError(t, err)
	assert.Equal(t, s, string(data))
}

func TestMarshalJsi(t *testing.T) {
	data, err := MarshalJsi(&js)
	assert.NoError(t, err)
	assert.Equal(t, s, string(data))
}

func TestUnMarshalStd(t *testing.T) {
	js2,err := UnMarshalStd(_b)
	assert.NoError(t, err)
	assert.Equal(t, js2, js)
}

// go test -bench=
func BenchmarkMarshalStd(b *testing.B) {
	for b.Loop() {
		_, _ = MarshalStd(&js)
	}
}

func BenchmarkMarshalJsi(b *testing.B) {
	for b.Loop() {
		_, _ = MarshalJsi(&js)
	}
}

func BenchmarkUnMarshalStd(b *testing.B) {
	for b.Loop() {
		_, _ = UnMarshalStd(_b)
	}
}

func BenchmarkUnMarshalJsi(b *testing.B) {
	for b.Loop() {
		_, _ = UnMarshalJsi(_b)
	}
}

