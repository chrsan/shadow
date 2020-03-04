package internal

import (
	"unsafe"
)

func f1059(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	var s1i64 int64
	_ = s1i64
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
	l3 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 52
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l3
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		f71(ctx, s0i32)
	}
	s0i32 = l1
	s1i32 = 80
	s2i32 = 0
	s3i32 = 22140
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
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
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l4 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	l9 = s0i32
	s0i32 = l1
	s1i32 = 44
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s0i32 = s0i32 & s1i32
	l7 = s0i32
lbl2:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s1i32 = l7
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			goto lbl1
		}
		goto lbl4
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l3
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l11
	s0i32 = f348(ctx, s0i32, s1i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l3
	s1i32 = 44
	s0i32 = s0i32 + s1i32
	s1i32 = l10
	s2i32 = 16
	s0i32 = f102(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l3
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s1i32 = l9
	s2i32 = 16
	s0i32 = f102(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
lbl6:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l4 = s0i32
lbl4:
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = l8
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l8 = s1i32
	s2i32 = l7
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 & s1i32
	l7 = s0i32
	s0i32 = l8
	s1i32 = l4
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	goto lbl1
lbl3:
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	l5 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
	l1 = s0i32
	s1i32 = 100
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 24
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+100)]))
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 28
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l5
	s2i32 = l4
	s3i32 = l1
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint32(s1i32)
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
lbl7:
	s0i32 = l3
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
lbl9:
	s0i32 = l2
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = 1
	l5 = s0i32
lbl1:
	s0i32 = l6
	s1i32 = l6
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
		s0i32 = l6
		s1i32 = l0
		f72(ctx, s0i32, s1i32)
	}
	s0i32 = l5
	return s0i32
}
