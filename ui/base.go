package ui

/*
#include "bridge.h"
*/
import "C"

import (
	"unsafe"
	//"fmt"
)

type Object interface {
	eo() *C.Eo
	//	Id() int
	IsValid() bool
	Dispose()
	Show()
	Hide()
	SetVisible(bool)
	IsVisible() bool
	SetData(string, interface{})
	Data(key string) interface{}
	DelData(key string) interface{}
	IsAntialiased() bool
	SetAntialiased(bool)
	Color() (int, int, int, int)
	SetColor(int, int, int, int)
	Name() string
	SetName(string)
	Move(int, int)
	Resize(int, int)
	Geometry() (int, int, int, int)
	Size() (int, int)
	SetSize(int, int)
	SetGeometry(int, int, int, int)
	Location() (int, int)
	SetLocation(int, int)
	Type() string
	Layer() int
	SetLayer(int)
	Raise()
	Lower()
	StackAbove(Object)
	StackBelow(Object)
	Clip() Object
	SetClip(Object)
	UnsetClip()
	Clipees() []Object
	Above() Object
	Below() Object
	SetRenderMode(RenderMode)
	RenderMode() RenderMode
	SetStaticClip(bool)
	IsStaticClip() bool
	Canvas() Canvas
	FindChildByName(string, int) Object
	SetHintFit()
	SetHintExpandBoth()
	SetHintExpandHorizontal()
	SetHintExpandVertical()
	SetHintFillBoth()
	SetHintFillHorizontal()
	SetHintFillVertical()
	HintMin() (int, int)
	SetHintMin(int, int)
	HintMax() (int, int)
	SetHintMax(int, int)
	HintRequest() (int, int)
	SetHintRequest(int, int)
	HintAspect() (AspectControl, int, int)
	SetHintAspect(AspectControl, int, int)
	HintAlign() (float64, float64)
	SetHintAlign(float64, float64)
	HintWeight() (float64, float64)
	SetHintWeight(float64, float64)
	HintPadding() (int, int, int, int)
	SetHintPadding(int, int, int, int)
	SetHintDisplayMode(DisplayMode)
	HintDisplayMode() DisplayMode
	SetParagraphDirection(BidiDirection)
	ParagraphDirection() BidiDirection
	UsePrecise(bool)
	IsPreciseUsing() bool
	//SetPointerMode(PointerMode)
	PointerMode() PointerMode
	SetFrameObject(bool)
	IsFrameObject() bool
	SetPassEvents(bool)
	IsPassEvents() bool
	SetRepeatEvents(bool)
	IsRepeatEvents() bool
	SetPropagateEvents(bool)
	IsPropagateEvents() bool
	SetFreezeEvents(bool)
	IsFreezeEvents() bool
}

type objectBase struct {
	obj *C.Eo
}

var _ Object = &objectBase{}

func wrapObjectBase(o *C.Eo) *objectBase {
	if o != nil {
		// TODO: use efl_key_data_get and set in 1.19
		/* TODO
		if C.evas_object_data_get(o, cObjectKey) == nil {
			id := objectID.Gen()
			C.evas_object_data_set(o, cObjectKey, unsafe.Pointer(&id))
		}*/
		//C.evas_object_show(o) default visibility causes unconsistency
		return &objectBase{obj: o}
	}
	return nil
}

//TODO remove when tests finished
func WrapTEST(o Object) Object {
	return wrapObjectBase(o.eo())
}

func (p *objectBase) eo() *C.Eo {
	return p.obj
}

/* TODO
func (p *objectBase) Id() int {
	id := C.evas_object_data_get(p.obj, cObjectKey)
	if id != nil {
		return int(*((*C.int)(id)))
	}
	return -1
}
*/

func (p *objectBase) IsValid() bool {
	return p.obj != nil
}

func (p *objectBase) Dispose() {
	C.evas_object_del(p.obj)
}

/*TODO: make children visible how?*/
func (p *objectBase) Show() {
	C.evas_object_show(p.obj)
}

func (p *objectBase) Hide() {
	C.evas_object_hide(p.obj)
}

func (p *objectBase) SetVisible(b bool) {
	if b {
		p.Show()
	} else {
		p.Hide()
	}
}

func (p *objectBase) IsVisible() bool {
	return C.evas_object_visible_get(p.obj) == eTrue
}

func (p *objectBase) SetData(key string, data interface{}) {
	objectDataMap.SetData(p.obj, key, data)
}

func (p *objectBase) Data(key string) interface{} {
	return objectDataMap.Data(p.obj, key)
}

func (p *objectBase) DelData(key string) interface{} {
	return objectDataMap.DelData(p.obj, key)
}

func (p *objectBase) IsAntialiased() bool {
	return C.evas_object_anti_alias_get(p.obj) == eTrue
}

func (p *objectBase) SetAntialiased(b bool) {
	C.evas_object_anti_alias_set(p.obj, eBool(b))
}

