package ui

/*
#include "bridge.h"
*/
import "C"

import "fmt"

const (
	HintExpand = 1.0
	HintFill   = -1.0
)

type Policy uint

const (
	PolicyQuit     Policy = C.ELM_POLICY_QUIT
	PolicyExit     Policy = C.ELM_POLICY_EXIT
	PolicyThrottle Policy = C.ELM_POLICY_THROTTLE
)

type PolicyVal int

const (
	QuitPolicyNone             PolicyVal = C.ELM_POLICY_QUIT_NONE
	QuitPolicyLastWindowClosed PolicyVal = C.ELM_POLICY_QUIT_LAST_WINDOW_CLOSED
	QuitPolicyLastWindowHidden PolicyVal = C.ELM_POLICY_QUIT_LAST_WINDOW_HIDDEN
	ExitPolicyNone             PolicyVal = C.ELM_POLICY_EXIT_NONE
	ExitPolicyWindowsDel       PolicyVal = C.ELM_POLICY_EXIT_WINDOWS_DEL
	ThrottlePolicyConfig       PolicyVal = C.ELM_POLICY_THROTTLE_CONFIG
	ThrottlePolicyHiddenAlways PolicyVal = C.ELM_POLICY_THROTTLE_HIDDEN_ALWAYS
	ThrottlePolicyNever        PolicyVal = C.ELM_POLICY_THROTTLE_NEVER
)

type AspectControl int

const (
	AspectControlNone       AspectControl = C.EVAS_ASPECT_CONTROL_NONE
	AspectControlNeither    AspectControl = C.EVAS_ASPECT_CONTROL_NEITHER
	AspectControlHorizontal AspectControl = C.EVAS_ASPECT_CONTROL_HORIZONTAL
	AspectControlVertical   AspectControl = C.EVAS_ASPECT_CONTROL_VERTICAL
	AspectControlBoth       AspectControl = C.EVAS_ASPECT_CONTROL_BOTH
)

type LabelSlideMode int

const (
	LabelSlideModeAlways LabelSlideMode = C.ELM_LABEL_SLIDE_MODE_ALWAYS
	LabelSlideModeAuto   LabelSlideMode = C.ELM_LABEL_SLIDE_MODE_AUTO
)

type WrapType int

const (
	WrapTypeNone  WrapType = C.ELM_WRAP_NONE
	WrapTypeChar  WrapType = C.ELM_WRAP_CHAR
	WrapTypeWord  WrapType = C.ELM_WRAP_WORD
	WrapTypeMixed WrapType = C.ELM_WRAP_MIXED

//	WrapTypeLast  WrapType = C.ELM_WRAP_LAST
)

type WindowType int

const (
	WindowTypeDialogBasic  WindowType = C.ELM_WIN_DIALOG_BASIC
	WindowTypeDesktop      WindowType = C.ELM_WIN_DESKTOP
	WindowTypeDock         WindowType = C.ELM_WIN_DOCK
	WindowTypeToolbar      WindowType = C.ELM_WIN_TOOLBAR
	WindowTypeMenu         WindowType = C.ELM_WIN_MENU
	WindowTypeUtility      WindowType = C.ELM_WIN_UTILITY
	WindowTypeSplash       WindowType = C.ELM_WIN_SPLASH
	WindowTypeDropdownMenu WindowType = C.ELM_WIN_DROPDOWN_MENU
	WindowTypePopupMenu    WindowType = C.ELM_WIN_POPUP_MENU
	WindowTypeTooltip      WindowType = C.ELM_WIN_TOOLTIP
	WindowTypeNotification WindowType = C.ELM_WIN_NOTIFICATION
	WindowTypeCombo        WindowType = C.ELM_WIN_COMBO
	WindowTypeDND          WindowType = C.ELM_WIN_DND
	WindowTypeInlinedImage WindowType = C.ELM_WIN_INLINED_IMAGE
	WindowTypeSocketImage  WindowType = C.ELM_WIN_SOCKET_IMAGE
	WindowTypeFake         WindowType = C.ELM_WIN_FAKE
)

type WindowKeyboardMode int

const (
	WindowKeyboardModeUnknown     WindowKeyboardMode = C.ELM_WIN_KEYBOARD_UNKNOWN
	WindowKeyboardModeOff         WindowKeyboardMode = C.ELM_WIN_KEYBOARD_OFF
	WindowKeyboardModeOn          WindowKeyboardMode = C.ELM_WIN_KEYBOARD_ON
	WindowKeyboardModeAlpha       WindowKeyboardMode = C.ELM_WIN_KEYBOARD_ALPHA
	WindowKeyboardModeNumeric     WindowKeyboardMode = C.ELM_WIN_KEYBOARD_NUMERIC
	WindowKeyboardModePin         WindowKeyboardMode = C.ELM_WIN_KEYBOARD_PIN
	WindowKeyboardModePhoneNumber WindowKeyboardMode = C.ELM_WIN_KEYBOARD_PHONE_NUMBER
	WindowKeyboardModeHex         WindowKeyboardMode = C.ELM_WIN_KEYBOARD_HEX
	WindowKeyboardModeTerminal    WindowKeyboardMode = C.ELM_WIN_KEYBOARD_TERMINAL
	WindowKeyboardModePassword    WindowKeyboardMode = C.ELM_WIN_KEYBOARD_PASSWORD
	WindowKeyboardModeIP          WindowKeyboardMode = C.ELM_WIN_KEYBOARD_IP
	WindowKeyboardModeHost        WindowKeyboardMode = C.ELM_WIN_KEYBOARD_HOST
	WindowKeyboardModeFile        WindowKeyboardMode = C.ELM_WIN_KEYBOARD_FILE
	WindowKeyboardModeURL         WindowKeyboardMode = C.ELM_WIN_KEYBOARD_URL
	WindowKeyboardModeKeypad      WindowKeyboardMode = C.ELM_WIN_KEYBOARD_KEYPAD
	WindowKeyboardModeJ2ME        WindowKeyboardMode = C.ELM_WIN_KEYBOARD_J2ME
)

type WindowIndicatorOpacityMode int

const (
	WindowIndicatorOpacityModeUnknown     WindowIndicatorOpacityMode = C.ELM_WIN_INDICATOR_OPACITY_UNKNOWN
	WindowIndicatorOpacityModeOpaque      WindowIndicatorOpacityMode = C.ELM_WIN_INDICATOR_OPAQUE
	WindowIndicatorOpacityModeTranslucent WindowIndicatorOpacityMode = C.ELM_WIN_INDICATOR_TRANSLUCENT
	WindowIndicatorOpacityModeTransparent WindowIndicatorOpacityMode = C.ELM_WIN_INDICATOR_TRANSPARENT
)

type WindowIndicatorMode int

const (
	WindowIndicatorModeUnknown WindowIndicatorMode = C.ELM_WIN_INDICATOR_UNKNOWN
	WindowIndicatorModeHide    WindowIndicatorMode = C.ELM_WIN_INDICATOR_HIDE
	WindowIndicatorModeShow    WindowIndicatorMode = C.ELM_WIN_INDICATOR_SHOW
)

type WindowKeygrabMode int

const (
	WindowKeygrabModeUnknown           WindowKeygrabMode = C.ELM_WIN_KEYGRAB_UNKNOWN
	WindowKeygrabModeShared            WindowKeygrabMode = C.ELM_WIN_KEYGRAB_SHARED
	WindowKeygrabModeTopmost           WindowKeygrabMode = C.ELM_WIN_KEYGRAB_TOPMOST
	WindowKeygrabModeExclusive         WindowKeygrabMode = C.ELM_WIN_KEYGRAB_EXCLUSIVE
	WindowKeygrabModeOverrideExclusive WindowKeygrabMode = C.ELM_WIN_KEYGRAB_OVERRIDE_EXCLUSIVE
)

type BgOption int

const (
	BgOptionCenter  BgOption = C.ELM_BG_OPTION_CENTER
	BgOptionScale   BgOption = C.ELM_BG_OPTION_SCALE
	BgOptionStretch BgOption = C.ELM_BG_OPTION_STRETCH
	BgOptionTile    BgOption = C.ELM_BG_OPTION_TILE

//	BgOptionLast    BgOption = C.ELM_BG_OPTION_LAST
)

type BubblePos int

const (
	BubblePosTopLeft     BubblePos = C.ELM_BUBBLE_POS_TOP_LEFT
	BubblePosTopRight    BubblePos = C.ELM_BUBBLE_POS_TOP_RIGHT
	BubblePosBottomLeft  BubblePos = C.ELM_BUBBLE_POS_BOTTOM_LEFT
	BubblePosBottomRight BubblePos = C.ELM_BUBBLE_POS_BOTTOM_RIGHT
)

type ImageOrient int

const (
	ImageOrientNone     ImageOrient = C.ELM_IMAGE_ORIENT_NONE
	ImageOrient0        ImageOrient = C.ELM_IMAGE_ORIENT_0
	ImageRotate90       ImageOrient = C.ELM_IMAGE_ROTATE_90
	ImageRotate180      ImageOrient = C.ELM_IMAGE_ROTATE_180
	ImageRotate270      ImageOrient = C.ELM_IMAGE_ROTATE_270
	ImageFlipHorizontal ImageOrient = C.ELM_IMAGE_FLIP_HORIZONTAL
	ImageFlipVertical   ImageOrient = C.ELM_IMAGE_FLIP_VERTICAL
	ImageFlipTranspose  ImageOrient = C.ELM_IMAGE_FLIP_TRANSPOSE
	ImageFlipTransverse ImageOrient = C.ELM_IMAGE_FLIP_TRANSVERSE
)

type ActionsliderPos int

const (
	ActionsliderPosLeft   ActionsliderPos = C.ELM_ACTIONSLIDER_LEFT
	ActionsliderPosCenter ActionsliderPos = C.ELM_ACTIONSLIDER_CENTER
	ActionsliderPosRight  ActionsliderPos = C.ELM_ACTIONSLIDER_RIGHT
	ActionsliderPosAll    ActionsliderPos = C.ELM_ACTIONSLIDER_ALL
)

type ScrollerPolicy int

const (
	ScrollerPolicyOn  ScrollerPolicy = C.ELM_SCROLLER_POLICY_ON
	ScrollerPolicyOff ScrollerPolicy = C.ELM_SCROLLER_POLICY_OFF

//	ScrollerPolicyLast ScrollerPolicy = C.ELM_SCROLLER_POLICY_LAST
)

type ScrollerSingleDirection int

