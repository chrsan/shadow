package internal

import (
	"math"
	"unsafe"
)

func f292(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
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
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 20
		s2i32 = 4
		s0i32 = f56(ctx, s0i32, s1i32, s2i32)
		l5 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l0 = s0i32
		s0i32 = l3
		s1i32 = l5
		s2i32 = 12
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 228
		s2i32 = l5
		s3i32 = l0
		s2i32 = s2i32 - s3i32
		f52(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = 20568
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		goto lbl0
	}
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	s1i32 = 261120
	s0i32 = s0i32 & s1i32
	s1i32 = 3072
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		l5 = s0i32
		goto lbl3
	}
	s0i32 = l2
	s1i32 = 691749
	s2i32 = l5
	s3i32 = 2097151
	s2i32 = s2i32 & s3i32
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 1
	s1i32 = s1i32 & s2i32
	s0i32 = f1595(ctx, s0i32, s1i32)
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = 1
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		l5 = s0i32
		goto lbl3
	}
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 20
		s2i32 = 4
		s0i32 = f56(ctx, s0i32, s1i32, s2i32)
		l5 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l0 = s0i32
		s0i32 = l3
		s1i32 = l5
		s2i32 = 12
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 228
		s2i32 = l5
		s3i32 = l0
		s2i32 = s2i32 - s3i32
		f52(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = 20568
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		goto lbl2
	}
	s0i32 = l6
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s1i32 = 4
	s0i32 = s0i32 | s1i32
	l5 = s0i32
	s1i32 = l2
	s0i32 = f46(ctx, s0i32, s1i32)
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
	s2i32 = -261121
	s1i32 = s1i32 & s2i32
	s2i32 = 3072
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
lbl3:
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	s1i32 = 261120
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = 40
			s0i32 = s0i32 + s1i32
			s1i32 = 4
			s0i32 = s0i32 | s1i32
			l2 = s0i32
			s1i32 = l5
			s0i32 = f46(ctx, s0i32, s1i32)
			l5 = s0i32
			s0i32 = l6
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		}
		s0i32 = l6
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l6
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		f83(ctx, s0i32, s1i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l2 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l2
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l7 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l2
		s1i32 = l2
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
	lbl10:
		s0i32 = l6
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l6
		s2i32 = 24
		s1i32 = s1i32 + s2i32
		f178(ctx, s0i32, s1i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l2 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
		}
		s0i32 = l2
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l7 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
		}
		s0i32 = l2
		s1i32 = l2
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
	lbl11:
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = -261121
		s1i32 = s1i32 & s2i32
		s2i32 = 1024
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 0
		f196(ctx, s0i32, s1i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l5 = s0i32
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	if s0i32 != 0 {
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = 40
			s0i32 = s0i32 + s1i32
			s1i32 = 4
			s0i32 = s0i32 | s1i32
			l2 = s0i32
			s1i32 = l5
			s0i32 = f46(ctx, s0i32, s1i32)
			l5 = s0i32
			s0i32 = l6
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		}
		s0i32 = l5
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		f2117(ctx, s0i32, s1i32)
	}
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = 1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l6
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
			s0i32 = l3
			s1i32 = 48
			s2i32 = 4
			s0i32 = f56(ctx, s0i32, s1i32, s2i32)
			l5 = s0i32
			s0i32 = l3
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l1 = s0i32
			s0i32 = l3
			s1i32 = l5
			s2i32 = 40
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = 229
			s2i32 = l5
			s3i32 = l1
			s2i32 = s2i32 - s3i32
			f52(ctx, s0i32, s1i32, s2i32)
			s0i32 = l5
			s1i32 = l0
			f1271(ctx, s0i32, s1i32)
			goto lbl2
		}
		s0i32 = l3
		s1i32 = 20
		s2i32 = 4
		s0i32 = f56(ctx, s0i32, s1i32, s2i32)
		l5 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l0 = s0i32
		s0i32 = l3
		s1i32 = l5
		s2i32 = 12
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 228
		s2i32 = l5
		s3i32 = l0
		s2i32 = s2i32 - s3i32
		f52(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = 20568
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		goto lbl2
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l2 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+44)])
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0i32 = f696(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l5 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 40
		s0i32 = s0i32 + s1i32
		s1i32 = 4
		s0i32 = s0i32 | s1i32
		l2 = s0i32
		s1i32 = l5
		s0i32 = f46(ctx, s0i32, s1i32)
		l5 = s0i32
		s0i32 = l6
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	}
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = -3
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
lbl16:
	s0i32 = 28004
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		s2i32 = l1
		s3i32 = l3
		s0i32 = f1643(ctx, s0i32, s1i32, s2i32, s3i32)
		l5 = s0i32
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l0
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	s2i32 = l1
	s0i32 = f1198(ctx, s0i32, s1i32, s2i32)
	l4 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l2 = s0i32
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s2i32 = l1
		s3i32 = l3
		s0i32 = f270(ctx, s0i32, s1i32, s2i32, s3i32)
		l5 = s0i32
		goto lbl2
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l2 = s0i32
		goto lbl20
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0i32
	s0i32 = l6
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l6
	s2i32 = l3
	s0i32 = f385(ctx, s0i32, s1i32, s2i32)
	l2 = s0i32
	if s0i32 != 0 {
		goto lbl20
	}
	s0i32 = l0
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	s2i32 = l1
	s3i32 = l3
	s0i32 = f270(ctx, s0i32, s1i32, s2i32, s3i32)
	l5 = s0i32
	goto lbl2