func (p *objectBase) Color() (int, int, int, int) {
	var r, g, b, a C.int
	C.evas_object_color_get(p.obj, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *objectBase) SetColor(r, g, b, a int) {
	C.evas_object_color_set(p.obj, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *objectBase) Name() string {
	return C.GoString(C.evas_object_name_get(p.obj))
}

func (p *objectBase) SetName(name string) {
	cname := C.CString(name)
	defer free(cname)
	C.evas_object_name_set(p.obj, cname)
}

func (p *objectBase) Move(x, y int) {
	C.evas_object_move(p.obj, C.Evas_Coord(x), C.Evas_Coord(y))
}

func (p *objectBase) Resize(w, h int) {
	C.evas_object_resize(p.obj, C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *objectBase) Geometry() (int, int, int, int) {
	var x, y, w, h C.Evas_Coord
	C.evas_object_geometry_get(p.obj, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *objectBase) Size() (int, int) {
	_, _, w, h := p.Geometry()
	return w, h
}

func (p *objectBase) SetSize(w, h int) {
	p.Resize(w, h)
}

func (p *objectBase) SetGeometry(x, y, w, h int) {
	C.evas_object_geometry_set(p.obj, C.Evas_Coord(x), C.Evas_Coord(y), C.Evas_Coord(w), C.Evas_Coord(h))
}

/*
func (p *objectBase) Position() (int, int) {
	x, y, _, _ := p.Geometry()
	return x, y
}

func (p *objectBase) SetPosition(x, y int) {
	p.Move(x, y)
}
*/

func (p *objectBase) Location() (int, int) {
	x, y, _, _ := p.Geometry()
	return x, y
}

func (p *objectBase) SetLocation(x, y int) {
	p.Move(x, y)
}

func (p *objectBase) Type() string {
	return C.GoString(C.evas_object_type_get(p.obj))
}

func (p *objectBase) Layer() int {
	return int(C.evas_object_layer_get(p.obj))
}

func (p *objectBase) SetLayer(l int) {
	C.evas_object_layer_set(p.obj, C.short(l))
}

func (p *objectBase) Raise() {
	C.evas_object_raise(p.obj)
}

func (p *objectBase) Lower() {
	C.evas_object_lower(p.obj)
}

func (p *objectBase) StackAbove(above Object) {
	C.evas_object_stack_above(p.obj, above.eo())
}

func (p *objectBase) StackBelow(below Object) {
	C.evas_object_stack_below(p.obj, below.eo())
}

func (p *objectBase) Clip() Object {
	return wrapObjectBase(C.evas_object_clip_get(p.obj))
}

func (p *objectBase) SetClip(o Object) {
	C.evas_object_clip_set(p.obj, o.eo())
}

func (p *objectBase) UnsetClip() {
	C.evas_object_clip_unset(p.obj)
}

func (p *objectBase) Clipees() []Object {
	lst := C.evas_object_clipees_get(p.obj)
	if lst != nil {
		return newListIterator(lst).ObjectSlice()
	}
	return nil
}

func (p *objectBase) Above() Object {
	return wrapObjectBase(C.evas_object_above_get(p.obj))
}

func (p *objectBase) Below() Object {
	return wrapObjectBase(C.evas_object_below_get(p.obj))
}

func (p *objectBase) SetRenderMode(mode RenderMode) {
	C.evas_object_render_op_set(p.obj, C.Evas_Render_Op(mode))
}

func (p *objectBase) RenderMode() RenderMode {
	return RenderMode(C.evas_object_render_op_get(p.obj))
}

func (p *objectBase) SetStaticClip(b bool) {
	C.evas_object_static_clip_set(p.obj, eBool(b))
}

func (p *objectBase) IsStaticClip() bool {
	return C.evas_object_static_clip_get(p.obj) == eTrue
}

func (p *objectBase) Canvas() Canvas {
	return wrapEvas(C.evas_object_evas_get(p.obj))
}

func (p *objectBase) FindChildByName(name string, recursive int) Object {
	cname := C.CString(name)
	defer free(name)
	return wrapObjectBase(C.evas_object_name_child_find(p.obj, cname, C.int(recursive)))
}

func (p *objectBase) SetHintFit() {
	p.SetHintAlign(HintFill, HintFill)
	p.SetHintWeight(HintExpand, HintExpand)
}

func (p *objectBase) SetHintExpandBoth() {
	p.SetHintWeight(HintExpand, HintExpand)
}

func (p *objectBase) SetHintExpandHorizontal() {
	p.SetHintWeight(HintExpand, 0.0)
}

func (p *objectBase) SetHintExpandVertical() {
	p.SetHintWeight(0.0, HintExpand)
}

func (p *objectBase) SetHintFillBoth() {
	p.SetHintAlign(HintFill, HintFill)
}

func (p *objectBase) SetHintFillHorizontal() {
	p.SetHintAlign(HintFill, 0.5)
}

func (p *objectBase) SetHintFillVertical() {
	p.SetHintAlign(0.5, HintFill)
}

func (p *objectBase) HintMin() (int, int) {
	var w, h C.Evas_Coord
	C.evas_object_size_hint_min_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *objectBase) SetHintMin(width, height int) {
	C.evas_object_size_hint_min_set(p.obj, C.Evas_Coord(width), C.Evas_Coord(height))
}

func (p *objectBase) HintMax() (int, int) {
	var w, h C.Evas_Coord
	C.evas_object_size_hint_max_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *objectBase) SetHintMax(w, h int) {
	C.evas_object_size_hint_max_set(p.obj, C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *objectBase) HintRequest() (int, int) {
	var w, h C.Evas_Coord
	C.evas_object_size_hint_request_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *objectBase) SetHintRequest(width, height int) {
	C.evas_object_size_hint_request_set(p.obj, C.Evas_Coord(width), C.Evas_Coord(height))
}

func (o *objectBase) HintAspect() (AspectControl, int, int) {
	var (
		a    C.Evas_Aspect_Control
		w, h C.Evas_Coord
	)
	C.evas_object_size_hint_aspect_get(o.obj, &a, &w, &h)
	return AspectControl(a), int(w), int(h)
}

func (p *objectBase) SetHintAspect(a AspectControl, w, h int) {
	C.evas_object_size_hint_aspect_set(p.obj, C.Evas_Aspect_Control(a), C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *objectBase) HintAlign() (float64, float64) {
	var x, y C.double
	C.evas_object_size_hint_align_get(p.obj, &x, &y)
	return float64(x), float64(y)
}

func (p *objectBase) SetHintAlign(x, y float64) {
	C.evas_object_size_hint_align_set(p.obj, C.double(x), C.double(y))
}

func (p *objectBase) HintWeight() (float64, float64) {
	var x, y C.double
	C.evas_object_size_hint_weight_get(p.obj, &x, &y)
	return float64(x), float64(y)
}

func (p *objectBase) SetHintWeight(x, y float64) {
	C.evas_object_size_hint_weight_set(p.obj, C.double(x), C.double(y))
}

// left right top bottom
func (p *objectBase) HintPadding() (int, int, int, int) {
	var l, r, t, b C.Evas_Coord
	C.evas_object_size_hint_padding_get(p.obj, &l, &r, &t, &b)
	return int(l), int(r), int(t), int(b)
}

func (p *objectBase) SetHintPadding(l, r, t, b int) {
	C.evas_object_size_hint_padding_set(p.obj, C.Evas_Coord(l), C.Evas_Coord(r), C.Evas_Coord(t), C.Evas_Coord(b))
}

func (p *objectBase) SetHintDisplayMode(mode DisplayMode) {
	C.evas_object_size_hint_display_mode_set(p.obj, C.Evas_Display_Mode(mode))
}

func (p *objectBase) HintDisplayMode() DisplayMode {
	return DisplayMode(C.evas_object_size_hint_display_mode_get(p.obj))
}

func (p *objectBase) KeyGrab(key string, modifiers uint64, not_modifiers uint64, exclusive bool) bool {
	ckey := C.CString(key)
	defer free(ckey)
	return C.evas_object_key_grab(p.obj, ckey, C.Evas_Modifier_Mask(modifiers), C.Evas_Modifier_Mask(not_modifiers), eBool(exclusive)) == eTrue
}

func (p *objectBase) KeyUngrab(key string, modifiers uint64, not_modifiers uint64) {
	ckey := C.CString(key)
	defer free(ckey)
	C.evas_object_key_ungrab(p.obj, ckey, C.Evas_Modifier_Mask(modifiers), C.Evas_Modifier_Mask(not_modifiers))
}

func (p *objectBase) SetParagraphDirection(dir BidiDirection) {
	C.evas_object_paragraph_direction_set(p.obj, C.Efl_Text_Bidirectional_Type(dir))
}

func (p *objectBase) ParagraphDirection() BidiDirection {
	return BidiDirection(C.evas_object_paragraph_direction_get(p.obj))
}

func (p *objectBase) UsePrecise(b bool) {
	C.evas_object_precise_is_inside_set(p.obj, eBool(b))
}

func (p *objectBase) IsPreciseUsing() bool {
	return C.evas_object_precise_is_inside_get(p.obj) == eTrue
}

/* TODO: in 1.19 Efl Input... used
https://abi-laboratory.pro/tracker/compat_report/efl/1.18.4/1.19.0-beta2/6c22a/abi_compat_report.html
func (p *objectBase) SetPointerMode(mode PointerMode) {
	//C.evas_object_pointer_mode_set(p.obj, C.Efl_Event_Object_Pointer_Mode(mode))
	C.evas_object_pointer_mode_set(p.obj, C.Efl_Input_Object_Pointer_Mode(mode))
}
*/

func (p *objectBase) PointerMode() PointerMode {
	return PointerMode(C.evas_object_pointer_mode_get(p.obj))
}

func (p *objectBase) SetFrameObject(b bool) {
	C.evas_object_is_frame_object_set(p.obj, eBool(b))
}

func (p *objectBase) IsFrameObject() bool {
	return C.evas_object_is_frame_object_get(p.obj) == eTrue
}

func (p *objectBase) SetPassEvents(b bool) {
	C.evas_object_pass_events_set(p.obj, eBool(b))
}

func (p *objectBase) IsPassEvents() bool {
	return C.evas_object_pass_events_get(p.obj) == eTrue
}

func (p *objectBase) SetRepeatEvents(b bool) {
	C.evas_object_repeat_events_set(p.obj, eBool(b))
}

func (p *objectBase) IsRepeatEvents() bool {
	return C.evas_object_repeat_events_get(p.obj) == eTrue
}

func (p *objectBase) SetPropagateEvents(b bool) {
	C.evas_object_propagate_events_set(p.obj, eBool(b))
}

func (p *objectBase) IsPropagateEvents() bool {
	return C.evas_object_propagate_events_get(p.obj) == eTrue
}

func (p *objectBase) SetFreezeEvents(b bool) {
	C.evas_object_freeze_events_set(p.obj, eBool(b))
}

func (p *objectBase) IsFreezeEvents() bool {
	return C.evas_object_freeze_events_get(p.obj) == eTrue
}

/*
	Widget
*/

type Widget interface {
	Object
	HasFocus() bool
	SetFocus()
	SetAllowFocus(bool)
	IsAllowFocus() bool
	SetFocusNext(FocusDirection)
	NextFocusObject(FocusDirection) Widget
	SetNextFocusObject(Widget, FocusDirection)
	SetFocusMovePolicy(FocusMovePolicyType)
	FocusMovePolicy() FocusMovePolicyType
	SetFocusHighlightStyle(string) bool
	FocusHighlightStyle() string
	FocusedObject() Widget
	IsDisabled() bool
	SetDisabled(bool)
	Scale() float64
	SetScale(float64)
	SetStyle(string)
	Style() string
	SetPartText(string, string)
	PartText(string) string
	Text() string
	SetText(string)
	IsMirrored() bool
	SetMirrored(bool)
	IsAutomaticMirrored() bool
	SetAutomaticMirrored(bool)
	FindObjectByName(string, int) Widget
	ParentWidget() Widget
	TopWidget() Widget
	WidgetType() string
	IsWidget() bool
	SetCursor(string)
	Cursor() string
	UnsetCursor()
	SetCursorStyle(string)
	CursorStyle() string
	SetAccessInfo(string)
	EmitSignal(string, string)
	SetOrientationModeDisabled(bool)
	IsOrientationModeDisabled() bool
	ShowTooltip()
	HideTooltip()
	SetTooltipText(string)
	UnsetTooltip()
	SetTooltipStyle(string)
	TooltipStyle() string
	SetTooltipWindowMode(bool) bool
	IsTooltipWindowMode() bool
	IncrementTooltipMovementFreezeCount()
	DecrementTooltipMovementFreezeCount()
	TooltipMovementFreezeCount() int
	SetTooltipOrient(TooltipOrient)
	TooltipOrient() TooltipOrient
	PushScrollHold()
	PopScrollHold()
	ScrollHold() int
	PushScrollFreeze()
	PopScrollFreeze()
	ScrollFreeze() int
	SetLockX(bool)
	SetLockY(bool)
	IsXLocked() bool
	IsYLocked() bool
	SetScrollItemLoopEnabled(bool)
	IsScrollItemLoopEnabled() bool
	SetTheme(*Theme)
	Theme() *Theme
	AddHandler(string, Handler)
	AddHandlerWithData(string, DataHandler, interface{})
	AddEventHandler(string, EventHandler)
	AddEventHandlerWithData(string, EventDataHandler, interface{})
	On(string, func())
	OnWithData(string, func(interface{}), interface{})
	OnEvent(string, func(interface{}))
	OnEventWithData(string, func(interface{}, interface{}), interface{})
	OnShow(func())
	OnHide(func())
	OnMove(func())
	OnResize(func())
	OnDispose(func())
	OnSizeHintsChanged(func())
	OnImagePreloaded(func())
	OnImageUnloaded(func())
	OnMouseIn(func(*MouseInEvent))
	OnMouseOut(func(*MouseOutEvent))
	OnMouseDown(func(*MouseDownEvent))
	OnMouseUp(func(*MouseUpEvent))
	OnMouseMove(func(*MouseMoveEvent))
	OnMouseWheel(func(*MouseWheelEvent))
	OnMultiDown(func(*MultiDownEvent))
	OnMultiUp(func(*MultiUpEvent))
	OnMultiMove(func(*MultiMoveEvent))
	OnKeyDown(func(*KeyDownEvent))
	OnKeyUp(func(*KeyUpEvent))
	OnHold(func(*HoldEvent))
	OnFocusIn(func(*MouseInEvent))
	OnFocusOut(func(*MouseInEvent))
	OnAxisUpdate(func(*AxisUpdateEvent))
	SetCnpSelection(SelectionType, SelectionFormat, string) bool
	CnpSelection(SelectionType, SelectionFormat, SelectionHandler) bool
	ClearCnpSelection(SelectionType) bool
	SetCnpSelectionLossHandler(SelectionType, SelectionLossHandler)
	SetDragAction(XDndAction) bool
	CancelDrag() bool
}

type widgetBase struct {
	*objectBase
}

var _ Widget = &widgetBase{}
var _ Object = &widgetBase{}

func wrapWidgetBase(o *C.Eo) *widgetBase {
	if o != nil {
		return &widgetBase{wrapObjectBase(o)}
	}
	return nil
}

func (p *widgetBase) HasFocus() bool {
	return C.elm_object_focus_get(p.obj) == eTrue
}

func (p *widgetBase) SetFocus() {
	C.elm_object_focus_set(p.obj, eTrue)
}

func (p *widgetBase) SetAllowFocus(b bool) {
	C.elm_object_focus_allow_set(p.obj, eBool(b))
}

func (p *widgetBase) IsAllowFocus() bool {
	return C.elm_object_focus_allow_get(p.obj) == eTrue
}

func (p *widgetBase) SetFocusNext(dir FocusDirection) {
	C.elm_object_focus_next(p.obj, C.Elm_Focus_Direction(dir))
}

func (p *widgetBase) NextFocusObject(dir FocusDirection) Widget {
	return wrapWidgetBase(C.elm_object_focus_next_object_get(p.obj, C.Elm_Focus_Direction(dir)))
}

func (p *widgetBase) SetNextFocusObject(o Widget, dir FocusDirection) {
	C.elm_object_focus_next_object_set(p.obj, o.eo(), C.Elm_Focus_Direction(dir))
}

func (p *widgetBase) SetFocusMovePolicy(policy FocusMovePolicyType) {
	C.elm_object_focus_move_policy_set(p.obj, C.Elm_Focus_Move_Policy(policy))
}

func (p *widgetBase) FocusMovePolicy() FocusMovePolicyType {
	return FocusMovePolicyType(C.elm_object_focus_move_policy_get(p.obj))
}

func (p *widgetBase) SetFocusHighlightStyle(style string) bool {
	cstyle := C.CString(style)
	defer free(cstyle)
	return C.elm_object_focus_highlight_style_set(p.obj, cstyle) == eTrue
}

func (p *widgetBase) FocusHighlightStyle() string {
	return C.GoString(C.elm_object_focus_highlight_style_get(p.obj))
}

func (p *widgetBase) FocusedObject() Widget {
	return wrapWidgetBase(C.elm_object_focused_object_get(p.obj))
}

func (p *widgetBase) IsDisabled() bool {
	return goBool(C.elm_object_disabled_get(p.obj))
}

func (p *widgetBase) SetDisabled(b bool) {
	C.elm_object_disabled_set(p.obj, eBool(b))
}

func (p *widgetBase) Scale() float64 {
	return float64(C.elm_object_scale_get(p.obj))
}

func (p *widgetBase) SetScale(scale float64) {
	C.elm_object_scale_set(p.obj, C.double(scale))
}

func (p *widgetBase) SetStyle(style string) {
	cstyle := C.CString(style)
	defer free(cstyle)
	C.elm_object_style_set(p.obj, cstyle)
}

func (p *widgetBase) Style() string {
	return C.GoString(C.elm_object_style_get(p.obj))
}

func (p *widgetBase) SetPartText(part string, text string) {
	var cpart *C.char
	if part != "" {
		cpart = C.CString(part)
		defer free(cpart)
	}
	ctext := C.CString(text)
	defer free(ctext)
	C.elm_object_part_text_set(p.obj, cpart, ctext)
}

func (p *widgetBase) PartText(part string) string {
	var cpart *C.char
	if part != "" {
		cpart = C.CString(part)
		defer free(cpart)
	}
	return C.GoString(C.elm_object_part_text_get(p.obj, cpart))
}

func (p *widgetBase) Text() string {
	return p.PartText("")
}

func (p *widgetBase) SetText(txt string) {
	p.SetPartText("", txt)
}

func (p *widgetBase) IsMirrored() bool {
	return C.elm_object_mirrored_get(p.obj) == eTrue
}

func (p *widgetBase) SetMirrored(b bool) {
	C.elm_object_mirrored_set(p.obj, eBool(b))
}

func (p *widgetBase) IsAutomaticMirrored() bool {
	return C.elm_object_mirrored_automatic_get(p.obj) == eTrue
}

func (p *widgetBase) SetAutomaticMirrored(b bool) {
	C.elm_object_mirrored_automatic_set(p.obj, eBool(b))
}

func (p *widgetBase) FindObjectByName(name string, recursive int) Widget {
	cname := C.CString(name)
	defer free(name)
	return wrapWidgetBase(C.elm_object_name_find(p.obj, cname, C.int(recursive)))
}

func (p *widgetBase) ParentWidget() Widget {
	return wrapWidgetBase(C.elm_object_parent_widget_get(p.obj))
}

func (p *widgetBase) TopWidget() Widget {
	return wrapWidgetBase(C.elm_object_top_widget_get(p.obj))
}

func (p *widgetBase) WidgetType() string {
	return C.GoString(C.elm_object_widget_type_get(p.obj))
}

func (p *widgetBase) IsWidget() bool {
	return C.elm_object_widget_check(p.obj) == eTrue
}

func (p *widgetBase) SetCursor(cursor string) {
	ccursor := C.CString(cursor)
	defer free(ccursor)
	C.elm_object_cursor_set(p.obj, ccursor)
}

func (p *widgetBase) Cursor() string {
	return C.GoString(C.elm_object_cursor_get(p.obj))
}

func (p *widgetBase) UnsetCursor() {
	C.elm_object_cursor_unset(p.obj)
}

func (p *widgetBase) SetCursorStyle(style string) {
	cstyle := C.CString(style)
	defer free(cstyle)
	C.elm_object_cursor_style_set(p.obj, cstyle)
}

func (p *widgetBase) CursorStyle() string {
	return C.GoString(C.elm_object_cursor_style_get(p.obj))
}

func (p *widgetBase) SetAccessInfo(txt string) {
	ctxt := C.CString(txt)
	defer free(ctxt)
	C.elm_object_access_info_set(p.obj, ctxt)
}

func (p *widgetBase) EmitSignal(emission, source string) {
	cemi := C.CString(emission)
	csrc := C.CString(source)
	defer free(cemi, csrc)
	C.elm_object_signal_emit(p.obj, cemi, csrc)
}

func (p *widgetBase) SetOrientationModeDisabled(b bool) {
	C.elm_object_orientation_mode_disabled_set(p.obj, eBool(b))
}

func (p *widgetBase) IsOrientationModeDisabled() bool {
	return C.elm_object_orientation_mode_disabled_get(p.obj) == eTrue
}

func (p *widgetBase) ShowTooltip() {
	C.elm_object_tooltip_show(p.obj)
}

func (p *widgetBase) HideTooltip() {
	C.elm_object_tooltip_hide(p.obj)
}

func (p *widgetBase) SetTooltipText(txt string) {
	ctxt := C.CString(txt)
	defer free(ctxt)
	C.elm_object_tooltip_text_set(p.obj, ctxt)
}

func (p *widgetBase) UnsetTooltip() {
	C.elm_object_tooltip_unset(p.obj)
}

func (p *widgetBase) SetTooltipStyle(style string) {
	cstyle := C.CString(style)
	defer free(cstyle)
	C.elm_object_tooltip_style_set(p.obj, cstyle)
}

func (p *widgetBase) TooltipStyle() string {
	return C.GoString(C.elm_object_tooltip_style_get(p.obj))
}

func (p *widgetBase) SetTooltipWindowMode(disable bool) bool {
	return C.elm_object_tooltip_window_mode_set(p.obj, eBool(disable)) == eTrue
}

func (p *widgetBase) IsTooltipWindowMode() bool {
	return C.elm_object_tooltip_window_mode_get(p.obj) == eTrue
}

func (p *widgetBase) IncrementTooltipMovementFreezeCount() {
	C.elm_object_tooltip_move_freeze_push(p.obj)
}

func (p *widgetBase) DecrementTooltipMovementFreezeCount() {
	C.elm_object_tooltip_move_freeze_pop(p.obj)
}

func (p *widgetBase) TooltipMovementFreezeCount() int {
	return int(C.elm_object_tooltip_move_freeze_get(p.obj))
}

func (p *widgetBase) SetTooltipOrient(orient TooltipOrient) {
	C.elm_object_tooltip_orient_set(p.obj, C.Elm_Tooltip_Orient(orient))
}

func (p *widgetBase) TooltipOrient() TooltipOrient {
	return TooltipOrient(C.elm_object_tooltip_orient_get(p.obj))
}

func (p *widgetBase) PushScrollHold() {
	C.elm_object_scroll_hold_push(p.obj)
}

func (p *widgetBase) PopScrollHold() {
	C.elm_object_scroll_hold_pop(p.obj)
}

func (p *widgetBase) ScrollHold() int {
	return int(C.elm_object_scroll_hold_get(p.obj))
}

func (p *widgetBase) PushScrollFreeze() {
	C.elm_object_scroll_freeze_push(p.obj)
}

func (p *widgetBase) PopScrollFreeze() {
	C.elm_object_scroll_freeze_pop(p.obj)
}

func (p *widgetBase) ScrollFreeze() int {
	return int(C.elm_object_scroll_freeze_get(p.obj))
}

func (p *widgetBase) SetLockX(lock bool) {
	C.elm_object_scroll_lock_x_set(p.obj, eBool(lock))
}

func (p *widgetBase) SetLockY(lock bool) {
	C.elm_object_scroll_lock_y_set(p.obj, eBool(lock))
}

func (p *widgetBase) IsXLocked() bool {
	return goBool(C.elm_object_scroll_lock_x_get(p.obj))
}

func (p *widgetBase) IsYLocked() bool {
	return goBool(C.elm_object_scroll_lock_y_get(p.obj))
}

func (p *widgetBase) SetScrollItemLoopEnabled(b bool) {
	C.elm_object_scroll_item_loop_enabled_set(p.obj, eBool(b))
}

func (p *widgetBase) IsScrollItemLoopEnabled() bool {
	return goBool(C.elm_object_scroll_item_loop_enabled_get(p.obj))
}

func (p *widgetBase) SetTheme(theme *Theme) {
	C.elm_object_theme_set(p.obj, theme.obj)
}

func (p *widgetBase) Theme() *Theme {
	return wrapTheme(C.elm_object_theme_get(p.obj))
}

func (p *widgetBase) AddHandler(ev string, h Handler) {
	addSmartCallback(p.obj, ev, h)
}

func (p *widgetBase) AddHandlerWithData(ev string, h DataHandler, data interface{}) {
	addSmartCallbackWithData(p.obj, ev, h, data)
}

func (p *widgetBase) AddEventHandler(ev string, h EventHandler) {
	addSmartCallbackWithEvent(p.Type(), p.obj, ev, h)
}

func (p *widgetBase) AddEventHandlerWithData(ev string, h EventDataHandler, data interface{}) {
	addSmartCallbackWithEventAndData(p.Type(), p.obj, ev, h, data)
}

/*
func (p *widgetBase) DelHandler(ev string, h Handler) {
	//TODO
}
*/

func (p *widgetBase) On(ev string, fn func()) {
	eh := HandleFunc(fn)
	p.AddHandler(ev, eh)
}

func (p *widgetBase) OnWithData(ev string, fn func(interface{}), data interface{}) {
	eh := DataHandleFunc(fn)
	p.AddHandlerWithData(ev, eh, data)
}

func (p *widgetBase) OnEvent(ev string, fn func(interface{})) {
	eh := EventHandleFunc(fn)
	p.AddEventHandler(ev, eh)
}

func (p *widgetBase) OnEventWithData(ev string, fn func(interface{}, interface{}), data interface{}) {
	eh := EventDataHandleFunc(fn)
	p.AddEventHandlerWithData(ev, eh, data)
}

func (p *widgetBase) OnShow(fn func()) {
	eh := HandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeShow, eh)
}

func (p *widgetBase) OnHide(fn func()) {
	eh := HandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeHide, eh)
}

func (p *widgetBase) OnMove(fn func()) {
	eh := HandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeMove, eh)
}

func (p *widgetBase) OnResize(fn func()) {
	eh := HandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeResize, eh)
}

func (p *widgetBase) OnDispose(fn func()) {
	eh := HandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeDel, eh)
}

func (p *widgetBase) OnSizeHintsChanged(fn func()) {
	eh := HandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeSizeHintsChanged, eh)
}

