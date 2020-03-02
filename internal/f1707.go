package internal

func f1707(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l5 = s0i32
lbl0:
	s0i32 = l5
	s1i32 = l2
	l3 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l1
	s3i32 = 10
	s2i32 = i32DivU(s2i32, s3i32)
	l6 = s2i32
	s3i32 = 10
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 - s2i32
	s2i32 = 48
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = 9
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l6
	l1 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s2i32 = l4
		s1i32 = s1i32 + s2i32
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		s2i32 = 1
		s3i32 = l3
		s2i32 = s2i32 - s3i32
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	lbl2:
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l2
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l3 = s0i32
		s0i32 = l1
		l2 = s0i32
		s0i32 = l3
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
