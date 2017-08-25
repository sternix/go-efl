package ui

/*
#include "bridge.h"
*/
import "C"

import (
	"unsafe"
)

//------------------------------------------------------------

type Box struct {
	*widgetBase
}

func NewBox(parent Object) *Box {
	eo := C.elm_box_add(parent.eo())
	return wrapBox(eo)
}

func wrapBox(o *C.Eo) *Box {
	if o != nil {
		return &Box{wrapWidgetBase(o)}
	}
	return nil
}

func (p *Box) SetHorizontal() {
	C.elm_box_horizontal_set(p.obj, eTrue)
}

func (p *Box) IsHorizontal() bool {
	return C.elm_box_horizontal_get(p.obj) == eTrue
}

func (p *Box) Add(o Object) {
	p.PackEnd(o)
}

func (p *Box) PackEnd(o Object) {
	C.elm_box_pack_end(p.obj, o.eo())
}

func (p *Box) PackStart(o Object) {
	C.elm_box_pack_start(p.obj, o.eo())
}

func (p *Box) PackBefore(subobj Object, before Object) {
	C.elm_box_pack_before(p.obj, subobj.eo(), before.eo())
}

func (p *Box) PackAfter(subobj Object, after Object) {
	C.elm_box_pack_after(p.obj, subobj.eo(), after.eo())
}

func (p *Box) Clear() {
	C.elm_box_clear(p.obj)
}

func (p *Box) Unpack(o Object) {
	C.elm_box_unpack(p.obj, o.eo())
}

func (p *Box) UnpackAll() {
	C.elm_box_unpack_all(p.obj)
}

func (p *Box) SetPadding(h, v int) {
	C.elm_box_padding_set(p.obj, C.int(h), C.int(v))
}

func (p *Box) Padding() (int, int) {
	var h, v C.int
	C.elm_box_padding_get(p.obj, &h, &v)
	return int(h), int(v)
}

func (p *Box) SetHomogeneous(b bool) {
	C.elm_box_homogeneous_set(p.obj, eBool(b))
}

func (p *Box) IsHomogeneous() bool {
	return C.elm_box_homogeneous_get(p.obj) == eTrue
}

func (p *Box) SetAlign(hor, ver float64) {
	C.elm_box_align_set(p.obj, C.double(hor), C.double(ver))
}

func (p *Box) Align() (float64, float64) {
	var hor, ver C.double
	C.elm_box_align_get(p.obj, &hor, &ver)
	return float64(hor), float64(ver)
}

func (p *Box) Children() []Object {
	lst := C.elm_box_children_get(p.obj)
	if lst != nil {
		return newListIterator(lst).ObjectSlice()
	}
	return nil
}

func (p *Box) Recalculate() {
	C.elm_box_recalculate(p.obj)
}

/*
TODO:
void 	elm_box_layout_set (Elm_Box *obj, Evas_Object_Box_Layout cb, const void *data, Ecore_Cb free_data)
void 	elm_box_layout_transition (Evas_Object *obj, Evas_Object_Box_Data *priv, void *data)
Elm_Box_Transition * 	elm_box_transition_new (const double duration, Evas_Object_Box_Layout start_layout, void *start_layout_data, Ecore_Cb start_layout_free_data, Evas_Object_Box_Layout end_layout, void *end_layout_data, Ecore_Cb end_layout_free_data, Ecore_Cb transition_end_cb, void *transition_end_data)
void 	elm_box_transition_free (void *data)
*/

//------------------------------------------------------------

var _ Widget = &Grid{}

type Grid struct {
	*widgetBase
}

func NewGrid(o Object) *Grid {
	eo := C.elm_grid_add(o.eo())
	return wrapGrid(eo)
}

func wrapGrid(o *C.Eo) *Grid {
	if o != nil {
		return &Grid{wrapWidgetBase(o)}
	}
	return nil
}

func (p *Grid) SetSize(w, h int) {
	C.elm_grid_size_set(p.obj, C.int(w), C.int(h))
}

func (p *Grid) Size() (int, int) {
	var w, h C.int
	C.elm_grid_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Grid) Children() []Object {
	lst := C.elm_grid_children_get(p.obj)
	if lst != nil {
		return newListIterator(lst).ObjectSlice()
	}
	return nil
}

func (p *Grid) Clear(clear bool) {
	C.elm_grid_clear(p.obj, eBool(clear))
}

func (p *Grid) Unpack(child Object) {
	C.elm_grid_unpack(p.obj, child.eo())
}

func (p *Grid) Pack(o Object, x, y, w, h int) {
	C.elm_grid_pack(p.obj, o.eo(), C.int(x), C.int(y), C.int(w), C.int(h))
}

