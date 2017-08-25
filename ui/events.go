package ui

/*
#include "bridge.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// -------------------------------------

func registerHandler(h Handler) int {
	return registry.Register(h)
}

func registerDataHandler(h DataHandler, data interface{}) int {
	dc := dataCallback{handler: h, data: data}
	return registry.Register(dc)
}

func addSmartCallback(obj *C.Eo, event string, h interface{}) {
	id := registry.Register(h)
	cevent := C.CString(event)
	defer free(cevent)
	C.cgo_smart_callback_add(obj, cevent, unsafe.Pointer(&id))
}

func addSmartCallbackWithData(obj *C.Eo, event string, h DataHandler, data interface{}) {
	id := registerDataHandler(h, data)
	cevent := C.CString(event)
	defer free(cevent)
	C.cgo_smart_callback_add(obj, cevent, unsafe.Pointer(&id))
}

func addSmartCallbackWithEvent(sender string, obj *C.Eo, event string, h EventHandler) {
	id := registry.Register(signalEventInfo{xsender: sender, xsignal: event, handler: h})
	cevent := C.CString(event)
	defer free(cevent)
	C.cgo_smart_callback_add(obj, cevent, unsafe.Pointer(&id))
}

func addSmartCallbackWithEventAndData(sender string, obj *C.Eo, event string, h EventDataHandler, data interface{}) {
	id := registry.Register(signalInfoWithData{xsender: sender, xsignal: event, handler: h, data: data})
	cevent := C.CString(event)
	defer free(cevent)
	C.cgo_smart_callback_add(obj, cevent, unsafe.Pointer(&id))
}

func addEvasObjectCallback(obj *C.Eo, cbtype CallbackType, h interface{}) {
	index := registry.Register(h)
	C.cgo_evas_object_event_callback_add(obj, C.Evas_Callback_Type(cbtype), unsafe.Pointer(&index))
}

func addEvasCallback(evas *C.Evas, cbtype CallbackType, h Handler) {
	index := registry.Register(h)
	C.cgo_evas_event_callback_add(evas, C.Evas_Callback_Type(cbtype), unsafe.Pointer(&index))
}

func addEvasCallbackWithPriority(evas *C.Evas, cbtype CallbackType, priority CanvasEventPriority, h Handler) {
	index := registry.Register(h)
	C.cgo_evas_event_callback_priority_add(evas, C.Evas_Callback_Type(cbtype), C.Evas_Callback_Priority(priority), unsafe.Pointer(&index))
}

//export go_call_handler
func go_call_handler(index C.int, obj *C.Eo, event_info unsafe.Pointer) {
	if h := registry.Lookup(int(index)); h != nil {
		switch cb := h.(type) {
		case Handler:
			cb.Handle()
		case dataCallback:
			cb.handler.Handle(cb.data)
		case signalEventInfo:
			cb.handler.Handle(wrapEvent(cb, event_info))
		case signalInfoWithData:
			cb.handler.Handle(wrapEvent(cb, event_info), cb.data)
		case MouseInHandler:
			cb.Handle(wrapMouseInEvent((*C.Evas_Event_Mouse_In)(event_info)))
		case MouseOutHandler:
			cb.Handle(wrapMouseOutEvent((*C.Evas_Event_Mouse_Out)(event_info)))
		case MouseDownHandler:
			cb.Handle(wrapMouseDownEvent((*C.Evas_Event_Mouse_Down)(event_info)))
		case MouseUpHandler:
			cb.Handle(wrapMouseUpEvent((*C.Evas_Event_Mouse_Up)(event_info)))
		case MouseMoveHandler:
			cb.Handle(wrapMouseMoveEvent((*C.Evas_Event_Mouse_Move)(event_info)))
		case MouseWheelHandler:
			cb.Handle(wrapMouseWheelEvent((*C.Evas_Event_Mouse_Wheel)(event_info)))
		case MultiDownHandler:
			cb.Handle(wrapMultiDownEvent((*C.Evas_Event_Multi_Down)(event_info)))
		case MultiUpHandler:
			cb.Handle(wrapMultiUpEvent((*C.Evas_Event_Multi_Up)(event_info)))
		case MultiMoveHandler:
			cb.Handle(wrapMultiMoveEvent((*C.Evas_Event_Multi_Move)(event_info)))
		case KeyDownHandler:
			cb.Handle(wrapKeyDownEvent((*C.Evas_Event_Key_Down)(event_info)))
		case KeyUpHandler:
			cb.Handle(wrapKeyUpEvent((*C.Evas_Event_Key_Up)(event_info)))
		case HoldHandler:
			cb.Handle(wrapHoldEvent((*C.Evas_Event_Hold)(event_info)))
		case AxisUpdateHandler:
			cb.Handle(wrapAxisUpdateEvent((*C.Evas_Event_Axis_Update)(event_info)))
		default:
			fmt.Println("Unknown Type")
		}
	}
}

// -------------------------------------

type esignalInfo interface {
	Sender() string
	Signal() string
}

type signalEventInfo struct {
	xsender string
	xsignal string
	handler EventHandler
}

func (p signalEventInfo) Sender() string {
	return p.xsender
}

func (p signalEventInfo) Signal() string {
	return p.xsignal
}

type signalInfoWithData struct {
	xsender string
	xsignal string
	handler EventDataHandler
	data    interface{}
}

func (p signalInfoWithData) Sender() string {
	return p.xsender
}

func (p signalInfoWithData) Signal() string {
	return p.xsignal
}

func wrapEvent(info esignalInfo, cinfo unsafe.Pointer) interface{} {
	switch info.Sender() {
	case "elm_list":
		switch info.Signal() {
		case "activated", "clicked,double", "clicked,right", "highlighted", "unhighlighted":
			return wrapListItem((*C.Elm_Widget_Item)(cinfo))
		}
	case "elm_actionslider":
		switch info.Signal() {
		case "selected":
			return C.GoString((*C.char)(cinfo))
		}
	case "elm_colorselector":
		switch info.Signal() {
		case "color,item,selected", "color,item,longpressed":
			return wrapColorselectorPaletteItem((*C.Elm_Widget_Item)(cinfo))
		}
	case "elm_ctxpopup":
		switch info.Signal() {
		case "geometry,update":
			return wrapRectangle((*C.Evas_Coord_Rectangle)(cinfo))
		}
	case "elm_entry":
		switch info.Signal() {
		case "anchor,clicked", "anchor,down", "anchor,in", "anchor,out", "anchor,up":
			return wrapEntryAnchorInfo((*C.Elm_Entry_Anchor_Info)(cinfo))
		}
	case "elm_fileselector_button":
		switch info.Signal() {
		case "file,chosen":
			return goString((*C.Eina_Stringshare)(cinfo))
		}
	case "elm_fileselector_entry":
		switch info.Signal() {
		case "file,chosen":
			return goString((*C.Eina_Stringshare)(cinfo))
		}
	case "elm_fileselector":
		switch info.Signal() {
		case "directory,open", "done", "activated":
			return goString((*C.Eina_Stringshare)(cinfo))
		}
	case "elm_flipselector":
		switch info.Signal() {
		case "selected":
			return wrapFlipselectorItem((*C.Elm_Widget_Item)(cinfo))
		}
	case "elm_hoversel":
		switch info.Signal() {
		case "selected":
			return wrapHoverselItem((*C.Elm_Widget_Item)(cinfo))
		}
	case "elm_image":
		switch info.Signal() {
		case "drop":
			return C.GoString((*C.char)(cinfo))
		case "download,progress":
			return wrapImageProgress((*C.Elm_Image_Progress)(cinfo))
		case "download,error":
			return wrapImageError((*C.Elm_Image_Error)(cinfo))
		}
	case "elm_index":
		switch info.Signal() { // item data pointer ???
		case "changed", "delay,changed", "selected":
			return wrapIndexItem((*C.Elm_Widget_Item)(cinfo))
		}
	case "elm_panel":
		switch info.Signal() {
		case "scroll":
			return wrapPanelScrollInfo((*C.Elm_Panel_Scroll_Info)(cinfo))
		}
	case "elm_segment_control":
		switch info.Signal() {
		case "changed":
			return wrapSegmentControlItem((*C.Elm_Widget_Item)(cinfo))
		}
	case "elm_slideshow":
		switch info.Signal() {
		case "changed", "transition,end":
			return wrapSlideshowItem((*C.Elm_Widget_Item)(cinfo))
		}
	case "elm_toolbar":
		switch info.Signal() {
		case "selected":
			return wrapToolbarItem((*C.Elm_Widget_Item)(cinfo))
		}
	case "elm_gengrid":
		switch info.Signal() {
		case "activated", "pressed", "released", "clicked,double", "clicked,right", "longpressed", "selected", "unselected", "realized", "unrealized", "moved", "highlighted", "unhighlighted", "item,reorder,anim,start", "item,reorder,anim,stop":
			return wrapGengridItem((*C.Elm_Widget_Item)(cinfo))
		}
	case "elm_genlist":
		switch info.Signal() {
		case "activated", "pressed", "released", "clicked,double", "clicked,right", "selected", "unselected", "expanded", "contracted", "expand,request", "contract,request", "realized", "longpressed", "moved", "moved,after", "moved,before", "highlighted", "unhighlighted":
			return wrapGenlistItem((*C.Elm_Widget_Item)(cinfo))
		}
	default:
		fmt.Printf("Unhandled sender: %s, Signal: %s\n", info.Sender(), info.Signal())
	}

	return nil
}

// -------------------------------------

type Handler interface {
	Handle()
}

type handler struct {
	fn func()
}

func (p *handler) Handle() {
	p.fn()
}

func HandleFunc(f func()) Handler {
	return &handler{fn: f}
}

// -------------------------------------

type dataCallback struct {
	handler DataHandler
	data    interface{}
}

type DataHandler interface {
	Handle(interface{})
}

type dataHandler struct {
	fn func(interface{})
}

func (p *dataHandler) Handle(data interface{}) {
	p.fn(data)
}

func DataHandleFunc(f func(interface{})) DataHandler {
	return &dataHandler{fn: f}
}

// -------------------------------------

type EventHandler interface {
	Handle(interface{})
}

type eventHandler struct {
	fn func(interface{})
}

func (p *eventHandler) Handle(event interface{}) {
	p.fn(event)
}

func EventHandleFunc(f func(interface{})) EventHandler {
	return &eventHandler{fn: f}
}

// -------------------------------------

type eventDataCallback struct {
	handler EventDataHandler
	data    interface{}
}

type EventDataHandler interface {
	Handle(interface{}, interface{})
}

type eventDataHandler struct {
	fn func(interface{}, interface{})
}

func (p *eventDataHandler) Handle(event interface{}, data interface{}) {
	p.fn(event, data)
}

func EventDataHandleFunc(f func(interface{}, interface{})) EventDataHandler {
	return &eventDataHandler{fn: f}
}

// -------------------------------------

type MouseInHandler interface {
	Handle(*MouseInEvent)
}

type mouseInEventHandler struct {
	fn func(*MouseInEvent)
}

func (p *mouseInEventHandler) Handle(ev *MouseInEvent) {
	p.fn(ev)
}

func MouseInHandleFunc(f func(*MouseInEvent)) MouseInHandler {
	return &mouseInEventHandler{fn: f}
}

// -------------------------------------
type MouseOutHandler interface {
	Handle(*MouseOutEvent)
}

type mouseOutEventHandler struct {
	fn func(*MouseOutEvent)
}

func (p *mouseOutEventHandler) Handle(ev *MouseOutEvent) {
	p.fn(ev)
}

func MouseOutHandleFunc(f func(*MouseOutEvent)) MouseOutHandler {
	return &mouseOutEventHandler{fn: f}
}

// -------------------------------------

type MouseDownHandler interface {
	Handle(*MouseDownEvent)
}

type mouseDownEventHandler struct {
	fn func(*MouseDownEvent)
}

func (p *mouseDownEventHandler) Handle(ev *MouseDownEvent) {
	p.fn(ev)
}

func MouseDownHandleFunc(f func(*MouseDownEvent)) MouseDownHandler {
	return &mouseDownEventHandler{fn: f}
}

// -------------------------------------

type MouseUpHandler interface {
	Handle(*MouseUpEvent)
}

type mouseUpEventHandler struct {
	fn func(*MouseUpEvent)
}

func (p *mouseUpEventHandler) Handle(ev *MouseUpEvent) {
	p.fn(ev)
}

func MouseUpHandleFunc(f func(*MouseUpEvent)) MouseUpHandler {
	return &mouseUpEventHandler{fn: f}
}

// -------------------------------------

type MouseMoveHandler interface {
	Handle(*MouseMoveEvent)
}

type mouseMoveEventHandler struct {
	fn func(*MouseMoveEvent)
}

func (p *mouseMoveEventHandler) Handle(ev *MouseMoveEvent) {
	p.fn(ev)
}

func MouseMoveHandleFunc(f func(*MouseMoveEvent)) MouseMoveHandler {
	return &mouseMoveEventHandler{fn: f}
}

// -------------------------------------

type MouseWheelHandler interface {
	Handle(*MouseWheelEvent)
}

type mouseWheelEventHandler struct {
	fn func(*MouseWheelEvent)
}

func (p *mouseWheelEventHandler) Handle(ev *MouseWheelEvent) {
	p.fn(ev)
}

func MouseWheelHandleFunc(f func(*MouseWheelEvent)) MouseWheelHandler {
	return &mouseWheelEventHandler{fn: f}
}

// -------------------------------------

type MultiDownHandler interface {
	Handle(*MultiDownEvent)
}

type multiDownEventHandler struct {
	fn func(*MultiDownEvent)
}

func (p *multiDownEventHandler) Handle(ev *MultiDownEvent) {
	p.fn(ev)
}

func MultiDownHandleFunc(f func(*MultiDownEvent)) MultiDownHandler {
	return &multiDownEventHandler{fn: f}
}

// -------------------------------------

type MultiUpHandler interface {
	Handle(*MultiUpEvent)
}

type multiUpEventHandler struct {
	fn func(*MultiUpEvent)
}

func (p *multiUpEventHandler) Handle(ev *MultiUpEvent) {
	p.fn(ev)
}

func MultiUpHandleFunc(f func(*MultiUpEvent)) MultiUpHandler {
	return &multiUpEventHandler{fn: f}
}

// -------------------------------------

type MultiMoveHandler interface {
	Handle(*MultiMoveEvent)
}

type multiMoveEventHandler struct {
	fn func(*MultiMoveEvent)
}

func (p *multiMoveEventHandler) Handle(ev *MultiMoveEvent) {
	p.fn(ev)
}

func MultiMoveHandleFunc(f func(*MultiMoveEvent)) MultiMoveHandler {
	return &multiMoveEventHandler{fn: f}
}

// -------------------------------------

type KeyDownHandler interface {
	Handle(*KeyDownEvent)
}

type keyDownEventHandler struct {
	fn func(*KeyDownEvent)
}

func (p *keyDownEventHandler) Handle(ev *KeyDownEvent) {
	p.fn(ev)
}

func KeyDownHandleFunc(f func(*KeyDownEvent)) KeyDownHandler {
	return &keyDownEventHandler{fn: f}
}

// -------------------------------------

type KeyUpHandler interface {
	Handle(*KeyUpEvent)
}

type keyUpEventHandler struct {
	fn func(*KeyUpEvent)
}

func (p *keyUpEventHandler) Handle(ev *KeyUpEvent) {
	p.fn(ev)
}

func KeyUpHandleFunc(f func(*KeyUpEvent)) KeyUpHandler {
	return &keyUpEventHandler{fn: f}
}

// -------------------------------------

type HoldHandler interface {
	Handle(*HoldEvent)
}

type holdEventHandler struct {
	fn func(*HoldEvent)
}

func (p *holdEventHandler) Handle(ev *HoldEvent) {
	p.fn(ev)
}

func HoldHandleFunc(f func(*HoldEvent)) HoldHandler {
	return &holdEventHandler{fn: f}
}

// -------------------------------------

type AxisUpdateHandler interface {
	Handle(*AxisUpdateEvent)
}

type axisUpdateEventHandler struct {
	fn func(*AxisUpdateEvent)
}

func (p *axisUpdateEventHandler) Handle(ev *AxisUpdateEvent) {
	p.fn(ev)
}

func AxisUpdateHandleFunc(f func(*AxisUpdateEvent)) AxisUpdateHandler {
	return &axisUpdateEventHandler{fn: f}
}

/*
	Event Types
*/

