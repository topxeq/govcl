//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"unsafe"

	. "github.com/topxeq/govcl/vcl/api"
	. "github.com/topxeq/govcl/vcl/types"
)

type TStringGrid struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewStringGrid(owner IComponent) *TStringGrid {
	s := new(TStringGrid)
	s.instance = StringGrid_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	// 不敢启用，因为不知道会发生什么...
	// runtime.SetFinalizer(s, (*TStringGrid).Free)
	return s
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsStringGrid(obj interface{}) *TStringGrid {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TStringGrid{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsStringGrid.
func StringGridFromInst(inst uintptr) *TStringGrid {
	return AsStringGrid(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsStringGrid.
func StringGridFromObj(obj IObject) *TStringGrid {
	return AsStringGrid(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsStringGrid.
func StringGridFromUnsafePointer(ptr unsafe.Pointer) *TStringGrid {
	return AsStringGrid(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (s *TStringGrid) Free() {
	if s.instance != 0 {
		StringGrid_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TStringGrid) Instance() uintptr {
	return s.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TStringGrid) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TStringGrid) IsValid() bool {
	return s.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (s *TStringGrid) Is() TIs {
	return TIs(s.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (s *TStringGrid) As() TAs {
//    return TAs(s.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TStringGridClass() TClass {
	return StringGrid_StaticClassType()
}

func (s *TStringGrid) CellRect(ACol int32, ARow int32) TRect {
	return StringGrid_CellRect(s.instance, ACol, ARow)
}

func (s *TStringGrid) MouseToCell(X int32, Y int32, ACol *int32, ARow *int32) {
	StringGrid_MouseToCell(s.instance, X, Y, ACol, ARow)
}

func (s *TStringGrid) MouseCoord(X int32, Y int32) TGridCoord {
	return StringGrid_MouseCoord(s.instance, X, Y)
}

// CN: 是否可以获得焦点。
// EN: .
func (s *TStringGrid) CanFocus() bool {
	return StringGrid_CanFocus(s.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (s *TStringGrid) ContainsControl(Control IControl) bool {
	return StringGrid_ContainsControl(s.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (s *TStringGrid) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(StringGrid_ControlAtPos(s.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (s *TStringGrid) DisableAlign() {
	StringGrid_DisableAlign(s.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (s *TStringGrid) EnableAlign() {
	StringGrid_EnableAlign(s.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (s *TStringGrid) FindChildControl(ControlName string) *TControl {
	return AsControl(StringGrid_FindChildControl(s.instance, ControlName))
}

func (s *TStringGrid) FlipChildren(AllLevels bool) {
	StringGrid_FlipChildren(s.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (s *TStringGrid) Focused() bool {
	return StringGrid_Focused(s.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (s *TStringGrid) HandleAllocated() bool {
	return StringGrid_HandleAllocated(s.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (s *TStringGrid) InsertControl(AControl IControl) {
	StringGrid_InsertControl(s.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (s *TStringGrid) Invalidate() {
	StringGrid_Invalidate(s.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (s *TStringGrid) RemoveControl(AControl IControl) {
	StringGrid_RemoveControl(s.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (s *TStringGrid) Realign() {
	StringGrid_Realign(s.instance)
}

// CN: 重绘。
// EN: Repaint.
func (s *TStringGrid) Repaint() {
	StringGrid_Repaint(s.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (s *TStringGrid) ScaleBy(M int32, D int32) {
	StringGrid_ScaleBy(s.instance, M, D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (s *TStringGrid) ScrollBy(DeltaX int32, DeltaY int32) {
	StringGrid_ScrollBy(s.instance, DeltaX, DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (s *TStringGrid) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	StringGrid_SetBounds(s.instance, ALeft, ATop, AWidth, AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (s *TStringGrid) SetFocus() {
	StringGrid_SetFocus(s.instance)
}

// CN: 控件更新。
// EN: Update.
func (s *TStringGrid) Update() {
	StringGrid_Update(s.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (s *TStringGrid) BringToFront() {
	StringGrid_BringToFront(s.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (s *TStringGrid) ClientToScreen(Point TPoint) TPoint {
	return StringGrid_ClientToScreen(s.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (s *TStringGrid) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return StringGrid_ClientToParent(s.instance, Point, CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (s *TStringGrid) Dragging() bool {
	return StringGrid_Dragging(s.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TStringGrid) HasParent() bool {
	return StringGrid_HasParent(s.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (s *TStringGrid) Hide() {
	StringGrid_Hide(s.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (s *TStringGrid) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return StringGrid_Perform(s.instance, Msg, WParam, LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (s *TStringGrid) Refresh() {
	StringGrid_Refresh(s.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (s *TStringGrid) ScreenToClient(Point TPoint) TPoint {
	return StringGrid_ScreenToClient(s.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (s *TStringGrid) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return StringGrid_ParentToClient(s.instance, Point, CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (s *TStringGrid) SendToBack() {
	StringGrid_SendToBack(s.instance)
}

// CN: 显示控件。
// EN: Show control.
func (s *TStringGrid) Show() {
	StringGrid_Show(s.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (s *TStringGrid) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return StringGrid_GetTextBuf(s.instance, Buffer, BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (s *TStringGrid) GetTextLen() int32 {
	return StringGrid_GetTextLen(s.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (s *TStringGrid) SetTextBuf(Buffer string) {
	StringGrid_SetTextBuf(s.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TStringGrid) FindComponent(AName string) *TComponent {
	return AsComponent(StringGrid_FindComponent(s.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TStringGrid) GetNamePath() string {
	return StringGrid_GetNamePath(s.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TStringGrid) Assign(Source IObject) {
	StringGrid_Assign(s.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TStringGrid) ClassType() TClass {
	return StringGrid_ClassType(s.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TStringGrid) ClassName() string {
	return StringGrid_ClassName(s.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TStringGrid) InstanceSize() int32 {
	return StringGrid_InstanceSize(s.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TStringGrid) InheritsFrom(AClass TClass) bool {
	return StringGrid_InheritsFrom(s.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TStringGrid) Equals(Obj IObject) bool {
	return StringGrid_Equals(s.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TStringGrid) GetHashCode() int32 {
	return StringGrid_GetHashCode(s.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (s *TStringGrid) ToString() string {
	return StringGrid_ToString(s.instance)
}

func (s *TStringGrid) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	StringGrid_AnchorToNeighbour(s.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (s *TStringGrid) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	StringGrid_AnchorParallel(s.instance, ASide, ASpace, CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (s *TStringGrid) AnchorHorizontalCenterTo(ASibling IControl) {
	StringGrid_AnchorHorizontalCenterTo(s.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (s *TStringGrid) AnchorVerticalCenterTo(ASibling IControl) {
	StringGrid_AnchorVerticalCenterTo(s.instance, CheckPtr(ASibling))
}

func (s *TStringGrid) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	StringGrid_AnchorAsAlign(s.instance, ATheAlign, ASpace)
}

func (s *TStringGrid) AnchorClient(ASpace int32) {
	StringGrid_AnchorClient(s.instance, ASpace)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (s *TStringGrid) Align() TAlign {
	return StringGrid_GetAlign(s.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (s *TStringGrid) SetAlign(value TAlign) {
	StringGrid_SetAlign(s.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (s *TStringGrid) Anchors() TAnchors {
	return StringGrid_GetAnchors(s.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (s *TStringGrid) SetAnchors(value TAnchors) {
	StringGrid_SetAnchors(s.instance, value)
}

func (s *TStringGrid) BiDiMode() TBiDiMode {
	return StringGrid_GetBiDiMode(s.instance)
}

func (s *TStringGrid) SetBiDiMode(value TBiDiMode) {
	StringGrid_SetBiDiMode(s.instance, value)
}

// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (s *TStringGrid) BorderStyle() TBorderStyle {
	return StringGrid_GetBorderStyle(s.instance)
}

// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (s *TStringGrid) SetBorderStyle(value TBorderStyle) {
	StringGrid_SetBorderStyle(s.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (s *TStringGrid) Color() TColor {
	return StringGrid_GetColor(s.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (s *TStringGrid) SetColor(value TColor) {
	StringGrid_SetColor(s.instance, value)
}

func (s *TStringGrid) ColCount() int32 {
	return StringGrid_GetColCount(s.instance)
}

func (s *TStringGrid) SetColCount(value int32) {
	StringGrid_SetColCount(s.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (s *TStringGrid) Constraints() *TSizeConstraints {
	return AsSizeConstraints(StringGrid_GetConstraints(s.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (s *TStringGrid) SetConstraints(value *TSizeConstraints) {
	StringGrid_SetConstraints(s.instance, CheckPtr(value))
}

func (s *TStringGrid) DefaultColWidth() int32 {
	return StringGrid_GetDefaultColWidth(s.instance)
}

func (s *TStringGrid) SetDefaultColWidth(value int32) {
	StringGrid_SetDefaultColWidth(s.instance, value)
}

func (s *TStringGrid) DefaultRowHeight() int32 {
	return StringGrid_GetDefaultRowHeight(s.instance)
}

func (s *TStringGrid) SetDefaultRowHeight(value int32) {
	StringGrid_SetDefaultRowHeight(s.instance, value)
}

func (s *TStringGrid) DefaultDrawing() bool {
	return StringGrid_GetDefaultDrawing(s.instance)
}

func (s *TStringGrid) SetDefaultDrawing(value bool) {
	StringGrid_SetDefaultDrawing(s.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (s *TStringGrid) DoubleBuffered() bool {
	return StringGrid_GetDoubleBuffered(s.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (s *TStringGrid) SetDoubleBuffered(value bool) {
	StringGrid_SetDoubleBuffered(s.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (s *TStringGrid) DragCursor() TCursor {
	return StringGrid_GetDragCursor(s.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (s *TStringGrid) SetDragCursor(value TCursor) {
	StringGrid_SetDragCursor(s.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (s *TStringGrid) DragKind() TDragKind {
	return StringGrid_GetDragKind(s.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (s *TStringGrid) SetDragKind(value TDragKind) {
	StringGrid_SetDragKind(s.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (s *TStringGrid) DragMode() TDragMode {
	return StringGrid_GetDragMode(s.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (s *TStringGrid) SetDragMode(value TDragMode) {
	StringGrid_SetDragMode(s.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TStringGrid) Enabled() bool {
	return StringGrid_GetEnabled(s.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TStringGrid) SetEnabled(value bool) {
	StringGrid_SetEnabled(s.instance, value)
}

func (s *TStringGrid) FixedColor() TColor {
	return StringGrid_GetFixedColor(s.instance)
}

func (s *TStringGrid) SetFixedColor(value TColor) {
	StringGrid_SetFixedColor(s.instance, value)
}

func (s *TStringGrid) FixedCols() int32 {
	return StringGrid_GetFixedCols(s.instance)
}

func (s *TStringGrid) SetFixedCols(value int32) {
	StringGrid_SetFixedCols(s.instance, value)
}

func (s *TStringGrid) RowCount() int32 {
	return StringGrid_GetRowCount(s.instance)
}

func (s *TStringGrid) SetRowCount(value int32) {
	StringGrid_SetRowCount(s.instance, value)
}

func (s *TStringGrid) FixedRows() int32 {
	return StringGrid_GetFixedRows(s.instance)
}

func (s *TStringGrid) SetFixedRows(value int32) {
	StringGrid_SetFixedRows(s.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (s *TStringGrid) Font() *TFont {
	return AsFont(StringGrid_GetFont(s.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (s *TStringGrid) SetFont(value *TFont) {
	StringGrid_SetFont(s.instance, CheckPtr(value))
}

func (s *TStringGrid) GridLineWidth() int32 {
	return StringGrid_GetGridLineWidth(s.instance)
}

func (s *TStringGrid) SetGridLineWidth(value int32) {
	StringGrid_SetGridLineWidth(s.instance, value)
}

func (s *TStringGrid) Options() TGridOptions {
	return StringGrid_GetOptions(s.instance)
}

func (s *TStringGrid) SetOptions(value TGridOptions) {
	StringGrid_SetOptions(s.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (s *TStringGrid) ParentColor() bool {
	return StringGrid_GetParentColor(s.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (s *TStringGrid) SetParentColor(value bool) {
	StringGrid_SetParentColor(s.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (s *TStringGrid) ParentDoubleBuffered() bool {
	return StringGrid_GetParentDoubleBuffered(s.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (s *TStringGrid) SetParentDoubleBuffered(value bool) {
	StringGrid_SetParentDoubleBuffered(s.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (s *TStringGrid) ParentFont() bool {
	return StringGrid_GetParentFont(s.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (s *TStringGrid) SetParentFont(value bool) {
	StringGrid_SetParentFont(s.instance, value)
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (s *TStringGrid) ParentShowHint() bool {
	return StringGrid_GetParentShowHint(s.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (s *TStringGrid) SetParentShowHint(value bool) {
	StringGrid_SetParentShowHint(s.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (s *TStringGrid) PopupMenu() *TPopupMenu {
	return AsPopupMenu(StringGrid_GetPopupMenu(s.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (s *TStringGrid) SetPopupMenu(value IComponent) {
	StringGrid_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TStringGrid) ScrollBars() TScrollStyle {
	return StringGrid_GetScrollBars(s.instance)
}

func (s *TStringGrid) SetScrollBars(value TScrollStyle) {
	StringGrid_SetScrollBars(s.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (s *TStringGrid) ShowHint() bool {
	return StringGrid_GetShowHint(s.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (s *TStringGrid) SetShowHint(value bool) {
	StringGrid_SetShowHint(s.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (s *TStringGrid) TabOrder() TTabOrder {
	return StringGrid_GetTabOrder(s.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (s *TStringGrid) SetTabOrder(value TTabOrder) {
	StringGrid_SetTabOrder(s.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TStringGrid) Visible() bool {
	return StringGrid_GetVisible(s.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TStringGrid) SetVisible(value bool) {
	StringGrid_SetVisible(s.instance, value)
}

func (s *TStringGrid) VisibleColCount() int32 {
	return StringGrid_GetVisibleColCount(s.instance)
}

func (s *TStringGrid) VisibleRowCount() int32 {
	return StringGrid_GetVisibleRowCount(s.instance)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (s *TStringGrid) SetOnClick(fn *TNotifyEvent) {
	StringGrid_SetOnClick(s.instance, fn)
}

func (s *TStringGrid) SetOnColumnMoved(fn *TMovedEvent) {
	StringGrid_SetOnColumnMoved(s.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (s *TStringGrid) SetOnContextPopup(fn *TContextPopupEvent) {
	StringGrid_SetOnContextPopup(s.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (s *TStringGrid) SetOnDblClick(fn *TNotifyEvent) {
	StringGrid_SetOnDblClick(s.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (s *TStringGrid) SetOnDragDrop(fn *TDragDropEvent) {
	StringGrid_SetOnDragDrop(s.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (s *TStringGrid) SetOnDragOver(fn *TDragOverEvent) {
	StringGrid_SetOnDragOver(s.instance, fn)
}

func (s *TStringGrid) SetOnDrawCell(fn *TDrawCellEvent) {
	StringGrid_SetOnDrawCell(s.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (s *TStringGrid) SetOnEndDock(fn *TEndDragEvent) {
	StringGrid_SetOnEndDock(s.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (s *TStringGrid) SetOnEndDrag(fn *TEndDragEvent) {
	StringGrid_SetOnEndDrag(s.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (s *TStringGrid) SetOnEnter(fn *TNotifyEvent) {
	StringGrid_SetOnEnter(s.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (s *TStringGrid) SetOnExit(fn *TNotifyEvent) {
	StringGrid_SetOnExit(s.instance, fn)
}

func (s *TStringGrid) SetOnGetEditMask(fn *TGetEditEvent) {
	StringGrid_SetOnGetEditMask(s.instance, fn)
}

func (s *TStringGrid) SetOnGetEditText(fn *TGetEditEvent) {
	StringGrid_SetOnGetEditText(s.instance, fn)
}

// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (s *TStringGrid) SetOnKeyDown(fn *TKeyEvent) {
	StringGrid_SetOnKeyDown(s.instance, fn)
}

func (s *TStringGrid) SetOnKeyPress(fn *TKeyPressEvent) {
	StringGrid_SetOnKeyPress(s.instance, fn)
}

// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (s *TStringGrid) SetOnKeyUp(fn *TKeyEvent) {
	StringGrid_SetOnKeyUp(s.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (s *TStringGrid) SetOnMouseDown(fn *TMouseEvent) {
	StringGrid_SetOnMouseDown(s.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (s *TStringGrid) SetOnMouseEnter(fn *TNotifyEvent) {
	StringGrid_SetOnMouseEnter(s.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (s *TStringGrid) SetOnMouseLeave(fn *TNotifyEvent) {
	StringGrid_SetOnMouseLeave(s.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (s *TStringGrid) SetOnMouseMove(fn *TMouseMoveEvent) {
	StringGrid_SetOnMouseMove(s.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (s *TStringGrid) SetOnMouseUp(fn *TMouseEvent) {
	StringGrid_SetOnMouseUp(s.instance, fn)
}

// CN: 设置鼠标滚轮按下事件。
// EN: .
func (s *TStringGrid) SetOnMouseWheelDown(fn *TMouseWheelUpDownEvent) {
	StringGrid_SetOnMouseWheelDown(s.instance, fn)
}

// CN: 设置鼠标滚轮抬起事件。
// EN: .
func (s *TStringGrid) SetOnMouseWheelUp(fn *TMouseWheelUpDownEvent) {
	StringGrid_SetOnMouseWheelUp(s.instance, fn)
}

func (s *TStringGrid) SetOnRowMoved(fn *TMovedEvent) {
	StringGrid_SetOnRowMoved(s.instance, fn)
}

func (s *TStringGrid) SetOnSelectCell(fn *TSelectCellEvent) {
	StringGrid_SetOnSelectCell(s.instance, fn)
}

func (s *TStringGrid) SetOnSetEditText(fn *TSetEditEvent) {
	StringGrid_SetOnSetEditText(s.instance, fn)
}

// CN: 设置启动停靠。
// EN: .
func (s *TStringGrid) SetOnStartDock(fn *TStartDockEvent) {
	StringGrid_SetOnStartDock(s.instance, fn)
}

func (s *TStringGrid) SetOnTopLeftChanged(fn *TNotifyEvent) {
	StringGrid_SetOnTopLeftChanged(s.instance, fn)
}

// CN: 获取画布。
// EN: .
func (s *TStringGrid) Canvas() *TCanvas {
	return AsCanvas(StringGrid_GetCanvas(s.instance))
}

func (s *TStringGrid) Col() int32 {
	return StringGrid_GetCol(s.instance)
}

func (s *TStringGrid) SetCol(value int32) {
	StringGrid_SetCol(s.instance, value)
}

func (s *TStringGrid) EditorMode() bool {
	return StringGrid_GetEditorMode(s.instance)
}

func (s *TStringGrid) SetEditorMode(value bool) {
	StringGrid_SetEditorMode(s.instance, value)
}

func (s *TStringGrid) GridHeight() int32 {
	return StringGrid_GetGridHeight(s.instance)
}

func (s *TStringGrid) GridWidth() int32 {
	return StringGrid_GetGridWidth(s.instance)
}

func (s *TStringGrid) LeftCol() int32 {
	return StringGrid_GetLeftCol(s.instance)
}

func (s *TStringGrid) SetLeftCol(value int32) {
	StringGrid_SetLeftCol(s.instance, value)
}

func (s *TStringGrid) Selection() TGridRect {
	return StringGrid_GetSelection(s.instance)
}

func (s *TStringGrid) SetSelection(value TGridRect) {
	StringGrid_SetSelection(s.instance, value)
}

func (s *TStringGrid) Row() int32 {
	return StringGrid_GetRow(s.instance)
}

func (s *TStringGrid) SetRow(value int32) {
	StringGrid_SetRow(s.instance, value)
}

func (s *TStringGrid) TopRow() int32 {
	return StringGrid_GetTopRow(s.instance)
}

func (s *TStringGrid) SetTopRow(value int32) {
	StringGrid_SetTopRow(s.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (s *TStringGrid) TabStop() bool {
	return StringGrid_GetTabStop(s.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (s *TStringGrid) SetTabStop(value bool) {
	StringGrid_SetTabStop(s.instance, value)
}

// CN: 获取依靠客户端总数。
// EN: .
func (s *TStringGrid) DockClientCount() int32 {
	return StringGrid_GetDockClientCount(s.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (s *TStringGrid) DockSite() bool {
	return StringGrid_GetDockSite(s.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (s *TStringGrid) SetDockSite(value bool) {
	StringGrid_SetDockSite(s.instance, value)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (s *TStringGrid) MouseInClient() bool {
	return StringGrid_GetMouseInClient(s.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (s *TStringGrid) VisibleDockClientCount() int32 {
	return StringGrid_GetVisibleDockClientCount(s.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (s *TStringGrid) Brush() *TBrush {
	return AsBrush(StringGrid_GetBrush(s.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (s *TStringGrid) ControlCount() int32 {
	return StringGrid_GetControlCount(s.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (s *TStringGrid) Handle() HWND {
	return StringGrid_GetHandle(s.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (s *TStringGrid) ParentWindow() HWND {
	return StringGrid_GetParentWindow(s.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (s *TStringGrid) SetParentWindow(value HWND) {
	StringGrid_SetParentWindow(s.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (s *TStringGrid) UseDockManager() bool {
	return StringGrid_GetUseDockManager(s.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (s *TStringGrid) SetUseDockManager(value bool) {
	StringGrid_SetUseDockManager(s.instance, value)
}

func (s *TStringGrid) Action() *TAction {
	return AsAction(StringGrid_GetAction(s.instance))
}

func (s *TStringGrid) SetAction(value IComponent) {
	StringGrid_SetAction(s.instance, CheckPtr(value))
}

func (s *TStringGrid) BoundsRect() TRect {
	return StringGrid_GetBoundsRect(s.instance)
}

func (s *TStringGrid) SetBoundsRect(value TRect) {
	StringGrid_SetBoundsRect(s.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (s *TStringGrid) ClientHeight() int32 {
	return StringGrid_GetClientHeight(s.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (s *TStringGrid) SetClientHeight(value int32) {
	StringGrid_SetClientHeight(s.instance, value)
}

func (s *TStringGrid) ClientOrigin() TPoint {
	return StringGrid_GetClientOrigin(s.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (s *TStringGrid) ClientRect() TRect {
	return StringGrid_GetClientRect(s.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (s *TStringGrid) ClientWidth() int32 {
	return StringGrid_GetClientWidth(s.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (s *TStringGrid) SetClientWidth(value int32) {
	StringGrid_SetClientWidth(s.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (s *TStringGrid) ControlState() TControlState {
	return StringGrid_GetControlState(s.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (s *TStringGrid) SetControlState(value TControlState) {
	StringGrid_SetControlState(s.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (s *TStringGrid) ControlStyle() TControlStyle {
	return StringGrid_GetControlStyle(s.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (s *TStringGrid) SetControlStyle(value TControlStyle) {
	StringGrid_SetControlStyle(s.instance, value)
}

func (s *TStringGrid) Floating() bool {
	return StringGrid_GetFloating(s.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TStringGrid) Parent() *TWinControl {
	return AsWinControl(StringGrid_GetParent(s.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TStringGrid) SetParent(value IWinControl) {
	StringGrid_SetParent(s.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (s *TStringGrid) Left() int32 {
	return StringGrid_GetLeft(s.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (s *TStringGrid) SetLeft(value int32) {
	StringGrid_SetLeft(s.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (s *TStringGrid) Top() int32 {
	return StringGrid_GetTop(s.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (s *TStringGrid) SetTop(value int32) {
	StringGrid_SetTop(s.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (s *TStringGrid) Width() int32 {
	return StringGrid_GetWidth(s.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (s *TStringGrid) SetWidth(value int32) {
	StringGrid_SetWidth(s.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (s *TStringGrid) Height() int32 {
	return StringGrid_GetHeight(s.instance)
}

// CN: 设置高度。
// EN: Set height.
func (s *TStringGrid) SetHeight(value int32) {
	StringGrid_SetHeight(s.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TStringGrid) Cursor() TCursor {
	return StringGrid_GetCursor(s.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TStringGrid) SetCursor(value TCursor) {
	StringGrid_SetCursor(s.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (s *TStringGrid) Hint() string {
	return StringGrid_GetHint(s.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (s *TStringGrid) SetHint(value string) {
	StringGrid_SetHint(s.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TStringGrid) ComponentCount() int32 {
	return StringGrid_GetComponentCount(s.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (s *TStringGrid) ComponentIndex() int32 {
	return StringGrid_GetComponentIndex(s.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (s *TStringGrid) SetComponentIndex(value int32) {
	StringGrid_SetComponentIndex(s.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TStringGrid) Owner() *TComponent {
	return AsComponent(StringGrid_GetOwner(s.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (s *TStringGrid) Name() string {
	return StringGrid_GetName(s.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (s *TStringGrid) SetName(value string) {
	StringGrid_SetName(s.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TStringGrid) Tag() int {
	return StringGrid_GetTag(s.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TStringGrid) SetTag(value int) {
	StringGrid_SetTag(s.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (s *TStringGrid) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(StringGrid_GetAnchorSideLeft(s.instance))
}

// CN: 设置左边锚点。
// EN: .
func (s *TStringGrid) SetAnchorSideLeft(value *TAnchorSide) {
	StringGrid_SetAnchorSideLeft(s.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (s *TStringGrid) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(StringGrid_GetAnchorSideTop(s.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (s *TStringGrid) SetAnchorSideTop(value *TAnchorSide) {
	StringGrid_SetAnchorSideTop(s.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (s *TStringGrid) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(StringGrid_GetAnchorSideRight(s.instance))
}

// CN: 设置右边锚点。
// EN: .
func (s *TStringGrid) SetAnchorSideRight(value *TAnchorSide) {
	StringGrid_SetAnchorSideRight(s.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (s *TStringGrid) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(StringGrid_GetAnchorSideBottom(s.instance))
}

// CN: 设置底边锚点。
// EN: .
func (s *TStringGrid) SetAnchorSideBottom(value *TAnchorSide) {
	StringGrid_SetAnchorSideBottom(s.instance, CheckPtr(value))
}

func (s *TStringGrid) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(StringGrid_GetChildSizing(s.instance))
}

func (s *TStringGrid) SetChildSizing(value *TControlChildSizing) {
	StringGrid_SetChildSizing(s.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (s *TStringGrid) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(StringGrid_GetBorderSpacing(s.instance))
}

// CN: 设置边框间距。
// EN: .
func (s *TStringGrid) SetBorderSpacing(value *TControlBorderSpacing) {
	StringGrid_SetBorderSpacing(s.instance, CheckPtr(value))
}

func (s *TStringGrid) Cells(ACol int32, ARow int32) string {
	return StringGrid_GetCells(s.instance, ACol, ARow)
}

func (s *TStringGrid) SetCells(ACol int32, ARow int32, value string) {
	StringGrid_SetCells(s.instance, ACol, ARow, value)
}

func (s *TStringGrid) Cols(Index int32) *TStrings {
	return AsStrings(StringGrid_GetCols(s.instance, Index))
}

func (s *TStringGrid) SetCols(Index int32, value IObject) {
	StringGrid_SetCols(s.instance, Index, CheckPtr(value))
}

func (s *TStringGrid) Objects(ACol int32, ARow int32) *TObject {
	return AsObject(StringGrid_GetObjects(s.instance, ACol, ARow))
}

func (s *TStringGrid) SetObjects(ACol int32, ARow int32, value IObject) {
	StringGrid_SetObjects(s.instance, ACol, ARow, CheckPtr(value))
}

func (s *TStringGrid) Rows(Index int32) *TStrings {
	return AsStrings(StringGrid_GetRows(s.instance, Index))
}

func (s *TStringGrid) SetRows(Index int32, value IObject) {
	StringGrid_SetRows(s.instance, Index, CheckPtr(value))
}

func (s *TStringGrid) ColWidths(Index int32) int32 {
	return StringGrid_GetColWidths(s.instance, Index)
}

func (s *TStringGrid) SetColWidths(Index int32, value int32) {
	StringGrid_SetColWidths(s.instance, Index, value)
}

func (s *TStringGrid) RowHeights(Index int32) int32 {
	return StringGrid_GetRowHeights(s.instance, Index)
}

func (s *TStringGrid) SetRowHeights(Index int32, value int32) {
	StringGrid_SetRowHeights(s.instance, Index, value)
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (s *TStringGrid) DockClients(Index int32) *TControl {
	return AsControl(StringGrid_GetDockClients(s.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (s *TStringGrid) Controls(Index int32) *TControl {
	return AsControl(StringGrid_GetControls(s.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TStringGrid) Components(AIndex int32) *TComponent {
	return AsComponent(StringGrid_GetComponents(s.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (s *TStringGrid) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(StringGrid_GetAnchorSide(s.instance, AKind))
}
