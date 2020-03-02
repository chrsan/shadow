package internal

func f1165(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = 255
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	s1i32 = l0
	s0i32 = s0i32 * s1i32
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l0
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	return s0i32
}
