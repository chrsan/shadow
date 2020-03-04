package internal

import (
	"math/bits"
	"unsafe"
)

func f1350(ctx *Context, l0 int32, l1 int32) int32 {
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
	var l10 int32
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l2 = s0i32
	s0i32 = l0
	s1i32 = l7
	s2i32 = -8
	s1i32 = s1i32 & s2i32
	l6 = s1i32
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = 30240
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s1i32 = l0
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl0:
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l2 = s0i32
		s0i32 = l1
		s1i32 = 256
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l6
		s1i32 = l1
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			l2 = s0i32
			s0i32 = l6
			s1i32 = l1
			s0i32 = s0i32 - s1i32
			s1i32 = 30704
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = 1
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			if uint32(s0i32) <= uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl1
			}
		}
		s0i32 = 0
		return s0i32
	}
	s0i32 = l6
	s1i32 = l1
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = l1
		s0i32 = s0i32 - s1i32
		l2 = s0i32
		s1i32 = 16
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l0
		s1i32 = l7
		s2i32 = 1
		s1i32 = s1i32 & s2i32
		s2i32 = l1
		s1i32 = s1i32 | s2i32
		s2i32 = 2
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l2
		s2i32 = 3
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = 1
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		f535(ctx, s0i32, s1i32)
		goto lbl4
	}
	s0i32 = 0
	l2 = s0i32
	s0i32 = l4
	s1i32 = 30248
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 30236
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l1
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l0
		s1i32 = l7
		s2i32 = 1
		s1i32 = s1i32 & s2i32
		s2i32 = l1
		s1i32 = s1i32 | s2i32
		s2i32 = 2
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l5
		s2i32 = l1
		s1i32 = s1i32 - s2i32
		l1 = s1i32
		s2i32 = 1
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = 30236
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = 30248
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl4
	}
	s0i32 = l4
	s1i32 = 30244
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 30232
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l1
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l5
		s1i32 = l1
		s0i32 = s0i32 - s1i32
		l2 = s0i32
		s1i32 = 16
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l7
			s2i32 = 1
			s1i32 = s1i32 & s2i32
			s2i32 = l1
			s1i32 = s1i32 | s2i32
			s2i32 = 2
			s1i32 = s1i32 | s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l1
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s1i32 = l2
			s2i32 = 1
			s1i32 = s1i32 | s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l5
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l5
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			s2i32 = -2
			s1i32 = s1i32 & s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			goto lbl8
		}
		s0i32 = l0
		s1i32 = l7
		s2i32 = 1
		s1i32 = s1i32 & s2i32
		s2i32 = l5
		s1i32 = s1i32 | s2i32
		s2i32 = 2
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = 1
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = 0
		l2 = s0i32
		s0i32 = 0
		l1 = s0i32
	lbl8:
		s0i32 = 30244
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = 30232
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl4
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0i32
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = -8
	s0i32 = s0i32 & s1i32
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s1i32 = l1
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l9
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l10 = s0i32
	s0i32 = l3
	s1i32 = 255
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l6 = s0i32
		s1i32 = l3
		s2i32 = 3
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		l5 = s1i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = 30264
		s1i32 = s1i32 + s2i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s0i32 = l6
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l8 = s1i32
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 30224
			s1i32 = 30224
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = -2
			s3i32 = l5
			s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
			s1i32 = s1i32 & s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			goto lbl10
		}
		s0i32 = l6
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		goto lbl10
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l8 = s0i32
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l3 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l2 = s1i32
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		}
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		goto lbl13
	}
	s0i32 = l4
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = 0
	l3 = s0i32
	goto lbl13
lbl16:
lbl17:
	s0i32 = l2
	l5 = s0i32
	s0i32 = l6
	l3 = s0i32
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l6 = s0i32
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl13:
	s0i32 = l8
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l5 = s1i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 30528
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		if s0i32 != 0 {
			goto lbl18
		}
		s0i32 = 30228
		s1i32 = 30228
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = -2
		s3i32 = l5
		s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
		s1i32 = s1i32 & s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl10
	}
	s0i32 = l8
	s1i32 = 16
	s2i32 = 20
	s3i32 = l8
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	s4i32 = l4
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
lbl18:
	s0i32 = l3
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l2 = s0i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l3
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
lbl10:
	s0i32 = l10
	s1i32 = 15
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l7
		s2i32 = 1
		s1i32 = s1i32 & s2i32
		s2i32 = l9
		s1i32 = s1i32 | s2i32
		s2i32 = 2
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l9
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = 1
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		goto lbl4
	}
	s0i32 = l0
	s1i32 = l7
	s2i32 = 1
	s1i32 = s1i32 & s2i32
	s2i32 = l1
	s1i32 = s1i32 | s2i32
	s2i32 = 2
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l10
	s2i32 = 3
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l9
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = 1
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l10
	f535(ctx, s0i32, s1i32)
lbl4:
	s0i32 = l0
	l2 = s0i32
lbl1:
	s0i32 = l2
	return s0i32
}
