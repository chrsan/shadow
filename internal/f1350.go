package internal

func f1350(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int64
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s0i64 int64
	_ = s0i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	s0i32 = 0
	s1i32 = l0
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0i64 = int64(uint32(s0i32))
	l2 = s0i64
	s0i32 = int32(s0i64)
	l1 = s0i32
	s1i32 = l0
	s2i32 = 1
	s1i32 = s1i32 | s2i32
	s2i32 = 65536
	if uint32(s1i32) < uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s0i32 = -1
	s1i32 = l1
	s2i64 = l2
	s3i64 = 32
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	s2i32 = int32(s2i64)
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
lbl1:
	l1 = s0i32
	s0i32 = f164(ctx, s0i32)
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = 0
	s2i32 = l1
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
lbl0:
	s0i32 = l0
	return s0i32
}