const (
	ScrollerSingleDirectionSoft ScrollerSingleDirection = C.ELM_SCROLLER_SINGLE_DIRECTION_SOFT
	ScrollerSingleDirectionHard ScrollerSingleDirection = C.ELM_SCROLLER_SINGLE_DIRECTION_HARD

//	ScrollerSingleDirectionLast ScrollerSingleDirection = C.ELM_SCROLLER_SINGLE_DIRECTION_LAST
)

type ScrollerMovementBlock int

const (
	ScrollerMovementBlockVertical   ScrollerMovementBlock = C.ELM_SCROLLER_MOVEMENT_BLOCK_VERTICAL
	ScrollerMovementBlockHorizontal ScrollerMovementBlock = C.ELM_SCROLLER_MOVEMENT_BLOCK_HORIZONTAL
)

type TooltipOrient int

const (
	TooltipOrientNone        TooltipOrient = C.ELM_TOOLTIP_ORIENT_NONE
	TooltipOrientTopLeft     TooltipOrient = C.ELM_TOOLTIP_ORIENT_TOP_LEFT
	TooltipOrientTop         TooltipOrient = C.ELM_TOOLTIP_ORIENT_TOP
	TooltipOrientTopRight    TooltipOrient = C.ELM_TOOLTIP_ORIENT_TOP_RIGHT
	TooltipOrientLeft        TooltipOrient = C.ELM_TOOLTIP_ORIENT_LEFT
	TooltipOrientCenter      TooltipOrient = C.ELM_TOOLTIP_ORIENT_CENTER
	TooltipOrientRight       TooltipOrient = C.ELM_TOOLTIP_ORIENT_RIGHT
	TooltipOrientBottomLeft  TooltipOrient = C.ELM_TOOLTIP_ORIENT_BOTTOM_LEFT
	TooltipOrientBottom      TooltipOrient = C.ELM_TOOLTIP_ORIENT_BOTTOM
	TooltipOrientBottomRight TooltipOrient = C.ELM_TOOLTIP_ORIENT_BOTTOM_RIGHT
	//TooltipOrient TooltipOrient = C.ELM_TOOLTIP_ORIENT_LAST
)

type ClockEditMode int

const (
	ClockEditModeHourDecimal ClockEditMode = C.ELM_CLOCK_EDIT_HOUR_DECIMAL
	ClockEditModeHourUnit    ClockEditMode = C.ELM_CLOCK_EDIT_HOUR_UNIT
	ClockEditModeMinDecimal  ClockEditMode = C.ELM_CLOCK_EDIT_MIN_DECIMAL
	ClockEditModeMinUnit     ClockEditMode = C.ELM_CLOCK_EDIT_MIN_UNIT
	ClockEditModeSecDecimal  ClockEditMode = C.ELM_CLOCK_EDIT_SEC_DECIMAL
	ClockEditModeSecUnit     ClockEditMode = C.ELM_CLOCK_EDIT_SEC_UNIT
	ClockEditModeAll         ClockEditMode = C.ELM_CLOCK_EDIT_ALL
)

type ColorselectorMode int

const (
	ColorselectorModeComponents ColorselectorMode = C.ELM_COLORSELECTOR_COMPONENTS
	ColorselectorModeBoth       ColorselectorMode = C.ELM_COLORSELECTOR_BOTH
	ColorselectorModePicker     ColorselectorMode = C.ELM_COLORSELECTOR_PICKER
	ColorselectorModeAll        ColorselectorMode = C.ELM_COLORSELECTOR_ALL
)

type DayselectorDay int

const (
	DayselectorDaySun DayselectorDay = C.ELM_DAYSELECTOR_SUN
	DayselectorDayMon DayselectorDay = C.ELM_DAYSELECTOR_MON
	DayselectorDayTue DayselectorDay = C.ELM_DAYSELECTOR_TUE
	DayselectorDayWed DayselectorDay = C.ELM_DAYSELECTOR_WED
	DayselectorDayThu DayselectorDay = C.ELM_DAYSELECTOR_THU
	DayselectorDayFri DayselectorDay = C.ELM_DAYSELECTOR_FRI
	DayselectorDaySat DayselectorDay = C.ELM_DAYSELECTOR_SAT
	DayselectorDayMax DayselectorDay = C.ELM_DAYSELECTOR_MAX
)

type FileselectorMode int

const (
	FileselectorModeList FileselectorMode = C.ELM_FILESELECTOR_LIST
	FileselectorModeGrid FileselectorMode = C.ELM_FILESELECTOR_GRID
	//FileselectorMode FileselectorMode = C.ELM_FILESELECTOR_LAST
)

type FileselectorSort int

const (
	FileselectorSortFilenameAsc  FileselectorSort = C.ELM_FILESELECTOR_SORT_BY_FILENAME_ASC
	FileselectorSortFilenameDesc FileselectorSort = C.ELM_FILESELECTOR_SORT_BY_FILENAME_DESC
	FileselectorSortTypeAsc      FileselectorSort = C.ELM_FILESELECTOR_SORT_BY_TYPE_ASC
	FileselectorSortTypeDesc     FileselectorSort = C.ELM_FILESELECTOR_SORT_BY_TYPE_DESC
	FileselectorSortSizeAsc      FileselectorSort = C.ELM_FILESELECTOR_SORT_BY_SIZE_ASC
	FileselectorSortSizeDesc     FileselectorSort = C.ELM_FILESELECTOR_SORT_BY_SIZE_DESC
	FileselectorSortModifiedAsc  FileselectorSort = C.ELM_FILESELECTOR_SORT_BY_MODIFIED_ASC
	FileselectorSortModifiedDesc FileselectorSort = C.ELM_FILESELECTOR_SORT_BY_MODIFIED_DESC

//	FileselectorSort FileselectorSort = C.ELM_FILESELECTOR_SORT_LAST
)

type CalendarWeekday int

const (
	CalendarWeekdaySunday    CalendarWeekday = C.ELM_DAY_SUNDAY
	CalendarWeekdayMonday    CalendarWeekday = C.ELM_DAY_MONDAY
	CalendarWeekdayTuesday   CalendarWeekday = C.ELM_DAY_TUESDAY
	CalendarWeekdayWednesday CalendarWeekday = C.ELM_DAY_WEDNESDAY
	CalendarWeekdayThursday  CalendarWeekday = C.ELM_DAY_THURSDAY
	CalendarWeekdayFriday    CalendarWeekday = C.ELM_DAY_FRIDAY
	CalendarWeekdaySaturday  CalendarWeekday = C.ELM_DAY_SATURDAY
)

type CalendarSelectable int

const (
	CalendarSelectableNone  CalendarSelectable = C.ELM_CALENDAR_SELECTABLE_NONE
	CalendarSelectableYear  CalendarSelectable = C.ELM_CALENDAR_SELECTABLE_YEAR
	CalendarSelectableMonth CalendarSelectable = C.ELM_CALENDAR_SELECTABLE_MONTH
	CalendarSelectableDay   CalendarSelectable = C.ELM_CALENDAR_SELECTABLE_DAY
)

type CalendarSelectMode int

const (
	CalendarSelectModeDefault  CalendarSelectMode = C.ELM_CALENDAR_SELECT_MODE_DEFAULT
	CalendarSelectModeAlways   CalendarSelectMode = C.ELM_CALENDAR_SELECT_MODE_ALWAYS
	CalendarSelectModeNone     CalendarSelectMode = C.ELM_CALENDAR_SELECT_MODE_NONE
	CalendarSelectModeOnDemand CalendarSelectMode = C.ELM_CALENDAR_SELECT_MODE_ONDEMAND
)

type CalendarMarkRepeatType int

const (
	CalendarMarkRepeatTypeUnique         CalendarMarkRepeatType = C.ELM_CALENDAR_UNIQUE
	CalendarMarkRepeatTypeDaily          CalendarMarkRepeatType = C.ELM_CALENDAR_DAILY
	CalendarMarkRepeatTypeWeekly         CalendarMarkRepeatType = C.ELM_CALENDAR_WEEKLY
	CalendarMarkRepeatTypeMonthly        CalendarMarkRepeatType = C.ELM_CALENDAR_MONTHLY
	CalendarMarkRepeatTypeAnnually       CalendarMarkRepeatType = C.ELM_CALENDAR_ANNUALLY
	CalendarMarkRepeatTypeLastDayOfMonth CalendarMarkRepeatType = C.ELM_CALENDAR_LAST_DAY_OF_MONTH
	// in 1.19
	//CalendarMarkRepeatTypeReverseDaily   CalendarMarkRepeatType = C.ELM_CALENDAR_REVERSE_DAILY
)

type ButtonFlags int

const (
	ButtonFlagsNone        ButtonFlags = C.EVAS_BUTTON_NONE
	ButtonFlagsDoubleClick ButtonFlags = C.EVAS_BUTTON_DOUBLE_CLICK
	ButtonFlagsTripleClick ButtonFlags = C.EVAS_BUTTON_TRIPLE_CLICK
)

type FontHinting int

const (
	FontHintingNone     FontHinting = C.EVAS_FONT_HINTING_NONE
	FontHintingAuto     FontHinting = C.EVAS_FONT_HINTING_AUTO
	FontHintingByteCode FontHinting = C.EVAS_FONT_HINTING_BYTECODE
)

type TableHomogeneousMode int

const (
	TableHomogeneousModeNone  TableHomogeneousMode = C.EVAS_OBJECT_TABLE_HOMOGENEOUS_NONE
	TableHomogeneousModeTable TableHomogeneousMode = C.EVAS_OBJECT_TABLE_HOMOGENEOUS_TABLE
	TableHomogeneousModeItem  TableHomogeneousMode = C.EVAS_OBJECT_TABLE_HOMOGENEOUS_ITEM
)

type TextStyle int

