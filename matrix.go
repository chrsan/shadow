package shadow

import "unsafe"

type Matrix struct {
	ScaleX float32
	SkewX  float32
	TransX float32
	SkewY  float32
	ScaleY float32
	TransY float32
}

func IdentityMatrix() Matrix {
	var m Matrix
	m.Reset()
	return m
}

func (m *Matrix) Reset() {
	m.ScaleX = 1
	m.SkewX = 0
	m.TransX = 0
	m.SkewY = 0
	m.ScaleY = 1
	m.TransY = 0
}

func init() {
	if unsafe.Sizeof(Matrix{}) != 24 {
		panic("unexpected Matrix size")
	}
}
