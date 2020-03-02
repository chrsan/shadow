package internal

func f1164(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l1
	s1i32 = s1i32 * s2i32
	s2i32 = 128
	s1i32 = s1i32 + s2i32
	l0 = s1i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = l0
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s0i32 = s0i32 - s1i32
	return s0i32
}