func (p *widgetBase) OnImagePreloaded(fn func()) {
	eh := HandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeImagePreloaded, eh)
}

func (p *widgetBase) OnImageUnloaded(fn func()) {
	eh := HandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeImageUnloaded, eh)
}

func (p *widgetBase) OnMouseIn(fn func(ev *MouseInEvent)) {
	eh := MouseInHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeMouseIn, eh)
}

func (p *widgetBase) OnMouseOut(fn func(ev *MouseOutEvent)) {
	eh := MouseOutHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeMouseOut, eh)
}

func (p *widgetBase) OnMouseDown(fn func(ev *MouseDownEvent)) {
	eh := MouseDownHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeMouseDown, eh)
}

func (p *widgetBase) OnMouseUp(fn func(ev *MouseUpEvent)) {
	eh := MouseUpHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeMouseUp, eh)
}

func (p *widgetBase) OnMouseMove(fn func(ev *MouseMoveEvent)) {
	eh := MouseMoveHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeMouseMove, eh)
}

func (p *widgetBase) OnMouseWheel(fn func(ev *MouseWheelEvent)) {
	eh := MouseWheelHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeMouseWheel, eh)
}

func (p *widgetBase) OnMultiDown(fn func(ev *MultiDownEvent)) {
	eh := MultiDownHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeMultiDown, eh)
}

