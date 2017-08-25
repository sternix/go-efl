package ui

/*
#include "bridge.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type MenuItem struct {
	*widgetItemBase
}

func wrapMenuItem(o *C.Elm_Widget_Item) *MenuItem {
	if o != nil {
		return &MenuItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *MenuItem) IsSeperator() bool {
	return C.elm_menu_item_is_separator(p.obj) == eTrue
}

func (p *MenuItem) Object() Object {
	return wrapObjectBase(C.elm_menu_item_object_get(p.obj))
}

func (p *MenuItem) SetIconName(name string) {
	cname := C.CString(name)
	defer free(cname)
	C.elm_menu_item_icon_name_set(p.obj, cname)
}

func (p *MenuItem) IconName() string {
	return C.GoString(C.elm_menu_item_icon_name_get(p.obj))
}

func (p *MenuItem) SetSelected(b bool) {
	C.elm_menu_item_selected_set(p.obj, eBool(b))
}

func (p *MenuItem) IsSelected() bool {
	return C.elm_menu_item_selected_get(p.obj) == eTrue
}

func (p *MenuItem) SubItems() []*MenuItem {
	return newListIterator(C.elm_menu_item_subitems_get(p.obj)).MenuItemSlice()
}

func (p *MenuItem) ClearSubitems() {
	C.elm_menu_item_subitems_clear(p.obj)
}

func (p *MenuItem) Index() int {
	return int(C.elm_menu_item_index_get(p.obj))
}

func (p *MenuItem) Next() *MenuItem {
	return wrapMenuItem(C.elm_menu_item_next_get(p.obj))
}

func (p *MenuItem) Prev() *MenuItem {
	return wrapMenuItem(C.elm_menu_item_prev_get(p.obj))
}

//------------------------------------------------------------

type ToolbarItem struct {
	*widgetItemBase
}

func wrapToolbarItem(o *C.Elm_Widget_Item) *ToolbarItem {
	if o != nil {
		return &ToolbarItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *ToolbarItem) Next() *ToolbarItem {
	return wrapToolbarItem(C.elm_toolbar_item_next_get(p.obj))
}

func (p *ToolbarItem) Prev() *ToolbarItem {
	return wrapToolbarItem(C.elm_toolbar_item_prev_get(p.obj))
}

func (p *ToolbarItem) SetPriority(pri int) {
	C.elm_toolbar_item_priority_set(p.obj, C.int(pri))
}

func (p *ToolbarItem) Priority() int {
	return int(C.elm_toolbar_item_priority_get(p.obj))
}

func (p *ToolbarItem) SetIcon(name string) {
	cname := C.CString(name)
	defer free(cname)
	C.elm_toolbar_item_icon_set(p.obj, cname)
}

func (p *ToolbarItem) Icon() string {
	return C.GoString(C.elm_toolbar_item_icon_get(p.obj))
}

func (p *ToolbarItem) IconObject() *Icon {
	return wrapIcon(C.elm_toolbar_item_icon_object_get(p.obj))
}

func (p *ToolbarItem) SetSeperator(b bool) {
	C.elm_toolbar_item_separator_set(p.obj, eBool(b))
}

func (p *ToolbarItem) IsSeperator() bool {
	return C.elm_toolbar_item_separator_get(p.obj) == eTrue
}

func (p *ToolbarItem) SetSelected(b bool) {
	C.elm_toolbar_item_selected_set(p.obj, eBool(b))
}

func (p *ToolbarItem) IsSelected() bool {
	return C.elm_toolbar_item_selected_get(p.obj) == eTrue
}

func (p *ToolbarItem) SetIsMenu(b bool) {
	C.elm_toolbar_item_menu_set(p.obj, eBool(b))
}

func (p *ToolbarItem) Menu() *Menu {
	return wrapMenu(C.elm_toolbar_item_menu_get(p.obj))
}

func (p *ToolbarItem) Show(t ToolbarItemScrolltoType) {
	C.elm_toolbar_item_show(p.obj, C.Elm_Toolbar_Item_Scrollto_Type(t))
}

func (p *ToolbarItem) Object() Object {
	return wrapObjectBase(C.elm_toolbar_item_object_get(p.obj))
}

func (p *ToolbarItem) SetIconFile(file string, key string) bool {
	cfile := C.CString(file)
	ckey := C.CString(key)
	defer free(cfile, ckey)
	return C.elm_toolbar_item_icon_file_set(p.obj, cfile, ckey) == eTrue
}

//TODO AddState with data
func (p *ToolbarItem) AddState(icon string, label string, handler Handler) *ToolbarItemState {
	var (
		cicon, clabel *C.char
		data          unsafe.Pointer
		smartcb       C.Evas_Smart_Cb
	)

	if icon != "" {
		cicon = C.CString(icon)
		defer free(cicon)
	}

	if label != "" {
		clabel = C.CString(label)
		defer free(clabel)
	}

	if handler != nil {
		handler_id := registerHandler(handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	return wrapToolbarItemState(C.elm_toolbar_item_state_add(p.obj, cicon, clabel, smartcb, data))
}

func (p *ToolbarItem) DelState(state *ToolbarItemState) bool {
	return C.elm_toolbar_item_state_del(p.obj, state.obj) == eTrue
}

func (p *ToolbarItem) SetState(state *ToolbarItemState) bool {
	return C.elm_toolbar_item_state_set(p.obj, state.obj) == eTrue
}

func (p *ToolbarItem) UnsetState() {
	C.elm_toolbar_item_state_unset(p.obj)
}

func (p *ToolbarItem) State() *ToolbarItemState {
	return wrapToolbarItemState(C.elm_toolbar_item_state_get(p.obj))
}

func (p *ToolbarItem) NextState() *ToolbarItemState {
	return wrapToolbarItemState(C.elm_toolbar_item_state_next(p.obj))
}

func (p *ToolbarItem) PrevState() *ToolbarItemState {
	return wrapToolbarItemState(C.elm_toolbar_item_state_prev(p.obj))
}

type ToolbarItemState struct {
	obj *C.Elm_Toolbar_Item_State
}

func wrapToolbarItemState(o *C.Elm_Toolbar_Item_State) *ToolbarItemState {
	if o != nil {
		return &ToolbarItemState{o}
	}
	return nil
}

//------------------------------------------------------------

type IndexItem struct {
	*widgetItemBase
}

func wrapIndexItem(o *C.Elm_Widget_Item) *IndexItem {
	if o != nil {
		return &IndexItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *IndexItem) SetSelected(b bool) {
	C.elm_index_item_selected_set(p.obj, eBool(b))
}

func (p *IndexItem) Letter() string {
	return C.GoString(C.elm_index_item_letter_get(p.obj))
}

func (p *IndexItem) SetPriority(pri int) {
	C.elm_index_item_priority_set(p.obj, C.int(pri))
}

//------------------------------------------------------------

type ListItem struct {
	*widgetItemBase
}

func wrapListItem(o *C.Elm_Widget_Item) *ListItem {
	if o != nil {
		return &ListItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *ListItem) SetSelected(b bool) {
	C.elm_list_item_selected_set(p.obj, eBool(b))
}

func (p *ListItem) IsSelected() bool {
	return C.elm_list_item_selected_get(p.obj) == eTrue
}

func (p *ListItem) SetSeperator(b bool) {
	C.elm_list_item_separator_set(p.obj, eBool(b))
}

func (p *ListItem) IsSeperator() bool {
	return C.elm_list_item_separator_get(p.obj) == eTrue
}

func (p *ListItem) Show() {
	C.elm_list_item_show(p.obj)
}

func (p *ListItem) BringIn() {
	C.elm_list_item_bring_in(p.obj)
}

func (p *ListItem) Object() Object {
	return wrapObjectBase(C.elm_list_item_object_get(p.obj))
}

func (p *ListItem) Prev() *ListItem {
	return wrapListItem(C.elm_list_item_prev(p.obj))
}

func (p *ListItem) Next() *ListItem {
	return wrapListItem(C.elm_list_item_next(p.obj))
}

//------------------------------------------------------------

type SlideshowItem struct {
	*widgetItemBase
}

func wrapSlideshowItem(o *C.Elm_Widget_Item) *SlideshowItem {
	if o != nil {
		return &SlideshowItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *SlideshowItem) Show() {
	C.elm_slideshow_item_show(p.obj)
}

func (p *SlideshowItem) Object() Object {
	return wrapObjectBase(C.elm_slideshow_item_object_get(p.obj))
}

//------------------------------------------------------------

type MultibuttonentryItem struct {
	*widgetItemBase
}

func wrapMultibuttonentryItem(o *C.Elm_Widget_Item) *MultibuttonentryItem {
	if o != nil {
		return &MultibuttonentryItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *MultibuttonentryItem) SetSelected(b bool) {
	C.elm_multibuttonentry_item_selected_set(p.obj, eBool(b))
}

func (p *MultibuttonentryItem) IsSelected() bool {
	return C.elm_multibuttonentry_item_selected_get(p.obj) == eTrue
}

func (p *MultibuttonentryItem) Prev() *MultibuttonentryItem {
	return wrapMultibuttonentryItem(C.elm_multibuttonentry_item_prev_get(p.obj))
}

func (p *MultibuttonentryItem) Next() *MultibuttonentryItem {
	return wrapMultibuttonentryItem(C.elm_multibuttonentry_item_next_get(p.obj))
}

//------------------------------------------------------------

type SegmentControlItem struct {
	*widgetItemBase
}

func wrapSegmentControlItem(o *C.Elm_Widget_Item) *SegmentControlItem {
	if o != nil {
		return &SegmentControlItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *SegmentControlItem) Index() int {
	return int(C.elm_segment_control_item_index_get(p.obj))
}

func (p *SegmentControlItem) Object() Object {
	return wrapObjectBase(C.elm_segment_control_item_object_get(p.obj))
}

func (p *SegmentControlItem) SetSelected(b bool) {
	C.elm_segment_control_item_selected_set(p.obj, eBool(b))
}

//------------------------------------------------------------

type FlipselectorItem struct {
	*widgetItemBase
}

func wrapFlipselectorItem(o *C.Elm_Widget_Item) *FlipselectorItem {
	if o != nil {
		return &FlipselectorItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *FlipselectorItem) SetSelected(b bool) {
	C.elm_flipselector_item_selected_set(p.obj, eBool(b))
}

func (p *FlipselectorItem) IsSelected() bool {
	return C.elm_flipselector_item_selected_get(p.obj) == eTrue
}

func (p *FlipselectorItem) Prev() *FlipselectorItem {
	return wrapFlipselectorItem(C.elm_flipselector_item_prev_get(p.obj))
}

func (p *FlipselectorItem) Next() *FlipselectorItem {
	return wrapFlipselectorItem(C.elm_flipselector_item_next_get(p.obj))
}

//------------------------------------------------------------

type NaviframeItem struct {
	*widgetItemBase
}

func wrapNaviframeItem(o *C.Elm_Widget_Item) *NaviframeItem {
	if o != nil {
		return &NaviframeItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *NaviframeItem) PopTo() {
	C.elm_naviframe_item_pop_to(p.obj)
}

func (p *NaviframeItem) Promote() {
	C.elm_naviframe_item_promote(p.obj)
}

func (p *NaviframeItem) SimplePromote(content Object) {
	C.elm_naviframe_item_simple_promote(p.obj, content.eo())
}

func (p *NaviframeItem) SetStyle(style string) {
	cstyle := C.CString(style)
	defer free(cstyle)
	C.elm_naviframe_item_style_set(p.obj, cstyle)
}

func (p *NaviframeItem) Style() string {
	return C.GoString(C.elm_naviframe_item_style_get(p.obj))
}

func (p *NaviframeItem) SetTitleEnabled(b bool, t bool) {
	C.elm_naviframe_item_title_enabled_set(p.obj, eBool(b), eBool(t))
}

func (p *NaviframeItem) IsTitleEnabled() bool {
	return C.elm_naviframe_item_title_enabled_get(p.obj) == eTrue
}

/*TODO:
void elm_naviframe_item_pop_cb_set(Elm_Object_Item *it, Elm_Naviframe_Item_Pop_Cb func, void *data)
*/

