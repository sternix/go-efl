package ui

/*
#include "bridge.h"
*/
import "C"

type Transit struct {
	obj *C.Elm_Transit
}

func NewTransit() *Transit {
	return &Transit{
		obj: C.elm_transit_add(),
	}
}

func wrapTransit(o *C.Elm_Transit) *Transit {
	return &Transit{
		obj: o,
	}
}

func (p *Transit) Del() {
	C.elm_transit_del(p.obj)
}

func (p *Transit) AddObject(o Object) {
	C.elm_transit_object_add(p.obj, o.eo())
}

func (p *Transit) RemoveObject(o Object) {
	C.elm_transit_object_remove(p.obj, o.eo())
}

func (p *Transit) Objects() []Object {
	lst := C.elm_transit_objects_get(p.obj)
	if lst != nil {
		return newListIterator(lst).ObjectSlice()
	}
	return nil
}

func (p *Transit) SetKeepObjectsFinalState(b bool) {
	C.elm_transit_objects_final_state_keep_set(p.obj, eBool(b))
}

func (p *Transit) IsKeepObjectsFinalState() bool {
	return C.elm_transit_objects_final_state_keep_get(p.obj) == eTrue
}

func (p *Transit) SetEnableEvents(b bool) {
	C.elm_transit_event_enabled_set(p.obj, eBool(b))
}

func (p *Transit) IsEventsEnabled() bool {
	return C.elm_transit_event_enabled_get(p.obj) == eTrue
}

func (p *Transit) SetAutoReverse(b bool) {
	C.elm_transit_auto_reverse_set(p.obj, eBool(b))
}

func (p *Transit) IsAutoReversed() bool {
	return C.elm_transit_auto_reverse_get(p.obj) == eTrue
}

func (p *Transit) SetRepeatCount(count int) {
	C.elm_transit_repeat_times_set(p.obj, C.int(count))
}

func (p *Transit) RepeatCount() int {
	return int(C.elm_transit_repeat_times_get(p.obj))
}

func (p *Transit) SetTweenMode(mode TransitTweenMode) {
	C.elm_transit_tween_mode_set(p.obj, C.Elm_Transit_Tween_Mode(mode))
}

func (p *Transit) TweenMode() TransitTweenMode {
	return TransitTweenMode(C.elm_transit_tween_mode_get(p.obj))
}

func (p *Transit) SetTweenModeFactor(v1, v2 float64) {
	C.elm_transit_tween_mode_factor_set(p.obj, C.double(v1), C.double(v2))
}

func (p *Transit) TweenModeFactor() (float64, float64) {
	var v1, v2 C.double
	C.elm_transit_tween_mode_factor_get(p.obj, &v1, &v2)
	return float64(v1), float64(v2)
}

func (p *Transit) SetTweenModeFactorN(v_size uint) float64 {
	var v C.double
	C.elm_transit_tween_mode_factor_n_set(p.obj, C.uint(v_size), &v)
	return float64(v)
}

func (p *Transit) SetDuration(dur float64) {
	C.elm_transit_duration_set(p.obj, C.double(dur))
}

func (p *Transit) Duration() float64 {
	return float64(C.elm_transit_duration_get(p.obj))
}

func (p *Transit) Start() {
	C.elm_transit_go(p.obj)
}

func (p *Transit) StartIn(seconds float64) {
	C.elm_transit_go_in(p.obj, C.double(seconds))
}

func (p *Transit) SetPaused(b bool) {
	C.elm_transit_paused_set(p.obj, eBool(b))
}

func (p *Transit) IsPaused() bool {
	return C.elm_transit_paused_get(p.obj) == eTrue
}

func (p *Transit) ProgressValue() float64 {
	return float64(C.elm_transit_progress_value_get(p.obj))
}

func (p *Transit) AddChain(tr *Transit) {
	C.elm_transit_chain_transit_add(p.obj, tr.obj)
}

func (p *Transit) DelChain(tr *Transit) {
	C.elm_transit_chain_transit_del(p.obj, tr.obj)
}

func (p *Transit) Chains() []*Transit {
	lst := C.elm_transit_chain_transits_get(p.obj)
	if lst != nil {
		return newListIterator(lst).TransitSlice()
	}
	return nil
}

func (p *Transit) SetSmooth(b bool) {
	C.elm_transit_smooth_set(p.obj, eBool(b))
}

func (p *Transit) IsSmooth() bool {
	return C.elm_transit_smooth_get(p.obj) == eTrue
}

func (p *Transit) AddFadeEffect() {
	C.elm_transit_effect_fade_add(p.obj)
}

func (p *Transit) AddBlendEffect() {
	C.elm_transit_effect_blend_add(p.obj)
}

func (p *Transit) AddRotationEffect(fd, td float32) {
	C.elm_transit_effect_rotation_add(p.obj, C.float(fd), C.float(td))
}

func (p *Transit) AddImageAnimationEffect(images []string) {
	C.elm_transit_effect_image_animation_add(p.obj, stringSliceToEList(images))
}

func (p *Transit) AddZoomEffect(fr, tr float32) {
	C.elm_transit_effect_zoom_add(p.obj, C.float(fr), C.float(tr))
}

func (p *Transit) AddResizingEffect(fw, fh, tw, th int) {
	C.elm_transit_effect_resizing_add(p.obj, C.Evas_Coord(fw), C.Evas_Coord(fh), C.Evas_Coord(tw), C.Evas_Coord(th))
}

func (p *Transit) AddTranslationEffect(fdx, fdy, tdx, tdy int) {
	C.elm_transit_effect_translation_add(p.obj, C.Evas_Coord(fdx), C.Evas_Coord(fdy), C.Evas_Coord(tdx), C.Evas_Coord(tdy))
}

func (p *Transit) AddFlipEffect(axis TransitEffectFlipAxis, cw bool) {
	C.elm_transit_effect_flip_add(p.obj, C.Elm_Transit_Effect_Flip_Axis(axis), eBool(cw))
}

func (p *Transit) AddResizableFlipEffect(axis TransitEffectFlipAxis, cw bool) {
	C.elm_transit_effect_resizable_flip_add(p.obj, C.Elm_Transit_Effect_Flip_Axis(axis), eBool(cw))
}

func (p *Transit) AddWipeEffect(typ TransitEffectWipeType, dir TransitEffectWipeDir) {
	C.elm_transit_effect_wipe_add(p.obj, C.Elm_Transit_Effect_Wipe_Type(typ), C.Elm_Transit_Effect_Wipe_Dir(dir))
}

func (p *Transit) AddColorEffect(fr, fg, fb, fa, tr, tg, tb, ta uint) {
	C.elm_transit_effect_color_add(p.obj, C.uint(fr), C.uint(fg), C.uint(fb), C.uint(fa), C.uint(tr), C.uint(tg), C.uint(tb), C.uint(ta))
}

/*
TODO:
void 	elm_transit_effect_add (Elm_Transit *transit, Elm_Transit_Effect_Transition_Cb transition_cb, Elm_Transit_Effect *effect, Elm_Transit_Effect_End_Cb end_cb)
void 	elm_transit_effect_del (Elm_Transit *transit, Elm_Transit_Effect_Transition_Cb transition_cb, Elm_Transit_Effect *effect)
void 	elm_transit_del_cb_set (Elm_Transit *transit, Elm_Transit_Del_Cb cb, void *data)
*/
