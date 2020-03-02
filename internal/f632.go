package internal

import (
	"unsafe"
)

func f632(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int64
	_ = l16
	var l17 int64
	_ = l17
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		s2i32 = 8
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = 128
		s2i32 = s2i32 + s3i32
		s3i32 = 8
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
		s4i32 = 128
		s3i32 = s3i32 + s4i32
		s4i32 = 8
		s3i32 = s3i32 >> (uint32(s4i32) & 31)
		s4i32 = l2
		s5i32 = 1
		f148(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l12 = s1i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l8 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l13 = s1i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l9 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l14 = s1i32
	s2i32 = 65535
	s1i32 = s1i32 + s2i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l10 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l15 = s1i32
	s2i32 = 65535
	s1i32 = s1i32 + s2i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l11 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l10
		s1i32 = l8
		s0i32 = s0i32 - s1i32
		l4 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l11
		s1i32 = l9
		s0i32 = s0i32 - s1i32
		l5 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l4
		s1i32 = l5
		s0i32 = s0i32 | s1i32
		s0i64 = int64(s0i32)
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967296
		if uint64(s0i64) < uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	lbl5:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l7 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l4 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l5 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
		goto lbl2
	lbl4:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l4 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l5 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l7 = s0i32
		s1i32 = l11
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l4
		s1i32 = l10
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l5
		s1i32 = l9
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l6
		s1i32 = l8
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l4
		s0i64 = int64(s0i32)
		s1i32 = l6
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l16 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l7
		s0i64 = int64(s0i32)
		s1i32 = l5
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l17 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i64 = l16
		s1i64 = l17
		s0i64 = s0i64 | s1i64
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967295
		if uint64(s0i64) > uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l12
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = l13
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		s2i32 = 8
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = l14
		s3i32 = 128
		s2i32 = s2i32 + s3i32
		s3i32 = 8
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s3i32 = l15
		s4i32 = 128
		s3i32 = s3i32 + s4i32
		s4i32 = 8
		s3i32 = s3i32 >> (uint32(s4i32) & 31)
		s4i32 = l2
		s5i32 = 1
		f148(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l3
	s3i32 = 80
	s2i32 = s2i32 + s3i32
	s0i32 = f111(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+60)])
	if s0i32 != 0 {
		goto lbl0
	}
lbl6:
	s0i32 = l3
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l3
	s2i32 = l0
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		s2i32 = 8
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = 128
		s2i32 = s2i32 + s3i32
		s3i32 = 8
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s3i32 = l3
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
		s4i32 = 128
		s3i32 = s3i32 + s4i32
		s4i32 = 8
		s3i32 = s3i32 >> (uint32(s4i32) & 31)
		s4i32 = l2
		s5i32 = 1
		f148(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	}
	s0i32 = l1
	f110(ctx, s0i32)
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+60)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	goto lbl0
lbl2:
	s0i32 = l3
	s1i32 = l7
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l4
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l5
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l6
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s2i32 = l0
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
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = 128
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	s3i32 = 128
	s2i32 = s2i32 + s3i32
	s3i32 = 8
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
	s4i32 = 128
	s3i32 = s3i32 + s4i32
	s4i32 = 8
	s3i32 = s3i32 >> (uint32(s4i32) & 31)
	s4i32 = l2
	s5i32 = 1
	f148(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
lbl0:
	s0i32 = l3
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