const (
	TextStylePlain                      TextStyle = C.EVAS_TEXT_STYLE_PLAIN
	TextStyleShadow                     TextStyle = C.EVAS_TEXT_STYLE_SHADOW
	TextStyleOutline                    TextStyle = C.EVAS_TEXT_STYLE_OUTLINE
	TextStyleSoftOutline                TextStyle = C.EVAS_TEXT_STYLE_SOFT_OUTLINE
	TextStyleGlow                       TextStyle = C.EVAS_TEXT_STYLE_GLOW
	TextStyleeOutlineShadow             TextStyle = C.EVAS_TEXT_STYLE_OUTLINE_SHADOW
	TextStyleFarShadow                  TextStyle = C.EVAS_TEXT_STYLE_FAR_SHADOW
	TextStyleOutlineSoftShadow          TextStyle = C.EVAS_TEXT_STYLE_OUTLINE_SOFT_SHADOW
	TextStyleSoftShadow                 TextStyle = C.EVAS_TEXT_STYLE_SOFT_SHADOW
	TextStyleFarSoftShadow              TextStyle = C.EVAS_TEXT_STYLE_FAR_SOFT_SHADOW
	TextStyleShadowDirectionBottomRight TextStyle = C.EVAS_TEXT_STYLE_SHADOW_DIRECTION_BOTTOM_RIGHT
	TextStyleShadowDirectionBottom      TextStyle = C.EVAS_TEXT_STYLE_SHADOW_DIRECTION_BOTTOM
	TextStyleShadowDirectionBottomLeft  TextStyle = C.EVAS_TEXT_STYLE_SHADOW_DIRECTION_BOTTOM_LEFT
	TextStyleShadowDirectionLeft        TextStyle = C.EVAS_TEXT_STYLE_SHADOW_DIRECTION_LEFT
	TextStyleShadowDirectionTopLeft     TextStyle = C.EVAS_TEXT_STYLE_SHADOW_DIRECTION_TOP_LEFT
	TextStyleShadowDirectionTop         TextStyle = C.EVAS_TEXT_STYLE_SHADOW_DIRECTION_TOP
	TextStyleShadowDirectionTopRight    TextStyle = C.EVAS_TEXT_STYLE_SHADOW_DIRECTION_TOP_RIGHT
	TextStyleShadowDirectionRight       TextStyle = C.EVAS_TEXT_STYLE_SHADOW_DIRECTION_RIGHT
)

type BorderFillMode int

const (
	BorderFillModeNone      BorderFillMode = C.EVAS_BORDER_FILL_NONE
	BorderFillModeDefault   BorderFillMode = C.EVAS_BORDER_FILL_DEFAULT
	BorderFillModeFillSolid BorderFillMode = C.EVAS_BORDER_FILL_SOLID
)

type LoadError int

const (
	LoadErrorNone             LoadError = C.EVAS_LOAD_ERROR_NONE
	LoadErrorGeneric          LoadError = C.EVAS_LOAD_ERROR_GENERIC
	LoadErrorDoesNotExist     LoadError = C.EVAS_LOAD_ERROR_DOES_NOT_EXIST
	LoadErrorPermissionDenied LoadError = C.EVAS_LOAD_ERROR_PERMISSION_DENIED
	LoadErrorAllocationFailed LoadError = C.EVAS_LOAD_ERROR_RESOURCE_ALLOCATION_FAILED
	LoadErrorCorruptFile      LoadError = C.EVAS_LOAD_ERROR_CORRUPT_FILE
	LoadErrorUnknownFormat    LoadError = C.EVAS_LOAD_ERROR_UNKNOWN_FORMAT
)

//TODO: const char *evas_load_error_str (Evas_Load_Error error)
func (e LoadError) Error() string {
	switch e {
	case LoadErrorNone:
		return ""
	case LoadErrorGeneric:
		return "Generic Evas Load Error"
	case LoadErrorDoesNotExist:
		return "Does not exist"
	case LoadErrorPermissionDenied:
		return "Permission Denied"
	case LoadErrorAllocationFailed:
		return "Allocation Failed"
	case LoadErrorCorruptFile:
		return "File Corrupt"
	case LoadErrorUnknownFormat:
		return "Unknown Format"
	}
	return "Unknown Evas Load Error"
}

type ImageScaleHint int

const (
	ImageScaleHintNone    ImageScaleHint = C.EVAS_IMAGE_SCALE_HINT_NONE
	ImageScaleHintDynamic ImageScaleHint = C.EVAS_IMAGE_SCALE_HINT_DYNAMIC
	ImageScaleHintStatic  ImageScaleHint = C.EVAS_IMAGE_SCALE_HINT_STATIC
)

type Colorspace int

const (
	ColorspaceARGB8888        Colorspace = C.EVAS_COLORSPACE_ARGB8888
	ColorspaceYCBCR422P601_PL Colorspace = C.EVAS_COLORSPACE_YCBCR422P601_PL
	ColorspaceYCBCR422P709_PL Colorspace = C.EVAS_COLORSPACE_YCBCR422P709_PL
	ColorspaceRGB565_A5P      Colorspace = C.EVAS_COLORSPACE_RGB565_A5P
	ColorspaceGRY8            Colorspace = C.EVAS_COLORSPACE_GRY8
)

type ImageContentHint int

const (
	ImageContentHintNone    ImageContentHint = C.EVAS_IMAGE_CONTENT_HINT_NONE
	ImageContentHintDynamic ImageContentHint = C.EVAS_IMAGE_CONTENT_HINT_DYNAMIC
	ImageContentHintStatic  ImageContentHint = C.EVAS_IMAGE_CONTENT_HINT_STATIC
)

type ImageAnimatedLoopHint int

const (
	ImageAnimatedLoopHintLoop     ImageAnimatedLoopHint = C.EVAS_IMAGE_ANIMATED_HINT_LOOP
	ImageAnimatedLoopHintPingPong ImageAnimatedLoopHint = C.EVAS_IMAGE_ANIMATED_HINT_PINGPONG
)

type FocusAutoscrollModeType int

const (
	FocusAutoscrollModeTypeShow    FocusAutoscrollModeType = C.ELM_FOCUS_AUTOSCROLL_MODE_SHOW
	FocusAutoscrollModeTypeNone    FocusAutoscrollModeType = C.ELM_FOCUS_AUTOSCROLL_MODE_NONE
	FocusAutoscrollModeTypeBringIn FocusAutoscrollModeType = C.ELM_FOCUS_AUTOSCROLL_MODE_BRING_IN
)

type FocusMovePolicyType int

const (
	FocusMovePolicyIn      FocusMovePolicyType = C.ELM_FOCUS_MOVE_POLICY_IN
	FocusMovePolicyClick   FocusMovePolicyType = C.ELM_FOCUS_MOVE_POLICY_CLICK
	FocusMovePolicyKeyOnly FocusMovePolicyType = C.ELM_FOCUS_MOVE_POLICY_KEY_ONLY
)

type TransitTweenMode int

const (
	TransitTweenModeLinear        TransitTweenMode = C.ELM_TRANSIT_TWEEN_MODE_LINEAR
	TransitTweenModeSinusoidal    TransitTweenMode = C.ELM_TRANSIT_TWEEN_MODE_SINUSOIDAL
	TransitTweenModeDecelerate    TransitTweenMode = C.ELM_TRANSIT_TWEEN_MODE_DECELERATE
	TransitTweenModeAccelerate    TransitTweenMode = C.ELM_TRANSIT_TWEEN_MODE_ACCELERATE
	TransitTweenModeDivisorInterp TransitTweenMode = C.ELM_TRANSIT_TWEEN_MODE_DIVISOR_INTERP
	TransitTweenModeBounce        TransitTweenMode = C.ELM_TRANSIT_TWEEN_MODE_BOUNCE
	TransitTweenModeSpring        TransitTweenMode = C.ELM_TRANSIT_TWEEN_MODE_SPRING
	TransitTweenModeBezierCurve   TransitTweenMode = C.ELM_TRANSIT_TWEEN_MODE_BEZIER_CURVE
)

type TransitEffectFlipAxis int

const (
	TransitEffectFlipAxisX TransitEffectFlipAxis = C.ELM_TRANSIT_EFFECT_FLIP_AXIS_X
	TransitEffectFlipAxisY TransitEffectFlipAxis = C.ELM_TRANSIT_EFFECT_FLIP_AXIS_Y
)

type TransitEffectWipeType int

const (
	TransitEffectWipeTypeHide TransitEffectWipeType = C.ELM_TRANSIT_EFFECT_WIPE_TYPE_HIDE
	TransitEffectWipeTypeShow TransitEffectWipeType = C.ELM_TRANSIT_EFFECT_WIPE_TYPE_SHOW
)

type TransitEffectWipeDir int

const (
	TransitEffectWipeDirLeft  TransitEffectWipeDir = C.ELM_TRANSIT_EFFECT_WIPE_DIR_LEFT
	TransitEffectWipeDirRight TransitEffectWipeDir = C.ELM_TRANSIT_EFFECT_WIPE_DIR_RIGHT
	TransitEffectWipeDirUp    TransitEffectWipeDir = C.ELM_TRANSIT_EFFECT_WIPE_DIR_UP
	TransitEffectWipeDirDown  TransitEffectWipeDir = C.ELM_TRANSIT_EFFECT_WIPE_DIR_DOWN
)

type RenderMode int

const ( // commenteds are deprecated
	RenderModeBlend    RenderMode = C.EVAS_RENDER_BLEND
	RenderModeBlendRel RenderMode = C.EVAS_RENDER_BLEND_REL
	//RenderModeCopy     RenderMode = C.EVAS_RENDER_COPY
	RenderModeCopyRel RenderMode = C.EVAS_RENDER_COPY_REL
	//RenderModeAdd      RenderMode = C.EVAS_RENDER_ADD
	//RenderModeAddRel   RenderMode = C.EVAS_RENDER_ADD_REL
	//RenderModeSub      RenderMode = C.EVAS_RENDER_SUB
	//RenderModeSubRel   RenderMode = C.EVAS_RENDER_SUB_REL
	//RenderModeTint     RenderMode = C.EVAS_RENDER_TINT
	//RenderModeTintRel  RenderMode = C.EVAS_RENDER_TINT_REL
	//RenderModeMask     RenderMode = C.EVAS_RENDER_MASK
	//RenderModeMul      RenderMode = C.EVAS_RENDER_MUL
)

type FocusDirection int

const (
	FocusDirectionPrevious FocusDirection = C.ELM_FOCUS_PREVIOUS
	FocusDirectionNext     FocusDirection = C.ELM_FOCUS_NEXT
	FocusDirectionUp       FocusDirection = C.ELM_FOCUS_UP
	FocusDirectionDown     FocusDirection = C.ELM_FOCUS_DOWN
	FocusDirectionRight    FocusDirection = C.ELM_FOCUS_RIGHT
	FocusDirectionLeft     FocusDirection = C.ELM_FOCUS_LEFT
)

type FlipInteraction int

const (
	FlipInteractionNone   FlipInteraction = C.ELM_FLIP_INTERACTION_NONE
	FlipInteractionRotate FlipInteraction = C.ELM_FLIP_INTERACTION_ROTATE
	FlipInteractionCube   FlipInteraction = C.ELM_FLIP_INTERACTION_CUBE
	FlipInteractionPage   FlipInteraction = C.ELM_FLIP_INTERACTION_PAGE
)

type FlipMode int

