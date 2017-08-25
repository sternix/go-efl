#include "bridge.h"
#include "_cgo_export.h"

void cgo_smart_callback_func(void *data , Eo *obj , void *event_info) {
	int *handle_id = (int *)data;
	go_call_handler(*handle_id , obj , event_info);
}

void cgo_smart_callback_add(Eo *obj , const char *event , void *data) {
	evas_object_smart_callback_add(obj , event , cgo_smart_callback_func , data);
}

void cgo_smart_callback_del(Eo *obj , const char *event) {
	evas_object_smart_callback_del(obj , event , cgo_smart_callback_func);
}

void cgo_evas_object_event_callback_func(void *data, Evas *e, Eo *obj, void *event_info) {
	int *handle_id = (int *)data;
	go_call_handler(*handle_id , obj , event_info);
}

void cgo_evas_object_event_callback_add(Eo *obj, Evas_Callback_Type type, void *data) {
	evas_object_event_callback_add(obj, type, cgo_evas_object_event_callback_func, data);
}

void cgo_evas_event_callback_func(void *data, Evas *e, void *event_info) {
	int *handle_id = (int *)data;
	go_call_handler(*handle_id , e , event_info);
	//go_call_evas_event_handler(*handle_id, e, event_info);
}

void cgo_evas_event_callback_add(Evas *e, Evas_Callback_Type type,void *data) {
	evas_event_callback_add(e, type, cgo_evas_event_callback_func, data);
}

void cgo_evas_event_callback_priority_add(Evas *e, Evas_Callback_Type type, Evas_Callback_Priority priority, void *data) {
	evas_event_callback_priority_add(e, type, priority, cgo_evas_event_callback_func, data);
}

char **cgo_null_term_string_get(char *str) {
	char **data = (char **)malloc(sizeof(char *) * 2);
	data[0] = str;
	data[1] = NULL;
	return data;
}

// array must be NULL terminated
Eina_Iterator *cgo_carray_iterator_new(char **array) {
	return eina_carray_iterator_new((void**)array);
}

Eina_List *cgo_carray_list_new(char **array, int len) {
	Eina_List *lst = NULL;
	for (int i = 0; i < len; i++) {
		lst = eina_list_append(lst,array[i]);
	}
	return lst;
}

//-------------------------------------------------------------

void cgo_elm_object_item_del_cb_func(void *data, Eo *obj, void *event_info) {
	int *id = (int *)data;
	if (id != NULL) {
		go_elm_object_item_del_cb_func(*id, obj);
	} 
}

void cgo_elm_object_item_del_cb_set(Elm_Object_Item *it) {
	elm_object_item_del_cb_set(it, cgo_elm_object_item_del_cb_func);
} 

//-------------------------------------------------------------

Eo *cgo_slideshow_item_get_func(void *data, Eo *obj) {
	int *id = (int *)data;
	return go_slideshow_item_get_func(*id,obj);
}

/*
void cgo_slideshow_item_del_func(void *data, Eo *obj) {
	int *id = (int *)data;
	go_slideshow_item_del_func(*id);
}
*/

Elm_Slideshow_Item_Class cgo_slideshow_item_class_get() {
	Elm_Slideshow_Item_Class ssic;
	ssic.func.get = cgo_slideshow_item_get_func;
	//ssic.func.del = cgo_slideshow_item_del_func;
	ssic.func.del = NULL; 
	return ssic;
}

//-------------------------------------------------------------

char *cgo_elm_gengrid_item_class_text_get(void *data, Eo *obj, const char *part) {
	int *id = (int *)data;
	char *p = (char *)part;
	return go_elm_gengrid_item_class_text_get(*id,obj,p);
}

Eo *cgo_elm_gengrid_item_class_content_get(void *data, Eo *obj, const char *part) {
	int *id = (int *)data;
	char *p = (char *)part;
	return go_elm_gengrid_item_class_content_get(*id,obj,p);
}

Eina_Bool cgo_elm_gengrid_item_class_state_get(void *data, Eo *obj, const char *part) {
	int *id = (int *)data;
	char *p = (char *)part;
	return go_elm_gengrid_item_class_state_get(*id,obj,p);
}

/*
void cgo_elm_gengrid_item_class_del(void *data, Eo *obj) {
	int *id = (int *)data;
	return go_elm_gengrid_item_class_del(*id,obj);
}
*/

