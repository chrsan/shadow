package internal

import (
	"math"
	"unsafe"
)

func f1701(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 float32
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
	var l17 float32
	_ = l17
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l13 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l14 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l15 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l16 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l12 = s1f32
	s0f32 = s0f32 - s1f32
	l11 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	l8 = s0i32
	s0i32 = 1
	l6 = s0i32
	s0i32 = 1
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l17 = s1f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l18 = s2f32
	s1f32 = s1f32 - s2f32
	l9 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = 2139095040
	s1i32 = s1i32 & s2i32
	s2i32 = 2139095040
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0i32 = 1
	s1f32 = l14
	s2f32 = l15
	s1f32 = s1f32 - s2f32
	l19 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = 2139095040
	s1i32 = s1i32 & s2i32
	s2i32 = 2139095040
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = l9
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l19
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
lbl0:
	l7 = s0i32
	s0f32 = l13
	s1f32 = l10
	s0f32 = s0f32 - s1f32
	l9 = s0f32
	s0i32 = l8
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l9
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l11
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l9
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	l6 = s0i32
lbl1:
	s0i32 = l6
	s1i32 = l7
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = 0
	l1 = s0i32
	s0i32 = l7
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = 1
	l1 = s0i32
	s0f32 = l12
	s1f32 = l18
	s0f32 = s0f32 - s1f32
	l12 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f32 = l10
	s1f32 = l15
	s0f32 = s0f32 - s1f32
	l10 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f32 = l12
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l10
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	l1 = s0i32
lbl4:
	s0i32 = l6
	if s0i32 != 0 {
		s0f32 = l16
		s1f32 = l17
		s0f32 = s0f32 - s1f32
		l11 = s0f32
		s0i32 = int32(math.Float32bits(s0f32))
		s1i32 = 2139095040
		s0i32 = s0i32 & s1i32
		s1i32 = 2139095040
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0f32 = l13
		s1f32 = l14
		s0f32 = s0f32 - s1f32
		l9 = s0f32
		s0i32 = int32(math.Float32bits(s0f32))
		s1i32 = 2139095040
		s0i32 = s0i32 & s1i32
		s1i32 = 2139095040
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0f32 = l11
		s1f32 = 0
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1f32 = l9
		s2f32 = 0
		if s1f32 == s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 & s1i32
	} else {
		s0i32 = 0
	}
	s1i32 = l1
	s0i32 = s0i32 | s1i32
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl3:
	s0i32 = l4
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	return
lbl2:
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0f32
	s0i32 = l5
	s1f32 = l11
	s2f32 = l9
	s0i32 = f327(ctx, s0i32, s1f32, s2f32)
	if s0i32 != 0 {
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0f32
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s1f32 = l9
		s1f32 = -s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0f32
		s0i32 = l4
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2f32 = l10
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l4
		s1f32 = l9
		s2f32 = l10
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	}
}
