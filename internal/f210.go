package internal

func f210(ctx *Context, l0 float64, l1 int32) float32 {
	var l2 float64
	_ = l2
	var l3 float64
	_ = l3
	var s2i32 int32
	_ = s2i32
	var s0f32 float32
	_ = s0f32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	var s5f64 float64
	_ = s5f64
	var s6f64 float64
	_ = s6f64
	s0f64 = -1
	s1f64 = l0
	s2f64 = l0
	s1f64 = s1f64 * s2f64
	l2 = s1f64
	s2f64 = l0
	s1f64 = s1f64 * s2f64
	l3 = s1f64
	s2f64 = l2
	s3f64 = 0.13339200271297674
	s2f64 = s2f64 * s3f64
	s3f64 = 0.3333313950307914
	s2f64 = s2f64 + s3f64
	s1f64 = s1f64 * s2f64
	s2f64 = l0
	s1f64 = s1f64 + s2f64
	s2f64 = l3
	s3f64 = l2
	s4f64 = l2
	s3f64 = s3f64 * s4f64
	l0 = s3f64
	s2f64 = s2f64 * s3f64
	s3f64 = l2
	s4f64 = 0.024528318116654728
	s3f64 = s3f64 * s4f64
	s4f64 = 0.05338123784456704
	s3f64 = s3f64 + s4f64
	s4f64 = l0
	s5f64 = l2
	s6f64 = 0.009465647849436732
	s5f64 = s5f64 * s6f64
	s6f64 = 0.002974357433599673
	s5f64 = s5f64 + s6f64
	s4f64 = s4f64 * s5f64
	s3f64 = s3f64 + s4f64
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	l0 = s1f64
	s0f64 = s0f64 / s1f64
	s1f64 = l0
	s2i32 = l1
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	s0f32 = float32(s0f64)
	return s0f32
}
