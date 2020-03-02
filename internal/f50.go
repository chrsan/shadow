package internal

func f50(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int64
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	s0i32 = -1
	s1i32 = l0
	s1i64 = int64(uint32(s1i32))
	s2i32 = l1
	s2i64 = int64(uint32(s2i32))
	s1i64 = s1i64 * s2i64
	l2 = s1i64
	s1i32 = int32(s1i64)
	s2i64 = l2
	s3i64 = 32
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	s2i32 = int32(s2i64)
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = 2
	s0i32 = f44(ctx, s0i32, s1i32)
	return s0i32
}