const (
	FlipModeRotateYCenterAxis  FlipMode = C.ELM_FLIP_ROTATE_Y_CENTER_AXIS
	FlipModeRotateXCenterAxis  FlipMode = C.ELM_FLIP_ROTATE_X_CENTER_AXIS
	FlipModeRotateXZCenterAxis FlipMode = C.ELM_FLIP_ROTATE_XZ_CENTER_AXIS
	FlipModeRotateYZCenterAxis FlipMode = C.ELM_FLIP_ROTATE_YZ_CENTER_AXIS
	FlipModeCubeLeft           FlipMode = C.ELM_FLIP_CUBE_LEFT
	FlipModeCubeRight          FlipMode = C.ELM_FLIP_CUBE_RIGHT
	FlipModeCubeUp             FlipMode = C.ELM_FLIP_CUBE_UP
	FlipModeCubeDown           FlipMode = C.ELM_FLIP_CUBE_DOWN
	FlipModePageLeft           FlipMode = C.ELM_FLIP_PAGE_LEFT
	FlipModePageRight          FlipMode = C.ELM_FLIP_PAGE_RIGHT
	FlipModePageUp             FlipMode = C.ELM_FLIP_PAGE_UP
	FlipModePageDown           FlipMode = C.ELM_FLIP_PAGE_DOWN
)

type FlipDirection int

const (
	FlipDirectionUp    FlipDirection = C.ELM_FLIP_DIRECTION_UP
	FlipDirectionDown  FlipDirection = C.ELM_FLIP_DIRECTION_DOWN
	FlipDirectionLeft  FlipDirection = C.ELM_FLIP_DIRECTION_LEFT
	FlipDirectionRight FlipDirection = C.ELM_FLIP_DIRECTION_RIGHT
)

type HoverAxis int

const (
	HoverAxisNone       HoverAxis = C.ELM_HOVER_AXIS_NONE
	HoverAxisHorizontal HoverAxis = C.ELM_HOVER_AXIS_HORIZONTAL
	HoverAxisVertical   HoverAxis = C.ELM_HOVER_AXIS_VERTICAL
	HoverAxisBoth       HoverAxis = C.ELM_HOVER_AXIS_BOTH
)

type PhotocamZoomMode int

const (
	PhotocamZoomModeManual    PhotocamZoomMode = C.ELM_PHOTOCAM_ZOOM_MODE_MANUAL
	PhotocamZoomModeAutoFit   PhotocamZoomMode = C.ELM_PHOTOCAM_ZOOM_MODE_AUTO_FIT
	PhotocamZoomModeAutoFill  PhotocamZoomMode = C.ELM_PHOTOCAM_ZOOM_MODE_AUTO_FILL
	PhotocamZoomModeAutoFitIn PhotocamZoomMode = C.ELM_PHOTOCAM_ZOOM_MODE_AUTO_FIT_IN
	//PhotocamZoomModeLast PhotocamZoomMode = C.ELM_PHOTOCAM_ZOOM_MODE_LAST
)

type SelectMode int

const (
	SelectModeDefault     SelectMode = C.ELM_OBJECT_SELECT_MODE_DEFAULT
	SelectModeAlways      SelectMode = C.ELM_OBJECT_SELECT_MODE_ALWAYS
	SelectModeNone        SelectMode = C.ELM_OBJECT_SELECT_MODE_NONE
	SelectModeDisplayOnly SelectMode = C.ELM_OBJECT_SELECT_MODE_DISPLAY_ONLY
	//SelectMode SelectMode = C.ELM_OBJECT_SELECT_MODE_MAX
)

type ToolbarShrinkMode int

const (
	ToolbarShrinkModeNone   ToolbarShrinkMode = C.ELM_TOOLBAR_SHRINK_NONE
	ToolbarShrinkModeHide   ToolbarShrinkMode = C.ELM_TOOLBAR_SHRINK_HIDE
	ToolbarShrinkModeScroll ToolbarShrinkMode = C.ELM_TOOLBAR_SHRINK_SCROLL
	ToolbarShrinkModeMenu   ToolbarShrinkMode = C.ELM_TOOLBAR_SHRINK_MENU
	ToolbarShrinkModeExpand ToolbarShrinkMode = C.ELM_TOOLBAR_SHRINK_EXPAND
	ToolbarShrinkModeLast   ToolbarShrinkMode = C.ELM_TOOLBAR_SHRINK_LAST
)

type IconLookupOrder int

const (
	IconLookupOrderThemeFDO IconLookupOrder = C.ELM_ICON_LOOKUP_THEME_FDO
	IconLookupOrderFDO      IconLookupOrder = C.ELM_ICON_LOOKUP_FDO
	IconLookupOrderTheme    IconLookupOrder = C.ELM_ICON_LOOKUP_THEME
)

type CNPMode int

const (
	CNPModeMarkup    CNPMode = C.ELM_CNP_MODE_MARKUP
	CNPModeNoImage   CNPMode = C.ELM_CNP_MODE_NO_IMAGE
	CNPModePlaintext CNPMode = C.ELM_CNP_MODE_PLAINTEXT
)

type TextFormat int

const (
	TextFormatMarkupUTF8 TextFormat = C.ELM_TEXT_FORMAT_MARKUP_UTF8
	TextFormatPlainUTF8  TextFormat = C.ELM_TEXT_FORMAT_PLAIN_UTF8
)

type InputPanelLanguage int

const (
	InputPanelLanguageAutomatic InputPanelLanguage = C.ELM_INPUT_PANEL_LANG_AUTOMATIC
	InputPanelLanguageAlphabet  InputPanelLanguage = C.ELM_INPUT_PANEL_LANG_ALPHABET
)

type AutocapitalType int

const (
	AutocapitalTypeNone     AutocapitalType = C.ELM_AUTOCAPITAL_TYPE_NONE
	AutocapitalTypeWord     AutocapitalType = C.ELM_AUTOCAPITAL_TYPE_WORD
	AutocapitalTypeSentence AutocapitalType = C.ELM_AUTOCAPITAL_TYPE_SENTENCE
	AutocapitalTypeAllChar  AutocapitalType = C.ELM_AUTOCAPITAL_TYPE_ALLCHARACTER
)

type InputHint int

const (
	InputHintNone          InputHint = C.ELM_INPUT_HINT_NONE
	InputHintAutoComplete  InputHint = C.ELM_INPUT_HINT_AUTO_COMPLETE
	InputHintSensitiveData InputHint = C.ELM_INPUT_HINT_SENSITIVE_DATA
)

type InputPanelLayout int

const (
	InputPanelLayoutNormal      InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_NORMAL
	InputPanelLayoutNumber      InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_NUMBER
	InputPanelLayoutEmail       InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_EMAIL
	InputPanelLayoutURL         InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_URL
	InputPanelLayoutPhoneNumber InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_PHONENUMBER
	InputPanelLayoutIP          InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_IP
	InputPanelLayoutMonth       InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_MONTH
	InputPanelLayoutNumberonly  InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_NUMBERONLY
	InputPanelLayoutInvalid     InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_INVALID
	InputPanelLayoutHex         InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_HEX
	InputPanelLayoutTerminal    InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_TERMINAL
	InputPanelLayoutPassword    InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_PASSWORD
	InputPanelLayoutDatetime    InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_DATETIME
	InputPanelLayoutEmotion     InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_EMOTICON

// in 1.19 InputPanelLayoutVoice InputPanelLayout = C.ELM_INPUT_PANEL_LAYOUT_VOICE
)

type InputPanelReturnKeyType int

const (
	InputPanelReturnKeyTypeDefault InputPanelReturnKeyType = C.ELM_INPUT_PANEL_RETURN_KEY_TYPE_DEFAULT
	InputPanelReturnKeyTypeDone    InputPanelReturnKeyType = C.ELM_INPUT_PANEL_RETURN_KEY_TYPE_DONE
	InputPanelReturnKeyTypeGo      InputPanelReturnKeyType = C.ELM_INPUT_PANEL_RETURN_KEY_TYPE_GO
	InputPanelReturnKeyTypeJoin    InputPanelReturnKeyType = C.ELM_INPUT_PANEL_RETURN_KEY_TYPE_JOIN
	InputPanelReturnKeyTypeLogin   InputPanelReturnKeyType = C.ELM_INPUT_PANEL_RETURN_KEY_TYPE_LOGIN
	InputPanelReturnKeyTypeNext    InputPanelReturnKeyType = C.ELM_INPUT_PANEL_RETURN_KEY_TYPE_NEXT
	InputPanelReturnKeyTypeSearch  InputPanelReturnKeyType = C.ELM_INPUT_PANEL_RETURN_KEY_TYPE_SEARCH
	InputPanelReturnKeyTypeSend    InputPanelReturnKeyType = C.ELM_INPUT_PANEL_RETURN_KEY_TYPE_SEND
	InputPanelReturnKeyTypeSignin  InputPanelReturnKeyType = C.ELM_INPUT_PANEL_RETURN_KEY_TYPE_SIGNIN
)

type PanelOrient int

const (
	PanelOrientTop    PanelOrient = C.ELM_PANEL_ORIENT_TOP
	PanelOrientBottom PanelOrient = C.ELM_PANEL_ORIENT_BOTTOM
	PanelOrientLeft   PanelOrient = C.ELM_PANEL_ORIENT_LEFT
	PanelOrientRight  PanelOrient = C.ELM_PANEL_ORIENT_RIGHT
)

type CtxpopupDirection int

const (
	CtxpopupDirectionDown    CtxpopupDirection = C.ELM_CTXPOPUP_DIRECTION_DOWN
	CtxpopupDirectionRight   CtxpopupDirection = C.ELM_CTXPOPUP_DIRECTION_RIGHT
	CtxpopupDirectionLeft    CtxpopupDirection = C.ELM_CTXPOPUP_DIRECTION_LEFT
	CtxpopupDirectionUp      CtxpopupDirection = C.ELM_CTXPOPUP_DIRECTION_UP
	CtxpopupDirectionUnknown CtxpopupDirection = C.ELM_CTXPOPUP_DIRECTION_UNKNOWN
)

type DatetimeField int

const (
	DatetimeFieldYear   DatetimeField = C.ELM_DATETIME_YEAR
	DatetimeFieldMonth  DatetimeField = C.ELM_DATETIME_MONTH
	DatetimeFieldDate   DatetimeField = C.ELM_DATETIME_DATE
	DatetimeFieldHour   DatetimeField = C.ELM_DATETIME_HOUR
	DatetimeFieldMinute DatetimeField = C.ELM_DATETIME_MINUTE
	DatetimeFieldAmPm   DatetimeField = C.ELM_DATETIME_AMPM
)

type MultiSelectMode int

const (
	MultiSelectModeDefault     MultiSelectMode = C.ELM_OBJECT_MULTI_SELECT_MODE_DEFAULT
	MultiSelectModeWithControl MultiSelectMode = C.ELM_OBJECT_MULTI_SELECT_MODE_WITH_CONTROL
)

