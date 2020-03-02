package internal

import (
	"unsafe"
)

func f2053(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
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
	var s3i64 int64
	_ = s3i64
	s0i32 = ctx.g0
	s1i32 = 224
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	s3i32 = 8191
	if s2i32 <= s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
		s3i32 = 8192
		if s2i32 < s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			goto lbl2
		}
	}
	s2i32 = l1
	s2i32 = int32(ctx.Mem[int(s2i32+10)])
	s3i32 = 2
	s2i32 = s2i32 & s3i32
	if s2i32 != 0 {
		goto lbl2
	}
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l0 = s2i32
	s2i32 = int32(ctx.Mem[int(s2i32+84)])
	if s2i32 != 0 {
		s2i32 = l0
		s3i32 = 4
		s2i32 = s2i32 + s3i32
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
		s4i32 = l0
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+28)]))
		s2i32 = f28(ctx, s2i32, s3i32, s4i32)
		l4 = s2i32
		s2i32 = l0
		s3i32 = 0
		ctx.Mem[int(s2i32+84)] = uint8(s3i32)
		s2i32 = l0
		s3i32 = l4
		ctx.Mem[int(s2i32+85)] = uint8(s3i32)
	}
	s2i32 = l5
	s3i32 = l2
	s3i32 = f65(ctx, s3i32)
	l4 = s3i32
	ctx.Mem[int(s2i32+24)] = uint8(s3i32)
	s2i32 = l4
	if s2i32 != 0 {
		s2i32 = l0
		s3i32 = 4
		s2i32 = s2i32 + s3i32
		l4 = s2i32
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
		s3i32 = 6
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 3
		s2i32 = s2i32 & s3i32
		l0 = s2i32
		if s2i32 != 0 {
			goto lbl7
		}
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
		s3i32 = l2
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s4i32 = l2
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
		s3i32 = s3i32 | s4i32
		s2i32 = s2i32 | s3i32
		if s2i32 != 0 {
			goto lbl7
		}
		s2i32 = 1
		goto lbl6
	lbl7:
		s2i32 = l2
		s3i32 = l4
		s4i32 = l5
		s5i32 = 8
		s4i32 = s4i32 + s5i32
		s5i32 = l0
		s2i32 = f64(ctx, s2i32, s3i32, s4i32, s5i32)
		l4 = s2i32
		s2i32 = l5
		s2i32 = int32(ctx.Mem[int(s2i32+24)])
	lbl6:
		l6 = s2i32
		s2i32 = l5
		s3i32 = l4
		s3i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)])) = uint64(s3i64)
		s2i32 = l5
		s3i32 = l4
		s3i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)])) = uint64(s3i64)
	}
	s2i32 = l5
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = 0
	s4i32 = l6
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	goto lbl1
lbl2:
	s2i32 = 0
lbl1:
	s0i32 = f229(ctx, s0i32, s1i32, s2i32)
	l0 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+188)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+189)])
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l3
	s0i32 = s0i32 & s1i32
	l6 = s0i32
	s0i32 = l0
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	l7 = s0i32
lbl8:
	s0i32 = l4
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
	lbl12:
		s0i32 = l0
		f228(ctx, s0i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+173)])
		l3 = s0i32
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+188)])
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l3
		if s0i32 != 0 {
			goto lbl12
		}
		goto lbl9
		panic("unreachable executed")
		panic("unreachable executed")
	}
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+188)] = uint8(s1i32)
	goto lbl9
lbl10:
	s0i32 = l3
	if s0i32 != 0 {
		goto lbl0
	}
lbl9:
	s0i32 = l7
	s1i32 = l1
	s2i32 = l2
	s3i32 = 0
	s4i32 = l6
	f132(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+188)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+189)])
	l4 = s0i32
	goto lbl8
	panic("unreachable executed")
	panic("unreachable executed")
lbl0:
	s0i32 = l0
	s1i32 = 132
	s0i32 = s0i32 + s1i32
	f112(ctx, s0i32)
	s0i32 = l0
	s1i32 = 21300
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l2 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l1
	f12(ctx, s0i32)
lbl13:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
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
		goto lbl14
	}
	s0i32 = l0
	f12(ctx, s0i32)
lbl14:
	s0i32 = l5
	s1i32 = 224
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
