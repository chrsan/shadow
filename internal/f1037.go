package internal

func f1037(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
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
	var s5i32 int32
	_ = s5i32
	s0i32 = l2
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l3 = s0i32
		s0i32 = 0
		l1 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		s1i32 = -1
		s2i32 = 0
		s3i32 = l3
		s4i32 = 7
		s5i32 = l1
		s4i32 = s4i32 - s5i32
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 1
		s3i32 = s3i32 & s4i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l2
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
