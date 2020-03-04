package internal

import (
	"unsafe"
)

func f1055(ctx *Context, l0 int32, l1 int32) {
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
	var l11 int32
	_ = l11
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
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 52
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		f71(ctx, s0i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l2 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0i32
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l5
	s0i32 = s0i32 ^ s1i32
	s1i32 = -2048144789
	s0i32 = s0i32 * s1i32
	l1 = s0i32
	s1i32 = 13
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l1
	s0i32 = s0i32 ^ s1i32
	s1i32 = -1028477387
	s0i32 = s0i32 * s1i32
	l1 = s0i32
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l1
	s0i32 = s0i32 ^ s1i32
	l1 = s0i32
	s1i32 = 1
	s2i32 = l1
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l9 = s0i32
	s1i32 = l2
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l6 = s0i32
lbl2:
	s0i32 = l6
	s1i32 = l1
	s2i32 = 20
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l11 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l9
	s1i32 = l11
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l10
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l5
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = l1
	s1i32 = 0
	s2i32 = l2
	s3i32 = l1
	s4i32 = 0
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
	s0i32 = s0i32 + s1i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l7
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = l2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	goto lbl1
lbl3:
	s0i32 = l6
	s1i32 = l1
	s2i32 = 20
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l1 = s0i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l2 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl6:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		f287(ctx, s0i32, s1i32)
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l2
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
	}
	s0i32 = l8
	s1i32 = l3
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	f489(ctx, s0i32, s1i32)
lbl1:
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 1
	s1i32 = 0
	s2i32 = l0
	s1i32 = s1i32 - s2i32
	s2i32 = l0
	s3i32 = -1
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
	l0 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l0
		f72(ctx, s0i32, s1i32)
	}
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
