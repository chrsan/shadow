package internal

import (
	"unsafe"
)

func f740(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s6i32 int32
	_ = s6i32
	var s0f32 float32
	_ = s0f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s1i32 = 24
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 24
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l6 = s0i32
	s1i32 = -97
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = 6
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l2
		l9 = s0f32
		s0i32 = l6
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = 1
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			s0f32 = 0
			l9 = s0f32
			goto lbl0
		}
		s0f32 = 1
		l9 = s0f32
		goto lbl0
	}
	s0f32 = l2
	l9 = s0f32
	s0i32 = l8
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl4
	case 1:
		goto lbl0
	case 2:
		goto lbl0
	case 3:
		goto lbl0
	case 4:
		goto lbl0
	case 5:
		goto lbl3
	default:
		goto lbl5
	}
lbl5:
	s0f32 = l5
	l9 = s0f32
	goto lbl0
lbl4:
	s0f32 = l4
	l9 = s0f32
	goto lbl0
lbl3:
	s0f32 = l3
	l9 = s0f32
lbl0:
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 24
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l6 = s0i32
	s1i32 = -97
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = 17
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l3
		l10 = s0f32
		s0i32 = l6
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = 1
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = 1
			l10 = s0f32
			goto lbl6
		}
		s0f32 = 0
		l10 = s0f32
		goto lbl6
	}
	s0f32 = l3
	l10 = s0f32
	s0i32 = l8
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl10
	case 1:
		goto lbl6
	case 2:
		goto lbl6
	case 3:
		goto lbl6
	case 4:
		goto lbl6
	case 5:
		goto lbl6
	case 6:
		goto lbl6
	case 7:
		goto lbl6
	case 8:
		goto lbl6
	case 9:
		goto lbl6
	case 10:
		goto lbl6
	case 11:
		goto lbl6
	case 12:
		goto lbl6
	case 13:
		goto lbl6
	case 14:
		goto lbl6
	case 15:
		goto lbl6
	case 16:
		goto lbl11
	default:
		goto lbl9
	}
lbl11:
	s0f32 = l2
	l10 = s0f32
	goto lbl6
lbl10:
	s0f32 = l4
	l10 = s0f32
	goto lbl6
lbl9:
	s0f32 = l5
	l10 = s0f32
lbl6:
	s0i32 = l7
	s1i32 = 8
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 24
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l6 = s0i32
	s1i32 = -97
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = 17
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l4
		l11 = s0f32
		s0i32 = l6
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = 1
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
		}
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = 1
			l11 = s0f32
			goto lbl12
		}
		s0f32 = 0
		l11 = s0f32
		goto lbl12
	}
	s0f32 = l4
	l11 = s0f32
	s0i32 = l8
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl12
	case 1:
		goto lbl12
	case 2:
		goto lbl12
	case 3:
		goto lbl12
	case 4:
		goto lbl12
	case 5:
		goto lbl16
	case 6:
		goto lbl12
	case 7:
		goto lbl12
	case 8:
		goto lbl12
	case 9:
		goto lbl12
	case 10:
		goto lbl12
	case 11:
		goto lbl12
	case 12:
		goto lbl12
	case 13:
		goto lbl12
	case 14:
		goto lbl12
	case 15:
		goto lbl12
	case 16:
		goto lbl17
	default:
		goto lbl15
	}
lbl17:
	s0f32 = l2
	l11 = s0f32
	goto lbl12
lbl16:
	s0f32 = l3
	l11 = s0f32
	goto lbl12
lbl15:
	s0f32 = l5
	l11 = s0f32
lbl12:
	s0i32 = l7
	s1i32 = 24
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l7 = s0i32
	s1i32 = -98
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = 16
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = 1
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
		s0i32 = l7
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = 1
			l5 = s0f32
			goto lbl18
		}
		s0f32 = 0
		l5 = s0f32
		goto lbl18
	}
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl18
	case 1:
		goto lbl18
	case 2:
		goto lbl18
	case 3:
		goto lbl18
	case 4:
		goto lbl22
	case 5:
		goto lbl18
	case 6:
		goto lbl18
	case 7:
		goto lbl18
	case 8:
		goto lbl18
	case 9:
		goto lbl18
	case 10:
		goto lbl18
	case 11:
		goto lbl18
	case 12:
		goto lbl18
	case 13:
		goto lbl18
	case 14:
		goto lbl18
	case 15:
		goto lbl23
	default:
		goto lbl21
	}
lbl23:
	s0f32 = l2
	l5 = s0f32
	goto lbl18
lbl22:
	s0f32 = l3
	l5 = s0f32
	goto lbl18
lbl21:
	s0f32 = l4
	l5 = s0f32
lbl18:
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2f32 = l9
	s3f32 = l10
	s4f32 = l11
	s5f32 = l5
	s6i32 = l1
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, float32, float32, float32, float32))(table[s6i32].f()))(ctx, s0i32, s1i32, s2f32, s3f32, s4f32, s5f32)
}
