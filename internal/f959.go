package internal

import (
	"unsafe"
)

func f959(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 float32
	_ = l2
	var l3 float32
	_ = l3
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
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
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l4 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s0f32 = s0f32 - s1f32
	l2 = s0f32
	s0f32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0f32 = s0f32 - s1f32
	l3 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l6 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0f32 = s0f32 - s1f32
	l5 = s0f32
	s0f32 = l6
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l10 = s1f32
	s0f32 = s0f32 - s1f32
	l8 = s0f32
	s0f32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 - s1f32
	l9 = s0f32
	s0f32 = l6
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	l7 = s1f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l2
		s1f32 = 0
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0f32 = l4
		s1f32 = l2
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		goto lbl3
	}
	s0f32 = l2
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l2
		s1f32 = l4
		s1f32 = -s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		goto lbl3
	}
	s0f32 = l4
	s1f32 = l2
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	goto lbl2
lbl4:
	s0f32 = l4
	s1f32 = l2
	s1f32 = -s1f32
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl3:
	s0i32 = 0
	s1f32 = l5
	s2f32 = l3
	s3f32 = l4
	s2f32 = s2f32 * s3f32
	s3f32 = l2
	s2f32 = s2f32 / s3f32
	s1f32 = s1f32 - s2f32
	l6 = s1f32
	s2f32 = l6
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = l8
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	s1f32 = l9
	s2f32 = l3
	s1f32 = s1f32 - s2f32
	s2f32 = l4
	s1f32 = s1f32 * s2f32
	s2f32 = l2
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l6
	s0f32 = s0f32 / s1f32
	goto lbl1
lbl2:
	s0i32 = 0
	s1f32 = l5
	s2f32 = l2
	s1f32 = s1f32 * s2f32
	s2f32 = l4
	s1f32 = s1f32 / s2f32
	s2f32 = l3
	s1f32 = s1f32 - s2f32
	l6 = s1f32
	s2f32 = l6
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = l3
	s1f32 = l8
	s2f32 = l5
	s1f32 = s1f32 - s2f32
	s2f32 = l2
	s1f32 = s1f32 * s2f32
	s2f32 = l4
	s1f32 = s1f32 / s2f32
	s2f32 = l9
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l6
	s0f32 = s0f32 / s1f32
lbl1:
	l6 = s0f32
	s0i32 = l1
	s1f32 = l7
	s2f32 = l7
	s3f32 = l5
	s4f32 = 0
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4i32 = 1
	s3i32 = s3i32 ^ s4i32
	if s3i32 == 0 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		s3f32 = l3
		s4f32 = 0
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		s4i32 = 1
		s3i32 = s3i32 ^ s4i32
		if s3i32 != 0 {
			goto lbl10
		}
		s3f32 = l5
		s4f32 = l3
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			goto lbl8
		}
		goto lbl9
	}
	s3f32 = l3
	s4f32 = 0
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4i32 = 1
	s3i32 = s3i32 ^ s4i32
	if s3i32 == 0 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		s3f32 = l3
		s4f32 = l5
		s4f32 = -s4f32
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			goto lbl8
		}
		goto lbl9
	}
	s3f32 = l5
	s4f32 = l3
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4i32 = 1
	s3i32 = s3i32 ^ s4i32
	if s3i32 != 0 {
		goto lbl9
	}
	goto lbl8
lbl10:
	s3f32 = l5
	s4f32 = l3
	s4f32 = -s4f32
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		goto lbl8
	}
lbl9:
	s3i32 = 0
	s4f32 = l5
	s5f32 = l2
	s4f32 = s4f32 * s5f32
	s5f32 = l3
	s4f32 = s4f32 / s5f32
	s5f32 = l4
	s4f32 = s4f32 - s5f32
	l7 = s4f32
	s5f32 = l7
	s4f32 = s4f32 * s5f32
	s5f32 = 0
	if s4f32 == s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		goto lbl0
	}
	s3f32 = l4
	s4f32 = l5
	s5f32 = l9
	s6f32 = l2
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = l3
	s4f32 = s4f32 / s5f32
	s5f32 = l8
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l7
	s3f32 = s3f32 / s4f32
	goto lbl7
lbl8:
	s3i32 = 0
	s4f32 = l2
	s5f32 = l3
	s6f32 = l4
	s5f32 = s5f32 * s6f32
	s6f32 = l5
	s5f32 = s5f32 / s6f32
	s4f32 = s4f32 - s5f32
	l7 = s4f32
	s5f32 = l7
	s4f32 = s4f32 * s5f32
	s5f32 = 0
	if s4f32 == s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		goto lbl0
	}
	s3f32 = l9
	s4f32 = l2
	s3f32 = s3f32 - s4f32
	s4f32 = l3
	s5f32 = l8
	s6f32 = l4
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = l5
	s4f32 = s4f32 / s5f32
	s3f32 = s3f32 - s4f32
	s4f32 = l7
	s3f32 = s3f32 / s4f32
lbl7:
	l2 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = l10
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l3 = s0f32
	s0i32 = l1
	s1f32 = l2
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l1
	s1f32 = l3
	s2f32 = l2
	s3f32 = l3
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = l5
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l1
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l2 = s1f32
	s2f32 = l6
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l2 = s0f32
	s0i32 = l1
	s1f32 = l6
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l1
	s1f32 = l2
	s2f32 = l6
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = l3
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s0i32 = l1
	s1i64 = 550821167104
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = 1
lbl0:
	return s0i32
}
