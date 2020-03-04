package internal

func f435(ctx *Context, l0 int32, l1 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+40)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	f143(ctx, s0i32)
	s0i32 = l0
	s1i32 = l0
	s2i32 = l1
	s1i32 = f81(ctx, s1i32, s2i32)
	l1 = s1i32
	ctx.Mem[int(s0i32+42)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	ctx.Mem[int(s0i32+41)] = uint8(s1i32)
}
