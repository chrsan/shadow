package internal

import (
	"unsafe"
)

func f1821(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		f128(ctx, s0i32)
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s1i32 = 0
	s2i32 = l1
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+24)]))
	s2i32 = s2i32 - s3i32
	l5 = s2i32
	s3i32 = l5
	s4i32 = 0
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	l9 = s3i32
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l8 = s1i32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	s1i32 = l5
	s2i32 = 31
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = l5
	s1i32 = s1i32 & s2i32
	s2i32 = l4
	s1i32 = s1i32 + s2i32
	l1 = s1i32
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
	l7 = s0i32
	s0i32 = l8
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
	l1 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		s0i32 = 0
		l1 = s0i32
	}
	s0i32 = l7
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		return
	}
	s0i32 = l3
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	l10 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l11 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l5 = s0i32
	s1i32 = l8
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l8
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l1
	s1i32 = l11
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l5
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l5 = s0i32
lbl4:
	s0i32 = l4
	s1i32 = l5
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l1 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l5
		s1i32 = l4
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l4
		s1i32 = s1i32 - s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		goto lbl3
	}
	s0i32 = l1
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l5
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l4
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
lbl3:
	s0i32 = l10
	s1i32 = l3
	s2i32 = l9
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l10 = s0i32
	s0i32 = l8
	s1i32 = l11
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l7
	s1i32 = l2
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l1 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		l5 = s0i32
		s0i32 = l4
		l6 = s0i32
		s0i32 = l2
		l3 = s0i32
		goto lbl7
	}
	s0i32 = l2
	l3 = s0i32
	s0i32 = l4
	l6 = s0i32
	s0i32 = l7
	l5 = s0i32
lbl9:
	s0i32 = l5
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l5
	s1i32 = l3
	s2i32 = l1
	s3i32 = 1
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l1 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
lbl7:
	s0i32 = l5
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l3
	s1i32 = l5
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l3
	s1i32 = l5
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l5
	s1i32 = s1i32 - s2i32
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
lbl6:
	s0i32 = l7
	l5 = s0i32
lbl10:
	s0i32 = l4
	s1i32 = l2
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l1 = s1i32
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l2
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = 0
	l3 = s0i32
	s0i32 = l5
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l0
	s1i32 = l4
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+48)]))
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
lbl12:
	s0i32 = 1
	l2 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l1 = s0i32
	s1i32 = l3
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	l9 = s1i32
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l6 = s1i32
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl14:
		s0i32 = l1
		s1i32 = l2
		s2i32 = l9
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = 1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l1 = s0i32
		s1i32 = l4
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l9
		s1i32 = s1i32 + s2i32
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		l1 = s1i32
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
	}
	s0i32 = l5
	s1i32 = 1
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l7
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = 0
	l2 = s0i32
	s0i32 = l7
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
lbl15:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	s1i32 = l2
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	s2i32 = l2
	s3i32 = l10
	s2i32 = s2i32 + s3i32
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s2i32 = l1
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l7
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
lbl11:
}
