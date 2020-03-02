package internal

import (
	"unsafe"
)

func f2088(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float32
	_ = l16
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l11 = s0f32
	s0i32 = l1
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 2
		l2 = s0i32
		s0f32 = l10
		l13 = s0f32
		s0f32 = l11
		l15 = s0f32
		s0i32 = 2
		l5 = s0i32
		goto lbl1
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = 2
	l5 = s0i32
	s0i32 = 2
	l2 = s0i32
	s0f32 = l11
	l12 = s0f32
	s0f32 = l10
	l14 = s0f32
lbl3:
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l15 = s0f32
	s1f32 = l12
	s0f32 = s0f32 - s1f32
	l12 = s0f32
	s0i32 = l1
	l0 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l13 = s0f32
	s1f32 = l14
	s0f32 = s0f32 - s1f32
	l14 = s0f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f32 = l12
	s3f32 = 0
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l14
		s1f32 = 0
		s0f32 = s0f32 * s1f32
		s1f32 = l12
		s0f32 = s0f32 * s1f32
		l16 = s0f32
		s1f32 = l16
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			return s0i32
		}
		s0i32 = 2
		l6 = s0i32
		s0i32 = l3
		s1i32 = l2
		s2f32 = l14
		s3f32 = 0
		if s2f32 < s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l1 = s2i32
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = 3
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l5
		s1f32 = l12
		s2f32 = 0
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		l8 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l9 = s0i32
		s0i32 = l1
		l2 = s0i32
		s0i32 = l8
		l5 = s0i32
		s0i32 = l4
		s1i32 = l9
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = 3
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0f32 = l15
	l12 = s0f32
	s0f32 = l13
	l14 = s0f32
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l7
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
lbl1:
	s0f32 = l10
	s1f32 = l13
	s0f32 = s0f32 - s1f32
	l10 = s0f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f32 = l11
	s3f32 = l15
	s2f32 = s2f32 - s3f32
	l11 = s2f32
	s3f32 = 0
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l10
		s1f32 = 0
		s0f32 = s0f32 * s1f32
		s1f32 = l11
		s0f32 = s0f32 * s1f32
		l13 = s0f32
		s1f32 = l13
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			return s0i32
		}
		s0i32 = 2
		l6 = s0i32
		s0i32 = l3
		s1i32 = l2
		s2f32 = l10
		s3f32 = 0
		if s2f32 < s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 + s1i32
		s1i32 = 3
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l4
		s1i32 = l5
		s2f32 = l11
		s3f32 = 0
		if s2f32 < s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 + s1i32
		s1i32 = 3
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = 1
	l6 = s0i32
lbl0:
	s0i32 = l6
	return s0i32
}
