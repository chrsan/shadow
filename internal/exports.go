package internal

func (c *Context) SkiaCanvasDrawPath(p0 int32, p1 int32, p2 int32) {
	f1311(c, p0, p1, p2)
}

func (c *Context) SkiaPaintIsAntiAlias(p0 int32) int32 {
	return f1246(c, p0)
}

func (c *Context) SkiaPointSize() int32 {
	return f464(c)
}

func (c *Context) SkiaPathLineTo(p0 int32, p1 float32, p2 float32) {
	f1085(c, p0, p1, p2)
}

func (c *Context) StackAlloc(p0 int32) int32 {
	return f1342(c, p0)
}

func (c *Context) SkiaMatrixCreateInverse(p0 int32, p1 int32) int32 {
	return f1116(c, p0, p1)
}

func (c *Context) SkiaPaintSetAlpha(p0 int32, p1 int32) {
	f1256(c, p0, p1)
}

func (c *Context) SkiaPathGetPoint(p0 int32, p1 int32, p2 int32) {
	f1020(c, p0, p1, p2)
}

func (c *Context) SkiaPathComputeTightBounds(p0 int32, p1 int32) {
	f976(c, p0, p1)
}

func (c *Context) SkiaCanvasDrawRect(p0 int32, p1 float32, p2 float32, p3 float32, p4 float32, p5 int32) {
	f1304(c, p0, p1, p2, p3, p4, p5)
}

func (c *Context) SkiaMatrixReset(p0 int32) {
	f967(c, p0)
}

func (c *Context) SkiaPaintSize() int32 {
	return f1298(c)
}

func (c *Context) SkiaPaintSetStrokeWidth(p0 int32, p1 float32) {
	f1199(c, p0, p1)
}

func (c *Context) SkiaMatrixSet(p0 int32, p1 float32, p2 float32, p3 float32, p4 float32, p5 float32, p6 float32) {
	f1046(c, p0, p1, p2, p3, p4, p5, p6)
}

func (c *Context) SkiaSurfaceDestroy(p0 int32) {
	f1911(c, p0)
}

func (c *Context) SkiaCanvasScale(p0 int32, p1 float32, p2 float32) {
	f1398(c, p0, p1, p2)
}

func (c *Context) SkiaCanvasGetTotalMatrix(p0 int32, p1 int32) {
	f1316(c, p0, p1)
}

func (c *Context) SkiaPaintGetColor(p0 int32) int32 {
	return f1278(c, p0)
}

func (c *Context) SkiaPaintGetStyle(p0 int32) int32 {
	return f1224(c, p0)
}

func (c *Context) SkiaPathClose(p0 int32) {
	f1055(c, p0)
}

func (c *Context) SkiaPathGetPoints(p0 int32, p1 int32, p2 int32) int32 {
	return f1010(c, p0, p1, p2)
}

func (c *Context) SkiaSurfaceCreateRgba(p0 int32, p1 int32) int32 {
	return f778(c, p0, p1)
}

func (c *Context) StackSave() int32 {
	return f1343(c)
}

func (c *Context) SkiaPaintGetStrokeJoin(p0 int32) int32 {
	return f1185(c, p0)
}

func (c *Context) SkiaPathApply(p0 int32, p1 int32, p2 int32, p3 int32) int32 {
	return f1049(c, p0, p1, p2, p3)
}

func (c *Context) SkiaPathCountVerbs(p0 int32) int32 {
	return f1000(c, p0)
}

func (c *Context) SkiaMatrixCreate() int32 {
	return f1399(c)
}

func (c *Context) SkiaSurfaceReadPixels(p0 int32, p1 int32) {
	f1976(c, p0, p1)
}

func (c *Context) SkiaIsSurfaceBgra() int32 {
	return f1780(c)
}

func (c *Context) SkiaCanvasSetMatrix(p0 int32, p1 int32) {
	f1332(c, p0, p1)
}

func (c *Context) SkiaSurfaceDataSize() int32 {
	return f464(c)
}

func (c *Context) SkiaCanvasFlush(p0 int32) {
	f1648(c, p0)
}

func (c *Context) SkiaPathSetFillType(p0 int32, p1 int32) {
	f1101(c, p0, p1)
}

func (c *Context) SkiaPathGetVerbs(p0 int32, p1 int32, p2 int32) int32 {
	return f991(c, p0, p1, p2)
}

func (c *Context) SkiaSurfaceCreateRgbaPremultiplied(p0 int32, p1 int32) int32 {
	return f2080(c, p0, p1)
}

func (c *Context) SkiaPaintCreate() int32 {
	return f1294(c)
}

