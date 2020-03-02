package internal

func f205(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = 12
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		s1i32 = l0
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l0 = s0i32
		s1i32 = 2824
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = 1
		s1i32 = l0
		s2i32 = 4116
		s1i32 = s1i32 & s2i32
		if s1i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = 0
lbl1:
	return s0i32
lbl0:
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	return s0i32
}
