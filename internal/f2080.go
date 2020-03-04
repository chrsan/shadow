package internal

import (
	"unsafe"
)

func f2080(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
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
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+52)]))
	s4i32 = l2
	s5i32 = l3
	f672(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	s0i32 = l0
	s1i32 = l1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l2 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			l3 = s0i32
			goto lbl3
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l4 = s0i32
		s0i32 = l2
		s1i32 = 2
		s0i32 = f44(ctx, s0i32, s1i32)
		l3 = s0i32
		s1i32 = l4
		s2i32 = l2
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	lbl3:
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l2 = s0i32
		s0i32 = l0
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l2
		f13(ctx, s0i32)
		goto lbl1
	}
	s0i32 = l0
	s1i32 = l2
	if s1i32 != 0 {
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
		s3i32 = l2
		s1i32 = f22(ctx, s1i32, s2i32, s3i32)
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	} else {
		s1i32 = 0
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
lbl1:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l2 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			l3 = s0i32
			goto lbl8
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l4 = s0i32
		s0i32 = l2
		s1i32 = 3
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l5 = s0i32
		s1i32 = 2
		s0i32 = f44(ctx, s0i32, s1i32)
		l3 = s0i32
		s1i32 = l4
		s2i32 = l5
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	lbl8:
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l2 = s0i32
		s0i32 = l0
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l2
		f13(ctx, s0i32)
		goto lbl6
	}
	s0i32 = l0
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l3 = s1i32
	if s1i32 != 0 {
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		s3i32 = l3
		s1i32 = f22(ctx, s1i32, s2i32, s3i32)
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	} else {
		s1i32 = l2
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
lbl6:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	l2 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			l3 = s0i32
			goto lbl12
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l4 = s0i32
		s0i32 = l2
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l5 = s0i32
		s1i32 = 2
		s0i32 = f44(ctx, s0i32, s1i32)
		l3 = s0i32
		s1i32 = l4
		s2i32 = l5
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	lbl12:
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l2 = s0i32
		s0i32 = l0
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l2
		f13(ctx, s0i32)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l2
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l3 = s1i32
	if s1i32 != 0 {
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
		s3i32 = l3
		s1i32 = f22(ctx, s1i32, s2i32, s3i32)
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
	} else {
		s1i32 = l2
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
lbl0:
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+84)])
	l2 = s1i32
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l1
		s1i32 = int32(ctx.Mem[int(s1i32+85)])
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+90)])
	ctx.Mem[int(s0i32+90)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+86)])
	ctx.Mem[int(s0i32+86)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+87)])
	ctx.Mem[int(s0i32+87)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+88)])
	ctx.Mem[int(s0i32+88)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+89)])
	ctx.Mem[int(s0i32+89)] = uint8(s1i32)
}
