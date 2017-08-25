package ui

/*
#cgo pkg-config: --cflags --libs elementary
#cgo CFLAGS: -O2 -g

#include "bridge.h"
*/
import "C"

import (
	"os"
	"sync"
	"unsafe"
)

func init() {
	argc := len(os.Args)
	argv := make([]*C.char, argc)
	for i, a := range os.Args {
		argv[i] = C.CString(a)
	}
	defer free(argv)
	C.elm_init(C.int(argc), &argv[0])
	SetPolicy(PolicyQuit, QuitPolicyLastWindowClosed)
}

var funcMap *_funcMap = newFuncMap()

type _funcMap struct {
	index int
	mu    sync.Mutex
	funcs map[int]func()
}

func newFuncMap() *_funcMap {
	fm := &_funcMap{}
	fm.funcs = make(map[int]func())
	return fm
}

func (p *_funcMap) Register(fn func()) int {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.index++
	p.funcs[p.index] = fn
	return p.index
}

func (p *_funcMap) Lookup(index int) func() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if fn, ok := p.funcs[index]; ok {
		delete(p.funcs, index)
		return fn
	}
	return nil
}

//export go_call_func_async
func go_call_func_async(index C.int) {
	if fn := funcMap.Lookup(int(index)); fn != nil {
		fn()
	}
}

func Async(fn func()) {
	index := funcMap.Register(fn)
	C.cgo_ecore_main_loop_thread_safe_call_async(unsafe.Pointer(&index))
}

//export go_call_func_sync
func go_call_func_sync(index C.int) {
	if fn := funcMap.Lookup(int(index)); fn != nil {
		fn()
	}
}

func Sync(fn func()) {
	index := funcMap.Register(fn)
	C.cgo_ecore_main_loop_thread_safe_call_sync(unsafe.Pointer(&index))
}

func Run() {
	C.elm_run()
	C.elm_shutdown()
}

func Exit() {
	C.elm_exit()
}

// -------------------------------------
// some ecore functions

func Restart() {
	C.ecore_app_restart()
}

// for power management
func Throttle() float64 {
	return float64(C.ecore_throttle_get())
}

func AdjustThrottle(t float64) {
	C.ecore_throttle_adjust(C.double(t))
}

func MemoryStateInfo() MemoryState {
	return MemoryState(C.ecore_memory_state_get())
}

func SetMemoryState(ms MemoryState) {
	C.ecore_memory_state_set(C.Ecore_Memory_State(ms))
}

func PowerStateInfo() PowerState {
	return PowerState(C.ecore_power_state_get())
}

func SetPowerState(ps PowerState) {
	C.ecore_power_state_set(C.Ecore_Power_State(ps))
}

// TODO : test it
type Job struct {
	obj *C.Ecore_Job
}

func wrapJob(o *C.Ecore_Job) *Job {
	if o != nil {
		return &Job{o}
	}
	return nil
}

//export go_call_ecore_callback_func
func go_call_ecore_callback_func(id C.int) {
	if f := registry.Lookup(int(id)); f != nil {
		if fn, ok := f.(func()); ok {
			fn()
		}
	}
}

//TODO: how we can delete id from map
// we omitted the data
func NewJob(fn func()) *Job {
	id := registry.Register(fn)
	o := C.cgo_ecore_job_add(unsafe.Pointer(&id))
	return wrapJob(o)
}

// if job is not executed, can be deleted
func (p *Job) Delete() {
	cid := C.ecore_job_del(p.obj)
	if cid != nil {
		id := int(*((*C.int)(cid)))
		registry.Delete(id)
	}
}

func DumpTimerEvents() string {
	return C.GoString(C.ecore_timer_dump())
}

//export go_ecore_task_callback_func
func go_ecore_task_callback_func(id C.int) C.Eina_Bool {
	if f := registry.Lookup(int(id)); f != nil {
		if fn, ok := f.(func() bool); ok {
			return eBool(fn())
		}
	}
	return eFalse
}

type Timer struct {
	obj *C.Ecore_Timer
}

func NewTimer(period float64, fn func() bool) *Timer {
	id := registry.Register(fn)
	o := C.cgo_ecore_timer_add(C.double(period), unsafe.Pointer(&id))
	return wrapTimer(o)
}

// This is the same as ecore_timer_add(), but "now" is the time from ecore_loop_time_get() not ecore_time_get() as ecore_timer_add() uses.
func NewLoopTimer(period float64, fn func() bool) *Timer {
	id := registry.Register(fn)
	o := C.cgo_ecore_timer_loop_add(C.double(period), unsafe.Pointer(&id))
	return wrapTimer(o)
}

func wrapTimer(o *C.Ecore_Timer) *Timer {
	if o != nil {
		return &Timer{o}
	}
	return nil
}

func (p *Timer) Delete() {
	cid := C.ecore_timer_del(p.obj)
	if cid != nil {
		id := int(*((*C.int)(cid)))
		registry.Delete(id)
	}
}

func (p *Timer) Freeze() {
	C.ecore_timer_freeze(p.obj)
}

func (p *Timer) IsFreezed() bool {
	return C.ecore_timer_freeze_get(p.obj) == eTrue
}

func (p *Timer) Thaw() {
	C.ecore_timer_thaw(p.obj)
}

func (p *Timer) Delay(add float64) {
	C.ecore_timer_delay(p.obj, C.double(add))
}

func (p *Timer) Reset() {
	C.ecore_timer_reset(p.obj)
}

func (p *Timer) Interval() float64 {
	return float64(C.ecore_timer_interval_get(p.obj))
}

func (p *Timer) SetInterval(ival float64) {
	C.ecore_timer_interval_set(p.obj, C.double(ival))
}

func (p *Timer) Pending() float64 {
	return float64(C.ecore_timer_pending_get(p.obj))
}

func TimerPrecision() float64 {
	return float64(C.ecore_timer_precision_get())
}

func SetTimerPrecision(p float64) {
	C.ecore_timer_precision_set(C.double(p))
}

/*
TODO: is this are nice values
void 	ecore_exe_run_priority_set (int pri)
int 	ecore_exe_run_priority_get (void)
*/

// -------------------------------------

func SetLanguage(lang string) {
	clang := C.CString(lang)
	defer free(clang)
	C.elm_language_set(clang)
}

func DataDir() string {
	return C.GoString(C.elm_app_data_dir_get())
}

// -------------------------------------

func LookupRenderMethod(name string) int {
	cname := C.CString(name)
	defer free(cname)
	return int(C.evas_render_method_lookup(cname))
}