func (p *widgetBase) OnMultiUp(fn func(ev *MultiUpEvent)) {
	eh := MultiUpHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeMultiUp, eh)
}

func (p *widgetBase) OnMultiMove(fn func(ev *MultiMoveEvent)) {
	eh := MultiMoveHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeMultiMove, eh)
}

func (p *widgetBase) OnKeyDown(fn func(ev *KeyDownEvent)) {
	eh := KeyDownHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeKeyDown, eh)
}

func (p *widgetBase) OnKeyUp(fn func(ev *KeyUpEvent)) {
	eh := KeyUpHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeKeyUp, eh)
}

func (p *widgetBase) OnHold(fn func(ev *HoldEvent)) {
	eh := HoldHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeHold, eh)
}

func (p *widgetBase) OnFocusIn(fn func(ev *MouseInEvent)) {
	eh := MouseInHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeFocusIn, eh)
}

func (p *widgetBase) OnFocusOut(fn func(ev *MouseInEvent)) {
	eh := MouseInHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeFocusOut, eh)
}

func (p *widgetBase) OnAxisUpdate(fn func(ev *AxisUpdateEvent)) {
	eh := AxisUpdateHandleFunc(fn)
	addEvasObjectCallback(p.obj, CallbackTypeAxisUpdate, eh)
}