// TODO item_style parametric ??
Elm_Gengrid_Item_Class * cgo_elm_gengrid_item_class_new() {
	Elm_Gengrid_Item_Class *cls = elm_gengrid_item_class_new();
	cls->item_style = "default";
	//cls->decorate_item_style = "";
	//cls->decorate_all_item_style = "";
	cls->func.text_get = cgo_elm_gengrid_item_class_text_get;
	cls->func.content_get = cgo_elm_gengrid_item_class_content_get;
	cls->func.state_get = cgo_elm_gengrid_item_class_state_get;
	//cls->func.del = cgo_elm_gengrid_item_class_del;
	return cls;
}

//-------------------------------------------------------------

char * cgo_elm_genlist_item_class_text_get(void *data, Eo *obj , const char *part) {
	int *id = (int *)data;
	char *p = (char *)part;
	return go_elm_genlist_item_class_text_get(*id,obj,p);
}

Eo * cgo_elm_genlist_item_class_content_get(void *data, Eo *obj, const char *part) {
	int *id = (int *)data;
	char *p = (char *)part;
	return go_elm_genlist_item_class_content_get(*id, obj, p);
}

Eo * cgo_elm_genlist_item_class_reusable_content_get (void *data, Eo *obj, const char *part, Eo *old) {
	int *id = (int *)data;
	char *p = (char *)part;
	return go_elm_genlist_item_class_reusable_content_get(*id,obj,p,old); 
}

Eina_Bool cgo_elm_genlist_item_class_state_get (void *data, Eo *obj, const char *part) {
	int *id = (int *)data;
	char *p = (char *)part;
	return go_elm_genlist_item_class_state_get(*id,obj,p);
}

Eina_Bool cgo_elm_genlist_item_class_filter_get(void *data, Eo *obj, void *key) {
	int *id = (int *)data;
	return go_elm_genlist_item_class_filter_get(*id,obj,key);
}

//TODO: id is deleted with cgo_elm_object_item_del_cb_set, do we need this
/*
void cgo_elm_genlist_item_class_del(void *data, Eo *obj) {
	int *id = (int *)data;
	go_elm_genlist_item_class_del(*id,obj);
}
*/

// TODO make styles parametric
Elm_Genlist_Item_Class * cgo_elm_genlist_item_class_new() {
	Elm_Genlist_Item_Class *cls = elm_genlist_item_class_new();
	cls->item_style = "default";
	//cls->decorate_item_style = "";
	//cls->decorate_all_item_style = "";
	cls->func.text_get = cgo_elm_genlist_item_class_text_get;
	cls->func.content_get = cgo_elm_genlist_item_class_content_get;
	cls->func.reusable_content_get = cgo_elm_genlist_item_class_reusable_content_get;
	cls->func.state_get = cgo_elm_genlist_item_class_state_get;
	cls->func.filter_get = cgo_elm_genlist_item_class_filter_get;
	//cls->func.del = cgo_elm_genlist_item_class_del;
	return cls;
}

//-------------------------------------------------------------

int cgo_compare_func(const void *data1 ,const void *data2) {
	int *d1 = (int *)data1;
	int *d2 = (int *)data2;
	return go_compare_func(*d1,*d2);
}

Elm_Widget_Item *go_elm_list_item_sorted_insert (Elm_List *obj, const char *label, Efl_Canvas_Object *icon, Efl_Canvas_Object *end, Evas_Smart_Cb func, const void *data) {
	return elm_list_item_sorted_insert (obj,label,icon,end,func,data, cgo_compare_func);
}

/*
void 	evas_object_event_callback_priority_add (Eo *obj, Evas_Callback_Type type, Evas_Callback_Priority priority, Eo_Event_Cb func, const void *data)
void * 	evas_object_event_callback_del (Eo *obj, Evas_Callback_Type type, Eo_Event_Cb func)
void * 	evas_object_event_callback_del_full (Eo *obj, Evas_Callback_Type type, Eo_Event_Cb func, const void *data)
*/

//-------------------------------------------------------------

void cgo_ecore_async_callback(void *data) {
	int *id = (int *)data;
	go_call_func_async(*id);
}

void cgo_ecore_main_loop_thread_safe_call_async(void *data) {
	ecore_main_loop_thread_safe_call_async(cgo_ecore_async_callback,data);
}

void *cgo_ecore_sync_callback(void *data) {
	int *id = (int *)data;
	go_call_func_sync(*id);
	return NULL;
}