func (c *Context) SkiaPaintGetStrokeWidth(p0 int32) float32 {
	return f1208(c, p0)
}

func (c *Context) SkiaPaintGetStrokeCap(p0 int32) int32 {
	return f1192(c, p0)
}

func (c *Context) SkiaPaintGetStrokeMiter(p0 int32) float32 {
	return f1167(c, p0)
}

func (c *Context) SkiaPathRewind(p0 int32) {
	f1038(c, p0)
}

func (c *Context) StackRestore(p0 int32) {
	f1341(c, p0)
}

func (c *Context) SkiaCanvasTranslate(p0 int32, p1 float32, p2 float32) {
	f1347(c, p0, p1, p2)
}

func (c *Context) SkiaCanvasConcat(p0 int32, p1 int32) {
	f1322(c, p0, p1)
}

func (c *Context) SkiaPaintSetColor(p0 int32, p1 int32, p2 int32, p3 int32, p4 int32) {
	f1273(c, p0, p1, p2, p3, p4)
}

func (c *Context) SkiaPathCountPoints(p0 int32) int32 {
	return f1029(c, p0)
}

func (c *Context) Free(p0 int32) {
	f12(c, p0)
}

func (c *Context) SkiaMatrixDestroy(p0 int32) {
	f13(c, p0)
}

func (c *Context) SkiaCanvasSave(p0 int32) {
	f1571(c, p0)
}

func (c *Context) SkiaCanvasRestore(p0 int32) {
	f1470(c, p0)
}

func (c *Context) SkiaPathSize() int32 {
	return f1133(c)
}

func (c *Context) SkiaPathReset(p0 int32) {
	f1045(c, p0)
}

func (c *Context) SkiaMatrixSize() int32 {
	return f2119(c)
}

func (c *Context) SkiaPaintReset(p0 int32) {
	f1151(c, p0)
}

func (c *Context) SkiaCanvasClear(p0 int32, p1 int32, p2 int32, p3 int32, p4 int32) {
	f1703(c, p0, p1, p2, p3, p4)
}

func (c *Context) SkiaPaintSetAntiAlias(p0 int32, p1 int32) {
	f1235(c, p0, p1)
}

func (c *Context) SkiaPaintSetStrokeMiter(p0 int32, p1 float32) {
	f1159(c, p0, p1)
}

func (c *Context) SkiaPathDestroy(p0 int32) {
	f1115(c, p0)
}

func (c *Context) SkiaPathMoveTo(p0 int32, p1 float32, p2 float32) {
	f1093(c, p0, p1, p2)
}

func (c *Context) SkiaPathCubicTo(p0 int32, p1 float32, p2 float32, p3 float32, p4 float32, p5 float32, p6 float32) {
	f1062(c, p0, p1, p2, p3, p4, p5, p6)
}

func (c *Context) SkiaSurfaceGetCanvas(p0 int32) int32 {
	return f1850(c, p0)
}

func (c *Context) SkiaPaintSetStrokeCap(p0 int32, p1 int32) {
	f1188(c, p0, p1)
}

func (c *Context) SkiaPaintGetAlpha(p0 int32) int32 {
	return f1265(c, p0)
}

func (c *Context) SkiaPathConicTo(p0 int32, p1 float32, p2 float32, p3 float32, p4 float32, p5 float32) {
	f1070(c, p0, p1, p2, p3, p4, p5)
}

func (c *Context) SkiaPathGetBounds(p0 int32, p1 int32) {
	f983(c, p0, p1)
}

func (c *Context) Malloc(p0 int32) int32 {
	return f164(c, p0)
}

func (c *Context) SkiaPathCreate() int32 {
	return f1127(c)
}

func (c *Context) SkiaPathQuadTo(p0 int32, p1 float32, p2 float32, p3 float32, p4 float32) {
	f1076(c, p0, p1, p2, p3, p4)
}

func (c *Context) SkiaRectSize() int32 {
	return f1141(c)
}

func (c *Context) SkiaMatrixCreateFrom(p0 float32, p1 float32, p2 float32, p3 float32, p4 float32, p5 float32) int32 {
	return f1193(c, p0, p1, p2, p3, p4, p5)
}

func (c *Context) SkiaPaintDestroy(p0 int32) {
	f1289(c, p0)
}

func (c *Context) SkiaPaintSetStyle(p0 int32, p1 int32) {
	f1215(c, p0, p1)
}

func (c *Context) SkiaPaintSetStrokeJoin(p0 int32, p1 int32) {
	f1176(c, p0, p1)
}

func (c *Context) SkiaPathGetFillType(p0 int32) int32 {
	return f1106(c, p0)
}

func (c *Context) Start() {
	f451(c)
}
