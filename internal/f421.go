package internal

import (
	"unsafe"
)

func f421(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int64
	_ = l7
	var l8 float32
	_ = l8
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = 144
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 2
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 65535
	s1i32 = s1i32 & s2i32
	s2i32 = l3
	s3i32 = 128
	s2i32 = s2i32 + s3i32
	s2i32 = f37(ctx, s2i32)
	l6 = s2i32
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32) int32)(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+54)])
	s1i32 = 16
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s1i32 = 14
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 49152
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	s1i32 = l1
	s2i32 = 4
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 49152
	s1i32 = s1i32 & s2i32
	l1 = s1i32
	s0i32 = s0i32 | s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l6
	s1i32 = l4
	s1f32 = float32(s1i32)
	s2f32 = 1.5258789e-05
	s1f32 = s1f32 * s2f32
	s2i32 = l1
	s2f32 = float32(s2i32)
	s3f32 = 1.5258789e-05
	s2f32 = s2f32 * s3f32
	s3i32 = l6
	f2091(ctx, s0i32, s1f32, s2f32, s3i32)
lbl1:
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s1f32 = 0
	if s0f32 > s1f32 {
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
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = l3
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l1 = s0i32
	s0i32 = l3
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l0
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l7 = s0i64
	s0i32 = l3
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 128
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 4575657221408423936
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = l7
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l3
	s2i32 = 72
	s1i32 = s1i32 + s2i32
	s1i32 = f24(ctx, s1i32)
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		goto lbl6
	}
	s0i32 = l3
	s1i32 = 72
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s0i32 = f84(ctx, s0i32, s1i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl6:
	s0i32 = l6
	s1i32 = l3
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	f138(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i64 = 1082130432
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 1065353216
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1f32 = -1
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l8 = s0f32
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1f32 = l8
		s2i32 = l0
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+54)])))
		s3i32 = 1
		s2i32 = s2i32 & s3i32
		f1686(ctx, s0i32, s1f32, s2i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+53)])
		l5 = s0i32
		s0i32 = l4
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = -2147483648
		s1i32 = s1i32 & s2i32
		s2i32 = l5
		s3i32 = 15
		s2i32 = s2i32 & s3i32
		s3i32 = 16
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = l5
		s4i32 = 4
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s2i32 = s2i32 | s3i32
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	if s0i32 != 0 {
		s0i32 = l3
		s0i32 = f37(ctx, s0i32)
		l5 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
		s1i32 = l5
		s2i32 = l1
		s3i32 = l4
		s4i32 = 0
		s0i32 = f674(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l5
			f194(ctx, s0i32, s1i32)
		}
		s0i32 = l5
		f34(ctx, s0i32)
	}
	s0i32 = l4
	s0i32 = f617(ctx, s0i32)
	s1i32 = 1
	s0i32 = s0i32 | s1i32
	s1i32 = 3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l3
		s1i32 = f37(ctx, s1i32)
		l0 = s1i32
		s2i32 = l1
		s0i32 = f616(ctx, s0i32, s1i32, s2i32)
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l0
			f194(ctx, s0i32, s1i32)
		}
		s0i32 = l0
		f34(ctx, s0i32)
	}
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l3
		s2i32 = 72
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		f138(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l1
	f34(ctx, s0i32)
	s0i32 = l2
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = 1
	l4 = s0i32
	goto lbl0
lbl5:
	s0i32 = l1
	f34(ctx, s0i32)
	s0i32 = 1
	l4 = s0i32
	goto lbl0
lbl3:
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l4 = s0i32
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l6
	f194(ctx, s0i32, s1i32)
lbl2:
	s0i32 = 1
	l4 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
lbl0:
	s0i32 = l6
	f34(ctx, s0i32)
	s0i32 = l3
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	return s0i32
}
