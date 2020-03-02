package internal

func f211(ctx *Context, l0 int64, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int64
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	s0i64 = l0
	s1i64 = 4294967296
	if uint64(s0i64) < uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i64 = l0
		l5 = s0i64
		goto lbl0
	}
lbl2:
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i64 = l0
	s2i64 = l0
	s3i64 = 10
	s2i64 = i64DivU(s2i64, s3i64)
	l5 = s2i64
	s3i64 = 10
	s2i64 = s2i64 * s3i64
	s1i64 = s1i64 - s2i64
	s1i32 = int32(s1i64)
	s2i32 = 48
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i64 = l0
	s1i64 = 42949672959
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l2 = s0i32
	s0i64 = l5
	l0 = s0i64
	s0i32 = l2
	if s0i32 != 0 {
		goto lbl2
	}
lbl0:
	s0i64 = l5
	s0i32 = int32(s0i64)
	l2 = s0i32
	if s0i32 != 0 {
	lbl4:
		s0i32 = l1
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l2
		s2i32 = l2
		s3i32 = 10
		s2i32 = i32DivU(s2i32, s3i32)
		l3 = s2i32
		s3i32 = 10
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 48
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l2
		s1i32 = 9
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l4 = s0i32
		s0i32 = l3
		l2 = s0i32
		s0i32 = l4
		if s0i32 != 0 {
			goto lbl4
		}
	}
	s0i32 = l1
	return s0i32
}