lbl20:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = -2
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl22
	}
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl22
	case 1:
		goto lbl24
	default:
		goto lbl23
	}
lbl24:
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l1 = s0i32
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 84
		s2i32 = 4
		s0i32 = f56(ctx, s0i32, s1i32, s2i32)
		l5 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l4 = s0i32
		s0i32 = l3
		s1i32 = l5
		s2i32 = 76
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 230
		s2i32 = l5
		s3i32 = l4
		s2i32 = s2i32 - s3i32
		f52(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = l0
		s2i32 = l1
		s3i32 = l2
		f1243(ctx, s0i32, s1i32, s2i32, s3i32)
		goto lbl2
	}
	s0i32 = l1
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = f209(ctx, s0i32)
	l2 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l1 = s0i32
	s0i32 = l2
	s1i32 = -16777216
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 72
		s2i32 = 4
		s0i32 = f56(ctx, s0i32, s1i32, s2i32)
		l5 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l2 = s0i32
		s0i32 = l3
		s1i32 = l5
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 231
		s2i32 = l5
		s3i32 = l2
		s2i32 = s2i32 - s3i32
		f52(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = l0
		s2i32 = l1
		f366(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = 20312
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl2
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l8 = s0f32
	s0i32 = l3
	s1i32 = 72
	s2i32 = 4
	s0i32 = f56(ctx, s0i32, s1i32, s2i32)
	l5 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l2 = s0i32
	s0i32 = l3
	s1i32 = l5
	s2i32 = -64
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l2
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	s0f32 = l8
	s1f32 = 255
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	l8 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l8
	s3f32 = 2.1474835e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l8 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l8
	s3f32 = -2.1474835e+09
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l8 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l8
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl27
	}
	s0i32 = -2147483648
lbl27:
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s1i32 = 255
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 232
		s2i32 = l2
		f52(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = l0
		s2i32 = l1
		f366(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = 20248
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl2
	}
	s0i32 = l3
	s1i32 = 233
	s2i32 = l2
	f52(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = l0
	s2i32 = l1
	f366(ctx, s0i32, s1i32, s2i32)
	goto lbl2
lbl23:
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
	s0i32 = l0
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	s0i32 = f1229(ctx, s0i32, s1i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l1 = s0i32
	s0i32 = l3
	s1i32 = 76
	s2i32 = 4
	s0i32 = f56(ctx, s0i32, s1i32, s2i32)
	l5 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l4 = s0i32
	s0i32 = l3
	s1i32 = l5
	s2i32 = 68
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 234
	s2i32 = l5
	s3i32 = l4
	s2i32 = s2i32 - s3i32
	f52(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = l0
	s2i32 = l1
	s3i32 = l2
	f1228(ctx, s0i32, s1i32, s2i32, s3i32)
	goto lbl2
lbl30:
	s0i32 = l0
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	s2i32 = l1
	s3i32 = l3
	s0i32 = f270(ctx, s0i32, s1i32, s2i32, s3i32)
	l5 = s0i32
	goto lbl2
lbl22:
	s0i32 = l3
	s1i32 = 20
	s2i32 = 4
	s0i32 = f56(ctx, s0i32, s1i32, s2i32)
	l5 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s0i32 = l3
	s1i32 = l5
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 228
	s2i32 = l5
	s3i32 = l0
	s2i32 = s2i32 - s3i32
	f52(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = 20568
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
lbl2:
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = f23(ctx, s0i32)
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
lbl0:
	s0i32 = l6
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l5
	return s0i32
}
