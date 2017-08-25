package ui

/*
#include "bridge.h"
*/
import "C"

import (
	//	"crypto/rand"
	//	"io"
	"fmt"
	//	"strconv"
	"sync"
	"time"
	"unsafe"
)

const (
	eTrue  = C.EINA_TRUE
	eFalse = C.EINA_FALSE
)

func eBool(b bool) C.Eina_Bool {
	if b {
		return eTrue
	}
	return eFalse
}

func goBool(b C.Eina_Bool) bool {
	return b != eFalse
}

func free(str ...interface{}) {
	for _, s := range str {
		switch x := s.(type) {
		case *C.char:
			C.free(unsafe.Pointer(x))
		case []*C.char:
			for _, cp := range x {
				C.free(unsafe.Pointer(cp))
			}
		default:
			fmt.Printf("utils.go free(): Unknown type: %T\n", x)
		}

	}
}

func goString(str *C.Eina_Stringshare, f ...bool) string {
	if len(f) > 0 {
		if f[0] == true {
			defer C.eina_stringshare_del(str)
		}
	}
	return C.GoString((*C.char)(str))
}

//NOTE: in windows wchar_t is uint16, other platforms use int32
/*
func wcharToRune(wc C.wchar_t) (rune, bool) {
	s := strconv.FormatInt(int64(wc), 10)
	if s == "" {
		return '\000', false
	}
	return rune(s[0]), true
}
*/

func toCTime(tm time.Time) *C.struct_tm {
	var ct C.time_t = C.time_t(tm.Unix())
	return C.localtime(&ct)
}

func toGoTime(ctm *C.struct_tm) time.Time {
	sec := int64(C.mktime(ctm))
	return time.Unix(sec, 0)
}

func CArrayToSlice(array **C.char, l int) []string {
	lst := C.cgo_carray_list_new(array, C.int(l))
	if lst != nil {
		return newListIterator(lst).StringSlice()
	}
	return nil
}

/*TODO: is leaking? test this */
func stringSliceToEList(sl []string) *C.Eina_List {
	var lst *C.Eina_List = nil
	for _, s := range sl {
		cstr := C.eina_stringshare_add(C.CString(s))
		lst = C.eina_list_append(lst, unsafe.Pointer(cstr))
	}
	return lst
}

type eIterator struct {
	itr  *C.Eina_Iterator
	data unsafe.Pointer
	list *C.Eina_List
}

func newIterator(itr *C.Eina_Iterator) *eIterator {
	return &eIterator{
		itr: itr,
	}
}

func newListIterator(list *C.Eina_List) *eIterator {
	return &eIterator{
		itr:  C.eina_list_iterator_new(list),
		list: list,
	}
}

func (p *eIterator) Next() bool {
	if p.itr != nil {
		if C.eina_iterator_next(p.itr, &p.data) == eTrue {
			return true
		} else {
			C.eina_iterator_free(p.itr)
			p.itr = nil
			if p.list != nil {
				C.eina_list_free(p.list)
			}
			return false
		}
	}
	return false
}

func (p *eIterator) StringSlice() (ret []string) {
	for p.Next() {
		ret = append(ret, C.GoString((*C.char)(p.data)))
	}
	return
}

func (p *eIterator) StringShareSlice() (ret []string) {
	for p.Next() {
		ret = append(ret, goString((*C.Eina_Stringshare)(p.data), true))
	}
	return
}

func (p *eIterator) IntSlice() (ret []int) {
	for p.Next() {
		ret = append(ret, int(*((*C.int)(p.data))))
	}
	return
}

func (p *eIterator) ObjectSlice() (ret []Object) {
	for p.Next() {
		ret = append(ret, wrapObjectBase((*C.Eo)(p.data)))
	}
	return
}

func (p *eIterator) CalendarMarkSlice() (ret []*CalendarMark) {
	for p.Next() {
		ret = append(ret, newCalendarMark((*C.Elm_Calendar_Mark)(p.data)))
	}
	return
}

func (p *eIterator) TransitSlice() (ret []*Transit) {
	for p.Next() {
		ret = append(ret, wrapTransit((*C.Elm_Transit)(p.data)))
	}
	return
}

