package internal

func f1371(ctx *Context, l0 int64, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i64 = l0
	if s0i64 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl1:
		s0i32 = l1
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i64 = l0
		s1i32 = int32(s1i64)
		s2i32 = 7
		s1i32 = s1i32 & s2i32
		s2i32 = 48
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i64 = l0
		s1i64 = 3
		s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
		l0 = s0i64
		s1i64 = 0
		if s0i64 != s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l1
	return s0i32
}