func SetGridPackLocation(o Object, x, y, w, h int) {
	C.elm_grid_pack_set(o.eo(), C.Evas_Coord(x), C.Evas_Coord(y), C.Evas_Coord(w), C.Evas_Coord(h))
}

func GridPackLocation(o Object) (int, int, int, int) {
	var x, y, w, h C.Evas_Coord
	C.elm_grid_pack_get(o.eo(), &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

//------------------------------------------------------------

type Table struct {
	*widgetBase
}

func NewTable(o Object) *Table {
	eo := C.elm_table_add(o.eo())
	return wrapTable(eo)
}

func wrapTable(o *C.Eo) *Table {
	if o != nil {
		return &Table{wrapWidgetBase(o)}
	}
	return nil
}

func (p *Table) SetHomogeneous(b bool) {
	C.elm_table_homogeneous_set(p.obj, eBool(b))
}

func (p *Table) IsHomogeneous() bool {
	return C.elm_table_homogeneous_get(p.obj) == eTrue
}

func (p *Table) SetPadding(h, v int) {
	C.elm_table_padding_set(p.obj, C.int(h), C.int(v))
}

func (p *Table) Padding() (int, int) {
	var h, v C.int
	C.elm_table_padding_get(p.obj, &h, &v)
	return int(h), int(v)
}

func (p *Table) SetAlign(h, v float64) {
	C.elm_table_align_set(p.obj, C.double(h), C.double(v))
}

func (p *Table) Align() (float64, float64) {
	var h, v C.double
	C.elm_table_align_get(p.obj, &h, &v)
	return float64(h), float64(v)
}

func (p *Table) Clear(clear bool) {
	C.elm_table_clear(p.obj, eBool(clear))
}

func (p *Table) Child(col, row int) Object {
	return wrapObjectBase(C.elm_table_child_get(p.obj, C.int(col), C.int(row)))
}

func (p *Table) Unpack(subobj Object) {
	C.elm_table_unpack(p.obj, subobj.eo())
}

func (p *Table) Pack(so Object, col, row, colspan, rowspan int) {
	C.elm_table_pack(p.obj, so.eo(), C.int(col), C.int(row), C.int(colspan), C.int(rowspan))
}

func (p *Table) Add(o Object, col, row int) {
	C.elm_table_pack(p.obj, o.eo(), C.int(col), C.int(row), 1, 1)
}

func SetTablePackLocation(o Object, col, row, colspan, rowspan int) {
	C.elm_table_pack_set(o.eo(), C.int(col), C.int(row), C.int(colspan), C.int(rowspan))
}

func TablePackLocation(o Object) (int, int, int, int) {
	var col, row, colspan, rowspan C.int
	C.elm_table_pack_get(o.eo(), &col, &row, &colspan, &rowspan)
	return int(col), int(row), int(colspan), int(rowspan)
}

//------------------------------------------------------------

var _ Widget = &Naviframe{}

type Naviframe struct {
	*Layout
}

func NewNaviframe(o Object) *Naviframe {
	eo := C.elm_naviframe_add(o.eo())
	return wrapNaviframe(eo)
}

func wrapNaviframe(o *C.Eo) *Naviframe {
	if o != nil {
		return &Naviframe{wrapLayout(o)}
	}
	return nil
}

func (p *Naviframe) SetEventEnabled(b bool) {
	C.elm_naviframe_event_enabled_set(p.obj, eBool(b))
}

func (p *Naviframe) IsEventEnabled() bool {
	return C.elm_naviframe_event_enabled_get(p.obj) == eTrue
}

func (p *Naviframe) SetPreserveContentOnPop(b bool) {
	C.elm_naviframe_content_preserve_on_pop_set(p.obj, eBool(b))
}

func (p *Naviframe) IsPreserveContentOnPop() bool {
	return C.elm_naviframe_content_preserve_on_pop_get(p.obj) == eTrue
}

func (p *Naviframe) SetPrevButtonAutoPushed(b bool) {
	C.elm_naviframe_prev_btn_auto_pushed_set(p.obj, eBool(b))
}

func (p *Naviframe) IsPrevButtonAutoPushed() bool {
	return C.elm_naviframe_prev_btn_auto_pushed_get(p.obj) == eTrue
}

func (p *Naviframe) SetAutoCreatePrevButton(b bool) {
	C.elm_naviframe_prev_btn_auto_pushed_set(p.obj, eBool(b))
}

func (p *Naviframe) IsAutoCreatePrevButton() bool {
	return C.elm_naviframe_prev_btn_auto_pushed_get(p.obj) == eTrue
}

func (p *Naviframe) Items() []*NaviframeItem {
	lst := C.elm_naviframe_items_get(p.obj)
	if lst != nil {
		return newListIterator(lst).NaviframeItemSlice()
	}
	return nil
}

func (p *Naviframe) TopItem() *NaviframeItem {
	return wrapNaviframeItem(C.elm_naviframe_top_item_get(p.obj))
}

func (p *Naviframe) BottomItem() *NaviframeItem {
	return wrapNaviframeItem(C.elm_naviframe_bottom_item_get(p.obj))
}

func (p *Naviframe) PopItem() Object {
	return wrapObjectBase(C.elm_naviframe_item_pop(p.obj))
}

func (p *Naviframe) ItemFactory() *NaviframeItemFactory {
	return NewNaviframeItemFactory(p)
}

//------------------------------------------------------------

type Layout struct {
	*containerBase
}

var _ Container = &Layout{}

func NewLayout(o Object) *Layout {
	eo := C.elm_layout_add(o.eo())
	return wrapLayout(eo)
}

func wrapLayout(o *C.Eo) *Layout {
	if o != nil {
		return &Layout{wrapContainerBase(o)}
	}
	return nil
}

func (p *Layout) SetEdjeObjectCanAccess(b bool) {
	C.elm_layout_edje_object_can_access_set(p.obj, eBool(b))
}

func (p *Layout) EdjeObjectCanAccess() bool {
	return C.elm_layout_edje_object_can_access_get(p.obj) == eTrue
}

func (p *Layout) SetLayoutTheme(klass string, group string, style string) bool {
	cklass := C.CString(klass)
	cgroup := C.CString(group)
	cstyle := C.CString(style)
	defer free(cklass, cgroup, cstyle)
	return C.elm_layout_theme_set(p.obj, cklass, cgroup, cstyle) == eTrue
}

func (p *Layout) Freeze() int {
	return int(C.elm_layout_freeze(p.obj))
}

func (p *Layout) Thaw() int {
	return int(C.elm_layout_thaw(p.obj))
}

func (p *Layout) EmitSignal(e string, source string) {
	ce := C.CString(e)
	csource := C.CString(source)
	defer free(ce, csource)
	C.elm_layout_signal_emit(p.obj, ce, csource)
}

type layoutBox struct {
	*Layout
}

func (p *Layout) Box() *layoutBox {
	return &layoutBox{
		Layout: p,
	}
}

func (p *layoutBox) Append(part string, o Object) bool {
	cpart := C.CString(part)
	defer free(cpart)
	return C.elm_layout_box_append(p.obj, cpart, o.eo()) == eTrue
}

func (p *layoutBox) Prepend(part string, o Object) bool {
	cpart := C.CString(part)
	defer free(cpart)
	return C.elm_layout_box_prepend(p.obj, cpart, o.eo()) == eTrue
}

func (p *layoutBox) InsertBefore(part string, child Object, reference Object) bool {
	cpart := C.CString(part)
	defer free(cpart)
	return C.elm_layout_box_insert_before(p.obj, cpart, child.eo(), reference.eo()) == eTrue
}

func (p *layoutBox) InsertAt(part string, o Object, pos uint) bool {
	cpart := C.CString(part)
	defer free(cpart)
	return C.elm_layout_box_insert_at(p.obj, cpart, o.eo(), C.uint(pos)) == eTrue
}

func (p *layoutBox) Remove(part string, o Object) bool {
	cpart := C.CString(part)
	defer free(cpart)
	return C.elm_layout_box_remove(p.obj, cpart, o.eo()) != nil // if != nil then found
}

func (p *layoutBox) RemoveAll(part string, clear bool) bool {
	cpart := C.CString(part)
	defer free(cpart)
	return C.elm_layout_box_remove_all(p.obj, cpart, eBool(clear)) == eTrue
}

//------------------------------------------------------------

type Panes struct {
	*Layout
}

func NewPanes(o Object) *Panes {
	eo := C.elm_panes_add(o.eo())
	return wrapPanes(eo)
}

func wrapPanes(o *C.Eo) *Panes {
	if o != nil {
		return &Panes{wrapLayout(o)}
	}
	return nil
}

func (p *Panes) SetLeftContentSize(size float64) {
	C.elm_panes_content_left_size_set(p.obj, C.double(size))
}

func (p *Panes) LeftContentSize() float64 {
	return float64(C.elm_panes_content_left_size_get(p.obj))
}

func (p *Panes) SetHorizontal(b bool) {
	C.elm_panes_horizontal_set(p.obj, eBool(b))
}

func (p *Panes) IsHorizontal() bool {
	return C.elm_panes_horizontal_get(p.obj) == eTrue
}

func (p *Panes) SetFixed(b bool) {
	C.elm_panes_fixed_set(p.obj, eBool(b))
}

func (p *Panes) IsFixed() bool {
	return C.elm_panes_fixed_get(p.obj) == eTrue
}

func (p *Panes) SetRightContentSize(size float64) {
	C.elm_panes_content_right_size_set(p.obj, C.double(size))
}

func (p *Panes) RightContentSize() float64 {
	return float64(C.elm_panes_content_right_size_get(p.obj))
}

func (p *Panes) SetLeftContentRelativeMinSize(size float64) {
	C.elm_panes_content_left_min_relative_size_set(p.obj, C.double(size))
}

func (p *Panes) LeftContentRelativeMinSize() float64 {
	return float64(C.elm_panes_content_left_min_relative_size_get(p.obj))
}

func (p *Panes) SetRightContentRelativeMinSize(size float64) {
	C.elm_panes_content_right_min_relative_size_set(p.obj, C.double(size))
}

func (p *Panes) RightContentRelativeMinSize() float64 {
	return float64(C.elm_panes_content_right_min_relative_size_get(p.obj))
}

func (p *Panes) SetLeftContentMinSize(size int) {
	C.elm_panes_content_left_min_size_set(p.obj, C.int(size))
}

func (p *Panes) LeftContentMinSize() int {
	return int(C.elm_panes_content_left_min_size_get(p.obj))
}

func (p *Panes) SetRightContentMinSize(size int) {
	C.elm_panes_content_right_min_size_set(p.obj, C.int(size))
}

func (p *Panes) RightContentMinSize() int {
	return int(C.elm_panes_content_right_min_size_get(p.obj))
}

//------------------------------------------------------------

var _ Scrollable = &Scroller{}

type Scroller struct {
	*Layout
	*scrollableBase
}

func NewScroller(o Object) *Scroller {
	eo := C.elm_scroller_add(o.eo())
	return &Scroller{
		Layout:         wrapLayout(eo),
		scrollableBase: wrapScrollableBase(eo),
	}
}

type Scrollable interface {
	SetPropagateEvents(bool)
	IsPropagateEvents() bool
	SetPageScrollLimit(int, int)
	PageScrollLimit() (int, int)
	SetMinContentLimit(bool, bool)
	ShowRegion(int, int, int, int)
	SetPolicy(ScrollerPolicy, ScrollerPolicy)
	Policy() (ScrollerPolicy, ScrollerPolicy)
	SetSingleDirection(ScrollerSingleDirection)
	SingleDirection() ScrollerSingleDirection
	Region() (int, int, int, int)
	ChildSize() (int, int)
	SetPageSnap(bool, bool)
	PageSnap() (bool, bool)
	SetBounce(bool, bool)
	Bounce() (bool, bool)
	SetPageRelative(float64, float64)
	PageRelative() (float64, float64)
	SetPageSize(int, int)
	PageSize() (int, int)
	CurrentPage() (int, int)
	LastPage() (int, int)
	ShowPage(int, int)
	BringPageIn(int, int)
	BringRegionIn(int, int, int, int)
	SetGravity(float64, float64)
	Gravity() (float64, float64)
	SetMovementBlock(ScrollerMovementBlock)
	MovementBlock() ScrollerMovementBlock
	SetStepSize(int, int)
	StepSize() (int, int)
	SetLoop(bool, bool)
	IsLoop() (bool, bool)
	SetWheelDisabled(bool)
	IsWheelDisabled() bool
}

type scrollableBase struct {
	obj *C.Eo
}

func wrapScrollableBase(o *C.Eo) *scrollableBase {
	if o != nil {
		return &scrollableBase{o}
	}
	return nil
}

func (p *scrollableBase) SetPropagateEvents(b bool) {
	C.elm_scroller_propagate_events_set(p.obj, eBool(b))
}

func (p *scrollableBase) IsPropagateEvents() bool {
	return C.elm_scroller_propagate_events_get(p.obj) == eTrue
}

func (p *scrollableBase) SetPageScrollLimit(horlimit, verlimit int) {
	C.elm_scroller_page_scroll_limit_set(p.obj, C.int(horlimit), C.int(verlimit))
}

func (p *scrollableBase) PageScrollLimit() (int, int) {
	var hlimit, vlimit C.int
	C.elm_scroller_page_scroll_limit_get(p.obj, &hlimit, &vlimit)
	return int(hlimit), int(vlimit)
}

func (p *scrollableBase) SetMinContentLimit(w, h bool) {
	C.elm_scroller_content_min_limit(p.obj, eBool(w), eBool(h))
}

func (p *scrollableBase) ShowRegion(x, y, w, h int) {
	C.elm_scroller_region_show(p.obj, C.Evas_Coord(x), C.Evas_Coord(y), C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *scrollableBase) SetPolicy(hpol ScrollerPolicy, vpol ScrollerPolicy) {
	C.elm_scroller_policy_set(p.obj, C.Elm_Scroller_Policy(hpol), C.Elm_Scroller_Policy(vpol))
}

func (p *scrollableBase) Policy() (ScrollerPolicy, ScrollerPolicy) {
	var hpol, vpol C.Elm_Scroller_Policy
	C.elm_scroller_policy_get(p.obj, &hpol, &vpol)
	return ScrollerPolicy(hpol), ScrollerPolicy(vpol)
}

func (p *scrollableBase) SetSingleDirection(dir ScrollerSingleDirection) {
	C.elm_scroller_single_direction_set(p.obj, C.Elm_Scroller_Single_Direction(dir))
}

func (p *scrollableBase) SingleDirection() ScrollerSingleDirection {
	return ScrollerSingleDirection(C.elm_scroller_single_direction_get(p.obj))
}

func (p *scrollableBase) Region() (int, int, int, int) {
	var x, y, w, h C.Evas_Coord
	C.elm_scroller_region_get(p.obj, &x, &y, &w, &h)
	return int(x), int(y), int(w), int(h)
}

func (p *scrollableBase) ChildSize() (int, int) {
	var w, h C.Evas_Coord
	C.elm_scroller_child_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *scrollableBase) SetPageSnap(hsnap, vsnap bool) {
	C.elm_scroller_page_snap_set(p.obj, eBool(hsnap), eBool(vsnap))
}

func (p *scrollableBase) PageSnap() (bool, bool) {
	var hsnap, vsnap C.Eina_Bool
	C.elm_scroller_page_snap_get(p.obj, &hsnap, &vsnap)
	return goBool(hsnap), goBool(vsnap)
}

func (p *scrollableBase) SetBounce(hbounce, vbounce bool) {
	C.elm_scroller_bounce_set(p.obj, eBool(hbounce), eBool(vbounce))
}

func (p *scrollableBase) Bounce() (bool, bool) {
	var hbounce, vbounce C.Eina_Bool
	C.elm_scroller_bounce_get(p.obj, &hbounce, &vbounce)
	return goBool(hbounce), goBool(vbounce)
}

func (p *scrollableBase) SetPageRelative(hrel, vrel float64) {
	C.elm_scroller_page_relative_set(p.obj, C.double(hrel), C.double(vrel))
}

func (p *scrollableBase) PageRelative() (float64, float64) {
	var hrel, vrel C.double
	C.elm_scroller_page_relative_get(p.obj, &hrel, &vrel)
	return float64(hrel), float64(vrel)
}

func (p *scrollableBase) SetPageSize(hsize, vsize int) {
	C.elm_scroller_page_size_set(p.obj, C.Evas_Coord(hsize), C.Evas_Coord(vsize))
}

func (p *scrollableBase) PageSize() (int, int) {
	var hsize, vsize C.Evas_Coord
	C.elm_scroller_page_size_get(p.obj, &hsize, &vsize)
	return int(hsize), int(vsize)
}

func (p *scrollableBase) CurrentPage() (int, int) {
	var hnum, vnum C.int
	C.elm_scroller_current_page_get(p.obj, &hnum, &vnum)
	return int(hnum), int(vnum)
}

func (p *scrollableBase) LastPage() (int, int) {
	var hnum, vnum C.int
	C.elm_scroller_last_page_get(p.obj, &hnum, &vnum)
	return int(hnum), int(vnum)
}

func (p *scrollableBase) ShowPage(hnum, vnum int) {
	C.elm_scroller_page_show(p.obj, C.int(hnum), C.int(vnum))
}

func (p *scrollableBase) BringPageIn(hnum, vnum int) {
	C.elm_scroller_page_bring_in(p.obj, C.int(hnum), C.int(vnum))
}

func (p *scrollableBase) BringRegionIn(x, y, w, h int) {
	C.elm_scroller_region_bring_in(p.obj, C.Evas_Coord(x), C.Evas_Coord(y), C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *scrollableBase) SetGravity(x, y float64) {
	C.elm_scroller_gravity_set(p.obj, C.double(x), C.double(y))
}

func (p *scrollableBase) Gravity() (float64, float64) {
	var x, y C.double
	C.elm_scroller_gravity_get(p.obj, &x, &y)
	return float64(x), float64(y)
}

func (p *scrollableBase) SetMovementBlock(typ ScrollerMovementBlock) {
	C.elm_scroller_movement_block_set(p.obj, C.Elm_Scroller_Movement_Block(typ))
}

func (p *scrollableBase) MovementBlock() ScrollerMovementBlock {
	return ScrollerMovementBlock(C.elm_scroller_movement_block_get(p.obj))
}

func (p *scrollableBase) SetStepSize(x, y int) {
	C.elm_scroller_step_size_set(p.obj, C.Evas_Coord(x), C.Evas_Coord(y))
}

func (p *scrollableBase) StepSize() (int, int) {
	var x, y C.Evas_Coord
	C.elm_scroller_step_size_get(p.obj, &x, &y)
	return int(x), int(y)
}

func (p *scrollableBase) SetLoop(looph, loopv bool) {
	C.elm_scroller_loop_set(p.obj, eBool(looph), eBool(loopv))
}

func (p *scrollableBase) IsLoop() (bool, bool) {
	var hloop, vloop C.Eina_Bool
	C.elm_scroller_loop_get(p.obj, &hloop, &vloop)
	return goBool(hloop), goBool(vloop)
}

func (p *scrollableBase) SetWheelDisabled(b bool) {
	C.elm_scroller_wheel_disabled_set(p.obj, eBool(b))
}

func (p *scrollableBase) IsWheelDisabled() bool {
	return C.elm_scroller_wheel_disabled_get(p.obj) == eTrue
}

//------------------------------------------------------------

type Genlist struct {
	*Layout
}

func NewGenlist(parent Object) *Genlist {
	return &Genlist{wrapLayout(C.elm_genlist_add(parent.eo()))}
}

func (p *Genlist) SetHomogeneous(b bool) {
	C.elm_genlist_homogeneous_set(p.obj, eBool(b))
}

func (p *Genlist) IsHomogeneous() bool {
	return C.elm_genlist_homogeneous_get(p.obj) == eTrue
}

func (p *Genlist) SetSelectMode(mode SelectMode) {
	C.elm_genlist_select_mode_set(p.obj, C.Elm_Object_Select_Mode(mode))
}

func (p *Genlist) SelectMode() SelectMode {
	return SelectMode(C.elm_genlist_select_mode_get(p.obj))
}

func (p *Genlist) SetFocusOnSelection(b bool) {
	C.elm_genlist_focus_on_selection_set(p.obj, eBool(b))
}

func (p *Genlist) IsFocusOnSelection() bool {
	return C.elm_genlist_focus_on_selection_get(p.obj) == eTrue
}

func (p *Genlist) SetLongpressTimeout(t float64) {
	C.elm_genlist_longpress_timeout_set(p.obj, C.double(t))
}

func (p *Genlist) LongpressTimeout() float64 {
	return float64(C.elm_genlist_longpress_timeout_get(p.obj))
}

func (p *Genlist) SetMultiSelect(b bool) {
	C.elm_genlist_multi_select_set(p.obj, eBool(b))
}

func (p *Genlist) IsMultiSelect() bool {
	return C.elm_genlist_multi_select_get(p.obj) == eTrue
}

func (p *Genlist) SetReorderMode(b bool) {
	C.elm_genlist_reorder_mode_set(p.obj, eBool(b))
}

func (p *Genlist) IsReorderMode() bool {
	return C.elm_genlist_reorder_mode_get(p.obj) == eTrue
}

func (p *Genlist) SetDecorateMode(b bool) {
	C.elm_genlist_decorate_mode_set(p.obj, eBool(b))
}

func (p *Genlist) IsDecorateMode() bool {
	return C.elm_genlist_decorate_mode_get(p.obj) == eTrue
}

func (p *Genlist) SetMultiSelectMode(mode MultiSelectMode) {
	C.elm_genlist_multi_select_mode_set(p.obj, C.Elm_Object_Multi_Select_Mode(mode))
}

func (p *Genlist) MultiSelectMode() MultiSelectMode {
	return MultiSelectMode(C.elm_genlist_multi_select_mode_get(p.obj))
}

func (p *Genlist) SetBlockCount(count int) {
	C.elm_genlist_block_count_set(p.obj, C.int(count))
}

func (p *Genlist) BlockCount() int {
	return int(C.elm_genlist_block_count_get(p.obj))
}

func (p *Genlist) SetTreeEffectEnabled(b bool) {
	C.elm_genlist_tree_effect_enabled_set(p.obj, eBool(b))
}

func (p *Genlist) IsTreeEffectEnabled() bool {
	return C.elm_genlist_tree_effect_enabled_get(p.obj) == eTrue
}

func (p *Genlist) SetHighlightMode(b bool) {
	C.elm_genlist_highlight_mode_set(p.obj, eBool(b))
}

func (p *Genlist) IsHighlightMode() bool {
	return C.elm_genlist_highlight_mode_get(p.obj) == eTrue
}

func (p *Genlist) SetMode(mode ListMode) {
	C.elm_genlist_mode_set(p.obj, C.Elm_List_Mode(mode))
}

func (p *Genlist) Mode() ListMode {
	return ListMode(C.elm_genlist_mode_get(p.obj))
}

func (p *Genlist) DecoratedItem() *GenlistItem {
	return wrapGenlistItem(C.elm_genlist_decorated_item_get(p.obj))
}

func (p *Genlist) SelectedItem() *GenlistItem {
	return wrapGenlistItem(C.elm_genlist_selected_item_get(p.obj))
}

func (p *Genlist) FirstItem() *GenlistItem {
	return wrapGenlistItem(C.elm_genlist_first_item_get(p.obj))
}

func (p *Genlist) LastItem() *GenlistItem {
	return wrapGenlistItem(C.elm_genlist_last_item_get(p.obj))
}

func (p *Genlist) RealizedItems() []*GenlistItem {
	return newListIterator(C.elm_genlist_realized_items_get(p.obj)).GenlistItemSlice()
}

func (p *Genlist) SelectedItems() []*GenlistItem {
	return newListIterator(C.elm_genlist_selected_items_get(p.obj)).GenlistItemSlice()
}

func (p *Genlist) UpdateRealizedItems() {
	C.elm_genlist_realized_items_update(p.obj)
}

func (p *Genlist) ItemAtXY(x, y int) (*GenlistItem, int) {
	var posrel C.int
	wi := C.elm_genlist_at_xy_item_get(p.obj, C.int(x), C.int(y), &posrel)
	return wrapGenlistItem(wi), int(posrel)
}

func (p *Genlist) NthItem(n uint) *GenlistItem {
	return wrapGenlistItem(C.elm_genlist_nth_item_get(p.obj, C.uint(n)))
}

// TODO: this isnot legal put this filter key to a map, call c function with map id
func (p *Genlist) SetFilter(key interface{}) {
	C.elm_genlist_filter_set(p.obj, unsafe.Pointer(&key))
}

func (p *Genlist) FilteredItemsCount() uint {
	return uint(C.elm_genlist_filtered_items_count(p.obj))
}

func (p *Genlist) ItemsCount() uint {
	return uint(C.elm_genlist_items_count(p.obj))
}

func (p *Genlist) Clear() {
	C.elm_genlist_clear(p.obj)
}

func (p *Genlist) SearchItemByText(from *GenlistItem, partname string, pattern string, flag GlobMatchFlag) *GenlistItem {
	cpartname := C.CString(partname)
	cpattern := C.CString(pattern)
	defer free(cpartname, cpattern)
	return wrapGenlistItem(C.elm_genlist_search_by_text_item_get(p.obj, from.obj, cpartname, cpattern, C.Elm_Glob_Match_Flags(flag)))
}

func (p *Genlist) ItemFactory() *GenlistItemFactory {
	return NewGenlistItemFactory(p)
}

/*
Eina_Iterator * 	elm_genlist_filter_iterator_new (Elm_Genlist *obj)

TODO:
void 	elm_genlist_item_tooltip_content_cb_set (Elm_Object_Item *it, Elm_Tooltip_Item_Content_Cb func, const void *data, Evas_Smart_Cb del_cb)

*/

//------------------------------------------------------------

type Gengrid struct {
	*Layout
}

func NewGengrid(parent Object) *Gengrid {
	return &Gengrid{wrapLayout(C.elm_gengrid_add(parent.eo()))}
}

func (p *Gengrid) SetAlign(x, y float64) {
	C.elm_gengrid_align_set(p.obj, C.double(x), C.double(y))
}

func (p *Gengrid) Align() (float64, float64) {
	var x, y C.double
	C.elm_gengrid_align_get(p.obj, &x, &y)
	return float64(x), float64(y)
}

func (p *Gengrid) SetFilled(b bool) {
	C.elm_gengrid_filled_set(p.obj, eBool(b))
}

func (p *Gengrid) IsFilled() bool {
	return C.elm_gengrid_filled_get(p.obj) == eTrue
}

func (p *Gengrid) SetPageRelative(x, y float64) {
	C.elm_gengrid_page_relative_set(p.obj, C.double(x), C.double(y))
}

func (p *Gengrid) PageRelative() (float64, float64) {
	var x, y C.double
	C.elm_gengrid_page_relative_get(p.obj, &x, &y)
	return float64(x), float64(y)
}

func (p *Gengrid) SetMultiSelect(b bool) {
	C.elm_gengrid_multi_select_set(p.obj, eBool(b))
}

func (p *Gengrid) IsMultiSelect() bool {
	return C.elm_gengrid_multi_select_get(p.obj) == eTrue
}

func (p *Gengrid) SetGroupItemSize(w, h int) {
	C.elm_gengrid_group_item_size_set(p.obj, C.int(w), C.int(h))
}

func (p *Gengrid) GroupItemSize() (int, int) {
	var w, h C.int
	C.elm_gengrid_group_item_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Gengrid) SetSelectMode(mode SelectMode) {
	C.elm_gengrid_select_mode_set(p.obj, C.Elm_Object_Select_Mode(mode))
}

func (p *Gengrid) SelectMode() SelectMode {
	return SelectMode(C.elm_gengrid_select_mode_get(p.obj))
}

func (p *Gengrid) SetReorderMode(b bool) {
	C.elm_gengrid_reorder_mode_set(p.obj, eBool(b))
}

func (p *Gengrid) IsReorderMode() bool {
	return C.elm_gengrid_reorder_mode_get(p.obj) == eTrue
}

func (p *Gengrid) SetHighlightMode(b bool) {
	C.elm_gengrid_highlight_mode_set(p.obj, eBool(b))
}

func (p *Gengrid) IsHighlightMode() bool {
	return C.elm_gengrid_highlight_mode_get(p.obj) == eTrue
}

func (p *Gengrid) SetReorderType(t GengridReorderType) {
	C.elm_gengrid_reorder_type_set(p.obj, C.Elm_Gengrid_Reorder_Type(t))
}

func (p *Gengrid) SetItemSize(w, h int) {
	C.elm_gengrid_item_size_set(p.obj, C.int(w), C.int(h))
}

func (p *Gengrid) ItemSize() (int, int) {
	var w, h C.int
	C.elm_gengrid_item_size_get(p.obj, &w, &h)
	return int(w), int(h)
}

func (p *Gengrid) SetMultiSelectMode(mode MultiSelectMode) {
	C.elm_gengrid_multi_select_mode_set(p.obj, C.Elm_Object_Multi_Select_Mode(mode))
}

func (p *Gengrid) MultiSelectMode() MultiSelectMode {
	return MultiSelectMode(C.elm_gengrid_multi_select_mode_get(p.obj))
}

func (p *Gengrid) SetHorizontal(b bool) {
	C.elm_gengrid_horizontal_set(p.obj, eBool(b))
}

func (p *Gengrid) IsHorizontal() bool {
	return C.elm_gengrid_horizontal_get(p.obj) == eTrue
}

func (p *Gengrid) SetPageSize(h, v int) {
	C.elm_gengrid_page_size_set(p.obj, C.int(h), C.int(h))
}

func (p *Gengrid) SelectedItem() *GengridItem {
	return wrapGengridItem(C.elm_gengrid_selected_item_get(p.obj))
}

func (p *Gengrid) RealizedItems() []*GengridItem {
	return newListIterator(C.elm_gengrid_realized_items_get(p.obj)).GengridItemSlice()
}

func (p *Gengrid) FirstItem() *GengridItem {
	return wrapGengridItem(C.elm_gengrid_first_item_get(p.obj))
}

func (p *Gengrid) SelectedItems() []*GengridItem {
	return newListIterator(C.elm_gengrid_selected_items_get(p.obj)).GengridItemSlice()
}

func (p *Gengrid) LastItem() *GengridItem {
	return wrapGengridItem(C.elm_gengrid_last_item_get(p.obj))
}

func (p *Gengrid) SetWheelDisabled(b bool) {
	C.elm_gengrid_wheel_disabled_set(p.obj, eBool(b))
}

func (p *Gengrid) IsWheelDisabled() bool {
	return C.elm_gengrid_wheel_disabled_get(p.obj) == eTrue
}

func (p *Gengrid) UpdateRealizedItems() {
	C.elm_gengrid_realized_items_update(p.obj)
}

func (p *Gengrid) ItemsCount() uint {
	return uint(C.elm_gengrid_items_count(p.obj))
}

func (p *Gengrid) ItemAtXY(x, y int) (*GengridItem, int, int) {
	var xposrel, yposrel C.int
	wi := C.elm_gengrid_at_xy_item_get(p.obj, C.int(x), C.int(y), &xposrel, &yposrel)
	return wrapGengridItem(wi), int(xposrel), int(yposrel)
}

func (p *Gengrid) Clear() {
	C.elm_gengrid_clear(p.obj)
}

func (p *Gengrid) StartReorderMode(mode PositionMap) {
	C.elm_gengrid_reorder_mode_start(p.obj, C.Ecore_Pos_Map(mode))
}

func (p *Gengrid) StopReorderMode() {
	C.elm_gengrid_reorder_mode_stop(p.obj)
}

func (p *Gengrid) ItemFactory() *GengridItemFactory {
	return NewGengridItemFactory(p)
}

/*
TODO:

Elm_Widget_Item * 	elm_gengrid_item_sorted_insert (Elm_Gengrid *obj, const Elm_Gengrid_Item_Class *itc, const void *data, Eina_Compare_Cb comp, Evas_Smart_Cb func, const void *func_data)

Elm_Widget_Item * 	elm_gengrid_search_by_text_item_get (Elm_Gengrid *obj, Elm_Widget_Item *item_to_search_from, const char *part_name, const char *pattern, Elm_Glob_Match_Flags flags)

void 	elm_gengrid_item_tooltip_content_cb_set (Elm_Object_Item *it, Elm_Tooltip_Item_Content_Cb func, const void *data, Evas_Smart_Cb del_cb)

*/

//------------------------------------------------------------
