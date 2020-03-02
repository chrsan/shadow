package internal

func f146(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
lbl0:
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l1 = s0i32
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl0
	}
}
