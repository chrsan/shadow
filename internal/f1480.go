package internal

import (
	"math"
	"unsafe"
)

func f1480(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int64
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
	var s5i32 int32
	_ = s5i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l6
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = 1
		s3i32 = l1
		s4i32 = l2
		s5i32 = l0
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+296)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l1 = s0i32
		s1i32 = 16383
		s0i32 = s0i32 & s1i32
		l5 = s0i32
		s0i32 = l1
		s1i32 = 18
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		l2 = s0i32
		s0i32 = l1
		s1i32 = 10
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 240
		s0i32 = s0i32 & s1i32
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
	l5 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s1i32 = f24(ctx, s1i32)
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	}
	s0i32 = l6
	s1i32 = l5
	s2i32 = 14
	s1i32 = s1i32 & s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+284)]))
		s2i32 = l2
		s1i32 = s1i32 + s2i32
		goto lbl3
	}
	s1i32 = l0
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s2f32 = float32(s2i32)
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	s3i32 = l2
	s3f32 = float32(s3i32)
	s4f32 = 0.5
	s3f32 = s3f32 + s4f32
	s4i32 = l6
	s5i32 = 8
	s4i32 = s4i32 + s5i32
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+260)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 4 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, float32, float32, int32))(table[s5i32].f()))(ctx, s1i32, s2f32, s3f32, s4i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		s2f32 = 0
		if s1f32 > s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		goto lbl5
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+284)]))
	s2i32 = 1
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
lbl5:
	l2 = s1i32
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2f32 = 4.2949673e+09
	s1f32 = s1f32 * s2f32
	l10 = s1f32
	s2f32 = 9.2233715e+18
	s3f32 = l10
	s4f32 = 9.2233715e+18
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
	l10 = s1f32
	s2f32 = -9.2233715e+18
	s3f32 = l10
	s4f32 = -9.2233715e+18
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
	l10 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 9.223372e+18
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l10
		s1i64 = int64(math.Trunc(float64(s1f32)))
		goto lbl7
	}
	s1i64 = -9223372036854775808
lbl7:
	s2i32 = l2
	s2i64 = int64(s2i32)
	s3i64 = 16
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 - s2i64
	l9 = s1i64
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl9
		}
	}
	s1i64 = l9
	s2i32 = l0
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)])))
	s1i64 = s1i64 * s2i64
	l9 = s1i64
lbl9:
	s1i64 = l9
	s2i64 = 32
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
lbl3:
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	l5 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = l1
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			l1 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			s1i32 = l6
			s2i32 = 8
			s1i32 = s1i32 + s2i32
			s2i32 = l6
			s3i32 = l2
			s4i32 = l1
			if s3i32 < s4i32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			s2i32 = l2
			s3i32 = 0
			if s2i32 < s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l2 = s0i32
			goto lbl11
		}
		s0i32 = l2
		s1i32 = l1
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
		}
		s0i32 = l2
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l2
			s2i32 = -1
			s1i32 = s1i32 ^ s2i32
			s2i32 = l1
			s1i32 = i32RemS(s1i32, s2i32)
			s2i32 = -1
			s1i32 = s1i32 ^ s2i32
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			goto lbl11
		}
		s0i32 = l2
		s1i32 = l1
		s0i32 = i32RemS(s0i32, s1i32)
		l2 = s0i32
		goto lbl11
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l5 = s0i32
	s1i32 = l2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l2
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l2
		s2i32 = -1
		s1i32 = s1i32 ^ s2i32
		s2i32 = l5
		s1i32 = i32RemS(s1i32, s2i32)
		s2i32 = -1
		s1i32 = s1i32 ^ s2i32
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		goto lbl15
	}
	s0i32 = l2
	s1i32 = l5
	s0i32 = i32RemS(s0i32, s1i32)
	l2 = s0i32
lbl15:
	s0i32 = l2
	s1i32 = l5
	s2i32 = l2
	s3i32 = -1
	s2i32 = s2i32 ^ s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = l1
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l2 = s0i32
lbl11:
	s0i32 = 0
	l5 = s0i32
	s0i32 = 0
lbl0:
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l8 = s1i32
	s2i32 = l2
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
	if s1i32 != 0 {
		s1i32 = l0
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+288)])))
		l0 = s1i32
		s2i32 = l7
		s3i32 = l5
		s4i32 = l8
		s3i32 = s3i32 * s4i32
		s2i32 = s2i32 + s3i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l5 = s2i32
		s3i32 = 16711935
		s2i32 = s2i32 & s3i32
		s3i32 = l1
		s2i32 = s2i32 * s3i32
		s3i32 = 256
		s4i32 = l1
		s3i32 = s3i32 - s4i32
		l7 = s3i32
		s4i32 = l2
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		l2 = s4i32
		s5i32 = 16711935
		s4i32 = s4i32 & s5i32
		s3i32 = s3i32 * s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 16711935
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 * s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 16711935
		s1i32 = s1i32 & s2i32
		s2i32 = l5
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 16711935
		s2i32 = s2i32 & s3i32
		s3i32 = l1
		s2i32 = s2i32 * s3i32
		s3i32 = l2
		s4i32 = 8
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 16711935
		s3i32 = s3i32 & s4i32
		s4i32 = l7
		s3i32 = s3i32 * s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 16711935
		s2i32 = s2i32 & s3i32
		s3i32 = l0
		s2i32 = s2i32 * s3i32
		s3i32 = -16711936
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		goto lbl17
	}
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l2 = s1i32
	s2i32 = l0
	s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+288)])))
	l0 = s2i32
	s3i32 = 255
	if uint32(s2i32) > uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl17
	}
	s1i32 = l2
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l0
	s1i32 = s1i32 * s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l2
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s3i32 = l0
	s2i32 = s2i32 * s3i32
	s3i32 = -16711936
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
lbl17:
	s2i32 = l4
	s3i32 = 22124
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32))(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	s0i32 = l6
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
