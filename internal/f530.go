package internal

func f530(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	s0i32 = l0
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	s0i32 = l0
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l3 = s0i32
	s0i32 = l0
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	s0i32 = l0
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l2 = s0i32
	s1i32 = 255
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l1
		s2i32 = 16
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 | s1i32
		s1i32 = l3
		s2i32 = 8
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 | s1i32
		s1i32 = -16777216
		s0i32 = s0i32 | s1i32
		return s0i32
	}
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 * s1i32
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 67043328
	s0i32 = s0i32 & s1i32
	s1i32 = l2
	s2i32 = l3
	s1i32 = s1i32 * s2i32
	s2i32 = 128
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = l1
	s1i32 = s1i32 + s2i32
	s2i32 = 261888
	s1i32 = s1i32 & s2i32
	s2i32 = l0
	s3i32 = -16777216
	s2i32 = s2i32 & s3i32
	s3i32 = l2
	s4i32 = l4
	s3i32 = s3i32 * s4i32
	s4i32 = 128
	s3i32 = s3i32 + s4i32
	l0 = s3i32
	s4i32 = 8
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = l0
	s3i32 = s3i32 + s4i32
	s4i32 = 8
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 | s3i32
	s1i32 = s1i32 | s2i32
	s0i32 = s0i32 | s1i32
	return s0i32
}
