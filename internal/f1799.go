package internal

import (
	"unsafe"
)

func f1799(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var l14 int64
	_ = l14
	var l15 int64
	_ = l15
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
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	var s6i64 int64
	_ = s6i64
	s0i32 = ctx.g0
	s1i32 = 1152
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l6 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l9 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l11 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l14 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l9
	s0i64 = int64(s0i32)
	s1i32 = l6
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l15 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l14
	s1i64 = l15
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
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s1i32 = l9
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l10 = s1i32
	s2i32 = l8
	if s1i32 >= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l12 = s2i32
	s3i32 = l6
	if s2i32 <= s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l13 = s3i32
	s4i32 = l11
	if s3i32 <= s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4i32 = l10
	s4i64 = int64(s4i32)
	s5i32 = l13
	s5i64 = int64(s5i32)
	s4i64 = s4i64 - s5i64
	l14 = s4i64
	s5i64 = 0
	if s4i64 > s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	s5i32 = l7
	s5i64 = int64(s5i32)
	s6i32 = l12
	s6i64 = int64(s6i32)
	s5i64 = s5i64 - s6i64
	l15 = s5i64
	s6i64 = 0
	if s5i64 > s6i64 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	s4i32 = s4i32 & s5i32
	s5i64 = l14
	s6i64 = l15
	s5i64 = s5i64 | s6i64
	s6i64 = 2147483648
	s5i64 = s5i64 + s6i64
	s6i64 = 4294967296
	if uint64(s5i64) < uint64(s6i64) {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	s4i32 = s4i32 & s5i32
	s3i32 = s3i32 & s4i32
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 & s1i32
	l7 = s0i32
lbl0:
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	l10 = s0i32
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l8
	s1i32 = l11
	s0i32 = s0i32 - s1i32
	l8 = s0i32
	s1i32 = 32
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l8
	s1i32 = 3
	s0i32 = s0i32 + s1i32
	s1i32 = -4
	s0i32 = s0i32 & s1i32
	s0i64 = int64(s0i32)
	s1i32 = l9
	s2i32 = l6
	s1i32 = s1i32 - s2i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 * s1i64
	s1i64 = 1024
	if s0i64 > s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l4
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 24580
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 1136
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = l3
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1140)]))
		l4 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1136)]))
		l6 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1144)]))
		goto lbl3
	}
	s0i32 = l5
	s1i32 = 1144
	s0i32 = s0i32 + s1i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1136)])) = uint64(s1i64)
	s0i32 = 0
	l4 = s0i32
	s0i32 = 0
	l6 = s0i32
	s0i32 = 0
lbl3:
	l1 = s0i32
	s0i32 = l5
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l6
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l1
	s2i32 = l6
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l4
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l4
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l5
	s2i32 = 108
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 24708
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint64(s1i64)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l1 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = l5
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l1
	s2i32 = l6
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 92
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l1
	s2i32 = l3
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	}
	s0i32 = l4
	s1i32 = 0
	s2i32 = l5
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+84)]))
	s3i32 = l5
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+80)]))
	s4i32 = l5
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+72)]))
	s3i32 = s3i32 - s4i32
	s2i32 = s2i32 * s3i32
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = l3
	s2i32 = l5
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s4i32 = l2
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s5i32 = 2
	s6i32 = l7
	f410(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	s0i32 = l5
	s1i32 = 24708
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l0 = s0i32
	s1i32 = l5
	s2i32 = -64
	s1i32 = s1i32 - s2i32
	s2i32 = l1
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32))(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s0i32 = f35(ctx, s0i32)
	goto lbl1
lbl2:
	s0i32 = l0
	s1i32 = l3
	s2i32 = l5
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	s4i32 = l2
	s5i32 = l3
	s6i32 = l10
	s7i32 = 0
	if s6i32 != s7i32 {
		s6i32 = 1
	} else {
		s6i32 = 0
	}
	s2i32 = f1803(ctx, s2i32, s3i32, s4i32, s5i32, s6i32)
	l0 = s2i32
	s3i32 = l2
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s4i32 = l2
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s5i32 = 2
	s6i32 = l7
	f410(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	s0i32 = l0
	s0i32 = f419(ctx, s0i32)
lbl1:
	s0i32 = l5
	s1i32 = 1152
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
