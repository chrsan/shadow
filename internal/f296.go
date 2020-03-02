package internal

func f296(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 8
		s2i32 = 0
		f16(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+5)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s1i32 = 17
	s2i32 = 0
	f16(ctx, s0i32, s1i32, s2i32)
	goto lbl1
lbl2:
	s0i32 = l1
	s1i32 = l0
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	f660(ctx, s0i32, s1i32)
lbl1:
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 128
		s2i32 = l0
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		f16(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+3)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+6)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l1
	s1i32 = 18
	s2i32 = 0
	f16(ctx, s0i32, s1i32, s2i32)
	goto lbl4
lbl5:
	s0i32 = l1
	s1i32 = l0
	s2i32 = 36
	s1i32 = s1i32 + s2i32
	f660(ctx, s0i32, s1i32)
lbl4:
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+4)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 9
		s2i32 = 0
		f16(ctx, s0i32, s1i32, s2i32)
	}
}