type Point struct {
	obj *C.Evas_Point
}

func wrapPoint(o *C.Evas_Point) *Point {
	if o != nil {
		return &Point{o}
	}
	return nil
}

func (p *Point) X() int {
	return int(p.obj.x)
}

func (p *Point) Y() int {
	return int(p.obj.y)
}

type CoordPoint struct {
	obj *C.Evas_Coord_Point
}

func wrapCoordPoint(o *C.Evas_Coord_Point) *CoordPoint {
	if o != nil {
		return &CoordPoint{o}
	}
	return nil
}

func (p *CoordPoint) X() int {
	return int(p.obj.x)
}

func (p *CoordPoint) Y() int {
	return int(p.obj.y)
}

type Size struct {
	obj *C.Evas_Coord_Size
}

func wrapSize(o *C.Evas_Coord_Size) *Size {
	if o != nil {
		return &Size{o}
	}
	return nil
}

func (p *Size) Width() int {
	return int(p.obj.w)
}

func (p *Size) Height() int {
	return int(p.obj.h)
}

type Rectangle struct {
	obj *C.Evas_Coord_Rectangle
}

func wrapRectangle(o *C.Evas_Coord_Rectangle) *Rectangle {
	if o != nil {
		return &Rectangle{o}
	}
	return nil
}

