package internal

func f1705(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s1i32 = l2
	s2i32 = 8
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l1
	s3i32 = 16
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	s2i32 = l3
	s1i32 = s1i32 | s2i32
	s2i32 = l4
	s3i32 = 24
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	f533(ctx, s0i32, s1i32)
}
