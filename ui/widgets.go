package ui

/*
#include "bridge.h"
*/
import "C"

import (
	"time"
	"unsafe"
)

//------------------------------------------------------------

type Window struct {
	*widgetBase
}

var _ Widget = &Window{}

func NewWindow(name, title string) *Window {
	cname := C.CString(name)
	ctitle := C.CString(title)
	defer free(cname, ctitle)

	eo := C.elm_win_util_standard_add(cname, ctitle)
	return &Window{wrapWidgetBase(eo)}
}

//TODO: is this parent a window,
func NewDialog(parent Object, name string, title string) *Window {
	cname := C.CString(name)
	ctitle := C.CString(title)
	defer free(cname, ctitle)

	eo := C.elm_win_util_dialog_add(parent.eo(), cname, ctitle)
	return &Window{wrapWidgetBase(eo)}
}

func NewWindowWithType(parent Object, name string, typ WindowType) *Window {
	cname := C.CString(name)
	defer free(cname)
	eo := C.elm_win_add(parent.eo(), cname, C.Elm_Win_Type(typ))
	return &Window{wrapWidgetBase(eo)}
}

func (p *Window) SetAutodel(b bool) {
	C.elm_win_autodel_set(p.obj, eBool(b))
}

func (p *Window) IsAutodel() bool {
	return C.elm_win_autodel_get(p.obj) == eTrue
}

func (p *Window) AddResizeObject(o Widget) {
	C.elm_win_resize_object_add(p.obj, o.eo())
}

func (p *Window) DelResizeObject(o Widget) {
	C.elm_win_resize_object_del(p.obj, o.eo())
}

func (p *Window) SetKeyboardWindow(b bool) {
	C.elm_win_keyboard_win_set(p.obj, eBool(b))
}

func (p *Window) IsKeyboardWindow() bool {
	return C.elm_win_keyboard_win_get(p.obj) == eTrue
}

func (p *Window) SetAutohide(b bool) {
	C.elm_win_autohide_set(p.obj, eBool(b))
}

func (p *Window) IsAutoHide() bool {
	return C.elm_win_autohide_get(p.obj) == eTrue
}

func (p *Window) SetOverride(b bool) {
	C.elm_win_override_set(p.obj, eBool(b))
}

func (p *Window) IsOverride() bool {
	return C.elm_win_override_get(p.obj) == eTrue
}

func (p *Window) SetIconified(b bool) {
	C.elm_win_iconified_set(p.obj, eBool(b))
}

func (p *Window) IsIconified() bool {
	return C.elm_win_iconified_get(p.obj) == eTrue
}

func (p *Window) SetQuickpanelZone(zone int) {
	C.elm_win_quickpanel_zone_set(p.obj, C.int(zone))
}

func (p *Window) QuickpanelZone() int {
	return int(C.elm_win_quickpanel_zone_get(p.obj))
}

func (p *Window) SetMaximized(b bool) {
	C.elm_win_maximized_set(p.obj, eBool(b))
}

func (p *Window) IsMaximized() bool {
	return C.elm_win_maximized_get(p.obj) == eTrue
}

func (p *Window) SetModal(b bool) {
	C.elm_win_modal_set(p.obj, eBool(b))
}

func (p *Window) IsModal() bool {
	return C.elm_win_modal_get(p.obj) == eTrue
}

func (p *Window) SetIconName(name string) {
	cname := C.CString(name)
	defer free(cname)
	C.elm_win_icon_name_set(p.obj, cname)
}

func (p *Window) IconName() string {
	return C.GoString(C.elm_win_icon_name_get(p.obj))
}

func (p *Window) SetWithdrawn(b bool) {
	C.elm_win_withdrawn_set(p.obj, eBool(b))
}

func (p *Window) IsWithdrawn() bool {
	return C.elm_win_withdrawn_get(p.obj) == eTrue
}

func (p *Window) SetRole(role string) {
	crole := C.CString(role)
	defer free(crole)
	C.elm_win_role_set(p.obj, crole)
}

func (p *Window) Role() string {
	return C.GoString(C.elm_win_role_get(p.obj))
}

func (p *Window) SetStepSize(w, h int) {
	C.elm_win_size_step_set(p.obj, C.int(w), C.int(h))
}

func (p *Window) StepSize() (int, int) {
	var w, h C.int
	C.elm_win_size_step_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Window) SetFocusHighlightStyle(style string) bool {
	cstyle := C.CString(style)
	defer free(cstyle)
	C.elm_win_focus_highlight_style_set(p.obj, cstyle)
	return true // for inheritance compatibility
}

func (p *Window) FocusHighlightStyle() string {
	return C.GoString(C.elm_win_focus_highlight_style_get(p.obj))
}

func (p *Window) SetBorderless(b bool) {
	C.elm_win_borderless_set(p.obj, eBool(b))
}

func (p *Window) IsBorderless() bool {
	return C.elm_win_borderless_get(p.obj) == eTrue
}

func (p *Window) SetScreenConstrain(b bool) {
	C.elm_win_screen_constrain_set(p.obj, eBool(b))
}

func (p *Window) IsScreenConstrain() bool {
	return C.elm_win_screen_constrain_get(p.obj) == eTrue
}

func (p *Window) SetFocusHighlightEnabled(b bool) {
	C.elm_win_focus_highlight_enabled_set(p.obj, eBool(b))
}

func (p *Window) IsFocusHighlightEnabled() bool {
	return C.elm_win_focus_highlight_enabled_get(p.obj) == eTrue
}

func (p *Window) SetTitle(title string) {
	ctitle := C.CString(title)
	defer free(ctitle)
	C.elm_win_title_set(p.obj, ctitle)
}

func (p *Window) Title() string {
	return C.GoString(C.elm_win_title_get(p.obj))
}

func (p *Window) SetAlpha(b bool) {
	C.elm_win_alpha_set(p.obj, eBool(b))
}

func (p *Window) IsAlpha() bool {
	return C.elm_win_alpha_get(p.obj) == eTrue
}

func (p *Window) SetUrgent(b bool) {
	C.elm_win_urgent_set(p.obj, eBool(b))
}

func (p *Window) IsUrgent() bool {
	return C.elm_win_urgent_get(p.obj) == eTrue
}

func (p *Window) SetRotation(r int) {
	C.elm_win_rotation_set(p.obj, C.int(r))
}

func (p *Window) Rotation() int {
	return int(C.elm_win_rotation_get(p.obj))
}

func (p *Window) SetQuickpanelMinorPriority(pri int) {
	C.elm_win_quickpanel_priority_minor_set(p.obj, C.int(pri))
}

func (p *Window) QuickpanelMinorPriority() int {
	return int(C.elm_win_quickpanel_priority_minor_get(p.obj))
}

func (p *Window) SetQuickpanelMajorPriority(pri int) {
	C.elm_win_quickpanel_priority_major_set(p.obj, C.int(pri))
}

func (p *Window) QuickpanelMajorPriority() int {
	return int(C.elm_win_quickpanel_priority_major_get(p.obj))
}

func (p *Window) SetSticky(b bool) {
	C.elm_win_sticky_set(p.obj, eBool(b))
}

func (p *Window) IsSticky() bool {
	return C.elm_win_sticky_get(p.obj) == eTrue
}

func (p *Window) SetFocusHighlightAnimate(b bool) {
	C.elm_win_focus_highlight_animate_set(p.obj, eBool(b))
}

func (p *Window) IsFocusHighlightAnimate() bool {
	return C.elm_win_focus_highlight_animate_get(p.obj) == eTrue
}

func (p *Window) SetAspect(a float64) {
	C.elm_win_aspect_set(p.obj, C.double(a))
}

func (p *Window) Aspect() float64 {
	return float64(C.elm_win_aspect_get(p.obj))
}

func (p *Window) SetDemandAttention(b bool) {
	C.elm_win_demand_attention_set(p.obj, eBool(b))
}

func (p *Window) IsDemandAttention() bool {
	return C.elm_win_demand_attention_get(p.obj) == eTrue
}

func (p *Window) SetLayer(layer int) {
	C.elm_win_layer_set(p.obj, C.int(layer))
}

func (p *Window) Layer() int {
	return int(C.elm_win_layer_get(p.obj))
}

func (p *Window) SetProfile(profile string) {
	cprofile := C.CString(profile)
	defer free(cprofile)
	C.elm_win_profile_set(p.obj, cprofile)
}

func (p *Window) Profile() string {
	return C.GoString(C.elm_win_profile_get(p.obj))
}

func (p *Window) SetShaped(b bool) {
	C.elm_win_shaped_set(p.obj, eBool(b))
}

func (p *Window) IsShaped() bool {
	return C.elm_win_shaped_get(p.obj) == eTrue
}

func (p *Window) SetFullscreen(b bool) {
	C.elm_win_fullscreen_set(p.obj, eBool(b))
}

func (p *Window) IsFullscreen() bool {
	return C.elm_win_fullscreen_get(p.obj) == eTrue
}

func (p *Window) SetConformant(b bool) {
	C.elm_win_conformant_set(p.obj, eBool(b))
}

func (p *Window) IsConformant() bool {
	return C.elm_win_conformant_get(p.obj) == eTrue
}

func (p *Window) SetBaseSize(w int, h int) {
	C.elm_win_size_base_set(p.obj, C.int(w), C.int(h))
}

func (p *Window) BaseSize() (int, int) {
	var w, h C.int
	C.elm_win_size_base_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Window) SetQuickpanel(b bool) {
	C.elm_win_quickpanel_set(p.obj, eBool(b))
}

func (p *Window) IsQuickpanel() bool {
	return C.elm_win_quickpanel_get(p.obj) == eTrue
}

func (p *Window) SetRotationWithResize(r int) {
	C.elm_win_rotation_with_resize_set(p.obj, C.int(r))
}

func (p *Window) SetSkipKeyboardFocus(b bool) {
	C.elm_win_prop_focus_skip_set(p.obj, eBool(b))
}

func (p *Window) ScreenPosition() (int, int) {
	var w, h C.int
	C.elm_win_screen_position_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Window) HasFocus() bool {
	return C.elm_win_focus_get(p.obj) == eTrue
}