func RenderMethods() []string {
	lst := C.evas_render_method_list()
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func ClearGlobalFontPath() {
	C.evas_font_path_global_clear()
}

func AppendToGlobalFontPath(path string) {
	cpath := C.CString(path)
	defer free(cpath)
	C.evas_font_path_global_append(cpath)
}

func PrependToGlobalFontPath(path string) {
	cpath := C.CString(path)
	defer free(cpath)
	C.evas_font_path_global_prepend(cpath)
}

func GlobalFontPaths() []string {
	lst := C.evas_font_path_global_list()
	return newListIterator(lst).StringSlice()
}

func SetBaseScale(scale float64) {
	C.elm_app_base_scale_set(C.double(scale))
}

func BaseScale() float64 {
	return float64(C.elm_app_base_scale_get())
}

func AppName() string {
	return C.GoString(C.elm_app_name_get())
}

func SetAppName(name string) {
	cname := C.CString(name)
	defer free(cname)
	C.elm_app_name_set(cname)
}

func SetDesktopFile(path string) {
	cpath := C.CString(path)
	defer free(cpath)
	C.elm_app_desktop_entry_set(cpath)
}

func DesktopFile() string {
	return C.GoString(C.elm_app_desktop_entry_get())
}

func ProcessState() ProcState {
	return ProcState(C.elm_process_state_get())
}

func SetPolicy(policy Policy, value PolicyVal) bool {
	return C.elm_policy_set(C.uint(policy), C.int(value)) == eTrue
}

func PolicyValue(policy Policy) PolicyVal {
	return PolicyVal(C.elm_policy_get(C.uint(policy)))
}

/*
TODO: is this already using
func SetQuicklaunchMode(b bool) {
	C.elm_quicklaunch_mode_set(eBool(b))
}

func IsQuicklaunchMode() bool {
	return C.elm_quicklaunch_mode_get() == eTrue
}

func ShutdownSubQuicklaunch() int {
	return int(C.elm_quicklaunch_sub_shutdown())
}

func ShutdownQuicklaunch() int {
	return int(C.elm_quicklaunch_shutdown())
}

func SeedQuicklaunch() {
	C.elm_quicklaunch_seed()
}

func CleanupQuicklaunch() {
	C.elm_quicklaunch_cleanup()
}
*/

/*
int 	elm_quicklaunch_init (int argc, char **argv)
int 	elm_quicklaunch_sub_init (int argc, char **argv)
int 	elm_quicklaunch_sub_shutdown (void)
Eina_Bool 	elm_quicklaunch_prepare (int argc, char **argv, const char *cwd)
Eina_Bool 	efl_quicklaunch_prepare (int argc, char **argv, const char *cwd)
Eina_Bool 	elm_quicklaunch_fork (int argc, char **argv, char *cwd, void(*postfork_func)(void *data), void *postfork_data)
int 	elm_quicklaunch_fallback (int argc, char **argv)
int 	efl_quicklaunch_fallback (int argc, char **argv)
char * 	elm_quicklaunch_exe_path_get (const char *exe, const char *cwd)


*/

func isKeyModifierSet(mod *C.Evas_Modifier, modifier string) bool {
	cmodifier := C.CString(modifier)
	defer free(cmodifier)
	return C.evas_key_modifier_is_set(mod, cmodifier) == eTrue
}

// -----------------------------------------------------

type Theme struct {
	obj *C.Elm_Theme
}

func NewTheme() *Theme {
	return &Theme{C.elm_theme_new()}
}

func wrapTheme(o *C.Elm_Theme) *Theme {
	if o != nil {
		return &Theme{o}
	}
	return nil
}

func DefaultTheme() *Theme {
	return wrapTheme(C.elm_theme_default_get())
}

func FullFlushTheme() {
	C.elm_theme_full_flush()
}

func ThemeAvailableNames() []string {
	lst := C.elm_theme_name_available_list_new()
	if lst != nil {
		defer C.elm_theme_name_available_list_free(lst)
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func SystemThemeDir() string {
	return C.GoString(C.elm_theme_system_dir_get())
}

func UserThemeDir() string {
	return C.GoString(C.elm_theme_user_dir_get())
}

//char * elm_theme_list_item_path_get (const char *f, Eina_Bool *in_search_path)

func (p *Theme) Free() {
	C.elm_theme_free(p.obj)
}

func (p *Theme) Copy(dest *Theme) {
	C.elm_theme_copy(p.obj, dest.obj)
}

func (p *Theme) SetRef(ref *Theme) {
	C.elm_theme_ref_set(p.obj, ref.obj)
}

func (p *Theme) Ref() *Theme {
	return wrapTheme(C.elm_theme_ref_get(p.obj))
}

func (p *Theme) AddOverlay(item string) {
	citem := C.CString(item)
	defer free(citem)
	C.elm_theme_overlay_add(p.obj, citem)
}

func (p *Theme) DelOverlay(item string) {
	citem := C.CString(item)
	defer free(citem)
	C.elm_theme_overlay_del(p.obj, citem)
}

func (p *Theme) Overlays() []string {
	lst := C.elm_theme_overlay_list_get(p.obj)
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func (p *Theme) AddExtension(item string) {
	citem := C.CString(item)
	defer free(citem)
	C.elm_theme_extension_add(p.obj, citem)
}

func (p *Theme) DelExtension(item string) {
	citem := C.CString(item)
	defer free(citem)
	C.elm_theme_extension_del(p.obj, citem)
}

func (p *Theme) Extensions() []string {
	lst := C.elm_theme_extension_list_get(p.obj)
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func (p *Theme) SetTheme(theme string) {
	ctheme := C.CString(theme)
	defer free(ctheme)
	C.elm_theme_set(p.obj, ctheme)
}

func (p *Theme) Theme() string {
	return C.GoString(C.elm_theme_get(p.obj))
}

func (p *Theme) Themes() []string {
	lst := C.elm_theme_list_get(p.obj)
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func (p *Theme) Flush() {
	C.elm_theme_flush(p.obj)
}

func (p *Theme) Data(key string) string {
	ckey := C.CString(key)
	defer free(ckey)
	return C.GoString(C.elm_theme_data_get(p.obj, ckey))
}

func (p *Theme) Groups(base string) []string {
	cbase := C.CString(base)
	defer free(cbase)
	lst := C.elm_theme_group_base_list(p.obj, cbase)
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

// -----------------------------------------------------

type FontProperty struct {
	obj *C.Elm_Font_Properties
}

func wrapFontProperty(o *C.Elm_Font_Properties) *FontProperty {
	if o != nil {
		return &FontProperty{o}
	}
	return nil
}

func (p *FontProperty) Name() string {
	return C.GoString(p.obj.name)
}

func (p *FontProperty) Styles() []string {
	lst := p.obj.styles
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func FontProperties(font string) *FontProperty {
	cfont := C.CString(font)
	defer free(cfont)
	cfp := C.elm_font_properties_get(cfont)
	fp := wrapFontProperty(cfp)
	C.elm_font_properties_free(cfp)
	return fp
}

// -----------------------------------------------------

type need struct{}

var Need = need{}

func (need) Efreet() bool {
	return C.elm_need_efreet() == eTrue
}

func (need) Systray() bool {
	return C.elm_need_systray() == eTrue
}

func (need) SysNotify() bool {
	return C.elm_need_sys_notify() == eTrue
}

func (need) Eldbus() bool {
	return C.elm_need_eldbus() == eTrue
}

func (need) Elocation() bool {
	return C.elm_need_elocation() == eTrue
}

func (need) Ethumb() bool {
	return C.elm_need_ethumb() == eTrue
}

func (need) Web() bool {
	return C.elm_need_web() == eTrue
}

// -----------------------------------------------------

type elmConfig struct{}

var Config = elmConfig{}

func (elmConfig) PreferredEngine() string {
	return C.GoString(C.elm_config_preferred_engine_get())
}

func (elmConfig) SetPreferredEngine(engine string) {
	cengine := C.CString(engine)
	defer free(cengine)
	C.elm_config_preferred_engine_set(cengine)
}

func (elmConfig) AccelPreference() string {
	return C.GoString(C.elm_config_accel_preference_get())
}

func (elmConfig) SetAccelPreference(pre string) {
	cpre := C.CString(pre)
	defer free(cpre)
	C.elm_config_accel_preference_set(cpre)
}

func (elmConfig) IsSelectionClearOnUnfocused() bool {
	return C.elm_config_selection_unfocused_clear_get() == eTrue
}

func (elmConfig) SetSelectionClearOnUnfocused(b bool) {
	C.elm_config_selection_unfocused_clear_set(eBool(b))
}

func (elmConfig) IsVsync() bool {
	return C.elm_config_vsync_get() == eTrue
}

func (elmConfig) SetVsync(b bool) {
	C.elm_config_vsync_set(eBool(b))
}

func (elmConfig) IsAccelPreferenceOverride() bool {
	return C.elm_config_accel_preference_override_get() == eTrue
}

func (elmConfig) SetAccelPreferenceOverride(b bool) {
	C.elm_config_accel_preference_override_set(eBool(b))
}

func (elmConfig) IsMirrored() bool {
	return C.elm_config_mirrored_get() == eTrue
}

func (elmConfig) SetMirrored(b bool) {
	C.elm_config_mirrored_set(eBool(b))
}

func (elmConfig) IsClouseauEnabled() bool {
	return C.elm_config_clouseau_enabled_get() == eTrue
}

func (elmConfig) SetClouseauEnabled(b bool) {
	C.elm_config_clouseau_enabled_set(eBool(b))
}

func (elmConfig) IndicatorService(rotation int) string {
	return C.GoString(C.elm_config_indicator_service_get(C.int(rotation)))
}

type ColorClass struct {
	obj *C.Elm_Color_Class
}

func wrapColorClass(o *C.Elm_Color_Class) *ColorClass {
	if o != nil {
		return &ColorClass{o}
	}
	return nil
}

func (p *ColorClass) Name() string {
	return C.GoString(p.obj.name)
}

func (p *ColorClass) Description() string {
	return C.GoString(p.obj.desc)
}

func (elmConfig) ColorClasses() []*ColorClass {
	lst := C.elm_config_color_classes_list_get()
	if lst != nil {
		defer C.elm_config_color_classes_list_free(lst)
		return newListIterator(lst).ColorClassSlice()
	}
	return nil
}

type ColorOverlayColor struct {
	r, g, b, a byte
}

func newColorOverlayColor(cr, cg, cb, ca C.uchar) ColorOverlayColor {
	return ColorOverlayColor{
		r: byte(cr),
		g: byte(cg),
		b: byte(cb),
		a: byte(ca),
	}
}

func (p ColorOverlayColor) R() byte {
	return p.r
}

func (p ColorOverlayColor) G() byte {
	return p.g
}

func (p ColorOverlayColor) B() byte {
	return p.b
}

func (p ColorOverlayColor) A() byte {
	return p.a
}

type ColorOverlay struct {
	obj *C.Elm_Color_Overlay
}

func wrapColorOverlay(o *C.Elm_Color_Overlay) *ColorOverlay {
	if o != nil {
		return &ColorOverlay{o}
	}
	return nil
}

func (p *ColorOverlay) ColorClass() string {
	return C.GoString(p.obj.color_class)
}

func (p *ColorOverlay) Color() ColorOverlayColor {
	return newColorOverlayColor(p.obj.color.r, p.obj.color.g, p.obj.color.b, p.obj.color.a)
}

func (p *ColorOverlay) Outline() ColorOverlayColor {
	return newColorOverlayColor(p.obj.outline.r, p.obj.outline.g, p.obj.outline.b, p.obj.outline.a)
}

func (p *ColorOverlay) Shadow() ColorOverlayColor {
	return newColorOverlayColor(p.obj.shadow.r, p.obj.shadow.g, p.obj.shadow.b, p.obj.shadow.a)
}

func (elmConfig) ColorOverlays() []*ColorOverlay {
	lst := C.elm_config_color_overlay_list_get()
	if lst != nil {
		return newListIterator(lst).ColorOverlaySlice()
	}
	return nil
}

func (elmConfig) SetColorOverlay(cls string, r, g, b, a, r2, g2, b2, a2, r3, g3, b3, a3 int) {
	ccls := C.CString(cls)
	defer free(ccls)
	C.elm_config_color_overlay_set(ccls, C.int(r), C.int(g), C.int(b), C.int(a), C.int(r2), C.int(g2), C.int(b2), C.int(a2), C.int(r3), C.int(g3), C.int(b3), C.int(a3))
}

func (elmConfig) UnsetOverlay(cls string) {
	ccls := C.CString(cls)
	defer free(ccls)
	C.elm_config_color_overlay_unset(ccls)
}

func (elmConfig) ApplyOverlayChanges() {
	C.elm_config_color_overlay_apply()
}

func (elmConfig) IsMagnifierEnabled() bool {
	return C.elm_config_magnifier_enable_get() == eTrue
}

func (elmConfig) SetMagnifierEnabled(b bool) {
	C.elm_config_magnifier_enable_set(eBool(b))
}

func (elmConfig) MagnifierScale() float64 {
	return float64(C.elm_config_magnifier_scale_get())
}
func (elmConfig) SetMagnifierScale(scale float64) {
	C.elm_config_magnifier_scale_set(C.double(scale))
}

func (elmConfig) IsAudioMute(channel EdjeChannel) bool {
	return C.elm_config_audio_mute_get(C.Edje_Channel(channel)) == eTrue
}

func (elmConfig) SetAudioMute(channel EdjeChannel, b bool) {
	C.elm_config_audio_mute_set(C.Edje_Channel(channel), eBool(b))
}

func (elmConfig) IsWindowAutoFocusEnabled() bool {
	return C.elm_config_window_auto_focus_enable_get() == eTrue
}

func (elmConfig) SetWindowAutoFocusEnable(b bool) {
	C.elm_config_window_auto_focus_enable_set(eBool(b))
}

func (elmConfig) IsWindowAutoFocusAnimated() bool {
	return C.elm_config_window_auto_focus_animate_get() == eTrue
}

func (elmConfig) SetWindowAutoFocusAnimate(b bool) {
	C.elm_config_window_auto_focus_animate_set(eBool(b))
}

func (elmConfig) IsPopupScrollable() bool {
	return C.elm_config_popup_scrollable_get() == eTrue
}

func (elmConfig) SetPopupScrollable(b bool) {
	C.elm_config_popup_scrollable_set(eBool(b))
}

func (elmConfig) SetTransitionDurationFactor(factor float64) {
	C.elm_config_transition_duration_factor_set(C.double(factor))
}

func (elmConfig) TransitionDurationFactor() float64 {
	return float64(C.elm_config_transition_duration_factor_get())
}

func (elmConfig) SetWebBackend(backend string) {
	cbackend := C.CString(backend)
	defer free(cbackend)
	C.elm_config_web_backend_set(cbackend)
}

func (elmConfig) WebBackend() string {
	return C.GoString(C.elm_config_web_backend_get())
}

func (elmConfig) Save() bool {
	return C.elm_config_save() == eTrue
}

func (elmConfig) Reload() {
	C.elm_config_reload()
}

func (elmConfig) FlushAll() {
	C.elm_config_all_flush()
}

func (elmConfig) FingerSize() int {
	return int(C.elm_config_finger_size_get())
}

func (elmConfig) SetFingerSize(size int) {
	C.elm_config_finger_size_set(C.Evas_Coord(size))
}

/*
This is kind of low level Elementary call, most useful on size evaluation times for widgets.
An external user wouldn't be calling, most of the time.

func (elmConfig) AdjustFingerSize(timesw, timesh int) (int,int) {
	var w,h C.Evas_Coord
	C.elm_coords_finger_size_adjust(C.int(timesw), &w, C.int(timesh), &h)
	return int(w), int(h)
}
*/

func (elmConfig) Profile() string {
	return C.GoString(C.elm_config_profile_get())
}

func (elmConfig) ProfileDir(profile string, is_user bool) string {
	var dir *C.char
	cprofile := C.CString(profile)
	defer free(cprofile, dir)
	dir = C.elm_config_profile_dir_get(cprofile, eBool(is_user))
	return C.GoString(dir)
}

func (elmConfig) Profiles(hiddens bool) []string {
	var lst *C.Eina_List
	if hiddens {
		lst = C.elm_config_profile_list_full_get()
	} else {
		lst = C.elm_config_profile_list_get()
	}

	if lst != nil {
		defer C.elm_config_profile_list_free(lst)
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func (elmConfig) SetProfile(profile string) {
	cprofile := C.CString(profile)
	defer free(cprofile)
	C.elm_config_profile_set(cprofile)
}

func (elmConfig) IsProfilesExists(profile string) bool {
	cprofile := C.CString(profile)
	defer free(cprofile)
	return C.elm_config_profile_exists(cprofile) == eTrue
}

func (elmConfig) AddDerivedProfile(profile, derive_options string) {
	cprofile := C.CString(profile)
	cderopts := C.CString(derive_options)
	defer free(cprofile, cderopts)
	C.elm_config_profile_derived_add(cprofile, cderopts)
}

func (elmConfig) DelDerivedProfile(profile string) {
	cprofile := C.CString(profile)
	defer free(cprofile)
	C.elm_config_profile_derived_del(cprofile)
}

func (elmConfig) SaveProfile(profile string) {
	cprofile := C.CString(profile)
	defer free(cprofile)
	C.elm_config_profile_save(cprofile)
}

func (elmConfig) LongpressTimeout() float64 {
	return float64(C.elm_config_longpress_timeout_get())
}

func (elmConfig) SetLongpressTimeout(timeout float64) {
	C.elm_config_longpress_timeout_set(C.double(timeout))
}

func (elmConfig) SetSoftcursorMode(mode SoftcursorMode) {
	C.elm_config_softcursor_mode_set(C.Elm_Softcursor_Mode(mode))
}

func (elmConfig) SoftcursorMode() SoftcursorMode {
	return SoftcursorMode(C.elm_config_softcursor_mode_get())
}

func (elmConfig) TooltipDelay() float64 {
	return float64(C.elm_config_tooltip_delay_get())
}

func (elmConfig) SetTooltipDelay(delay float64) {
	C.elm_config_tooltip_delay_set(C.double(delay))
}

func (elmConfig) IsOnlyEngineCursor() bool {
	return C.elm_config_cursor_engine_only_get() == eTrue
}

func (elmConfig) SetOnlyEngineCursor(b bool) {
	C.elm_config_cursor_engine_only_set(eBool(b))
}

func (elmConfig) Scale() float64 {
	return float64(C.elm_config_scale_get())
}

func (elmConfig) SetScale(scale float64) {
	C.elm_config_scale_set(C.double(scale))
}

func (elmConfig) IsContextMenuDisabled() bool {
	return C.elm_config_context_menu_disabled_get() == eTrue
}

func (elmConfig) SetContextMenuDisabled(b bool) {
	C.elm_config_context_menu_disabled_set(eBool(b))
}

func (elmConfig) IsAtspiMode() bool {
	return C.elm_config_atspi_mode_get() == eTrue
}

func (elmConfig) SetAtspiMode(b bool) {
	C.elm_config_atspi_mode_set(eBool(b))
}

func (elmConfig) IconTheme() string {
	return C.GoString(C.elm_config_icon_theme_get())
}

func (elmConfig) SetIconTheme(theme string) {
	ctheme := C.CString(theme)
	defer free(ctheme)
	C.elm_config_icon_theme_set(ctheme)
}

type TextClass struct {
	obj *C.Elm_Text_Class
}

func wrapTextClass(o *C.Elm_Text_Class) *TextClass {
	if o != nil {
		return &TextClass{o}
	}
	return nil
}

func (p *TextClass) Name() string {
	return C.GoString(p.obj.name)
}

func (p *TextClass) Description() string {
	return C.GoString(p.obj.desc)
}

func (elmConfig) TextClasses() []*TextClass {
	lst := C.elm_config_text_classes_list_get()
	if lst != nil {
		defer C.elm_config_text_classes_list_free(lst)
		return newListIterator(lst).TextClassSlice()
	}
	return nil
}

type FontOverlay struct {
	obj *C.Elm_Font_Overlay
}

func wrapFontOverlay(o *C.Elm_Font_Overlay) *FontOverlay {
	if o != nil {
		return &FontOverlay{o}
	}
	return nil
}

func (p *FontOverlay) TextClass() string {
	return C.GoString(p.obj.text_class)
}

func (p *FontOverlay) Font() string {
	return C.GoString(p.obj.font)
}

func (p *FontOverlay) FontSize() int {
	return int(p.obj.size)
}

func (elmConfig) FontOverlays() []*FontOverlay {
	lst := C.elm_config_font_overlay_list_get()
	if lst != nil {
		return newListIterator(lst).FontOverlaySlice()
	}
	return nil
}

func (elmConfig) SetFontOverlay(cls, font string, size int) {
	ccls := C.CString(cls)
	cfont := C.CString(font)
	defer free(ccls, cfont)
	C.elm_config_font_overlay_set(ccls, cfont, C.Evas_Font_Size(size))
}

func (elmConfig) UnsetFontOverlay(cls string) {
	ccls := C.CString(cls)
	defer free(ccls)
	C.elm_config_font_overlay_unset(ccls)
}

func (elmConfig) ApplyFontOverlay() {
	C.elm_config_font_overlay_apply()
}

//in efl defined as void elm_config_font_hint_type_set(int type)
func (elmConfig) SetFontHintType(fh FontHinting) {
	C.elm_config_font_hint_type_set(C.int(fh))
}

func (elmConfig) IsAccessEnabled() bool {
	return C.elm_config_access_get() == eTrue
}

func (elmConfig) SetAccessEnabled(b bool) {
	C.elm_config_access_set(eBool(b))
}

func (elmConfig) IsSelectionUnfocusedClear() bool {
	return C.elm_config_selection_unfocused_clear_get() == eTrue
}

func (elmConfig) SetSelectionUnfocusedClear(b bool) {
	C.elm_config_selection_unfocused_clear_set(eBool(b))
}

// -----------------------------------------------------

type passwordShowLastConfig struct{}

var PasswordShowLastConfig = passwordShowLastConfig{}

func (passwordShowLastConfig) IsEnabled() bool {
	return C.elm_config_password_show_last_get() == eTrue
}

func (passwordShowLastConfig) SetEnabled(b bool) {
	C.elm_config_password_show_last_set(eBool(b))
}

func (passwordShowLastConfig) Timeout() float64 {
	return float64(C.elm_config_password_show_last_timeout_get())
}

func (passwordShowLastConfig) SetTimeout(timeout float64) {
	C.elm_config_password_show_last_timeout_set(C.double(timeout))
}

// -----------------------------------------------------

type scrollConfig struct{}

var ScrollConfig = scrollConfig{}

func (scrollConfig) IsBounceEnabled() bool {
	return C.elm_config_scroll_bounce_enabled_get() == eTrue
}

func (scrollConfig) SetBounceEnabled(b bool) {
	C.elm_config_scroll_bounce_enabled_set(eBool(b))
}

func (scrollConfig) BounceFriction() float64 {
	return float64(C.elm_config_scroll_bounce_friction_get())
}

func (scrollConfig) SetBounceFriction(bf float64) {
	C.elm_config_scroll_bounce_friction_set(C.double(bf))
}

func (scrollConfig) PageScrollFriction() float64 {
	return float64(C.elm_config_scroll_page_scroll_friction_get())
}

func (scrollConfig) SetPageScrollFriction(friction float64) {
	C.elm_config_scroll_page_scroll_friction_set(C.double(friction))
}

func (scrollConfig) BringInScrollFriction() float64 {
	return float64(C.elm_config_scroll_bring_in_scroll_friction_get())
}

func (scrollConfig) SetBringInScrollFriction(friction float64) {
	C.elm_config_scroll_bring_in_scroll_friction_set(C.double(friction))
}

func (scrollConfig) ZoomFriction() float64 {
	return float64(C.elm_config_scroll_zoom_friction_get())
}

func (scrollConfig) SetZoomFriction(friction float64) {
	C.elm_config_scroll_zoom_friction_set(C.double(friction))
}

func (scrollConfig) AccelFactor() float64 {
	return float64(C.elm_config_scroll_accel_factor_get())
}

func (scrollConfig) SetAccelFactor(factor float64) {
	C.elm_config_scroll_accel_factor_set(C.double(factor))
}

func (scrollConfig) IsAnimationDisabled() bool {
	return C.elm_config_scroll_animation_disabled_get() == eTrue
}

func (scrollConfig) SetAnimationDisabled(b bool) {
	C.elm_config_scroll_animation_disabled_set(eBool(b))
}

// -----------------------------------------------------

type elmScrollThumbscrollConfig struct{}

var ScrollThumbScrollConfig = elmScrollThumbscrollConfig{}

func (elmScrollThumbscrollConfig) IsEnabled() bool {
	return C.elm_config_scroll_thumbscroll_enabled_get() == eTrue
}

func (elmScrollThumbscrollConfig) SetEnabled(b bool) {
	C.elm_config_scroll_thumbscroll_enabled_set(eBool(b))
}

func (elmScrollThumbscrollConfig) Threshold() uint {
	return uint(C.elm_config_scroll_thumbscroll_threshold_get())
}

func (elmScrollThumbscrollConfig) SetThreshold(t uint) {
	C.elm_config_scroll_thumbscroll_threshold_set(C.uint(t))
}

func (elmScrollThumbscrollConfig) HoldThreshold() uint {
	return uint(C.elm_config_scroll_thumbscroll_hold_threshold_get())
}

func (elmScrollThumbscrollConfig) SetHoldThreshold(t uint) {
	C.elm_config_scroll_thumbscroll_hold_threshold_set(C.uint(t))
}

func (elmScrollThumbscrollConfig) MomentumThreshold() float64 {
	return float64(C.elm_config_scroll_thumbscroll_momentum_threshold_get())
}

func (elmScrollThumbscrollConfig) SetMomentumThreshold(t float64) {
	C.elm_config_scroll_thumbscroll_momentum_threshold_set(C.double(t))
}

func (elmScrollThumbscrollConfig) FlickDistanceTolerance() uint {
	return uint(C.elm_config_scroll_thumbscroll_flick_distance_tolerance_get())
}

func (elmScrollThumbscrollConfig) SetFlickDistanceTolerance(t uint) {
	C.elm_config_scroll_thumbscroll_flick_distance_tolerance_set(C.uint(t))
}

func (elmScrollThumbscrollConfig) Friction() float64 {
	return float64(C.elm_config_scroll_thumbscroll_friction_get())
}

func (elmScrollThumbscrollConfig) SetFriction(f float64) {
	C.elm_config_scroll_thumbscroll_friction_set(C.double(f))
}

func (elmScrollThumbscrollConfig) MinFriction() float64 {
	return float64(C.elm_config_scroll_thumbscroll_min_friction_get())
}

func (elmScrollThumbscrollConfig) SetMinFriction(f float64) {
	C.elm_config_scroll_thumbscroll_min_friction_set(C.double(f))
}

func (elmScrollThumbscrollConfig) StandardFriction() float64 {
	return float64(C.elm_config_scroll_thumbscroll_friction_standard_get())
}

func (elmScrollThumbscrollConfig) SetStandardFriction(f float64) {
	C.elm_config_scroll_thumbscroll_friction_standard_set(C.double(f))
}

func (elmScrollThumbscrollConfig) BorderFriction() float64 {
	return float64(C.elm_config_scroll_thumbscroll_border_friction_get())
}

func (elmScrollThumbscrollConfig) SetBorderFriction(f float64) {
	C.elm_config_scroll_thumbscroll_border_friction_set(C.double(f))
}

func (elmScrollThumbscrollConfig) SensitivityFriction() float64 {
	return float64(C.elm_config_scroll_thumbscroll_sensitivity_friction_get())
}

func (elmScrollThumbscrollConfig) SetSensitivityFriction(f float64) {
	C.elm_config_scroll_thumbscroll_sensitivity_friction_set(C.double(f))
}

func (elmScrollThumbscrollConfig) AccelerationThreshold() float64 {
	return float64(C.elm_config_scroll_thumbscroll_acceleration_threshold_get())
}

func (elmScrollThumbscrollConfig) SetAccelerationThreshold(t float64) {
	C.elm_config_scroll_thumbscroll_acceleration_threshold_set(C.double(t))
}

func (elmScrollThumbscrollConfig) AccelerationTimeLimit() float64 {
	return float64(C.elm_config_scroll_thumbscroll_acceleration_time_limit_get())
}

func (elmScrollThumbscrollConfig) SetAccelerationTimeLimit(l float64) {
	C.elm_config_scroll_thumbscroll_acceleration_time_limit_set(C.double(l))
}

func (elmScrollThumbscrollConfig) AccelerationWeight() float64 {
	return float64(C.elm_config_scroll_thumbscroll_acceleration_weight_get())
}

func (elmScrollThumbscrollConfig) SetAccelerationWeight(w float64) {
	C.elm_config_scroll_thumbscroll_acceleration_weight_set(C.double(w))
}

func (elmScrollThumbscrollConfig) IsSmoothStart() bool {
	return C.elm_config_scroll_thumbscroll_smooth_start_get() == eTrue
}

func (elmScrollThumbscrollConfig) SetSmoothStart(b bool) {
	C.elm_config_scroll_thumbscroll_smooth_start_set(eBool(b))
}

func (elmScrollThumbscrollConfig) SmoothAmount() float64 {
	return float64(C.elm_config_scroll_thumbscroll_smooth_amount_get())
}

func (elmScrollThumbscrollConfig) SetSmoothAmount(a float64) {
	C.elm_config_scroll_thumbscroll_smooth_amount_set(C.double(a))
}

func (elmScrollThumbscrollConfig) SmoothTimeWindow() float64 {
	return float64(C.elm_config_scroll_thumbscroll_smooth_time_window_get())
}

func (elmScrollThumbscrollConfig) SetSmoothTimeWindow(a float64) {
	C.elm_config_scroll_thumbscroll_smooth_time_window_set(C.double(a))
}

// -----------------------------------------------------

type cacheConfig struct{}

var CacheConfig = cacheConfig{}

func (cacheConfig) FlushAll() {
	C.elm_cache_all_flush()
}

func (cacheConfig) FlushInterval() int {
	return int(C.elm_config_cache_flush_interval_get())
}

func (cacheConfig) SetFlushInterval(ival int) {
	C.elm_config_cache_flush_interval_set(C.int(ival))
}

func (cacheConfig) IsFlushEnabled() bool {
	return C.elm_config_cache_flush_enabled_get() == eTrue
}

func (cacheConfig) SetFlushEnabled(b bool) {
	C.elm_config_cache_flush_enabled_set(eBool(b))
}

func (cacheConfig) FontCacheSize() int {
	return int(C.elm_config_cache_font_cache_size_get())
}

func (cacheConfig) SetFontCacheSize(size int) {
	C.elm_config_cache_font_cache_size_set(C.int(size))
}

func (cacheConfig) ImageCacheSize() int {
	return int(C.elm_config_cache_image_cache_size_get())
}

func (cacheConfig) SetImageCacheSize(size int) {
	C.elm_config_cache_image_cache_size_set(C.int(size))
}

func (cacheConfig) EdjeFileCacheSize() int {
	return int(C.elm_config_cache_edje_file_cache_size_get())
}

func (cacheConfig) SetEdjeFileCacheSize(size int) {
	C.elm_config_cache_edje_file_cache_size_set(C.int(size))
}

func (cacheConfig) EdjeCollectionCacheSize() int {
	return int(C.elm_config_cache_edje_collection_cache_size_get())
}

func (cacheConfig) SetEdjeCollectionCacheSize(size int) {
	C.elm_config_cache_edje_collection_cache_size_set(C.int(size))
}

// -----------------------------------------------------

type focusConfig struct{}

var FocusConfig = focusConfig{}

func (focusConfig) AutoscrollMode() FocusAutoscrollModeType {
	return FocusAutoscrollModeType(C.elm_config_focus_autoscroll_mode_get())
}

func (focusConfig) SetAutoscrollMode(mode FocusAutoscrollModeType) {
	C.elm_config_focus_autoscroll_mode_set(C.Elm_Focus_Autoscroll_Mode(mode))
}

func (focusConfig) IsHighlightEnabled() bool {
	return C.elm_config_focus_highlight_enabled_get() == eTrue
}

func (focusConfig) SetHighlightEnabled(b bool) {
	C.elm_config_focus_highlight_enabled_set(eBool(b))
}

func (focusConfig) IsHighlightAnimate() bool {
	return C.elm_config_focus_highlight_animate_get() == eTrue
}

func (focusConfig) SetHighlightAnimate(b bool) {
	C.elm_config_focus_highlight_animate_set(eBool(b))
}

func (focusConfig) IsHighlightClipDisabled() bool {
	return C.elm_config_focus_highlight_clip_disabled_get() == eTrue
}

func (focusConfig) SetHighlightClipDisabled(b bool) {
	C.elm_config_focus_highlight_clip_disabled_set(eBool(b))
}

func (focusConfig) MovePolicy() FocusMovePolicyType {
	return FocusMovePolicyType(C.elm_config_focus_move_policy_get())
}

func (focusConfig) SetMovePolicy(policy FocusMovePolicyType) {
	C.elm_config_focus_move_policy_set(C.Elm_Focus_Move_Policy(policy))
}

func (focusConfig) IsItemSelectOnFocusDisabled() bool {
	return C.elm_config_item_select_on_focus_disabled_get() == eTrue
}

func (focusConfig) SetItemSelectOnFocusDisabled(b bool) {
	C.elm_config_item_select_on_focus_disabled_set(eBool(b))
}

func (focusConfig) IsFirstItemFocusOnFirstFocusin() bool {
	return C.elm_config_first_item_focus_on_first_focusin_get() == eTrue
}

func (focusConfig) SetFirstItemFocusOnFirstFocusin(b bool) {
	C.elm_config_first_item_focus_on_first_focusin_set(eBool(b))
}

type gestureLayerConfig struct{}

var GenstureLayerConfig = gestureLayerConfig{}

func (gestureLayerConfig) LongTapStartTimeout() float64 {
	return float64(C.elm_config_glayer_long_tap_start_timeout_get())
}

func (gestureLayerConfig) SetLongTapStartTimeout(t float64) {
	C.elm_config_glayer_long_tap_start_timeout_set(C.double(t))
}

func (gestureLayerConfig) DoubleTimeout() float64 {
	return float64(C.elm_config_glayer_double_tap_timeout_get())
}

func (gestureLayerConfig) SetDoubleTimeout(t float64) {
	C.elm_config_glayer_double_tap_timeout_set(C.double(t))
}

// -----------------------------------------------------

type GestureLayer struct {
	*widgetBase
}

func NewGestureLayer(w Widget) *GestureLayer {
	return wrapGestureLayer(C.elm_gesture_layer_add(w.eo()))
}

func wrapGestureLayer(o *C.Eo) *GestureLayer {
	if o != nil {
		return &GestureLayer{wrapWidgetBase(o)}
	}
	return nil
}

func (p *GestureLayer) SetZoomStep(step float64) {
	C.elm_gesture_layer_zoom_step_set(p.obj, C.double(step))
}

func (p *GestureLayer) ZoomStep() float64 {
	return float64(C.elm_gesture_layer_zoom_step_get(p.obj))
}

func (p *GestureLayer) SetTapFingerSize(size int) {
	C.elm_gesture_layer_tap_finger_size_set(p.obj, C.Evas_Coord(size))
}

func (p *GestureLayer) TapFingerSize() int {
	return int(C.elm_gesture_layer_tap_finger_size_get(p.obj))
}

func (p *GestureLayer) SetHoldEvents(b bool) {
	C.elm_gesture_layer_hold_events_set(p.obj, eBool(b))
}

func (p *GestureLayer) IsHoldEvents() bool {
	return C.elm_gesture_layer_hold_events_get(p.obj) == eTrue
}

func (p *GestureLayer) SetRotateStep(step float64) {
	C.elm_gesture_layer_rotate_step_set(p.obj, C.double(step))
}

func (p *GestureLayer) RotateStep() float64 {
	return float64(C.elm_gesture_layer_rotate_step_get(p.obj))
}

func (p *GestureLayer) Attach(target Canvas) {
	C.elm_gesture_layer_attach(p.obj, target.eo())
}

func (p *GestureLayer) SetLineMinLength(l int) {
	C.elm_gesture_layer_line_min_length_set(p.obj, C.int(l))
}

func (p *GestureLayer) LineMinLength() int {
	return int(C.elm_gesture_layer_line_min_length_get(p.obj))
}

func (p *GestureLayer) SetZoomDistanceTolerance(t int) {
	C.elm_gesture_layer_zoom_distance_tolerance_set(p.obj, C.Evas_Coord(t))
}

func (p *GestureLayer) ZoomDistanceTolerance() int {
	return int(C.elm_gesture_layer_zoom_distance_tolerance_get(p.obj))
}

func (p *GestureLayer) SetLineDistanceTolerance(t int) {
	C.elm_gesture_layer_line_distance_tolerance_set(p.obj, C.Evas_Coord(t))
}

func (p *GestureLayer) LineDistanceTolerance() int {
	return int(C.elm_gesture_layer_line_distance_tolerance_get(p.obj))
}

func (p *GestureLayer) SetLineAngularTolerance(t float64) {
	C.elm_gesture_layer_line_angular_tolerance_set(p.obj, C.double(t))
}

func (p *GestureLayer) LineAngularTolerance() float64 {
	return float64(C.elm_gesture_layer_line_angular_tolerance_get(p.obj))
}

func (p *GestureLayer) SetZoomWheelFactor(f float64) {
	C.elm_gesture_layer_zoom_wheel_factor_set(p.obj, C.double(f))
}

func (p *GestureLayer) ZoomWheelFactor() float64 {
	return float64(C.elm_gesture_layer_zoom_wheel_factor_get(p.obj))
}

func (p *GestureLayer) SetZoomFingerFactor(f float64) {
	C.elm_gesture_layer_zoom_finger_factor_set(p.obj, C.double(f))
}

func (p *GestureLayer) ZoomFingerFactor() float64 {
	return float64(C.elm_gesture_layer_zoom_finger_factor_get(p.obj))
}

func (p *GestureLayer) SetRotateAngularTolerance(t float64) {
	C.elm_gesture_layer_rotate_angular_tolerance_set(p.obj, C.double(t))
}

func (p *GestureLayer) RotateAngularTolerance() float64 {
	return float64(C.elm_gesture_layer_rotate_angular_tolerance_get(p.obj))
}

func (p *GestureLayer) SetFlickTimeLimit(ms uint) {
	C.elm_gesture_layer_flick_time_limit_ms_set(p.obj, C.uint(ms))
}

func (p *GestureLayer) FlickTimeLimit() uint {
	return uint(C.elm_gesture_layer_flick_time_limit_ms_get(p.obj))
}

func (p *GestureLayer) SetLongTapStartTimeout(t float64) {
	C.elm_gesture_layer_long_tap_start_timeout_set(p.obj, C.double(t))
}

func (p *GestureLayer) LongTapStartTimeout() float64 {
	return float64(C.elm_gesture_layer_long_tap_start_timeout_get(p.obj))
}

func (p *GestureLayer) SetContinuesEnabled(b bool) {
	C.elm_gesture_layer_continues_enable_set(p.obj, eBool(b))
}

func (p *GestureLayer) IsContinuesEnabled() bool {
	return C.elm_gesture_layer_continues_enable_get(p.obj) == eTrue
}

func (p *GestureLayer) SetDoubleTapTimeout(t float64) {
	C.elm_gesture_layer_double_tap_timeout_set(p.obj, C.double(t))
}

func (p *GestureLayer) DoubleTapTimeout() float64 {
	return float64(C.elm_gesture_layer_double_tap_timeout_get(p.obj))
}

/*
TODO:
void 	elm_gesture_layer_cb_set (Elm_Gesture_Layer *obj, Elm_Gesture_Type idx, Elm_Gesture_State cb_type, Elm_Gesture_Event_Cb cb, void *data)
void 	elm_gesture_layer_cb_del (Elm_Gesture_Layer *obj, Elm_Gesture_Type idx, Elm_Gesture_State cb_type, Elm_Gesture_Event_Cb cb, void *data)
void 	elm_gesture_layer_cb_add (Elm_Gesture_Layer *obj, Elm_Gesture_Type idx, Elm_Gesture_State cb_type, Elm_Gesture_Event_Cb cb, void *data)
void 	elm_gesture_layer_tap_longpress_cb_add (Evas_Object *obj, Elm_Gesture_State state, Elm_Gesture_Event_Cb cb, void *data)
void 	elm_gesture_layer_tap_longpress_cb_del (Evas_Object *obj, Elm_Gesture_State state, Elm_Gesture_Event_Cb cb, void *data)
*/

// -----------------------------------------------------

type SelectionData struct {
	obj *C.Elm_Selection_Data
}

func wrapSelectionData(o *C.Elm_Selection_Data) *SelectionData {
	if o != nil {
		return &SelectionData{o}
	}
	return nil
}

func (p *SelectionData) X() int {
	return int(p.obj.x)
}

func (p *SelectionData) Y() int {
	return int(p.obj.y)
}

func (p *SelectionData) Format() SelectionFormat {
	return SelectionFormat(p.obj.format)
}

/*
TODO: p.obj.data type is void *, but generally used as char *
*/
func (p *SelectionData) Data() string {
	return C.GoString((*C.char)(p.obj.data))
}

/* we dont need len in go side
func (p *SelectionData) Len() int {
	return int(p.obj.len)
}
*/

func (p *SelectionData) Action() XDndAction {
	return XDndAction(p.obj.action)
}

// -----------------------------------------------------

type DragUserInfo struct {
	obj *C.Elm_Drag_User_Info
}

func wrapDragUserInfo(o *C.Elm_Drag_User_Info) *DragUserInfo {
	if o != nil {
		return &DragUserInfo{o}
	}
	return nil
}

func (p *DragUserInfo) Format() SelectionFormat {
	return SelectionFormat(p.obj.format)
}

func (p *DragUserInfo) Data() string {
	return C.GoString(p.obj.data)
}

func (p *DragUserInfo) Icons() []Object {
	if p.obj.icons != nil {
		return newListIterator(p.obj.icons).ObjectSlice()
	}
	return nil
}

func (p *DragUserInfo) Action() XDndAction {
	return XDndAction(p.obj.action)
}

/*
func (p *DragInfo) DragIconCreateHandler() DragIconCreateHandler {

}
*/

/*
Elm_Drag_User_Info

Elm_Sel_Format 	format
const char * 	data
Eina_List * 	icons
Elm_Xdnd_Action 	action
Elm_Drag_Icon_Create_Cb 	createicon
void * 	createdata
Elm_Drag_Start 	dragstart
void * 	startcbdata
Elm_Drag_Pos 	dragpos
void * 	dragdata
Elm_Drag_Accept 	acceptcb
void * 	acceptdata
Elm_Drag_Done 	dragdone
void * 	donecbdata
*/

type (
	// for use in elm_cnp_selection_get, defined as Elm_Drop_Cb but clarification we define another interface
	SelectionHandler interface {
		Handle(Widget, *SelectionData) bool
	}

	DropHandler interface {
		Handle(Widget, SelectionData)
	}

	SelectionLossHandler interface {
		Handle(SelectionType)
	}

	// TODO: is input Object is win type
	DragIconCreateHandler interface {
		Handle(Widget) (Object, int, int)
	}

	DragStateHandler interface {
		Handle(Widget)
	}

	DragDoneHandler interface {
		Handle(Widget, bool)
	}

	DragAcceptHandler interface {
		Handle(Widget, bool)
	}

	DragPosHandler interface {
		Handle(Widget, int, int, XDndAction)
	}

	DragStartHandler interface {
		Handle(Widget)
	}

	DragItemContainerPosHandler interface {
		Handle(Widget, WidgetItem, int, int, int, int, XDndAction)
	}

	DropItemContainerHandler interface {
		Handle(Widget, WidgetItem, SelectionData, int, int)
	}
)

// -----------------------------------------------------

//export go_elm_selection_cb_func
func go_elm_selection_cb_func(id C.int, obj *C.Eo, ev *C.Elm_Selection_Data) C.Eina_Bool {
	x := registry.Lookup(int(id))
	if x != nil {
		if h, ok := x.(SelectionHandler); ok {
			return eBool(h.Handle(wrapWidgetBase(obj), wrapSelectionData(ev)))
		}
	}
	return eFalse
}

//export go_elm_selection_loss_cb_func
func go_elm_selection_loss_cb_func(id C.int, t C.Elm_Sel_Type) {
	x := registry.Lookup(int(id))
	if x != nil {
		if h, ok := x.(SelectionLossHandler); ok {
			h.Handle(SelectionType(t))
		}
	}
}

//export go_elm_xy_item_get_cb_func
func go_elm_xy_item_get_cb_func(obj *C.Eo, x C.Evas_Coord, y C.Evas_Coord, xposret *C.int, yposret *C.int) *C.Elm_Widget_Item {
	//TODO test this...

	if obj != nil {
		var w interface{} = wrapWidgetBase(obj)

		if g, ok := w.(interface {
			ItemAtXY(int, int) (WidgetItem, int, int)
		}); ok { // this is Gengrid
			it, _, _ := g.ItemAtXY(int(x), int(y))
			return it.eo()
		}

		if l, ok := w.(interface {
			ItemAtXY(int, int) (WidgetItem, int)
		}); ok { // this is Genlist or List
			it, _ := l.ItemAtXY(int(x), int(y))
			return it.eo()
		}
	}
	return nil
}
