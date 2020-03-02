package internal

func f125(ctx *Context, l0 float64) float32 {
	var l1 float64
	_ = l1
	var l2 float64
	_ = l2
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
	s0f64 = l0
	s1f64 = l0
	s0f64 = s0f64 * s1f64
	l1 = s0f64
	s1f64 = l0
	s0f64 = s0f64 * s1f64
	l2 = s0f64
	s1f64 = l1
	s2f64 = l1
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 * s1f64
	s1f64 = l1
	s2f64 = 2.718311493989822e-06
	s1f64 = s1f64 * s2f64
	s2f64 = -0.00019839334836096632
	s1f64 = s1f64 + s2f64
	s0f64 = s0f64 * s1f64
	s1f64 = l2
	s2f64 = l1
	s3f64 = 0.008333329385889463
	s2f64 = s2f64 * s3f64
	s3f64 = -0.16666666641626524
	s2f64 = s2f64 + s3f64
	s1f64 = s1f64 * s2f64
	s2f64 = l0
	s1f64 = s1f64 + s2f64
	s0f64 = s0f64 + s1f64
	s0f32 = float32(s0f64)
	return s0f32
}
