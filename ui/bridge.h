#pragma once

#include <Elementary.h>
#include <Emotion.h>

void cgo_smart_callback_func(void *data , Eo *obj , void *event_info);
void cgo_smart_callback_add(Eo *obj , const char *event , void *data);
void cgo_smart_callback_del(Eo *obj , const char *event);

void cgo_evas_object_event_callback_add(Eo *obj, Evas_Callback_Type type, void *data);
void cgo_evas_object_event_callback_func(void *data, Evas *e, Eo *obj, void *event_info);


void cgo_evas_event_callback_add(Evas *e, Evas_Callback_Type type,void *data);
void cgo_evas_event_callback_priority_add(Evas *e, Evas_Callback_Type type, Evas_Callback_Priority priority, void *data);

Eina_Iterator *cgo_carray_iterator_new(char **array);
Eina_List *cgo_carray_list_new(char **array, int len);


char **cgo_null_term_string_get(char *str);

Eo *cgo_slideshow_item_get_func(void *data, Eo *obj);
Elm_Slideshow_Item_Class cgo_slideshow_item_class_get();
Elm_Gengrid_Item_Class * cgo_elm_gengrid_item_class_new();
Elm_Genlist_Item_Class * cgo_elm_genlist_item_class_new();

int cgo_compare_func(const void *data1 ,const void *data2);
/*
TODO:
get evas_version
https://github.com/tasn/efl/blob/master/src/lib/evas/Evas_Common.h
evas_version
*/


void cgo_elm_object_item_del_cb_set(Elm_Object_Item *it);

void cgo_ecore_main_loop_thread_safe_call_async(void *data);
void cgo_ecore_main_loop_thread_safe_call_sync(void *data);


Ecore_Job *cgo_ecore_job_add(void *data);
Ecore_Timer * cgo_ecore_timer_add (double in, const void *data);
Ecore_Timer * cgo_ecore_timer_loop_add (double in, const void *data);


void cgo_web_console_message_hook_set(Elm_Web *obj, void *data);
void cgo_web_window_create_hook_set(Elm_Web *obj, void *data);
void cgo_web_dialog_confirm_hook_set(Elm_Web *obj, void *data);
void cgo_web_dialog_prompt_hook_set(Elm_Web *obj, void *data);
void cgo_web_dialog_alert_hook_set(Elm_Web *obj, void *data);



static inline int cgo_evas_textgrid_cell_bold_get(Evas_Textgrid_Cell *cell) {
	return cell->bold;
}

static inline void cgo_evas_textgrid_cell_bold_set(Evas_Textgrid_Cell *cell, unsigned short v) {
	cell->bold = v;
}

static inline int cgo_evas_textgrid_cell_italic_get(Evas_Textgrid_Cell *cell) {
	return cell->italic;
}

static inline void cgo_evas_textgrid_cell_italic_set(Evas_Textgrid_Cell *cell, unsigned short v) {
	cell->italic = v;
}

static inline int cgo_evas_textgrid_cell_underline_get(Evas_Textgrid_Cell *cell) {
	return cell->underline;
}

static inline void cgo_evas_textgrid_cell_underline_set(Evas_Textgrid_Cell *cell, unsigned short v) {
	cell->underline = v;
}

static inline int cgo_evas_textgrid_cell_strikethrough_get(Evas_Textgrid_Cell *cell) {
	return cell->strikethrough;
}

static inline void cgo_evas_textgrid_cell_strikethrough_set(Evas_Textgrid_Cell *cell, unsigned short v) {
	cell->strikethrough = v;
}

static inline int cgo_evas_textgrid_cell_fg_extended_get(Evas_Textgrid_Cell *cell) {
	return cell->fg_extended;
}

static inline void cgo_evas_textgrid_cell_fg_extended_set(Evas_Textgrid_Cell *cell,unsigned short v) {
	cell->fg_extended = v;
}

static inline int cgo_evas_textgrid_cell_bg_extended_get(Evas_Textgrid_Cell *cell) {
	return cell->bg_extended;
}

static inline void cgo_evas_textgrid_cell_bg_extended_set(Evas_Textgrid_Cell *cell,unsigned short v) {
	cell->bg_extended = v;
}

static inline int cgo_evas_textgrid_cell_double_width_get(Evas_Textgrid_Cell *cell) {
	return cell->double_width;
}

static inline void cgo_evas_textgrid_cell_double_width_set(Evas_Textgrid_Cell *cell, unsigned short v) {
	cell->double_width = v;
}


Eina_Bool cgo_elm_cnp_selection_get(const Eo *obj, Elm_Sel_Type selection, Elm_Sel_Format format, void *data);
Eina_Bool cgo_elm_cnp_selection_set(Eo *obj, Elm_Sel_Type selection, Elm_Sel_Format format, char *buf);
void cgo_elm_selection_loss_callback_set(Eo *obj, Elm_Sel_Type selection,void *data);




void cgo_elm_win_wm_rotation_available_rotations_set(Elm_Win *obj, Eina_List *list);
Eina_List *cgo_elm_win_wm_rotation_available_rotations_get(const Elm_Win *obj);
