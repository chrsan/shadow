package internal

import (
	"unsafe"
)

func f2042(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s1i64 int64
	_ = s1i64
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 208
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	l6 = s0i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s1i32 = 8191
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l6 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		s1i32 = 8192
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l6 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l7 = s0i32
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l7
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l5
	s1i32 = l6
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l2
	s1i32 = l5
	s2i32 = 192
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s0i32 = f42(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = l4
	s1i32 = f65(ctx, s1i32)
	l7 = s1i32
	ctx.Mem[int(s0i32+16)] = uint8(s1i32)
	s0i32 = 0
	l6 = s0i32
	s0i32 = l7
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	s1i32 = 6
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l7 = s0i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s1i32 = s1i32 | s2i32
	s0i32 = s0i32 | s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+200)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+192)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	goto lbl2
lbl3:
	s0i32 = l5
	s1i32 = l4
	s2i32 = l5
	s3i32 = 192
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = l7
	s1i32 = f64(ctx, s1i32, s2i32, s3i32, s4i32)
	l7 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l5
	s0i32 = int32(ctx.Mem[int(s0i32+16)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl2:
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 192
	s0i32 = s0i32 + s1i32
	l6 = s0i32
lbl0:
	s0i32 = l5
	s1i32 = l0
	s2i32 = l6
	s0i32 = f229(ctx, s0i32, s1i32, s2i32)
	l0 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+188)])
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l0
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	l7 = s0i32
lbl5:
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+189)])
	if s0i32 != 0 {
	lbl8:
		s0i32 = l0
		f228(ctx, s0i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+173)])
		l6 = s0i32
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+188)])
		if s0i32 != 0 {
			goto lbl9
		}
		s0i32 = l6
		if s0i32 != 0 {
			goto lbl8
		}
		goto lbl6
	lbl9:
		s0i32 = l6
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		goto lbl4
	}
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+188)] = uint8(s1i32)
lbl6:
	s0i32 = l7
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
	f502(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+188)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl4:
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
		goto lbl10
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
		goto lbl10
	}
	s0i32 = l1
	f12(ctx, s0i32)
lbl10:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
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
		goto lbl11
	}
	s0i32 = l0
	f12(ctx, s0i32)
lbl11:
	s0i32 = l5
	s1i32 = 208
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
