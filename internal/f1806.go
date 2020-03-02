package internal

func f1806(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 ^ s1i32
	s1i32 = 65536
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		f128(ctx, s0i32)
	}
}
