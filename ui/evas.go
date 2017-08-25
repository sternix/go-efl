package ui

/*
#include "bridge.h"
*/
import "C"

import (
	"unsafe"
)

func ColorArgbPremul(r, g, b, a int) (int, int, int, int) {
	var (
		_r = C.int(r)
		_g = C.int(g)
		_b = C.int(b)
		_a = C.int(a)
	)
	C.evas_color_argb_premul(_a, &_r, &_g, &_b)
	return int(_r), int(_g), int(_b), int(_a)
}

func ColorArgbUnpremul(r, g, b, a int) (int, int, int, int) {
	var (
		_r = C.int(r)
		_g = C.int(g)
		_b = C.int(b)
		_a = C.int(a)
	)
	C.evas_color_argb_unpremul(_a, &_r, &_g, &_b)
	return int(_r), int(_g), int(_b), int(_a)
}

func ConvertHsvToRgb(h, s, v float32) (int, int, int) {
	var r, g, b C.int
	C.evas_color_hsv_to_rgb(C.float(h), C.float(s), C.float(v), &r, &g, &b)
	return int(r), int(g), int(b)
}

func ConvertRgbToHsv(r, g, b int) (float32, float32, float32) {
	var h, s, v C.float
	C.evas_color_rgb_to_hsv(C.int(r), C.int(g), C.int(b), &h, &s, &v)
	return float32(h), float32(s), float32(v)
}

type Canvas interface {
	eo() *C.Evas
	AddLine() *EvasLine
	AddBox() *EvasBox
	AddGrid() *EvasGrid
	AddText() *EvasText
	AddImage() *EvasImage
	AddFilledImage() *EvasImage
	AddMap(int) *EvasMap
	AddPolygon() *EvasPolygon
	AddTable() *EvasTable
	AddRectangle() *EvasRectangle
	AddTextblock() *EvasTextblock
	AddTextgrid() *EvasTextgrid
	SetSize(int, int)
	Size() (int, int)
	SetViewport(int, int, int, int)
	Viewport() (int, int, int, int)
	CoordScreenXToWorld(int) int
	CoordScreenYToWorld(int) int
	CoordWorldXToScreen(int) int
	CoordWorldYToScreen(int) int
	PointerOutputXY() (int, int)
	PointerCanvasXY() (int, int)
	PointerButtonDownMask() int
	IsPointerInside() bool
	TopObjectAtXY(int, int, bool, bool) Object
	TopObjectInRectangle(int, int, int, int, bool, bool) Object
	ObjectsAtXY(int, int, bool, bool) []Object
	ObjectsInRectangle(int, int, int, int, bool, bool) []Object
	BottomObject() Object
	TopObject() Object
	FocusIn()
	FocusOut()
	HasFocus() bool
	FocusedObject() Object
	FindObjectByName(string) Object
	PushNoChange()
	PopNoChange()
	AddDamageRect(int, int, int, int)
	AddObscuredRect(int, int, int, int)
	ClearObscured()
	RenderUpdates() []string
	Render()
	NoRender()
	RenderIdleFlush()
	FlushImageCache()
	ReloadImageCache()
	SetImageCacheSize(int)
	ImageCacheSize() int
	FlushFontCache()
	FontCacheSize() int
	SetFontCacheSize(int)
	ClearFontPath()
	AppendToFontPath(string)
	PrependToFontPath(string)
	FontPaths() []string
	AvailableFonts() []string
	IsFontHintingSupported(FontHinting) bool
	SetFontHinting(FontHinting)
	FontHinting() FontHinting
	Freeze()
	Thaw()
	ThawEval()
	FreezeCount() int
	IsKeyModifierSet(string) bool
	IsKeyLockSet(string) bool
	AddKeyModifier(string)
	DelKeyModifier(string)
	AddKeyLock(string)
	DelKeyLock(string)
	SetKeyModifier(string, bool)
	SetKeyLock(string, bool)
	KeyModifierMask(string) uint64
	FeedMouseDown(int, ButtonFlags, int64, interface{})
	FeedMouseUp(int, ButtonFlags, int64, interface{})
	FeedMouseCancel(int64, interface{})
	FeedMouseWheel(int, int, int64, interface{})
	FeedMouseMove(int, int, int64, interface{})
	FeedMouseIn(int64, interface{})
	FeedMouseOut(int64, interface{})
	FeedMultiDown(int, int, int, float64, float64, float64, float64, float64, float64, float64, ButtonFlags, int64, interface{})
	FeedMultiUp(int, int, int, float64, float64, float64, float64, float64, float64, float64, ButtonFlags, int64, interface{})
	FeedMultiMove(int, int, int, float64, float64, float64, float64, float64, float64, float64, int64, interface{})
	FeedKeyDown(string, string, string, string, int64, interface{})
	FeedKeyDownWithKeycode(string, string, string, string, int64, interface{}, uint)
	FeedKeyUp(string, string, string, string, int64, interface{})
	FeedKeyUpWithKeycode(string, string, string, string, int64, interface{}, uint)
	FeedHold(int, int64, interface{})
}

type evasBase struct {
	obj *C.Evas
}

func wrapEvas(o *C.Evas) *evasBase {
	if o != nil {
		return &evasBase{o}
	}
	return nil
}

func (p *evasBase) eo() *C.Evas {
	return p.obj
}

func (p *evasBase) AddLine() *EvasLine {
	return newEvasLine(p.obj)
}

func (p *evasBase) AddBox() *EvasBox {
	return newEvasBox(p.obj)
}

func (p *evasBase) AddGrid() *EvasGrid {
	return newEvasGrid(p.obj)
}

func (p *evasBase) AddText() *EvasText {
	return newEvasText(p.obj)
}

func (p *evasBase) AddImage() *EvasImage {
	return newEvasImage(p.obj)
}

func (p *evasBase) AddFilledImage() *EvasImage {
	return newFilledEvasImage(p.obj)
}

func (p *evasBase) AddMap(count int) *EvasMap {
	return newEvasMap(count)
}

func (p *evasBase) AddPolygon() *EvasPolygon {
	return newEvasPolygon(p.obj)
}

func (p *evasBase) AddTable() *EvasTable {
	return newEvasTable(p.obj)
}

func (p *evasBase) AddRectangle() *EvasRectangle {
	return newEvasRectangle(p.obj)
}

func (p *evasBase) AddTextblock() *EvasTextblock {
	return newEvasTextblock(p.obj)
}

func (p *evasBase) AddTextgrid() *EvasTextgrid {
	return newEvasTextgrid(p.obj)
}

func (p *evasBase) Dispose() {
	C.evas_free(p.obj)
}

/*
func (p *evasBase) SetOutputMethod(eid int) {
	C.evas_output_method_set(p.obj, C.int(eid))
}

func (p *evasBase) SetOutputMethodByName(name string) bool {
	rm := LookupRenderMethod(name)
	if rm < 0 {
		return false
	}
	C.evas_output_method_set(p.obj, C.int(rm))
	return true
}

func (p *evasBase) OutputMethod() int {
	return int(C.evas_output_method_get(p.obj))
}
*/

func (p *evasBase) SetSize(w, h int) {
	C.evas_output_size_set(p.obj, C.int(w), C.int(h))
}