func (p *widgetBase) SetCnpSelection(st SelectionType, sf SelectionFormat, txt string) bool {
	ctxt := C.CString(txt)
	defer free(ctxt)
	return C.cgo_elm_cnp_selection_set(p.obj, C.Elm_Sel_Type(st), C.Elm_Sel_Format(sf), ctxt) == eTrue
}

func (p *widgetBase) CnpSelection(st SelectionType, sf SelectionFormat, h SelectionHandler) bool {
	var data unsafe.Pointer
	if h != nil {
		id := registry.Register(h)
		data = unsafe.Pointer(&id)
	}
	return C.cgo_elm_cnp_selection_get(p.obj, C.Elm_Sel_Type(st), C.Elm_Sel_Format(sf), data) == eTrue
}

func (p *widgetBase) ClearCnpSelection(st SelectionType) bool {
	return C.elm_object_cnp_selection_clear(p.obj, C.Elm_Sel_Type(st)) == eTrue
}

func (p *widgetBase) SetCnpSelectionLossHandler(st SelectionType, h SelectionLossHandler) {
	id := registry.Register(h)
	C.cgo_elm_selection_loss_callback_set(p.obj, C.Elm_Sel_Type(st), unsafe.Pointer(&id))
}

func (p *widgetBase) SetDragAction(action XDndAction) bool {
	return C.elm_drag_action_set(p.obj, C.Elm_Xdnd_Action(action)) == eTrue
}

