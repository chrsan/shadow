#ifndef SKIA_BINDINGS_H_
#define SKIA_BINDINGS_H_

#include <stddef.h>
#include <stdint.h>
#include <sys/types.h>

#define INIT_STRUCT(x) \
  struct x;            \
  typedef struct x x;

INIT_STRUCT(skia_surface)
INIT_STRUCT(skia_canvas)
INIT_STRUCT(skia_paint)
INIT_STRUCT(skia_path)

#undef INIT_STRUCT

#ifdef __cplusplus
extern "C" {
#endif

// The same order as in core/SkPaint.h
typedef enum {
  PAINT_STYLE_FILL,
  PAINT_STYLE_STROKE,
  PAINT_STYLE_STROKE_AND_FILL,
} PaintStyle;

// The same order as in core/SkPaint.h
typedef enum {
  STROKE_CAP_BUTT,
  STROKE_CAP_ROUND,
  STROKE_CAP_SQUARE,
} StrokeCap;

// The same order as in core/SkPaint.h
typedef enum {
  STROKE_JOIN_MITER,
  STROKE_JOIN_ROUND,
  STROKE_JOIN_BEVEL,
} StrokeJoin;

// The same order as in core/SkPath.h
typedef enum {
  PATH_FILL_TYPE_WINDING,
  PATH_FILL_TYPE_EVEN_ODD,
} PathFillType;

// Matrix

typedef struct {
  float scale_x;
  float skew_x;
  float trans_x;
  float skew_y;
  float scale_y;
  float trans_y;
} skia_matrix;

#ifdef WASM_BUILD
size_t skia_matrix_size();
#endif

skia_matrix* skia_matrix_create();
void skia_matrix_destroy(skia_matrix* matrix);
skia_matrix* skia_matrix_create_from(float scale_x,
                                     float skew_x,
                                     float trans_x,
                                     float skew_y,
                                     float scale_y,
                                     float trans_y);
uint8_t skia_matrix_create_inverse(const skia_matrix* matrix,
                                   skia_matrix* inverse);
void skia_matrix_set(skia_matrix* matrix,
                     float scale_x,
                     float skew_x,
                     float trans_x,
                     float skew_y,
                     float scale_y,
                     float trans_y);
void skia_matrix_reset(skia_matrix* matrix);

// Surface

#ifdef WASM_BUILD
typedef struct {
  uint8_t* ptr;
  uint32_t size;
} skia_surface_data;

size_t skia_surface_data_size();

skia_surface* skia_surface_create_rgba(int width, int height);
skia_surface* skia_surface_create_rgba_premultiplied(int width, int height);
void skia_surface_read_pixels(skia_surface* surface, skia_surface_data* data);
#else
skia_surface* skia_surface_create_rgba(int width, int height, void* pixels);
skia_surface* skia_surface_create_rgba_premultiplied(int width,
                                                     int height,
                                                     void* pixels);
#endif

void skia_surface_destroy(skia_surface* surface);
skia_canvas* skia_surface_get_canvas(skia_surface* surface);
uint8_t skia_is_surface_bgra();

// Canvas

void skia_canvas_clear(skia_canvas* canvas,
                       uint8_t r,
                       uint8_t g,
                       uint8_t b,
                       uint8_t a);
void skia_canvas_flush(skia_canvas* canvas);
void skia_canvas_save(skia_canvas* canvas);
void skia_canvas_restore(skia_canvas* canvas);
void skia_canvas_scale(skia_canvas* canvas, float sx, float sy);
void skia_canvas_translate(skia_canvas* canvas, float dx, float dy);
void skia_canvas_set_matrix(skia_canvas* canvas, const skia_matrix* matrix);
void skia_canvas_concat(skia_canvas* canvas, const skia_matrix* matrix);
void skia_canvas_get_total_matrix(const skia_canvas* canvas,
                                  skia_matrix* matrix);
void skia_canvas_draw_path(skia_canvas* canvas,
                           const skia_path* path,
                           const skia_paint* paint);
void skia_canvas_draw_rect(skia_canvas* canvas,
                           float x,
                           float y,
                           float w,
                           float h,
                           const skia_paint* paint);

// Paint

#ifdef WASM_BUILD
size_t skia_paint_size();
#endif

skia_paint* skia_paint_create();
void skia_paint_destroy(skia_paint* paint);
PaintStyle skia_paint_get_style(const skia_paint* paint);
void skia_paint_set_style(skia_paint* paint, PaintStyle style);
uint32_t skia_paint_get_color(const skia_paint* paint);
void skia_paint_set_color(skia_paint* paint,
                          uint8_t r,
                          uint8_t g,
                          uint8_t b,
                          uint8_t a);
uint8_t skia_paint_get_alpha(const skia_paint* paint);
void skia_paint_set_alpha(skia_paint* paint, uint8_t a);
uint8_t skia_paint_is_anti_alias(const skia_paint* paint);
void skia_paint_set_anti_alias(skia_paint* paint, uint8_t aa);
float skia_paint_get_stroke_width(const skia_paint* paint);
void skia_paint_set_stroke_width(skia_paint* paint, float width);
StrokeCap skia_paint_get_stroke_cap(const skia_paint* paint);
void skia_paint_set_stroke_cap(skia_paint* paint, StrokeCap cap);
StrokeJoin skia_paint_get_stroke_join(const skia_paint* paint);
void skia_paint_set_stroke_join(skia_paint* paint, StrokeJoin join);
float skia_paint_get_stroke_miter(const skia_paint* paint);
void skia_paint_set_stroke_miter(skia_paint* paint, float miter);
void skia_paint_reset(skia_paint* paint);

typedef struct {
  float x;
  float y;
} skia_point;

#ifdef WASM_BUILD
size_t skia_point_size();
#endif

typedef struct {
  float left;
  float top;
  float right;
  float bottom;
} skia_rect;

#ifdef WASM_BUILD
size_t skia_rect_size();
#endif

// Path

#ifdef WASM_BUILD
size_t skia_path_size();
#endif

skia_path* skia_path_create();
void skia_path_destroy(skia_path* path);
PathFillType skia_path_get_fill_type(const skia_path* path);
void skia_path_set_fill_type(skia_path* path, PathFillType type);
void skia_path_move_to(skia_path* path, float x, float y);
void skia_path_line_to(skia_path* path, float x, float y);
void skia_path_quad_to(skia_path* path, float x1, float y1, float x2, float y2);
void skia_path_conic_to(skia_path* path,
                        float x1,
                        float y1,
                        float x2,
                        float y2,
                        float w);
void skia_path_cubic_to(skia_path* path,
                        float x1,
                        float y1,
                        float x2,
                        float y2,
                        float x3,
                        float y3);
void skia_path_close(skia_path* path);
uint8_t skia_path_apply(skia_path* path,
                        const uint8_t* verbs,
                        int num_verbs,
                        const float* values);
void skia_path_reset(skia_path* path);
void skia_path_rewind(skia_path* path);
int skia_path_count_points(const skia_path* path);
skia_point skia_path_get_point(const skia_path* path, int index);
int skia_path_get_points(const skia_path* path, skia_point* points, int max);
int skia_path_count_verbs(const skia_path* path);
int skia_path_get_verbs(const skia_path* path, uint8_t* verbs, int max);
skia_rect skia_path_get_bounds(const skia_path* path);
skia_rect skia_path_compute_tight_bounds(const skia_path* path);

#ifdef __cplusplus
}
#endif

#endif  // SKIA_BINDINGS_H_
