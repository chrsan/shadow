package internal

func f1106(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	return s0i32
}