func (p *widgetBase) CancelDrag() bool {
	return C.elm_drag_cancel(p.obj) == eTrue
}

/* in 1.19
func (p *widgetBase) Children() []Object {
	it := C.efl_children_iterator_new(p.obj)
	if it != nil {
		return NewIterator(it).ObjectSlice()
	}
	return nil
}
*/

type Container interface {
	SetPartContent(string, Object)
	PartContent(string) Object
	UnsetPartContent(string) Object
	SetContent(Object)
	Content() Object
	UnsetContent() Object
	/*
		SetDragSource(float64, float64, bool, bool)
		AddDragItem()
		DelDragItem()
		AddDropItem()
		DelDropItem()
	*/
}

var _ Container = &containerBase{}
var _ Widget = &containerBase{}
var _ Object = &containerBase{}

type containerBase struct {
	*widgetBase
}

func wrapContainerBase(o *C.Eo) *containerBase {
	if o != nil {
		return &containerBase{wrapWidgetBase(o)}
	}
	return nil
}

func (p *containerBase) SetPartContent(part string, content Object) {
	var cpart *C.char
	if part != "" {
		cpart = C.CString(part)
		defer free(cpart)
	}
	C.elm_object_part_content_set(p.obj, cpart, content.eo())
}

func (p *containerBase) PartContent(part string) Object {
	var cpart *C.char
	if part != "" {
		cpart = C.CString(part)
		defer free(cpart)
	}
	return wrapObjectBase(C.elm_object_part_content_get(p.obj, cpart))
}

func (p *containerBase) UnsetPartContent(part string) Object {
	var cpart *C.char
	if part != "" {
		cpart = C.CString(part)
		defer free(cpart)
	}
	return wrapObjectBase(C.elm_object_part_content_unset(p.obj, cpart))
}

func (p *containerBase) SetContent(o Object) {
	p.SetPartContent("", o)
}

func (p *containerBase) Content() Object {
	return p.PartContent("")
}

func (p *containerBase) UnsetContent() Object {
	return p.UnsetPartContent("")
}

func (p *containerBase) SetStandardIcon(name string) {
	icon := NewIcon(p)
	icon.SetStandard(name)
	p.SetPartContent("icon", icon)
}

