package internal

import (
	"unsafe"
)

func f264(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	if int(s2i32) < 0 || int(s2i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s2i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s2i32].numParams != 2 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32))(table[s2i32].f()))(ctx, s0i32, s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	l14 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l15 = s0i32
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32) int32)(table[s1i32].f()))(ctx, s0i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l5 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l6 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		l7 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l8 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		l9 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l10 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l11 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
		l12 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l16 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l5 = s0i32
		s0i32 = l4
		s1i32 = 128
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l16
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l12
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s0i32 = f24(ctx, s0i32)
		s1i32 = 12
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l0 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l5 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		l6 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l7 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		l8 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l9 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l10 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
		l11 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l12 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l0 = s0i32
		s0i32 = l4
		s1i32 = 128
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l12
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l1
		f966(ctx, s0i32, s1i32, s2i32)
		s0i32 = l15
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l14
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l17 = s0f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l18 = s1f32
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
		s0f32 = l17
		s1f32 = 0
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l17 = s0f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l19 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0f32 = l19
		s1i32 = l14
		s1f32 = float32(s1i32)
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl0
		}
		s0f32 = l18
		s1i32 = l15
		s1f32 = float32(s1i32)
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl0
		}
		s0f32 = l17
		s1f32 = 0
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		s1i32 = 6
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l0 = s0i32
		s1i32 = 3
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l2
	s1i32 = l3
	s0i32 = f452(ctx, s0i32, s1i32)
	l13 = s0i32
lbl0:
	s0i32 = l4
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
	s0i32 = l13
	return s0i32
}