func (p *Rectangle) X() int {
	return int(p.obj.x)
}

func (p *Rectangle) Y() int {
	return int(p.obj.y)
}

func (p *Rectangle) Width() int {
	return int(p.obj.w)
}

func (p *Rectangle) Height() int {
	return int(p.obj.h)
}

type PrecisionPoint struct {
	obj *C.Evas_Coord_Precision_Point
}

func wrapPrecisionPoint(o *C.Evas_Coord_Precision_Point) *PrecisionPoint {
	if o != nil {
		return &PrecisionPoint{o}
	}
	return nil
}

func (p *PrecisionPoint) X() int {
	return int(p.obj.x)
}

func (p *PrecisionPoint) Y() int {
	return int(p.obj.y)
}

func (p *PrecisionPoint) XSub() float64 {
	return float64(p.obj.xsub)
}

func (p *PrecisionPoint) YSub() float64 {
	return float64(p.obj.ysub)
}

type Position struct {
	obj *C.Evas_Position
}

func wrapPosition(o *C.Evas_Position) *Position {
	if o != nil {
		return &Position{o}
	}
	return nil
}

func (p *Position) Output() *Point {
	return wrapPoint(&p.obj.output)
}

func (p *Position) Canvas() *CoordPoint {
	return wrapCoordPoint(&p.obj.canvas)
}

