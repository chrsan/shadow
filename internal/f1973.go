package internal

import (
	"math"
	"unsafe"
)

func f1973(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l9 int64
	_ = l9
	var l10 int64
	_ = l10
	var l11 float32
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
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = l3
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		s1i32 = 16
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 32767
		s0i32 = s0i32 & s1i32
		l4 = s0i32
		s1i32 = 13
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 939524096
		s0i32 = s0i32 + s1i32
		s1i32 = 0
		s2i32 = l4
		s3i32 = 1023
		if uint32(s2i32) > uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l2
		s2i32 = -2147483648
		s1i32 = s1i32 & s2i32
		s0i32 = s0i32 | s1i32
		l4 = s0i32
		s0i64 = 2147475456
		s1i64 = 0
		s2i32 = l2
		s2i64 = int64(uint32(s2i32))
		l9 = s2i64
		s3i64 = 32767
		s2i64 = s2i64 & s3i64
		l10 = s2i64
		s2i32 = int32(s2i64)
		s3i32 = 1023
		if uint32(s2i32) > uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		s1i64 = l10
		s2i64 = 13
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i64 = 939524096
		s1i64 = s1i64 + s2i64
		s0i64 = s0i64 & s1i64
		s1i64 = l9
		s2i64 = 16
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i64 = 2147483648
		s1i64 = s1i64 & s2i64
		s0i64 = s0i64 | s1i64
		s0i32 = int32(s0i64)
		l6 = s0i32
		s0i32 = 0
		l2 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = l2
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s1f32 = math.Float32frombits(uint32(s1i32))
		s2i64 = 2147475456
		s3i64 = 0
		s4i32 = l1
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
		l5 = s4i32
		s4i64 = int64(uint32(s4i32))
		l9 = s4i64
		s5i64 = 32767
		s4i64 = s4i64 & s5i64
		l10 = s4i64
		s4i32 = int32(s4i64)
		s5i32 = 1023
		if uint32(s4i32) > uint32(s5i32) {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i64 = s2i64
		} else {
			s2i64 = s3i64
		}
		s3i64 = l10
		s4i64 = 13
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 939524096
		s3i64 = s3i64 + s4i64
		s2i64 = s2i64 & s3i64
		s3i64 = l9
		s4i64 = 16
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 2147483648
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s2i32 = int32(s2i64)
		s2f32 = math.Float32frombits(uint32(s2i32))
		l11 = s2f32
		s1f32 = s1f32 + s2f32
		s2f32 = l11
		s1f32 = s1f32 + s2f32
		s2i64 = 2147475456
		s3i64 = 0
		s4i32 = l1
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
		l7 = s4i32
		s4i64 = int64(uint32(s4i32))
		l9 = s4i64
		s5i64 = 32767
		s4i64 = s4i64 & s5i64
		l10 = s4i64
		s4i32 = int32(s4i64)
		s5i32 = 1023
		if uint32(s4i32) > uint32(s5i32) {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i64 = s2i64
		} else {
			s2i64 = s3i64
		}
		s3i64 = l10
		s4i64 = 13
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 939524096
		s3i64 = s3i64 + s4i64
		s2i64 = s2i64 & s3i64
		s3i64 = l9
		s4i64 = 16
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 2147483648
		s3i64 = s3i64 & s4i64
		s2i64 = s2i64 | s3i64
		s2i32 = int32(s2i64)
		l6 = s2i32
		s2f32 = math.Float32frombits(uint32(s2i32))
		s1f32 = s1f32 + s2f32
		s2f32 = 0.25
		s1f32 = s1f32 * s2f32
		s1i32 = int32(math.Float32bits(s1f32))
		l8 = s1i32
		s2i32 = 16
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 32768
		s1i32 = s1i32 & s2i32
		s2i32 = l4
		s2f32 = math.Float32frombits(uint32(s2i32))
		s3i32 = l5
		s4i32 = -2147483648
		s3i32 = s3i32 & s4i32
		s4i32 = l5
		s5i32 = 16
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s5i32 = 32767
		s4i32 = s4i32 & s5i32
		l4 = s4i32
		s5i32 = 13
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s5i32 = 939524096
		s4i32 = s4i32 + s5i32
		s5i32 = 0
		s6i32 = l4
		s7i32 = 1023
		if uint32(s6i32) > uint32(s7i32) {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			// s4i32 = s4i32
		} else {
			s4i32 = s5i32
		}
		s3i32 = s3i32 | s4i32
		s3f32 = math.Float32frombits(uint32(s3i32))
		l11 = s3f32
		s2f32 = s2f32 + s3f32
		s3f32 = l11
		s2f32 = s2f32 + s3f32
		s3i32 = l7
		s4i32 = -2147483648
		s3i32 = s3i32 & s4i32
		s4i32 = l7
		s5i32 = 16
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s5i32 = 32767
		s4i32 = s4i32 & s5i32
		l4 = s4i32
		s5i32 = 13
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s5i32 = 939524096
		s4i32 = s4i32 + s5i32
		s5i32 = 0
		s6i32 = l4
		s7i32 = 1023
		if uint32(s6i32) > uint32(s7i32) {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			// s4i32 = s4i32
		} else {
			s4i32 = s5i32
		}
		s3i32 = s3i32 | s4i32
		l4 = s3i32
		s3f32 = math.Float32frombits(uint32(s3i32))
		s2f32 = s2f32 + s3f32
		s3f32 = 0.25
		s2f32 = s2f32 * s3f32
		s2i32 = int32(math.Float32bits(s2f32))
		l5 = s2i32
		s3i32 = -2147483648
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		s2i32 = l5
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = 1073741824
		s2i32 = s2i32 + s3i32
		s3i32 = -65536
		s2i32 = s2i32 & s3i32
		s3i32 = 0
		s4i32 = l5
		s5i32 = 2147475456
		s4i32 = s4i32 & s5i32
		s5i32 = 947904511
		if uint32(s4i32) > uint32(s5i32) {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s1i32 = s1i32 | s2i32
		s2i32 = l8
		s3i32 = -134217728
		s2i32 = s2i32 - s3i32
		s3i32 = 13
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 65535
		s4i32 = 0
		s5i32 = l8
		s6i32 = 2147475456
		s5i32 = s5i32 & s6i32
		s6i32 = 947904511
		if uint32(s5i32) > uint32(s6i32) {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