void cgo_ecore_main_loop_thread_safe_call_sync(void *data) {
	ecore_main_loop_thread_safe_call_sync(cgo_ecore_sync_callback,data);
	// TODO: this returns void *, but discarded
}

//-------------------------------------------------------------

void cgo_ecore_callback_func(void *data) {
	int *id = (int *)data;
	go_call_ecore_callback_func(*id);
}

Ecore_Job *cgo_ecore_job_add(void *data) {
	return ecore_job_add(cgo_ecore_callback_func,data);
}

//-------------------------------------------------------------

Eina_Bool cgo_ecore_task_callback_func(void *data) {
	int *id = (int *)data;
	return go_ecore_task_callback_func(*id);
}

Ecore_Timer * cgo_ecore_timer_add (double in, const void *data) {
	return ecore_timer_add (in, cgo_ecore_task_callback_func, data);
}

Ecore_Timer * cgo_ecore_timer_loop_add (double in, const void *data) {
	return ecore_timer_add (in, cgo_ecore_task_callback_func, data);
}

//-------------------------------------------------------------

void cgo_web_console_message_func(void *data, Eo *obj, const char *message, unsigned int line_number, const char *source_id) {
	int *id = (int *)data;
	char *msg = (char *)message;
	char *srcid = (char *)source_id;
	go_web_console_message_func(*id, obj, msg, line_number, srcid);
}

void cgo_web_console_message_hook_set(Elm_Web *obj, void *data) {
	elm_web_console_message_hook_set(obj, cgo_web_console_message_func, data);
}

Eo *cgo_web_window_open_func(void *data, Eo *obj, Eina_Bool js, const Elm_Web_Window_Features *window_features) {
	int *id = (int *)data;
	Elm_Web_Window_Features *wf = (Elm_Web_Window_Features *)window_features;
	return go_web_window_open_func(*id, obj, js, wf);
}

void cgo_web_window_create_hook_set(Elm_Web *obj, void *data) {
	elm_web_window_create_hook_set (obj, cgo_web_window_open_func, data);
}

/* TODO
Eo *cgo_web_dialog_file_selector_func(void *data, Eo *obj, Eina_Bool always_multiple, Eina_List accept_types, Eina_List **selected, Eina_Bool *ret) {
	int *id = (int *)data;
	return go_web_dialog_file_selector_func(*id, obj, always_multiple, accept_types, selected, ret);
}
*/

Eo *cgo_web_dialog_confirm_func(void *data, Eo *obj, const char *message, Eina_Bool *ret) {
	int *id = (int *)data;
	char *msg = (char *)message;
	return go_web_dialog_confirm_func(*id, obj, msg, ret);
}

void cgo_web_dialog_confirm_hook_set(Elm_Web *obj, void *data) {
	elm_web_dialog_confirm_hook_set(obj, cgo_web_dialog_confirm_func, data);
}

Eo *cgo_web_dialog_prompt_func(void *data, Eo *obj, const char *message, const char *def_value, const char **value, Eina_Bool *ret) {
	int *id = (int *)data;
	char *msg = (char *)message;
	char *dv = (char *)def_value;
	char **val = (char **)value;
	return go_web_dialog_prompt_func(*id, obj, msg, dv, val, ret);
}

void cgo_web_dialog_prompt_hook_set(Elm_Web *obj, void *data) {
	elm_web_dialog_prompt_hook_set(obj, cgo_web_dialog_prompt_func, data);
}

Eo *cgo_web_dialog_alert_func(void *data, Eo *obj, const char *message) {
	int *id = (int *)data;
	char *msg = (char *)message;
	return go_web_dialog_alert_func(*id, obj, msg);
}

void cgo_web_dialog_alert_hook_set(Elm_Web *obj, void *data) {
	elm_web_dialog_alert_hook_set(obj, cgo_web_dialog_alert_func, data);
}

//-------------------------------------------------------------

Eina_Bool cgo_elm_drop_cb_func(void *data, Eo *obj, Elm_Selection_Data *ev) {
	return EINA_FALSE; 
}

Eo *cgo_elm_drag_icon_create_cb_func(void *data, Eo *win, Evas_Coord *xoff, Evas_Coord *yoff) {
	return NULL;		
}

void cgo_elm_drag_state_func(void *data, Eo *obj) {

}

void cgo_elm_drag_done_func(void *data, Eo *obj, Eina_Bool accepted) {

}

