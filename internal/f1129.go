package internal

import (
	"math"
	"unsafe"
)

func f1129(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l10 int32
	_ = l10
	var l11 int64
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
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
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
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
	s1i32 = 4640
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l4 = s0i32
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l3
	s2i32 = 120
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s0i32 = f346(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l4 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l12 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l13 = s0f32
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
		s0f32 = l13
		s1f32 = 0.5
		s0f32 = s0f32 + s1f32
		s0f32 = float32(math.Floor(float64(s0f32)))
		l13 = s0f32
		s1f32 = 2.1474835e+09
		s2f32 = l13
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
		l13 = s0f32
		s1f32 = -2.1474835e+09
		s2f32 = l13
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
		l13 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l13
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl2
		}
		s0i32 = -2147483648
	lbl2:
		l4 = s0i32
		s0f32 = l12
		s1f32 = 0.5
		s0f32 = s0f32 + s1f32
		s0f32 = float32(math.Floor(float64(s0f32)))
		l12 = s0f32
		s1f32 = 2.1474835e+09
		s2f32 = l12
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
		l12 = s0f32
		s1f32 = -2.1474835e+09
		s2f32 = l12
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
		l12 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l12
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl4
		}
		s0i32 = -2147483648
	lbl4:
		l5 = s0i32
		s0i32 = l1
		s1i32 = l3
		s2i32 = 120
		s1i32 = s1i32 + s2i32
		s0i32 = f156(ctx, s0i32, s1i32)
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l4
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l3
			s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+144)])))
			s2i32 = l4
			s2i64 = int64(s2i32)
			s1i64 = s1i64 + s2i64
			l11 = s1i64
			s2i64 = 2147483647
			s3i64 = l11
			s4i64 = 2147483647
			if s3i64 < s4i64 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1i64 = s1i64
			} else {
				s1i64 = s2i64
			}
			l11 = s1i64
			s2i64 = -2147483647
			s3i64 = l11
			s4i64 = -2147483647
			if s3i64 > s4i64 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1i64 = s1i64
			} else {
				s1i64 = s2i64
			}
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i64)
			s0i32 = l3
			s1i32 = l5
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l3
			s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)])))
			s2i32 = l5
			s2i64 = int64(s2i32)
			s1i64 = s1i64 + s2i64
			l11 = s1i64
			s2i64 = 2147483647
			s3i64 = l11
			s4i64 = 2147483647
			if s3i64 < s4i64 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1i64 = s1i64
			} else {
				s1i64 = s2i64
			}
			l11 = s1i64
			s2i64 = -2147483647
			s3i64 = l11
			s4i64 = -2147483647
			if s3i64 > s4i64 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1i64 = s1i64
			} else {
				s1i64 = s2i64
			}
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i64)
			s0i32 = l3
			s1i32 = 1
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l3
			s2i32 = 72
			s1i32 = s1i32 + s2i32
			s2i32 = l2
			f504(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l1 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l0
		f12(ctx, s0i32)
	lbl7:
		s0i32 = l3
		s1i32 = 4640
		s0i32 = s0i32 + s1i32
		ctx.g0 = s0i32
		return
	}
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4592)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l4 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l5 = s0i32
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4624)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l5
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4636)])) = s1f32
	s0i32 = l3
	s1i32 = l4
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4632)])) = s1f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	s1i32 = l3
	s2i32 = 4624
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s3i32 = 4624
	s2i32 = s2i32 + s3i32
	s0i32 = f42(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s1i32 = 4608
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4636)]))
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l12 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l12
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
	l12 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l12
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
	l12 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l12
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl8
	}
	s1i32 = -2147483648
lbl8:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 4604
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4632)]))
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l12 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l12
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
	l12 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l12
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
	l12 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l12
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl10
	}
	s1i32 = -2147483648
lbl10:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 4600
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4628)]))
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l12 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l12
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
	l12 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l12
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
	l12 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l12
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl12
	}
	s1i32 = -2147483648
lbl12:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4624)]))
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l12 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l12
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
	l12 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l12
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
	l12 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l12
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl14
	}
	s1i32 = -2147483648