func (p *evasBase) Size() (int, int) {
	var w, h C.int
	C.evas_output_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *evasBase) SetViewport(x, y, w, h int) {
	C.evas_output_viewport_set(p.obj, C.Evas_Coord(x), C.Evas_Coord(y), C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *evasBase) Viewport() (int, int, int, int) {
	var x, y, w, h C.Evas_Coord
	C.evas_output_viewport_get(p.obj, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *evasBase) CoordScreenXToWorld(x int) int {
	return int(C.evas_coord_screen_x_to_world(p.obj, C.int(x)))
}

func (p *evasBase) CoordScreenYToWorld(y int) int {
	return int(C.evas_coord_screen_y_to_world(p.obj, C.int(y)))
}

func (p *evasBase) CoordWorldXToScreen(x int) int {
	return int(C.evas_coord_world_x_to_screen(p.obj, C.int(x)))
}

func (p *evasBase) CoordWorldYToScreen(y int) int {
	return int(C.evas_coord_world_y_to_screen(p.obj, C.int(y)))
}

func (p *evasBase) PointerOutputXY() (int, int) {
	var x, y C.int
	C.evas_pointer_output_xy_get(p.obj, &x, &y)
	return int(x), int(y)
}

func (p *evasBase) PointerCanvasXY() (int, int) {
	var x, y C.int
	C.evas_pointer_canvas_xy_get(p.obj, &x, &y)
	return int(x), int(y)
}

func (p *evasBase) PointerButtonDownMask() int {
	return int(C.evas_pointer_button_down_mask_get(p.obj))
}

func (p *evasBase) IsPointerInside() bool {
	return C.evas_pointer_inside_get(p.obj) == eTrue
}

func (p *evasBase) TopObjectAtXY(x int, y int, peo bool, hid bool) Object {
	return wrapObjectBase(C.evas_object_top_at_xy_get(p.obj, C.int(x), C.int(y), eBool(peo), eBool(hid)))
}

func (p *evasBase) TopObjectInRectangle(x int, y int, w int, h int, ipeo bool, hid bool) Object {
	return wrapObjectBase(C.evas_object_top_in_rectangle_get(p.obj, C.int(x), C.int(y), C.int(w), C.int(h), eBool(ipeo), eBool(hid)))
}

func (p *evasBase) ObjectsAtXY(x int, y int, ipeo bool, hid bool) []Object {
	lst := C.evas_objects_at_xy_get(p.obj, C.int(x), C.int(y), eBool(ipeo), eBool(hid))
	return newListIterator(lst).ObjectSlice()
}

func (p *evasBase) ObjectsInRectangle(x int, y int, w int, h int, ipeo bool, hid bool) []Object {
	lst := C.evas_objects_in_rectangle_get(p.obj, C.int(x), C.int(y), C.int(w), C.int(h), eBool(ipeo), eBool(hid))
	return newListIterator(lst).ObjectSlice()
}

func (p *evasBase) BottomObject() Object {
	return wrapObjectBase(C.evas_object_bottom_get(p.obj))
}

func (p *evasBase) TopObject() Object {
	return wrapObjectBase(C.evas_object_top_get(p.obj))
}

/*
//TODO is this required
func (p *evasBase) AttachedData() unsafe.Pointer {
	return unsafe.Pointer(C.evas_data_attach_get(p.obj))
}

func (p *evasBase) AttachData(data unsafe.Pointer) {
	C.evas_data_attach_set(p.obj, data)
}
*/

func (p *evasBase) FocusIn() {
	C.evas_focus_in(p.obj)
}

func (p *evasBase) FocusOut() {
	C.evas_focus_out(p.obj)
}

func (p *evasBase) HasFocus() bool {
	return C.evas_focus_state_get(p.obj) == eTrue
}

func (p *evasBase) FocusedObject() Object {
	return wrapObjectBase(C.evas_focus_get(p.obj))
}

func (p *evasBase) FindObjectByName(name string) Object {
	cname := C.CString(name)
	defer free(cname)
	return wrapObjectBase(C.evas_object_name_find(p.obj, cname))
}

func (p *evasBase) PushNoChange() {
	C.evas_nochange_push(p.obj)
}

func (p *evasBase) PopNoChange() {
	C.evas_nochange_pop(p.obj)
}

func (p *evasBase) AddDamageRect(x, y, w, h int) {
	C.evas_damage_rectangle_add(p.obj, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (p *evasBase) AddObscuredRect(x, y, w, h int) {
	C.evas_obscured_rectangle_add(p.obj, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (p *evasBase) ClearObscured() {
	C.evas_obscured_clear(p.obj)
}

func (p *evasBase) RenderUpdates() []string {
	lst := C.evas_render_updates(p.obj)
	return newListIterator(lst).StringSlice()
}

func (p *evasBase) Render() {
	C.evas_render(p.obj)
}

func (p *evasBase) NoRender() {
	C.evas_norender(p.obj)
}

func (p *evasBase) RenderIdleFlush() {
	C.evas_render_idle_flush(p.obj)
}

func (p *evasBase) FlushImageCache() {
	C.evas_image_cache_flush(p.obj)
}

func (p *evasBase) ReloadImageCache() {
	C.evas_image_cache_reload(p.obj)
}

func (p *evasBase) SetImageCacheSize(size int) {
	C.evas_image_cache_set(p.obj, C.int(size))
}

func (p *evasBase) ImageCacheSize() int {
	return int(C.evas_image_cache_get(p.obj))
}

func (p *evasBase) FlushFontCache() {
	C.evas_font_cache_flush(p.obj)
}

func (p *evasBase) FontCacheSize() int {
	return int(C.evas_font_cache_get(p.obj))
}

func (p *evasBase) SetFontCacheSize(size int) {
	C.evas_font_cache_set(p.obj, C.int(size))
}

func (p *evasBase) ClearFontPath() {
	C.evas_font_path_clear(p.obj)
}

func (p *evasBase) AppendToFontPath(path string) {
	cpath := C.CString(path)
	defer free(cpath)
	C.evas_font_path_append(p.obj, cpath)
}

func (p *evasBase) PrependToFontPath(path string) {
	cpath := C.CString(path)
	defer free(cpath)
	C.evas_font_path_prepend(p.obj, cpath)
}

func (p *evasBase) FontPaths() []string {
	return newListIterator(C.evas_font_path_list(p.obj)).StringSlice()
}

func (p *evasBase) AvailableFonts() []string {
	lst := C.evas_font_available_list(p.obj)
	return newListIterator(lst).StringSlice()
}

func (p *evasBase) IsFontHintingSupported(flags FontHinting) bool {
	return C.evas_font_hinting_can_hint(p.obj, C.Evas_Font_Hinting_Flags(flags)) == eTrue
}

func (p *evasBase) SetFontHinting(flags FontHinting) {
	C.evas_font_hinting_set(p.obj, C.Evas_Font_Hinting_Flags(flags))
}

func (p *evasBase) FontHinting() FontHinting {
	return FontHinting(C.evas_font_hinting_get(p.obj))
}

func (p *evasBase) Freeze() {
	C.evas_event_freeze(p.obj)
}

func (p *evasBase) Thaw() {
	C.evas_event_thaw(p.obj)
}

func (p *evasBase) ThawEval() {
	C.evas_event_thaw_eval(p.obj)
}

func (p *evasBase) FreezeCount() int {
	return int(C.evas_event_freeze_get(p.obj))
}

func (p *evasBase) IsKeyModifierSet(keyname string) bool {
	ckeyname := C.CString(keyname)
	defer free(ckeyname)
	mods := C.evas_key_modifier_get(p.obj)
	return C.evas_key_modifier_is_set(mods, ckeyname) == eTrue
}

func (p *evasBase) IsKeyLockSet(keyname string) bool {
	ckeyname := C.CString(keyname)
	defer free(ckeyname)
	lock := C.evas_key_lock_get(p.obj)
	return C.evas_key_lock_is_set(lock, ckeyname) == eTrue
}

func (p *evasBase) AddKeyModifier(keyname string) {
	ckeyname := C.CString(keyname)
	defer free(ckeyname)
	C.evas_key_modifier_add(p.obj, ckeyname)
}

func (p *evasBase) DelKeyModifier(keyname string) {
	ckeyname := C.CString(keyname)
	defer free(ckeyname)
	C.evas_key_modifier_del(p.obj, ckeyname)
}

func (p *evasBase) AddKeyLock(keyname string) {
	ckeyname := C.CString(keyname)
	defer free(ckeyname)
	C.evas_key_lock_add(p.obj, ckeyname)
}

func (p *evasBase) DelKeyLock(keyname string) {
	ckeyname := C.CString(keyname)
	defer free(ckeyname)
	C.evas_key_lock_del(p.obj, ckeyname)
}

func (p *evasBase) SetKeyModifier(keyname string, on bool) {
	ckeyname := C.CString(keyname)
	defer free(ckeyname)
	if on {
		C.evas_key_modifier_on(p.obj, ckeyname)
	} else {
		C.evas_key_modifier_off(p.obj, ckeyname)
	}
}

func (p *evasBase) SetKeyLock(keyname string, on bool) {
	ckeyname := C.CString(keyname)
	defer free(ckeyname)
	if on {
		C.evas_key_lock_on(p.obj, ckeyname)
	} else {
		C.evas_key_lock_off(p.obj, ckeyname)
	}
}

func (p *evasBase) KeyModifierMask(keyname string) uint64 {
	ckeyname := C.CString(keyname)
	defer free(ckeyname)
	return uint64(C.evas_key_modifier_mask_get(p.obj, ckeyname))
}

/* TODO
void 	evas_seat_key_lock_on (Evas_Canvas *obj, const char *keyname, Efl_Input_Device *seat);
void 	evas_seat_key_lock_off (Evas_Canvas *obj, const char *keyname, Efl_Input_Device *seat);
Eina_Bool evas_seat_key_lock_is_set (const Evas_Lock * l,const char * keyname,const Evas_Device *seat);
void *evas_event_callback_del_full (Evas *e, Evas_Callback_Type type, Evas_Event_Cb func, const void *data)
void *evas_event_callback_del (Evas *e, Evas_Callback_Type type, Evas_Event_Cb func)
void  evas_post_event_callback_push (Evas *e, Evas_Object_Event_Post_Cb func, const void *data)
void  evas_post_event_callback_remove (Evas *e, Evas_Object_Event_Post_Cb func)
void  evas_post_event_callback_remove_full (Evas *e, Evas_Object_Event_Post_Cb func, const void *data)
*/

func (p *evasBase) FeedMouseDown(button int, flag ButtonFlags, ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_mouse_down(p.obj, C.int(button), C.Evas_Button_Flags(flag), C.uint(ts), ptr)
}

func (p *evasBase) FeedMouseUp(button int, flag ButtonFlags, ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_mouse_up(p.obj, C.int(button), C.Evas_Button_Flags(flag), C.uint(ts), ptr)
}

func (p *evasBase) FeedMouseCancel(ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_mouse_cancel(p.obj, C.uint(ts), ptr)
}

func (p *evasBase) FeedMouseWheel(direction int, z int, ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_mouse_wheel(p.obj, C.int(direction), C.int(z), C.uint(ts), ptr)
}

func (p *evasBase) FeedMouseMove(x int, y int, ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_mouse_move(p.obj, C.int(x), C.int(y), C.uint(ts), ptr)
}

func (p *evasBase) FeedMouseIn(ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_mouse_in(p.obj, C.uint(ts), ptr)
}

func (p *evasBase) FeedMouseOut(ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_mouse_out(p.obj, C.uint(ts), ptr)
}

func (p *evasBase) FeedMultiDown(d int, x int, y int, r float64, rx float64, ry float64, press float64, ang float64, fx float64, fy float64, flag ButtonFlags, ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_multi_down(p.obj, C.int(d), C.int(x), C.int(y), C.double(r), C.double(rx), C.double(ry), C.double(press), C.double(ang), C.double(fx), C.double(fy), C.Evas_Button_Flags(flag), C.uint(ts), ptr)
}

func (p *evasBase) FeedMultiUp(d int, x int, y int, r float64, rx float64, ry float64, press float64, ang float64, fx float64, fy float64, flag ButtonFlags, ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_multi_up(p.obj, C.int(d), C.int(x), C.int(y), C.double(r), C.double(rx), C.double(ry), C.double(press), C.double(ang), C.double(fx), C.double(fy), C.Evas_Button_Flags(flag), C.uint(ts), ptr)
}

func (p *evasBase) FeedMultiMove(d int, x int, y int, r float64, rx float64, ry float64, press float64, ang float64, fx float64, fy float64, ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_multi_move(p.obj, C.int(d), C.int(x), C.int(y), C.double(r), C.double(rx), C.double(ry), C.double(press), C.double(ang), C.double(fx), C.double(fy), C.uint(ts), ptr)
}

func (p *evasBase) FeedKeyDown(keyname string, key string, str string, compose string, ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	ckeyname := C.CString(keyname)
	ckey := C.CString(key)
	cstr := C.CString(str)
	ccompose := C.CString(compose)
	cts := C.uint(ts)
	defer free(ckeyname, ckey, cstr, ccompose)
	C.evas_event_feed_key_down(p.obj, ckeyname, ckey, cstr, ccompose, cts, ptr)
}

func (p *evasBase) FeedKeyDownWithKeycode(keyname string, key string, str string, compose string, ts int64, data interface{}, keycode uint) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	ckeyname := C.CString(keyname)
	ckey := C.CString(key)
	cstr := C.CString(str)
	ccompose := C.CString(compose)
	cts := C.uint(ts)
	ckeycode := C.uint(keycode)
	defer free(ckeyname, ckey, cstr, ccompose)
	C.evas_event_feed_key_down_with_keycode(p.obj, ckeyname, ckey, cstr, ccompose, cts, ptr, ckeycode)
}

func (p *evasBase) FeedKeyUp(keyname string, key string, str string, compose string, ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	ckeyname := C.CString(keyname)
	ckey := C.CString(key)
	cstr := C.CString(str)
	ccompose := C.CString(compose)
	cts := C.uint(ts)
	defer free(ckeyname, ckey, cstr, ccompose)
	C.evas_event_feed_key_up(p.obj, ckeyname, ckey, cstr, ccompose, cts, ptr)
}

func (p *evasBase) FeedKeyUpWithKeycode(keyname string, key string, str string, compose string, ts int64, data interface{}, keycode uint) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	ckeyname := C.CString(keyname)
	ckey := C.CString(key)
	cstr := C.CString(str)
	ccompose := C.CString(compose)
	cts := C.uint(ts)
	ckeycode := C.uint(keycode)
	defer free(ckeyname, ckey, cstr, ccompose)
	C.evas_event_feed_key_up_with_keycode(p.obj, ckeyname, ckey, cstr, ccompose, cts, ptr, ckeycode)
}

func (p *evasBase) FeedHold(hold int, ts int64, data interface{}) {
	var ptr unsafe.Pointer
	if data != nil {
		id := registry.Register(data)
		ptr = unsafe.Pointer(&id)
	}
	C.evas_event_feed_hold(p.obj, C.int(hold), C.uint(ts), ptr)
}

/*
TODO
void evas_event_feed_axis_update (Evas *obj, unsigned int timestamp, int device, int toolid, int naxes, const Evas_Axis *axis, const void *data)
void evas_event_refeed_event (Evas *obj, void *event_copy, Evas_Callback_Type event_type)
void 	evas_event_input_mouse_move (Evas *obj, int x, int y, unsigned int timestamp, const void *data)
void 	evas_event_input_multi_down (Evas *obj, int d, int x, int y, double rad, double radx, double rady, double pres, double ang, double fx, double fy, Evas_Button_Flags flags, unsigned int timestamp, const void *data)
void 	evas_event_input_multi_move (Evas *obj, int d, int x, int y, double rad, double radx, double rady, double pres, double ang, double fx, double fy, unsigned int timestamp, const void *data)
void 	evas_event_input_multi_up (Evas *obj, int d, int x, int y, double rad, double radx, double rady, double pres, double ang, double fx, double fy, Evas_Button_Flags flags, unsigned int timestamp, const void *data)
*/

//-------8<-------------------------------------------------------

type EvasBox struct {
	*objectBase
}

func newEvasBox(evas *C.Evas) *EvasBox {
	eo := C.evas_object_box_add(evas)
	return &EvasBox{wrapObjectBase(eo)}
}

func (p *EvasBox) Padding() (int, int) {
	var h, v C.int
	C.evas_object_box_padding_get(p.obj, &h, &v)
	return int(h), int(v)
}

func (p *EvasBox) SetPadding(h, v int) {
	C.evas_object_box_padding_set(p.obj, C.int(h), C.int(v))
}

func (p *EvasBox) Align() (float64, float64) {
	var h, v C.double
	C.evas_object_box_align_get(p.obj, &h, &v)
	return float64(h), float64(v)
}

func (p *EvasBox) SetAlign(h, v float64) {
	C.evas_object_box_align_set(p.obj, C.double(h), C.double(v))
}

func (p *EvasBox) Append(child Object) bool {
	return C.evas_object_box_append(p.obj, child.eo()) != nil
}

func (p *EvasBox) Prepend(child Object) bool {
	return C.evas_object_box_prepend(p.obj, child.eo()) != nil
}

func (p *EvasBox) InsertBefore(child Object, reference Object) bool {
	return C.evas_object_box_insert_before(p.obj, child.eo(), reference.eo()) != nil
}

func (p *EvasBox) InsertAfter(child Object, reference Object) bool {
	return C.evas_object_box_insert_after(p.obj, child.eo(), reference.eo()) != nil
}

func (p *EvasBox) InsertAt(child Object, pos uint) bool {
	return C.evas_object_box_insert_at(p.obj, child.eo(), C.uint(pos)) != nil
}

func (p *EvasBox) Remove(child Object) bool {
	return C.evas_object_box_remove(p.obj, child.eo()) == eTrue
}

func (p *EvasBox) RemoveAt(pos uint) bool {
	return C.evas_object_box_remove_at(p.obj, C.uint(pos)) == eTrue
}

func (p *EvasBox) RemoveAll(del bool) bool {
	return C.evas_object_box_remove_all(p.obj, eBool(del)) == eTrue
}

func (p *EvasBox) Children() []Object {
	lst := C.evas_object_box_children_get(p.obj)
	return newListIterator(lst).ObjectSlice()
}

//-------8<-------------------------------------------------------

type EvasGrid struct {
	*objectBase
}

func newEvasGrid(evas *C.Evas) *EvasGrid {
	eo := C.evas_object_grid_add(evas)
	return &EvasGrid{wrapObjectBase(eo)}
}

//evas_object_grid_add_to

func (p *EvasGrid) Size() (int, int) {
	var w, h C.int
	C.evas_object_grid_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *EvasGrid) SetSize(w, h int) {
	C.evas_object_grid_size_set(p.obj, C.int(w), C.int(h))
}

func (p *EvasGrid) IsMirrored() bool {
	return C.evas_object_grid_mirrored_get(p.obj) == eTrue
}

func (p *EvasGrid) SetMirrored(m bool) {
	C.evas_object_grid_mirrored_set(p.obj, eBool(m))
}

func (p *EvasGrid) Pack(child Object, x int, y int, w int, h int) bool {
	return C.evas_object_grid_pack(p.obj, child.eo(), C.int(x), C.int(y), C.int(w), C.int(h)) == eTrue
}

func (p *EvasGrid) Unpack(child Object) {
	C.evas_object_grid_unpack(p.obj, child.eo())
}

func (p *EvasGrid) PackLocation(child Object) (int, int, int, int, bool) {
	var x, y, w, h C.int
	b := C.evas_object_grid_pack_get(p.obj, child.eo(), &x, &y, &w, &h)
	if b == eFalse {
		return 0, 0, 0, 0, false
	}
	return int(x), int(y), int(w), int(h), true
}

func (p *EvasGrid) Clear(del bool) {
	C.evas_object_grid_clear(p.obj, eBool(del))
}

func (p *EvasGrid) Children() []Object {
	lst := C.evas_object_grid_children_get(p.obj)
	return newListIterator(lst).ObjectSlice()
}

//-------8<-------------------------------------------------------

type EvasImage struct {
	*objectBase
}

func IsImageFormatSupported(fileName string) bool {
	cfileName := C.CString(fileName)
	defer free(cfileName)
	return C.evas_object_image_extension_can_load_get(cfileName) == eTrue
}

func newEvasImage(evas *C.Evas) *EvasImage {
	eo := C.evas_object_image_add(evas)
	return &EvasImage{wrapObjectBase(eo)}
}

func newFilledEvasImage(evas *C.Evas) *EvasImage {
	eo := C.evas_object_image_filled_add(evas)
	return &EvasImage{wrapObjectBase(eo)}
}

func (p *EvasImage) SetFile(file, key string) {
	var ckey *C.char
	cfile := C.CString(file)
	defer free(cfile)
	if key != "" {
		ckey = C.CString(key)
		defer free(ckey)
	}
	C.evas_object_image_file_set(p.obj, cfile, ckey)
}

func (p *EvasImage) File() (string, string) {
	var cfile, ckey *C.char
	C.evas_object_image_file_get(p.obj, &cfile, &ckey)
	return C.GoString(cfile), C.GoString(ckey)
}

func (p *EvasImage) Border() (int, int, int, int) {
	var l, r, t, b C.int
	C.evas_object_image_border_get(p.obj, &l, &r, &t, &b)
	return int(l), int(r), int(t), int(b)
}

func (p *EvasImage) SetBorder(l, r, t, b int) {
	C.evas_object_image_border_set(p.obj, C.int(l), C.int(r), C.int(t), C.int(b))
}

func (p *EvasImage) BorderCenterFill() BorderFillMode {
	return BorderFillMode(C.evas_object_image_border_center_fill_get(p.obj))
}

func (p *EvasImage) SetBorderCenterFill(mode BorderFillMode) {
	C.evas_object_image_border_center_fill_set(p.obj, C.Evas_Border_Fill_Mode(mode))
}

func (p *EvasImage) IsFilled() bool {
	return C.evas_object_image_filled_get(p.obj) == eTrue
}

func (p *EvasImage) SetFilled(b bool) {
	C.evas_object_image_filled_set(p.obj, eBool(b))
}

func (p *EvasImage) Fill() (int, int, int, int) {
	var x, y, w, h C.Evas_Coord
	C.evas_object_image_fill_get(p.obj, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *EvasImage) SetFill(x, y, w, h int) {
	C.evas_object_image_fill_set(p.obj, C.Evas_Coord(x), C.Evas_Coord(y), C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *EvasImage) BorderScale() float64 {
	return float64(C.evas_object_image_border_scale_get(p.obj))
}

func (p *EvasImage) SetBorderScale(scale float64) {
	C.evas_object_image_border_scale_set(p.obj, C.double(scale))
}

func (p *EvasImage) ImageSize() (int, int) {
	var w, h C.int
	C.evas_object_image_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *EvasImage) SetImageSize(w, h int) {
	C.evas_object_image_size_set(p.obj, C.int(w), C.int(h))
}

func (p *EvasImage) Stride() int {
	return int(C.evas_object_image_stride_get(p.obj))
}

func (p *EvasImage) Alpha() bool {
	return C.evas_object_image_alpha_get(p.obj) == eTrue
}

func (p *EvasImage) SetAlpha(b bool) {
	C.evas_object_image_alpha_set(p.obj, eBool(b))
}

func (p *EvasImage) SmoothScale() bool {
	return C.evas_object_image_smooth_scale_get(p.obj) == eTrue
}

func (p *EvasImage) SetSmoothScale(b bool) {
	C.evas_object_image_smooth_scale_set(p.obj, eBool(b))
}

func (p *EvasImage) Preload(cancel bool) {
	C.evas_object_image_preload(p.obj, eBool(cancel))
}

/*
flags parameter like "quality=100 compress=8"
*/
func (p *EvasImage) Save(file, key, flags string) bool {
	var ckey *C.char

	cfile := C.CString(file)
	if key != "" {
		ckey = C.CString(key)
		defer free(ckey)
	}
	cflags := C.CString(flags)
	defer free(cfile, cflags)
	return C.evas_object_image_save(p.obj, cfile, ckey, cflags) == eTrue
}

func (p *EvasImage) IsPixelsDirty() bool {
	return C.evas_object_image_pixels_dirty_get(p.obj) == eTrue
}

func (p *EvasImage) SetPixelsDirty(b bool) {
	C.evas_object_image_pixels_dirty_set(p.obj, eBool(b))
}

type imageLoad struct {
	*EvasImage
}

func (p *EvasImage) Load() *imageLoad {
	return &imageLoad{
		EvasImage: p,
	}
}

func (p *imageLoad) Error() error {
	le := LoadError(C.evas_object_image_load_error_get(p.obj))
	if le == LoadErrorNone {
		return nil
	}
	return le
}

func (p *imageLoad) SetDpi(dpi float64) {
	C.evas_object_image_load_dpi_set(p.obj, C.double(dpi))
}

func (p *imageLoad) Dpi() float64 {
	return float64(C.evas_object_image_load_dpi_get(p.obj))
}

func (p *imageLoad) Size() (int, int) {
	var w, h C.int
	C.evas_object_image_load_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *imageLoad) SetSize(w, h int) {
	C.evas_object_image_load_size_set(p.obj, C.int(w), C.int(h))
}

func (p *imageLoad) ScaleDown() int {
	return int(C.evas_object_image_load_scale_down_get(p.obj))
}

func (p *imageLoad) SetScaleDown(sd int) {
	C.evas_object_image_load_scale_down_set(p.obj, C.int(sd))
}

func (p *imageLoad) Region() (int, int, int, int) {
	var x, y, w, h C.int
	C.evas_object_image_load_region_get(p.obj, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *imageLoad) SetRegion(x, y, w, h int) {
	C.evas_object_image_load_region_set(p.obj, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (p *imageLoad) Orientation() bool {
	return C.evas_object_image_load_orientation_get(p.obj) == eTrue
}

func (p *imageLoad) SetOrientation(b bool) {
	C.evas_object_image_load_orientation_set(p.obj, eBool(b))
}

func (p *EvasImage) Colorspace() Colorspace {
	return Colorspace(C.evas_object_image_colorspace_get(p.obj))
}

func (p *EvasImage) SetColorspace(cs Colorspace) {
	C.evas_object_image_colorspace_set(p.obj, C.Evas_Colorspace(cs))
}

func (p *EvasImage) IsRegionSupported() bool {
	return C.evas_object_image_region_support_get(p.obj) == eTrue
}

func (p *EvasImage) SetScaleHint(hint ImageScaleHint) {
	C.evas_object_image_scale_hint_set(p.obj, C.Evas_Image_Scale_Hint(hint))
}

func (p *EvasImage) ScaleHint() ImageScaleHint {
	return ImageScaleHint(C.evas_object_image_scale_hint_get(p.obj))
}

func (p *EvasImage) SetContentHint(hint ImageContentHint) {
	C.evas_object_image_content_hint_set(p.obj, C.Evas_Image_Content_Hint(hint))
}

func (p *EvasImage) ContentHint() ImageContentHint {
	return ImageContentHint(C.evas_object_image_content_hint_get(p.obj))
}

func (p *EvasImage) Source() Object {
	return wrapObjectBase(C.evas_object_image_source_get(p.obj))
}

func (p *EvasImage) SetSource(src Object) bool {
	return C.evas_object_image_source_set(p.obj, src.eo()) == eTrue
}

func (p *EvasImage) UnsetSource() bool {
	return C.evas_object_image_source_unset(p.obj) == eTrue
}

func (p *EvasImage) SetSourceVisible(b bool) {
	C.evas_object_image_source_visible_set(p.obj, eBool(b))
}

func (p *EvasImage) IsSourceVisible() bool {
	return C.evas_object_image_source_visible_get(p.obj) == eTrue
}

func (p *EvasImage) SetSourceEvents(b bool) {
	C.evas_object_image_source_events_set(p.obj, eBool(b))
}

func (p *EvasImage) IsSourceEvents() bool {
	return C.evas_object_image_source_events_get(p.obj) == eTrue
}

func (p *EvasImage) SourceClip() bool {
	return C.evas_object_image_source_clip_get(p.obj) == eTrue
}

func (p *EvasImage) SetSourceClip(b bool) {
	C.evas_object_image_source_clip_set(p.obj, eBool(b))
}

func (p *EvasImage) IsAnimated() bool {
	return C.evas_object_image_animated_get(p.obj) == eTrue
}

/*
returned data is C data, if you manipulate slice you changed the C data
*/
func (p *EvasImage) ImageData(for_writing bool) []uint {
	vptr := C.evas_object_image_data_get(p.obj, eBool(for_writing))
	_, h := p.ImageSize()
	length := h * p.Stride() / int(unsafe.Sizeof(int(0)))
	return (*[1 << 30]uint)(unsafe.Pointer(vptr))[:length:length]
}

func (p *EvasImage) AddUpdateData(x, y, w, h int) {
	C.evas_object_image_data_update_add(p.obj, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (p *EvasImage) SetImageData(data []uint) {
	C.evas_object_image_data_set(p.obj, unsafe.Pointer(&data[0]))
}

type animatedImage struct {
	*EvasImage
}

func (p *EvasImage) Animated() (*animatedImage, bool) {
	if p.IsAnimated() {
		return &animatedImage{
			EvasImage: p,
		}, true
	}

	return nil, false
}

func (p *animatedImage) FrameCount() int {
	return int(C.evas_object_image_animated_get(p.obj))
}

func (p *animatedImage) LoopType() ImageAnimatedLoopHint {
	return ImageAnimatedLoopHint(C.evas_object_image_animated_loop_type_get(p.obj))
}

func (p *animatedImage) LoopCount() int {
	return int(C.evas_object_image_animated_loop_count_get(p.obj))
}

func (p *animatedImage) FrameDuration(startFrame, frameNumber int) float64 {
	return float64(C.evas_object_image_animated_frame_duration_get(p.obj, C.int(startFrame), C.int(frameNumber)))
}

func (p *animatedImage) SetFrame(frameNumber int) {
	C.evas_object_image_animated_frame_set(p.obj, C.int(frameNumber))
}

/*
TODO
void 	evas_object_image_memfile_set (Evas_Object *obj, void *data, int size, char *format, char *key)
 	Sets the data for an image from memory to be loaded.

void 	evas_object_image_data_set (Evas_Object *obj, void *data)
 	Sets the raw image data of the given image object.

void * 	evas_object_image_data_get (const Evas_Object *obj, Eina_Bool for_writing)
 	Gets a pointer to the raw image data of the given image object.

void * 	evas_object_image_data_convert (Evas_Object *obj, Evas_Colorspace to_cspace)
 	Converts the raw image data of the given image object to the specified colorspace.

void 	evas_object_image_data_copy_set (Evas_Object *obj, void *data)
 	Replaces the raw image data of the given image object.

Eina_Bool 	evas_object_image_pixels_import (Evas_Object *obj, Evas_Pixel_Import_Source *pixels)
 	Imports pixels from given source to a given canvas image object.

*/

//-------8<-------------------------------------------------------

type EvasLine struct {
	*objectBase
}

func newEvasLine(evas *C.Evas) *EvasLine {
	eo := C.evas_object_line_add(evas)
	return &EvasLine{
		wrapObjectBase(eo),
	}
}

func (p *EvasLine) XY() (int, int, int, int) {
	var x1, y1, x2, y2 C.int
	C.evas_object_line_xy_get(p.obj, &x1, &y1, &x2, &y2)
	return int(x1), int(y1), int(x2), int(y2)
}

func (p *EvasLine) SetXY(x1, y1, x2, y2 int) {
	C.evas_object_line_xy_set(p.obj, C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

//-------8<-------------------------------------------------------

type EvasMap struct {
	emap *C.Evas_Map
}

func newEvasMap(count int) *EvasMap {
	emap := C.evas_map_new(C.int(count))
	return &EvasMap{
		emap: emap,
	}
}

func (p *EvasMap) Dup() *EvasMap {
	em := C.evas_map_dup(p.emap)
	return &EvasMap{
		emap: em,
	}
}

func (p *EvasMap) Dispose() {
	C.evas_map_free(p.emap)
}

func (p *EvasMap) PopulateFromObjectFull(obj Object, z int) {
	C.evas_map_util_points_populate_from_object_full(p.emap, obj.eo(), C.Evas_Coord(z))
}

func (p *EvasMap) PopulateFromObject(obj Object) {
	C.evas_map_util_points_populate_from_object(p.emap, obj.eo())
}

func (p *EvasMap) PopulateFromGeometry(x, y, w, h, z int) {
	C.evas_map_util_points_populate_from_geometry(p.emap, C.Evas_Coord(x), C.Evas_Coord(y), C.Evas_Coord(w), C.Evas_Coord(h), C.Evas_Coord(z))
}

func (p *EvasMap) Rotate(degrees float64, cx, cy int) {
	C.evas_map_util_rotate(p.emap, C.double(degrees), C.Evas_Coord(cx), C.Evas_Coord(cy))
}

func (p *EvasMap) Zoom(zoomx, zoomy float64, cx, cy int) {
	C.evas_map_util_zoom(p.emap, C.double(zoomx), C.double(zoomy), C.Evas_Coord(cx), C.Evas_Coord(cy))
}

func (p *EvasMap) SetPointsColor(r, g, b, a int) {
	C.evas_map_util_points_color_set(p.emap, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *EvasMap) SetPointColor(idx int, r, g, b, a int) {
	C.evas_map_point_color_set(p.emap, C.int(idx), C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *EvasMap) PointColor(idx int) (int, int, int, int) {
	var r, g, b, a C.int
	C.evas_map_point_color_get(p.emap, C.int(idx), &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *EvasMap) Rotate3D(dx, dy, dz float64, cx, cy, cz int) {
	C.evas_map_util_3d_rotate(p.emap, C.double(dx), C.double(dy), C.double(dz), C.Evas_Coord(cx), C.Evas_Coord(cy), C.Evas_Coord(cz))
}

func (p *EvasMap) QuatRotate(qx, qy, qz, qw, cx, cy, cz float64) {
	C.evas_map_util_quat_rotate(p.emap, C.double(qx), C.double(qy), C.double(qz), C.double(qw), C.double(cx), C.double(cy), C.double(cz))
}

func (p *EvasMap) Lighting3D(lx, ly, lz, lr, lg, lb, ar, ag, ab int) {
	C.evas_map_util_3d_lighting(p.emap, C.Evas_Coord(lx), C.Evas_Coord(ly), C.Evas_Coord(lz), C.int(lr), C.int(lg), C.int(lb), C.int(ar), C.int(ag), C.int(ab))
}

func (p *EvasMap) Perspective3D(px, py, z0, foc int) {
	C.evas_map_util_3d_perspective(p.emap, C.Evas_Coord(px), C.Evas_Coord(py), C.Evas_Coord(z0), C.Evas_Coord(foc))
}

func (p *EvasMap) IsClockwise() bool {
	return C.evas_map_util_clockwise_get(p.emap) == eTrue
}

func (p *EvasMap) MoveSyncFlag() bool {
	return C.evas_map_util_object_move_sync_get(p.emap) == eTrue
}

func (p *EvasMap) SetMoveSyncFlag(b bool) {
	C.evas_map_util_object_move_sync_set(p.emap, eBool(b))
}

func (p *EvasMap) AlphaFlag() bool {
	return C.evas_map_alpha_get(p.emap) == eTrue
}

func (p *EvasMap) SetAlphaFlag(b bool) {
	C.evas_map_alpha_set(p.emap, eBool(b))
}

func (p *EvasMap) Count() int {
	return int(C.evas_map_count_get(p.emap))
}

func (p *EvasMap) IsSmooth() bool {
	return C.evas_map_smooth_get(p.emap) == eTrue
}

func (p *EvasMap) SetSmooth(b bool) {
	C.evas_map_smooth_set(p.emap, eBool(b))
}

func (p *EvasMap) SetPointCoord(idx, x, y, z int) {
	C.evas_map_point_coord_set(p.emap, C.int(idx), C.Evas_Coord(x), C.Evas_Coord(y), C.Evas_Coord(z))
}

func (p *EvasMap) PointCoord(idx int) (int, int, int) {
	var x, y, z C.Evas_Coord
	C.evas_map_point_coord_get(p.emap, C.int(idx), &x, &y, &z)
	return int(x), int(y), int(z)
}

func (p *EvasMap) SetPointImageUV(idx int, u, v float64) {
	C.evas_map_point_image_uv_set(p.emap, C.int(idx), C.double(u), C.double(v))
}

func (p *EvasMap) PointImageUV(idx int) (float64, float64) {
	var u, v C.double
	C.evas_map_point_image_uv_get(p.emap, C.int(idx), &u, &v)
	return float64(u), float64(v)
}

//-------8<-------------------------------------------------------

type EvasPolygon struct {
	*objectBase
}

func newEvasPolygon(evas *C.Evas) *EvasPolygon {
	eo := C.evas_object_polygon_add(evas)
	return &EvasPolygon{
		wrapObjectBase(eo),
	}
}

func (p *EvasPolygon) AddPoint(x, y int) {
	C.evas_object_polygon_point_add(p.obj, C.Evas_Coord(x), C.Evas_Coord(y))
}

func (p *EvasPolygon) ClearPoints() {
	C.evas_object_polygon_points_clear(p.obj)
}

//-------8<-------------------------------------------------------

type EvasRectangle struct {
	*objectBase
}

func newEvasRectangle(evas *C.Evas) *EvasRectangle {
	eo := C.evas_object_rectangle_add(evas)
	return &EvasRectangle{
		wrapObjectBase(eo),
	}
}

//-------8<-------------------------------------------------------

type EvasTable struct {
	*objectBase
}

func newEvasTable(evas *C.Evas) *EvasTable {
	eo := C.evas_object_table_add(evas)
	return &EvasTable{
		wrapObjectBase(eo),
	}
}

func (p *EvasTable) Homogeneous() TableHomogeneousMode {
	return TableHomogeneousMode(C.evas_object_table_homogeneous_get(p.obj))
}

func (p *EvasTable) SetHomogeneous(mode TableHomogeneousMode) {
	C.evas_object_table_homogeneous_set(p.obj, C.Evas_Object_Table_Homogeneous_Mode(mode))
}

func (p *EvasTable) Padding() (int, int) {
	var h, v C.int
	C.evas_object_table_padding_get(p.obj, &h, &v)
	return int(h), int(v)
}

func (p *EvasTable) SetPadding(h, v int) {
	C.evas_object_table_padding_set(p.obj, C.int(h), C.int(v))
}

func (p *EvasTable) Align() (float64, float64) {
	var h, v C.double
	C.evas_object_table_align_get(p.obj, &h, &v)
	return float64(h), float64(v)
}

func (p *EvasTable) SetAlign(h, v float64) {
	C.evas_object_table_align_set(p.obj, C.double(h), C.double(v))
}

func (p *EvasTable) IsMirrored() bool {
	return C.evas_object_table_mirrored_get(p.obj) == eTrue
}

func (p *EvasTable) SetMirrored(m bool) {
	C.evas_object_table_mirrored_set(p.obj, eBool(m))
}

func (p *EvasTable) Pack(child Object, col int, row int, colspan int, rowspan int) bool {
	return C.evas_object_table_pack(p.obj, child.eo(), C.ushort(col), C.ushort(row), C.ushort(colspan), C.ushort(rowspan)) == eTrue
}

func (p *EvasTable) Unpack(child Object) {
	C.evas_object_table_unpack(p.obj, child.eo())
}

func (p *EvasTable) PackLocation(child Object) (int, int, int, int) {
	var col, row, colspan, rowspan C.ushort
	b := C.evas_object_table_pack_get(p.obj, child.eo(), &col, &row, &colspan, &rowspan)
	if b == eFalse {
		return 0, 0, 0, 0
	}
	return int(col), int(row), int(colspan), int(rowspan)
}

func (p *EvasTable) Clear(del bool) {
	C.evas_object_table_clear(p.obj, eBool(del))
}

func (p *EvasTable) ColRowSize() (int, int) {
	var cols, rows C.int
	C.evas_object_table_col_row_size_get(p.obj, &cols, &rows)
	return int(cols), int(rows)
}

func (p *EvasTable) Children() []Object {
	lst := C.evas_object_table_children_get(p.obj)
	return newListIterator(lst).ObjectSlice()
}

func (p *EvasTable) Child(col, row int) Object {
	return wrapObjectBase(C.evas_object_table_child_get(p.obj, C.ushort(col), C.ushort(row)))
}

//-------8<-------------------------------------------------------

type EvasText struct {
	*objectBase
}

func newEvasText(evas *C.Evas) *EvasText {
	eo := C.evas_object_text_add(evas)
	return &EvasText{
		wrapObjectBase(eo),
	}
}

func (p *EvasText) SetFontSource(file string) {
	cfile := C.CString(file)
	defer free(cfile)
	C.evas_object_text_font_source_set(p.obj, cfile)
}

func (p *EvasText) Font() (string, int) {
	var (
		cfname *C.char
		csize  C.Evas_Font_Size
	)
	C.evas_object_text_font_get(p.obj, &cfname, &csize)
	return C.GoString(cfname), int(csize)
}

func (p *EvasText) SetFont(font string, size int) {
	cfont := C.CString(font)
	defer free(cfont)
	C.evas_object_text_font_set(p.obj, cfont, C.Evas_Font_Size(size))
}

func (p *EvasText) Text() string {
	return C.GoString(C.evas_object_text_text_get(p.obj))
}

func (p *EvasText) SetText(txt string) {
	ctxt := C.CString(txt)
	defer free(ctxt)
	C.evas_object_text_text_set(p.obj, ctxt)
}

func (p *EvasText) Ascent() int {
	return int(C.evas_object_text_ascent_get(p.obj))
}

func (p *EvasText) Descent() int {
	return int(C.evas_object_text_descent_get(p.obj))
}

func (p *EvasText) MaxAscent() int {
	return int(C.evas_object_text_max_ascent_get(p.obj))
}

func (p *EvasText) MaxDescent() int {
	return int(C.evas_object_text_max_descent_get(p.obj))
}

func (p *EvasText) HorizAdvance() int {
	return int(C.evas_object_text_horiz_advance_get(p.obj))
}

func (p *EvasText) VertAdvance() int {
	return int(C.evas_object_text_vert_advance_get(p.obj))
}

func (p *EvasText) Inset() int {
	return int(C.evas_object_text_inset_get(p.obj))
}

func (p *EvasText) CharPos(index int) (int, int, int, int, bool) {
	var x, y, w, h C.int
	r := int(C.evas_object_text_char_pos_get(p.obj, C.int(index), &x, &y, &w, &h))
	if r == 0 {
		return -1, -1, -1, -1, false
	}
	return int(x), int(y), int(w), int(h), true
}

func (p *EvasText) LastUpToPos(x, y int) int {
	return int(C.evas_object_text_last_up_to_pos(p.obj, C.int(x), C.int(y)))
}

//EAPI int 	evas_object_text_char_coords_get (const Evas_Object *obj, Evas_Coord x, Evas_Coord y, Evas_Coord *cx, Evas_Coord *cy, Evas_Coord *cw, Evas_Coord *ch)

func (p *EvasText) TextStyle() TextStyle {
	return TextStyle(C.evas_object_text_style_get(p.obj))
}

func (p *EvasText) SetStyle(style TextStyle) {
	C.evas_object_text_style_set(p.obj, C.Evas_Text_Style_Type(style))
}

func (p *EvasText) ShadowColor() (int, int, int, int) {
	var r, g, b, a C.int
	C.evas_object_text_shadow_color_get(p.obj, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *EvasText) SetShadowColor(r, b, g, a int) {
	C.evas_object_text_shadow_color_set(p.obj, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *EvasText) GlowColor() (int, int, int, int) {
	var r, g, b, a C.int
	C.evas_object_text_glow_color_get(p.obj, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *EvasText) SetGlowColor(r, b, g, a int) {
	C.evas_object_text_glow_color_set(p.obj, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *EvasText) Glow2Color() (int, int, int, int) {
	var r, g, b, a C.int
	C.evas_object_text_glow2_color_get(p.obj, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *EvasText) SetGlow2Color(r, b, g, a int) {
	C.evas_object_text_glow2_color_set(p.obj, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *EvasText) OutlineColor() (int, int, int, int) {
	var r, g, b, a C.int
	C.evas_object_text_outline_color_get(p.obj, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *EvasText) SetOutlineColor(r, b, g, a int) {
	C.evas_object_text_outline_color_set(p.obj, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *EvasText) StylePad() (int, int, int, int) {
	var left, right, top, bottom C.int
	C.evas_object_text_style_pad_get(p.obj, &left, &right, &top, &bottom)
	return int(left), int(right), int(top), int(bottom)
}

//-------8<-------------------------------------------------------

type EvasTextblock struct {
	*evasBase
}

func newEvasTextblock(evas *C.Evas) *EvasTextblock {
	eo := C.evas_object_textblock_add(evas)
	return &EvasTextblock{
		wrapEvas(eo),
	}
}

func wrapEvasTextblock(o *C.Eo) *EvasTextblock {
	if o != nil {
		return &EvasTextblock{
			wrapEvas(o),
		}
	}
	return nil
}

func (p *EvasTextblock) Cursor() *EvasTextblockCursor {
	c := C.evas_object_textblock_cursor_get(p.obj)
	if c != nil {
		return wrapEvasTextblockCursor(c)
	}
	return nil
}

func (p *EvasTextblock) LineGeometry(idx int) (int, int, int, int) {
	var x, y, w, h C.Evas_Coord
	C.evas_object_textblock_line_number_geometry_get(p.obj, C.int(idx), &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *EvasTextblock) Clear() {
	C.evas_object_textblock_clear(p.obj)
}

func (p *EvasTextblock) FormattedSize() (int, int) {
	var w, h C.int
	C.evas_object_textblock_size_formatted_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *EvasTextblock) NativeSize() (int, int) {
	var w, h C.int
	C.evas_object_textblock_size_native_get(p.obj, &w, &h)
	return int(w), int(h)
}

//TODO

/*
struct _Efl_Text_Cursor_Cursor
{
   Evas_Object                     *obj;
   size_t                           pos;
   Evas_Object_Textblock_Node_Text *node;
   Eina_Bool                        changed : 1;
};

*/
type EvasTextblockCursor struct {
	obj *C.Efl_Text_Cursor_Cursor
}

func newEvasTextblockCursor(tb *EvasTextblock) *EvasTextblockCursor {
	tbc := C.evas_object_textblock_cursor_new(tb.obj)
	return wrapEvasTextblockCursor(tbc)
}

func wrapEvasTextblockCursor(o *C.Efl_Text_Cursor_Cursor) *EvasTextblockCursor {
	return &EvasTextblockCursor{
		obj: o,
	}
}

func (p *EvasTextblockCursor) Dispose() {
	C.evas_textblock_cursor_free(p.obj)
}

//-------8<-------------------------------------------------------

type EvasTextgridCell struct {
	obj *C.Evas_Textgrid_Cell
}

func wrapEvasTextgridCell(o *C.Evas_Textgrid_Cell) *EvasTextgridCell {
	if o != nil {
		return &EvasTextgridCell{o}
	}
	return nil
}

//TODO: this is not true, make conversion
func (p *EvasTextgridCell) Codepoint() rune {
	return rune(p.obj.codepoint)
}

func (p *EvasTextgridCell) SetCodepoint(r rune) {
	p.obj.codepoint = C.Eina_Unicode(r)
}

func (p *EvasTextgridCell) FgIndex() byte {
	return byte(p.obj.fg)
}

func (p *EvasTextgridCell) SetFgIndex(i byte) {
	p.obj.fg = C.uchar(i)
}

func (p *EvasTextgridCell) BgIndex() byte {
	return byte(p.obj.bg)
}

func (p *EvasTextgridCell) SetBgIndex(i byte) {
	p.obj.bg = C.uchar(i)
}

func (p *EvasTextgridCell) IsBold() bool {
	return C.cgo_evas_textgrid_cell_bold_get(p.obj) != 0
}

func (p *EvasTextgridCell) SetBold(b bool) {
	var val C.ushort
	if b {
		val = 1
	}
	C.cgo_evas_textgrid_cell_bold_set(p.obj, val)
}

func (p *EvasTextgridCell) IsItalic() bool {
	return C.cgo_evas_textgrid_cell_italic_get(p.obj) != 0
}

func (p *EvasTextgridCell) SetItalic(b bool) {
	var val C.ushort
	if b {
		val = 1
	}
	C.cgo_evas_textgrid_cell_italic_set(p.obj, val)
}

func (p *EvasTextgridCell) IsUnderline() bool {
	return C.cgo_evas_textgrid_cell_underline_get(p.obj) != 0
}

func (p *EvasTextgridCell) SetUnderline(b bool) {
	var val C.ushort
	if b {
		val = 1
	}
	C.cgo_evas_textgrid_cell_underline_set(p.obj, val)
}

func (p *EvasTextgridCell) IsStrikethrough() bool {
	return C.cgo_evas_textgrid_cell_strikethrough_get(p.obj) != 0
}

func (p *EvasTextgridCell) SetStrikethrough(b bool) {
	var val C.ushort
	if b {
		val = 1
	}
	C.cgo_evas_textgrid_cell_strikethrough_set(p.obj, val)
}

func (p *EvasTextgridCell) IsFgExtended() bool {
	return C.cgo_evas_textgrid_cell_fg_extended_get(p.obj) != 0
}

func (p *EvasTextgridCell) SetFgExtended(b bool) {
	var val C.ushort
	if b {
		val = 1
	}
	C.cgo_evas_textgrid_cell_fg_extended_set(p.obj, val)
}
func (p *EvasTextgridCell) IsBgExtended() bool {
	return C.cgo_evas_textgrid_cell_bg_extended_get(p.obj) != 0
}

func (p *EvasTextgridCell) SetBgExtended(b bool) {
	var val C.ushort
	if b {
		val = 1
	}
	C.cgo_evas_textgrid_cell_bg_extended_set(p.obj, val)
}

func (p *EvasTextgridCell) IsDoubleWidth() bool {
	return C.cgo_evas_textgrid_cell_double_width_get(p.obj) != 0
}

func (p *EvasTextgridCell) SetDoubleWidth(b bool) {
	var val C.ushort
	if b {
		val = 1
	}
	C.cgo_evas_textgrid_cell_double_width_set(p.obj, val)
}

type EvasTextgrid struct {
	*objectBase
}

func newEvasTextgrid(evas *C.Evas) *EvasTextgrid {
	o := C.evas_object_textgrid_add(evas)
	return wrapEvasTextgrid(o)
}

func wrapEvasTextgrid(o *C.Eo) *EvasTextgrid {
	if o != nil {
		return &EvasTextgrid{wrapObjectBase(o)}
	}
	return nil
}

func (p *EvasTextgrid) SetTextgridSize(w, h int) {
	C.evas_object_textgrid_size_set(p.obj, C.int(w), C.int(h))
}

func (p *EvasTextgrid) TextgridSize() (int, int) {
	var w, h C.int
	C.evas_object_textgrid_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *EvasTextgrid) SetFontSource(src string) {
	var csrc *C.char
	if src != "" {
		csrc = C.CString(src)
		defer free(csrc)
	}
	C.evas_object_textgrid_font_source_set(p.obj, csrc)
}

func (p *EvasTextgrid) FontSource() string {
	return C.GoString(C.evas_object_textgrid_font_source_get(p.obj))
}

func (p *EvasTextgrid) SetFont(fnt string, size int) {
	var cfnt *C.char
	if fnt != "" {
		cfnt = C.CString(fnt)
		defer free(cfnt)
	}
	C.evas_object_textgrid_font_set(p.obj, cfnt, C.Evas_Font_Size(size))
}

func (p *EvasTextgrid) Font() (string, int) {
	var (
		cfnt  *C.char
		csize C.Evas_Font_Size
	)
	C.evas_object_textgrid_font_get(p.obj, &cfnt, &csize)
	return C.GoString(cfnt), int(csize)
}

func (p *EvasTextgrid) CellSize() (int, int) {
	var w, h C.int
	C.evas_object_textgrid_cell_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *EvasTextgrid) SetPaletteColor(t EvasTextgridPalette, idx, r, g, b, a int) {
	C.evas_object_textgrid_palette_set(p.obj, C.Evas_Textgrid_Palette(t), C.int(idx), C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *EvasTextgrid) PaletteColor(t EvasTextgridPalette, idx int) (int, int, int, int) {
	var r, g, b, a C.int
	C.evas_object_textgrid_palette_get(p.obj, C.Evas_Textgrid_Palette(t), C.int(idx), &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

// not implemented in efl always return Normal
func (p *EvasTextgrid) SetSupportedFontStyles(style EvasTextgridFontStyle) {
	C.evas_object_textgrid_supported_font_styles_set(p.obj, C.Evas_Textgrid_Font_Style(style))
}

func (p *EvasTextgrid) SupportedFontStyles() EvasTextgridFontStyle {
	return EvasTextgridFontStyle(C.evas_object_textgrid_supported_font_styles_get(p.obj))
}

func (p *EvasTextgrid) SetCellRow(y int, cell *EvasTextgridCell) {
	C.evas_object_textgrid_cellrow_set(p.obj, C.int(y), cell.obj)
}

func (p *EvasTextgrid) CellRow(y int) *EvasTextgridCell {
	return wrapEvasTextgridCell(C.evas_object_textgrid_cellrow_get(p.obj, C.int(y)))
}

func (p *EvasTextgrid) SetUpdated(x, y, w, h int) {
	C.evas_object_textgrid_update_add(p.obj, C.int(x), C.int(y), C.int(w), C.int(h))
}

//-------8<-------------------------------------------------------

/*
type Smart struct {}

func (p *Smart) AddMember(o Object) {
	C.evas_object_smart_member_add(p.obj, o.Obj())
}

func SmartObjectDelMember(o Object) {
	C.evas_object_smart_member_del(o.Obj())
}

func (p *Smart) Members() []Object {
	lst := C.evas_object_smart_members_get(p.obj)
	return ObjectListToSlice(lst, true)
}

func (p *Smart) SetData(data interface{}) {
	C.evas_object_smart_data_set(p.obj, unsafe.Pointer(&data))
}

func (p *Smart) Data() interface{} {
	v := C.evas_object_smart_data_get(p.obj)
	if v != nil {
		return unsafe.Pointer(v)
	}
	return nil
}
*/

//-------8<-------------------------------------------------------

/*
struct _Emotion_Webcam
{
   EINA_REFCOUNT;

   const char *syspath;
   const char *device;
   const char *name;

   const char *custom;

   const char *filename;
};
*/

type Webcam struct {
	obj *C.Emotion_Webcam
}

func wrapWebcam(o *C.Emotion_Webcam) *Webcam {
	if o != nil {
		return &Webcam{o}
	}
	return nil
}

func (p *Webcam) Name() string {
	return C.GoString(C.emotion_webcam_name_get(p.obj))
}

func (p *Webcam) Device() string {
	return C.GoString(C.emotion_webcam_device_get(p.obj))
}

func Webcams() []*Webcam {
	lst := C.emotion_webcams_get()
	if lst != nil {
		return newListIterator(lst).WebcamSlice()
	}
	return nil
}

//-------8<-------------------------------------------------------

type Emotion struct {
	*objectBase
}

func NewEmotion(canvas *evasBase) *Emotion {
	eo := C.emotion_object_add(canvas.obj)
	return wrapEmotion(eo)
}

func wrapEmotion(o *C.Eo) *Emotion {
	if o != nil {
		return &Emotion{wrapObjectBase(o)}
	}
	return nil
}

// only check file extension, not mime type
func IsExtensionSupportedByEmotion(filename string) bool {
	cfilename := C.CString(filename)
	defer free(cfilename)
	return C.emotion_object_extension_may_play_get(cfilename) == eTrue
}

func IsExtensionSupportedFastByEmotion(filename string) bool {
	cfilename := C.CString(filename)
	defer free(cfilename)
	return C.emotion_object_extension_may_play_fast_get(cfilename) == eTrue
}

// modulename: gstreamer, vlc etc
func (p *Emotion) Init(modulename string) int {
	var cmodulename *C.char
	if modulename != "" {
		cmodulename = C.CString(modulename)
		defer free(cmodulename)
	}
	return int(C.emotion_object_init(p.obj, cmodulename))
}

func (p *Emotion) SetFile(filename string) {
	cfilename := C.CString(filename)
	defer free(cfilename)
	C.emotion_object_file_set(p.obj, cfilename)
}

func (p *Emotion) File() string {
	return C.GoString(C.emotion_object_file_get(p.obj))
}

func (p *Emotion) SetPlay(b bool) {
	C.emotion_object_play_set(p.obj, eBool(b))
}

func (p *Emotion) IsPlaying() bool {
	return C.emotion_object_play_get(p.obj) == eTrue
}

func (p *Emotion) SetPosition(pos float64) {
	C.emotion_object_position_set(p.obj, C.double(pos))
}

func (p *Emotion) Position() float64 {
	return float64(C.emotion_object_position_get(p.obj))
}

func (p *Emotion) BufferSize() float64 {
	return float64(C.emotion_object_buffer_size_get(p.obj))
}

func (p *Emotion) IsSeekable() bool {
	return C.emotion_object_seekable_get(p.obj) == eTrue
}

func (p *Emotion) PlayLength() float64 {
	return float64(C.emotion_object_play_length_get(p.obj))
}

func (p *Emotion) SetPlaySpeed(sp float64) {
	C.emotion_object_play_speed_set(p.obj, C.double(sp))
}

func (p *Emotion) PlaySpeed() float64 {
	return float64(C.emotion_object_play_speed_get(p.obj))
}

func (p *Emotion) ProgressInfo() string {
	return C.GoString(C.emotion_object_progress_info_get(p.obj))
}

func (p *Emotion) ProgressStatus() float64 {
	return float64(C.emotion_object_progress_status_get(p.obj))
}

func (p *Emotion) SetBorder(l, r, t, b int) {
	C.emotion_object_border_set(p.obj, C.int(l), C.int(r), C.int(t), C.int(b))
}

func (p *Emotion) Border() (int, int, int, int) {
	var l, r, t, b C.int
	C.emotion_object_border_get(p.obj, &l, &r, &t, &b)
	return int(l), int(r), int(t), int(b)
}

func (p *Emotion) SetBgColor(r, g, b, a int) {
	C.emotion_object_bg_color_set(p.obj, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *Emotion) BgColor() (int, int, int, int) {
	var r, g, b, a C.int
	C.emotion_object_bg_color_get(p.obj, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *Emotion) SetKeepAspect(aspect EmotionAspect) {
	C.emotion_object_keep_aspect_set(p.obj, C.Emotion_Aspect(aspect))
}

func (p *Emotion) KeepAspect() EmotionAspect {
	return EmotionAspect(C.emotion_object_keep_aspect_get(p.obj))
}

func (p *Emotion) AspectRatio() float64 {
	return float64(C.emotion_object_ratio_get(p.obj))
}

func (p *Emotion) VideoSize() (int, int) {
	var w, h C.int
	C.emotion_object_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Emotion) SetSmoothScale(b bool) {
	C.emotion_object_smooth_scale_set(p.obj, eBool(b))
}

func (p *Emotion) IsSmoothScale() bool {
	return C.emotion_object_smooth_scale_get(p.obj) == eTrue
}

func (p *Emotion) SetVideoMute(b bool) {
	C.emotion_object_video_mute_set(p.obj, eBool(b))
}

func (p *Emotion) IsVideoMuted() bool {
	return C.emotion_object_video_mute_get(p.obj) == eTrue
}

func (p *Emotion) SetSubtitle(filepath string) {
	cfilepath := C.CString(filepath)
	defer free(cfilepath)
	C.emotion_object_video_subtitle_file_set(p.obj, cfilepath)
}

func (p *Emotion) Subtitle() string {
	return C.GoString(C.emotion_object_video_subtitle_file_get(p.obj))
}

func (p *Emotion) VideoChannelCount() int {
	return int(C.emotion_object_video_channel_count(p.obj))
}

func (p *Emotion) VideoChannelName(ch int) string {
	return C.GoString(C.emotion_object_video_channel_name_get(p.obj, C.int(ch)))
}

func (p *Emotion) SetVideoChannel(ch int) {
	C.emotion_object_video_channel_set(p.obj, C.int(ch))
}

func (p *Emotion) VideoChannel() int {
	return int(C.emotion_object_video_channel_get(p.obj))
}

func (p *Emotion) SetAudioVolume(vol float64) {
	C.emotion_object_audio_volume_set(p.obj, C.double(vol))
}

func (p *Emotion) AudioVolume() float64 {
	return float64(C.emotion_object_audio_volume_get(p.obj))
}

func (p *Emotion) SetAudioMute(b bool) {
	C.emotion_object_audio_mute_set(p.obj, eBool(b))
}

func (p *Emotion) IsAudioMuted() bool {
	return C.emotion_object_audio_mute_get(p.obj) == eTrue
}

func (p *Emotion) SetVisualization(vis EmotionVisualization) {
	C.emotion_object_vis_set(p.obj, C.Emotion_Vis(vis))
}

func (p *Emotion) Visualization() EmotionVisualization {
	return EmotionVisualization(C.emotion_object_vis_get(p.obj))
}

func (p *Emotion) IsVisualizationSupported(vis EmotionVisualization) bool {
	return C.emotion_object_vis_supported(p.obj, C.Emotion_Vis(vis)) == eTrue
}

func (p *Emotion) SetPriority(b bool) {
	C.emotion_object_priority_set(p.obj, eBool(b))
}

func (p *Emotion) IsPriority() bool {
	return C.emotion_object_priority_get(p.obj) == eTrue
}

func (p *Emotion) SetSuspend(es EmotionSuspend) {
	C.emotion_object_suspend_set(p.obj, C.Emotion_Suspend(es))
}

func (p *Emotion) Suspend() EmotionSuspend {
	return EmotionSuspend(C.emotion_object_suspend_get(p.obj))
}

func (p *Emotion) Title() string {
	return C.GoString(C.emotion_object_title_get(p.obj))
}

func (p *Emotion) MetaInfo(mi EmotionMetaInfo) string {
	return C.GoString(C.emotion_object_meta_info_get(p.obj, C.Emotion_Meta_Info(mi)))
}

func (p *Emotion) LoadLastPostion() {
	C.emotion_object_last_position_load(p.obj)
}

func (p *Emotion) SaveLastPosition() {
	C.emotion_object_last_position_save(p.obj)
}

func (p *Emotion) IsVideoHandled() bool {
	return C.emotion_object_video_handled_get(p.obj) == eTrue
}

func (p *Emotion) IsAudioHandled() bool {
	return C.emotion_object_audio_handled_get(p.obj) == eTrue
}

func (p *Emotion) SendSimpleEvent(ev EmotionEvent) {
	C.emotion_object_event_simple_send(p.obj, C.Emotion_Event(ev))
}

func (p *Emotion) AudioChannelCount() int {
	return int(C.emotion_object_audio_channel_count(p.obj))
}

func (p *Emotion) AudioChannelName(ch int) string {
	return C.GoString(C.emotion_object_audio_channel_name_get(p.obj, C.int(ch)))
}

func (p *Emotion) SetAudioChannel(ch int) {
	C.emotion_object_audio_channel_set(p.obj, C.int(ch))
}

func (p *Emotion) AudioChannel() int {
	return int(C.emotion_object_audio_channel_get(p.obj))
}

func (p *Emotion) SetSpuMute(b bool) {
	C.emotion_object_spu_mute_set(p.obj, eBool(b))
}

func (p *Emotion) IsSpuMuted() bool {
	return C.emotion_object_spu_mute_get(p.obj) == eTrue
}

func (p *Emotion) SpuChannelCount() int {
	return int(C.emotion_object_spu_channel_count(p.obj))
}

func (p *Emotion) SpuChannelName(ch int) string {
	return C.GoString(C.emotion_object_spu_channel_name_get(p.obj, C.int(ch)))
}

func (p *Emotion) SetSpuChannel(ch int) {
	C.emotion_object_spu_channel_set(p.obj, C.int(ch))
}

func (p *Emotion) SpuChannel() int {
	return int(C.emotion_object_spu_channel_get(p.obj))
}

func (p *Emotion) ChapterCount() int {
	return int(C.emotion_object_chapter_count(p.obj))
}

func (p *Emotion) SetChapter(cp int) {
	C.emotion_object_chapter_set(p.obj, C.int(cp))
}

func (p *Emotion) Chapter() int {
	return int(C.emotion_object_chapter_get(p.obj))
}

func (p *Emotion) ChapterName(cp int) string {
	return C.GoString(C.emotion_object_chapter_name_get(p.obj, C.int(cp)))
}

func (p *Emotion) Eject() {
	C.emotion_object_eject(p.obj)
}

func (p *Emotion) RefFile() string {
	return C.GoString(C.emotion_object_ref_file_get(p.obj))
}

func (p *Emotion) RefNum() int {
	return int(C.emotion_object_ref_num_get(p.obj))
}

func (p *Emotion) SpuButtonCount() int {
	return int(C.emotion_object_spu_button_count_get(p.obj))
}

func (p *Emotion) SpuButton() int {
	return int(C.emotion_object_spu_button_get(p.obj))
}

func (p *Emotion) ImageObject() Object {
	return wrapObjectBase(C.emotion_object_image_get(p.obj))
}

/*
TODO: in 1.19
Evas_Object* emotion_file_meta_artwork_get (const Evas_Object *	obj,const char *path,Emotion_Artwork_Info type)
*/

//-------8<-------------------------------------------------------
//-------8<-------------------------------------------------------
//-------8<-------------------------------------------------------