func (p *Window) ScreenSize() (int, int, int, int) {
	var x, y, w, h C.int
	C.elm_win_screen_size_get(p.obj, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *Window) DPI() (int, int) {
	var xdpi, ydpi C.int
	C.elm_win_screen_dpi_get(p.obj, &xdpi, &ydpi)
	return int(xdpi), int(ydpi)
}

func (p *Window) SetNonblank(b bool) {
	C.elm_win_noblank_set(p.obj, eBool(b))
}

func (p *Window) IsNonblank() bool {
	return C.elm_win_noblank_get(p.obj) == eTrue
}

func (p *Window) Activate() {
	C.elm_win_activate(p.obj)
}

func (p *Window) Raise() {
	C.elm_win_raise(p.obj)
}

func (p *Window) Lower() {
	C.elm_win_lower(p.obj)
}

func (p *Window) SetCenter(hor, ver bool) {
	C.elm_win_center(p.obj, eBool(hor), eBool(ver))
}

func (p *Window) MainMenu() *Menu {
	return wrapMenu(C.elm_win_main_menu_get(p.obj))
}

/* since 1.9
func (p *Window) PreferredRotationValue() int {
	return int(C.elm_win_wm_rotation_preferred_rotation_get(p.obj))
}
*/

// FIXME: this functions wants image data
func (p *Window) SetIcon(icon *Icon) {
	C.elm_win_icon_object_set(p.obj, C.elm_image_object_get(icon.obj))
}

func (p *Window) Icon() *Icon {
	return wrapIcon(C.elm_win_icon_object_get(p.obj))
}

func (p *Window) SetKeyboardMode(mode WindowKeyboardMode) {
	C.elm_win_keyboard_mode_set(p.obj, C.Efl_Ui_Win_Keyboard_Mode(mode))
}

func (p *Window) KeyboardMode() WindowKeyboardMode {
	return WindowKeyboardMode(C.elm_win_keyboard_mode_get(p.obj))
}

func (p *Window) SetIndicatorOpacity(mode WindowIndicatorOpacityMode) {
	C.elm_win_indicator_opacity_set(p.obj, C.Elm_Win_Indicator_Opacity_Mode(mode))
}

func (p *Window) IndicatorOpacityMode() WindowIndicatorOpacityMode {
	return WindowIndicatorOpacityMode(C.elm_win_indicator_opacity_get(p.obj))
}

func (p *Window) SetIndicatorMode(mode WindowIndicatorMode) {
	C.elm_win_indicator_mode_set(p.obj, C.Elm_Win_Indicator_Mode(mode))
}

func (p *Window) IndicatorMode() WindowIndicatorMode {
	return WindowIndicatorMode(C.elm_win_indicator_mode_get(p.obj))
}

func (p *Window) SetKeygrab(key string, mode WindowKeygrabMode) {
	ckey := C.CString(key)
	defer free(ckey)
	//C.elm_win_keygrab_set(p.obj, ckey, 0, 0, 0, C.Efl_Ui_Win_Keygrab_Mode(mode))
	C.elm_win_keygrab_set(p.obj, ckey, 0, 0, 0, C.Elm_Win_Keygrab_Mode(mode))
}

func (p *Window) UnsetKeygrab(key string) {
	ckey := C.CString(key)
	defer free(ckey)
	C.elm_win_keygrab_unset(p.obj, ckey, 0, 0)
}

func (p *Window) ListenSocket(svcname string, svcnum int, svcsys bool) bool {
	csvcname := C.CString(svcname)
	defer free(csvcname)
	return C.elm_win_socket_listen(p.obj, csvcname, C.int(svcnum), eBool(svcsys)) == eTrue
}

func (p *Window) InlinedImageObject() Object {
	return wrapObjectBase(C.elm_win_inlined_image_object_get(p.obj))
}

// this is XWindows related, params paremeter not used,
func (p *Window) SendIllumeCommand(cmd IllumeCommand) {
	C.elm_win_illume_command_send(p.obj, C.Elm_Illume_Command(cmd), nil)
}

func (p *Window) IsWmRotationManualRotationDone() bool {
	return C.elm_win_wm_rotation_manual_rotation_done_get(p.obj) == eTrue
}

func (p *Window) SetWmRotationManualRotationDone(b bool) {
	C.elm_win_wm_rotation_manual_rotation_done_set(p.obj, eBool(b))
}

func (p *Window) IsWmRotationSupported() bool {
	return C.elm_win_wm_rotation_supported_get(p.obj) == eTrue
}

func (p *Window) WmAvailableRotations() ([]int, bool) {
	list := C.cgo_elm_win_wm_rotation_available_rotations_get(p.obj)
	if list == nil {
		return nil, false
	}

	slice := newListIterator(list).IntSlice()
	return slice, true
}

func (p *Window) SetWmAvailableRotations(rotations []int) {
	list := newIntList()
	for _, r := range rotations {
		list.Append(r)
	}
	C.cgo_elm_win_wm_rotation_available_rotations_set(p.obj, list.obj)
}

//------------------------------------------------------------

type SegmentControl struct {
	*Layout
}

var _ Container = &SegmentControl{}

func NewSegmentControl(parent Object) *SegmentControl {
	return &SegmentControl{wrapLayout(C.elm_segment_control_add(parent.eo()))}
}

func (p *SegmentControl) DelItemAt(idx int) {
	C.elm_segment_control_item_del_at(p.obj, C.int(idx))
}

func (p *SegmentControl) ItemCount() int {
	return int(C.elm_segment_control_item_count_get(p.obj))
}

func (p *SegmentControl) Item(idx int) *SegmentControlItem {
	return wrapSegmentControlItem(C.elm_segment_control_item_get(p.obj, C.int(idx)))
}

func (p *SegmentControl) ItemLabel(idx int) string {
	return C.GoString(C.elm_segment_control_item_label_get(p.obj, C.int(idx)))
}

func (p *SegmentControl) ItemIcon(idx int) *Icon {
	return wrapIcon(C.elm_segment_control_item_icon_get(p.obj, C.int(idx)))
}

func (p *SegmentControl) SelectedItem() *SegmentControlItem {
	return wrapSegmentControlItem(C.elm_segment_control_item_selected_get(p.obj))
}

func (p *SegmentControl) ItemFactory() *SegmentControlItemFactory {
	return NewSegmentControlItemFactory(p)
}

//------------------------------------------------------------

type Seperator struct {
	*Layout
}

func NewSeperator(parent Object) *Seperator {
	return &Seperator{wrapLayout(C.elm_separator_add(parent.eo()))}
}

func (p *Seperator) SetHorizontal(b bool) {
	C.elm_separator_horizontal_set(p.obj, eBool(b))
}

func (p *Seperator) IsHorizontal() bool {
	return C.elm_separator_horizontal_get(p.obj) == eTrue
}

//------------------------------------------------------------

type Check struct {
	*Layout
}

func NewCheck(parent Object) *Check {
	return &Check{wrapLayout(C.elm_check_add(parent.eo()))}
}

/* in 1.19
func (p *Check) SetSelected(b bool) {
	C.elm_check_selected_set(p.obj, eBool(b))
}

func (p *Check) IsSelected() bool {
	return C.elm_check_selected_get(p.obj) == eTrue
}
*/

func (p *Check) SetState(state bool) {
	C.elm_check_state_set(p.obj, eBool(state))
}

func (p *Check) State() bool {
	return C.elm_check_state_get(p.obj) == eTrue
}

func (p *Check) SetIcon(icon *Icon) {
	p.SetPartContent("icon", icon)
}

//------------------------------------------------------------

type CheckToggle struct {
	*Check
}

func NewCheckToggle(parent Object) *CheckToggle {
	chk := &CheckToggle{NewCheck(parent)}
	chk.SetStyle("toggle")
	return chk
}

func (p *CheckToggle) SetOnText(txt string) {
	p.SetPartText("on", txt)
}

func (p *CheckToggle) OnText() string {
	return p.PartText("on")
}

func (p *CheckToggle) SetOffText(txt string) {
	p.SetPartText("off", txt)
}

func (p *CheckToggle) OffText() string {
	return p.PartText("off")
}

//------------------------------------------------------------

type Bubble struct {
	*Layout
}

func NewBubble(o Object) *Bubble {
	eo := C.elm_bubble_add(o.eo())
	return &Bubble{wrapLayout(eo)}
}

func (p *Bubble) SetPos(pos BubblePos) {
	C.elm_bubble_pos_set(p.obj, C.Elm_Bubble_Pos(pos))
}

func (p *Bubble) Pos() BubblePos {
	return BubblePos(C.elm_bubble_pos_get(p.obj))
}

//------------------------------------------------------------

type Frame struct {
	*Layout
}

func NewFrame(o Object) *Frame {
	return &Frame{wrapLayout(C.elm_frame_add(o.eo()))}
}

func (p *Frame) SetAutocollapse(b bool) {
	C.elm_frame_autocollapse_set(p.obj, eBool(b))
}

func (p *Frame) IsAutocollapse() bool {
	return C.elm_frame_autocollapse_get(p.obj) == eTrue
}

func (p *Frame) SetCollapse(b bool) {
	C.elm_frame_collapse_set(p.obj, eBool(b))
}

func (p *Frame) IsCollapsed() bool {
	return C.elm_frame_collapse_get(p.obj) == eTrue
}

func (p *Frame) Collapse(b bool) {
	C.elm_frame_collapse_go(p.obj, eBool(b))
}

//------------------------------------------------------------

type Inwin struct {
	*Layout
}

func NewInwin(w *Window) *Inwin {
	return &Inwin{wrapLayout(C.elm_win_inwin_add(w.obj))}
}

func (p *Inwin) Activate() {
	C.elm_win_inwin_activate(p.obj)
}

func (p *Inwin) SetContent(o Object) {
	C.elm_win_inwin_content_set(p.obj, o.eo())
}

func (p *Inwin) Content() Object {
	return wrapObjectBase(C.elm_win_inwin_content_get(p.obj))
}

func (p *Inwin) UnsetContent() Object {
	return wrapObjectBase(C.elm_win_inwin_content_unset(p.obj))
}

//------------------------------------------------------------

type Radio struct {
	*Layout
}

func NewRadio(parent Object) *Radio {
	return &Radio{wrapLayout(C.elm_radio_add(parent.eo()))}
}

func wrapRadio(o *C.Eo) *Radio {
	if o != nil {
		return &Radio{wrapLayout(o)}
	}
	return nil
}

func (p *Radio) Add(r *Radio) {
	C.elm_radio_group_add(r.obj, p.obj)
}

func (p *Radio) SetValue(v int) {
	C.elm_radio_state_value_set(p.obj, C.int(v))
}

func (p *Radio) Value() int {
	return int(C.elm_radio_state_value_get(p.obj))
}

func (p *Radio) SetSelected(i int) {
	C.elm_radio_value_set(p.obj, C.int(i))
}

func (p *Radio) Selected() int {
	return int(C.elm_radio_value_get(p.obj))
}

func (p *Radio) SelectedItem() *Radio {
	return wrapRadio(C.elm_radio_selected_object_get(p.obj))
}

//void 	elm_radio_value_pointer_set (Elm_Radio *obj, int *valuep)

//------------------------------------------------------------

type Button struct {
	*Layout
}

func NewButton(parent Object) *Button {
	return &Button{wrapLayout(C.elm_button_add(parent.eo()))}
}

func wrapButton(o *C.Eo) *Button {
	if o != nil {
		return &Button{wrapLayout(o)}
	}
	return nil
}

func (p *Button) SetAutorepeatInitialTimeout(to float64) {
	C.elm_button_autorepeat_initial_timeout_set(p.obj, C.double(to))
}

func (p *Button) AutorepeatInitialTimeout() float64 {
	return float64(C.elm_button_autorepeat_initial_timeout_get(p.obj))
}

func (p *Button) SetAutorepeatGapTimeout(to float64) {
	C.elm_button_autorepeat_gap_timeout_set(p.obj, C.double(to))
}

func (p *Button) AutorepeatGapTimeout() float64 {
	return float64(C.elm_button_autorepeat_gap_timeout_get(p.obj))
}

func (p *Button) SetAutorepeat(b bool) {
	C.elm_button_autorepeat_set(p.obj, eBool(b))
}

func (p *Button) IsAutorepeat() bool {
	return C.elm_button_autorepeat_get(p.obj) == eTrue
}

//------------------------------------------------------------

type Notify struct {
	*widgetBase
}

func NewNotify(o Object) *Notify {
	return &Notify{wrapWidgetBase(C.elm_notify_add(o.eo()))}
}

func (p *Notify) SetAlign(hor, ver float64) {
	C.elm_notify_align_set(p.obj, C.double(hor), C.double(ver))
}

func (p *Notify) Align() (float64, float64) {
	var hor, ver C.double
	C.elm_notify_align_get(p.obj, &hor, &ver)
	return float64(hor), float64(ver)
}

func (p *Notify) SetAllowEvents(b bool) {
	C.elm_notify_allow_events_set(p.obj, eBool(b))
}

func (p *Notify) IsAllowEvents() bool {
	return C.elm_notify_allow_events_get(p.obj) == eTrue
}

func (p *Notify) SetTimeout(to float64) {
	C.elm_notify_timeout_set(p.obj, C.double(to))
}

func (p *Notify) Timeout() float64 {
	return float64(C.elm_notify_timeout_get(p.obj))
}

func (p *Notify) Dismiss() {
	C.elm_notify_dismiss(p.obj)
}

func (p *Notify) SetParent(parent Object) {
	C.elm_notify_parent_set(p.obj, parent.eo())
}

func (p *Notify) Parent() Widget {
	return wrapWidgetBase(C.elm_notify_parent_get(p.obj))
}

//------------------------------------------------------------

type Photo struct {
	*widgetBase
}

func NewPhoto(o Object) *Photo {
	return &Photo{wrapWidgetBase(C.elm_photo_add(o.eo()))}
}

func (p *Photo) SetFile(file string) bool {
	cfile := C.CString(file)
	defer free(cfile)
	return C.elm_photo_file_set(p.obj, cfile) == eTrue
}

func (p *Photo) SetEditable(b bool) {
	C.elm_photo_editable_set(p.obj, eBool(b))
}

func (p *Photo) IsEditable() bool {
	return C.elm_photo_editable_get(p.obj) == eTrue
}

func (p *Photo) SetFillInside(b bool) {
	C.elm_photo_fill_inside_set(p.obj, eBool(b))
}

func (p *Photo) IsFillInside() bool {
	return C.elm_photo_fill_inside_get(p.obj) == eTrue
}

func (p *Photo) SetFixedAspect(b bool) {
	C.elm_photo_aspect_fixed_set(p.obj, eBool(b))
}

func (p *Photo) IsFixedAspect() bool {
	return C.elm_photo_aspect_fixed_get(p.obj) == eTrue
}

func (p *Photo) SetPhotoSize(size int) {
	C.elm_photo_size_set(p.obj, C.int(size))
}

func (p *Photo) PhotoSize() int {
	return int(C.elm_photo_size_get(p.obj))
}

func (p *Photo) SetThumb(file, group string) {
	cfile := C.CString(file)
	cgroup := C.CString(group)
	defer free(cfile, cgroup)
	C.elm_photo_thumb_set(p.obj, cfile, cgroup)
}

//------------------------------------------------------------

type Slider struct {
	*widgetBase
}

func NewSlider(parent Object) *Slider {
	eo := C.elm_slider_add(parent.eo())
	return &Slider{wrapWidgetBase(eo)}
}

func (p *Slider) SetHorizontal(b bool) {
	C.elm_slider_horizontal_set(p.obj, eBool(b))
}

func (p *Slider) IsHorizontal() bool {
	return C.elm_slider_horizontal_get(p.obj) == eTrue
}

func (p *Slider) SetInverted(b bool) {
	C.elm_slider_inverted_set(p.obj, eBool(b))
}

func (p *Slider) IsInverted() bool {
	return C.elm_slider_inverted_get(p.obj) == eTrue
}

func (p *Slider) SetMinMax(min, max float64) {
	C.elm_slider_min_max_set(p.obj, C.double(min), C.double(max))
}

func (p *Slider) MinMax() (float64, float64) {
	var min, max C.double
	C.elm_slider_min_max_get(p.obj, &min, &max)
	return float64(min), float64(max)
}

func (p *Slider) SetValue(val float64) {
	C.elm_slider_value_set(p.obj, C.double(val))
}

func (p *Slider) Value() float64 {
	return float64(C.elm_slider_value_get(p.obj))
}

func (p *Slider) SetSpanSize(size int) {
	C.elm_slider_span_size_set(p.obj, C.Evas_Coord(size))
}

func (p *Slider) SpanSize() int {
	return int(C.elm_slider_span_size_get(p.obj))
}

func (p *Slider) SetUnitFormat(format string) {
	cformat := C.CString(format)
	defer free(cformat)
	C.elm_slider_unit_format_set(p.obj, cformat)
}

func (p *Slider) UnitFormat() string {
	return C.GoString(C.elm_slider_unit_format_get(p.obj))
}

func (p *Slider) SetIndicatorformat(format string) {
	cformat := C.CString(format)
	defer free(cformat)
	C.elm_slider_indicator_format_set(p.obj, cformat)
}

func (p *Slider) IndicatorFormat() string {
	return C.GoString(C.elm_slider_indicator_format_get(p.obj))
}

func (p *Slider) SetIndicatorVisible(b bool) {
	C.elm_slider_indicator_show_set(p.obj, eBool(b))
}

func (p *Slider) IsIndicatorVisible() bool {
	return C.elm_slider_indicator_show_get(p.obj) == eTrue
}

func (p *Slider) SetIndicatorVisibleMode(mode SliderIndicatorVisibleMode) {
	C.elm_slider_indicator_visible_mode_set(p.obj, C.Efl_Ui_Slider_Indicator_Visible_Mode(mode))
}

func (p *Slider) IndicatorVisibleMode() SliderIndicatorVisibleMode {
	return SliderIndicatorVisibleMode(C.elm_slider_indicator_visible_mode_get(p.obj))
}

func (p *Slider) SetIndicatorVisibleOnFocus(b bool) {
	C.elm_slider_indicator_show_on_focus_set(p.obj, eBool(b))
}

func (p *Slider) IsIndicatorVisibleOnFocus() bool {
	return C.elm_slider_indicator_show_on_focus_get(p.obj) == eTrue
}

func (p *Slider) SetRangeEnabled(b bool) {
	C.elm_slider_range_enabled_set(p.obj, eBool(b))
}

func (p *Slider) IsRangeEnabled() bool {
	return C.elm_slider_range_enabled_get(p.obj) == eTrue
}

func (p *Slider) SetRange(from, to float64) {
	C.elm_slider_range_set(p.obj, C.double(from), C.double(to))
}

func (p *Slider) Range() (float64, float64) {
	var from, to C.double
	C.elm_slider_range_get(p.obj, &from, &to)
	return float64(from), float64(to)
}

func (p *Slider) SetStep(step float64) {
	C.elm_slider_step_set(p.obj, C.double(step))
}

func (p *Slider) Step() float64 {
	return float64(C.elm_slider_step_get(p.obj))
}

/*
TODO

void 	elm_slider_indicator_format_function_set (Elm_Slider *obj, slider_func_type func, slider_freefunc_type free_func)
void 	elm_slider_units_format_function_set (Evas_Object *obj, slider_func_type func, slider_freefunc_type free_func)

*/

//------------------------------------------------------------

type Label struct {
	*Layout
}

func NewLabel(parent Object) *Label {
	eo := C.elm_label_add(parent.eo())
	return &Label{wrapLayout(eo)}
}

func (p *Label) SetSlideSpeed(s float64) {
	C.elm_label_slide_speed_set(p.obj, C.double(s))
}

func (p *Label) SlideSpeed() float64 {
	return float64(C.elm_label_slide_speed_get(p.obj))
}

func (p *Label) SetSlideDuration(d float64) {
	C.elm_label_slide_duration_set(p.obj, C.double(d))
}

func (p *Label) SlideDuration() float64 {
	return float64(C.elm_label_slide_duration_get(p.obj))
}

func (p *Label) SetSlideMode(mode LabelSlideMode) {
	C.elm_label_slide_mode_set(p.obj, C.Elm_Label_Slide_Mode(mode))
}

func (p *Label) SlideMode() LabelSlideMode {
	return LabelSlideMode(C.elm_label_slide_mode_get(p.obj))
}

func (p *Label) SetEllipsis(b bool) {
	C.elm_label_ellipsis_set(p.obj, eBool(b))
}

func (p *Label) Ellipsis() bool {
	return C.elm_label_ellipsis_get(p.obj) == eTrue
}

func (p *Label) StartSlide() {
	C.elm_label_slide_go(p.obj)
}

func (p *Label) SetWrapWidth(w int) {
	C.elm_label_wrap_width_set(p.obj, C.int(w))
}

func (p *Label) WrapWidth() int {
	return int(C.elm_label_wrap_width_get(p.obj))
}

func (p *Label) SetLineWrap(wt WrapType) {
	C.elm_label_line_wrap_set(p.obj, C.Elm_Wrap_Type(wt))
}

func (p *Label) LineWrap() WrapType {
	return WrapType(C.elm_label_line_wrap_get(p.obj))
}

//------------------------------------------------------------

type Datetime struct {
	*Layout
}

func NewDatetime(o Object) *Datetime {
	eo := C.elm_datetime_add(o.eo())
	return &Datetime{wrapLayout(eo)}
}

func (p *Datetime) SetFormat(fmtstr string) {
	cfmtstr := C.CString(fmtstr)
	defer free(cfmtstr)
	C.elm_datetime_format_set(p.obj, cfmtstr)
}

func (p *Datetime) Format() string {
	return C.GoString(C.elm_datetime_format_get(p.obj))
}

func (p *Datetime) SetFieldLimit(field DatetimeField, min, max int) {
	C.elm_datetime_field_limit_set(p.obj, C.Elm_Datetime_Field_Type(field), C.int(min), C.int(max))
}

func (p *Datetime) FieldLimit(field DatetimeField) (int, int) {
	var min, max C.int
	C.elm_datetime_field_limit_get(p.obj, C.Elm_Datetime_Field_Type(field), &min, &max)
	return int(min), int(max)
}

func (p *Datetime) SetMinValue(t time.Time) bool {
	return C.elm_datetime_value_min_set(p.obj, toCTime(t)) == eTrue
}

func (p *Datetime) MinValue() (time.Time, bool) {
	var ctm C.struct_tm
	b := C.elm_datetime_value_min_get(p.obj, &ctm)
	return toGoTime(&ctm), goBool(b)
}

func (p *Datetime) SetMaxValue(t time.Time) bool {
	return C.elm_datetime_value_max_set(p.obj, toCTime(t)) == eTrue
}

func (p *Datetime) MaxValue() (time.Time, bool) {
	var ctm C.struct_tm
	b := C.elm_datetime_value_max_get(p.obj, &ctm)
	return toGoTime(&ctm), goBool(b)
}

func (p *Datetime) SetValue(t time.Time) bool {
	return C.elm_datetime_value_set(p.obj, toCTime(t)) == eTrue
}

func (p *Datetime) Value() (time.Time, bool) {
	var ctm C.struct_tm
	b := C.elm_datetime_value_get(p.obj, &ctm)
	return toGoTime(&ctm), goBool(b)
}

func (p *Datetime) SetFieldVisible(field DatetimeField, b bool) {
	C.elm_datetime_field_visible_set(p.obj, C.Elm_Datetime_Field_Type(field), eBool(b))
}

func (p *Datetime) IsFieldVisible(field DatetimeField) bool {
	return C.elm_datetime_field_visible_get(p.obj, C.Elm_Datetime_Field_Type(field)) == eTrue
}

//------------------------------------------------------------

type Popup struct {
	*Layout
}

func NewPopup(parent Object) *Popup {
	eo := C.elm_popup_add(parent.eo())
	return &Popup{wrapLayout(eo)}
}

func (p *Popup) SetAlign(h, v float64) {
	C.elm_popup_align_set(p.obj, C.double(h), C.double(v))
}

func (p *Popup) Align() (float64, float64) {
	var h, v C.double
	C.elm_popup_align_get(p.obj, &h, &v)
	return float64(h), float64(v)
}

func (p *Popup) SetAllowEvents(b bool) {
	C.elm_popup_allow_events_set(p.obj, eBool(b))
}

func (p *Popup) IsAllowEvents() bool {
	return C.elm_popup_allow_events_get(p.obj) == eTrue
}

func (p *Popup) SetContentTextWrapType(t WrapType) {
	C.elm_popup_content_text_wrap_type_set(p.obj, C.Elm_Wrap_Type(t))
}

func (p *Popup) ContentTextWrapType() WrapType {
	return WrapType(C.elm_popup_content_text_wrap_type_get(p.obj))
}

func (p *Popup) SetOrient(t PopupOrient) {
	C.elm_popup_orient_set(p.obj, C.Elm_Popup_Orient(t))
}

func (p *Popup) Orient() PopupOrient {
	return PopupOrient(C.elm_popup_orient_get(p.obj))
}

func (p *Popup) SetScrollable(b bool) {
	C.elm_popup_scrollable_set(p.obj, eBool(b))
}

func (p *Popup) IsScrollable() bool {
	return C.elm_popup_scrollable_get(p.obj) == eTrue
}

func (p *Popup) Dismiss() {
	C.elm_popup_dismiss(p.obj)
}

func (p *Popup) SetTimeout(t float64) {
	C.elm_popup_timeout_set(p.obj, C.double(t))
}

func (p *Popup) Timeout() float64 {
	return float64(C.elm_popup_timeout_get(p.obj))
}

func (p *Popup) SetTitle(title string) {
	p.SetPartText("title,text", title)
}

func (p *Popup) SetIcon(icon *Icon) {
	p.SetPartContent("title,icon", icon)
}

func (p *Popup) SetButton1(btn *Button) {
	p.SetPartContent("button1", btn)
}

func (p *Popup) SetButton2(btn *Button) {
	p.SetPartContent("button2", btn)
}

func (p *Popup) SetButton3(btn *Button) {
	p.SetPartContent("button3", btn)
}

func (p *Popup) ItemFactory() *PopupItemFactory {
	return NewPopupItemFactory(p)
}

//------------------------------------------------------------

type Bg struct {
	*Layout
}

func NewBg(o Object) *Bg {
	eo := C.elm_bg_add(o.eo())
	return &Bg{wrapLayout(eo)}
}

func (p *Bg) SetBgColor(r, g, b int) {
	C.elm_bg_color_set(p.obj, C.int(r), C.int(g), C.int(b))
}

func (p *Bg) BgColor() (int, int, int) {
	var r, g, b C.int
	C.elm_bg_color_get(p.obj, &r, &g, &b)
	return int(r), int(g), int(b)
}

func (p *Bg) SetLoadSize(w, h int) {
	C.elm_bg_load_size_set(p.obj, C.int(w), C.int(h))
}

func (p *Bg) SetOption(option BgOption) {
	C.elm_bg_option_set(p.obj, C.Elm_Bg_Option(option))
}

func (p *Bg) Option() BgOption {
	return BgOption(C.elm_bg_option_get(p.obj))
}

func (p *Bg) SetFile(file string, group ...string) {
	var cgroup *C.char
	cfile := C.CString(file)
	defer free(cfile)
	if len(group) > 0 {
		if group[0] != "" {
			cgroup = C.CString(group[0])
			defer free(cgroup)
		}
	}
	C.elm_bg_file_set(p.obj, cfile, cgroup)
}

func (p *Bg) File() (string, string) {
	var file, group *C.char
	C.elm_bg_file_get(p.obj, &file, &group)
	return C.GoString(file), C.GoString(group)
}

//------------------------------------------------------------

type Spinner struct {
	*Layout
}

func NewSpinner(parent Object) *Spinner {
	eo := C.elm_spinner_add(parent.eo())
	return &Spinner{wrapLayout(eo)}
}

func (p *Spinner) SetValue(val float64) {
	C.elm_spinner_value_set(p.obj, C.double(val))
}

func (p *Spinner) Value() float64 {
	return float64(C.elm_spinner_value_get(p.obj))
}

func (p *Spinner) SetStep(step float64) {
	C.elm_spinner_step_set(p.obj, C.double(step))
}

func (p *Spinner) SetWrap(b bool) {
	C.elm_spinner_wrap_set(p.obj, eBool(b))
}

func (p *Spinner) SetMinMax(min, max float64) {
	C.elm_spinner_min_max_set(p.obj, C.double(min), C.double(max))
}

func (p *Spinner) SetInterval(iv float64) {
	C.elm_spinner_interval_set(p.obj, C.double(iv))
}

func (p *Spinner) SetLabelFormat(format string) {
	cfmt := C.CString(format)
	defer free(cfmt)
	C.elm_spinner_label_format_set(p.obj, cfmt)
}

func (p *Spinner) SetVertical() {
	p.SetStyle("vertical")
}

func (p *Spinner) AddSpecialValue(index float64, val string) {
	cstr := C.CString(val)
	defer free(cstr)
	C.elm_spinner_special_value_add(p.obj, C.double(index), cstr)
}

func (p *Spinner) SetEditable(b bool) {
	C.elm_spinner_editable_set(p.obj, eBool(b))
}

//------------------------------------------------------------

type Clock struct {
	*Layout
}

func NewClock(o Object) *Clock {
	eo := C.elm_clock_add(o.eo())
	return &Clock{wrapLayout(eo)}
}

func (p *Clock) SetShowAmPm(b bool) {
	C.elm_clock_show_am_pm_set(p.obj, eBool(b))
}

func (p *Clock) IsShowAmPm() bool {
	return C.elm_clock_show_am_pm_get(p.obj) == eTrue
}

func (p *Clock) SetFirstInterval(interval float64) {
	C.elm_clock_first_interval_set(p.obj, C.double(interval))
}

func (p *Clock) FirstInterval() float64 {
	return float64(C.elm_clock_first_interval_get(p.obj))
}

func (p *Clock) SetShowSeconds(b bool) {
	C.elm_clock_show_seconds_set(p.obj, eBool(b))
}

func (p *Clock) IsShowSeconds() bool {
	return C.elm_clock_show_seconds_get(p.obj) == eTrue
}

func (p *Clock) SetEdit(b bool) {
	C.elm_clock_edit_set(p.obj, eBool(b))
}

func (p *Clock) IsEdit() bool {
	return C.elm_clock_edit_get(p.obj) == eTrue
}

func (p *Clock) SetPause(b bool) {
	C.elm_clock_pause_set(p.obj, eBool(b))
}

func (p *Clock) IsPaused() bool {
	return C.elm_clock_pause_get(p.obj) == eTrue
}

func (p *Clock) SetTime(hour, min, sec int) {
	C.elm_clock_time_set(p.obj, C.int(hour), C.int(min), C.int(sec))
}

func (p *Clock) Time() (int, int, int) {
	var hour, min, sec C.int
	C.elm_clock_time_get(p.obj, &hour, &min, &sec)
	return int(hour), int(min), int(sec)
}

func (p *Clock) SetEditMode(mode ClockEditMode) {
	C.elm_clock_edit_mode_set(p.obj, C.Elm_Clock_Edit_Mode(mode))
}

func (p *Clock) EditMode() ClockEditMode {
	return ClockEditMode(C.elm_clock_edit_mode_get(p.obj))
}

//------------------------------------------------------------

type Colorselector struct {
	*Layout
}

func NewColorselector(o Object) *Colorselector {
	eo := C.elm_colorselector_add(o.eo())
	return &Colorselector{wrapLayout(eo)}
}

func (p *Colorselector) SetColor(r, g, b, a int) {
	C.elm_colorselector_color_set(p.obj, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *Colorselector) Color() (int, int, int, int) {
	var r, g, b, a C.int
	C.elm_colorselector_color_get(p.obj, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *Colorselector) SetPaletteName(palette string) {
	cpalette := C.CString(palette)
	defer free(cpalette)
	C.elm_colorselector_palette_name_set(p.obj, cpalette)
}

func (p *Colorselector) PaletteName() string {
	return C.GoString(C.elm_colorselector_palette_name_get(p.obj))
}

func (p *Colorselector) SetMode(mode ColorselectorMode) {
	C.elm_colorselector_mode_set(p.obj, C.Elm_Colorselector_Mode(mode))
}

func (p *Colorselector) Mode() ColorselectorMode {
	return ColorselectorMode(C.elm_colorselector_mode_get(p.obj))
}

func (p *Colorselector) Items() []*ColorselectorPaletteItem {
	lst := C.elm_colorselector_palette_items_get(p.obj)
	if lst != nil {
		return newListIterator(lst).ColorselectorPaletteItemSlice()
	}
	return nil
}

func (p *Colorselector) SelectedItem() *ColorselectorPaletteItem {
	return wrapColorselectorPaletteItem(C.elm_colorselector_palette_selected_item_get(p.obj))
}

func (p *Colorselector) AddColor(r, g, b, a int) *ColorselectorPaletteItem {
	return wrapColorselectorPaletteItem(C.elm_colorselector_palette_color_add(p.obj, C.int(r), C.int(g), C.int(b), C.int(a)))
}

func (p *Colorselector) ClearPalette() {
	C.elm_colorselector_palette_clear(p.obj)
}

//------------------------------------------------------------

type Calendar struct {
	*Layout
}

type CalendarMark struct {
	obj *C.Elm_Calendar_Mark
}

func newCalendarMark(p *C.Elm_Calendar_Mark) *CalendarMark {
	return &CalendarMark{
		obj: p,
	}
}

func NewCalendar(o Object) *Calendar {
	eo := C.elm_calendar_add(o.eo())
	return &Calendar{wrapLayout(eo)}
}

func (p *Calendar) SetFirstDayOfWeek(day CalendarWeekday) {
	C.elm_calendar_first_day_of_week_set(p.obj, C.Elm_Calendar_Weekday(day))
}

func (p *Calendar) FirstDayOfWeek() CalendarWeekday {
	return CalendarWeekday(C.elm_calendar_first_day_of_week_get(p.obj))
}

func (p *Calendar) SetSelectable(s CalendarSelectable) {
	C.elm_calendar_selectable_set(p.obj, C.Elm_Calendar_Selectable(s))
}

func (p *Calendar) Selectable() CalendarSelectable {
	return CalendarSelectable(C.elm_calendar_selectable_get(p.obj))
}

func (p *Calendar) SetInterval(ival float64) {
	C.elm_calendar_interval_set(p.obj, C.double(ival))
}

func (p *Calendar) Interval() float64 {
	return float64(C.elm_calendar_interval_get(p.obj))
}

func (p *Calendar) SetWeekdaysNames(names []string) {
	//names len must be 7
	if len(names) != 7 {
		return
	}
	cnames := make([]*C.char, len(names))
	for i, name := range names {
		cnames[i] = C.CString(name)
	}
	defer free(cnames)
	C.elm_calendar_weekdays_names_set(p.obj, (**C.char)(unsafe.Pointer(&cnames[0])))
}

func (p *Calendar) WeekdaysNames() []string {
	return CArrayToSlice(C.elm_calendar_weekdays_names_get(p.obj), 7)
}

func (p *Calendar) SetSelectMode(mode CalendarSelectMode) {
	C.elm_calendar_select_mode_set(p.obj, C.Elm_Calendar_Select_Mode(mode))
}

func (p *Calendar) SelectMode() CalendarSelectMode {
	return CalendarSelectMode(C.elm_calendar_select_mode_get(p.obj))
}

func (p *Calendar) SetMinMaxYear(min, max int) {
	C.elm_calendar_min_max_year_set(p.obj, C.int(min), C.int(max))
}

func (p *Calendar) MinMaxYear() (int, int) {
	var min, max C.int
	C.elm_calendar_min_max_year_get(p.obj, &min, &max)
	return int(min), int(max)
}

func (p *Calendar) SetSelectedTime(tm time.Time) {
	C.elm_calendar_selected_time_set(p.obj, toCTime(tm))
}

func (p *Calendar) SelectedTime() (time.Time, bool) {
	var ctm C.struct_tm
	b := C.elm_calendar_selected_time_get(p.obj, &ctm)
	return toGoTime(&ctm), goBool(b)
}

func (p *Calendar) CalendarMarks() []*CalendarMark {
	lst := C.elm_calendar_marks_get(p.obj)
	if lst != nil {
		return newListIterator(lst).CalendarMarkSlice()
	}
	return nil
}

func (p *Calendar) ClearMarks() {
	C.elm_calendar_marks_clear(p.obj)
}

func (p *Calendar) DrawMarks() {
	C.elm_calendar_marks_draw(p.obj)
}

func (p *Calendar) AddMark(mark_type string, mark_time time.Time, rt CalendarMarkRepeatType) *CalendarMark {
	cmt := C.CString(mark_type)
	defer free(cmt)
	return newCalendarMark(C.elm_calendar_mark_add(p.obj, cmt, toCTime(mark_time), C.Elm_Calendar_Mark_Repeat_Type(rt)))
}

func (p *Calendar) DisplayedTime() (time.Time, bool) {
	var ctm C.struct_tm
	b := C.elm_calendar_displayed_time_get(p.obj, &ctm)
	return toGoTime(&ctm), goBool(b)
}

func DeleteCalendarMark(mark *CalendarMark) {
	C.elm_calendar_mark_del(mark.obj)
}

/*
TODO:
void 	elm_calendar_format_function_set (Elm_Calendar *obj, Elm_Calendar_Format_Cb format_function)

in 1.19
void 	elm_calendar_date_min_set (Elm_Calendar *obj, const Efl_Time *min)
const Efl_Time * 	elm_calendar_date_min_get (const Elm_Calendar *obj)
void 	elm_calendar_date_max_set (Elm_Calendar *obj, const Efl_Time *max)
const Efl_Time * 	elm_calendar_date_max_get (const Elm_Calendar *obj)
*/

//------------------------------------------------------------

type Dayselector struct {
	*Layout
}

func NewDayselector(o Object) *Dayselector {
	eo := C.elm_dayselector_add(o.eo())
	return &Dayselector{wrapLayout(eo)}
}

func (p *Dayselector) SetWeekStart(day DayselectorDay) {
	C.elm_dayselector_week_start_set(p.obj, C.Elm_Dayselector_Day(day))
}

func (p *Dayselector) WeekStart() DayselectorDay {
	return DayselectorDay(C.elm_dayselector_week_start_get(p.obj))
}

func (p *Dayselector) SetWeekendLength(l uint) {
	C.elm_dayselector_weekend_length_set(p.obj, C.uint(l))
}

func (p *Dayselector) WeekendLength() uint {
	return uint(C.elm_dayselector_weekend_length_get(p.obj))
}

func (p *Dayselector) SetWeekendStart(day DayselectorDay) {
	C.elm_dayselector_weekend_start_set(p.obj, C.Elm_Dayselector_Day(day))
}

func (p *Dayselector) WeekendStart() DayselectorDay {
	return DayselectorDay(C.elm_dayselector_weekend_start_get(p.obj))
}

func (p *Dayselector) SetWeekdaysNames(names []string) {
	cnames := make([]*C.char, len(names)+1) // +1 for null term
	for i, name := range names {
		cnames[i] = C.CString(name)
	}
	defer free(cnames)
	C.elm_dayselector_weekdays_names_set(p.obj, (**C.char)(unsafe.Pointer(&cnames[0])))
}

func (p *Dayselector) WeekdaysNames() []string {
	lst := C.elm_dayselector_weekdays_names_get(p.obj)
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func (p *Dayselector) SetSelectedDay(day DayselectorDay, selected bool) {
	C.elm_dayselector_day_selected_set(p.obj, C.Elm_Dayselector_Day(day), eBool(selected))
}

func (p *Dayselector) IsDaySelected(day DayselectorDay) bool {
	return C.elm_dayselector_day_selected_get(p.obj, C.Elm_Dayselector_Day(day)) == eTrue
}

//------------------------------------------------------------

type Fileselector struct {
	*Layout
}

func NewFileselector(o Object) *Fileselector {
	eo := C.elm_fileselector_add(o.eo())
	return &Fileselector{wrapLayout(eo)}
}

func (p *Fileselector) SetIsSave(b bool) {
	C.elm_fileselector_is_save_set(p.obj, eBool(b))
}

func (p *Fileselector) IsSave() bool {
	return C.elm_fileselector_is_save_get(p.obj) == eTrue
}

func (p *Fileselector) SetFolderOnly(b bool) {
	C.elm_fileselector_folder_only_set(p.obj, eBool(b))
}

func (p *Fileselector) IsFolderOnly() bool {
	return C.elm_fileselector_folder_only_get(p.obj) == eTrue
}

func (p *Fileselector) SetExpandable(b bool) {
	C.elm_fileselector_expandable_set(p.obj, eBool(b))
}

func (p *Fileselector) IsExpandable() bool {
	return C.elm_fileselector_expandable_get(p.obj) == eTrue
}

func (p *Fileselector) SetPath(path string) {
	cpath := C.CString(path)
	defer free(cpath)
	C.elm_fileselector_path_set(p.obj, cpath)
}

func (p *Fileselector) Path() string {
	return C.GoString(C.elm_fileselector_path_get(p.obj))
}

func (p *Fileselector) SetMode(mode FileselectorMode) {
	C.elm_fileselector_mode_set(p.obj, C.Elm_Fileselector_Mode(mode))
}

func (p *Fileselector) Mode() FileselectorMode {
	return FileselectorMode(C.elm_fileselector_mode_get(p.obj))
}

func (p *Fileselector) SetMultiSelect(b bool) {
	C.elm_fileselector_multi_select_set(p.obj, eBool(b))
}

func (p *Fileselector) IsMultiSelect() bool {
	return C.elm_fileselector_multi_select_get(p.obj) == eTrue
}

func (p *Fileselector) SetSelected(path string) bool {
	cpath := C.CString(path)
	defer free(cpath)
	return C.elm_fileselector_selected_set(p.obj, cpath) == eTrue
}

func (p *Fileselector) Selected() string {
	return C.GoString(C.elm_fileselector_selected_get(p.obj))
}

func (p *Fileselector) SelectedPaths() []string {
	lst := C.elm_fileselector_selected_paths_get(p.obj)
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

func (p *Fileselector) SetHiddenVisible(b bool) {
	C.elm_fileselector_hidden_visible_set(p.obj, eBool(b))
}

func (p *Fileselector) IsHiddenVisible() bool {
	return C.elm_fileselector_hidden_visible_get(p.obj) == eTrue
}

func (p *Fileselector) SetThumbnailSize(w, h int) {
	C.elm_fileselector_thumbnail_size_set(p.obj, C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *Fileselector) ThumbnailSize() (int, int) {
	var w, h C.Evas_Coord
	C.elm_fileselector_thumbnail_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Fileselector) SetSortMethod(method FileselectorSort) {
	C.elm_fileselector_sort_method_set(p.obj, C.Elm_Fileselector_Sort(method))
}

func (p *Fileselector) SortMethod() FileselectorSort {
	return FileselectorSort(C.elm_fileselector_sort_method_get(p.obj))
}

func (p *Fileselector) SetOkCancelButtonsEnabled(b bool) {
	C.elm_fileselector_buttons_ok_cancel_set(p.obj, eBool(b))
}

func (p *Fileselector) IsOkCancelButtonsEnabled() bool {
	return C.elm_fileselector_buttons_ok_cancel_get(p.obj) == eTrue
}

/*
Eina_Bool 	elm_fileselector_mime_types_filter_append (Evas_Object *obj, const char *mime_types, const char *filter_name)
Eina_Bool 	elm_fileselector_custom_filter_append (Evas_Object *obj, Elm_Fileselector_Filter_Func func, void *data, const char *filter_name)
void 	elm_fileselector_filters_clear (Evas_Object *obj)
*/

//------------------------------------------------------------

type FileselectorButton struct {
	*Layout
}

func NewFileselectorButton(parent Object) *FileselectorButton {
	return &FileselectorButton{wrapLayout(C.elm_fileselector_button_add(parent.eo()))}
}

func (p *FileselectorButton) SetWindowTitle(title string) {
	ctitle := C.CString(title)
	defer free(ctitle)
	C.elm_fileselector_button_window_title_set(p.obj, ctitle)
}

func (p *FileselectorButton) WindowTitle() string {
	return C.GoString(C.elm_fileselector_button_window_title_get(p.obj))
}

func (p *FileselectorButton) SetWindowSize(w, h int) {
	C.elm_fileselector_button_window_size_set(p.obj, C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *FileselectorButton) WindowSize() (int, int) {
	var w, h C.Evas_Coord
	C.elm_fileselector_button_window_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *FileselectorButton) SetInwinMode(b bool) {
	C.elm_fileselector_button_inwin_mode_set(p.obj, eBool(b))
}

func (p *FileselectorButton) IsInwinMode() bool {
	return C.elm_fileselector_button_inwin_mode_get(p.obj) == eTrue
}

//------------------------------------------------------------

type FileselectorEntry struct {
	*Layout
}

func NewFileselectorEntry(parent Object) *FileselectorEntry {
	return &FileselectorEntry{wrapLayout(C.elm_fileselector_entry_add(parent.eo()))}
}

func (p *FileselectorEntry) SetWindowTitle(title string) {
	ctitle := C.CString(title)
	defer free(ctitle)
	C.elm_fileselector_entry_window_title_set(p.obj, ctitle)
}

func (p *FileselectorEntry) WindowTitle() string {
	return C.GoString(C.elm_fileselector_entry_window_title_get(p.obj))
}

func (p *FileselectorEntry) SetWindowSize(w, h int) {
	C.elm_fileselector_entry_window_size_set(p.obj, C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *FileselectorEntry) WindowSize() (int, int) {
	var w, h C.Evas_Coord
	C.elm_fileselector_entry_window_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *FileselectorEntry) SetInwinMode(b bool) {
	C.elm_fileselector_entry_inwin_mode_set(p.obj, eBool(b))
}

func (p *FileselectorEntry) IsInwinMode() bool {
	return C.elm_fileselector_entry_inwin_mode_get(p.obj) == eTrue
}

//------------------------------------------------------------

type Image struct {
	*widgetBase
}

func NewImage(o Object) *Image {
	eo := C.elm_image_add(o.eo())
	return &Image{wrapWidgetBase(eo)}
}

func wrapImage(o *C.Eo) *Image {
	if o != nil {
		return &Image{wrapWidgetBase(o)}
	}
	return nil
}

func (p *Image) SetEditable(b bool) {
	C.elm_image_editable_set(p.obj, eBool(b))
}

func (p *Image) IsEditable() bool {
	return C.elm_image_editable_get(p.obj) == eTrue
}

func (p *Image) SetSmooth(b bool) {
	C.elm_image_smooth_set(p.obj, eBool(b))
}

func (p *Image) IsSmooth() bool {
	return C.elm_image_smooth_get(p.obj) == eTrue
}

func (p *Image) SetNoScale(b bool) {
	C.elm_image_no_scale_set(p.obj, eBool(b))
}

func (p *Image) IsNoScale() bool {
	return C.elm_image_no_scale_get(p.obj) == eTrue
}

func (p *Image) SetFixedAspect(b bool) {
	C.elm_image_aspect_fixed_set(p.obj, eBool(b))
}

func (p *Image) SetOrient(orient ImageOrient) {
	C.elm_image_orient_set(p.obj, C.Elm_Image_Orient(orient))
}

func (p *Image) Orient() ImageOrient {
	return ImageOrient(C.elm_image_orient_get(p.obj))
}

func (p *Image) SetFillOutside(b bool) {
	C.elm_image_fill_outside_set(p.obj, eBool(b))
}

func (p *Image) IsFillOutside() bool {
	return C.elm_image_fill_outside_get(p.obj) == eTrue
}

func (p *Image) SetResizable(up bool, down bool) {
	C.elm_image_resizable_set(p.obj, eBool(up), eBool(down))
}

func (p *Image) IsResizable() (bool, bool) {
	var up, down C.Eina_Bool
	C.elm_image_resizable_get(p.obj, &up, &down)
	return goBool(up), goBool(down)
}

func (p *Image) SetPreloadDisabled(b bool) {
	C.elm_image_preload_disabled_set(p.obj, eBool(b))
}

/* size of the image not widget */
func (p *Image) ImageSize() (int, int) {
	var w, h C.int
	C.elm_image_object_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Image) SetFile(file string, group string) {
	var cgroup *C.char
	if group != "" {
		cgroup := C.CString(group)
		defer free(cgroup)
	}
	cfile := C.CString(file)
	defer free(cfile)
	C.elm_image_file_set(p.obj, cfile, cgroup)
}

func (p *Image) File() (string, string) {
	var file, group *C.char
	C.elm_image_file_get(p.obj, &file, &group)
	return C.GoString(file), C.GoString(group)
}

func (p *Image) SetPrescale(scale int) {
	C.elm_image_prescale_set(p.obj, C.int(scale))
}

func (p *Image) Prescale() int {
	return int(C.elm_image_prescale_get(p.obj))
}

func (p *Image) SetPlayAnimated(b bool) {
	C.elm_image_animated_play_set(p.obj, eBool(b))
}

func (p *Image) IsPlayAnimated() bool {
	return C.elm_image_animated_play_get(p.obj) == eTrue
}

func (p *Image) SetAnimated(b bool) {
	C.elm_image_animated_set(p.obj, eBool(b))
}

func (p *Image) IsAnimated() bool {
	return C.elm_image_animated_get(p.obj) == eTrue
}

func (p *Image) IsAnimatedAvailable() bool {
	return C.elm_image_animated_available_get(p.obj) == eTrue
}

func (p *Image) ImageObject() Object {
	return wrapObjectBase(C.elm_image_object_get(p.obj))
}

/*
Eina_Bool 	elm_image_memfile_set (Elm_Image *obj, const void *img, size_t size, const char *format, const char *key)


Evas_Object * 	elm_image_object_get (const Elm_Image *obj)


Eina_Bool 	elm_image_mmap_set (Evas_Object *obj, const Eina_File *file, const char *group)
*/

type ImageProgress struct {
	obj *C.Elm_Image_Progress
}

func wrapImageProgress(o *C.Elm_Image_Progress) *ImageProgress {
	if o != nil {
		return &ImageProgress{o}
	}
	return nil
}

func (p *ImageProgress) Now() float64 {
	return float64(p.obj.now)
}

func (p *ImageProgress) Total() float64 {
	return float64(p.obj.total)
}

type ImageError struct {
	obj *C.Elm_Image_Error
}

func wrapImageError(o *C.Elm_Image_Error) *ImageError {
	if o != nil {
		return &ImageError{o}
	}
	return nil
}

// TODO : is this LoadError ?
func (p *ImageError) Status() int {
	return int(p.obj.status)
}

func (p *ImageError) OpenError() bool {
	return p.obj.open_error == eTrue
}

//------------------------------------------------------------

var _ Object = &Icon{}

type Icon struct {
	*Image
}

func NewIcon(parent Object) *Icon {
	return &Icon{wrapImage(C.elm_icon_add(parent.eo()))}
}

func wrapIcon(o *C.Eo) *Icon {
	if o != nil {
		return &Icon{wrapImage(o)}
	}
	return nil
}

func (p *Icon) SetStandard(name string) bool {
	cname := C.CString(name)
	defer free(cname)
	return C.elm_icon_standard_set(p.obj, cname) == eTrue
}

func (p *Icon) Standard() string {
	return C.GoString(C.elm_icon_standard_get(p.obj))
}

func (p *Icon) SetThumb(file, group string) {
	cfile := C.CString(file)
	cgroup := C.CString(group)
	defer free(cfile, cgroup)
	C.elm_icon_thumb_set(p.obj, cfile, cgroup)
}

//------------------------------------------------------------

type Progressbar struct {
	*Layout
}

func NewProgressbar(parent Object) *Progressbar {
	eo := C.elm_progressbar_add(parent.eo())
	return &Progressbar{wrapLayout(eo)}
}

func (p *Progressbar) SetSpanSize(s int) {
	C.elm_progressbar_span_size_set(p.obj, C.Evas_Coord(s))
}

func (p *Progressbar) SpanSize() int {
	return int(C.elm_progressbar_span_size_get(p.obj))
}

func (p *Progressbar) SetPulse(b bool) {
	C.elm_progressbar_pulse_set(p.obj, eBool(b))
}

func (p *Progressbar) IsPulsing() bool {
	return C.elm_progressbar_pulse_get(p.obj) == eTrue
}

func (p *Progressbar) SetValue(val float64) {
	C.elm_progressbar_value_set(p.obj, C.double(val))
}

func (p *Progressbar) Value() float64 {
	return float64(C.elm_progressbar_value_get(p.obj))
}

func (p *Progressbar) SetInverted(b bool) {
	C.elm_progressbar_inverted_set(p.obj, eBool(b))
}

func (p *Progressbar) IsInverted() bool {
	return C.elm_progressbar_inverted_get(p.obj) == eTrue
}

func (p *Progressbar) SetHorizontal(b bool) {
	C.elm_progressbar_horizontal_set(p.obj, eBool(b))
}

func (p *Progressbar) IsHorizontal() bool {
	return C.elm_progressbar_horizontal_get(p.obj) == eTrue
}

func (p *Progressbar) SetUnitFormat(f string) {
	cf := C.CString(f)
	defer free(cf)
	C.elm_progressbar_unit_format_set(p.obj, cf)
}

func (p *Progressbar) UnitFormat() string {
	return C.GoString(C.elm_progressbar_unit_format_get(p.obj))
}

func (p *Progressbar) Pulse(b bool) {
	C.elm_progressbar_pulse(p.obj, eBool(b))
}

//------------------------------------------------------------

type Actionslider struct {
	*Layout
}

func NewActionslider(o Object) *Actionslider {
	eo := C.elm_actionslider_add(o.eo())
	return &Actionslider{wrapLayout(eo)}
}

func (p *Actionslider) SetIndicatorPos(pos ActionsliderPos) {
	C.elm_actionslider_indicator_pos_set(p.obj, C.Elm_Actionslider_Pos(pos))
}

func (p *Actionslider) IndicatorPos() ActionsliderPos {
	return ActionsliderPos(C.elm_actionslider_indicator_pos_get(p.obj))
}

func (p *Actionslider) SetMagnetPos(pos ActionsliderPos) {
	C.elm_actionslider_magnet_pos_set(p.obj, C.Elm_Actionslider_Pos(pos))
}

func (p *Actionslider) MagnetPos() ActionsliderPos {
	return ActionsliderPos(C.elm_actionslider_magnet_pos_get(p.obj))
}

func (p *Actionslider) SetPosEnabled(pos ActionsliderPos) {
	C.elm_actionslider_enabled_pos_set(p.obj, C.Elm_Actionslider_Pos(pos))
}

func (p *Actionslider) PosEnabled() ActionsliderPos {
	return ActionsliderPos(C.elm_actionslider_enabled_pos_get(p.obj))
}

func (p *Actionslider) SelectedLabel() string {
	return C.GoString(C.elm_actionslider_selected_label_get(p.obj))
}

//------------------------------------------------------------

type Entry struct {
	*Layout
}

func NewEntry(parent Object) *Entry {
	eo := C.elm_entry_add(parent.eo())
	return &Entry{wrapLayout(eo)}
}

func wrapEntry(o *C.Eo) *Entry {
	if o != nil {
		return &Entry{wrapLayout(o)}
	}
	return nil
}

func NewSingleLineEntry(parent Object) *Entry {
	entry := NewEntry(parent)
	entry.SetSingleLine(true)
	entry.SetScrollable(true)
	return entry
}

func NewPasswordEntry(parent Object) *Entry {
	entry := NewSingleLineEntry(parent)
	entry.SetPasswordMode(true)
	return entry
}

func (p *Entry) SetEditable(b bool) {
	C.elm_entry_editable_set(p.obj, eBool(b))
}

func (p *Entry) IsEditable() bool {
	return C.elm_entry_editable_get(p.obj) == eTrue
}

func (p *Entry) SetSingleLine(b bool) {
	C.elm_entry_single_line_set(p.obj, eBool(b))
}

func (p *Entry) SetPasswordMode(b bool) {
	C.elm_entry_password_set(p.obj, eBool(b))
}

func (p *Entry) SetPlaceholder(ph string) {
	p.SetPartText("guide", ph)
}

func (p *Entry) SetScrollable(b bool) {
	C.elm_entry_scrollable_set(p.obj, eBool(b))
}

func (p *Entry) IsScrollable() bool {
	return C.elm_entry_scrollable_get(p.obj) == eTrue
}

func (p *Entry) SetShowInputPanelOnDemand(b bool) {
	C.elm_entry_input_panel_show_on_demand_set(p.obj, eBool(b))
}

func (p *Entry) IsShowInputPanelOnDemand() bool {
	return C.elm_entry_input_panel_show_on_demand_get(p.obj) == eTrue
}

func (p *Entry) SetContextMenuDisabled(b bool) {
	C.elm_entry_context_menu_disabled_set(p.obj, eBool(b))
}

func (p *Entry) IsContextMenuDisabled() bool {
	return C.elm_entry_context_menu_disabled_get(p.obj) == eTrue
}

func (p *Entry) SetCNPMode(mode CNPMode) {
	C.elm_entry_cnp_mode_set(p.obj, C.Elm_Cnp_Mode(mode))
}

func (p *Entry) CNPMode() CNPMode {
	return CNPMode(C.elm_entry_cnp_mode_get(p.obj))
}

func (p *Entry) SetFileTextFormat(format TextFormat) {
	C.elm_entry_file_text_format_set(p.obj, C.Elm_Text_Format(format))
}

func (p *Entry) SetInputPanelLanguage(lang InputPanelLanguage) {
	C.elm_entry_input_panel_language_set(p.obj, C.Elm_Input_Panel_Lang(lang))
}

func (p *Entry) InputPanelLanguage() InputPanelLanguage {
	return InputPanelLanguage(C.elm_entry_input_panel_language_get(p.obj))
}

func (p *Entry) SetSelectionHandlerDisabled(b bool) {
	C.elm_entry_selection_handler_disabled_set(p.obj, eBool(b))
}

func (p *Entry) SetInputPanelLayoutVariation(v int) {
	C.elm_entry_input_panel_layout_variation_set(p.obj, C.int(v))
}

func (p *Entry) InputPanelVariation() int {
	return int(C.elm_entry_input_panel_layout_variation_get(p.obj))
}

func (p *Entry) SetAutocapitalType(t AutocapitalType) {
	C.elm_entry_autocapital_type_set(p.obj, C.Elm_Autocapital_Type(t))
}

func (p *Entry) AutocapitalType() AutocapitalType {
	return AutocapitalType(C.elm_entry_autocapital_type_get(p.obj))
}

func (p *Entry) SetAnchorStyle(style string) {
	cstyle := C.CString(style)
	defer free(cstyle)
	C.elm_entry_anchor_hover_style_set(p.obj, cstyle)
}

func (p *Entry) AnchorHoverStyle() string {
	return C.GoString(C.elm_entry_anchor_hover_style_get(p.obj))
}

func (p *Entry) IsSingleLine() bool {
	return C.elm_entry_single_line_get(p.obj) == eTrue
}

func (p *Entry) IsPasswordMode() bool {
	return C.elm_entry_password_get(p.obj) == eTrue
}

func (p *Entry) SetInputPanelReturnKeyDisabled(b bool) {
	C.elm_entry_input_panel_return_key_disabled_set(p.obj, eBool(b))
}

func (p *Entry) IsInputPanelReturnKeyDisabled() bool {
	return C.elm_entry_input_panel_return_key_disabled_get(p.obj) == eTrue
}

func (p *Entry) SetAutosave(b bool) {
	C.elm_entry_autosave_set(p.obj, eBool(b))
}

func (p *Entry) IsAutoSave() bool {
	return C.elm_entry_autosave_get(p.obj) == eTrue
}

func (p *Entry) SetAnchorHoverParent(canvas Canvas) {
	C.elm_entry_anchor_hover_parent_set(p.obj, canvas.eo())
}

func (p *Entry) AnchorHoverParent() Canvas {
	return wrapEvas(C.elm_entry_anchor_hover_parent_get(p.obj))
}

func (p *Entry) SetAllowPrediction(b bool) {
	C.elm_entry_prediction_allow_set(p.obj, eBool(b))
}

func (p *Entry) IsAllowPrediction() bool {
	return C.elm_entry_prediction_allow_get(p.obj) == eTrue
}

func (p *Entry) SetInputHint(hint InputHint) {
	C.elm_entry_input_hint_set(p.obj, C.Elm_Input_Hints(hint))
}

func (p *Entry) InputHint() InputHint {
	return InputHint(C.elm_entry_input_hint_get(p.obj))
}

func (p *Entry) SetInputPanelLayout(layout InputPanelLayout) {
	C.elm_entry_input_panel_layout_set(p.obj, C.Elm_Input_Panel_Layout(layout))
}

func (p *Entry) InputPanelLayout() InputPanelLayout {
	return InputPanelLayout(C.elm_entry_input_panel_layout_get(p.obj))
}

func (p *Entry) SetInputPanelReturnKeyType(t InputPanelReturnKeyType) {
	C.elm_entry_input_panel_return_key_type_set(p.obj, C.Elm_Input_Panel_Return_Key_Type(t))
}

func (p *Entry) InputPanelReturnKeyType() InputPanelReturnKeyType {
	return InputPanelReturnKeyType(C.elm_entry_input_panel_return_key_type_get(p.obj))
}

func (p *Entry) SetInputPanelEnabled(b bool) {
	C.elm_entry_input_panel_enabled_set(p.obj, eBool(b))
}

func (p *Entry) IsInputPanelEnabled() bool {
	return C.elm_entry_input_panel_enabled_get(p.obj) == eTrue
}

func (p *Entry) SetLineWrap(t WrapType) {
	C.elm_entry_line_wrap_set(p.obj, C.Elm_Wrap_Type(t))
}

func (p *Entry) LineWrap() WrapType {
	return WrapType(C.elm_entry_line_wrap_get(p.obj))
}

func (p *Entry) SetCursorPos(pos int) {
	C.elm_entry_cursor_pos_set(p.obj, C.int(pos))
}

func (p *Entry) CursorPos() int {
	return int(C.elm_entry_cursor_pos_get(p.obj))
}

func (p *Entry) SetIconVisible(b bool) {
	C.elm_entry_icon_visible_set(p.obj, eBool(b))
}

func (p *Entry) MoveCursorToLineEnd() {
	C.elm_entry_cursor_line_end_set(p.obj)
}

func (p *Entry) SetSelectRegion(start, end int) {
	C.elm_entry_select_region_set(p.obj, C.int(start), C.int(end))
}

func (p *Entry) SelectRegion() (int, int) {
	var start, end C.int
	C.elm_entry_select_region_get(p.obj, &start, &end)
	return int(start), int(end)
}

func (p *Entry) SetInputPanelReturnKeyAutoEnabled(b bool) {
	C.elm_entry_input_panel_return_key_autoenabled_set(p.obj, eBool(b))
}

func (p *Entry) SetEndPartContentVisible(b bool) {
	C.elm_entry_end_visible_set(p.obj, eBool(b))
}

func (p *Entry) MoveCursorToBegin() {
	C.elm_entry_cursor_begin_set(p.obj)
}

func (p *Entry) MoveCursorToLineBegin() {
	C.elm_entry_cursor_line_begin_set(p.obj)
}

func (p *Entry) MoveCursorToEnd() {
	C.elm_entry_cursor_end_set(p.obj)
}

func (p *Entry) Textblock() *EvasTextblock {
	return wrapEvasTextblock(C.elm_entry_textblock_get(p.obj))
}

func (p *Entry) CursorGeometry() (int, int, int, int) {
	var x, y, w, h C.int
	C.elm_entry_cursor_geometry_get(p.obj, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *Entry) IsCursorOverFormatNode() bool {
	return C.elm_entry_cursor_is_format_get(p.obj) == eTrue
}

func (p *Entry) CharAtCursor() string {
	return C.GoString(C.elm_entry_cursor_content_get(p.obj))
}

func (p *Entry) Selection() string {
	return C.GoString(C.elm_entry_selection_get(p.obj))
}

func (p *Entry) IsCursorOverVisibleFormatNode() bool {
	return C.elm_entry_cursor_is_visible_format_get(p.obj) == eTrue
}

func (p *Entry) SetAllowSelect(b bool) {
	C.elm_entry_select_allow_set(p.obj, eBool(b))
}

func (p *Entry) IsAllowSelect() bool {
	return C.elm_entry_select_allow_get(p.obj) == eTrue
}

func (p *Entry) MoveCursorPrev() {
	C.elm_entry_cursor_prev(p.obj)
}

func (p *Entry) RemoveUserStyleOnTop() {
	C.elm_entry_text_style_user_pop(p.obj)
}

func (p *Entry) PushUserStyle(style string) {
	cstyle := C.CString(style)
	defer free(cstyle)
	C.elm_entry_text_style_user_push(p.obj, cstyle)
}

func (p *Entry) PeekUserStyle() string {
	return C.GoString(C.elm_entry_text_style_user_peek(p.obj))
}

func (p *Entry) ShowInputPanel() {
	C.elm_entry_input_panel_show(p.obj)
}

func (p *Entry) ResetImfContext() {
	C.elm_entry_imf_context_reset(p.obj)
}

func (p *Entry) ForceCalc() {
	C.elm_entry_calc_force(p.obj)
}

func (p *Entry) EndAnchorHover() {
	C.elm_entry_anchor_hover_end(p.obj)
}

func (p *Entry) BeginCursorSelection() {
	C.elm_entry_cursor_selection_begin(p.obj)
}

func (p *Entry) MoveCursorDown() bool {
	return C.elm_entry_cursor_down(p.obj) == eTrue
}

func (p *Entry) MoveCursorUp() bool {
	return C.elm_entry_cursor_up(p.obj) == eTrue
}

func (p *Entry) SaveFile() {
	C.elm_entry_file_save(p.obj)
}

func (p *Entry) CopySelection() {
	C.elm_entry_selection_copy(p.obj)
}

func (p *Entry) ClearContextMenu() {
	C.elm_entry_context_menu_clear(p.obj)
}

func (p *Entry) InsertText(txt string) {
	ctxt := C.CString(txt)
	defer free(ctxt)
	C.elm_entry_entry_insert(p.obj, ctxt)
}

func (p *Entry) PasteSelection() {
	C.elm_entry_selection_paste(p.obj)
}

func (p *Entry) MoveCursorNext() bool {
	return C.elm_entry_cursor_next(p.obj) == eTrue
}

func (p *Entry) SelectNone() {
	C.elm_entry_select_none(p.obj)
}

func (p *Entry) HideInputPanel() {
	C.elm_entry_input_panel_hide(p.obj)
}

func (p *Entry) SelectAll() {
	C.elm_entry_select_all(p.obj)
}

func (p *Entry) EndCursorSelection() {
	C.elm_entry_cursor_selection_end(p.obj)
}

func (p *Entry) CutSelection() {
	C.elm_entry_selection_cut(p.obj)
}

func (p *Entry) IsEmpty() bool {
	return C.elm_entry_is_empty(p.obj) == eTrue
}

func (p *Entry) AppendText(txt string) {
	ctxt := C.CString(txt)
	defer free(ctxt)
	C.elm_entry_entry_append(p.obj, ctxt)
}

func MarkupToUtf8(str string) string {
	cstr := C.CString(str)
	defer free(cstr)
	return C.GoString(C.elm_entry_markup_to_utf8(cstr))
}

func Utf8ToMarkup(str string) string {
	cstr := C.CString(str)
	defer free(cstr)
	return C.GoString(C.elm_entry_utf8_to_markup(cstr))
}

func (p *Entry) SetFile(file string, format TextFormat) bool {
	cfile := C.CString(file)
	defer free(cfile)
	return C.elm_entry_file_set(p.obj, cfile, C.Elm_Text_Format(format)) == eTrue
}

func (p *Entry) File() (string, TextFormat) {
	var (
		cfile  *C.char
		format C.Elm_Text_Format
	)

	C.elm_entry_file_get(p.obj, &cfile, &format)
	return C.GoString(cfile), TextFormat(format)
}

func (p *Entry) SetEntryText(txt string) {
	ctxt := C.CString(txt)
	defer free(ctxt)
	C.elm_entry_entry_set(p.obj, ctxt)
}

func (p *Entry) EntryText() string {
	return C.GoString(C.elm_entry_entry_get(p.obj))
}

func (p *Entry) SetFilterLimitSize(o EntryFilterLimitSize, txt string) {
	ctxt := C.CString(txt)
	defer free(ctxt) //test this is legal?
	C.elm_entry_filter_limit_size(unsafe.Pointer(&o.obj), p.obj, &ctxt)
}

func (p *Entry) SetAcceptFilter(o EntryFilterAcceptSet, txt string) {
	ctxt := C.CString(txt)
	defer free(ctxt) //test this is legal?
	C.elm_entry_filter_accept_set(unsafe.Pointer(&o.obj), p.obj, &ctxt)
}

func (p *Entry) AddContextMenuItem(label string, iconfile string, ictype IconType, handler Handler) {
	var (
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
	)

	if handler != nil {
		handler_id := registerHandler(handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	clabel := C.CString(label)
	ciconfile := C.CString(iconfile)
	defer free(clabel, ciconfile)
	C.elm_entry_context_menu_item_add(p.obj, clabel, ciconfile, C.Elm_Icon_Type(ictype), smartcb, data)
}

//void 	elm_entry_markup_filter_prepend (Elm_Entry *obj, Elm_Entry_Filter_Cb func, void *data)

/*
TODO:
is this byte[] or interface
elm_entry_imf_context_get

void 	elm_entry_item_provider_prepend (Elm_Entry *obj, Elm_Entry_Item_Provider_Cb func, void *data)
void 	elm_entry_item_provider_remove (Elm_Entry *obj, Elm_Entry_Item_Provider_Cb func, void *data)

void 	elm_entry_input_panel_imdata_set (Elm_Entry *obj, const void *data, int len)
void 	elm_entry_input_panel_imdata_get (const Elm_Entry *obj, void *data, int *len)

void 	elm_entry_markup_filter_remove (Elm_Entry *obj, Elm_Entry_Filter_Cb func, void *data)
void 	elm_entry_item_provider_append (Elm_Entry *obj, Elm_Entry_Item_Provider_Cb func, void *data)
void 	elm_entry_markup_filter_append (Elm_Entry *obj, Elm_Entry_Filter_Cb func, void *data)



*/
//------------------------------------------------------------

type Flip struct {
	*widgetBase
}

func NewFlip(parent Object) *Flip {
	return &Flip{wrapWidgetBase(C.elm_flip_add(parent.eo()))}
}

func (p *Flip) SetPerspective(foc, x, y int) {
	C.elm_flip_perspective_set(p.obj, C.Evas_Coord(foc), C.Evas_Coord(x), C.Evas_Coord(y))
}

func (p *Flip) SetInteraction(interaction FlipInteraction) {
	C.elm_flip_interaction_set(p.obj, C.Efl_Ui_Flip_Interaction(interaction))
}

func (p *Flip) Interaction() FlipInteraction {
	return FlipInteraction(C.elm_flip_interaction_get(p.obj))
}

func (p *Flip) IsFrontVisible() bool {
	return C.elm_flip_front_visible_get(p.obj) == eTrue
}

func (p *Flip) SetDirectionInteractionEnabled(dir FlipDirection, b bool) {
	C.elm_flip_interaction_direction_enabled_set(p.obj, C.Elm_Flip_Direction(dir), eBool(b))
}

func (p *Flip) SetDirectionInteractionHitsize(dir FlipDirection, size float64) {
	C.elm_flip_interaction_direction_hitsize_set(p.obj, C.Elm_Flip_Direction(dir), C.double(size))
}

func (p *Flip) DirectionInteractionHitsize(dir FlipDirection) float64 {
	return float64(C.elm_flip_interaction_direction_hitsize_get(p.obj, C.Elm_Flip_Direction(dir)))
}

func (p *Flip) Run(mode FlipMode) {
	C.elm_flip_go(p.obj, C.Efl_Ui_Flip_Mode(mode))
}

func (p *Flip) RunTo(front bool, mode FlipMode) {
	C.elm_flip_go_to(p.obj, eBool(front), C.Efl_Ui_Flip_Mode(mode))
}

//------------------------------------------------------------

type Hover struct {
	*Layout
}

func NewHover(parent Object) *Hover {
	return &Hover{wrapLayout(C.elm_hover_add(parent.eo()))}
}

func wrapHover(o *C.Eo) *Hover {
	if o != nil {
		return &Hover{wrapLayout(o)}
	}
	return nil
}

//TODO make Evas Canvas interface
func (p *Hover) SetTarget(canvas Canvas) {
	C.elm_hover_target_set(p.obj, canvas.eo())
}

func (p *Hover) Target() Canvas {
	return wrapEvas(C.elm_hover_target_get(p.obj))
}

func (p *Hover) BestContentLocation(axis HoverAxis) string {
	return C.GoString(C.elm_hover_best_content_location_get(p.obj, C.Elm_Hover_Axis(axis)))
}

func (p *Hover) Dismiss() {
	C.elm_hover_dismiss(p.obj)
}

func (p *Hover) SetParent(parent Widget) {
	C.elm_hover_parent_set(p.obj, parent.eo())
}

func (p *Hover) Parent() Widget {
	return wrapWidgetBase(C.elm_hover_parent_get(p.obj))
}

//------------------------------------------------------------

type Flipselector struct {
	*Layout
}

func NewFlipselector(parent Object) *Flipselector {
	return &Flipselector{wrapLayout(C.elm_flipselector_add(parent.eo()))}
}

func (p *Flipselector) Items() []*FlipselectorItem {
	return newListIterator(C.elm_flipselector_items_get(p.obj)).FlipselectorItemSlice()
}

func (p *Flipselector) FirstItem() *FlipselectorItem {
	return wrapFlipselectorItem(C.elm_flipselector_first_item_get(p.obj))
}

func (p *Flipselector) LastItem() *FlipselectorItem {
	return wrapFlipselectorItem(C.elm_flipselector_last_item_get(p.obj))
}

func (p *Flipselector) SelectedItem() *FlipselectorItem {
	return wrapFlipselectorItem(C.elm_flipselector_selected_item_get(p.obj))
}

func (p *Flipselector) FlipNext() {
	C.elm_flipselector_flip_next(p.obj)
}

func (p *Flipselector) FlipPrev() {
	C.elm_flipselector_flip_prev(p.obj)
}

func (p *Flipselector) SetFirstInterval(interval float64) {
	C.elm_flipselector_first_interval_set(p.obj, C.double(interval))
}

func (p *Flipselector) FirstInterval() float64 {
	return float64(C.elm_flipselector_first_interval_get(p.obj))
}

func (p *Flipselector) ItemFactory() *FlipselectorItemFactory {
	return NewFlipselectorItemFactory(p)
}

//------------------------------------------------------------

type Hoversel struct {
	*Button
}

func NewHoversel(parent Object) *Hoversel {
	return &Hoversel{wrapButton(C.elm_hoversel_add(parent.eo()))}
}

func (p *Hoversel) SetHorizontal(b bool) {
	C.elm_hoversel_horizontal_set(p.obj, eBool(b))
}

func (p *Hoversel) IsHorizontal() bool {
	return C.elm_hoversel_horizontal_get(p.obj) == eTrue
}

func (p *Hoversel) SetParent(parent Object) {
	C.elm_hoversel_hover_parent_set(p.obj, parent.eo())
}

func (p *Hoversel) Parent() Widget {
	return wrapWidgetBase(C.elm_hoversel_hover_parent_get(p.obj))
}

func (p *Hoversel) IsExpanded() bool {
	return C.elm_hoversel_expanded_get(p.obj) == eTrue
}

func (p *Hoversel) Items() []*HoverselItem {
	return newListIterator(C.elm_hoversel_items_get(p.obj)).HoverselItemSlice()
}

func (p *Hoversel) SetAutoUpdate(b bool) {
	C.elm_hoversel_auto_update_set(p.obj, eBool(b))
}

func (p *Hoversel) IsAutoUpdate() bool {
	return C.elm_hoversel_auto_update_get(p.obj) == eTrue
}

func (p *Hoversel) BeginHover() {
	C.elm_hoversel_hover_begin(p.obj)
}

func (p *Hoversel) EndHover() {
	C.elm_hoversel_hover_end(p.obj)
}

func (p *Hoversel) Clear() {
	C.elm_hoversel_clear(p.obj)
}

func (p *Hoversel) ItemFactory() *HoverselItemFactory {
	return NewHoverselItemFactory(p)
}

//------------------------------------------------------------

type Photocam struct {
	*widgetBase
}

func NewPhotocam(parent Object) *Photocam {
	return &Photocam{wrapWidgetBase(C.elm_photocam_add(parent.eo()))}
}

func (p *Photocam) SetPaused(b bool) {
	C.elm_photocam_paused_set(p.obj, eBool(b))
}

func (p *Photocam) IsPaused() bool {
	return C.elm_photocam_paused_get(p.obj) == eTrue
}

func (p *Photocam) SetGestureEnabled(b bool) {
	C.elm_photocam_gesture_enabled_set(p.obj, eBool(b))
}

func (p *Photocam) IsGestureEnabled() bool {
	return C.elm_photocam_gesture_enabled_get(p.obj) == eTrue
}

func (p *Photocam) SetZoomLevel(zoom float64) {
	C.elm_photocam_zoom_set(p.obj, C.double(zoom))
}

func (p *Photocam) ZoomLevel() float64 {
	return float64(C.elm_photocam_zoom_get(p.obj))
}

func (p *Photocam) SetZoomMode(mode PhotocamZoomMode) {
	C.elm_photocam_zoom_mode_set(p.obj, C.Elm_Photocam_Zoom_Mode(mode))
}

func (p *Photocam) ZoomMode() PhotocamZoomMode {
	return PhotocamZoomMode(C.elm_photocam_zoom_mode_get(p.obj))
}

func (p *Photocam) ImageRegion() (int, int, int, int) {
	var x, y, w, h C.int
	C.elm_photocam_image_region_get(p.obj, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *Photocam) InternalImage() Canvas {
	return wrapEvas(C.elm_photocam_internal_image_get(p.obj))
}

func (p *Photocam) ImageSize() (int, int) {
	var w, h C.int
	C.elm_photocam_image_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Photocam) SetImageOrient(orient ImageOrient) {
	C.elm_photocam_image_orient_set(p.obj, C.Evas_Image_Orient(orient))
}

func (p *Photocam) ImageOrient() ImageOrient {
	return ImageOrient(C.elm_photocam_image_orient_get(p.obj))
}

func (p *Photocam) ShowImageRegion(x, y, w, h int) {
	C.elm_photocam_image_region_show(p.obj, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (p *Photocam) BringInImageRegion(x, y, w, h int) {
	C.elm_photocam_image_region_bring_in(p.obj, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (p *Photocam) SetFile(imgFile string) error {
	cfile := C.CString(imgFile)
	defer free(cfile)
	le := LoadError(C.elm_photocam_file_set(p.obj, cfile))
	if le == LoadErrorNone {
		return nil
	}
	return le
}

func (p *Photocam) File() string {
	return C.GoString(C.elm_photocam_file_get(p.obj))
}

//------------------------------------------------------------

type Slideshow struct {
	*Layout
	//itemMap *slideshowItemMap
}

func NewSlideshow(parent Object) *Slideshow {
	eo := C.elm_slideshow_add(parent.eo())
	return wrapSlideshow(eo)
}

func wrapSlideshow(o *C.Eo) *Slideshow {
	if o != nil {
		return &Slideshow{
			Layout: wrapLayout(o),
			//itemMap = newSlideshowItemMap()
		}
	}
	return nil
}

func (p *Slideshow) SetCacheAfter(count int) {
	C.elm_slideshow_cache_after_set(p.obj, C.int(count))
}

func (p *Slideshow) CacheAfter() int {
	return int(C.elm_slideshow_cache_after_get(p.obj))
}

func (p *Slideshow) SetCacheBefore(count int) {
	C.elm_slideshow_cache_before_set(p.obj, C.int(count))
}

func (p *Slideshow) CacheBefore() int {
	return int(C.elm_slideshow_cache_before_get(p.obj))
}

func (p *Slideshow) SetLayout(layout string) {
	clayout := C.CString(layout)
	defer free(clayout)
	C.elm_slideshow_layout_set(p.obj, clayout)
}

func (p *Slideshow) LayoutName() string {
	return C.GoString(C.elm_slideshow_layout_get(p.obj))
}

func (p *Slideshow) SetTransition(transition string) {
	ctr := C.CString(transition)
	defer free(ctr)
	C.elm_slideshow_transition_set(p.obj, ctr)
}

func (p *Slideshow) Transition() string {
	return C.GoString(C.elm_slideshow_transition_get(p.obj))
}

func (p *Slideshow) SetLoop(b bool) {
	C.elm_slideshow_loop_set(p.obj, eBool(b))
}

func (p *Slideshow) IsLoop() bool {
	return C.elm_slideshow_loop_get(p.obj) == eTrue
}

func (p *Slideshow) SetTimeout(to float64) {
	C.elm_slideshow_timeout_set(p.obj, C.double(to))
}

func (p *Slideshow) Timeout() float64 {
	return float64(C.elm_slideshow_timeout_get(p.obj))
}

func (p *Slideshow) Items() []*SlideshowItem {
	return newListIterator(C.elm_slideshow_items_get(p.obj)).SlideshowItemSlice()
}

func (p *Slideshow) AvailableTransitions() []string {
	return newListIterator(C.elm_slideshow_transitions_get(p.obj)).StringSlice()
}

func (p *Slideshow) ItemsCount() uint {
	return uint(C.elm_slideshow_count_get(p.obj))
}

func (p *Slideshow) CurrentItem() *SlideshowItem {
	return wrapSlideshowItem(C.elm_slideshow_item_current_get(p.obj))
}

func (p *Slideshow) AvailableLayouts() []string {
	return newListIterator(C.elm_slideshow_layouts_get(p.obj)).StringSlice()
}

func (p *Slideshow) Previous() {
	C.elm_slideshow_previous(p.obj)
}

func (p *Slideshow) NthItem(n uint) *SlideshowItem {
	return wrapSlideshowItem(C.elm_slideshow_item_nth_get(p.obj, C.uint(n)))
}

func (p *Slideshow) Next() {
	C.elm_slideshow_next(p.obj)
}

func (p *Slideshow) Clear() {
	C.elm_slideshow_clear(p.obj)
}

func (p *Slideshow) ItemFactory(fn func(interface{}, *Slideshow) Object) *SlideshowItemFactory {
	return NewSlideshowItemFactory(p, fn)
}

//------------------------------------------------------------

type Toolbar struct {
	*Layout
}

func NewToolbar(parent Object) *Toolbar {
	return &Toolbar{wrapLayout(C.elm_toolbar_add(parent.eo()))}
}

func (p *Toolbar) SetHomogeneous(b bool) {
	C.elm_toolbar_homogeneous_set(p.obj, eBool(b))
}

func (p *Toolbar) IsHomogeneous() bool {
	return C.elm_toolbar_homogeneous_get(p.obj) == eTrue
}

func (p *Toolbar) SetAlign(align float64) {
	C.elm_toolbar_align_set(p.obj, C.double(align))
}

func (p *Toolbar) Align() float64 {
	return float64(C.elm_toolbar_align_get(p.obj))
}

func (p *Toolbar) SetSelectMode(mode SelectMode) {
	C.elm_toolbar_select_mode_set(p.obj, C.Elm_Object_Select_Mode(mode))
}

func (p *Toolbar) SelectMode() SelectMode {
	return SelectMode(C.elm_toolbar_select_mode_get(p.obj))
}

func (p *Toolbar) SetIconSize(size int) {
	C.elm_toolbar_icon_size_set(p.obj, C.int(size))
}

func (p *Toolbar) IconSize() int {
	return int(C.elm_toolbar_icon_size_get(p.obj))
}

func (p *Toolbar) SetShrinkMode(mode ToolbarShrinkMode) {
	C.elm_toolbar_shrink_mode_set(p.obj, C.Elm_Toolbar_Shrink_Mode(mode))
}

func (p *Toolbar) ShrinkMode() ToolbarShrinkMode {
	return ToolbarShrinkMode(C.elm_toolbar_shrink_mode_get(p.obj))
}

func (p *Toolbar) SetMenuParent(parent Widget) {
	C.elm_toolbar_menu_parent_set(p.obj, parent.eo())
}

func (p *Toolbar) MenuParent() Widget {
	return wrapWidgetBase(C.elm_toolbar_menu_parent_get(p.obj))
}

func (p *Toolbar) SetStandardPriority(priority int) {
	C.elm_toolbar_standard_priority_set(p.obj, C.int(priority))
}

func (p *Toolbar) StandardPriority() int {
	return int(C.elm_toolbar_standard_priority_get(p.obj))
}

func (p *Toolbar) MoreItem() *ToolbarItem {
	return wrapToolbarItem(C.elm_toolbar_more_item_get(p.obj))
}

func (p *Toolbar) ItemsCount() uint {
	return uint(C.elm_toolbar_items_count(p.obj))
}

func (p *Toolbar) ItemByLabel(lbl string) *ToolbarItem {
	clbl := C.CString(lbl)
	defer free(clbl)
	return wrapToolbarItem(C.elm_toolbar_item_find_by_label(p.obj, clbl))
}

func (p *Toolbar) SetReorderMode(b bool) {
	C.elm_toolbar_reorder_mode_set(p.obj, eBool(b))
}

func (p *Toolbar) IsReorderMode() bool {
	return C.elm_toolbar_reorder_mode_get(p.obj) == eTrue
}

func (p *Toolbar) SetTransverseExpanded(b bool) {
	C.elm_toolbar_transverse_expanded_set(p.obj, eBool(b))
}

func (p *Toolbar) IsTransverseExpanded() bool {
	return C.elm_toolbar_transverse_expanded_get(p.obj) == eTrue
}

func (p *Toolbar) SetIconOrderLookup(order IconLookupOrder) {
	C.elm_toolbar_icon_order_lookup_set(p.obj, C.Elm_Icon_Lookup_Order(order))
}

func (p *Toolbar) IconLookupOrder() IconLookupOrder {
	return IconLookupOrder(C.elm_toolbar_icon_order_lookup_get(p.obj))
}

func (p *Toolbar) SetHorizontal(b bool) {
	C.elm_toolbar_horizontal_set(p.obj, eBool(b))
}

func (p *Toolbar) IsHorizontal() bool {
	return C.elm_toolbar_horizontal_get(p.obj) == eTrue
}

func (p *Toolbar) SelectedItem() *ToolbarItem {
	return wrapToolbarItem(C.elm_toolbar_selected_item_get(p.obj))
}

func (p *Toolbar) FirstItem() *ToolbarItem {
	return wrapToolbarItem(C.elm_toolbar_first_item_get(p.obj))
}

func (p *Toolbar) LastItem() *ToolbarItem {
	return wrapToolbarItem(C.elm_toolbar_last_item_get(p.obj))
}

func (p *Toolbar) ItemFactory() *ToolbarItemFactory {
	return NewToolbarItemFactory(p)
}

//------------------------------------------------------------

type PanelScrollInfo struct {
	obj *C.Elm_Panel_Scroll_Info
}

func wrapPanelScrollInfo(o *C.Elm_Panel_Scroll_Info) *PanelScrollInfo {
	if o != nil {
		return &PanelScrollInfo{o}
	}
	return nil
}

func (p *PanelScrollInfo) RelX() float64 {
	return float64(p.obj.rel_x)
}

func (p *PanelScrollInfo) RelY() float64 {
	return float64(p.obj.rel_y)
}

type Panel struct {
	*Layout
}

func NewPanel(parent Object) *Panel {
	return &Panel{wrapLayout(C.elm_panel_add(parent.eo()))}
}

func (p *Panel) SetOrient(orient PanelOrient) {
	C.elm_panel_orient_set(p.obj, C.Elm_Panel_Orient(orient))
}

func (p *Panel) Orient() PanelOrient {
	return PanelOrient(C.elm_panel_orient_get(p.obj))
}

func (p *Panel) SetHidden(b bool) {
	C.elm_panel_hidden_set(p.obj, eBool(b))
}

func (p *Panel) IsHidden() bool {
	return C.elm_panel_hidden_get(p.obj) == eTrue
}

func (p *Panel) SetScrollable(b bool) {
	C.elm_panel_scrollable_set(p.obj, eBool(b))
}

func (p *Panel) IsScrollable() bool {
	return C.elm_panel_scrollable_get(p.obj) == eTrue
}

func (p *Panel) SetScrollableContentSize(ratio float64) {
	C.elm_panel_scrollable_content_size_set(p.obj, C.double(ratio))
}

/* in 1.19
func (p *Panel) ScrollableContentSize() float64 {
	return float64(C.elm_panel_scrollable_content_size_get(p.obj))
}
*/

func (p *Panel) Toggle() {
	C.elm_panel_toggle(p.obj)
}

//------------------------------------------------------------

type Ctxpopup struct {
	*Layout
}

func NewCtxpopup(parent Object) *Ctxpopup {
	return &Ctxpopup{wrapLayout(C.elm_ctxpopup_add(parent.eo()))}
}

func (p *Ctxpopup) SetHorizontal(b bool) {
	C.elm_ctxpopup_horizontal_set(p.obj, eBool(b))
}

func (p *Ctxpopup) IsHorizontal() bool {
	return C.elm_ctxpopup_horizontal_get(p.obj) == eTrue
}

func (p *Ctxpopup) Items() []*CtxpopupItem {
	return newListIterator(C.elm_ctxpopup_items_get(p.obj)).CtxpopupItemSlice()
}

func (p *Ctxpopup) FirstItem() *CtxpopupItem {
	return wrapCtxpopupItem(C.elm_ctxpopup_first_item_get(p.obj))
}

func (p *Ctxpopup) LastItem() *CtxpopupItem {
	return wrapCtxpopupItem(C.elm_ctxpopup_last_item_get(p.obj))
}

func (p *Ctxpopup) SetAutoHideDisabled(b bool) {
	C.elm_ctxpopup_auto_hide_disabled_set(p.obj, eBool(b))
}

func (p *Ctxpopup) IsAutoHideDisabled() bool {
	return C.elm_ctxpopup_auto_hide_disabled_get(p.obj) == eTrue
}

func (p *Ctxpopup) SetHoverParent(canvas Canvas) {
	C.elm_ctxpopup_hover_parent_set(p.obj, canvas.eo())
}

func (p *Ctxpopup) HoverParent() Canvas {
	return wrapEvas(C.elm_ctxpopup_hover_parent_get(p.obj))
}

func (p *Ctxpopup) SetDirectionPriority(first, second, third, fourth CtxpopupDirection) {
	C.elm_ctxpopup_direction_priority_set(p.obj, C.Elm_Ctxpopup_Direction(first), C.Elm_Ctxpopup_Direction(second), C.Elm_Ctxpopup_Direction(third), C.Elm_Ctxpopup_Direction(fourth))
}

func (p *Ctxpopup) DirectionPriority() (CtxpopupDirection, CtxpopupDirection, CtxpopupDirection, CtxpopupDirection) {
	var first, second, third, fourth C.Elm_Ctxpopup_Direction
	C.elm_ctxpopup_direction_priority_get(p.obj, &first, &second, &third, &fourth)
	return CtxpopupDirection(first), CtxpopupDirection(second), CtxpopupDirection(third), CtxpopupDirection(fourth)
}

func (p *Ctxpopup) Direction() CtxpopupDirection {
	return CtxpopupDirection(C.elm_ctxpopup_direction_get(p.obj))
}

func (p *Ctxpopup) Dismiss() {
	C.elm_ctxpopup_dismiss(p.obj)
}

func (p *Ctxpopup) Clear() {
	C.elm_ctxpopup_clear(p.obj)
}

func (p *Ctxpopup) ItemFactory() *CtxpopupItemFactory {
	return NewCtxpopupItemFactory(p)
}

//------------------------------------------------------------

type Index struct {
	*Layout
}

func NewIndex(parent Object) *Index {
	return &Index{wrapLayout(C.elm_index_add(parent.eo()))}
}

func (p *Index) SetAutohideDisabled(b bool) {
	C.elm_index_autohide_disabled_set(p.obj, eBool(b))
}

func (p *Index) IsAutohideDisabled() bool {
	return C.elm_index_autohide_disabled_get(p.obj) == eTrue
}

func (p *Index) SetOmitEnabled(b bool) {
	C.elm_index_omit_enabled_set(p.obj, eBool(b))
}

func (p *Index) IsOmitEnabled() bool {
	return C.elm_index_omit_enabled_get(p.obj) == eTrue
}

func (p *Index) SetStandardPriority(pri int) {
	C.elm_index_standard_priority_set(p.obj, C.int(pri))
}

func (p *Index) StandardPriority() int {
	return int(C.elm_index_standard_priority_get(p.obj))
}

func (p *Index) SetDelayChangeTime(deltm float64) {
	C.elm_index_delay_change_time_set(p.obj, C.double(deltm))
}

func (p *Index) DelayChangeTime() float64 {
	return float64(C.elm_index_delay_change_time_get(p.obj))
}

func (p *Index) SetIndicatorDisabled(b bool) {
	C.elm_index_indicator_disabled_set(p.obj, eBool(b))
}

func (p *Index) IsIndicatorDisabled() bool {
	return C.elm_index_indicator_disabled_get(p.obj) == eTrue
}

func (p *Index) SetItemLevel(level int) {
	C.elm_index_item_level_set(p.obj, C.int(level))
}

func (p *Index) ItemLevel() int {
	return int(C.elm_index_item_level_get(p.obj))
}

func (p *Index) FlushLevel(level int) {
	C.elm_index_level_go(p.obj, C.int(level))
}

func (p *Index) ClearItem() {
	C.elm_index_item_clear(p.obj)
}

func (p *Index) SelectedItem(level int) *IndexItem {
	return wrapIndexItem(C.elm_index_selected_item_get(p.obj, C.int(level)))
}

func (p *Index) SetHorizontal(b bool) {
	C.elm_index_horizontal_set(p.obj, eBool(b))
}

func (p *Index) IsHorizontal() bool {
	return C.elm_index_horizontal_get(p.obj) == eTrue
}

//TODO this is not legal, make map
func (p *Index) Find(data interface{}) *IndexItem {
	return wrapIndexItem(C.elm_index_item_find(p.obj, unsafe.Pointer(&data)))
}

func (p *Index) ItemFactory() *IndexItemFactory {
	return NewIndexItemFactory(p)
}

//------------------------------------------------------------

type List struct {
	*Layout
}

func NewList(parent Object) *List {
	return &List{wrapLayout(C.elm_list_add(parent.eo()))}
}

func (p *List) Start() {
	C.elm_list_go(p.obj)
}

func (p *List) SetHorizontal(b bool) {
	C.elm_list_horizontal_set(p.obj, eBool(b))
}

func (p *List) IsHorizontal() bool {
	return C.elm_list_horizontal_get(p.obj) == eTrue
}

func (p *List) SetSelectMode(mode SelectMode) {
	C.elm_list_select_mode_set(p.obj, C.Elm_Object_Select_Mode(mode))
}

func (p *List) SelectMode() SelectMode {
	return SelectMode(C.elm_list_select_mode_get(p.obj))
}

func (p *List) SetFocusOnSelection(b bool) {
	C.elm_list_focus_on_selection_set(p.obj, eBool(b))
}

func (p *List) IsFocusOnSelection() bool {
	return C.elm_list_focus_on_selection_get(p.obj) == eTrue
}

func (p *List) SetMultiSelect(b bool) {
	C.elm_list_multi_select_set(p.obj, eBool(b))
}

func (p *List) IsMultiSelect() bool {
	return C.elm_list_multi_select_get(p.obj) == eTrue
}

func (p *List) SetMultiSelectMode(mode MultiSelectMode) {
	C.elm_list_multi_select_mode_set(p.obj, C.Elm_Object_Multi_Select_Mode(mode))
}

func (p *List) MultiSelectMode() MultiSelectMode {
	return MultiSelectMode(C.elm_list_multi_select_mode_get(p.obj))
}

func (p *List) SetMode(mode ListMode) {
	C.elm_list_mode_set(p.obj, C.Elm_List_Mode(mode))
}

func (p *List) Mode() ListMode {
	return ListMode(C.elm_list_mode_get(p.obj))
}

func (p *List) SelectedItem() *ListItem {
	return wrapListItem(C.elm_list_selected_item_get(p.obj))
}

func (p *List) Items() []*ListItem {
	return newListIterator(C.elm_list_items_get(p.obj)).ListItemSlice()
}

func (p *List) FirstItem() *ListItem {
	return wrapListItem(C.elm_list_first_item_get(p.obj))
}

func (p *List) SelectedItems() []*ListItem {
	return newListIterator(C.elm_list_selected_items_get(p.obj)).ListItemSlice()
}

func (p *List) LastItem() *ListItem {
	return wrapListItem(C.elm_list_last_item_get(p.obj))
}

func (p *List) ItemAtXY(x, y int) (*ListItem, int) {
	var pos_relative C.int
	wi := C.elm_list_at_xy_item_get(p.obj, C.int(x), C.int(y), &pos_relative)
	return wrapListItem(wi), int(pos_relative)
}

func (p *List) Clear() {
	C.elm_list_clear(p.obj)
}

func (p *List) ItemFactory() *ListItemFactory {
	return NewListItemFactory(p)
}

/* DELETE
func (p *List) OnDoubleClicked(fn func(*ListItem)) {
	p.OnEvent("clicked,double", fn)
}
*/

//------------------------------------------------------------

type Menu struct {
	*widgetBase
}

func NewMenu(parent Object) *Menu {
	return &Menu{wrapWidgetBase(C.elm_menu_add(parent.eo()))}
}

func wrapMenu(o *C.Eo) *Menu {
	if o != nil {
		return &Menu{wrapWidgetBase(o)}
	}
	return nil
}

func (p *Menu) Move(x, y int) {
	C.elm_menu_move(p.obj, C.int(x), C.int(y))
}

func (p *Menu) Close() {
	C.elm_menu_close(p.obj)
}

func (p *Menu) AddSeperator(parent *MenuItem) *MenuItem {
	return wrapMenuItem(C.elm_menu_item_separator_add(p.obj, parent.obj))
}

func (p *Menu) SetParent(parent Object) {
	C.elm_menu_parent_set(p.obj, parent.eo())
}

func (p *Menu) Parent() Widget {
	return wrapWidgetBase(C.elm_menu_parent_get(p.obj))
}

func (p *Menu) SelectedItem() *MenuItem {
	return wrapMenuItem(C.elm_menu_selected_item_get(p.obj))
}

func (p *Menu) Items() []*MenuItem {
	return newListIterator(C.elm_menu_items_get(p.obj)).MenuItemSlice()
}

func (p *Menu) FirstItem() *MenuItem {
	return wrapMenuItem(C.elm_menu_first_item_get(p.obj))
}

func (p *Menu) LastItem() *MenuItem {
	return wrapMenuItem(C.elm_menu_last_item_get(p.obj))
}

func (p *Menu) ItemFactory() *MenuItemFactory {
	return NewMenuItemFactory(p)
}

//------------------------------------------------------------

type Mapbuf struct {
	*widgetBase
}

func NewMapbuf(parent Object) *Mapbuf {
	return &Mapbuf{wrapWidgetBase(C.elm_mapbuf_add(parent.eo()))}
}

func (p *Mapbuf) SetAuto(b bool) {
	C.elm_mapbuf_auto_set(p.obj, eBool(b))
}

func (p *Mapbuf) IsAuto() bool {
	return C.elm_mapbuf_auto_get(p.obj) == eTrue
}

func (p *Mapbuf) SetSmooth(b bool) {
	C.elm_mapbuf_smooth_set(p.obj, eBool(b))
}

func (p *Mapbuf) IsSmooth() bool {
	return C.elm_mapbuf_smooth_get(p.obj) == eTrue
}

func (p *Mapbuf) SetAlpha(b bool) {
	C.elm_mapbuf_alpha_set(p.obj, eBool(b))
}

func (p *Mapbuf) IsAlpha() bool {
	return C.elm_mapbuf_alpha_get(p.obj) == eTrue
}

func (p *Mapbuf) SetEnabled(b bool) {
	C.elm_mapbuf_enabled_set(p.obj, eBool(b))
}

func (p *Mapbuf) IsEnabled() bool {
	return C.elm_mapbuf_enabled_get(p.obj) == eTrue
}

func (p *Mapbuf) SetPointColor(index, r, g, b, a int) {
	C.elm_mapbuf_point_color_set(p.obj, C.int(index), C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *Mapbuf) PointColor(index int) (int, int, int, int) {
	var r, g, b, a C.int
	C.elm_mapbuf_point_color_get(p.obj, C.int(index), &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

//------------------------------------------------------------

type Thumb struct {
	*widgetBase
}

func NewThumb(parent Object) *Thumb {
	return &Thumb{wrapWidgetBase(C.elm_thumb_add(parent.eo()))}
}

func (p *Thumb) SetFile(file, key string) {
	cfile := C.CString(file)
	ckey := C.CString(key)
	defer free(cfile, ckey)
	C.elm_thumb_file_set(p.obj, cfile, ckey)
}

func (p *Thumb) File() (string, string) {
	var file, key *C.char
	C.elm_thumb_file_get(p.obj, &file, &key)
	return C.GoString(file), C.GoString(key)
}

func (p *Thumb) SetEditable(b bool) {
	C.elm_thumb_editable_set(p.obj, eBool(b))
}

func (p *Thumb) IsEditable() bool {
	return C.elm_thumb_editable_get(p.obj) == eTrue
}

func (p *Thumb) SetCompress(comp int) {
	C.elm_thumb_compress_set(p.obj, C.int(comp))
}

func (p *Thumb) Compress() int {
	var comp C.int
	C.elm_thumb_compress_get(p.obj, &comp)
	return int(comp)
}

func (p *Thumb) SetFormat(format ThumbFormat) {
	C.elm_thumb_format_set(p.obj, C.Ethumb_Thumb_Format(format))
}

func (p *Thumb) Format() ThumbFormat {
	return ThumbFormat(C.elm_thumb_format_get(p.obj))
}

func (p *Thumb) SetAnimateState(state ThumbAnimationState) {
	C.elm_thumb_animate_set(p.obj, C.Elm_Thumb_Animation_Setting(state))
}

func (p *Thumb) AnimateState() ThumbAnimationState {
	return ThumbAnimationState(C.elm_thumb_animate_get(p.obj))
}

func (p *Thumb) SetFDOSize(typ ThumbFDOSizeType) {
	C.elm_thumb_fdo_size_set(p.obj, C.Ethumb_Thumb_FDO_Size(typ))
}

func (p *Thumb) FDOSize() ThumbFDOSizeType {
	return ThumbFDOSizeType(C.elm_thumb_fdo_size_get(p.obj))
}

func (p *Thumb) SetOrientation(o ThumbOrientation) {
	C.elm_thumb_orientation_set(p.obj, C.Ethumb_Thumb_Orientation(o))
}

func (p *Thumb) Orientation() ThumbOrientation {
	return ThumbOrientation(C.elm_thumb_orientation_get(p.obj))
}

func (p *Thumb) SetAspect(a ThumbAspect) {
	C.elm_thumb_aspect_set(p.obj, C.Ethumb_Thumb_Aspect(a))
}

func (p *Thumb) Aspect() ThumbAspect {
	return ThumbAspect(C.elm_thumb_aspect_get(p.obj))
}

func (p *Thumb) SetQuality(q int) {
	C.elm_thumb_quality_set(p.obj, C.int(q))
}

func (p *Thumb) Quality() int {
	var q C.int
	C.elm_thumb_quality_get(p.obj, &q)
	return int(q)
}

func (p *Thumb) SetSize(w, h int) {
	C.elm_thumb_size_set(p.obj, C.int(w), C.int(h))
}

func (p *Thumb) Size() (int, int) {
	var w, h C.int
	C.elm_thumb_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Thumb) SetCropAlign(x, y float64) {
	C.elm_thumb_crop_align_set(p.obj, C.double(x), C.double(y))
}

func (p *Thumb) CropAlign() (float64, float64) {
	var x, y C.double
	C.elm_thumb_crop_align_get(p.obj, &x, &y)
	return float64(x), float64(y)
}

func (p *Thumb) Path() (string, string) {
	var file, key *C.char
	C.elm_thumb_path_get(p.obj, &file, &key)
	return C.GoString(file), C.GoString(key)
}

func (p *Thumb) Reload() {
	C.elm_thumb_reload(p.obj)
}

/*
TODO:
void * 	elm_thumb_ethumb_client_get (void)
Eina_Bool 	elm_thumb_ethumb_client_connected_get (void)

*/

//------------------------------------------------------------

type Multibuttonentry struct {
	*Layout
}

func NewMultibuttonentry(parent Object) *Multibuttonentry {
	return &Multibuttonentry{wrapLayout(C.elm_multibuttonentry_add(parent.eo()))}
}

func (p *Multibuttonentry) Entry() *Entry {
	return wrapEntry(C.elm_multibuttonentry_entry_get(p.obj))
}

func (p *Multibuttonentry) IsExpanded() bool {
	return C.elm_multibuttonentry_expanded_get(p.obj) == eTrue
}

func (p *Multibuttonentry) SetExpanded(b bool) {
	C.elm_multibuttonentry_expanded_set(p.obj, eBool(b))
}

func (p *Multibuttonentry) Items() []*MultibuttonentryItem {
	return newListIterator(C.elm_multibuttonentry_items_get(p.obj)).MultibuttonentryItemSlice()
}

func (p *Multibuttonentry) FirstItem() *MultibuttonentryItem {
	return wrapMultibuttonentryItem(C.elm_multibuttonentry_first_item_get(p.obj))
}

func (p *Multibuttonentry) LastItem() *MultibuttonentryItem {
	return wrapMultibuttonentryItem(C.elm_multibuttonentry_last_item_get(p.obj))
}

func (p *Multibuttonentry) SelectedItem() *MultibuttonentryItem {
	return wrapMultibuttonentryItem(C.elm_multibuttonentry_selected_item_get(p.obj))
}

func (p *Multibuttonentry) Clear() {
	C.elm_multibuttonentry_clear(p.obj)
}

func (p *Multibuttonentry) SetEditable(b bool) {
	C.elm_multibuttonentry_editable_set(p.obj, eBool(b))
}

func (p *Multibuttonentry) IsEditable() bool {
	return C.elm_multibuttonentry_editable_get(p.obj) == eTrue
}

func (p *Multibuttonentry) ItemFactory() *MultibuttonentryItemFactory {
	return NewMultibuttonentryItemFactory(p)
}

/*
TODO: requires callbacks

void                     elm_multibuttonentry_item_filter_append(Evas_Object *obj, Elm_Multibuttonentry_Item_Filter_Cb func, void *data)
void                     elm_multibuttonentry_item_filter_prepend(Evas_Object *obj, Elm_Multibuttonentry_Item_Filter_Cb func, void *data)
void                     elm_multibuttonentry_item_filter_remove(Evas_Object *obj, Elm_Multibuttonentry_Item_Filter_Cb func, void *data)
void                    elm_multibuttonentry_format_function_set(Evas_Object *obj, Elm_Multibuttonentry_Format_Cb f_func, const void *data)
*/

//------------------------------------------------------------

type Diskselector struct {
	*widgetBase
}

func NewDiskselector(parent Object) *Diskselector {
	return &Diskselector{wrapWidgetBase(C.elm_diskselector_add(parent.eo()))}
}

func (p *Diskselector) SetRoundEnabled(b bool) {
	C.elm_diskselector_round_enabled_set(p.obj, eBool(b))
}

func (p *Diskselector) IsRoundEnabled() bool {
	return C.elm_diskselector_round_enabled_get(p.obj) == eTrue
}

func (p *Diskselector) SideTextMaxLength() int {
	return int(C.elm_diskselector_side_text_max_length_get(p.obj))
}

func (p *Diskselector) SetSideTextMaxLength(l int) {
	C.elm_diskselector_side_text_max_length_set(p.obj, C.int(l))
}

func (p *Diskselector) SetDisplayItemCount(c int) {
	C.elm_diskselector_display_item_num_set(p.obj, C.int(c))
}

func (p *Diskselector) DisplayItemCount() int {
	return int(C.elm_diskselector_display_item_num_get(p.obj))
}

func (p *Diskselector) Clear() {
	C.elm_diskselector_clear(p.obj)
}

func (p *Diskselector) Items() []*DiskselectorItem {
	return newListIterator(C.elm_diskselector_items_get(p.obj)).DiskselectorItemSlice()
}

func (p *Diskselector) SelectedItem() *DiskselectorItem {
	return wrapDiskselectorItem(C.elm_diskselector_selected_item_get(p.obj))
}

func (p *Diskselector) FirstItem() *DiskselectorItem {
	return wrapDiskselectorItem(C.elm_diskselector_first_item_get(p.obj))
}

func (p *Diskselector) LastItem() *DiskselectorItem {
	return wrapDiskselectorItem(C.elm_diskselector_last_item_get(p.obj))
}

/*
TODO:
Elm_Object_Item *elm_diskselector_item_append(Evas_Object *obj, const char *label, Evas_Object *icon, Evas_Smart_Cb func, void *data)
*/

//------------------------------------------------------------

type MapOverlay struct {
	obj *C.Elm_Map_Overlay
}

func NewMapOverlay(m *Map, lon, lat float64) *MapOverlay {
	return &MapOverlay{
		C.elm_map_overlay_add(m.obj, C.double(lon), C.double(lat)),
	}
}

func wrapMapOverlay(o *C.Elm_Map_Overlay) *MapOverlay {
	if o != nil {
		return &MapOverlay{o}
	}
	return nil
}

func (p *MapOverlay) Del() {
	C.elm_map_overlay_del(p.obj)
}

func (p *MapOverlay) Type() MapOverlayType {
	return MapOverlayType(C.elm_map_overlay_type_get(p.obj))
}

//TODO this is not legal, make map
func (p *MapOverlay) SetData(data interface{}) {
	C.elm_map_overlay_data_set(p.obj, unsafe.Pointer(&data))
}

func (p *MapOverlay) Data() interface{} {
	return unsafe.Pointer(C.elm_map_overlay_data_get(p.obj))
}

func (p *MapOverlay) SetHidden(b bool) {
	C.elm_map_overlay_hide_set(p.obj, eBool(b))
}

func (p *MapOverlay) IsHidden() bool {
	return C.elm_map_overlay_hide_get(p.obj) == eTrue
}

func (p *MapOverlay) SetDisplayedZoomMin(zm int) {
	C.elm_map_overlay_displayed_zoom_min_set(p.obj, C.int(zm))
}

func (p *MapOverlay) DisplayedZoomMin() int {
	return int(C.elm_map_overlay_displayed_zoom_min_get(p.obj))
}

func (p *MapOverlay) SetPaused(b bool) {
	C.elm_map_overlay_paused_set(p.obj, eBool(b))
}

func (p *MapOverlay) IsPaused() bool {
	return C.elm_map_overlay_paused_get(p.obj) == eTrue
}

func (p *MapOverlay) IsVisible() bool {
	return C.elm_map_overlay_visible_get(p.obj) == eTrue
}

func (p *MapOverlay) SetContent(co Object) {
	C.elm_map_overlay_content_set(p.obj, co.eo())
}

func (p *MapOverlay) Content() Object {
	return wrapObjectBase(C.elm_map_overlay_content_get(p.obj))
}

func (p *MapOverlay) SetIcon(icon *Icon) {
	C.elm_map_overlay_icon_set(p.obj, icon.obj)
}

func (p *MapOverlay) Icon() *Icon {
	return wrapIcon(C.elm_map_overlay_icon_get(p.obj))
}

func (p *MapOverlay) SetRegion(lon, lat float64) {
	C.elm_map_overlay_region_set(p.obj, C.double(lon), C.double(lat))
}

func (p *MapOverlay) Region() (float64, float64) {
	var lon, lat C.double
	C.elm_map_overlay_region_get(p.obj, &lon, &lat)
	return float64(lon), float64(lat)
}

func (p *MapOverlay) SetColor(r, g, b, a int) {
	C.elm_map_overlay_color_set(p.obj, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *MapOverlay) Color() (int, int, int, int) {
	var r, g, b, a C.int
	C.elm_map_overlay_color_get(p.obj, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *MapOverlay) Show() {
	C.elm_map_overlay_show(p.obj)
}

func (p *MapOverlay) GroupMembers() []*MapOverlay {
	return newListIterator(C.elm_map_overlay_group_members_get(p.obj)).MapOverlaySlice()
}

/*
TODO:
void elm_map_overlay_get_cb_set(Elm_Map_Overlay *overlay, Elm_Map_Overlay_Get_Cb get_cb, void *data)
void elm_map_overlay_del_cb_set(Elm_Map_Overlay *overlay, Elm_Map_Overlay_Del_Cb del_cb, void *data)
*/

type MapOverlayClass struct {
	*MapOverlay
}

func NewMapOverlayClass(m *Map) *MapOverlayClass {
	return &MapOverlayClass{
		wrapMapOverlay(C.elm_map_overlay_class_add(m.obj)),
	}
}

func (p *MapOverlayClass) Append(overlay MapOverlay) {
	C.elm_map_overlay_class_append(p.obj, overlay.obj)
}

func (p *MapOverlayClass) Remove(overlay MapOverlay) {
	C.elm_map_overlay_class_remove(p.obj, overlay.obj)
}

func (p *MapOverlayClass) SetZoomMax(zm int) {
	C.elm_map_overlay_class_zoom_max_set(p.obj, C.int(zm))
}

func (p *MapOverlayClass) ZoomMax() int {
	return int(C.elm_map_overlay_class_zoom_max_get(p.obj))
}

//------------------------------------------------------------

type MapOverlayBubble struct {
	*MapOverlay
}

func NewMapOverlayBubble(m *Map) *MapOverlayBubble {
	return &MapOverlayBubble{
		wrapMapOverlay(C.elm_map_overlay_bubble_add(m.obj)),
	}
}

func (p *MapOverlayBubble) Follow(overlay MapOverlay) {
	C.elm_map_overlay_bubble_follow(p.obj, overlay.obj)
}

func (p *MapOverlayBubble) AppendContent(content Object) {
	C.elm_map_overlay_bubble_content_append(p.obj, content.eo())
}

func (p *MapOverlayBubble) ClearContent() {
	C.elm_map_overlay_bubble_content_clear(p.obj)
}

//------------------------------------------------------------

type MapOverlayLine struct {
	*MapOverlay
}

func NewMapOverlayLine(m *Map, flon, flat, tlon, tlat float64) *MapOverlayLine {
	return &MapOverlayLine{
		wrapMapOverlay(C.elm_map_overlay_line_add(m.obj, C.double(flon), C.double(flat), C.double(tlon), C.double(tlat))),
	}
}

type MapOverlayPolygon struct {
	*MapOverlay
}

func NewMapOverlayPolygon(m *Map) *MapOverlayPolygon {
	return &MapOverlayPolygon{
		wrapMapOverlay(C.elm_map_overlay_polygon_add(m.obj)),
	}
}

func (p *MapOverlayPolygon) AddRegion(lon, lat float64) {
	C.elm_map_overlay_polygon_region_add(p.obj, C.double(lon), C.double(lat))
}

//------------------------------------------------------------

type MapOverlayCircle struct {
	*MapOverlay
}

func NewMapOverlayCircle(m *Map, lon, lat, radius float64) *MapOverlayCircle {
	return &MapOverlayCircle{
		wrapMapOverlay(C.elm_map_overlay_circle_add(m.obj, C.double(lon), C.double(lat), C.double(radius))),
	}
}

//------------------------------------------------------------

type MapOverlayScale struct {
	*MapOverlay
}

func NewMapOverlayScale(m *Map, x, y int) *MapOverlayScale {
	return &MapOverlayScale{
		wrapMapOverlay(C.elm_map_overlay_scale_add(m.obj, C.int(x), C.int(y))),
	}
}

//------------------------------------------------------------

type MapRoute struct {
	obj *C.Elm_Map_Route
}

func wrapMapRoute(o *C.Elm_Map_Route) *MapRoute {
	if o != nil {
		return &MapRoute{o}
	}
	return nil
}

func (p *MapRoute) Del() {
	C.elm_map_route_del(p.obj)
}

func (p *MapRoute) Distance() float64 {
	return float64(C.elm_map_route_distance_get(p.obj))
}

func (p *MapRoute) Node() string {
	return C.GoString(C.elm_map_route_node_get(p.obj))
}

func (p *MapRoute) Waypoint() string {
	return C.GoString(C.elm_map_route_waypoint_get(p.obj))
}

//------------------------------------------------------------

type MapOverlayRoute struct {
	*MapOverlay
}

func NewMapOverlayRoute(m *Map, route *MapRoute) *MapOverlayRoute {
	return &MapOverlayRoute{
		wrapMapOverlay(C.elm_map_overlay_route_add(m.obj, route.obj)),
	}
}

//------------------------------------------------------------

type MapName struct {
	obj *C.Elm_Map_Name
}

func wrapMapName(o *C.Elm_Map_Name) *MapName {
	if o != nil {
		return &MapName{o}
	}
	return nil
}

func (p *MapName) Del() {
	C.elm_map_name_del(p.obj)
}

func (p *MapName) Address() string {
	return C.GoString(C.elm_map_name_address_get(p.obj))
}

func (p *MapName) Region() (float64, float64) {
	var lon, lat C.double
	C.elm_map_name_region_get(p.obj, &lon, &lat)
	return float64(lon), float64(lat)
}

//------------------------------------------------------------

type Map struct {
	*widgetBase
}

func NewMap(parent Object) *Map {
	return &Map{wrapWidgetBase(C.elm_map_add(parent.eo()))}
}

func (p *Map) SetZoom(zoom int) {
	C.elm_map_zoom_set(p.obj, C.int(zoom))
}

func (p *Map) Zoom() int {
	return int(C.elm_map_zoom_get(p.obj))
}

func (p *Map) SetPaused(b bool) {
	C.elm_map_paused_set(p.obj, eBool(b))
}

func (p *Map) IsPaused() bool {
	return C.elm_map_paused_get(p.obj) == eTrue
}

func (p *Map) SetWheelDisabled(b bool) {
	C.elm_map_wheel_disabled_set(p.obj, eBool(b))
}

func (p *Map) IsWheelDisabled() bool {
	return C.elm_map_wheel_disabled_get(p.obj) == eTrue
}

func (p *Map) SetMinZoom(mz int) {
	C.elm_map_zoom_min_set(p.obj, C.int(mz))
}

func (p *Map) MinZoom() int {
	return int(C.elm_map_zoom_min_get(p.obj))
}

func (p *Map) SetRotate(degree float64, cx, cy int) {
	C.elm_map_rotate_set(p.obj, C.double(degree), C.int(cx), C.int(cy))
}

func (p *Map) Rotate() (float64, int, int) {
	var (
		degree C.double
		cx, cy C.int
	)
	C.elm_map_rotate_get(p.obj, &degree, &cx, &cy)
	return float64(degree), int(cx), int(cy)
}

func (p *Map) SetUserAgent(agent string) {
	cagent := C.CString(agent)
	defer free(cagent)
	C.elm_map_user_agent_set(p.obj, cagent)
}

func (p *Map) UserAgent() string {
	return C.GoString(C.elm_map_user_agent_get(p.obj))
}

func (p *Map) SetMaxZoom(mz int) {
	C.elm_map_zoom_max_set(p.obj, C.int(mz))
}

func (p *Map) MaxZoom() int {
	return int(C.elm_map_zoom_max_get(p.obj))
}

func (p *Map) SetZoomMode(mode MapZoomMode) {
	C.elm_map_zoom_mode_set(p.obj, C.Elm_Map_Zoom_Mode(mode))
}

func (p *Map) ZoomMode() MapZoomMode {
	return MapZoomMode(C.elm_map_zoom_mode_get(p.obj))
}

func (p *Map) Region() (float64, float64) {
	var lon, lat C.double
	C.elm_map_region_get(p.obj, &lon, &lat)
	return float64(lon), float64(lat)
}

func (p *Map) TileLoadStatus() (int, int) {
	var try, finish C.int
	C.elm_map_tile_load_status_get(p.obj, &try, &finish)
	return int(try), int(finish)
}

func (p *Map) SetSource(t MapSourceType, src_name string) {
	csrc_name := C.CString(src_name)
	defer free(csrc_name)
	C.elm_map_source_set(p.obj, C.Elm_Map_Source_Type(t), csrc_name)
}

func (p *Map) Source(t MapSourceType) string {
	return C.GoString(C.elm_map_source_get(p.obj, C.Elm_Map_Source_Type(t)))
}

func (p *Map) ConvertRegionToCanvas(lon, lat float64) (int, int) {
	var x, y C.int
	C.elm_map_region_to_canvas_convert(p.obj, C.double(lon), C.double(lat), &x, &y)
	return int(x), int(y)
}

func (p *Map) ConvertCanvasToRegion(x, y int) (float64, float64) {
	var lon, lat C.double
	C.elm_map_canvas_to_region_convert(p.obj, C.int(x), C.int(y), &lon, &lat)
	return float64(lon), float64(lat)
}

func (p *Map) ShowRegion(lon, lat float64) {
	C.elm_map_region_show(p.obj, C.double(lon), C.double(lat))
}

func (p *Map) BringInRegion(lon, lat float64) {
	C.elm_map_region_bring_in(p.obj, C.double(lon), C.double(lat))
}

func (p *Map) BringInZoom(zoom int, lon, lat float64) {
	C.elm_map_region_zoom_bring_in(p.obj, C.int(zoom), C.double(lon), C.double(lat))
}

func (p *Map) RemoveTrack(track Object) {
	C.elm_map_track_remove(p.obj, track.eo())
}

/*
TODO:
void                     elm_map_overlays_show(Eina_List *overlays)
Eina_List * 	elm_map_overlays_get (const Elm_Map *obj)
Elm_Map_Route * 	elm_map_route_add (Elm_Map *obj, Elm_Map_Route_Type type, Elm_Map_Route_Method method, double flon, double flat, double tlon, double tlat, Elm_Map_Route_Cb route_cb, void *data)
Efl_Canvas_Object * 	elm_map_track_add (Elm_Map *obj, void *emap)

Elm_Map_Route           *elm_map_route_add(Evas_Object *obj, Elm_Map_Route_Type type, Elm_Map_Route_Method method, double flon, double flat, double tlon, double tlat, Elm_Map_Route_Cb route_cb, void *data)
Elm_Map_Name            *elm_map_name_add(const Evas_Object *obj, const char *address, double lon, double lat, Elm_Map_Name_Cb name_cb, void *data)

const char ** 	elm_map_sources_get (const Elm_Map *obj, Elm_Map_Source_Type type)


*/

//------------------------------------------------------------

type Player struct {
	*Layout
}

func NewPlayer(parent Object) *Player {
	return &Player{wrapLayout(C.elm_player_add(parent.eo()))}
}

func (p *Player) SetVideo(video *Video) {
	p.SetContent(video)
}

type Video struct {
	*Layout
}

func NewVideo(parent Object) *Video {
	return &Video{wrapLayout(C.elm_video_add(parent.eo()))}
}

func (p *Video) SetFile(file string) bool {
	cfile := C.CString(file)
	defer free(cfile)
	return C.elm_video_file_set(p.obj, cfile) == eTrue
}

func (p *Video) File() string {
	var cfile *C.char
	C.elm_video_file_get(p.obj, &cfile)
	return C.GoString(cfile)
}

func (p *Video) Emotion() *Emotion {
	return wrapEmotion(C.elm_video_emotion_get(p.obj))
}

func (p *Video) Play() {
	C.elm_video_play(p.obj)
}

func (p *Video) Pause() {
	C.elm_video_pause(p.obj)
}

func (p *Video) Stop() {
	C.elm_video_stop(p.obj)
}

func (p *Video) IsPlaying() bool {
	return C.elm_video_is_playing_get(p.obj) == eTrue
}

func (p *Video) IsSeekable() bool {
	return C.elm_video_is_seekable_get(p.obj) == eTrue
}

func (p *Video) IsAudioMute() bool {
	return C.elm_video_audio_mute_get(p.obj) == eTrue
}

func (p *Video) SetAudioMute(b bool) {
	C.elm_video_audio_mute_set(p.obj, eBool(b))
}

func (p *Video) AudioLevel() float64 {
	return float64(C.elm_video_audio_level_get(p.obj))
}

func (p *Video) SetAudioLevel(level float64) {
	C.elm_video_audio_level_set(p.obj, C.double(level))
}

func (p *Video) PlayPosition() float64 {
	return float64(C.elm_video_play_position_get(p.obj))
}

func (p *Video) SetPlayPosition(pos float64) {
	C.elm_video_play_position_set(p.obj, C.double(pos))
}

func (p *Video) PlayLength() float64 {
	return float64(C.elm_video_play_length_get(p.obj))
}

func (p *Video) SetRememberPosition(b bool) {
	C.elm_video_remember_position_set(p.obj, eBool(b))
}

func (p *Video) IsRememberPosition() bool {
	return C.elm_video_remember_position_get(p.obj) == eTrue
}

func (p *Video) Title() string {
	return C.GoString(C.elm_video_title_get(p.obj))
}

//------------------------------------------------------------

/*
win type must be,
ELM_WIN_SOCKET_IMAGE

Elm_Win *win = elm_win_add(NULL, "Image Socket Window",ELM_WIN_SOCKET_IMAGE);
elm_win_socket_listen(remote_win, "socket_name", 0, EINA_FALSE);

elm_plug_connect(plug, "socket_name", 0, EINA_FALSE);
Evas_Object *img = elm_plug_image_object_get(plug);
*/

type Plug struct {
	*objectBase
}

func NewPlug(parent Object) *Plug {
	eo := C.elm_plug_add(parent.eo())
	return wrapPlug(eo)
}

func wrapPlug(o *C.Eo) *Plug {
	if o != nil {
		return &Plug{wrapObjectBase(o)}
	}
	return nil
}

func (p *Plug) Connect(svcname string, svcnum int, svcsys bool) bool {
	csvcname := C.CString(svcname)
	defer free(csvcname)
	return C.elm_plug_connect(p.obj, csvcname, C.int(svcnum), eBool(svcsys)) == eTrue
}

func (p *Plug) ImageObject() Object {
	return wrapObjectBase(C.elm_plug_image_object_get(p.obj))
}

//------------------------------------------------------------

type Web struct {
	*widgetBase
}

func NewWeb(parent Object) *Web {
	o := C.elm_web_add(parent.eo())
	return wrapWeb(o)
}

func wrapWeb(o *C.Eo) *Web {
	if o != nil {
		return &Web{wrapWidgetBase(o)}
	}
	return nil
}

func (p *Web) SetHighlightMatchesText(b bool) {
	C.elm_web_text_matches_highlight_set(p.obj, eBool(b))
}

func (p *Web) IsHighlightMatchesText() bool {
	return C.elm_web_text_matches_highlight_get(p.obj) == eTrue
}

func (p *Web) SetUseragent(agent string) {
	cagent := C.CString(agent)
	defer free(cagent)
	C.elm_web_useragent_set(p.obj, cagent)
}

func (p *Web) Useragent() string {
	return C.GoString(C.elm_web_useragent_get(p.obj))
}

func (p *Web) SetURL(url string) bool {
	curl := C.CString(url)
	defer free(curl)
	return C.elm_web_url_set(p.obj, curl) == eTrue
}

func (p *Web) URL() string {
	return C.GoString(C.elm_web_url_get(p.obj))
}

func (p *Web) SetBgColor(r, g, b, a int) {
	C.elm_web_bg_color_set(p.obj, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *Web) BgColor() (int, int, int, int) {
	var r, g, b, a C.int
	C.elm_web_bg_color_get(p.obj, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *Web) SetInwinMode(b bool) {
	C.elm_web_inwin_mode_set(p.obj, eBool(b))
}

func (p *Web) IsInwinMode() bool {
	return C.elm_web_inwin_mode_get(p.obj) == eTrue
}

func (p *Web) SetTabPropagate(b bool) {
	C.elm_web_tab_propagate_set(p.obj, eBool(b))
}

func (p *Web) IsTabPropagate() bool {
	return C.elm_web_tab_propagate_get(p.obj) == eTrue
}

func (p *Web) SetHistoryEnabled(b bool) {
	C.elm_web_history_enabled_set(p.obj, eBool(b))
}

func (p *Web) IsHistoryEnabled() bool {
	return C.elm_web_history_enabled_get(p.obj) == eTrue
}

func (p *Web) SetZoomMode(mode WebZoomMode) {
	C.elm_web_zoom_mode_set(p.obj, C.Elm_Web_Zoom_Mode(mode))
}

func (p *Web) ZoomMode() WebZoomMode {
	return WebZoomMode(C.elm_web_zoom_mode_get(p.obj))
}

func (p *Web) SetZoom(zoom float64) {
	C.elm_web_zoom_set(p.obj, C.double(zoom))
}

func (p *Web) Zoom() float64 {
	return float64(C.elm_web_zoom_get(p.obj))
}

func (p *Web) SetPopupSelected(index int) {
	C.elm_web_popup_selected_set(p.obj, C.int(index))
}

func (p *Web) IsForwardPossible() bool {
	return C.elm_web_forward_possible_get(p.obj) == eTrue
}

func (p *Web) IsBackPossible() bool {
	return C.elm_web_back_possible_get(p.obj) == eTrue
}

func (p *Web) LoadProgress() float64 {
	return float64(C.elm_web_load_progress_get(p.obj))
}

func (p *Web) Selection() string {
	return C.GoString(C.elm_web_selection_get(p.obj))
}

func (p *Web) Title() string {
	return C.GoString(C.elm_web_title_get(p.obj))
}

func (p *Web) Navigate(step int) {
	C.elm_web_navigate(p.obj, C.int(step))
}

func (p *Web) Back() bool {
	return C.elm_web_back(p.obj) == eTrue
}

func (p *Web) LoadHtmlString(html string, base_url string, unreachable_url string) bool {
	chtml := C.CString(html)
	cbase_url := C.CString(base_url)
	cunreachable_url := C.CString(unreachable_url)
	defer free(chtml, cbase_url, cunreachable_url)
	return C.elm_web_html_string_load(p.obj, chtml, cbase_url, cunreachable_url) == eTrue
}

func (p *Web) SearchText(txt string, case_sensitive bool, forward bool, wrap bool) bool {
	ctxt := C.CString(txt)
	defer free(ctxt)
	return C.elm_web_text_search(p.obj, ctxt, eBool(case_sensitive), eBool(forward), eBool(wrap)) == eTrue
}

func (p *Web) DestroyPopup() {
	C.elm_web_popup_destroy(p.obj)
}

func (p *Web) ShowRegion(x, y, w, h int) {
	C.elm_web_region_show(p.obj, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (p *Web) Forward() bool {
	return C.elm_web_forward(p.obj) == eTrue
}

func (p *Web) MarkTextMatches(txt string, case_sensitive bool, highlight bool, limit uint) uint {
	ctxt := C.CString(txt)
	defer free(ctxt)
	return uint(C.elm_web_text_matches_mark(p.obj, ctxt, eBool(case_sensitive), eBool(highlight), C.uint(limit)))
}

func (p *Web) BringIn(x, y, w, h int) {
	C.elm_web_region_bring_in(p.obj, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (p *Web) Stop() bool {
	return C.elm_web_stop(p.obj) == eTrue
}

func (p *Web) IsNavigatePossible(steps int) bool {
	return C.elm_web_navigate_possible_get(p.obj, C.int(steps)) == eTrue
}

func (p *Web) FullReload() bool {
	return C.elm_web_reload_full(p.obj) == eTrue
}

func (p *Web) UnmarkAllTextMatches() bool {
	return C.elm_web_text_matches_unmark_all(p.obj) == eTrue
}

func (p *Web) Reload() bool {
	return C.elm_web_reload(p.obj) == eTrue
}

//export go_web_console_message_func
func go_web_console_message_func(cid C.int, obj *C.Eo, cmessage *C.char, cline_number C.uint, csource_id *C.char) {
	f := registry.Lookup(int(cid))
	if f != nil {
		// obj is omitted
		if fn, ok := f.(WebConsoleMessageFunc); ok {
			message := C.GoString(cmessage)
			source_id := C.GoString(csource_id)
			fn(message, uint(cline_number), source_id)
		}
	}
}

func (p *Web) SetWebConsoleMessageHook(fn WebConsoleMessageFunc) {
	id := registry.Register(fn)
	C.cgo_web_console_message_hook_set(p.obj, unsafe.Pointer(&id))
}

type WebWindowFeatures struct {
	obj *C.Elm_Web_Window_Features
}

func wrapWebWindowFeatures(o *C.Elm_Web_Window_Features) *WebWindowFeatures {
	if o != nil {
		return &WebWindowFeatures{o}
	}
	return nil
}

func (p *WebWindowFeatures) Property(flag WebWindowFeatureFlag) bool {
	return C.elm_web_window_features_property_get(p.obj, C.Elm_Web_Window_Feature_Flag(flag)) == eTrue
}

func (p *WebWindowFeatures) Region() (int, int, int, int) {
	var x, y, w, h C.Evas_Coord
	C.elm_web_window_features_region_get(p.obj, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *WebWindowFeatures) Ref() {
	C.elm_web_window_features_ref(p.obj)
}

func (p *WebWindowFeatures) Unref() {
	C.elm_web_window_features_unref(p.obj)
}

//export go_web_window_open_func
func go_web_window_open_func(cid C.int, obj *C.Eo, js C.Eina_Bool, wf *C.Elm_Web_Window_Features) *C.Eo {
	f := registry.Lookup(int(cid))
	if f != nil {
		// obj is omitted
		if fn, ok := f.(WebWindowOpenFunc); ok {
			o := fn(goBool(js), wrapWebWindowFeatures(wf))
			if o != nil && o.IsValid() {
				return o.eo()
			}
		}
	}
	return nil
}

func (p *Web) SetWindowCreateHook(fn WebWindowOpenFunc) {
	id := registry.Register(fn)
	C.cgo_web_window_create_hook_set(p.obj, unsafe.Pointer(&id))
}

//export go_web_dialog_confirm_func
func go_web_dialog_confirm_func(cid C.int, obj *C.Eo, cmessage *C.char, ret *C.Eina_Bool) *C.Eo {
	var cb C.Eina_Bool
	f := registry.Lookup(int(cid))
	if f != nil {
		// obj is omitted
		if fn, ok := f.(WebConfirmDialogFunc); ok {
			b, o := fn(C.GoString(cmessage))
			if b {
				cb = eTrue
			} else {
				cb = eFalse
			}
			ret = &cb
			if o != nil && o.IsValid() {
				return o.eo()
			}
		}
	}
	cb = eFalse
	ret = &cb
	return nil
}

func (p *Web) SetConfirmDialogHook(fn WebConfirmDialogFunc) {
	id := registry.Register(fn)
	C.cgo_web_dialog_confirm_hook_set(p.obj, unsafe.Pointer(&id))
}

//export go_web_dialog_prompt_func
func go_web_dialog_prompt_func(cid C.int, obj *C.Eo, cmessage *C.char, cdefval *C.char, value **C.char, ret *C.Eina_Bool) *C.Eo {
	var (
		cb  C.Eina_Bool
		val *C.char
	)
	f := registry.Lookup(int(cid))
	if f != nil {
		// obj is omitted
		if fn, ok := f.(WebPromptDialogFunc); ok {
			msg := C.GoString(cmessage)
			defval := C.GoString(cdefval)
			v, b, o := fn(msg, defval)
			if b {
				cb = eTrue
				if v != "" {
					val = C.CString(v)
					defer free(val)
				}
			} else {
				cb = eFalse
			}
			ret = &cb
			value = &val
			//TODO is return object necessary
			if o != nil && o.IsValid() {
				return o.eo()
			} else {
				return nil
			}
		}
	}
	cb = eFalse
	ret = &cb
	value = &val
	return nil
}

func (p *Web) SetPromptDialogHook(fn WebPromptDialogFunc) {
	id := registry.Register(fn)
	C.cgo_web_dialog_prompt_hook_set(p.obj, unsafe.Pointer(&id))
}

//export go_web_dialog_alert_func
func go_web_dialog_alert_func(cid C.int, obj *C.Eo, msg *C.char) *C.Eo {
	f := registry.Lookup(int(cid))
	if f != nil {
		// obj is omitted
		if fn, ok := f.(WebAlertDialogFunc); ok {
			o := fn(C.GoString(msg))
			if o != nil && o.IsValid() {
				return o.eo()
			}
		}
	}
	return nil
}

func (p *Web) SetAlertDialogHook(fn WebAlertDialogFunc) {
	id := registry.Register(fn)
	C.cgo_web_dialog_alert_hook_set(p.obj, unsafe.Pointer(&id))
}

/* TODO
//export go_web_dialog_file_selector_func
func go_web_dialog_file_selector_func(cid C.int, obj *C.Eo, always_multiple C.Eina_Bool, accept_types C.Eina_List, selected **C.Eina_List, ret *C.Eina_Bool) {
	f := registry.Lookup(int(cid))
	if f != nil {
		// obj is omitted
		if fn, ok := f.(WebDialogFileSelectorFunc); ok {
			....
		}
	}
	return nil
}
*/

/*
TODO:
void 	elm_web_dialog_file_selector_hook_set (Elm_Web *obj, Elm_Web_Dialog_File_Selector func, void *data)
Efl_Canvas_Object * 	elm_web_webkit_view_get (const Elm_Web *obj)
*/

//------------------------------------------------------------
