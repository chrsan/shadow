package internal

import (
	"math"
	"unsafe"
)

func f219(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
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
	s0i32 = ctx.g0
	s1i32 = 1136
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l7 = s1f32
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l8 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l9 = s1f32
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+40)])
	if s0i32 != 0 {
		s0i32 = l3
		s1f32 = l9
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l9 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l9
		s4f32 = 2.1474835e+09
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l9 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l9
		s4f32 = -2.1474835e+09
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l9 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l9
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl2
		}
		s1i32 = -2147483648
	lbl2:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l3
		s1f32 = l7
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l7 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l7
		s4f32 = 2.1474835e+09
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l7 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l7
		s4f32 = -2.1474835e+09
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l7 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l7
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl4
		}
		s1i32 = -2147483648
	lbl4:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l3
		s1f32 = l8
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l7 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l7
		s4f32 = 2.1474835e+09
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l7 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l7
		s4f32 = -2.1474835e+09
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l7 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l7
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl6
		}
		s1i32 = -2147483648
	lbl6:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l3
		s1f32 = l6
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l6 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l6
		s4f32 = 2.1474835e+09
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l6 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l6
		s4f32 = -2.1474835e+09
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l6 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l6
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl8
		}
		s1i32 = -2147483648
	lbl8:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l2
		f220(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l2
	s0i32 = f151(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1104)]))
	l4 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1100)]))
	l5 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	l6 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l6
	s3f32 = 2.1474835e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l6 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l6
	s3f32 = -2.1474835e+09
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l6 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l6
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl10
	}
	s0i32 = -2147483648
lbl10:
	l2 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0f32
	s0i32 = l3
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1132)])) = uint32(s1i32)
	s0i32 = l3
	s1f32 = l8
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l8 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l8
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l8 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l8
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l8 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l8
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl12
	}
	s1i32 = -2147483648
lbl12:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1128)])) = uint32(s1i32)
	s0i32 = l3
	s1f32 = l7
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l7 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l7
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l7 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l7
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l7 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l7
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl14
	}
	s1i32 = -2147483648
lbl14:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1124)])) = uint32(s1i32)
	s0i32 = l3
	s1f32 = l6
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l6 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l6
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l6
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l6
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl16
	}
	s1i32 = -2147483648
lbl16:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1120)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 1120
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = l4
	f220(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = f74(ctx, s0i32)
	s0i32 = l1
	f40(ctx, s0i32)
lbl0:
	s0i32 = l3
	s1i32 = 1136
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
