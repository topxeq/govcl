//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"fmt"
	"unsafe"

	. "github.com/topxeq/govcl/vcl/api"
	. "github.com/topxeq/govcl/vcl/types"
)

// 外部回调事件
// 参数一：函数地址
// 参数二：获取参数值的函数
// 返回值：如果为true则终止事件传递，如果为false则继续向后转发事件。
type TExtEventCallback func(fn interface{}, getVal func(idx int) uintptr) bool

// 外部扩展的事件回调，先不管重复注册的问题
var extEventCallback []TExtEventCallback

// CN: 注册外部扩展回调事件
// EN: Registering external extension callback events.
func RegisterExtEventCallback(callback TExtEventCallback) {
	extEventCallback = append(extEventCallback, callback)
}

// getParam 从指定索引和地址获取事件中的参数
// 不再使用Delphi导出的了，直接在这处理
func getParamOf(index int, ptr uintptr) uintptr {
	return *(*uintptr)(unsafe.Pointer(ptr + uintptr(index)*unsafe.Sizeof(ptr)))
}

// 回调过程
func eventCallbackProc(f uintptr, args uintptr, argcount int) uintptr {
	v, ok := EventCallbackOf(f)

	switch v.(type) {
	// func(sender IObject)
	case *TNotifyEvent:
		v = *(v.(*TNotifyEvent))

	case *TUDClickEvent:
		v = *(v.(*TUDClickEvent))

	case *TLVChangeEvent:
		v = *(v.(*TLVChangeEvent))

	case *TCloseEvent:
		v = *(v.(*TCloseEvent))

	case *TCloseQueryEvent:
		v = *(v.(*TCloseQueryEvent))

	case *TMenuChangeEvent:
		v = *(v.(*TMenuChangeEvent))

	case *TTVChangedEvent:
		v = *(v.(*TTVChangedEvent))

	case *TSysLinkEvent:
		v = *(v.(*TSysLinkEvent))

	case *TExceptionEvent:
		v = *(v.(*TExceptionEvent))

	case *TKeyEvent:
		v = *(v.(*TKeyEvent))

	case *TKeyPressEvent:
		v = *(v.(*TKeyPressEvent))

	case *TMouseEvent:
		v = *(v.(*TMouseEvent))

	case *TMouseMoveEvent:
		v = *(v.(*TMouseMoveEvent))

	case *TMouseWheelEvent:
		v = *(v.(*TMouseWheelEvent))

	case *TDrawItemEvent:
		v = *(v.(*TDrawItemEvent))

	case *TMenuDrawItemEvent:
		v = *(v.(*TMenuDrawItemEvent))

	case *TLVNotifyEvent:
		v = *(v.(*TLVNotifyEvent))

	case *TLVColumnClickEvent:
		v = *(v.(*TLVColumnClickEvent))

	case *TLVColumnRClickEvent:
		v = *(v.(*TLVColumnRClickEvent))

	case *TLVSelectItemEvent:
		v = *(v.(*TLVSelectItemEvent))

	case *TLVCheckedItemEvent:
		v = *(v.(*TLVCheckedItemEvent))

	case *TTabGetImageEvent:
		v = *(v.(*TTabGetImageEvent))

	case *TTVExpandedEvent:
		v = *(v.(*TTVExpandedEvent))

	case *TLVCompareEvent:
		v = *(v.(*TLVCompareEvent))

	case *TTVCompareEvent:
		v = *(v.(*TTVCompareEvent))

	case *TTVAdvancedCustomDrawEvent:
		v = *(v.(*TTVAdvancedCustomDrawEvent))

	case *TTVAdvancedCustomDrawItemEvent:
		v = *(v.(*TTVAdvancedCustomDrawItemEvent))

	case *TLVAdvancedCustomDrawEvent:
		v = *(v.(*TLVAdvancedCustomDrawEvent))

	case *TLVAdvancedCustomDrawItemEvent:
		v = *(v.(*TLVAdvancedCustomDrawItemEvent))

	case *TLVAdvancedCustomDrawSubItemEvent:
		v = *(v.(*TLVAdvancedCustomDrawSubItemEvent))

	case *TTBAdvancedCustomDrawEvent:
		v = *(v.(*TTBAdvancedCustomDrawEvent))

	case *TTBAdvancedCustomDrawBtnEvent:
		v = *(v.(*TTBAdvancedCustomDrawBtnEvent))

	case *TDropFilesEvent:
		v = *(v.(*TDropFilesEvent))

	case *TConstrainedResizeEvent:
		v = *(v.(*TConstrainedResizeEvent))

	case *THelpEvent:
		v = *(v.(*THelpEvent))

	case *TShortCutEvent:
		v = *(v.(*TShortCutEvent))

	case *TContextPopupEvent:
		v = *(v.(*TContextPopupEvent))

	case *TDragOverEvent:
		v = *(v.(*TDragOverEvent))

	case *TDragDropEvent:
		v = *(v.(*TDragDropEvent))

	case *TStartDragEvent:
		v = *(v.(*TStartDragEvent))

	case *TEndDragEvent:
		v = *(v.(*TEndDragEvent))

	case *TDockDropEvent:
		v = *(v.(*TDockDropEvent))

	case *TDockOverEvent:
		v = *(v.(*TDockOverEvent))

	case *TUnDockEvent:
		v = *(v.(*TUnDockEvent))

	case *TStartDockEvent:
		v = *(v.(*TStartDockEvent))

	case *TGetSiteInfoEvent:
		v = *(v.(*TGetSiteInfoEvent))

	case *TMouseWheelUpDownEvent:
		v = *(v.(*TMouseWheelUpDownEvent))

	case *TMessageEvent:
		v = *(v.(*TMessageEvent))

	case *TMovedEvent:
		v = *(v.(*TMovedEvent))

	case *TDrawCellEvent:
		v = *(v.(*TDrawCellEvent))

	case *TFixedCellClickEvent:
		v = *(v.(*TFixedCellClickEvent))

	case *TGetEditEvent:
		v = *(v.(*TGetEditEvent))

	case *TSelectCellEvent:
		v = *(v.(*TSelectCellEvent))

	case *TSetEditEvent:
		v = *(v.(*TSetEditEvent))

	case *TDrawSectionEvent:
		v = *(v.(*TDrawSectionEvent))

	case *TSectionNotifyEvent:
		v = *(v.(*TSectionNotifyEvent))

	case *TSectionTrackEvent:
		v = *(v.(*TSectionTrackEvent))

	case *TSectionDragEvent:
		v = *(v.(*TSectionDragEvent))

	case *TCustomSectionNotifyEvent:
		v = *(v.(*TCustomSectionNotifyEvent))

	// case *TGestureEvent:
	// 	v = *(v.(*TGestureEvent))

	case *TMouseActivateEvent:
		v = *(v.(*TMouseActivateEvent))

	case *TLBGetDataEvent:
		v = *(v.(*TLBGetDataEvent))

	case *TLBGetDataObjectEvent:
		v = *(v.(*TLBGetDataObjectEvent))

	case *TLBFindDataEvent:
		v = *(v.(*TLBFindDataEvent))

	case *TMeasureItemEvent:
		v = *(v.(*TMeasureItemEvent))

	case *TLVChangingEvent:
		v = *(v.(*TLVChangingEvent))

	case *TLVOwnerDataEvent:
		v = *(v.(*TLVOwnerDataEvent))

	case *TLVOwnerDataFindEvent:
		v = *(v.(*TLVOwnerDataFindEvent))

	case *TLVDeletedEvent:
		v = *(v.(*TLVDeletedEvent))

	case *TLVEditingEvent:
		v = *(v.(*TLVEditingEvent))

	case *TLVEditedEvent:
		v = *(v.(*TLVEditedEvent))

	case *TMenuMeasureItemEvent:
		v = *(v.(*TMenuMeasureItemEvent))

	case *TTabChangingEvent:
		v = *(v.(*TTabChangingEvent))

	case *TTVChangingEvent:
		v = *(v.(*TTVChangingEvent))

	case *TTVCollapsingEvent:
		v = *(v.(*TTVCollapsingEvent))

	case *TTVEditedEvent:
		v = *(v.(*TTVEditedEvent))

	case *TTVEditingEvent:
		v = *(v.(*TTVEditingEvent))

	case *TTVExpandingEvent:
		v = *(v.(*TTVExpandingEvent))

	case *TTVHintEvent:
		v = *(v.(*TTVHintEvent))

	case *TUDChangingEvent:
		v = *(v.(*TUDChangingEvent))

	case *TCreatingListErrorEvent:
		v = *(v.(*TCreatingListErrorEvent))

	case *TThumbPreviewItemRequestEvent:
		v = *(v.(*TThumbPreviewItemRequestEvent))

	case *TWindowPreviewItemRequestEvent:
		v = *(v.(*TWindowPreviewItemRequestEvent))

	case *TThumbButtonNotifyEvent:
		v = *(v.(*TThumbButtonNotifyEvent))

	case *TLVCustomDrawEvent:
		v = *(v.(*TLVCustomDrawEvent))

	case *TLVCustomDrawItemEvent:
		v = *(v.(*TLVCustomDrawItemEvent))

	case *TLVCustomDrawSubItemEvent:
		v = *(v.(*TLVCustomDrawSubItemEvent))

	case *TLVDrawItemEvent:
		v = *(v.(*TLVDrawItemEvent))

	case *TLVDataHintEvent:
		v = *(v.(*TLVDataHintEvent))

	case *TTVCustomDrawEvent:
		v = *(v.(*TTVCustomDrawEvent))

	case *TTVCustomDrawItemEvent:
		v = *(v.(*TTVCustomDrawItemEvent))

	case *TWebTitleChangeEvent:
		v = *(v.(*TWebTitleChangeEvent))

	case *TWebJSExternalEvent:
		v = *(v.(*TWebJSExternalEvent))

	case *TTaskDlgClickEvent:
		v = *(v.(*TTaskDlgClickEvent))

	case *TTaskDlgTimerEvent:
		v = *(v.(*TTaskDlgTimerEvent))

	case *TAlignPositionEvent:
		v = *(v.(*TAlignPositionEvent))

	case TNotifyEvent:
	default:
		fmt.Printf("unknown event func pointer: %#v, %T\n", v, v)
	}

	if ok {

		getVal := func(i int) uintptr {
			return getParamOf(i, args)
		}

		// 调用外部注册的事件回调过程
		for n := 0; n < len(extEventCallback); n++ {
			// 外部返回True则不继续下去了
			if extEventCallback[n](v, getVal) {
				return 0
			}
		}

		switch v.(type) {
		// func(sender IObject)
		case TNotifyEvent:
			v.(TNotifyEvent)(
				AsForm(getVal(0)))

		// func(sender IObject, button TUDBtnType)
		case TUDClickEvent:
			v.(TUDClickEvent)(
				AsObject(getVal(0)),
				TUDBtnType(getVal(1)))

		// func(sender IObject, item *TListItem, change int32)
		case TLVChangeEvent:
			v.(TLVChangeEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				TItemChange(getVal(2)))

		// func(sender IObject, action *TCloseAction) // Action *uintptr
		case TCloseEvent:
			v.(TCloseEvent)(
				AsObject(getVal(0)),
				(*TCloseAction)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, canClose *bool) //CanClose *uintptr
		case TCloseQueryEvent:
			v.(TCloseQueryEvent)(
				AsObject(getVal(0)),
				(*bool)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, source *TMenuItem, rebuild bool)
		case TMenuChangeEvent:
			v.(TMenuChangeEvent)(
				AsObject(getVal(0)),
				AsMenuItem(getVal(1)),
				DBoolToGoBool(getVal(2)))

		// func(sender IObject, node *TreeNode)
		case TTVChangedEvent:
			v.(TTVChangedEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)))

		// func(sender IObject, link string, linkType TSysLinkType) // TSysLinkType
		case TSysLinkEvent:
			v.(TSysLinkEvent)(
				AsObject(getVal(0)),
				DStrToGoStr(getVal(1)),
				TSysLinkType(getVal(2)))

		// func(sender, e IObject)
		case TExceptionEvent:
			v.(TExceptionEvent)(
				AsObject(getVal(0)),
				AsException(getVal(1)))

		// func(sender IObject, key *Char, shift TShiftState)
		case TKeyEvent:

			v.(TKeyEvent)(
				AsObject(getVal(0)),
				(*Char)(unsafe.Pointer(getVal(1))),
				TShiftState(getVal(2)))

		// func(sender IObject, key *Char)
		case TKeyPressEvent:

			v.(TKeyPressEvent)(
				AsObject(getVal(0)),
				(*Char)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)
		case TMouseEvent:
			v.(TMouseEvent)(
				AsObject(getVal(0)),
				TMouseButton(getVal(1)),
				TShiftState(getVal(2)),
				int32(getVal(3)),
				int32(getVal(4)))

		// func(sender IObject, shift TShiftState, x, y int32)
		case TMouseMoveEvent:
			v.(TMouseMoveEvent)(
				AsObject(getVal(0)),
				TShiftState(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

		// func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)
		case TMouseWheelEvent:
			v.(TMouseWheelEvent)(
				AsObject(getVal(0)),
				TShiftState(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)),
				int32(getVal(4)),
				(*bool)(unsafe.Pointer(getVal(5))))

			// func(control IWinControl, index int32, aRect TRect, state TOwnerDrawState)
		case TDrawItemEvent:
			v.(TDrawItemEvent)(
				AsWinControl(getVal(0)),
				int32(getVal(1)),
				*(*TRect)(unsafe.Pointer(getVal(2))),
				TOwnerDrawState(getVal(3)))

			// func(sender IObject, aCanvas *TCanvas, aRect TRect, selected bool)
		case TMenuDrawItemEvent:
			v.(TMenuDrawItemEvent)(
				AsObject(getVal(0)),
				AsCanvas(getVal(1)),
				*(*TRect)(unsafe.Pointer(getVal(2))),
				DBoolToGoBool(getVal(3)))

			// type TLVNotifyEvent func(sender IObject, item *TListItem)
		case TLVNotifyEvent:
			v.(TLVNotifyEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)))

			// type TLVColumnClickEvent func(sender IObject, column *TListColumn)
		case TLVColumnClickEvent:
			v.(TLVColumnClickEvent)(
				AsObject(getVal(0)),
				AsListColumn(getVal(1)))

			// type TLVColumnRClickEvent func(sender IObject, column *TListColumn, point TPoint)
		case TLVColumnRClickEvent:
			v.(TLVColumnRClickEvent)(
				AsObject(getVal(0)),
				AsListColumn(getVal(1)),
				TPoint{X: int32(getVal(2)), Y: int32(getVal(3))})

			// type TLVSelectItemEvent func(sender IObject, item *TListItem, selected bool)
		case TLVSelectItemEvent:
			v.(TLVSelectItemEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				DBoolToGoBool(getVal(2)))

			// type TLVCheckedItemEvent func(sender IObject, item *TListItem)
		case TLVCheckedItemEvent:
			v.(TLVCheckedItemEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)))

			// type TTabGetImageEvent func(sender IObject, tabIndex int32, imageIndex *int32)
		case TTabGetImageEvent:
			v.(TTabGetImageEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				(*int32)(unsafe.Pointer(getVal(2))))

			// type TTVExpandedEvent func(sender IObject, node *TTreeNode)
		case TTVExpandedEvent:
			v.(TTVExpandedEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)))

		//type TLVCompareEvent func(sender IObject, item1, item2 *TListItem, data int32, compare *int32)
		case TLVCompareEvent:
			v.(TLVCompareEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				AsListItem(getVal(2)),
				int32(getVal(3)),
				(*int32)(unsafe.Pointer(getVal(4))))

		//type TTVCompareEvent func(sender IObject, node1, node2 *TTreeNode, data int32, compare *int32)
		case TTVCompareEvent:
			v.(TTVCompareEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				AsTreeNode(getVal(2)),
				int32(getVal(3)),
				(*int32)(unsafe.Pointer(getVal(4))))

		//type TTVAdvancedCustomDrawEvent func(sender *TTreeView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TTVAdvancedCustomDrawEvent:

			v.(TTVAdvancedCustomDrawEvent)(
				AsTreeView(getVal(0)),
				*(*TRect)(unsafe.Pointer(getVal(1))),
				TCustomDrawStage(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

		//type TTVAdvancedCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, stage TCustomDrawStage, paintImages, defaultDraw *bool)
		case TTVAdvancedCustomDrawItemEvent:
			v.(TTVAdvancedCustomDrawItemEvent)(
				AsTreeView(getVal(0)),
				AsTreeNode(getVal(1)),
				TCustomDrawState(getVal(2)),
				TCustomDrawStage(getVal(3)),
				(*bool)(unsafe.Pointer(getVal(4))),
				(*bool)(unsafe.Pointer(getVal(5))))

			//---------------------------------------

		//type TLVAdvancedCustomDrawEvent func(sender *TListView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawEvent:
			v.(TLVAdvancedCustomDrawEvent)(
				AsListView(getVal(0)),
				*(*TRect)(unsafe.Pointer(getVal(1))),
				TCustomDrawStage(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

		//type TLVAdvancedCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawState, Stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawItemEvent:
			v.(TLVAdvancedCustomDrawItemEvent)(
				AsListView(getVal(0)),
				AsListItem(getVal(1)),
				TCustomDrawState(getVal(2)),
				TCustomDrawStage(getVal(3)),
				(*bool)(unsafe.Pointer(getVal(4))))

		//type TLVAdvancedCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawSubItemEvent:
			v.(TLVAdvancedCustomDrawSubItemEvent)(
				AsListView(getVal(0)),
				AsListItem(getVal(1)),
				int32(getVal(2)),
				TCustomDrawState(getVal(3)),
				TCustomDrawStage(getVal(4)),
				(*bool)(unsafe.Pointer(getVal(5))))

		//-----------------------------
		//type TTBAdvancedCustomDrawEvent func(sender *TToolBar, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TTBAdvancedCustomDrawEvent:
			v.(TTBAdvancedCustomDrawEvent)(
				AsToolBar(getVal(0)),
				*(*TRect)(unsafe.Pointer(getVal(1))),
				TCustomDrawStage(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

		//type TTBAdvancedCustomDrawBtnEvent func(sender *TToolBar, button *TToolButton, state TCustomDrawState, stage TCustomDrawStage, flags *TTBCustomDrawFlags, defaultDraw *bool)
		case TTBAdvancedCustomDrawBtnEvent:
			v.(TTBAdvancedCustomDrawBtnEvent)(
				AsToolBar(getVal(0)),
				AsToolButton(getVal(1)),
				TCustomDrawState(getVal(2)),
				TCustomDrawStage(getVal(3)),
				(*TTBCustomDrawFlags)(unsafe.Pointer(getVal(4))),
				(*bool)(unsafe.Pointer(getVal(5))))

		// TDropFilesEvent
		case TDropFilesEvent:
			nLen := int(getVal(2))
			tempArr := make([]string, nLen)
			p := getVal(1)
			for i := 0; i < nLen; i++ {
				tempArr[i] = DGetStringArrOf(p, i)
			}
			v.(TDropFilesEvent)(
				AsObject(getVal(0)),
				tempArr)

			// TConstrainedResizeEvent
		case TConstrainedResizeEvent:
			v.(TConstrainedResizeEvent)(
				AsObject(getVal(0)),
				(*int32)(unsafe.Pointer(getVal(1))),
				(*int32)(unsafe.Pointer(getVal(2))),
				(*int32)(unsafe.Pointer(getVal(3))),
				(*int32)(unsafe.Pointer(getVal(4))))

			// func(command uint16, data THelpEventData, callhelp *bool) bool
		case THelpEvent:
			v.(THelpEvent)(
				uint16(getVal(0)),
				THelpEventData(getVal(1)),
				(*bool)(unsafe.Pointer(getVal(2))),
				(*bool)(unsafe.Pointer(getVal(3))))

			// func(msg *TWMKey, handled *bool)
		case TShortCutEvent:
			v.(TShortCutEvent)(
				(*TWMKey)(unsafe.Pointer(getVal(0))),
				(*bool)(unsafe.Pointer(getVal(1))))

			// func(sender IObject, mousePos TPoint, handled *bool)
		case TContextPopupEvent:
			v.(TContextPopupEvent)(
				AsObject(getVal(0)),
				*(*TPoint)(unsafe.Pointer(getVal(1))),
				(*bool)(unsafe.Pointer(getVal(2))))

			// func(sender, source IObject, x, y int32, state TDragState, accept *bool)
		case TDragOverEvent:
			v.(TDragOverEvent)(
				AsObject(getVal(0)),
				AsObject(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)),
				TDragState(getVal(4)),
				(*bool)(unsafe.Pointer(getVal(5))))

			//func(sender, source IObject, x, y int32)
		case TDragDropEvent:
			v.(TDragDropEvent)(
				AsObject(getVal(0)),
				AsObject(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

			//func(sender IObject, dragObject *TDragObject)
		case TStartDragEvent:
			v.(TStartDragEvent)(
				AsObject(getVal(0)),
				AsDragObject(getVal(1)))

			//func(sender, target IObject, x, y int32)
		case TEndDragEvent:
			v.(TEndDragEvent)(
				AsObject(getVal(0)),
				AsObject(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

			// func(sender IObject, source *TDragDockObject, x, y int32)
		case TDockDropEvent:
			v.(TDockDropEvent)(
				AsObject(getVal(0)),
				AsDragDockObject(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

			//func(sender IObject, source *TDragDockObject, x, y int32, state TDragState, accept *bool)
		case TDockOverEvent:
			v.(TDockOverEvent)(
				AsObject(getVal(0)),
				AsDragDockObject(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)),
				TDragState(getVal(4)),
				(*bool)(unsafe.Pointer(getVal(5))))

			//func(sender IObject, client *TControl, newTarget *TControl, allow *bool)
		case TUnDockEvent:
			v.(TUnDockEvent)(
				AsObject(getVal(0)),
				AsControl(getVal(1)),
				AsControl(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

			//func(sender IObject, dragObject *TDragDockObject)
		case TStartDockEvent:
			v.(TStartDockEvent)(
				AsObject(getVal(0)),
				AsDragDockObject(getVal(1)))

			//func(sender IObject, dockClient *TControl, influenceRect *TRect, mousePos TPoint, canDock *bool)
		case TGetSiteInfoEvent:
			v.(TGetSiteInfoEvent)(
				AsObject(getVal(0)),
				AsControl(getVal(1)),
				(*TRect)(unsafe.Pointer(getVal(2))),
				*(*TPoint)(unsafe.Pointer(getVal(3))),
				(*bool)(unsafe.Pointer(getVal(4))))

			//func(sender IObject, shift TShiftState, mousePos TPoint, handled *bool)
		case TMouseWheelUpDownEvent:
			v.(TMouseWheelUpDownEvent)(
				AsObject(getVal(0)),
				TShiftState(getVal(1)),
				*(*TPoint)(unsafe.Pointer(getVal(2))),
				(*bool)(unsafe.Pointer(getVal(3))))

			// TMessageEvent = procedure (var Msg: TMsg; var Handled: Boolean) of object;
		case TMessageEvent: // func(msg *TMsg, handled *bool)
			v.(TMessageEvent)(
				(*TMsg)(unsafe.Pointer(getVal(0))),
				(*bool)(unsafe.Pointer(getVal(1))))

			// ---- grid
			//type TMovedEvent func(sender IObject, fromIndex, toIndex int32)
		case TMovedEvent:
			v.(TMovedEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)))

			//type TDrawCellEvent func(sender IObject, aCol, aRow int32, aRect TRect, state TGridDrawState)
		case TDrawCellEvent:
			v.(TDrawCellEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				*(*TRect)(unsafe.Pointer(getVal(3))),
				TGridDrawState(getVal(4)))

			//type TFixedCellClickEvent func(sender IObject, aCol, aRow int32)
		case TFixedCellClickEvent:
			v.(TFixedCellClickEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)))

			//type TGetEditEvent func(sender IObject, aCol, aRow int32, value *string)
		case TGetEditEvent:
			str := DStrToGoStr(*(*uintptr)(unsafe.Pointer(getVal(3))))
			v.(TGetEditEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				&str)
			*(*uintptr)(unsafe.Pointer(getVal(3))) = GoStrToDStr(str)

			//type TSelectCellEvent func(sender IObject, aCol, aRow int32, canSelect *bool)
		case TSelectCellEvent:
			v.(TSelectCellEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

			//type TSetEditEvent func(sender IObject, aCol, aRow int32, value string)
		case TSetEditEvent:
			v.(TSetEditEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				DStrToGoStr(getVal(3)))

			// ---- headercontrol
			//type TDrawSectionEvent func(headerControl *THeaderControl, section *THeaderSection, aRect TRect, pressed bool)
		case TDrawSectionEvent:
			v.(TDrawSectionEvent)(
				AsHeaderControl(getVal(0)),
				AsHeaderSection(getVal(1)),
				*(*TRect)(unsafe.Pointer(getVal(2))),
				getVal(3) != 0)

			//type TSectionNotifyEvent func(headerControl *THeaderControl, section *THeaderSection)
		case TSectionNotifyEvent:
			v.(TSectionNotifyEvent)(
				AsHeaderControl(getVal(0)),
				AsHeaderSection(getVal(1)))

			//type TSectionTrackEvent func(headerControl *THeaderControl, section *THeaderSection, width int32, state TSectionTrackState)
		case TSectionTrackEvent:
			v.(TSectionTrackEvent)(
				AsHeaderControl(getVal(0)),
				AsHeaderSection(getVal(1)),
				int32(getVal(2)),
				TSectionTrackState(getVal(3)))

			//type TSectionDragEvent func(sender IObject, fromSection, toSection *THeaderSection, allowDrag *bool)
		case TSectionDragEvent:
			v.(TSectionDragEvent)(
				AsObject(getVal(0)),
				AsHeaderSection(getVal(1)),
				AsHeaderSection(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

			//type TCustomSectionNotifyEvent func(headerControl *THeaderControl, section *THeaderSection)
		case TCustomSectionNotifyEvent:
			v.(TCustomSectionNotifyEvent)(
				AsHeaderControl(getVal(0)),
				AsHeaderSection(getVal(1)))

			////
			//type TGestureEvent func(sender IObject, eventInfo TGestureEventInfo, handled *bool)
		//case TGestureEvent:
		//	v.(TGestureEvent)(
		//		AsObject(getVal(0)),
		//		*(*TGestureEventInfo)(unsafe.Pointer(getVal(1))),
		//		(*bool)(unsafe.Pointer(getVal(2))))

		//type TMouseActivateEvent func(sender IObject, button TMouseButton, shift TShiftState, x, y int32, hitTest int32, mouseActivate *TMouseActivate)
		case TMouseActivateEvent:
			v.(TMouseActivateEvent)(
				AsObject(getVal(0)),
				TMouseButton(getVal(1)),
				TShiftState(getVal(2)),
				int32(getVal(3)),
				int32(getVal(4)),
				int32(getVal(5)),
				(*TMouseActivate)(unsafe.Pointer(getVal(6))))

			//type TLBGetDataEvent func(control *TWinControl, index int32, data *string)
		case TLBGetDataEvent:
			str := DStrToGoStr(*(*uintptr)(unsafe.Pointer(getVal(2))))
			v.(TLBGetDataEvent)(
				AsWinControl(getVal(0)),
				int32(getVal(1)),
				&str)
			*(*uintptr)(unsafe.Pointer(getVal(2))) = GoStrToDStr(str)

			//type TLBGetDataObjectEvent func(control *TWinControl, index int32, dataObject IObject)
		case TLBGetDataObjectEvent:
			//ptr := *(*uintptr)(unsafe.Pointer(getVal(2)))
			v.(TLBGetDataObjectEvent)(
				AsWinControl(getVal(0)),
				int32(getVal(1)),
				AsObject(getVal(2))) // 这个参数要改，先这样

			//type TLBFindDataEvent func(control *TWinControl, findString string) int32
		case TLBFindDataEvent:

			result := v.(TLBFindDataEvent)(
				AsWinControl(getVal(0)),
				DStrToGoStr(getVal(1)))
			*(*int32)(unsafe.Pointer(getVal(2))) = result

			//type TMeasureItemEvent func(control *TWinControl, index int32, height *int32)
		case TMeasureItemEvent:
			v.(TMeasureItemEvent)(
				AsWinControl(getVal(0)),
				int32(getVal(1)),
				(*int32)(unsafe.Pointer(getVal(2))))

			//type TLVChangingEvent func(sender IObject, item *TListItem, change TItemChange, allowChange *bool)
		case TLVChangingEvent:
			v.(TLVChangingEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				TItemChange(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

			//type TLVOwnerDataEvent func(sender IObject, item *TListItem)
		case TLVOwnerDataEvent:
			v.(TLVOwnerDataEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)))

			//type TLVOwnerDataFindEvent func(sender IObject, find TItemFind, findString string, findPosition TPoint, findData TCustomData, startIndex int32,
			//	direction TSearchDirection, warp bool, index *int32)
		case TLVOwnerDataFindEvent:
			v.(TLVOwnerDataFindEvent)(
				AsObject(getVal(0)),
				TItemFind(getVal(1)),
				DStrToGoStr(getVal(2)),
				*(*TPoint)(unsafe.Pointer(getVal(3))),
				TCustomData(getVal(4)),
				int32(getVal(5)),
				TSearchDirection(getVal(6)),
				DBoolToGoBool(getVal(7)),
				(*int32)(unsafe.Pointer(getVal(8))))

			//type TLVDeletedEvent func(sender IObject, item *TListItem)
		case TLVDeletedEvent:
			v.(TLVDeletedEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)))

			//type TLVEditingEvent func(sender IObject, item *TListItem, allowEdit *bool)
		case TLVEditingEvent:
			v.(TLVEditingEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				(*bool)(unsafe.Pointer(getVal(2))))

			//type TLVEditedEvent func(sender IObject, item *TListItem, s *string)
		case TLVEditedEvent:
			str := DStrToGoStr(*(*uintptr)(unsafe.Pointer(getVal(2))))
			v.(TLVEditedEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				&str)
			*(*uintptr)(unsafe.Pointer(getVal(2))) = GoStrToDStr(str)

			//type TMenuMeasureItemEvent func(sender IObject, aCanvas *TCanvas, width, height *int32)
		case TMenuMeasureItemEvent:
			v.(TMenuMeasureItemEvent)(
				AsObject(getVal(0)),
				AsCanvas(getVal(1)),
				(*int32)(unsafe.Pointer(getVal(2))),
				(*int32)(unsafe.Pointer(getVal(3))))

			//type TTabChangingEvent func(sender IObject, allowChange *bool)
		case TTabChangingEvent:
			v.(TTabChangingEvent)(
				AsObject(getVal(0)),
				(*bool)(unsafe.Pointer(getVal(1))))

			//type TTVChangingEvent func(sender IObject, node *TTreeNode, allowChange *bool)
		case TTVChangingEvent:
			v.(TTVChangingEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				(*bool)(unsafe.Pointer(getVal(2))))

			//type TTVCollapsingEvent func(sender IObject, node *TTreeNode, allowCollapse *bool)
		case TTVCollapsingEvent:
			v.(TTVCollapsingEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				(*bool)(unsafe.Pointer(getVal(2))))

			//type TTVEditedEvent func(sender IObject, node *TTreeNode, s *string)
		case TTVEditedEvent:
			str := DStrToGoStr(*(*uintptr)(unsafe.Pointer(getVal(2))))
			v.(TTVEditedEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				&str)
			*(*uintptr)(unsafe.Pointer(getVal(2))) = GoStrToDStr(str)

			//type TTVEditingEvent func(sender IObject, node *TTreeNode, allowEdit *bool)
		case TTVEditingEvent:
			v.(TTVEditingEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				(*bool)(unsafe.Pointer(getVal(2))))

			//type TTVExpandingEvent func(sender IObject, node *TTreeNode, allowExpansion *bool)
		case TTVExpandingEvent:
			v.(TTVExpandingEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				(*bool)(unsafe.Pointer(getVal(2))))

			//type TTVHintEvent func(sender IObject, node *TTreeNode, hint *string)
		case TTVHintEvent:
			str := DStrToGoStr(*(*uintptr)(unsafe.Pointer(getVal(2))))
			v.(TTVHintEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				&str)
			*(*uintptr)(unsafe.Pointer(getVal(2))) = GoStrToDStr(str)

			//type TUDChangingEvent func(sender IObject, allowChange *bool)
		case TUDChangingEvent:
			v.(TUDChangingEvent)(
				AsObject(getVal(0)),
				(*bool)(unsafe.Pointer(getVal(1))))

			//type TCreatingListErrorEvent func(sender IObject, winErrorCode uint32, errorDescription string, handled *bool)
		case TCreatingListErrorEvent:
			v.(TCreatingListErrorEvent)(
				AsObject(getVal(0)),
				uint32(getVal(1)),
				DStrToGoStr(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

			//type TThumbPreviewItemRequestEvent func(sender IObject, aPreviewHeight, aPreviewWidth int32, previewBitmap *TBitmap)
		case TThumbPreviewItemRequestEvent:
			v.(TThumbPreviewItemRequestEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				AsBitmap(getVal(3)))

			//type TWindowPreviewItemRequestEvent func(sender IObject, position *TPoint, previewBitmap *TBitmap)
		case TWindowPreviewItemRequestEvent:
			v.(TWindowPreviewItemRequestEvent)(
				AsObject(getVal(0)),
				(*TPoint)(unsafe.Pointer(getVal(1))),
				AsBitmap(getVal(2)))

			//type TThumbButtonNotifyEvent func(sender IObject, aButtonID int32)
		case TThumbButtonNotifyEvent:
			v.(TThumbButtonNotifyEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)))

			//--

			//type TLVCustomDrawEvent func(sender *TListView, aRect TRect, defaultDraw *bool)
		case TLVCustomDrawEvent:
			v.(TLVCustomDrawEvent)(
				AsListView(getVal(0)),
				*(*TRect)(unsafe.Pointer(getVal(1))),
				(*bool)(unsafe.Pointer(getVal(2))))

			//type TLVCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawStage, defaultDraw *bool)
		case TLVCustomDrawItemEvent:
			v.(TLVCustomDrawItemEvent)(
				AsListView(getVal(0)),
				AsListItem(getVal(1)),
				TCustomDrawStage(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

			//type TLVCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawStage, defaultDraw *bool)
		case TLVCustomDrawSubItemEvent:
			v.(TLVCustomDrawSubItemEvent)(
				AsListView(getVal(0)),
				AsListItem(getVal(1)),
				int32(getVal(2)),
				TCustomDrawStage(getVal(3)),
				(*bool)(unsafe.Pointer(getVal(4))))

			//type TLVDrawItemEvent func(sender *TListView, item *TListItem, rect TRect, state TOwnerDrawState)
		case TLVDrawItemEvent:
			v.(TLVDrawItemEvent)(
				AsListView(getVal(0)),
				AsListItem(getVal(1)),
				*(*TRect)(unsafe.Pointer(getVal(2))),
				TOwnerDrawState(getVal(3)))

		//type TLVDataHintEvent = func(sender IObject, startIndex, endIndex int32)
		case TLVDataHintEvent:
			v.(TLVDataHintEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)))

			//type TTVCustomDrawEvent func(sender *TTreeView, aRect TRect, defaultDraw *bool)
		case TTVCustomDrawEvent:
			v.(TTVCustomDrawEvent)(
				AsTreeView(getVal(0)),
				*(*TRect)(unsafe.Pointer(getVal(1))),
				(*bool)(unsafe.Pointer(getVal(2))))

			//type TTVCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawStage, defaultDraw *bool)
		case TTVCustomDrawItemEvent:
			v.(TTVCustomDrawItemEvent)(
				AsTreeView(getVal(0)),
				AsTreeNode(getVal(1)),
				TCustomDrawStage(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

			// type TWebTitleChangeEvent func(sender IObject, text string)
		case TWebTitleChangeEvent:
			v.(TWebTitleChangeEvent)(
				AsObject(getVal(0)),
				DStrToGoStr(getVal(1)))

		case TWebJSExternalEvent:
			str := DStrToGoStr(*(*uintptr)(unsafe.Pointer(getVal(3))))
			v.(TWebJSExternalEvent)(
				AsObject(getVal(0)),
				DStrToGoStr(getVal(1)),
				DStrToGoStr(getVal(2)),
				&str)
			*(*uintptr)(unsafe.Pointer(getVal(3))) = GoStrToDStr(str)

			//type TTaskDlgClickEvent func(sender IObject, modalResult TModalResult, canClose *bool)
		case TTaskDlgClickEvent:
			v.(TTaskDlgClickEvent)(
				AsObject(getVal(0)),
				TModalResult(getVal(1)),
				(*bool)(unsafe.Pointer(getVal(2))))

			//type TTaskDlgTimerEvent func(sender IObject, tickCount uint32, reset *bool)
		case TTaskDlgTimerEvent:
			v.(TTaskDlgTimerEvent)(
				AsObject(getVal(0)),
				uint32(getVal(1)),
				(*bool)(unsafe.Pointer(getVal(2))))

		// type TAlignPositionEvent func(sender *TWinControl, control *TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *TRect, alignInfo TAlignInfo)
		case TAlignPositionEvent:
			v.(TAlignPositionEvent)(
				AsWinControl(getVal(0)),
				AsControl(getVal(1)),
				(*int32)(unsafe.Pointer(getVal(2))),
				(*int32)(unsafe.Pointer(getVal(3))),
				(*int32)(unsafe.Pointer(getVal(4))),
				(*int32)(unsafe.Pointer(getVal(5))),
				(*TRect)(unsafe.Pointer(getVal(6))),
				*(*TAlignInfo)(unsafe.Pointer(getVal(7))))
		default:
		}
	}
	return 0
}