type PrecisionPosition struct {
	obj *C.Evas_Precision_Position
}

func wrapPrecisionPosition(o *C.Evas_Precision_Position) *PrecisionPosition {
	if o != nil {
		return &PrecisionPosition{o}
	}
	return nil
}

func (p *PrecisionPosition) Output() *Point {
	return wrapPoint(&p.obj.output)
}

func (p *PrecisionPosition) Canvas() *PrecisionPoint {
	return wrapPrecisionPoint(&p.obj.canvas)
}

type MouseInEvent struct {
	obj *C.Evas_Event_Mouse_In
}

func wrapMouseInEvent(o *C.Evas_Event_Mouse_In) *MouseInEvent {
	if o != nil {
		return &MouseInEvent{o}
	}
	return nil
}

func (p *MouseInEvent) Buttons() int {
	return int(p.obj.buttons)
}

func (p *MouseInEvent) Output() *Point {
	return wrapPoint(&p.obj.output)
}

func (p *MouseInEvent) Canvas() *CoordPoint {
	return wrapCoordPoint(&p.obj.canvas)
}

func (p *MouseInEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

func (p *MouseInEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

func (p *MouseInEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

type MouseOutEvent struct {
	obj *C.Evas_Event_Mouse_Out
}

func wrapMouseOutEvent(o *C.Evas_Event_Mouse_Out) *MouseOutEvent {
	if o != nil {
		return &MouseOutEvent{o}
	}
	return nil
}

func (p *MouseOutEvent) Buttons() int {
	return int(p.obj.buttons)
}

func (p *MouseOutEvent) Output() *Point {
	return wrapPoint(&p.obj.output)
}

func (p *MouseOutEvent) Canvas() *CoordPoint {
	return wrapCoordPoint(&p.obj.canvas)
}

func (p *MouseOutEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

func (p *MouseOutEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

func (p *MouseOutEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

type MouseDownEvent struct {
	obj *C.Evas_Event_Mouse_Down
}

func wrapMouseDownEvent(o *C.Evas_Event_Mouse_Down) *MouseDownEvent {
	if o != nil {
		return &MouseDownEvent{o}
	}
	return nil
}

func (p *MouseDownEvent) Button() int {
	return int(p.obj.button)
}

func (p *MouseDownEvent) Output() *Point {
	return wrapPoint(&p.obj.output)
}

func (p *MouseDownEvent) Canvas() *CoordPoint {
	return wrapCoordPoint(&p.obj.canvas)
}

func (p *MouseDownEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

func (p *MouseDownEvent) ButtonFlags() ButtonFlags {
	return ButtonFlags(p.obj.flags)
}

func (p *MouseDownEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

func (p *MouseDownEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

type MouseUpEvent struct {
	obj *C.Evas_Event_Mouse_Up
}

func wrapMouseUpEvent(o *C.Evas_Event_Mouse_Up) *MouseUpEvent {
	if o != nil {
		return &MouseUpEvent{o}
	}
	return nil
}

func (p *MouseUpEvent) Buttons() int {
	return int(p.obj.button)
}

func (p *MouseUpEvent) Output() *Point {
	return wrapPoint(&p.obj.output)
}

func (p *MouseUpEvent) Canvas() *CoordPoint {
	return wrapCoordPoint(&p.obj.canvas)
}

func (p *MouseUpEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

func (p *MouseUpEvent) ButtonFlags() ButtonFlags {
	return ButtonFlags(p.obj.flags)
}

func (p *MouseUpEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

func (p *MouseUpEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

type MouseMoveEvent struct {
	obj *C.Evas_Event_Mouse_Move
}

func wrapMouseMoveEvent(o *C.Evas_Event_Mouse_Move) *MouseMoveEvent {
	if o != nil {
		return &MouseMoveEvent{o}
	}
	return nil
}

func (p *MouseMoveEvent) Buttons() int {
	return int(p.obj.buttons)
}

func (p *MouseMoveEvent) CurPos() *Position {
	return wrapPosition(&p.obj.cur)
}

func (p *MouseMoveEvent) PrevPos() *Position {
	return wrapPosition(&p.obj.prev)
}

func (p *MouseMoveEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

func (p *MouseMoveEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

func (p *MouseMoveEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

type MultiDownEvent struct {
	obj *C.Evas_Event_Multi_Down
}

func wrapMultiDownEvent(o *C.Evas_Event_Multi_Down) *MultiDownEvent {
	if o != nil {
		return &MultiDownEvent{o}
	}
	return nil
}

func (p *MultiDownEvent) Device() int {
	return int(p.obj.device)
}

func (p *MultiDownEvent) Radius() float64 {
	return float64(p.obj.radius)
}

func (p *MultiDownEvent) XRadius() float64 {
	return float64(p.obj.radius_x)
}

func (p *MultiDownEvent) YRadius() float64 {
	return float64(p.obj.radius_y)
}

func (p *MultiDownEvent) Pressure() float64 {
	return float64(p.obj.pressure)
}

func (p *MultiDownEvent) Angle() float64 {
	return float64(p.obj.angle)
}

func (p *MultiDownEvent) Output() *Point {
	return wrapPoint(&p.obj.output)
}

func (p *MultiDownEvent) Canvas() *PrecisionPoint {
	return wrapPrecisionPoint(&p.obj.canvas)
}

func (p *MultiDownEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

func (p *MultiDownEvent) ButtonFlags() ButtonFlags {
	return ButtonFlags(p.obj.flags)
}

func (p *MultiDownEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

func (p *MultiDownEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

type MultiUpEvent struct {
	obj *C.Evas_Event_Multi_Up
}

func wrapMultiUpEvent(o *C.Evas_Event_Multi_Up) *MultiUpEvent {
	if o != nil {
		return &MultiUpEvent{o}
	}
	return nil
}

func (p *MultiUpEvent) Device() int {
	return int(p.obj.device)
}

func (p *MultiUpEvent) Radius() float64 {
	return float64(p.obj.radius)
}

func (p *MultiUpEvent) XRadius() float64 {
	return float64(p.obj.radius_x)
}

func (p *MultiUpEvent) YRadius() float64 {
	return float64(p.obj.radius_y)
}

func (p *MultiUpEvent) Pressure() float64 {
	return float64(p.obj.pressure)
}

func (p *MultiUpEvent) Angle() float64 {
	return float64(p.obj.angle)
}

func (p *MultiUpEvent) Output() *Point {
	return wrapPoint(&p.obj.output)
}

func (p *MultiUpEvent) Canvas() *PrecisionPoint {
	return wrapPrecisionPoint(&p.obj.canvas)
}

func (p *MultiUpEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

func (p *MultiUpEvent) ButtonFlags() ButtonFlags {
	return ButtonFlags(p.obj.flags)
}

func (p *MultiUpEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

func (p *MultiUpEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

type MultiMoveEvent struct {
	obj *C.Evas_Event_Multi_Move
}

func wrapMultiMoveEvent(o *C.Evas_Event_Multi_Move) *MultiMoveEvent {
	if o != nil {
		return &MultiMoveEvent{o}
	}
	return nil
}

func (p *MultiMoveEvent) Radius() float64 {
	return float64(p.obj.radius)
}

func (p *MultiMoveEvent) XRadius() float64 {
	return float64(p.obj.radius_x)
}

func (p *MultiMoveEvent) YRadius() float64 {
	return float64(p.obj.radius_y)
}

func (p *MultiMoveEvent) Pressure() float64 {
	return float64(p.obj.pressure)
}

func (p *MultiMoveEvent) Angle() float64 {
	return float64(p.obj.angle)
}

func (p *MultiMoveEvent) CurPos() *PrecisionPosition {
	return wrapPrecisionPosition(&p.obj.cur)
}

func (p *MultiMoveEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

func (p *MultiMoveEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

func (p *MultiMoveEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

type MouseWheelEvent struct {
	obj *C.Evas_Event_Mouse_Wheel
}

func wrapMouseWheelEvent(o *C.Evas_Event_Mouse_Wheel) *MouseWheelEvent {
	if o != nil {
		return &MouseWheelEvent{o}
	}
	return nil
}

func (p *MouseWheelEvent) Direction() int {
	return int(p.obj.direction)
}

func (p *MouseWheelEvent) Z() int {
	return int(p.obj.z)
}

func (p *MouseWheelEvent) Output() *Point {
	return wrapPoint(&p.obj.output)
}

func (p *MouseWheelEvent) Canvas() *CoordPoint {
	return wrapCoordPoint(&p.obj.canvas)
}

func (p *MouseWheelEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

func (p *MouseWheelEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

func (p *MouseWheelEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

type KeyDownEvent struct {
	obj *C.Evas_Event_Key_Down
}

func wrapKeyDownEvent(o *C.Evas_Event_Key_Down) *KeyDownEvent {
	if o != nil {
		return &KeyDownEvent{o}
	}
	return nil
}

func (p *KeyDownEvent) KeyName() string {
	return C.GoString(p.obj.keyname)
}

func (p *KeyDownEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

func (p *KeyDownEvent) Key() string {
	return C.GoString(p.obj.key)
}

func (p *KeyDownEvent) Str() string {
	return C.GoString(p.obj.string)
}

func (p *KeyDownEvent) Compose() string {
	return C.GoString(p.obj.compose)
}

func (p *KeyDownEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

func (p *KeyDownEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

type KeyUpEvent struct {
	obj *C.Evas_Event_Key_Up
}

func wrapKeyUpEvent(o *C.Evas_Event_Key_Up) *KeyUpEvent {
	if o != nil {
		return &KeyUpEvent{o}
	}
	return nil
}

func (p *KeyUpEvent) KeyName() string {
	return C.GoString(p.obj.keyname)
}

func (p *KeyUpEvent) IsModifierSet(modifier string) bool {
	return isKeyModifierSet(p.obj.modifiers, modifier)
}

func (p *KeyUpEvent) Key() string {
	return C.GoString(p.obj.key)
}

func (p *KeyUpEvent) Str() string {
	return C.GoString(p.obj.string)
}

func (p *KeyUpEvent) Compose() string {
	return C.GoString(p.obj.compose)
}

func (p *KeyUpEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

func (p *KeyUpEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

type HoldEvent struct {
	obj *C.Evas_Event_Hold
}

func wrapHoldEvent(o *C.Evas_Event_Hold) *HoldEvent {
	if o != nil {
		return &HoldEvent{o}
	}
	return nil
}

func (p *HoldEvent) Hold() int {
	return int(p.obj.hold)
}

func (p *HoldEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

func (p *HoldEvent) Flags() EventFlags {
	return EventFlags(p.obj.event_flags)
}

type Axis struct {
	obj *C.Evas_Axis
}

func wrapAxis(o *C.Evas_Axis) *Axis {
	if o != nil {
		return &Axis{o}
	}
	return nil
}

func (p *Axis) Label() AxisLabel {
	return AxisLabel(p.obj.label)
}

func (p *Axis) Value() float64 {
	return float64(p.obj.value)
}

type AxisUpdateEvent struct {
	obj *C.Evas_Event_Axis_Update
}

func wrapAxisUpdateEvent(o *C.Evas_Event_Axis_Update) *AxisUpdateEvent {
	if o != nil {
		return &AxisUpdateEvent{o}
	}
	return nil
}

func (p *AxisUpdateEvent) Timestamp() uint {
	return uint(p.obj.timestamp)
}

func (p *AxisUpdateEvent) Device() int {
	return int(p.obj.device)
}

func (p *AxisUpdateEvent) ToolId() int {
	return int(p.obj.toolid)
}

func (p *AxisUpdateEvent) NAxis() int {
	return int(p.obj.naxis)
}

func (p *AxisUpdateEvent) Axis() *Axis {
	return wrapAxis(p.obj.axis)
}