type ListMode int

const (
	ListModeCompress ListMode = C.ELM_LIST_COMPRESS
	ListModeScroll   ListMode = C.ELM_LIST_SCROLL
	ListModeLimit    ListMode = C.ELM_LIST_LIMIT
	ListModeExpand   ListMode = C.ELM_LIST_EXPAND
)

type ThumbFormat int

const (
	ThumbFormatFDO  ThumbFormat = C.ETHUMB_THUMB_FDO
	ThumbFormatJpeg ThumbFormat = C.ETHUMB_THUMB_JPEG
	ThumbFormatEet  ThumbFormat = C.ETHUMB_THUMB_EET
)

type ThumbAnimationState int

const (
	ThumbAnimationStateStart ThumbAnimationState = C.ELM_THUMB_ANIMATION_START
	ThumbAnimationStateStop  ThumbAnimationState = C.ELM_THUMB_ANIMATION_STOP
	ThumbAnimationStateLoop  ThumbAnimationState = C.ELM_THUMB_ANIMATION_LOOP
)

type ThumbFDOSizeType int

const (
	ThumbFDOSizeTypeNormal ThumbFDOSizeType = C.ETHUMB_THUMB_NORMAL
	ThumbFDOSizeTypeLarge  ThumbFDOSizeType = C.ETHUMB_THUMB_LARGE
)

type ThumbOrientation int

const (
	ThumbOrientationNone       ThumbOrientation = C.ETHUMB_THUMB_ORIENT_NONE
	ThumbOrientation90CW       ThumbOrientation = C.ETHUMB_THUMB_ROTATE_90_CW
	ThumbOrientation180        ThumbOrientation = C.ETHUMB_THUMB_ROTATE_180
	ThumbOrientation90CCW      ThumbOrientation = C.ETHUMB_THUMB_ROTATE_90_CCW
	ThumbOrientationHorizontal ThumbOrientation = C.ETHUMB_THUMB_FLIP_HORIZONTAL
	ThumbOrientationVertical   ThumbOrientation = C.ETHUMB_THUMB_FLIP_VERTICAL
	ThumbOrientationTranspose  ThumbOrientation = C.ETHUMB_THUMB_FLIP_TRANSPOSE
	ThumbOrientationTransverse ThumbOrientation = C.ETHUMB_THUMB_FLIP_TRANSVERSE
	ThumbOrientationOriginal   ThumbOrientation = C.ETHUMB_THUMB_ORIENT_ORIGINAL
)

type ThumbAspect int

const (
	ThumbAspectKeep   ThumbAspect = C.ETHUMB_THUMB_KEEP_ASPECT
	ThumbAspectIgnore ThumbAspect = C.ETHUMB_THUMB_IGNORE_ASPECT
	ThumbAspectCrop   ThumbAspect = C.ETHUMB_THUMB_CROP
)

type CallbackType int

const (
	CallbackTypeMouseIn              CallbackType = C.EVAS_CALLBACK_MOUSE_IN
	CallbackTypeMouseOut             CallbackType = C.EVAS_CALLBACK_MOUSE_OUT
	CallbackTypeMouseDown            CallbackType = C.EVAS_CALLBACK_MOUSE_DOWN
	CallbackTypeMouseUp              CallbackType = C.EVAS_CALLBACK_MOUSE_UP
	CallbackTypeMouseMove            CallbackType = C.EVAS_CALLBACK_MOUSE_MOVE
	CallbackTypeMouseWheel           CallbackType = C.EVAS_CALLBACK_MOUSE_WHEEL
	CallbackTypeMultiDown            CallbackType = C.EVAS_CALLBACK_MULTI_DOWN
	CallbackTypeMultiUp              CallbackType = C.EVAS_CALLBACK_MULTI_UP
	CallbackTypeMultiMove            CallbackType = C.EVAS_CALLBACK_MULTI_MOVE
	CallbackTypeKeyDown              CallbackType = C.EVAS_CALLBACK_KEY_DOWN
	CallbackTypeKeyUp                CallbackType = C.EVAS_CALLBACK_KEY_UP
	CallbackTypeFocusIn              CallbackType = C.EVAS_CALLBACK_FOCUS_IN
	CallbackTypeFocusOut             CallbackType = C.EVAS_CALLBACK_FOCUS_OUT
	CallbackTypeShow                 CallbackType = C.EVAS_CALLBACK_SHOW
	CallbackTypeHide                 CallbackType = C.EVAS_CALLBACK_HIDE
	CallbackTypeMove                 CallbackType = C.EVAS_CALLBACK_MOVE
	CallbackTypeResize               CallbackType = C.EVAS_CALLBACK_RESIZE
	CallbackTypeRestack              CallbackType = C.EVAS_CALLBACK_RESTACK
	CallbackTypeDel                  CallbackType = C.EVAS_CALLBACK_DEL
	CallbackTypeFree                 CallbackType = C.EVAS_CALLBACK_FREE
	CallbackTypeHold                 CallbackType = C.EVAS_CALLBACK_HOLD
	CallbackTypeSizeHintsChanged     CallbackType = C.EVAS_CALLBACK_CHANGED_SIZE_HINTS
	CallbackTypeImagePreloaded       CallbackType = C.EVAS_CALLBACK_IMAGE_PRELOADED
	CallbackTypeImageUnloaded        CallbackType = C.EVAS_CALLBACK_IMAGE_UNLOADED
	CallbackTypeImageResize          CallbackType = C.EVAS_CALLBACK_IMAGE_RESIZE
	CallbackTypeCanvasFocusIn        CallbackType = C.EVAS_CALLBACK_CANVAS_FOCUS_IN
	CallbackTypeCanvasFocusOut       CallbackType = C.EVAS_CALLBACK_CANVAS_FOCUS_OUT
	CallbackTypeRenderFlushPre       CallbackType = C.EVAS_CALLBACK_RENDER_FLUSH_PRE
	CallbackTypeRenderFlushPost      CallbackType = C.EVAS_CALLBACK_RENDER_FLUSH_POST
	CallbackTypeCanvasObjectFocusIn  CallbackType = C.EVAS_CALLBACK_CANVAS_OBJECT_FOCUS_IN
	CallbackTypeCanvasObjectFocusOut CallbackType = C.EVAS_CALLBACK_CANVAS_OBJECT_FOCUS_OUT
	CallbackTypePreRender            CallbackType = C.EVAS_CALLBACK_RENDER_PRE
	CallbackTypePostRender           CallbackType = C.EVAS_CALLBACK_RENDER_POST
	CallbackTypeDeviceChanged        CallbackType = C.EVAS_CALLBACK_DEVICE_CHANGED
	CallbackTypeAxisUpdate           CallbackType = C.EVAS_CALLBACK_AXIS_UPDATE
	CallbackTypeViewportResize       CallbackType = C.EVAS_CALLBACK_CANVAS_VIEWPORT_RESIZE

//	CallbackTypeLast                 CallbackType = C.EVAS_CALLBACK_LAST
)

type TouchPointState int

const (
	TouchPointDown   TouchPointState = C.EVAS_TOUCH_POINT_DOWN
	TouchPointUp     TouchPointState = C.EVAS_TOUCH_POINT_UP
	TouchPointMove   TouchPointState = C.EVAS_TOUCH_POINT_MOVE
	TouchPointStill  TouchPointState = C.EVAS_TOUCH_POINT_STILL
	TouchPointCancel TouchPointState = C.EVAS_TOUCH_POINT_CANCEL
)

type EventFlags int

const (
	EventFlagsNone     EventFlags = C.EVAS_EVENT_FLAG_NONE
	EventFlagsOnHold   EventFlags = C.EVAS_EVENT_FLAG_ON_HOLD
	EventFlagsOnScroll EventFlags = C.EVAS_EVENT_FLAG_ON_SCROLL
)

type AllocError int

const (
	AllocErrorNone      AllocError = C.EVAS_ALLOC_ERROR_NONE
	AllocErrorFatal     AllocError = C.EVAS_ALLOC_ERROR_FATAL
	AllocErrorRecovered AllocError = C.EVAS_ALLOC_ERROR_RECOVERED
)

func (e AllocError) Error() string {
	switch e {
	case AllocErrorNone:
		return ""
	case AllocErrorFatal:
		return "Fatal Alloc Error"
	case AllocErrorRecovered:
		return "Recovered Error"
	}

	return "Unknown Evas Alloc Error"
}

//Evas_Fill_Spread
type Texture int

const (
	TextureReflect         Texture = C.EVAS_TEXTURE_REFLECT
	TextureRepeat          Texture = C.EVAS_TEXTURE_REPEAT
	TextureRestrict        Texture = C.EVAS_TEXTURE_RESTRICT
	TextureRestrictReflect Texture = C.EVAS_TEXTURE_RESTRICT_REFLECT
	TextureRestrictRepeat  Texture = C.EVAS_TEXTURE_RESTRICT_REPEAT
	TexturePad             Texture = C.EVAS_TEXTURE_PAD
)

type PixelFormat int

const (
	PixelFormatNone        PixelFormat = C.EVAS_PIXEL_FORMAT_NONE
	PixelFormatARGB32      PixelFormat = C.EVAS_PIXEL_FORMAT_ARGB32
	PixelFormatYUV420P_601 PixelFormat = C.EVAS_PIXEL_FORMAT_YUV420P_601
)

type GengridReorderType int

const (
	GengridReorderTypeNormal GengridReorderType = C.ELM_GENGRID_REORDER_TYPE_NORMAL
	GengridReorderTypeSwap   GengridReorderType = C.ELM_GENGRID_REORDER_TYPE_SWAP
)

type PositionMap int

const (
	PositionMapLinear           PositionMap = C.ECORE_POS_MAP_LINEAR
	PositionMapAccelerate       PositionMap = C.ECORE_POS_MAP_ACCELERATE
	PositionMapDecelerate       PositionMap = C.ECORE_POS_MAP_DECELERATE
	PositionMapSinusoidal       PositionMap = C.ECORE_POS_MAP_SINUSOIDAL
	PositionMapAccelerateFactor PositionMap = C.ECORE_POS_MAP_ACCELERATE_FACTOR
	PositionMapDecelerateFactor PositionMap = C.ECORE_POS_MAP_DECELERATE_FACTOR
	PositionMapSinusoidalFactor PositionMap = C.ECORE_POS_MAP_SINUSOIDAL_FACTOR
	PositionMapDivisorInterp    PositionMap = C.ECORE_POS_MAP_DIVISOR_INTERP
	PositionMapBounce           PositionMap = C.ECORE_POS_MAP_BOUNCE
	PositionMapSpring           PositionMap = C.ECORE_POS_MAP_SPRING
	PositionMapCubicBezier      PositionMap = C.ECORE_POS_MAP_CUBIC_BEZIER
)