//------------------------------------------------------------

type GengridItem struct {
	*widgetItemBase
}

func wrapGengridItem(o *C.Elm_Widget_Item) *GengridItem {
	if o != nil {
		return &GengridItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *GengridItem) Index() int {
	return int(C.elm_gengrid_item_index_get(p.obj))
}

func (p *GengridItem) SetSelectMode(mode SelectMode) {
	C.elm_gengrid_item_select_mode_set(p.obj, C.Elm_Object_Select_Mode(mode))
}

func (p *GengridItem) SelectMode() SelectMode {
	return SelectMode(C.elm_gengrid_item_select_mode_get(p.obj))
}

func (p *GengridItem) Next() *GengridItem {
	return wrapGengridItem(C.elm_gengrid_item_next_get(p.obj))
}

func (p *GengridItem) Prev() *GengridItem {
	return wrapGengridItem(C.elm_gengrid_item_prev_get(p.obj))
}

func (p *GengridItem) SetSelected(b bool) {
	C.elm_gengrid_item_selected_set(p.obj, eBool(b))
}

func (p *GengridItem) IsSelected() bool {
	return C.elm_gengrid_item_selected_get(p.obj) == eTrue
}

func (p *GengridItem) Show(t GengridItemScrolltoType) {
	C.elm_gengrid_item_show(p.obj, C.Elm_Gengrid_Item_Scrollto_Type(t))
}

func (p *GengridItem) BringIn(t GengridItemScrolltoType) {
	C.elm_gengrid_item_bring_in(p.obj, C.Elm_Gengrid_Item_Scrollto_Type(t))
}

func (p *GengridItem) Update() {
	C.elm_gengrid_item_update(p.obj)
}

func (p *GengridItem) Pos() (uint, uint) {
	var x, y C.uint
	C.elm_gengrid_item_pos_get(p.obj, &x, &y)
	return uint(x), uint(y)
}

/* in 1.19
func (p *GengridItem) SetCustomSize(w, h int) {
	C.elm_gengrid_item_custom_size_set(p.obj, C.Evas_Coord(w), C.Evas_Coord(h))
}

func (p *GengridItem) CustomSize() (int, int) {
	var w, h C.Evas_Coord
	C.elm_gengrid_item_custom_size_get(p.obj, &w, &h)
	return int(w), int(h)
}
*/
/*
TODO:
void elm_gengrid_item_all_contents_unset(Elm_Object_Item *obj, Eina_List **l)
*/

//------------------------------------------------------------

type HoverselItem struct {
	*widgetItemBase
}

func wrapHoverselItem(o *C.Elm_Widget_Item) *HoverselItem {
	if o != nil {
		return &HoverselItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *HoverselItem) SetIcon(file string, group string, t IconType) {
	cfile := C.CString(file)
	cgroup := C.CString(group)
	defer free(cfile, cgroup)
	C.elm_hoversel_item_icon_set(p.obj, cfile, cgroup, C.Elm_Icon_Type(t))
}

func (p *HoverselItem) Icon() (string, string, IconType) {
	var (
		file, group *C.char
		t           C.Elm_Icon_Type
	)
	C.elm_hoversel_item_icon_get(p.obj, &file, &group, &t)
	return C.GoString(file), C.GoString(group), IconType(t)
}

//------------------------------------------------------------

type CtxpopupItem struct {
	*widgetItemBase
}

func wrapCtxpopupItem(o *C.Elm_Widget_Item) *CtxpopupItem {
	if o != nil {
		return &CtxpopupItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *CtxpopupItem) Prev() *CtxpopupItem {
	return wrapCtxpopupItem(C.elm_ctxpopup_item_prev_get(p.obj))
}

func (p *CtxpopupItem) Next() *CtxpopupItem {
	return wrapCtxpopupItem(C.elm_ctxpopup_item_next_get(p.obj))
}

//------------------------------------------------------------

type ColorselectorPaletteItem struct {
	*widgetItemBase
}

func wrapColorselectorPaletteItem(o *C.Elm_Widget_Item) *ColorselectorPaletteItem {
	if o != nil {
		return &ColorselectorPaletteItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *ColorselectorPaletteItem) Color() (int, int, int, int) {
	var r, g, b, a C.int
	C.elm_colorselector_palette_item_color_get(p.obj, &r, &g, &b, &a)
	return int(r), int(g), int(b), int(a)
}

func (p *ColorselectorPaletteItem) SetColor(r, g, b, a int) {
	C.elm_colorselector_palette_item_color_set(p.obj, C.int(r), C.int(g), C.int(b), C.int(a))
}

func (p *ColorselectorPaletteItem) IsSelected() bool {
	return C.elm_colorselector_palette_item_selected_get(p.obj) == eTrue
}

func (p *ColorselectorPaletteItem) SetSelected(b bool) {
	C.elm_colorselector_palette_item_selected_set(p.obj, eBool(b))
}

//------------------------------------------------------------

type PopupItem struct {
	*widgetItemBase
}

func wrapPopupItem(o *C.Elm_Widget_Item) *PopupItem {
	if o != nil {
		return &PopupItem{wrapWidgetItemBase(o)}
	}
	return nil
}

//------------------------------------------------------------

type GenlistItem struct {
	*widgetItemBase
}

func wrapGenlistItem(o *C.Elm_Widget_Item) *GenlistItem {
	if o != nil {
		return &GenlistItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *GenlistItem) Next() *GenlistItem {
	return wrapGenlistItem(C.elm_genlist_item_next_get(p.obj))
}

func (p *GenlistItem) Prev() *GenlistItem {
	return wrapGenlistItem(C.elm_genlist_item_prev_get(p.obj))
}

func (p *GenlistItem) SetSelected(b bool) {
	C.elm_genlist_item_selected_set(p.obj, eBool(b))
}

func (p *GenlistItem) IsSelected() bool {
	return C.elm_genlist_item_selected_get(p.obj) == eTrue
}

func (p *GenlistItem) Show(t GenlistItemScrolltoType) {
	C.elm_genlist_item_show(p.obj, C.Elm_Genlist_Item_Scrollto_Type(t))
}

func (p *GenlistItem) BringIn(t GenlistItemScrolltoType) {
	C.elm_genlist_item_bring_in(p.obj, C.Elm_Genlist_Item_Scrollto_Type(t))
}

func (p *GenlistItem) Update() {
	C.elm_genlist_item_update(p.obj)
}

func (p *GenlistItem) Index() int {
	return int(C.elm_genlist_item_index_get(p.obj))
}

func (p *GenlistItem) Parent() *GenlistItem {
	return wrapGenlistItem(C.elm_genlist_item_parent_get(p.obj))
}

func (p *GenlistItem) ClearSubitems() {
	C.elm_genlist_item_subitems_clear(p.obj)
}

func (p *GenlistItem) SubitemsCount() uint {
	return uint(C.elm_genlist_item_subitems_count(p.obj))
}

func (p *GenlistItem) Subitems() []*GenlistItem {
	return newListIterator(C.elm_genlist_item_subitems_get(p.obj)).GenlistItemSlice()
}

func (p *GenlistItem) SetExpanded(b bool) {
	C.elm_genlist_item_expanded_set(p.obj, eBool(b))
}

func (p *GenlistItem) IsExpanded() bool {
	return C.elm_genlist_item_expanded_get(p.obj) == eTrue
}

func (p *GenlistItem) ExpandedDepth() int {
	return int(C.elm_genlist_item_expanded_depth_get(p.obj))
}

func (p *GenlistItem) Promote() {
	C.elm_genlist_item_promote(p.obj)
}

func (p *GenlistItem) Demote() {
	C.elm_genlist_item_promote(p.obj)
}

func (p *GenlistItem) UpdateFields(parts string, ft GenlistItemFieldType) {
	var cparts *C.char

	if parts != "" {
		cparts = C.CString(parts)
		defer free(cparts)
	}

	C.elm_genlist_item_fields_update(p.obj, cparts, C.Elm_Genlist_Item_Field_Type(ft))
}

func (p *GenlistItem) SetDecorateMode(dit string, b bool) {
	cdit := C.CString(dit)
	defer free(cdit)
	C.elm_genlist_item_decorate_mode_set(p.obj, cdit, eBool(b))
}

func (p *GenlistItem) DecorateMode() string {
	return C.GoString(C.elm_genlist_item_decorate_mode_get(p.obj))
}

func (p *GenlistItem) Type() GenlistItemType {
	return GenlistItemType(C.elm_genlist_item_type_get(p.obj))
}

func (p *GenlistItem) SetFlip(b bool) {
	C.elm_genlist_item_flip_set(p.obj, eBool(b))
}

func (p *GenlistItem) IsFlip() bool {
	return C.elm_genlist_item_flip_get(p.obj) == eTrue
}

func (p *GenlistItem) SetSelectMode(mode SelectMode) {
	C.elm_genlist_item_select_mode_set(p.obj, C.Elm_Object_Select_Mode(mode))
}

func (p *GenlistItem) SelectMode() SelectMode {
	return SelectMode(C.elm_genlist_item_select_mode_get(p.obj))
}

/*
TODO:
void elm_genlist_item_item_class_update(Elm_Object_Item *it, Elm_Genlist_Item_Class *itc)
Elm_Genlist_Item_Class *elm_genlist_item_item_class_get(const Elm_Object_Item *it)
void elm_genlist_item_all_contents_unset(Elm_Object_Item *it, Eina_List **l)
*/

//------------------------------------------------------------

type DiskselectorItem struct {
	*widgetItemBase
}

func wrapDiskselectorItem(o *C.Elm_Widget_Item) *DiskselectorItem {
	if o != nil {
		return &DiskselectorItem{wrapWidgetItemBase(o)}
	}
	return nil
}

func (p *DiskselectorItem) SetSelected(b bool) {
	C.elm_diskselector_item_selected_set(p.obj, eBool(b))
}

func (p *DiskselectorItem) IsSelected() bool {
	return C.elm_diskselector_item_selected_get(p.obj) == eTrue
}

func (p *DiskselectorItem) Prev() *DiskselectorItem {
	return wrapDiskselectorItem(C.elm_diskselector_item_prev_get(p.obj))
}

func (p *DiskselectorItem) Next() *DiskselectorItem {
	return wrapDiskselectorItem(C.elm_diskselector_item_next_get(p.obj))
}

//------------------------------------------------------------

/*
TODO:
Item interface
*/

/*
Item Factories
*/

type MenuItemFactory struct {
	menu        *Menu
	parent      *MenuItem
	icon        string
	label       string
	handler     Handler
	dataHandler DataHandler
	data        interface{}
}

func NewMenuItemFactory(menu *Menu) *MenuItemFactory {
	return &MenuItemFactory{menu: menu}
}

func (p *MenuItemFactory) SetParent(parent *MenuItem) {
	p.parent = parent
}

func (p *MenuItemFactory) SetIcon(icon string) {
	p.icon = icon
}

func (p *MenuItemFactory) SetLabel(lbl string) {
	p.label = lbl
}

func (p *MenuItemFactory) SetHandler(handler Handler) {
	p.handler = handler
}

func (p *MenuItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *MenuItemFactory) Add() *MenuItem {
	var (
		cicon   *C.char
		clabel  *C.char
		cparent *C.Elm_Widget_Item
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
	)
	if p.icon != "" {
		cicon = C.CString(p.icon)
		defer free(cicon)
	}
	if p.label != "" {
		clabel = C.CString(p.label)
		defer free(clabel)
	}
	if p.parent != nil {
		cparent = p.parent.obj
	}

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	mi := C.elm_menu_item_add(p.menu.obj, cparent, cicon, clabel, smartcb, data)
	p.icon = ""
	p.label = ""
	p.handler = nil
	p.parent = nil
	p.dataHandler = nil
	p.data = nil
	return wrapMenuItem(mi)
}

type ToolbarItemFactory struct {
	toolbar     *Toolbar
	icon        string
	label       string
	handler     Handler
	dataHandler DataHandler
	data        interface{}
}

func NewToolbarItemFactory(toolbar *Toolbar) *ToolbarItemFactory {
	return &ToolbarItemFactory{toolbar: toolbar}
}

func (p *ToolbarItemFactory) SetLabel(lbl string) {
	p.label = lbl
}

func (p *ToolbarItemFactory) SetIcon(icon string) {
	p.icon = icon
}

func (p *ToolbarItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *ToolbarItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *ToolbarItemFactory) add(isbefore bool, item *ToolbarItem) *ToolbarItem {
	var (
		cicon   *C.char
		clabel  *C.char
		ti      *C.Elm_Widget_Item
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
	)
	if p.icon != "" {
		cicon = C.CString(p.icon)
		defer free(cicon)
	}
	if p.label != "" {
		clabel = C.CString(p.label)
		defer free(clabel)
	}
	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	if isbefore {
		if item != nil {
			ti = C.elm_toolbar_item_insert_before(p.toolbar.obj, item.obj, cicon, clabel, smartcb, data)
		} else {
			ti = C.elm_toolbar_item_prepend(p.toolbar.obj, cicon, clabel, smartcb, data)
		}
	} else {
		if item != nil {
			ti = C.elm_toolbar_item_insert_after(p.toolbar.obj, item.obj, cicon, clabel, smartcb, data)
		} else {
			ti = C.elm_toolbar_item_append(p.toolbar.obj, cicon, clabel, smartcb, data)
		}
	}
	p.icon = ""
	p.label = ""
	p.handler = nil
	p.dataHandler = nil
	p.data = nil
	return wrapToolbarItem(ti)
}

func (p *ToolbarItemFactory) Append() *ToolbarItem {
	return p.add(false, nil)
}
func (p *ToolbarItemFactory) Prepend() *ToolbarItem {
	return p.add(true, nil)
}

func (p *ToolbarItemFactory) InsertBefore(item *ToolbarItem) *ToolbarItem {
	return p.add(true, item)
}

func (p *ToolbarItemFactory) InsertAfter(item *ToolbarItem) *ToolbarItem {
	return p.add(false, item)
}

type IndexItemFactory struct {
	index       *Index
	handler     Handler
	dataHandler DataHandler
	data        interface{}
}

func NewIndexItemFactory(index *Index) *IndexItemFactory {
	return &IndexItemFactory{index: index}
}

func (p *IndexItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *IndexItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *IndexItemFactory) add(letter string, isbefore bool, item *IndexItem) *IndexItem {
	var (
		ii      *C.Elm_Widget_Item
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
	)

	cletter := C.CString(letter)
	defer free(cletter)

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	if isbefore {
		if item != nil {
			ii = C.elm_index_item_insert_before(p.index.obj, item.obj, cletter, smartcb, data)
		} else {
			ii = C.elm_index_item_prepend(p.index.obj, cletter, smartcb, data)
		}
	} else {
		if item != nil {
			ii = C.elm_index_item_insert_after(p.index.obj, item.obj, cletter, smartcb, data)
		} else {
			ii = C.elm_index_item_append(p.index.obj, cletter, smartcb, data)
		}
	}
	p.handler = nil
	p.dataHandler = nil
	p.data = nil
	return wrapIndexItem(ii)
}

func (p *IndexItemFactory) Append(letter string) *IndexItem {
	return p.add(letter, false, nil)
}

func (p *IndexItemFactory) Prepend(letter string) *IndexItem {
	return p.add(letter, true, nil)
}

func (p *IndexItemFactory) InsertAfter(item *IndexItem, letter string) *IndexItem {
	return p.add(letter, false, item)
}

func (p *IndexItemFactory) InsertBefore(item *IndexItem, letter string) *IndexItem {
	return p.add(letter, true, item)
}

/*
Elm_Widget_Item * 	elm_index_item_sorted_insert (Elm_Index *obj, const char *letter, Evas_Smart_Cb func, const void *data, Eina_Compare_Cb cmp_func, Eina_Compare_Cb cmp_data_func)
*/

type ListItemFactory struct {
	list        *List
	righticon   *Icon
	lefticon    *Icon
	label       string
	handler     Handler
	dataHandler DataHandler
	data        interface{}
}

func NewListItemFactory(list *List) *ListItemFactory {
	return &ListItemFactory{list: list}
}

func (p *ListItemFactory) SetLabel(label string) {
	p.label = label
}

func (p *ListItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *ListItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *ListItemFactory) SetRightIcon(icon *Icon) {
	p.righticon = icon
}

func (p *ListItemFactory) SetLeftIcon(icon *Icon) {
	p.lefticon = icon
}

func (p *ListItemFactory) add(isbefore bool, item *ListItem) *ListItem {
	var (
		li      *C.Elm_Widget_Item
		lic     *C.Eo
		ric     *C.Eo
		clabel  *C.char
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
	)

	if p.label != "" {
		clabel = C.CString(p.label)
		defer free(clabel)
	}

	if p.righticon != nil {
		ric = p.righticon.obj
	}

	if p.lefticon != nil {
		lic = p.lefticon.obj
	}

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	if isbefore {
		if item != nil {
			li = C.elm_list_item_insert_before(p.list.obj, item.obj, clabel, lic, ric, smartcb, data)
		} else {
			li = C.elm_list_item_prepend(p.list.obj, clabel, lic, ric, smartcb, data)
		}
	} else {
		if item != nil {
			li = C.elm_list_item_insert_after(p.list.obj, item.obj, clabel, lic, ric, smartcb, data)
		} else {
			li = C.elm_list_item_append(p.list.obj, clabel, lic, ric, smartcb, data)
		}
	}

	p.label = ""
	p.righticon = nil
	p.lefticon = nil
	p.handler = nil
	p.dataHandler = nil
	p.data = nil

	return wrapListItem(li)
}

func (p *ListItemFactory) Append() *ListItem {
	return p.add(false, nil)
}

func (p *ListItemFactory) Prepend() *ListItem {
	return p.add(true, nil)
}

func (p *ListItemFactory) InsertBefore(item *ListItem) *ListItem {
	return p.add(true, item)
}

func (p *ListItemFactory) InsertAfter(item *ListItem) *ListItem {
	return p.add(false, item)
}

/*
TODO
Elm_Widget_Item * 	elm_list_item_sorted_insert (Elm_List *obj, const char *label, Efl_Canvas_Object *icon, Efl_Canvas_Object *end, Evas_Smart_Cb func, const void *data, Eina_Compare_Cb cmp_func)

*/

/*
TODO:
Elm_Widget_Item * 	elm_slideshow_item_sorted_insert (Elm_Slideshow *obj, const Elm_Slideshow_Item_Class *itc, const void *data, Eina_Compare_Cb func)
*/

type SlideshowItemModel interface {
	Content(interface{}, *Slideshow) Object
}

//export go_slideshow_item_get_func
func go_slideshow_item_get_func(id C.int, obj *C.Eo) *C.Eo {
	if x := registry.Lookup(int(id)); x != nil {
		if ssid, ok := x.(*slideShowItemFuncData); ok {
			return ssid.get_func(ssid.data, ssid.slideshow).eo()
		}
	}
	return nil
}

/* FIXME: every time a item gets and sets, it freed when item is deleted this is not required
//export go_slideshow_item_del_func
func go_slideshow_item_del_func(id C.int) {
	remove(int(id))
	fmt.Println("remove called")
}
*/

type slideShowItemFuncData struct {
	get_func  func(data interface{}, slideshow *Slideshow) Object
	data      interface{}
	slideshow *Slideshow
}

type SlideshowItemFactory struct {
	get_func  func(data interface{}, slideshow *Slideshow) Object
	slideshow *Slideshow
	ssic      C.Elm_Slideshow_Item_Class
}

func NewSlideshowItemFactory(slideshow *Slideshow, fn func(data interface{}, slideshow *Slideshow) Object) *SlideshowItemFactory {
	fact := &SlideshowItemFactory{
		slideshow: slideshow,
		get_func:  fn,
		ssic:      C.cgo_slideshow_item_class_get(),
	}

	return fact
}

func (p *SlideshowItemFactory) Add(data interface{}) *SlideshowItem {
	ifd := &slideShowItemFuncData{
		slideshow: p.slideshow,
		data:      data,
		get_func:  p.get_func,
	}
	id := registry.Register(ifd)
	return wrapSlideshowItem(C.elm_slideshow_item_add(p.slideshow.obj, &p.ssic, unsafe.Pointer(&id)))
}

type MultibuttonentryItemFactory struct {
	mbentry     *Multibuttonentry
	handler     Handler
	dataHandler DataHandler
	data        interface{}
}

func NewMultibuttonentryItemFactory(mbentry *Multibuttonentry) *MultibuttonentryItemFactory {
	return &MultibuttonentryItemFactory{mbentry: mbentry}
}

func (p *MultibuttonentryItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *MultibuttonentryItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *MultibuttonentryItemFactory) add(label string, isbefore bool, item *MultibuttonentryItem) *MultibuttonentryItem {
	var (
		mbei    *C.Elm_Widget_Item
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
	)

	clabel := C.CString(label)
	defer free(clabel)

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	if isbefore {
		if item != nil {
			mbei = C.elm_multibuttonentry_item_insert_before(p.mbentry.obj, item.obj, clabel, smartcb, data)
		} else {
			mbei = C.elm_multibuttonentry_item_prepend(p.mbentry.obj, clabel, smartcb, data)
		}
	} else {
		if item != nil {
			mbei = C.elm_multibuttonentry_item_insert_after(p.mbentry.obj, item.obj, clabel, smartcb, data)
		} else {
			mbei = C.elm_multibuttonentry_item_append(p.mbentry.obj, clabel, smartcb, data)
		}
	}

	p.handler = nil
	p.dataHandler = nil
	p.data = nil

	return wrapMultibuttonentryItem(mbei)
}

func (p *MultibuttonentryItemFactory) Append(label string) *MultibuttonentryItem {
	return p.add(label, false, nil)
}

func (p *MultibuttonentryItemFactory) Prepend(label string) *MultibuttonentryItem {
	return p.add(label, true, nil)
}

func (p *MultibuttonentryItemFactory) InsertBefore(item *MultibuttonentryItem, label string) *MultibuttonentryItem {
	return p.add(label, true, item)
}

func (p *MultibuttonentryItemFactory) InsertAfter(item *MultibuttonentryItem, label string) *MultibuttonentryItem {
	return p.add(label, false, item)
}

type SegmentControlItemFactory struct {
	sc    *SegmentControl
	icon  *Icon
	label string
}

func NewSegmentControlItemFactory(sc *SegmentControl) *SegmentControlItemFactory {
	return &SegmentControlItemFactory{sc: sc}
}

func (p *SegmentControlItemFactory) SetLabel(label string) {
	p.label = label
}

func (p *SegmentControlItemFactory) SetIcon(icon *Icon) {
	p.icon = icon
}

func (p *SegmentControlItemFactory) add(isinsert bool, index int) *SegmentControlItem {
	var (
		clabel *C.char
		ieo    *C.Eo
		sci    *C.Elm_Widget_Item
	)

	if p.label != "" {
		clabel = C.CString(p.label)
		defer free(clabel)
	}

	if p.icon != nil {
		ieo = p.icon.obj
	}

	if isinsert {
		sci = C.elm_segment_control_item_insert_at(p.sc.obj, ieo, clabel, C.int(index))
	} else {
		sci = C.elm_segment_control_item_add(p.sc.obj, ieo, clabel)
	}

	p.icon = nil
	p.label = ""

	return wrapSegmentControlItem(sci)
}

func (p *SegmentControlItemFactory) Add() *SegmentControlItem {
	return p.add(false, 0)
}

func (p *SegmentControlItemFactory) InsertAt(index int) *SegmentControlItem {
	return p.add(true, index)
}

/*
TODO:
Elm_Widget_Item * 	elm_flipselector_item_prepend (Elm_Flipselector *obj, const char *label, Evas_Smart_Cb func, void *data)
Elm_Widget_Item * 	elm_flipselector_item_append (Elm_Flipselector *obj, const char *label, Evas_Smart_Cb func, const void *data)
*/

type FlipselectorItemFactory struct {
	flipselector *Flipselector
	handler      Handler
	dataHandler  DataHandler
	data         interface{}
}

func NewFlipselectorItemFactory(fs *Flipselector) *FlipselectorItemFactory {
	return &FlipselectorItemFactory{flipselector: fs}
}

func (p *FlipselectorItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *FlipselectorItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *FlipselectorItemFactory) add(pre bool, label string) *FlipselectorItem {
	var (
		fsi     *C.Elm_Widget_Item
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
	)

	clabel := C.CString(label)
	defer free(clabel)

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	if pre {
		fsi = C.elm_flipselector_item_prepend(p.flipselector.obj, clabel, smartcb, data)
	} else {
		fsi = C.elm_flipselector_item_append(p.flipselector.obj, clabel, smartcb, data)
	}

	p.handler = nil
	p.dataHandler = nil
	p.data = nil

	return wrapFlipselectorItem(fsi)
}

func (p *FlipselectorItemFactory) Append(label string) *FlipselectorItem {
	return p.add(false, label)
}

func (p *FlipselectorItemFactory) Prepend(label string) *FlipselectorItem {
	return p.add(true, label)
}

//Elm_Widget_Item * 	elm_hoversel_item_add (Elm_Hoversel *obj, const char *label, const char *icon_file, Elm_Icon_Type icon_type, Evas_Smart_Cb func, const void *data)

type HoverselItemFactory struct {
	hoversel    *Hoversel
	label       string
	iconfile    string
	icontype    IconType
	handler     Handler
	dataHandler DataHandler
	data        interface{}
}

func NewHoverselItemFactory(hoversel *Hoversel) *HoverselItemFactory {
	return &HoverselItemFactory{
		hoversel: hoversel,
		icontype: IconTypeStandard,
	}
}

func (p *HoverselItemFactory) SetLabel(label string) {
	p.label = label
}

func (p *HoverselItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *HoverselItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *HoverselItemFactory) SetIcon(file string, typ ...IconType) {
	p.iconfile = file
	if len(typ) > 0 {
		p.icontype = typ[0]
	}
}

func (p *HoverselItemFactory) Add() *HoverselItem {
	var (
		clabel    *C.char
		ciconfile *C.char
		smartcb   C.Evas_Smart_Cb
		data      unsafe.Pointer
	)

	if p.label != "" {
		clabel = C.CString(p.label)
		defer free(clabel)
	}

	if p.iconfile != "" {
		ciconfile = C.CString(p.iconfile)
		defer free(ciconfile)
	}

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	it := C.elm_hoversel_item_add(p.hoversel.obj, clabel, ciconfile, C.Elm_Icon_Type(p.icontype), smartcb, data)
	p.label = ""
	p.iconfile = ""
	p.handler = nil
	p.dataHandler = nil
	p.data = nil

	return wrapHoverselItem(it)
}

/*
TODO:
Elm_Widget_Item * 	elm_ctxpopup_item_append (Elm_Ctxpopup *obj, const char *label, Efl_Canvas_Object *icon, Evas_Smart_Cb func, const void *data)
Elm_Widget_Item * 	elm_ctxpopup_item_prepend (Elm_Ctxpopup *obj, const char *label, Efl_Canvas_Object *icon, Evas_Smart_Cb func, const void *data)
*/

type CtxpopupItemFactory struct {
	ctxpopup    *Ctxpopup
	label       string
	icon        *Icon
	handler     Handler
	dataHandler DataHandler
	data        interface{}
}

func NewCtxpopupItemFactory(ctxpopup *Ctxpopup) *CtxpopupItemFactory {
	return &CtxpopupItemFactory{ctxpopup: ctxpopup}
}

func (p *CtxpopupItemFactory) SetLabel(label string) {
	p.label = label
}

func (p *CtxpopupItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *CtxpopupItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *CtxpopupItemFactory) SetIcon(icon *Icon) {
	p.icon = icon
}

func (p *CtxpopupItemFactory) add(ispre bool) *CtxpopupItem {
	var (
		clabel  *C.char
		cpi     *C.Elm_Widget_Item
		ic      *C.Eo
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
	)

	if p.label != "" {
		clabel = C.CString(p.label)
		defer free(clabel)
	}

	if p.icon != nil {
		ic = p.icon.obj
	}

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	if ispre {
		cpi = C.elm_ctxpopup_item_prepend(p.ctxpopup.obj, clabel, ic, smartcb, data)
	} else {
		cpi = C.elm_ctxpopup_item_append(p.ctxpopup.obj, clabel, ic, smartcb, data)
	}

	p.label = ""
	p.icon = nil
	p.handler = nil
	p.dataHandler = nil
	p.data = nil

	return wrapCtxpopupItem(cpi)
}

func (p *CtxpopupItemFactory) Append() *CtxpopupItem {
	return p.add(false)
}

func (p *CtxpopupItemFactory) Prepend() *CtxpopupItem {
	return p.add(true)
}

type NaviframeItemFactory struct {
	naviframe  *Naviframe
	title      string
	prevbutton *Button
	nextbutton *Button
	style      string
}

func NewNaviframeItemFactory(naviframe *Naviframe) *NaviframeItemFactory {
	return &NaviframeItemFactory{naviframe: naviframe}
}

func (p *NaviframeItemFactory) SetTitle(title string) {
	p.title = title
}

func (p *NaviframeItemFactory) SetPrevButton(btn *Button) {
	p.prevbutton = btn
}

func (p *NaviframeItemFactory) SetNextButton(btn *Button) {
	p.nextbutton = btn
}

func (p *NaviframeItemFactory) SetStyle(style string) {
	p.style = style
}

func (p *NaviframeItemFactory) add(isbefore bool, item *NaviframeItem, content Object) *NaviframeItem {
	var (
		ctitle  *C.char
		cstyle  *C.char
		prevbtn *C.Eo
		nextbtn *C.Eo
		nit     *C.Elm_Widget_Item
	)

	if p.title != "" {
		ctitle = C.CString(p.title)
		defer free(ctitle)
	}

	if p.style != "" {
		cstyle = C.CString(p.style)
		defer free(cstyle)
	}

	if p.prevbutton != nil {
		prevbtn = p.prevbutton.obj
	}

	if p.nextbutton != nil {
		nextbtn = p.nextbutton.obj
	}

	if item != nil {
		if isbefore {
			nit = C.elm_naviframe_item_insert_before(p.naviframe.obj, item.obj, ctitle, prevbtn, nextbtn, content.eo(), cstyle)
		} else {
			nit = C.elm_naviframe_item_insert_after(p.naviframe.obj, item.obj, ctitle, prevbtn, nextbtn, content.eo(), cstyle)
		}
	} else {
		nit = C.elm_naviframe_item_push(p.naviframe.obj, ctitle, prevbtn, nextbtn, content.eo(), cstyle)
	}

	return wrapNaviframeItem(nit)
}

func (p *NaviframeItemFactory) Push(content Object) *NaviframeItem {
	return p.add(false, nil, content)
}

func (p *NaviframeItemFactory) InsertBefore(item *NaviframeItem, content Object) *NaviframeItem {
	return p.add(true, item, content)
}

func (p *NaviframeItemFactory) InsertAfter(item *NaviframeItem, content Object) *NaviframeItem {
	return p.add(false, item, content)
}

/*
TODO:
Elm_Widget_Item * 	elm_naviframe_item_insert_before (Elm_Naviframe *obj, Elm_Widget_Item *before, const char *title_label, Evas_Object *prev_btn, Evas_Object *next_btn, Evas_Object *content, const char *item_style)
Elm_Widget_Item * 	elm_naviframe_item_push (Elm_Naviframe *obj, const char *title_label, Evas_Object *prev_btn, Evas_Object *next_btn, Evas_Object *content, const char *item_style)
Elm_Widget_Item * 	elm_naviframe_item_insert_after (Elm_Naviframe *obj, Elm_Widget_Item *after, const char *title_label, Evas_Object *prev_btn, Evas_Object *next_btn, Evas_Object *content, const char *item_style)
*/

//Elm_Object_Item         *elm_popup_item_append(Evas_Object *obj, const char *label, Evas_Object *icon, Evas_Smart_Cb func, void *data)

type PopupItemFactory struct {
	popup       *Popup
	label       string
	icon        *Icon
	handler     Handler
	dataHandler DataHandler
	data        interface{}
}

func NewPopupItemFactory(popup *Popup) *PopupItemFactory {
	return &PopupItemFactory{popup: popup}
}

func (p *PopupItemFactory) SetLabel(label string) {
	p.label = label
}

func (p *PopupItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *PopupItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *PopupItemFactory) SetIcon(icon *Icon) {
	p.icon = icon
}

func (p *PopupItemFactory) Append() *PopupItem {
	var (
		clabel  *C.char
		cicon   *C.Eo
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
	)

	if p.label != "" {
		clabel = C.CString(p.label)
		defer free(clabel)
	}

	if p.icon != nil {
		cicon = p.icon.obj
	}

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	p.label = ""
	p.icon = nil
	p.handler = nil
	p.dataHandler = nil
	p.data = nil

	return wrapPopupItem(C.elm_popup_item_append(p.popup.obj, clabel, cicon, smartcb, data))
}

//Elm_Object_Item         *elm_diskselector_item_append(Evas_Object *obj, const char *label, Evas_Object *icon, Evas_Smart_Cb func, void *data)

type DiskselectorItemFactory struct {
	diskselector *Diskselector
	label        string
	icon         *Icon
	handler      Handler
	dataHandler  DataHandler
	data         interface{}
}

func NewDiskselectorItemFactory(diskselector *Diskselector) *DiskselectorItemFactory {
	return &DiskselectorItemFactory{diskselector: diskselector}
}

func (p *DiskselectorItemFactory) SetLabel(label string) {
	p.label = label
}

func (p *DiskselectorItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *DiskselectorItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *DiskselectorItemFactory) SetIcon(icon *Icon) {
	p.icon = icon
}

func (p *DiskselectorItemFactory) Append() *DiskselectorItem {
	var (
		clabel  *C.char
		ic      *C.Eo
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
	)

	if p.label != "" {
		clabel = C.CString(p.label)
		defer free(clabel)
	}

	if p.icon != nil {
		ic = p.icon.obj
	}

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	p.handler = nil
	p.dataHandler = nil
	p.data = nil

	return wrapDiskselectorItem(C.elm_diskselector_item_append(p.diskselector.obj, clabel, ic, smartcb, data))
}

//TODO: elm_gengrid_item_class_free ???

type GengridModel interface {
	Text(string) string
	Content(Object, string) Object
	State(string) bool
}

type AbstractGengridModel struct{}

var _ GengridModel = AbstractGengridModel{}

func (AbstractGengridModel) Text(str string) string {
	return ""
}

func (AbstractGengridModel) Content(parent Object, str string) Object {
	return nil
}

func (AbstractGengridModel) State(str string) bool {
	return false
}

//TODO var gengridModelmap = make(map[int]GengridModel)

//export go_elm_gengrid_item_class_text_get
func go_elm_gengrid_item_class_text_get(id C.int, obj *C.Eo, part *C.char) *C.char {
	fmt.Println("gengrid text_get called")
	if x := registry.Lookup(int(id)); x != nil {
		if model, ok := x.(GengridModel); ok {
			mt := model.Text(C.GoString(part))
			//fmt.Printf("model text: %s\n",mt)
			txt := C.CString(mt)
			// This function must return a strdup'()ed string, as the caller will free() it when done.
			//defer free(txt) => Segfault FIXME is efl freed ???
			return txt
		}
	}
	//fmt.Println("gengrid text_get return nil")
	return nil
}

//export go_elm_gengrid_item_class_content_get
func go_elm_gengrid_item_class_content_get(id C.int, obj *C.Eo, part *C.char) *C.Eo {
	fmt.Println("gengrid content_get called")
	if x := registry.Lookup(int(id)); x != nil {
		if model, ok := x.(GengridModel); ok {
			c := model.Content(wrapWidgetBase(obj), C.GoString(part))
			if c != nil {
				return c.eo()
			}
		}
	}

	//fmt.Println("gengrid content_get return nil")
	return nil
}

//export go_elm_gengrid_item_class_state_get
func go_elm_gengrid_item_class_state_get(id C.int, obj *C.Eo, part *C.char) C.Eina_Bool {
	fmt.Println("gengrid state_get called")
	if x := registry.Lookup(int(id)); x != nil {
		if model, ok := x.(GengridModel); ok {
			return eBool(model.State(C.GoString(part)))
		}
	}
	return eFalse
}

/*
//export go_elm_gengrid_item_class_del
func go_elm_gengrid_item_class_del(id C.int, obj *C.Eo) {
	fmt.Println("gengrid del called")
	registry.Delete(int(id))
}
*/

type GengridItemFactory struct {
	gengrid     *Gengrid
	cls         *C.Elm_Gengrid_Item_Class
	handler     Handler
	dataHandler DataHandler
	data        interface{}
}

func NewGengridItemFactory(gengrid *Gengrid) *GengridItemFactory {
	f := &GengridItemFactory{
		gengrid: gengrid,
		cls:     C.cgo_elm_gengrid_item_class_new(),
	}
	return f
}

func (p *GengridItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *GengridItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *GengridItemFactory) add(model GengridModel, isbefore bool, item *GengridItem) *GengridItem {
	var (
		it      *C.Elm_Widget_Item
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
		id      int
	)

	id = registry.Register(model)

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	if isbefore {
		if item != nil {
			it = C.elm_gengrid_item_insert_before(p.gengrid.obj, p.cls, unsafe.Pointer(&id), item.obj, smartcb, data)
		} else {
			it = C.elm_gengrid_item_prepend(p.gengrid.obj, p.cls, unsafe.Pointer(&id), smartcb, data)
		}
	} else {
		if item != nil {
			it = C.elm_gengrid_item_insert_after(p.gengrid.obj, p.cls, unsafe.Pointer(&id), item.obj, smartcb, data)
		} else {
			it = C.elm_gengrid_item_append(p.gengrid.obj, p.cls, unsafe.Pointer(&id), smartcb, data)
		}
	}

	p.handler = nil
	p.dataHandler = nil
	p.data = nil

	return wrapGengridItem(it)
}

func (p *GengridItemFactory) Append(model GengridModel) *GengridItem {
	return p.add(model, false, nil)
}

func (p *GengridItemFactory) Prepend(model GengridModel) *GengridItem {
	return p.add(model, true, nil)
}

func (p *GengridItemFactory) InsertBefore(model GengridModel, before *GengridItem) *GengridItem {
	return p.add(model, true, before)
}
func (p *GengridItemFactory) InsertAfter(model GengridModel, after *GengridItem) *GengridItem {
	return p.add(model, false, after)
}

/*
Elm_Widget_Item *elm_gengrid_item_sorted_insert (Elm_Gengrid *obj, const Elm_Gengrid_Item_Class *itc, const void *data, Eina_Compare_Cb comp, Evas_Smart_Cb func, const void *func_data)
*/

//-----------------------------------------------------------------------------------

//TODO var gengridModelmap = make(map[int]GengridModel)

//export go_elm_genlist_item_class_text_get
func go_elm_genlist_item_class_text_get(id C.int, obj *C.Eo, part *C.char) *C.char {
	fmt.Println("genlist text_get called")
	if x := registry.Lookup(int(id)); x != nil {
		if model, ok := x.(GenlistModel); ok {
			mt := model.Text(C.GoString(part))
			txt := C.CString(mt)
			//defer free(txt) => Segfault FIXME is efl freed ???
			return txt
		}
	}
	fmt.Println("genlist text_get return nil")
	return nil
}

//export go_elm_genlist_item_class_content_get
func go_elm_genlist_item_class_content_get(id C.int, obj *C.Eo, part *C.char) *C.Eo {
	fmt.Println("genlist content_get called")
	if x := registry.Lookup(int(id)); x != nil {
		if model, ok := x.(GenlistModel); ok {
			//FIXME: is this widget or object ?
			c := model.Content(wrapWidgetBase(obj), C.GoString(part))
			if c != nil {
				return c.eo()
			}
		}
	}
	fmt.Println("genlist content_get return nil")
	return nil
}

//export go_elm_genlist_item_class_reusable_content_get
func go_elm_genlist_item_class_reusable_content_get(id C.int, obj *C.Eo, part *C.char, old *C.Eo) *C.Eo {
	//TODO
	return nil
}

//export go_elm_genlist_item_class_state_get
func go_elm_genlist_item_class_state_get(id C.int, obj *C.Eo, part *C.char) C.Eina_Bool {
	fmt.Println("genlist state_get called")
	if x := registry.Lookup(int(id)); x != nil {
		if model, ok := x.(GenlistModel); ok {
			return eBool(model.State(C.GoString(part)))
		}
	}
	return eFalse
}

//export go_elm_genlist_item_class_filter_get
func go_elm_genlist_item_class_filter_get(id C.int, obj *C.Eo, key unsafe.Pointer) C.Eina_Bool {
	//TODO
	return eFalse
}

/*
//export go_elm_genlist_item_class_del
func go_elm_genlist_item_class_del(id C.int, obj *C.Eo) {
	//fmt.Printf("genlist del called: id = %d\n",int(id))
	//registry.Delete(int(id))
	// TODO: id is deleted with cgo_elm_object_item_del_cb_set, do we need this cb
}
*/

type GenlistModel interface {
	GengridModel
	ReusableContent(string, Object) Object
	Filter(interface{}) bool
}

// override which is needed
type AbstractGenlistModel struct{}

var _ GengridModel = AbstractGengridModel{}

func (AbstractGenlistModel) Text(str string) string {
	return ""
}

func (AbstractGenlistModel) Content(parent Object, str string) Object {
	return nil
}

func (AbstractGenlistModel) State(str string) bool {
	return false
}

func (AbstractGenlistModel) ReusableContent(str string, o Object) Object {
	return nil
}

func (AbstractGenlistModel) Filter(key interface{}) bool {
	return false
}

type GenlistItemFactory struct {
	genlist     *Genlist
	cls         *C.Elm_Genlist_Item_Class
	handler     Handler
	dataHandler DataHandler
	data        interface{}
	itemType    GenlistItemType
	parent      *GenlistItem
}

func NewGenlistItemFactory(genlist *Genlist) *GenlistItemFactory {
	f := &GenlistItemFactory{
		genlist:  genlist,
		cls:      C.cgo_elm_genlist_item_class_new(),
		itemType: GenlistItemTypeNone,
	}
	return f
}

func (p *GenlistItemFactory) SetItemType(t GenlistItemType) {
	p.itemType = t
}

func (p *GenlistItemFactory) SetParent(parent *GenlistItem) {
	p.parent = parent
}

func (p *GenlistItemFactory) SetHandler(h Handler) {
	p.handler = h
}

func (p *GenlistItemFactory) SetDataHandler(h DataHandler, data interface{}) {
	p.dataHandler = h
	p.data = data
}

func (p *GenlistItemFactory) add(model GenlistModel, isbefore bool, item *GenlistItem) *GenlistItem {
	var (
		it      *C.Elm_Widget_Item
		cparent *C.Elm_Widget_Item
		smartcb C.Evas_Smart_Cb
		data    unsafe.Pointer
		id      int
	)

	id = registry.Register(model)

	if p.parent != nil {
		cparent = p.parent.obj
	}

	if p.handler != nil {
		handler_id := registerHandler(p.handler)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	} else if p.dataHandler != nil {
		handler_id := registerDataHandler(p.dataHandler, p.data)
		smartcb = (C.Evas_Smart_Cb)(C.cgo_smart_callback_func)
		data = unsafe.Pointer(&handler_id)
	}

	if isbefore {
		if item != nil {
			it = C.elm_genlist_item_insert_before(p.genlist.obj, p.cls, unsafe.Pointer(&id), cparent, item.obj, C.Elm_Genlist_Item_Type(p.itemType), smartcb, data)
		} else {
			it = C.elm_genlist_item_prepend(p.genlist.obj, p.cls, unsafe.Pointer(&id), cparent, C.Elm_Genlist_Item_Type(p.itemType), smartcb, data)
		}
	} else {
		if item != nil {
			it = C.elm_genlist_item_insert_after(p.genlist.obj, p.cls, unsafe.Pointer(&id), cparent, item.obj, C.Elm_Genlist_Item_Type(p.itemType), smartcb, data)
		} else {
			it = C.elm_genlist_item_append(p.genlist.obj, p.cls, unsafe.Pointer(&id), cparent, C.Elm_Genlist_Item_Type(p.itemType), smartcb, data)
		}
	}

	p.handler = nil
	p.dataHandler = nil
	p.data = nil
	p.parent = nil
	p.itemType = GenlistItemTypeNone

	return wrapGenlistItem(it)
}

func (p *GenlistItemFactory) Append(model GenlistModel) *GenlistItem {
	return p.add(model, false, nil)
}

func (p *GenlistItemFactory) Prepend(model GenlistModel) *GenlistItem {
	return p.add(model, true, nil)
}

func (p *GenlistItemFactory) InsertBefore(model GenlistModel, before *GenlistItem) *GenlistItem {
	return p.add(model, true, before)
}
func (p *GenlistItemFactory) InsertAfter(model GenlistModel, after *GenlistItem) *GenlistItem {
	return p.add(model, false, after)
}

/*
Elm_Widget_Item * 	elm_genlist_item_sorted_insert (Elm_Genlist *obj, const Elm_Genlist_Item_Class *itc, const void *data, Elm_Widget_Item *parent, Elm_Genlist_Item_Type type, Eina_Compare_Cb comp, Evas_Smart_Cb func, const void *func_data)

void 	elm_genlist_item_class_free (Elm_Genlist_Item_Class *itc)
*/

type EntryContextMenuItem struct {
	obj *C.Elm_Entry_Context_Menu_Item
}

func wrapEntryContextMenuItem(o *C.Elm_Entry_Context_Menu_Item) *EntryContextMenuItem {
	if o != nil {
		return &EntryContextMenuItem{o}
	}
	return nil
}

func (p *EntryContextMenuItem) Label() string {
	return C.GoString(C.elm_entry_context_menu_item_label_get(p.obj))
}

func (p *EntryContextMenuItem) Icon() (string, string, IconType) {
	var (
		ifile, igroup *C.char
		it            C.Elm_Icon_Type
	)
	C.elm_entry_context_menu_item_icon_get(p.obj, &ifile, &igroup, &it)
	return C.GoString(ifile), C.GoString(igroup), IconType(it)
}

type EntryFilterLimitSize struct {
	obj C.Elm_Entry_Filter_Limit_Size
}

func NewEntryFilterLimitSize() *EntryFilterLimitSize {
	return &EntryFilterLimitSize{}
}

func (p *EntryFilterLimitSize) SetMaxCharCount(count int) {
	p.obj.max_char_count = C.int(count)
}

func (p *EntryFilterLimitSize) MaxCharCount() int {
	return int(p.obj.max_char_count)
}

func (p *EntryFilterLimitSize) SetMaxByteCount(count int) {
	p.obj.max_byte_count = C.int(count)
}

func (p *EntryFilterLimitSize) MaxByteCount() int {
	return int(p.obj.max_byte_count)
}

type EntryFilterAcceptSet struct {
	obj C.Elm_Entry_Filter_Accept_Set
}

func NewEntryFilterAcceptSet() *EntryFilterAcceptSet {
	return &EntryFilterAcceptSet{}
}

func (p *EntryFilterAcceptSet) SetAccepted(str string) {
	ca := C.CString(str)
	defer free(ca) // test this is legal ??
	p.obj.accepted = ca
}

func (p *EntryFilterAcceptSet) Accepted() string {
	return C.GoString(p.obj.accepted)
}

func (p *EntryFilterAcceptSet) SetRejected(str string) {
	cr := C.CString(str)
	defer free(cr) // test this is legal
	p.obj.rejected = cr
}

func (p *EntryFilterAcceptSet) Rejected() string {
	return C.GoString(p.obj.rejected)
}

type EntryAnchorInfo struct {
	obj *C.Elm_Entry_Anchor_Info
}

func wrapEntryAnchorInfo(o *C.Elm_Entry_Anchor_Info) *EntryAnchorInfo {
	if o != nil {
		return &EntryAnchorInfo{o}
	}
	return nil
}

func (p *EntryAnchorInfo) Name() string {
	return C.GoString(p.obj.name)
}

func (p *EntryAnchorInfo) Button() int {
	return int(p.obj.button)
}

func (p *EntryAnchorInfo) Geometry() (int, int, int, int) {
	return int(p.obj.x), int(p.obj.y), int(p.obj.w), int(p.obj.h)
}

type EntryAnchorHoverInfo struct {
	obj *C.Elm_Entry_Anchor_Hover_Info
}

func wrapEntryAnchorHoverInfo(o *C.Elm_Entry_Anchor_Hover_Info) *EntryAnchorHoverInfo {
	if o != nil {
		return &EntryAnchorHoverInfo{o}
	}
	return nil
}

func (p *EntryAnchorHoverInfo) AnchorInfo() *EntryAnchorInfo {
	return wrapEntryAnchorInfo(p.obj.anchor_info)
}

func (p *EntryAnchorHoverInfo) Hover() *Hover {
	return wrapHover(p.obj.hover)
}

/*
TODO: is it necessary
hover_parent field
*/
