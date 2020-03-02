package internal

func f87(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	s0i32 = ctx.g0
	s1i32 = 256
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = l3
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 73728
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s2i32 = s2i32 - s3i32
	l4 = s2i32
	s3i32 = 256
	s4i32 = l4
	s5i32 = 256
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l1 = s4i32
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = l5
	s2i32 = l1
	if s2i32 != 0 {
		s2i32 = l4
	} else {
		s2i32 = l2
		s3i32 = l3
		s2i32 = s2i32 - s3i32
		l1 = s2i32
	lbl2:
		s2i32 = l0
		s3i32 = l5
		s4i32 = 256
		f77(ctx, s2i32, s3i32, s4i32)
		s2i32 = l4
		s3i32 = -256
		s2i32 = s2i32 + s3i32
		l4 = s2i32
		s3i32 = 255
		if uint32(s2i32) > uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			goto lbl2
		}
		s2i32 = l1
		s3i32 = 255
		s2i32 = s2i32 & s3i32
	}
	f77(ctx, s0i32, s1i32, s2i32)
lbl0:
	s0i32 = l5
	s1i32 = 256
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