type PopupOrient int

const (
	PopupOrientTop         PopupOrient = C.ELM_POPUP_ORIENT_TOP
	PopupOrientCenter      PopupOrient = C.ELM_POPUP_ORIENT_CENTER
	PopupOrientBottom      PopupOrient = C.ELM_POPUP_ORIENT_BOTTOM
	PopupOrientLeft        PopupOrient = C.ELM_POPUP_ORIENT_LEFT
	PopupOrientRight       PopupOrient = C.ELM_POPUP_ORIENT_RIGHT
	PopupOrientTopLeft     PopupOrient = C.ELM_POPUP_ORIENT_TOP_LEFT
	PopupOrientTopRight    PopupOrient = C.ELM_POPUP_ORIENT_TOP_RIGHT
	PopupOrientBottomLeft  PopupOrient = C.ELM_POPUP_ORIENT_BOTTOM_LEFT
	PopupOrientBottomRight PopupOrient = C.ELM_POPUP_ORIENT_BOTTOM_RIGHT
)

type SliderIndicatorVisibleMode int

const (
	SliderIndicatorVisibleModeDefault SliderIndicatorVisibleMode = C.ELM_SLIDER_INDICATOR_VISIBLE_MODE_DEFAULT
	SliderIndicatorVisibleModeAlways  SliderIndicatorVisibleMode = C.ELM_SLIDER_INDICATOR_VISIBLE_MODE_ALWAYS
	SliderIndicatorVisibleModeOnFocus SliderIndicatorVisibleMode = C.ELM_SLIDER_INDICATOR_VISIBLE_MODE_ON_FOCUS
	SliderIndicatorVisibleModeNone    SliderIndicatorVisibleMode = C.ELM_SLIDER_INDICATOR_VISIBLE_MODE_NONE
)

type ProcState int

const (
	ProcStateForeground ProcState = C.ELM_PROCESS_STATE_FOREGROUND
	ProcStateBackground ProcState = C.ELM_PROCESS_STATE_BACKGROUND
)

type ToolbarItemScrolltoType int

const (
	ToolbarItemScrolltoTypeNone   ToolbarItemScrolltoType = C.ELM_TOOLBAR_ITEM_SCROLLTO_NONE
	ToolbarItemScrolltoTypeIn     ToolbarItemScrolltoType = C.ELM_TOOLBAR_ITEM_SCROLLTO_IN
	ToolbarItemScrolltoTypeFirst  ToolbarItemScrolltoType = C.ELM_TOOLBAR_ITEM_SCROLLTO_FIRST
	ToolbarItemScrolltoTypeMiddle ToolbarItemScrolltoType = C.ELM_TOOLBAR_ITEM_SCROLLTO_MIDDLE
	ToolbarItemScrolltoTypeLast   ToolbarItemScrolltoType = C.ELM_TOOLBAR_ITEM_SCROLLTO_LAST
)

type GengridItemScrolltoType int

const (
	GengridItemScrolltoTypeNone   GengridItemScrolltoType = C.ELM_GENGRID_ITEM_SCROLLTO_NONE
	GengridItemScrolltoTypeIn     GengridItemScrolltoType = C.ELM_GENGRID_ITEM_SCROLLTO_IN
	GengridItemScrolltoTypeTop    GengridItemScrolltoType = C.ELM_GENGRID_ITEM_SCROLLTO_TOP
	GengridItemScrolltoTypeMiddle GengridItemScrolltoType = C.ELM_GENGRID_ITEM_SCROLLTO_MIDDLE
	GengridItemScrolltoTypeBottom GengridItemScrolltoType = C.ELM_GENGRID_ITEM_SCROLLTO_BOTTOM
)

type IconType int

const (
	IconTypeNone     IconType = C.ELM_ICON_NONE
	IconTypeFile     IconType = C.ELM_ICON_FILE
	IconTypeStandard IconType = C.ELM_ICON_STANDARD
)

type GenlistItemScrolltoType int

const (
	GenlistItemScrolltoTypeNone   GenlistItemScrolltoType = C.ELM_GENLIST_ITEM_SCROLLTO_NONE
	GenlistItemScrolltoTypeIn     GenlistItemScrolltoType = C.ELM_GENLIST_ITEM_SCROLLTO_IN
	GenlistItemScrolltoTypeTop    GenlistItemScrolltoType = C.ELM_GENLIST_ITEM_SCROLLTO_TOP
	GenlistItemScrolltoTypeMiddle GenlistItemScrolltoType = C.ELM_GENLIST_ITEM_SCROLLTO_MIDDLE
	GenlistItemScrolltoTypeBottom GenlistItemScrolltoType = C.ELM_GENLIST_ITEM_SCROLLTO_BOTTOM
)

type GenlistItemFieldType int

const (
	GenlistItemFieldTypeAll     GenlistItemFieldType = C.ELM_GENLIST_ITEM_FIELD_ALL
	GenlistItemFieldTypeText    GenlistItemFieldType = C.ELM_GENLIST_ITEM_FIELD_TEXT
	GenlistItemFieldTypeContent GenlistItemFieldType = C.ELM_GENLIST_ITEM_FIELD_CONTENT
	GenlistItemFieldTypeState   GenlistItemFieldType = C.ELM_GENLIST_ITEM_FIELD_STATE
)

type GenlistItemType int

const (
	GenlistItemTypeNone  GenlistItemType = C.ELM_GENLIST_ITEM_NONE
	GenlistItemTypeTree  GenlistItemType = C.ELM_GENLIST_ITEM_TREE
	GenlistItemTypeGroup GenlistItemType = C.ELM_GENLIST_ITEM_GROUP
)

type GlobMatchFlag int

const (
	GlobMatchFlagNoEscape GlobMatchFlag = C.ELM_GLOB_MATCH_NO_ESCAPE
	GlobMatchFlagPath     GlobMatchFlag = C.ELM_GLOB_MATCH_PATH
	GlobMatchFlagPeriod   GlobMatchFlag = C.ELM_GLOB_MATCH_PERIOD
	GlobMatchFlagNocase   GlobMatchFlag = C.ELM_GLOB_MATCH_NOCASE
)

type DisplayMode int

const (
	DisplayModeNone       DisplayMode = C.EVAS_DISPLAY_MODE_NONE
	DisplayModeCompress   DisplayMode = C.EVAS_DISPLAY_MODE_COMPRESS
	DisplayModeExpand     DisplayMode = C.EVAS_DISPLAY_MODE_EXPAND
	DisplayModeDontChange DisplayMode = C.EVAS_DISPLAY_MODE_DONT_CHANGE
)

type MapZoomMode int

const (
	MapZoomModeManual MapZoomMode = C.ELM_MAP_ZOOM_MODE_MANUAL
	MapZoomModeFit    MapZoomMode = C.ELM_MAP_ZOOM_MODE_AUTO_FIT
	MapZoomModeFill   MapZoomMode = C.ELM_MAP_ZOOM_MODE_AUTO_FILL
	//ELM_MAP_ZOOM_MODE_LAST
)

type MapSourceType int

const (
	MapSourceTypeTile  MapSourceType = C.ELM_MAP_SOURCE_TYPE_TILE
	MapSourceTypeRoute MapSourceType = C.ELM_MAP_SOURCE_TYPE_ROUTE
	MapSourceTypeName  MapSourceType = C.ELM_MAP_SOURCE_TYPE_NAME
	//ELM_MAP_SOURCE_TYPE_LAST
)

type MapOverlayType int

const (
	MapOverlayTypeNone    MapOverlayType = C.ELM_MAP_OVERLAY_TYPE_NONE
	MapOverlayTypeDefault MapOverlayType = C.ELM_MAP_OVERLAY_TYPE_DEFAULT
	MapOverlayTypeClass   MapOverlayType = C.ELM_MAP_OVERLAY_TYPE_CLASS
	MapOverlayTypeGroup   MapOverlayType = C.ELM_MAP_OVERLAY_TYPE_GROUP
	MapOverlayTypeBubble  MapOverlayType = C.ELM_MAP_OVERLAY_TYPE_BUBBLE
	MapOverlayTypeRoute   MapOverlayType = C.ELM_MAP_OVERLAY_TYPE_ROUTE
	MapOverlayTypeLine    MapOverlayType = C.ELM_MAP_OVERLAY_TYPE_LINE
	MapOverlayTypePolygon MapOverlayType = C.ELM_MAP_OVERLAY_TYPE_POLYGON
	MapOverlayTypeCircle  MapOverlayType = C.ELM_MAP_OVERLAY_TYPE_CIRCLE
	MapOverlayTypeScale   MapOverlayType = C.ELM_MAP_OVERLAY_TYPE_SCALE
)

type MapRouteType int

const (
	MapRouteTypeMotocar MapRouteType = C.ELM_MAP_ROUTE_TYPE_MOTOCAR
	MapRouteTypeBicycle MapRouteType = C.ELM_MAP_ROUTE_TYPE_BICYCLE
	MapRouteTypeFoot    MapRouteType = C.ELM_MAP_ROUTE_TYPE_FOOT
	//ELM_MAP_ROUTE_TYPE_LAST
)

type MapRouteMethod int

const (
	MapRouteMethodFastest  MapRouteMethod = C.ELM_MAP_ROUTE_METHOD_FASTEST
	MapRouteMethodShortest MapRouteMethod = C.ELM_MAP_ROUTE_METHOD_SHORTEST
	//ELM_MAP_ROUTE_METHOD_LAST
)

type AxisLabel int

const (
	AxisLabelUnknown         AxisLabel = C.EVAS_AXIS_LABEL_UNKNOWN
	AxisLabelX               AxisLabel = C.EVAS_AXIS_LABEL_X
	AxisLabelY               AxisLabel = C.EVAS_AXIS_LABEL_Y
	AxisLabelPressure        AxisLabel = C.EVAS_AXIS_LABEL_PRESSURE
	AxisLabelDistance        AxisLabel = C.EVAS_AXIS_LABEL_DISTANCE
	AxisLabelAzimuth         AxisLabel = C.EVAS_AXIS_LABEL_AZIMUTH
	AxisLabelTilt            AxisLabel = C.EVAS_AXIS_LABEL_TILT
	AxisLabelTwist           AxisLabel = C.EVAS_AXIS_LABEL_TWIST
	AxisLabelTouchWidthMajor AxisLabel = C.EVAS_AXIS_LABEL_TOUCH_WIDTH_MAJOR
	AxisLabelTouchWidthMinor AxisLabel = C.EVAS_AXIS_LABEL_TOUCH_WIDTH_MINOR
	AxisLabelToolWidthMajor  AxisLabel = C.EVAS_AXIS_LABEL_TOOL_WIDTH_MAJOR
	AxisLabelToolWidthMinor  AxisLabel = C.EVAS_AXIS_LABEL_TOOL_WIDTH_MINOR
	/* since 1.19
	AxisLabelWindowX         AxisLabel = C.EVAS_AXIS_LABEL_WINDOW_X
	AxisLabelWindowY         AxisLabel = C.EVAS_AXIS_LABEL_WINDOW_Y
	AxisLabelNormalX         AxisLabel = C.EVAS_AXIS_LABEL_NORMAL_X
	AxisLabelNormalY         AxisLabel = C.EVAS_AXIS_LABEL_NORMAL_Y
	*/
)

