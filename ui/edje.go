package ui

/*
#include "bridge.h"
*/
import "C"

type edjeConfig struct{}

var EdjeConfig = edjeConfig{}

func (edjeConfig) SetFrametime(ft float64) {
	C.edje_frametime_set(C.double(ft))
}

func (edjeConfig) Frametime() float64 {
	return float64(C.edje_frametime_get())
}

func (edjeConfig) Freeze() {
	C.edje_freeze()
}

func (edjeConfig) Thaw() {
	C.edje_thaw()
}

func (edjeConfig) SetFontsetAppend(font string) {
	cfont := C.CString(font)
	defer free(cfont)
	C.edje_fontset_append_set(cfont)
}

func (edjeConfig) FontsetAppend() string {
	return C.GoString(C.edje_fontset_append_get())
}

func (edjeConfig) FileCollection(file string) []string {
	cfile := C.CString(file)
	defer free(cfile)

	lst := C.edje_file_collection_list(cfile)
	if lst != nil {
		defer C.edje_file_collection_list_free(lst)
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func (edjeConfig) IsFileGroupExists(file, group string) bool {
	cfile := C.CString(file)
	cgroup := C.CString(group)
	defer free(cfile, cgroup)
	return C.edje_file_group_exists(cfile, cgroup) == eTrue
}

func (edjeConfig) FileData(file, key string) string {
	cfile := C.CString(file)
	ckey := C.CString(key)
	ret := C.edje_file_data_get(cfile, ckey)
	defer free(cfile, ckey, ret)
	return C.GoString(ret)
}

func (edjeConfig) SetFileCache(count int) {
	C.edje_file_cache_set(C.int(count))
}

func (edjeConfig) FileCache() int {
	return int(C.edje_file_cache_get())
}

func (edjeConfig) FlushFileCache() {
	C.edje_file_cache_flush()
}

func (edjeConfig) SetCollectionCache(count int) {
	C.edje_collection_cache_set(C.int(count))
}

func (edjeConfig) CollectionCache() int {
	return int(C.edje_collection_cache_get())
}

func (edjeConfig) FlushCollectionCache() {
	C.edje_collection_cache_flush()
}

func (edjeConfig) SetScale(scale float64) {
	C.edje_scale_set(C.double(scale))
}

func (edjeConfig) Scale() float64 {
	return float64(C.edje_scale_get())
}

func (edjeConfig) SetPasswordShowLast(b bool) {
	C.edje_password_show_last_set(eBool(b))
}

func (edjeConfig) SetPasswordShowLastTimeout(to float64) {
	C.edje_password_show_last_timeout_set(C.double(to))
}

func (edjeConfig) SetColorClass(colorclass string, r, g, b, a, r2, g2, b2, a2, r3, g3, b3, a3 int) {
	ccc := C.CString(colorclass)
	defer free(ccc)
	C.edje_color_class_set(ccc, C.int(r), C.int(g), C.int(b), C.int(a), C.int(r2), C.int(g2), C.int(b2), C.int(a2), C.int(r3), C.int(g3), C.int(b3), C.int(a3))
}

func (edjeConfig) ColorClass(colorclass string) (int, int, int, int, int, int, int, int, int, int, int, int) {
	var r, g, b, a, r2, g2, b2, a2, r3, g3, b3, a3 C.int
	ccc := C.CString(colorclass)
	defer free(ccc)
	C.edje_color_class_get(ccc, &r, &g, &b, &a, &r2, &g2, &b2, &a2, &r3, &g3, &b3, &a3)
	return int(r), int(g), int(b), int(a), int(r2), int(g2), int(b2), int(a2), int(r3), int(g3), int(b3), int(a3)
}

func (edjeConfig) DelColorClass(colorclass string) {
	ccc := C.CString(colorclass)
	defer free(ccc)
	C.edje_color_class_del(ccc)
}

func (edjeConfig) ColorClasses() []string {
	lst := C.edje_color_class_list()
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func (edjeConfig) SetTextClass(textclass, font string, size int) {
	ctc := C.CString(textclass)
	cfont := C.CString(font)
	defer free(ctc, cfont)
	C.edje_text_class_set(ctc, cfont, C.Evas_Font_Size(size))
}

func (edjeConfig) TextClass(textclass string) (string, int) {
	var (
		cfont *C.char
		size  C.Evas_Font_Size
	)
	ctc := C.CString(textclass)
	defer free(ctc)
	C.edje_text_class_get(ctc, &cfont, &size)
	return C.GoString(cfont), int(size)
}

func (edjeConfig) DelTextClass(textclass string) {
	ctc := C.CString(textclass)
	defer free(ctc)
	C.edje_text_class_del(ctc)
}

func (edjeConfig) TextClasses() []string {
	lst := C.edje_text_class_list()
	if lst != nil {
		return newListIterator(lst).StringShareSlice()
	}
	return nil
}

func (edjeConfig) SetSizeClass(sizeclass string, minw, minh, maxw, maxh int) bool {
	csc := C.CString(sizeclass)
	defer free(csc)
	return C.edje_size_class_set(csc, C.Evas_Coord(minw), C.Evas_Coord(minh), C.Evas_Coord(maxw), C.Evas_Coord(maxh)) == eTrue
}

func (edjeConfig) SizeClass(sizeclass string) (bool, int, int, int, int) {
	var minw, minh, maxw, maxh C.Evas_Coord
	csc := C.CString(sizeclass)
	defer free(csc)
	ret := C.edje_size_class_get(csc, &minw, &minh, &maxw, &maxh)
	if ret != eTrue {
		return false, 0, 0, 0, 0
	}
	return true, int(minw), int(minh), int(maxw), int(maxh)
}

func (edjeConfig) DelSizeClass(sizeclass string) {
	csc := C.CString(sizeclass)
	defer free(csc)
	C.edje_size_class_del(csc)
}

func (edjeConfig) SizeClasses() []string {
	lst := C.edje_size_class_list()
	if lst != nil {
		return newListIterator(lst).StringShareSlice()
	}
	return nil
}

func (edjeConfig) ProcessMessageSignal() {
	C.edje_message_signal_process()
}

/* deprecated
func (edjeConfig) SetExternObjectMinSize(o Object, w, h int) {
	C.edje_extern_object_min_size_set(o.eo(), C.int(w), C.int(h))
}
*/

func (edjeConfig) SetExternObjectMaxSize(o Object, w, h int) {
	C.edje_extern_object_max_size_set(o.eo(), C.Evas_Coord(w), C.Evas_Coord(h))
}

func (edjeConfig) SetExternObjectAspect(o Object, control EdjeAspectControl, w, h int) {
	C.edje_extern_object_aspect_set(o.eo(), C.Edje_Aspect_Control(control), C.Evas_Coord(w), C.Evas_Coord(h))
}

func (edjeConfig) AvailableModules() []string {
	lst := C.edje_available_modules_get()
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func (edjeConfig) LoadModule(name string) {
	cname := C.CString(name)
	defer free(cname)
	C.edje_module_load(cname)
}

func (edjeConfig) SetLanguage(lang string) {
	clang := C.CString(lang)
	defer free(clang)
	C.edje_language_set(clang)
}

func (edjeConfig) SetTransitionDurationFactor(scale float64) {
	C.edje_transition_duration_factor_set(C.double(scale))
}

func (edjeConfig) TransitionDurationFactor() float64 {
	return float64(C.edje_transition_duration_factor_get())
}

type Perspective struct {
	obj *C.Edje_Perspective
}

func NewPerspective(canvas Canvas) *Perspective {
	o := C.edje_perspective_new(canvas.eo())
	return wrapPerspective(o)
}

func GlobalPerspective(canvas Canvas) *Perspective {
	p := C.edje_evas_global_perspective_get(canvas.eo())
	return wrapPerspective(p)
}

func wrapPerspective(o *C.Edje_Perspective) *Perspective {
	if o != nil {
		return &Perspective{o}
	}
	return nil
}

func (p *Perspective) Dispose() {
	C.edje_perspective_free(p.obj)
}

func (p *Perspective) Set(px, py, z0, foc int) {
	C.edje_perspective_set(p.obj, C.Evas_Coord(px), C.Evas_Coord(py), C.Evas_Coord(z0), C.Evas_Coord(foc))
}

func (p *Perspective) SetGlobal(b bool) {
	C.edje_perspective_global_set(p.obj, eBool(b))
}

func (p *Perspective) IsGlobal() bool {
	return C.edje_perspective_global_get(p.obj) == eTrue
}

type Edje struct {
	*objectBase
}

func NewEdje(canvas Canvas) *Edje {
	eo := C.edje_object_add(canvas.eo())
	return wrapEdje(eo)
}

func wrapEdje(o *C.Eo) *Edje {
	if o != nil {
		return &Edje{wrapObjectBase(o)}
	}
	return nil
}

func (p *Edje) CollectionData(key string) string {
	ckey := C.CString(key)
	defer free(ckey)
	return C.GoString(C.edje_object_data_get(p.obj, ckey))
}

func (p *Edje) SetFile(file, group string) error {
	var cgroup *C.char

	cfile := C.CString(file)
	defer free(cfile)
	if group != "" {
		cgroup = C.CString(group)
		defer free(cgroup)
	}
	if C.edje_object_file_set(p.obj, cfile, cgroup) != eTrue {
		return p.LoadError()
	}
	return nil
}

func (p *Edje) File() (string, string) {
	var cfile, cgroup *C.char
	C.edje_object_file_get(p.obj, &cfile, &cgroup)
	return C.GoString(cfile), C.GoString(cgroup)
}

func (p *Edje) LoadError() EdjeLoadError {
	return EdjeLoadError(C.edje_object_load_error_get(p.obj))
}

func (p *Edje) IsPlaying() bool {
	return C.edje_object_play_get(p.obj) == eTrue
}

func (p *Edje) SetPlay(b bool) {
	C.edje_object_play_set(p.obj, eBool(b))
}

func (p *Edje) SetAnimation(b bool) {
	C.edje_object_animation_set(p.obj, eBool(b))
}

func (p *Edje) IsAnimation() bool {
	return C.edje_object_animation_get(p.obj) == eTrue
}

func (p *Edje) Freeze() {
	C.edje_object_freeze(p.obj)
}

func (p *Edje) Thaw() {
	C.edje_object_thaw(p.obj)
}

func (p *Edje) Preload(cancel bool) {
	C.edje_object_preload(p.obj, eBool(cancel))
}

func (p *Edje) SetColorClass(colorclass string, r, g, b, a, r2, g2, b2, a2, r3, g3, b3, a3 int) {
	ccc := C.CString(colorclass)
	defer free(ccc)
	C.edje_object_color_class_set(p.obj, ccc, C.int(r), C.int(g), C.int(b), C.int(a), C.int(r2), C.int(g2), C.int(b2), C.int(a2), C.int(r3), C.int(g3), C.int(b3), C.int(a3))
}

func (p *Edje) ColorClass(colorclass string) (int, int, int, int, int, int, int, int, int, int, int, int) {
	var r, g, b, a, r2, g2, b2, a2, r3, g3, b3, a3 C.int
	ccc := C.CString(colorclass)
	defer free(ccc)
	C.edje_object_color_class_get(p.obj, ccc, &r, &g, &b, &a, &r2, &g2, &b2, &a2, &r3, &g3, &b3, &a3)
	return int(r), int(g), int(b), int(a), int(r2), int(g2), int(b2), int(a2), int(r3), int(g3), int(b3), int(a3)
}

func (p *Edje) DelColorClass(colorclass string) {
	ccc := C.CString(colorclass)
	defer free(ccc)
	C.edje_object_color_class_del(p.obj, ccc)
}

func (p *Edje) ClearColorClass() bool {
	return C.edje_object_color_class_clear(p.obj) == eTrue
}

func (p *Edje) SetTextClass(textclass, font string, size int) {
	ctc := C.CString(textclass)
	cfont := C.CString(font)
	defer free(ctc, cfont)
	C.edje_object_text_class_set(p.obj, ctc, cfont, C.int(size))
}

func (p *Edje) TextClass(textclass string) (string, int) {
	var (
		cfont *C.char
		size  C.int
	)
	ctc := C.CString(textclass)
	defer free(ctc)
	C.edje_object_text_class_get(p.obj, ctc, &cfont, &size)
	return C.GoString(cfont), int(size)
}

func (p *Edje) DelTextClass(textclass string) {
	ctc := C.CString(textclass)
	defer free(ctc)
	C.edje_object_text_class_del(p.obj, ctc)
}

func (p *Edje) SetSizeClass(sizeclass string, minw, minh, maxw, maxh int) bool {
	csc := C.CString(sizeclass)
	defer free(csc)
	return C.edje_object_size_class_set(p.obj, csc, C.int(minw), C.int(minh), C.int(maxw), C.int(maxh)) == eTrue
}

func (p *Edje) SizeClass(sizeclass string) (bool, int, int, int, int) {
	var minw, minh, maxw, maxh C.int
	csc := C.CString(sizeclass)
	defer free(csc)
	ret := C.edje_object_size_class_get(p.obj, csc, &minw, &minh, &maxw, &maxh)
	if ret != eTrue {
		return false, 0, 0, 0, 0
	}
	return true, int(minw), int(minh), int(maxw), int(maxh)
}

func (p *Edje) DelSizeClass(sizeclass string) {
	csc := C.CString(sizeclass)
	defer free(csc)
	C.edje_object_size_class_del(p.obj, csc)
}

func (p *Edje) SetScale(scale float64) {
	C.edje_object_scale_set(p.obj, C.double(scale))
}

func (p *Edje) Scale() float64 {
	return float64(C.edje_object_scale_get(p.obj))
}

func (p *Edje) BaseScale() float64 {
	return float64(C.edje_object_base_scale_get(p.obj))
}

func (p *Edje) SetMirrored(b bool) {
	C.edje_object_mirrored_set(p.obj, eBool(b))
}

func (p *Edje) IsMirrored() bool {
	return C.edje_object_mirrored_get(p.obj) == eTrue
}

func (p *Edje) MinSize() (int, int) {
	var w, h C.int
	C.edje_object_size_min_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Edje) MaxSize() (int, int) {
	var w, h C.int
	C.edje_object_size_max_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Edje) ForceCalc() {
	C.edje_object_calc_force(p.obj)
}

func (p *Edje) CalcMinSize() (int, int) {
	var w, h C.int
	C.edje_object_size_min_calc(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Edje) CalcRestrictedMinSize(minw, minh int) (int, int) {
	var w, h C.int
	C.edje_object_size_min_restricted_calc(p.obj, &w, &h, C.int(minw), C.int(minh))
	return int(w), int(h)
}

func (p *Edje) CalcPartExtends() (int, int, int, int) {
	var x, y, w, h C.int
	C.edje_object_parts_extends_calc(p.obj, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *Edje) SetUpdateHints(b bool) {
	C.edje_object_update_hints_set(p.obj, eBool(b))
}

func (p *Edje) IsUpdateHints() bool {
	return C.edje_object_update_hints_get(p.obj) == eTrue
}

func (p *Edje) PartExists(part string) bool {
	cpart := C.CString(part)
	defer free(cpart)
	return C.edje_object_part_exists(p.obj, cpart) == eTrue
}

func (p *Edje) PartObject(part string) Object {
	cpart := C.CString(part)
	defer free(cpart)
	return wrapObjectBase(C.edje_object_part_object_get(p.obj, cpart))
}

func (p *Edje) PartGeometry(part string) (int, int, int, int) {
	var x, y, w, h C.int

	cpart := C.CString(part)
	defer free(cpart)

	C.edje_object_part_geometry_get(p.obj, cpart, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *Edje) SetPartText(part, txt string) bool {
	cpart := C.CString(part)
	ctxt := C.CString(txt)
	defer free(cpart, ctxt)
	return C.edje_object_part_text_set(p.obj, cpart, ctxt) == eTrue
}

func (p *Edje) PartText(part string) string {
	cpart := C.CString(part)
	defer free(cpart)
	return C.GoString(C.edje_object_part_text_get(p.obj, cpart))
}

func (p *Edje) SelectAllPartText(part string) {
	cpart := C.CString(part)
	defer free(cpart)
	C.edje_object_part_text_select_all(p.obj, cpart)
}

func (p *Edje) DeselectPartText(part string) {
	cpart := C.CString(part)
	defer free(cpart)
	C.edje_object_part_text_select_none(p.obj, cpart)
}

func (p *Edje) SetPartUnescapedText(part, txt string) {
	cpart := C.CString(part)
	ctxt := C.CString(txt)
	defer free(cpart, ctxt)
	C.edje_object_part_text_unescaped_set(p.obj, cpart, ctxt)
}

func (p *Edje) PartUnescapedText(part string) string {
	cpart := C.CString(part)
	ret := C.edje_object_part_text_unescaped_get(p.obj, cpart)
	defer free(cpart, ret)
	return C.GoString(ret)
}

func (p *Edje) SetPartTextInputHint(part string, hint EdjeInputHint) {
	cpart := C.CString(part)
	defer free(cpart)
	C.edje_object_part_text_input_hint_set(p.obj, cpart, C.Edje_Input_Hints(hint))
}

func (p *Edje) PartTextInputHint(part string) EdjeInputHint {
	cpart := C.CString(part)
	defer free(cpart)
	return EdjeInputHint(C.edje_object_part_text_input_hint_get(p.obj, cpart))
}

func (p *Edje) SwallowPart(part string, o Object) {
	cpart := C.CString(part)
	defer free(cpart)
	C.edje_object_part_swallow(p.obj, cpart, o.eo())
}

func (p *Edje) UnswallowPart(o Object) {
	C.edje_object_part_unswallow(p.obj, o.eo())
}

func (p *Edje) PartSwallow(part string) Object {
	cpart := C.CString(part)
	defer free(cpart)
	return wrapObjectBase(C.edje_object_part_swallow_get(p.obj, cpart))
}

func (p *Edje) PartExternalObject(part string) Object {
	cpart := C.CString(part)
	defer free(cpart)
	return wrapObjectBase(C.edje_object_part_external_object_get(p.obj, cpart))
}

/*
TODO:
edje_object_text_change_cb_set

*/
