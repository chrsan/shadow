package internal

func f1346(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = int32(len(ctx.Mem) / 65536)
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 - s1i32
	s1i32 = 65535
	s0i32 = s0i32 + s1i32
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0i32 = int32(len(ctx.Mem) / 65536)
	panic("growMem not supported when using mmap yet")
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = 0
	f9(ctx, s0i32)
	s0i32 = 1
	return s0i32
}