type BidiDirection int

const (
	BidiDirectionNatural BidiDirection = C.EVAS_BIDI_DIRECTION_NATURAL
	BidiDirectionNeutral BidiDirection = C.EVAS_BIDI_DIRECTION_NEUTRAL
	BidiDirectionLTR     BidiDirection = C.EVAS_BIDI_DIRECTION_LTR
	BidiDirectionRTL     BidiDirection = C.EVAS_BIDI_DIRECTION_RTL
	BidiDirectionInherit BidiDirection = C.EVAS_BIDI_DIRECTION_INHERIT
)

type PointerMode int

const (
	PointerModeAutograb             PointerMode = C.EVAS_OBJECT_POINTER_MODE_AUTOGRAB
	PointerModeNograb               PointerMode = C.EVAS_OBJECT_POINTER_MODE_NOGRAB
	PointerModeNograbNoRepeatUpDown PointerMode = C.EVAS_OBJECT_POINTER_MODE_NOGRAB_NO_REPEAT_UPDOWN
)

type IllumeCommand int

const (
	IllumeCommandFocusBack    IllumeCommand = C.ELM_ILLUME_COMMAND_FOCUS_BACK
	IllumeCommandFocusForward IllumeCommand = C.ELM_ILLUME_COMMAND_FOCUS_FORWARD
	IllumeCommandFocusHome    IllumeCommand = C.ELM_ILLUME_COMMAND_FOCUS_HOME
	IllumeCommandClose        IllumeCommand = C.ELM_ILLUME_COMMAND_CLOSE
)

type EmotionAspect int

const (
	EmotionAspectKeepNone   EmotionAspect = C.EMOTION_ASPECT_KEEP_NONE
	EmotionAspectKeepWidth  EmotionAspect = C.EMOTION_ASPECT_KEEP_WIDTH
	EmotionAspectKeepHeight EmotionAspect = C.EMOTION_ASPECT_KEEP_HEIGHT
	EmotionAspectKeepBoth   EmotionAspect = C.EMOTION_ASPECT_KEEP_BOTH
	EmotionAspectCrop       EmotionAspect = C.EMOTION_ASPECT_CROP
	EmotionAspectCustom     EmotionAspect = C.EMOTION_ASPECT_CUSTOM
)

type EmotionVisualization int

const (
	EmotionVisualizationNone                      EmotionVisualization = C.EMOTION_VIS_NONE
	EmotionVisualizationGoom                      EmotionVisualization = C.EMOTION_VIS_GOOM
	EmotionVisualizationLibvisualBumpscore        EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_BUMPSCOPE
	EmotionVisualizationLibvisualCorona           EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_CORONA
	EmotionVisualizationLibvisualDancingParticles EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_DANCING_PARTICLES
	EmotionVisualizationLibvisualGdkpixbuf        EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_GDKPIXBUF
	EmotionVisualizationLibvisualGForce           EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_G_FORCE
	EmotionVisualizationLibvisualGoom             EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_GOOM
	EmotionVisualizationLibvisualInfinite         EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_INFINITE
	EmotionVisualizationLibvisualJakdaw           EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_JAKDAW
	EmotionVisualizationLibvisualJess             EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_JESS
	EmotionVisualizationLibvisualLvAnalyser       EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_LV_ANALYSER
	EmotionVisualizationLibvisualLvFlower         EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_LV_FLOWER
	EmotionVisualizationLibvisualLvGlTest         EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_LV_GLTEST
	EmotionVisualizationLibvisualLvScope          EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_LV_SCOPE
	EmotionVisualizationLibvisualMadspin          EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_MADSPIN
	EmotionVisualizationLibvisualNebulus          EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_NEBULUS
	EmotionVisualizationLibvisualOinksie          EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_OINKSIE
	EmotionVisualizationLibvisualPlasma           EmotionVisualization = C.EMOTION_VIS_LIBVISUAL_PLASMA
)

type EmotionSuspend int

const (
	EmotionSuspendWakeup    EmotionSuspend = C.EMOTION_WAKEUP
	EmotionSuspendSleep     EmotionSuspend = C.EMOTION_SLEEP
	EmotionSuspendDeepSleep EmotionSuspend = C.EMOTION_DEEP_SLEEP
	EmotionSuspendHibernate EmotionSuspend = C.EMOTION_HIBERNATE
)

type EmotionMetaInfo int

const (
	EmotionMetaInfoTrackTitle   EmotionMetaInfo = C.EMOTION_META_INFO_TRACK_TITLE
	EmotionMetaInfoTrackArtist  EmotionMetaInfo = C.EMOTION_META_INFO_TRACK_ARTIST
	EmotionMetaInfoTrackAlbum   EmotionMetaInfo = C.EMOTION_META_INFO_TRACK_ALBUM
	EmotionMetaInfoTrackYear    EmotionMetaInfo = C.EMOTION_META_INFO_TRACK_YEAR
	EmotionMetaInfoTrackGenre   EmotionMetaInfo = C.EMOTION_META_INFO_TRACK_GENRE
	EmotionMetaInfoTrackComment EmotionMetaInfo = C.EMOTION_META_INFO_TRACK_COMMENT
	EmotionMetaInfoTrackDiskId  EmotionMetaInfo = C.EMOTION_META_INFO_TRACK_DISC_ID
	EmotionMetaInfoTrackCount   EmotionMetaInfo = C.EMOTION_META_INFO_TRACK_COUNT
)

type EmotionEvent int

const (
	EmotionEventUp        EmotionEvent = C.EMOTION_EVENT_UP
	EmotionEventDown      EmotionEvent = C.EMOTION_EVENT_DOWN
	EmotionEventLeft      EmotionEvent = C.EMOTION_EVENT_LEFT
	EmotionEventRight     EmotionEvent = C.EMOTION_EVENT_RIGHT
	EmotionEventSelect    EmotionEvent = C.EMOTION_EVENT_SELECT
	EmotionEventNext      EmotionEvent = C.EMOTION_EVENT_NEXT
	EmotionEventPrev      EmotionEvent = C.EMOTION_EVENT_PREV
	EmotionEventAngleNext EmotionEvent = C.EMOTION_EVENT_ANGLE_NEXT
	EmotionEventAnglePrev EmotionEvent = C.EMOTION_EVENT_ANGLE_PREV
	EmotionEventForce     EmotionEvent = C.EMOTION_EVENT_FORCE
	EmotionEventMenu1     EmotionEvent = C.EMOTION_EVENT_MENU1
	EmotionEventMenu2     EmotionEvent = C.EMOTION_EVENT_MENU2
	EmotionEventMenu3     EmotionEvent = C.EMOTION_EVENT_MENU3
	EmotionEventMenu4     EmotionEvent = C.EMOTION_EVENT_MENU4
	EmotionEventMenu5     EmotionEvent = C.EMOTION_EVENT_MENU5
	EmotionEventMenu6     EmotionEvent = C.EMOTION_EVENT_MENU6
	EmotionEventMenu7     EmotionEvent = C.EMOTION_EVENT_MENU7
	EmotionEvent0         EmotionEvent = C.EMOTION_EVENT_0
	EmotionEvent1         EmotionEvent = C.EMOTION_EVENT_1
	EmotionEvent2         EmotionEvent = C.EMOTION_EVENT_2
	EmotionEvent3         EmotionEvent = C.EMOTION_EVENT_3
	EmotionEvent4         EmotionEvent = C.EMOTION_EVENT_4
	EmotionEvent5         EmotionEvent = C.EMOTION_EVENT_5
	EmotionEvent6         EmotionEvent = C.EMOTION_EVENT_6
	EmotionEvent7         EmotionEvent = C.EMOTION_EVENT_7
	EmotionEvent8         EmotionEvent = C.EMOTION_EVENT_8
	EmotionEvent9         EmotionEvent = C.EMOTION_EVENT_9
	EmotionEvent10        EmotionEvent = C.EMOTION_EVENT_10
)

/* TODO: in 1.19
type EmotionArtworkInfo int
const (
	EmotionArtworkInfoImage EmotionArtworkInfo = C.EMOTION_ARTWORK_IMAGE
	EmotionArtworkInfoPreviewImage EmotionArtworkInfo = C.EMOTION_ARTWORK_PREVIEW_IMAGE
)
*/

type MemoryState int

const (
	MemoryStateNormal MemoryState = C.ECORE_MEMORY_STATE_NORMAL
	MemoryStateLow    MemoryState = C.ECORE_MEMORY_STATE_LOW
)

type PowerState int

const (
	PowerStateMains   PowerState = C.ECORE_POWER_STATE_MAINS
	PowerStateBattery PowerState = C.ECORE_POWER_STATE_BATTERY
	PowerStateLow     PowerState = C.ECORE_POWER_STATE_LOW
)

type WebZoomMode int

const (
	WebZoomModeManual   WebZoomMode = C.ELM_WEB_ZOOM_MODE_MANUAL
	WebZoomModeAutoFit  WebZoomMode = C.ELM_WEB_ZOOM_MODE_AUTO_FIT
	WebZoomModeAutoFill WebZoomMode = C.ELM_WEB_ZOOM_MODE_AUTO_FILL

//WebZoomModeLast WebZoomMode = C.ELM_WEB_ZOOM_MODE_LAST
)

type WebMenuItemType int

const (
	WebMenuItemTypeSeperator WebMenuItemType = C.ELM_WEB_MENU_SEPARATOR
	WebMenuItemTypeGroup     WebMenuItemType = C.ELM_WEB_MENU_GROUP
	WebMenuItemTypeOption    WebMenuItemType = C.ELM_WEB_MENU_OPTION
)

type WebWindowFeatureFlag int

