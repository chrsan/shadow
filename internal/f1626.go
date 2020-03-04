package internal

import (
	"unsafe"
)

func f1626(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
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
	var s1i64 int64
	_ = s1i64
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 144
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+86)])
	if s0i32 != 0 {
		s0i32 = l3
		s0i32 = int32(ctx.Mem[int(s0i32+84)])
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
			s2i32 = l3
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
			s0i32 = f28(ctx, s0i32, s1i32, s2i32)
			l4 = s0i32
			s0i32 = l3
			s1i32 = 0
			ctx.Mem[int(s0i32+84)] = uint8(s1i32)
			s0i32 = l3
			s1i32 = l4
			ctx.Mem[int(s0i32+85)] = uint8(s1i32)
			s0i32 = l3
			s0i32 = int32(ctx.Mem[int(s0i32+86)])
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l4 = s0i32
		}
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
		s0i32 = l4
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l3 = s0i32
	}
	s0i32 = 0
	l4 = s0i32
	s0i32 = l3
	s0i32 = int32(ctx.Mem[int(s0i32+87)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 88
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	f2071(ctx, s0i32, s1i32)
	s0i32 = l2
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+136)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s0i32 = int32(ctx.Mem[int(s0i32+87)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = int32(ctx.Mem[int(s0i32+89)])
	l4 = s0i32
	s0i32 = l3
	s0i32 = int32(ctx.Mem[int(s0i32+88)])
	l3 = s0i32
	goto lbl1
lbl2:
	s0i32 = l3
	s0i32 = int32(ctx.Mem[int(s0i32+89)])
	l4 = s0i32
	s0i32 = l3
	s0i32 = int32(ctx.Mem[int(s0i32+88)])
	l3 = s0i32
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = -64
	s1i32 = s1i32 - s2i32
	s0i32 = f1881(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s1f32 = s1f32 - s2f32
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		l5 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = s1f32
		s0i32 = l2
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s1f32 = s1f32 - s2f32
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		l6 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = s1f32
		s0i32 = l2
		s1f32 = l5
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
		s0i32 = l2
		s1f32 = l6
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
		s0i32 = l2
		s1f32 = l5
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
		s0i32 = l2
		s1f32 = l6
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = l2
		s1f32 = l5
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l2
		s1i32 = 2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
		s0i32 = l2
		s1f32 = l6
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	}
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l4 = s0i32
lbl1:
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 56
		l4 = s0i32
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	l0 = s0i32
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l0
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	s2i32 = 8
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 335544325
	s3i32 = 268435461
	s4i32 = l3
	s5i32 = 255
	s4i32 = s4i32 & s5i32
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 88
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 84
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	f82(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = 88
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = 48
	f82(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 88
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 84
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	f82(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = 88
	s0i32 = s0i32 + s1i32
	f673(ctx, s0i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
	s0i32 = s0i32 - s1i32
	l4 = s0i32
lbl0:
	s0i32 = l2
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	return s0i32
}
