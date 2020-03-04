package internal

func f44(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = f1352(ctx, s0i32)
		goto lbl0
	}
	s0i32 = l0
	s0i32 = f164(ctx, s0i32)
lbl0:
	l2 = s0i32
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l2
	if s0i32 != 0 {
		goto lbl2
	}
	f104(ctx)
	panic("unreachable executed")
lbl2:
	s0i32 = l2
	return s0i32
}
