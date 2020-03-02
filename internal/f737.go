package internal

func f737(ctx *Context) {
	var l0 int32
	_ = l0
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = 28784
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l0 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl1
	case 1:
		goto lbl0
	default:
		goto lbl2
	}
lbl2:
	s0i32 = 28784
	s1i32 = 28784
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l0 = s1i32
	s2i32 = 1
	s3i32 = l0
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 28784
	s1i32 = 2
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	return
lbl1:
lbl3:
	s0i32 = 28784
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
lbl0:
}