void cgo_elm_drag_accept_func(void *data, Eo *obj, Eina_Bool doaccept) {

}

void cgo_elm_drag_pos_func(void *data, Eo *obj, Evas_Coord x, Evas_Coord y, Elm_Xdnd_Action action) {

}

void cgo_elm_drag_start_func(void *data, Eo *obj) {

}

void cgo_elm_drag_item_container_pos_func(void *data, Eo *cont, Elm_Object_Item *it, Evas_Coord x, Evas_Coord y, int xposret, int yposret, Elm_Xdnd_Action action) {

}

Eina_Bool cgo_elm_drop_item_container_cb_func(void *data, Eo *obj, Elm_Object_Item *it, Elm_Selection_Data *ev, int xposret, int yposret) {
	return EINA_FALSE; 
}

/*
Eina_Bool elm_cnp_selection_set(Eo *obj, Elm_Sel_Type selection, Elm_Sel_Format format, const void *buf, size_t buflen);
we cast void * to char *
*/
Eina_Bool cgo_elm_cnp_selection_set(Eo *obj, Elm_Sel_Type selection, Elm_Sel_Format format, char *buf) {
	return elm_cnp_selection_set(obj, selection, format, buf, strlen(buf));
}

// normally elm_drop_cb
Eina_Bool cgo_elm_selection_cb_func(void *data, Eo *obj, Elm_Selection_Data *ev) {
	int *id = (int *)data;
	return go_elm_selection_cb_func(*id, obj, ev);
}

Eina_Bool cgo_elm_cnp_selection_get(const Eo *obj, Elm_Sel_Type selection, Elm_Sel_Format format, void *data) {
	/*
	Elm_Drop_Cb cb = NULL;
	if (data != NULL) {
		cb = cgo_elm_selection_cb_func;
	}
	return elm_cnp_selection_get (obj, selection, format, cb, data);
	*/
	return elm_cnp_selection_get (obj, selection, format, cgo_elm_selection_cb_func, data);
}

void cgo_elm_selection_loss_cb_func(void *data, Elm_Sel_Type selection) {
	int *id = (int *)data;
	go_elm_selection_loss_cb_func(*id,selection);
}

void cgo_elm_selection_loss_callback_set(Eo *obj, Elm_Sel_Type selection,void *data) {
	elm_cnp_selection_loss_callback_set (obj, selection, cgo_elm_selection_loss_cb_func, data);
}

Elm_Widget_Item *cgo_elm_xy_item_get_cb_func(Eo *obj, Evas_Coord x, Evas_Coord y, int *xposret, int *yposret) {
	return go_elm_xy_item_get_cb_func(obj, x, y, xposret, yposret);
}

/* TODO
Eina_Bool cgo_elm_item_container_data_get_cb_func(Eo *obj, Elm_Object_Item *it, Elm_Drag_User_Info *info) {
	return EINA_FALSE;
}

Eina_Bool cgo_elm_drag_item_container_add(Eo *obj, double tm_to_anim, double tm_to_drag, Eina_Bool usedatagetcb) {
	Elm_Item_Container_Data_Get_Cb datagetcb = NULL;
	
	if (usedatagetcb) {
		datagetcb = cgo_elm_item_container_data_get_cb_func;
	}

	elm_drag_item_container_add(obj, tm_to_anim, tm_to_drag, cgo_elm_xy_item_get_cb_func, datagetcb);
}
*/

void cgo_elm_win_wm_rotation_available_rotations_set(Elm_Win *obj, Eina_List *list) {
	unsigned int count = (unsigned int)eina_list_count(list);
	int * rotations = malloc(sizeof(int) * count);
	void *list_data;
	int i = 0;
	Eina_List *l;
	EINA_LIST_FOREACH(list, l, list_data) {
		rotations[i++] = *((int *)list_data);
	}
	elm_win_wm_rotation_available_rotations_set(obj, (const int *)rotations, count);
	eina_list_free(list);
	free(rotations);
}

Eina_List *cgo_elm_win_wm_rotation_available_rotations_get(const Elm_Win *obj) {
	int * rotations = NULL;
	unsigned int count;
	Eina_List *list = NULL;

	Eina_Bool ret = elm_win_wm_rotation_available_rotations_get(obj, &rotations, &count);
	if (ret != EINA_TRUE || count <= 0) {
		return NULL;
	}

	for (int i = 0; i < count; i++) {
		list = eina_list_append(list, &rotations[i]);
	}

	return list;
}