func (p *eIterator) WidgetItemSlice() (ret []WidgetItem) {
	for p.Next() {
		ret = append(ret, wrapWidgetItemBase((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) MenuItemSlice() (ret []*MenuItem) {
	for p.Next() {
		ret = append(ret, wrapMenuItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) ToolbarItemSlice() (ret []*ToolbarItem) {
	for p.Next() {
		ret = append(ret, wrapToolbarItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) IndexItemSlice() (ret []*IndexItem) {
	for p.Next() {
		ret = append(ret, wrapIndexItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) ListItemSlice() (ret []*ListItem) {
	for p.Next() {
		ret = append(ret, wrapListItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) SlideshowItemSlice() (ret []*SlideshowItem) {
	for p.Next() {
		ret = append(ret, wrapSlideshowItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) MultibuttonentryItemSlice() (ret []*MultibuttonentryItem) {
	for p.Next() {
		ret = append(ret, wrapMultibuttonentryItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) SegmentControlItemSlice() (ret []*SegmentControlItem) {
	for p.Next() {
		ret = append(ret, wrapSegmentControlItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) FlipselectorItemSlice() (ret []*FlipselectorItem) {
	for p.Next() {
		ret = append(ret, wrapFlipselectorItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) NaviframeItemSlice() (ret []*NaviframeItem) {
	for p.Next() {
		ret = append(ret, wrapNaviframeItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) GengridItemSlice() (ret []*GengridItem) {
	for p.Next() {
		ret = append(ret, wrapGengridItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) HoverselItemSlice() (ret []*HoverselItem) {
	for p.Next() {
		ret = append(ret, wrapHoverselItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) CtxpopupItemSlice() (ret []*CtxpopupItem) {
	for p.Next() {
		ret = append(ret, wrapCtxpopupItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) ColorselectorPaletteItemSlice() (ret []*ColorselectorPaletteItem) {
	for p.Next() {
		ret = append(ret, wrapColorselectorPaletteItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) GenlistItemSlice() (ret []*GenlistItem) {
	for p.Next() {
		ret = append(ret, wrapGenlistItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) DiskselectorItemSlice() (ret []*DiskselectorItem) {
	for p.Next() {
		ret = append(ret, wrapDiskselectorItem((*C.Elm_Widget_Item)(p.data)))
	}
	return
}

func (p *eIterator) MapOverlaySlice() (ret []*MapOverlay) {
	for p.Next() {
		ret = append(ret, wrapMapOverlay((*C.Elm_Map_Overlay)(p.data)))
	}
	return
}

func (p *eIterator) WebcamSlice() (ret []*Webcam) {
	for p.Next() {
		ret = append(ret, wrapWebcam((*C.Emotion_Webcam)(p.data)))
	}
	return
}

func (p *eIterator) ColorClassSlice() (ret []*ColorClass) {
	for p.Next() {
		ret = append(ret, wrapColorClass((*C.Elm_Color_Class)(p.data)))
	}
	return
}

func (p *eIterator) ColorOverlaySlice() (ret []*ColorOverlay) {
	for p.Next() {
		ret = append(ret, wrapColorOverlay((*C.Elm_Color_Overlay)(p.data)))
	}
	return
}

func (p *eIterator) TextClassSlice() (ret []*TextClass) {
	for p.Next() {
		ret = append(ret, wrapTextClass((*C.Elm_Text_Class)(p.data)))
	}
	return
}

func (p *eIterator) FontOverlaySlice() (ret []*FontOverlay) {
	for p.Next() {
		ret = append(ret, wrapFontOverlay((*C.Elm_Font_Overlay)(p.data)))
	}
	return
}

func eListToSlice(l *C.Eina_List) []unsafe.Pointer {
	var (
		s    []unsafe.Pointer
		data unsafe.Pointer
	)

	for l != nil {
		data = C.eina_list_data_get(l)
		s = append(s, data)
		l = C.eina_list_next(l)
	}

	return s
}

func eStringListToSlice(l *C.Eina_List, f bool) []string {
	var (
		s    []string
		data unsafe.Pointer
	)
	if f {
		lstHead := l
		defer C.eina_list_free(lstHead)
	}

	for l != nil {
		data = C.eina_list_data_get(l)
		s = append(s, C.GoString((*C.char)(unsafe.Pointer(data))))
		l = C.eina_list_next(l)
	}

	return s
}

func eListToObjectSlice(l *C.Eina_List, f bool) []Object {
	var (
		s    []Object
		data unsafe.Pointer
	)

	if f {
		lstHead := l
		defer C.eina_list_free(lstHead)
	}

	for l != nil {
		data = C.eina_list_data_get(l)
		s = append(s, wrapObjectBase((*C.Eo)(unsafe.Pointer(data))))
		l = C.eina_list_next(l)
	}

	return s
}

// 64 bit string for object and callback identification
/*
func newID() (string, error) {
	id := make([]byte, 8)
	n, err := io.ReadFull(rand.Reader, id)
	if n != len(id) || err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", id[:]), nil
}
*/

var objectID _objectID
var objectItemID _objectID

type _objectID struct {
	id int
	mu sync.Mutex
}

func (p *_objectID) Gen() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.id++
	return p.id
}

type _objectItemID struct {
	id int
	mu sync.Mutex
}

func (p *_objectItemID) Gen() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.id++
	return p.id
}

type Comparer interface {
	Compare(interface{}) int
}

//export go_compare_func
func go_compare_func(data1 C.int, data2 C.int) int {
	d1 := registry.Lookup(int(data1))
	d2 := registry.Lookup(int(data2))
	if d1 == nil || d2 == nil {
		return 0
	}

	if i1, ok := d1.(Comparer); ok {
		return i1.Compare(d2)
	}
	return 0
}

var (
	objectDataMap = newObjDataMap()
	registry      = newRegistryMap()
)

type objDataMap struct {
	m  map[*C.Eo]map[string]interface{}
	mu sync.RWMutex
}

func newObjDataMap() *objDataMap {
	om := &objDataMap{}
	om.m = make(map[*C.Eo]map[string]interface{})
	return om
}

func (p *objDataMap) SetData(obj *C.Eo, key string, data interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.m[obj]; !ok {
		p.m[obj] = make(map[string]interface{})
	}
	p.m[obj][key] = data
}

func (p *objDataMap) Data(obj *C.Eo, key string) interface{} {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if _, ok := p.m[obj]; !ok {
		return nil
	}
	return p.m[obj][key]
}

func (p *objDataMap) DelData(obj *C.Eo, key string) interface{} {
	p.mu.Lock()
	defer p.mu.Unlock()
	if mm, ok := p.m[obj]; ok {
		if data, ok := mm[key]; ok {
			delete(mm, key)
			return data
		}
	}
	return nil
}

func (p *objDataMap) DelEo(obj *C.Eo) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.m, obj)
}

type registryMap struct {
	m     map[int]interface{}
	mu    sync.RWMutex
	index int
}

func newRegistryMap() *registryMap {
	rm := &registryMap{}
	rm.m = make(map[int]interface{})
	return rm
}

func (p *registryMap) Register(x interface{}) int {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.index++
	p.m[p.index] = x
	return p.index
}

func (p *registryMap) Lookup(index int) interface{} {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.m[index]
}

func (p *registryMap) Delete(index int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.m, index)
}

func (p *registryMap) Len() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return len(p.m)
}

// TODO : dump registry info

type slideshowItemMap struct {
	items map[int]*SlideshowItem
	index int
	mu    sync.RWMutex
}

func newSlideshowItemMap() *slideshowItemMap {
	ssim := &slideshowItemMap{}
	ssim.items = make(map[int]*SlideshowItem)
	return ssim
}

func (p *slideshowItemMap) Set(item *SlideshowItem) int {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.index++
	p.items[p.index] = item
	return p.index
}

func (p *slideshowItemMap) Get(index int) *SlideshowItem {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.items[index]
}

//--------------------------------------------------

type eIntList struct {
	obj *C.Eina_List
}

// it can append to nil list te object is nil
func newIntList() *eIntList {
	return &eIntList{}
}

func wrapIntList(o *C.Eina_List) *eIntList {
	if o != nil {
		return &eIntList{o}
	}
	return nil
}

func (p *eIntList) Append(i int) {
	p.obj = C.eina_list_append(p.obj, unsafe.Pointer(&i))
}

func (p *eIntList) Count() int {
	return int(C.eina_list_count(p.obj))
}

//--------------------------------------------------