/*
Eina_Bool 	elm_config_cursor_engine_only_get (void)
void 	elm_config_cursor_engine_only_set (Eina_Bool engine_only)
void 	elm_object_cursor_theme_search_enabled_set (Evas_Object *obj, Eina_Bool theme_search)
Eina_Bool 	elm_object_cursor_theme_search_enabled_get (const Evas_Object *obj)
*/

/*
void                    elm_object_tooltip_domain_translatable_text_set(Evas_Object *obj, const char *domain, const char *text)
void                    elm_object_tooltip_translatable_text_set(Evas_Object *obj, const char *text)
void                    elm_object_tooltip_content_cb_set(Evas_Object *obj, Elm_Tooltip_Content_Cb func, void *data, Evas_Smart_Cb del_cb)
*/

type WidgetItem interface {
	eo() *C.Elm_Widget_Item
	Widget() Widget
	SetPartContent(string, Object)
	PartContent(string) Object
	UnsetPartContent(string) Object
	SetContent(Object)
	Content() Object
	UnsetContent() Object
	SetPartText(string, string)
	PartText(string) string
	SetText(string)
	Text() string
	SetAccessInfo(string)
	AccessInfo() string
	EmitSignal(string, string)
	SetDisabled(bool)
	IsDisabled() bool
	Dispose()
	SetFocus(bool)
	HasFocus() bool
	NextFocusObject(FocusDirection) Widget
	SetNextFocusObject(Widget, FocusDirection)
	NextFocusItem(FocusDirection) WidgetItem
	SetNextFocusItem(WidgetItem, FocusDirection)
	SetTooltipText(string)
	SetTooltipWindowMode(bool)
	IsTooltipWindowMode() bool
	UnsetTooltip()
	SetTooltipStyle(string)
	TooltipStyle() string
	SetCursor(string)
	Cursor() string
	UnsetCursor()
	SetCursorStyle(string)
	CursorStyle() string
	SetCursorEngineOnly(bool)
	IsCursorEngineOnly() bool
	SetStyle(string)
	Style() string
}

type widgetItemBase struct {
	obj *C.Elm_Widget_Item
}

//export go_elm_object_item_del_cb_func
func go_elm_object_item_del_cb_func(id C.int, obj *C.Eo) {
	//fmt.Printf("widget item %d is deleted from registry\n",int(id))
	registry.Delete(int(id))
}

func wrapWidgetItemBase(o *C.Elm_Widget_Item) *widgetItemBase {
	if o != nil {
		C.cgo_elm_object_item_del_cb_set(o)
		return &widgetItemBase{obj: o}
	}
	return nil
}

func (p *widgetItemBase) eo() *C.Elm_Widget_Item {
	return p.obj
}

func (p *widgetItemBase) Widget() Widget {
	return wrapWidgetBase(C.elm_object_item_widget_get(p.obj))
}

func (p *widgetItemBase) SetPartContent(part string, o Object) {
	var cpart *C.char
	if part != "" {
		cpart = C.CString(part)
		defer free(cpart)
	}
	C.elm_object_item_part_content_set(p.obj, cpart, o.eo())
}

func (p *widgetItemBase) PartContent(part string) Object {
	var cpart *C.char
	if part != "" {
		cpart = C.CString(part)
		defer free(cpart)
	}
	return wrapObjectBase(C.elm_object_item_part_content_get(p.obj, cpart))
}

func (p *widgetItemBase) UnsetPartContent(part string) Object {
	var cpart *C.char
	if part != "" {
		cpart = C.CString(part)
		defer free(cpart)
	}
	return wrapObjectBase(C.elm_object_item_part_content_unset(p.obj, cpart))
}

func (p *widgetItemBase) SetContent(o Object) {
	p.SetPartContent("", o)
}

func (p *widgetItemBase) Content() Object {
	return p.PartContent("")
}

func (p *widgetItemBase) UnsetContent() Object {
	return p.UnsetPartContent("")
}

func (p *widgetItemBase) SetPartText(part, txt string) {
	var cpart *C.char
	if part != "" {
		cpart = C.CString(part)
		defer free(cpart)
	}
	ctxt := C.CString(txt)
	defer free(ctxt)
	C.elm_object_item_part_text_set(p.obj, cpart, ctxt)
}

func (p *widgetItemBase) PartText(part string) string {
	var cpart *C.char
	if part != "" {
		cpart = C.CString(part)
		defer free(cpart)
	}
	return C.GoString(C.elm_object_item_part_text_get(p.obj, cpart))
}

func (p *widgetItemBase) SetText(txt string) {
	p.SetPartText("", txt)
}

func (p *widgetItemBase) Text() string {
	return p.PartText("")
}

func (p *widgetItemBase) SetAccessInfo(txt string) {
	ctxt := C.CString(txt)
	defer free(ctxt)
	C.elm_object_item_access_info_set(p.obj, ctxt)
}

func (p *widgetItemBase) AccessInfo() string {
	return C.GoString(C.elm_object_access_info_get(p.obj))
}

func (p *widgetItemBase) Data() interface{} {
	v := C.elm_object_item_data_get(p.obj)
	if v != nil {
		return registry.Lookup(int(*((*C.int)(v))))
	}
	return nil
}

func (p *widgetItemBase) SetData(data interface{}) {
	id := registry.Register(data)
	C.elm_object_item_data_set(p.obj, unsafe.Pointer(&id))
}

func (p *widgetItemBase) EmitSignal(emi, src string) {
	cemi := C.CString(emi)
	csrc := C.CString(src)
	defer free(cemi, csrc)
	C.elm_object_item_signal_emit(p.obj, cemi, csrc)
}

func (p *widgetItemBase) SetDisabled(b bool) {
	C.elm_object_item_disabled_set(p.obj, eBool(b))
}

func (p *widgetItemBase) IsDisabled() bool {
	return C.elm_object_item_disabled_get(p.obj) == eTrue
}

func (p *widgetItemBase) Dispose() {
	C.elm_object_item_del(p.obj)
}

func (p *widgetItemBase) SetFocus(b bool) {
	C.elm_object_item_focus_set(p.obj, eBool(b))
}

func (p *widgetItemBase) HasFocus() bool {
	return C.elm_object_item_focus_get(p.obj) == eTrue
}

func (p *widgetItemBase) NextFocusObject(dir FocusDirection) Widget {
	return wrapWidgetBase(C.elm_object_item_focus_next_object_get(p.obj, C.Elm_Focus_Direction(dir)))
}

func (p *widgetItemBase) SetNextFocusObject(o Widget, dir FocusDirection) {
	C.elm_object_item_focus_next_object_set(p.obj, o.eo(), C.Elm_Focus_Direction(dir))
}

