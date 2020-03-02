package internal

import (
	"math"
)

func f481(ctx *Context, l0 int32, l1 int32, l2 int32, l3 float32) {
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
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	s0f32 = l3
	s1f32 = 6
	s0f32 = s0f32 * s1f32
	s0f32 = float32(math.Ceil(float64(s0f32)))
	l10 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l10
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
	l10 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l10
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
	l10 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l10
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl0
	}
	s0i32 = -2147483648
lbl0:
	l6 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l6
	s1i32 = -2
	s0i32 = s0i32 & s1i32
	s1i32 = l2
	s2i32 = l6
	s1i32 = s1i32 - s2i32
	l7 = s1i32
	s2i32 = -1
	s1i32 = s1i32 ^ s2i32
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s0i32 = l2
	s1i32 = 1
	s0i32 = f50(ctx, s0i32, s1i32)
	l8 = s0i32
	s0i32 = l7
	s0f32 = float32(uint32(s0i32))
	s1f32 = l3
	s2f32 = l3
	s1f32 = s1f32 + s2f32
	l13 = s1f32
	s0f32 = s0f32 / s1f32
	l14 = s0f32
lbl3:
	s0i32 = l0
	s1i32 = l4
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s2i32 = l6
	if uint32(s1i32) >= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l1
		s2i32 = l9
		s3i32 = l4
		s4i32 = 1
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s4i32 = 1
		s3i32 = s3i32 | s4i32
		s4i32 = l2
		s3i32 = s3i32 - s4i32
		l5 = s3i32
		s4i32 = l5
		s5i32 = 31
		s4i32 = s4i32 >> (uint32(s5i32) & 31)
		l5 = s4i32
		s3i32 = s3i32 + s4i32
		s4i32 = l5
		s3i32 = s3i32 ^ s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = 1
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l5 = s2i32
		s3i32 = 0
		s4i32 = l5
		s5i32 = 0
		if s4i32 > s5i32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s1i32 = s1i32 + s2i32
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		goto lbl4
	}
	s1f32 = 0
	l10 = s1f32
	s1f32 = 0
	s2f32 = 1.5
	s3i32 = l4
	s3f32 = float32(uint32(s3i32))
	s4f32 = 0.5
	s3f32 = s3f32 + s4f32
	s4f32 = l13
	s3f32 = s3f32 / s4f32
	s2f32 = s2f32 - s3f32
	l3 = s2f32
	s3f32 = 1.5
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl6
	}
	s1f32 = 1
	s2f32 = l3
	s3f32 = -1.5
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl6
	}
	s1f32 = l3
	s2f32 = l3
	s3f32 = l3
	s2f32 = s2f32 * s3f32
	l11 = s2f32
	s1f32 = s1f32 * s2f32
	l12 = s1f32
	s1f32 = 0.5625
	s2f32 = l3
	s3f32 = 1.125
	s2f32 = s2f32 * s3f32
	s3f32 = l12
	s4f32 = 6
	s3f32 = s3f32 / s4f32
	s4f32 = l11
	s5f32 = -3
	s4f32 = s4f32 * s5f32
	s5f32 = 0.25
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = l3
	s3f32 = 0.5
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3i32 = 1
	s2i32 = s2i32 ^ s3i32
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl6
	}
	s1f32 = l12
	s2f32 = 3
	s1f32 = s1f32 / s2f32
	s2f32 = l3
	s3f32 = -0.75
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s2f32 = l3
	s3f32 = -0.5
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3i32 = 1
	s2i32 = s2i32 ^ s3i32
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl6
	}
	s1f32 = l12
	s2f32 = -6
	s1f32 = s1f32 / s2f32
	s2f32 = l11
	s3f32 = -3
	s2f32 = s2f32 * s3f32
	s3f32 = 0.25
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = l3
	s3f32 = -1.125
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = 0.4375
	s1f32 = s1f32 + s2f32
lbl6:
	l12 = s1f32
	s1f32 = l14
	s2f32 = l3
	s1f32 = s1f32 + s2f32
	l3 = s1f32
	s2f32 = 1.5
	if s1f32 > s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl7
	}
	s1f32 = 1
	l10 = s1f32
	s1f32 = l3
	s2f32 = -1.5
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl7
	}
	s1f32 = l3
	s2f32 = l3
	s3f32 = l3
	s2f32 = s2f32 * s3f32
	l11 = s2f32
	s1f32 = s1f32 * s2f32
	l10 = s1f32
	s1f32 = l3
	s2f32 = 0.5
	if s1f32 > s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = 0.5625
		s2f32 = l3
		s3f32 = 1.125
		s2f32 = s2f32 * s3f32
		s3f32 = l10
		s4f32 = 6
		s3f32 = s3f32 / s4f32
		s4f32 = l11
		s5f32 = -3
		s4f32 = s4f32 * s5f32
		s5f32 = 0.25
		s4f32 = s4f32 * s5f32
		s3f32 = s3f32 + s4f32
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 - s2f32
		l10 = s1f32
		goto lbl7
	}
	s1f32 = l3
	s2f32 = -0.5
	if s1f32 > s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l10
		s2f32 = 3
		s1f32 = s1f32 / s2f32
		s2f32 = l3
		s3f32 = -0.75
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		l10 = s1f32
		goto lbl7
	}
	s1f32 = l10
	s2f32 = -6
	s1f32 = s1f32 / s2f32
	s2f32 = l11
	s3f32 = -3
	s2f32 = s2f32 * s3f32
	s3f32 = 0.25
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = l3
	s3f32 = -1.125
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = 0.4375
	s1f32 = s1f32 + s2f32
	l10 = s1f32
lbl7:
	s1f32 = l12
	s2f32 = l10
	s1f32 = s1f32 - s2f32
	s2f32 = 255
	s1f32 = s1f32 * s2f32
	l3 = s1f32
	s2f32 = 4.2949673e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2f32 = l3
	s3f32 = 0
	if s2f32 >= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s1i32 = s1i32 & s2i32
	if s1i32 != 0 {
		s1f32 = l3
		s1i32 = int32(uint32(math.Trunc(float64(s1f32))))
		goto lbl4
	}
	s1i32 = 0
lbl4:
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l8
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l8
	f13(ctx, s0i32)
lbl2:
}
