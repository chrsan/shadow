package internal

func f1662(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 0
		s2i32 = l0
		s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	}
}