func (p *widgetItemBase) NextFocusItem(dir FocusDirection) WidgetItem {
	return wrapWidgetItemBase(C.elm_object_item_focus_next_item_get(p.obj, C.Elm_Focus_Direction(dir)))
}

func (p *widgetItemBase) SetNextFocusItem(item WidgetItem, dir FocusDirection) {
	C.elm_object_item_focus_next_item_set(p.obj, item.eo(), C.Elm_Focus_Direction(dir))
}

func (p *widgetItemBase) SetTooltipText(txt string) {
	ctxt := C.CString(txt)
	defer free(ctxt)
	C.elm_object_item_tooltip_text_set(p.obj, ctxt)
}

func (p *widgetItemBase) SetTooltipWindowMode(b bool) {
	C.elm_object_item_tooltip_window_mode_set(p.obj, eBool(b))
}

func (p *widgetItemBase) IsTooltipWindowMode() bool {
	return C.elm_object_item_tooltip_window_mode_get(p.obj) == eTrue
}

func (p *widgetItemBase) UnsetTooltip() {
	C.elm_object_item_tooltip_unset(p.obj)
}

func (p *widgetItemBase) SetTooltipStyle(style string) {
	cstyle := C.CString(style)
	defer free(cstyle)
	C.elm_object_item_tooltip_style_set(p.obj, cstyle)
}

func (p *widgetItemBase) TooltipStyle() string {
	return C.GoString(C.elm_object_item_tooltip_style_get(p.obj))
}

func (p *widgetItemBase) SetCursor(cursor string) {
	ccursor := C.CString(cursor)
	defer free(ccursor)
	C.elm_object_item_cursor_set(p.obj, ccursor)
}

func (p *widgetItemBase) Cursor() string {
	return C.GoString(C.elm_object_item_cursor_get(p.obj))
}

func (p *widgetItemBase) UnsetCursor() {
	C.elm_object_item_cursor_unset(p.obj)
}

func (p *widgetItemBase) SetCursorStyle(style string) {
	cstyle := C.CString(style)
	defer free(cstyle)
	C.elm_object_item_cursor_style_set(p.obj, cstyle)
}

func (p *widgetItemBase) CursorStyle() string {
	return C.GoString(C.elm_object_item_cursor_style_get(p.obj))
}

func (p *widgetItemBase) SetCursorEngineOnly(b bool) {
	C.elm_object_item_cursor_engine_only_set(p.obj, eBool(b))
}

func (p *widgetItemBase) IsCursorEngineOnly() bool {
	return C.elm_object_item_cursor_engine_only_get(p.obj) == eTrue
}

func (p *widgetItemBase) SetStyle(style string) {
	cstyle := C.CString(style)
	defer free(cstyle)
	C.elm_object_item_style_set(p.obj, cstyle)
}

func (p *widgetItemBase) Style() string {
	return C.GoString(C.elm_object_item_style_get(p.obj))
}

/*
TODO:
void            elm_object_item_domain_translatable_part_text_set(Elm_Object_Item *it, const char *part, const char *domain, const char *text)
const char *    elm_object_item_translatable_part_text_get(const Elm_Object_Item *it, const char *part)
void            elm_object_item_domain_part_text_translatable_set(Elm_Object_Item *it, const char *part, const char *domain, Eina_Bool translatable)
void            elm_object_item_del_cb_set(Elm_Object_Item *it, Evas_Smart_Cb del_cb)
void            elm_object_item_tooltip_content_cb_set(Elm_Object_Item *it, Elm_Tooltip_Item_Content_Cb func, void *data, Evas_Smart_Cb del_cb)

Evas_Object *   elm_object_item_access_register(Elm_Object_Item *item)
void            elm_object_item_access_unregister(Elm_Object_Item *item)
Evas_Object *   elm_object_item_access_object_get(const Elm_Object_Item *item)
void            elm_object_item_access_order_set(Elm_Object_Item *item, Eina_List *objs)
const Eina_List *elm_object_item_access_order_get(const Elm_Object_Item *item)
void            elm_object_item_access_order_unset(Elm_Object_Item *item)

Evas_Object *   elm_object_item_track(Elm_Object_Item *it)
void            elm_object_item_untrack(Elm_Object_Item *it)
int             elm_object_item_track_get(const Elm_Object_Item *it)


void 	elm_object_focus_custom_chain_set (Evas_Object *obj, Eina_List *objs)
void 	elm_object_focus_custom_chain_unset (Evas_Object *obj)
const Eina_List * 	elm_object_focus_custom_chain_get (const Evas_Object *obj)

void 	elm_object_focus_custom_chain_append (Evas_Object *obj, Evas_Object *child, Evas_Object *relative_child)
void 	elm_object_focus_custom_chain_prepend (Evas_Object *obj, Evas_Object *child, Evas_Object *relative_child)
void 	elm_object_tree_focus_allow_set (Evas_Object *obj, Eina_Bool focusable)
Eina_Bool 	elm_object_tree_focus_allow_get (const Evas_Object *obj)

void 	elm_object_focus_region_show_mode_set (Evas_Object *obj, Elm_Focus_Region_Show_Mode mode)
Elm_Focus_Region_Show_Mode 	elm_object_focus_region_show_mode_get (const Evas_Object *obj)

Elm_Object_Item * 	elm_object_focused_item_get (const Evas_Object *obj)
Elm_Object_Item * 	elm_object_focus_next_item_get (const Evas_Object *obj, Elm_Focus_Direction dir)
void 	elm_object_focus_next_item_set (Evas_Object *obj, Elm_Object_Item *next_item, Elm_Focus_Direction dir)

void 	elm_object_domain_translatable_part_text_set (Evas_Object *obj, const char *part, const char *domain, const char *text)
const char * 	elm_object_translatable_part_text_get (const Evas_Object *obj, const char *part)
void 	elm_object_domain_part_text_translatable_set (Evas_Object *obj, const char *part, const char *domain, Eina_Bool translatable)

void 	elm_object_signal_callback_add (Evas_Object *obj, const char *emission, const char *source, Edje_Signal_Cb func, void *data)
void * 	elm_object_signal_callback_del (Evas_Object *obj, const char *emission, const char *source, Edje_Signal_Cb func)
void 	elm_object_event_callback_add (Evas_Object *obj, Elm_Event_Cb func, const void *data)
void * 	elm_object_event_callback_del (Evas_Object *obj, Elm_Event_Cb func, const void *data)

*/

//------------------------------------------------------------