const (
	WebWindowFeatureFlagToolbar     WebWindowFeatureFlag = C.ELM_WEB_WINDOW_FEATURE_TOOLBAR
	WebWindowFeatureFlagStatusbar   WebWindowFeatureFlag = C.ELM_WEB_WINDOW_FEATURE_STATUSBAR
	WebWindowFeatureFlagScrollbars  WebWindowFeatureFlag = C.ELM_WEB_WINDOW_FEATURE_SCROLLBARS
	WebWindowFeatureFlagMenubar     WebWindowFeatureFlag = C.ELM_WEB_WINDOW_FEATURE_MENUBAR
	WebWindowFeatureFlagLocationbar WebWindowFeatureFlag = C.ELM_WEB_WINDOW_FEATURE_LOCATIONBAR
	WebWindowFeatureFlagFullscreen  WebWindowFeatureFlag = C.ELM_WEB_WINDOW_FEATURE_FULLSCREEN
)

type WebConsoleMessageFunc func(string, uint, string)
type WebWindowOpenFunc func(bool, *WebWindowFeatures) Object
type WebConfirmDialogFunc func(string) (bool, Object)
type WebPromptDialogFunc func(string, string) (string, bool, Object)
type WebAlertDialogFunc func(string) Object

type EvasTextgridPalette int

const (
	EvasTextgridPaletteNone     EvasTextgridPalette = C.EVAS_TEXTGRID_PALETTE_NONE
	EvasTextgridPaletteStandard EvasTextgridPalette = C.EVAS_TEXTGRID_PALETTE_STANDARD
	EvasTextgridPaletteExtended EvasTextgridPalette = C.EVAS_TEXTGRID_PALETTE_EXTENDED

//EVAS_TEXTGRID_PALETTE_LAST
)

type EvasTextgridFontStyle int

const (
	EvasTextgridFontStyleNormal EvasTextgridFontStyle = C.EVAS_TEXTGRID_FONT_STYLE_NORMAL
)

type EdjeChannel int

const (
	EdjeChannelEffect     EdjeChannel = C.EDJE_CHANNEL_EFFECT
	EdjeChannelBackground EdjeChannel = C.EDJE_CHANNEL_BACKGROUND
	EdjeChannelMusic      EdjeChannel = C.EDJE_CHANNEL_MUSIC
	EdjeChannelForeground EdjeChannel = C.EDJE_CHANNEL_FOREGROUND
	EdjeChannelInterface  EdjeChannel = C.EDJE_CHANNEL_INTERFACE
	EdjeChannelInput      EdjeChannel = C.EDJE_CHANNEL_INPUT
	EdjeChannelAlert      EdjeChannel = C.EDJE_CHANNEL_ALERT
	EdjeChannelAll        EdjeChannel = C.EDJE_CHANNEL_ALL
)

type SoftcursorMode int

const (
	SoftcursorModeAuto SoftcursorMode = C.EFL_UI_SOFTCURSOR_MODE_AUTO
	SoftcursorModeOn   SoftcursorMode = C.EFL_UI_SOFTCURSOR_MODE_ON
	SoftcursorModeOff  SoftcursorMode = C.EFL_UI_SOFTCURSOR_MODE_OFF
)

type GestureState int

const (
	GestureStateUndefined GestureState = C.ELM_GESTURE_STATE_UNDEFINED
	GestureStateStart     GestureState = C.ELM_GESTURE_STATE_START
	GestureStateMove      GestureState = C.ELM_GESTURE_STATE_MOVE
	GestureStateEnd       GestureState = C.ELM_GESTURE_STATE_END
	GestureStateAbort     GestureState = C.ELM_GESTURE_STATE_ABORT
)

type GestureType int

const (
	GestureTypeFirst       GestureType = C.ELM_GESTURE_FIRST
	GestureTypeNTaps       GestureType = C.ELM_GESTURE_N_TAPS
	GestureTypeNLongTaps   GestureType = C.ELM_GESTURE_N_LONG_TAPS
	GestureTypeNDoubleTaps GestureType = C.ELM_GESTURE_N_DOUBLE_TAPS
	GestureTypeNTripleTaps GestureType = C.ELM_GESTURE_N_TRIPLE_TAPS
	GestureTypeMomentum    GestureType = C.ELM_GESTURE_MOMENTUM
	GestureTypeNLines      GestureType = C.ELM_GESTURE_N_LINES
	GestureTypeNFlicks     GestureType = C.ELM_GESTURE_N_FLICKS
	GestureTypeZoom        GestureType = C.ELM_GESTURE_ZOOM
	GestureTypeRotate      GestureType = C.ELM_GESTURE_ROTATE

//GestureTypeLast GestureType = C.ELM_GESTURE_LAST
)

type SelectionType int

const (
	SelectionTypePrimary   SelectionType = C.ELM_SEL_TYPE_PRIMARY
	SelectionTypeSecondary SelectionType = C.ELM_SEL_TYPE_SECONDARY
	SelectionTypeXDnd      SelectionType = C.ELM_SEL_TYPE_XDND
	SelectionTypeClipboard SelectionType = C.ELM_SEL_TYPE_CLIPBOARD
)

type SelectionFormat int

const (
	SelectionFormatTargets SelectionFormat = C.ELM_SEL_FORMAT_TARGETS
	SelectionFormatNone    SelectionFormat = C.ELM_SEL_FORMAT_NONE
	SelectionFormatText    SelectionFormat = C.ELM_SEL_FORMAT_TEXT
	SelectionFormatMarkup  SelectionFormat = C.ELM_SEL_FORMAT_MARKUP
	SelectionFormatImage   SelectionFormat = C.ELM_SEL_FORMAT_IMAGE
	SelectionFormatVcard   SelectionFormat = C.ELM_SEL_FORMAT_VCARD
	SelectionFormatHtml    SelectionFormat = C.ELM_SEL_FORMAT_HTML
)

type XDndAction int

const (
	XDndActionUnknown     XDndAction = C.ELM_XDND_ACTION_UNKNOWN
	XDndActionCopy        XDndAction = C.ELM_XDND_ACTION_COPY
	XDndActionMove        XDndAction = C.ELM_XDND_ACTION_MOVE
	XDndActionPrivate     XDndAction = C.ELM_XDND_ACTION_PRIVATE
	XDndActionAsk         XDndAction = C.ELM_XDND_ACTION_ASK
	XDndActionList        XDndAction = C.ELM_XDND_ACTION_LIST
	XDndActionLink        XDndAction = C.ELM_XDND_ACTION_LINK
	XDndActionDescription XDndAction = C.ELM_XDND_ACTION_DESCRIPTION
)

type CanvasEventPriority int

const (
	CanvasEventPriorityDefault CanvasEventPriority = C.EVAS_CALLBACK_PRIORITY_DEFAULT
	CanvasEventPriorityAfter   CanvasEventPriority = C.EVAS_CALLBACK_PRIORITY_AFTER
	CanvasEventPriorityBefore  CanvasEventPriority = C.EVAS_CALLBACK_PRIORITY_BEFORE
)

type EdjeAspectControl int

const (
	EdjeAspectControlNone       EdjeAspectControl = C.EDJE_ASPECT_CONTROL_NONE
	EdjeAspectControlNeither    EdjeAspectControl = C.EDJE_ASPECT_CONTROL_NEITHER
	EdjeAspectControlHorizontal EdjeAspectControl = C.EDJE_ASPECT_CONTROL_HORIZONTAL
	EdjeAspectControlVertical   EdjeAspectControl = C.EDJE_ASPECT_CONTROL_VERTICAL
	EdjeAspectControlBoth       EdjeAspectControl = C.EDJE_ASPECT_CONTROL_BOTH
)

type EdjeLoadError int

const (
	EdjeLoadErrorNone               EdjeLoadError = C.EDJE_LOAD_ERROR_NONE
	EdjeLoadErrorGeneric            EdjeLoadError = C.EDJE_LOAD_ERROR_GENERIC
	EdjeLoadErrorDoesNotExist       EdjeLoadError = C.EDJE_LOAD_ERROR_DOES_NOT_EXIST
	EdjeLoadErrorPermissionDenied   EdjeLoadError = C.EDJE_LOAD_ERROR_PERMISSION_DENIED
	EdjeLoadErrorAllocationFailed   EdjeLoadError = C.EDJE_LOAD_ERROR_RESOURCE_ALLOCATION_FAILED
	EdjeLoadErrorCorruptFile        EdjeLoadError = C.EDJE_LOAD_ERROR_CORRUPT_FILE
	EdjeLoadErrorUnknownFormat      EdjeLoadError = C.EDJE_LOAD_ERROR_UNKNOWN_FORMAT
	EdjeLoadErrorIncompatibleFile   EdjeLoadError = C.EDJE_LOAD_ERROR_INCOMPATIBLE_FILE
	EdjeLoadErrorUnknownCollection  EdjeLoadError = C.EDJE_LOAD_ERROR_UNKNOWN_COLLECTION
	EdjeLoadErrorRecursiveReference EdjeLoadError = C.EDJE_LOAD_ERROR_RECURSIVE_REFERENCE
)

/*
TODO:
const char *edje_load_error_str (Edje_Load_Error error);
*/
func (e EdjeLoadError) Error() string {
	switch e {
	case EdjeLoadErrorNone:
		return ""
	case EdjeLoadErrorGeneric:
		return "Generic Error"
	case EdjeLoadErrorDoesNotExist:
		return "Does Not Exist"
	case EdjeLoadErrorPermissionDenied:
		return "Permission Denied"
	case EdjeLoadErrorAllocationFailed:
		return "Allocation Failed"
	case EdjeLoadErrorCorruptFile:
		return "Corrupt File"
	case EdjeLoadErrorUnknownFormat:
		return "Unknown Format"
	case EdjeLoadErrorIncompatibleFile:
		return "Incompatible File"
	case EdjeLoadErrorUnknownCollection:
		return "Unknown Collection"
	case EdjeLoadErrorRecursiveReference:
		return "Recursive Reference"
	default:
		return fmt.Sprintf("Unknown Error: Code = %d", e)
	}
}

type EdjeInputHint int

const (
	EdjeInputHintNone          EdjeInputHint = C.EDJE_INPUT_HINT_NONE
	EdjeInputHintAutoComplete  EdjeInputHint = C.EDJE_INPUT_HINT_AUTO_COMPLETE
	EdjeInputHintSensitiveData EdjeInputHint = C.EDJE_INPUT_HINT_SENSITIVE_DATA
)

//TODO: type WebDialogFileSelectorFunc func()

//https://github.com/tasn/efl/blob/master/src/lib/elementary/elm_general.eot

/*
typedef unsigned long long 	Evas_Modifier_Mask
typedef int 	Evas_Coord
typedef double 	Evas_Real

*/