lbl14:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4596)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	l4 = s1i32
	s2i32 = l4
	s3i32 = 20
	s2i32 = s2i32 + s3i32
	s3i32 = l4
	s3i32 = int32(ctx.Mem[int(s3i32+40)])
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l4 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l5 = s1i32
	s2i32 = -2147483646
	s3i32 = l5
	s4i32 = -2147483646
	if s3i32 > s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l5 = s1i32
	s2i32 = -2147483646
	s3i32 = l5
	s4i32 = -2147483646
	if s3i32 > s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = -2147483647
	s2i32 = l4
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)])))
	l11 = s2i64
	s3i64 = 2147483646
	s4i64 = l11
	s5i64 = 2147483646
	if s4i64 < s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	l11 = s2i64
	s2i32 = int32(s2i64)
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	s3i64 = l11
	s4i64 = -2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = -2147483647
	s2i32 = l4
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)])))
	l11 = s2i64
	s3i64 = 2147483646
	s4i64 = l11
	s5i64 = 2147483646
	if s4i64 < s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	l11 = s2i64
	s2i32 = int32(s2i64)
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	s3i64 = l11
	s4i64 = -2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 120
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = 120
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s3i32 = 72
	s2i32 = s2i32 + s3i32
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s1i32 = 4592
	s0i32 = s0i32 + s1i32
	s1i32 = 4
	s0i32 = s0i32 | s1i32
	l4 = s0i32
	s1i32 = l4
	s2i32 = l3
	s3i32 = 120
	s2i32 = s2i32 + s3i32
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4616)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4604)]))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4596)]))
	s1i32 = s1i32 - s2i32
	s2i32 = 3
	s1i32 = s1i32 + s2i32
	s2i32 = -4
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4612)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 4592
	s0i32 = s0i32 + s1i32
	s0i32 = f118(ctx, s0i32)
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l4
	s2i32 = 1
	s1i32 = f50(ctx, s1i32, s2i32)
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4592)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	s2i32 = l4
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l9 = s0i32
	s0i32 = l3
	s1i32 = 4552
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 0
	ctx.Mem[int(s0i32+32)] = uint8(s1i32)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l4
	l6 = s0i32
	s0i32 = l3
	s1i64 = 8589934593
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4604)]))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4596)]))
	s1i32 = s1i32 - s2i32
	s1i64 = int64(uint32(s1i32))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4608)]))
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4600)]))
	s2i32 = s2i32 - s3i32
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l3
	s2i32 = 120
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4592)]))
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4612)]))
	s4i32 = 0
	s5i32 = 0
	s0i32 = f157(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l5 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = l4
	f12(ctx, s0i32)
lbl16:
	s0i32 = l3
	s1i32 = 120
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s0i32 = f406(ctx, s0i32, s1i32)
	l7 = s0i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4596)]))
	s1f32 = float32(s1i32)
	s1f32 = -s1f32
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4600)]))
	s2f32 = float32(s2i32)
	s2f32 = -s2f32
	f169(ctx, s0i32, s1f32, s2f32)
	s0i32 = l7
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	f190(ctx, s0i32, s1i32)
	s0i32 = l3
	s1i32 = 72
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 13195221663744
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	l5 = s1i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
	s3i32 = -4
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = -769
	s1i32 = s1i32 & s2i32
	s2i32 = l5
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 3
	s2i32 = s2i32 & s3i32
	s3i32 = 8
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 768
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s0i32 = f46(ctx, s0i32, s1i32)
	l8 = s0i32
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = l1
	s3i32 = 0
	f300(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l8
	s1i32 = l3
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	f83(ctx, s0i32, s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l10 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l10
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
lbl17:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l5 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l1 = s0i32
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l1
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l3
	s1i32 = l5
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l7
	s1i32 = l3
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l8
	f385(ctx, s0i32, s1i32, s2i32)
	s0i32 = l8
	s0i32 = f23(ctx, s0i32)
	s0i32 = l4
	s0i32 = f23(ctx, s0i32)
	s0i32 = l7
	s0i32 = f258(ctx, s0i32)
	s0i32 = l6
	s0i32 = f41(ctx, s0i32)
	s0i32 = l0
	s1i32 = l3
	s2i32 = 4592
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	f504(ctx, s0i32, s1i32, s2i32)
	s0i32 = l9
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l9
	f13(ctx, s0i32)
lbl0:
	s0i32 = l3
	s1i32 = 4640
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
