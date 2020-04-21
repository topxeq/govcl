
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
	. "github.com/topxeq/govcl/vcl/api"
    . "github.com/topxeq/govcl/vcl/types"
    "unsafe"
)

type TMemo struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMemo(owner IComponent) *TMemo {
    m := new(TMemo)
    m.instance = Memo_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsMemo(obj interface{}) *TMemo {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TMemo{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsMemo.
func MemoFromInst(inst uintptr) *TMemo {
    return AsMemo(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsMemo.
func MemoFromObj(obj IObject) *TMemo {
    return AsMemo(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsMemo.
func MemoFromUnsafePointer(ptr unsafe.Pointer) *TMemo {
    return AsMemo(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (m *TMemo) Free() {
    if m.instance != 0 {
        Memo_Free(m.instance)
        m.instance, m.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMemo) Instance() uintptr {
    return m.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMemo) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMemo) IsValid() bool {
    return m.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (m *TMemo) Is() TIs {
    return TIs(m.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (m *TMemo) As() TAs {
//    return TAs(m.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMemoClass() TClass {
    return Memo_StaticClassType()
}

// CN: 清除。
// EN: .
func (m *TMemo) Clear() {
    Memo_Clear(m.instance)
}

// CN: 清除选择。
// EN: .
func (m *TMemo) ClearSelection() {
    Memo_ClearSelection(m.instance)
}

// CN: 复制到粘贴板。
// EN: .
func (m *TMemo) CopyToClipboard() {
    Memo_CopyToClipboard(m.instance)
}

// CN: 剪切到粘贴板。
// EN: .
func (m *TMemo) CutToClipboard() {
    Memo_CutToClipboard(m.instance)
}

// CN: 从剪切板粘贴。
// EN: .
func (m *TMemo) PasteFromClipboard() {
    Memo_PasteFromClipboard(m.instance)
}

// CN: 撤销上一次操作。
// EN: .
func (m *TMemo) Undo() {
    Memo_Undo(m.instance)
}

// CN: 全选。
// EN: .
func (m *TMemo) SelectAll() {
    Memo_SelectAll(m.instance)
}

// CN: 是否可以获得焦点。
// EN: .
func (m *TMemo) CanFocus() bool {
    return Memo_CanFocus(m.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (m *TMemo) ContainsControl(Control IControl) bool {
    return Memo_ContainsControl(m.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (m *TMemo) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(Memo_ControlAtPos(m.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (m *TMemo) DisableAlign() {
    Memo_DisableAlign(m.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (m *TMemo) EnableAlign() {
    Memo_EnableAlign(m.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (m *TMemo) FindChildControl(ControlName string) *TControl {
    return AsControl(Memo_FindChildControl(m.instance, ControlName))
}

func (m *TMemo) FlipChildren(AllLevels bool) {
    Memo_FlipChildren(m.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (m *TMemo) Focused() bool {
    return Memo_Focused(m.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (m *TMemo) HandleAllocated() bool {
    return Memo_HandleAllocated(m.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (m *TMemo) InsertControl(AControl IControl) {
    Memo_InsertControl(m.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (m *TMemo) Invalidate() {
    Memo_Invalidate(m.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (m *TMemo) RemoveControl(AControl IControl) {
    Memo_RemoveControl(m.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (m *TMemo) Realign() {
    Memo_Realign(m.instance)
}

// CN: 重绘。
// EN: Repaint.
func (m *TMemo) Repaint() {
    Memo_Repaint(m.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (m *TMemo) ScaleBy(M int32, D int32) {
    Memo_ScaleBy(m.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (m *TMemo) ScrollBy(DeltaX int32, DeltaY int32) {
    Memo_ScrollBy(m.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (m *TMemo) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Memo_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (m *TMemo) SetFocus() {
    Memo_SetFocus(m.instance)
}

// CN: 控件更新。
// EN: Update.
func (m *TMemo) Update() {
    Memo_Update(m.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (m *TMemo) BringToFront() {
    Memo_BringToFront(m.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (m *TMemo) ClientToScreen(Point TPoint) TPoint {
    return Memo_ClientToScreen(m.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (m *TMemo) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Memo_ClientToParent(m.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (m *TMemo) Dragging() bool {
    return Memo_Dragging(m.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (m *TMemo) HasParent() bool {
    return Memo_HasParent(m.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (m *TMemo) Hide() {
    Memo_Hide(m.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (m *TMemo) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Memo_Perform(m.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (m *TMemo) Refresh() {
    Memo_Refresh(m.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (m *TMemo) ScreenToClient(Point TPoint) TPoint {
    return Memo_ScreenToClient(m.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (m *TMemo) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Memo_ParentToClient(m.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (m *TMemo) SendToBack() {
    Memo_SendToBack(m.instance)
}

// CN: 显示控件。
// EN: Show control.
func (m *TMemo) Show() {
    Memo_Show(m.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (m *TMemo) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Memo_GetTextBuf(m.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (m *TMemo) GetTextLen() int32 {
    return Memo_GetTextLen(m.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (m *TMemo) SetTextBuf(Buffer string) {
    Memo_SetTextBuf(m.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (m *TMemo) FindComponent(AName string) *TComponent {
    return AsComponent(Memo_FindComponent(m.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (m *TMemo) GetNamePath() string {
    return Memo_GetNamePath(m.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (m *TMemo) Assign(Source IObject) {
    Memo_Assign(m.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMemo) ClassType() TClass {
    return Memo_ClassType(m.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMemo) ClassName() string {
    return Memo_ClassName(m.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMemo) InstanceSize() int32 {
    return Memo_InstanceSize(m.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMemo) InheritsFrom(AClass TClass) bool {
    return Memo_InheritsFrom(m.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMemo) Equals(Obj IObject) bool {
    return Memo_Equals(m.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMemo) GetHashCode() int32 {
    return Memo_GetHashCode(m.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (m *TMemo) ToString() string {
    return Memo_ToString(m.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (m *TMemo) Align() TAlign {
    return Memo_GetAlign(m.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (m *TMemo) SetAlign(value TAlign) {
    Memo_SetAlign(m.instance, value)
}

// CN: 获取文字对齐。
// EN: Get Text alignment.
func (m *TMemo) Alignment() TAlignment {
    return Memo_GetAlignment(m.instance)
}

// CN: 设置文字对齐。
// EN: Set Text alignment.
func (m *TMemo) SetAlignment(value TAlignment) {
    Memo_SetAlignment(m.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (m *TMemo) Anchors() TAnchors {
    return Memo_GetAnchors(m.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (m *TMemo) SetAnchors(value TAnchors) {
    Memo_SetAnchors(m.instance, value)
}

func (m *TMemo) BiDiMode() TBiDiMode {
    return Memo_GetBiDiMode(m.instance)
}

func (m *TMemo) SetBiDiMode(value TBiDiMode) {
    Memo_SetBiDiMode(m.instance, value)
}

// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (m *TMemo) BorderStyle() TBorderStyle {
    return Memo_GetBorderStyle(m.instance)
}

// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (m *TMemo) SetBorderStyle(value TBorderStyle) {
    Memo_SetBorderStyle(m.instance, value)
}

func (m *TMemo) CharCase() TEditCharCase {
    return Memo_GetCharCase(m.instance)
}

func (m *TMemo) SetCharCase(value TEditCharCase) {
    Memo_SetCharCase(m.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (m *TMemo) Color() TColor {
    return Memo_GetColor(m.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (m *TMemo) SetColor(value TColor) {
    Memo_SetColor(m.instance, value)
}

func (m *TMemo) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Memo_GetConstraints(m.instance))
}

func (m *TMemo) SetConstraints(value *TSizeConstraints) {
    Memo_SetConstraints(m.instance, CheckPtr(value))
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (m *TMemo) DoubleBuffered() bool {
    return Memo_GetDoubleBuffered(m.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (m *TMemo) SetDoubleBuffered(value bool) {
    Memo_SetDoubleBuffered(m.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (m *TMemo) DragCursor() TCursor {
    return Memo_GetDragCursor(m.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (m *TMemo) SetDragCursor(value TCursor) {
    Memo_SetDragCursor(m.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (m *TMemo) DragKind() TDragKind {
    return Memo_GetDragKind(m.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (m *TMemo) SetDragKind(value TDragKind) {
    Memo_SetDragKind(m.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (m *TMemo) DragMode() TDragMode {
    return Memo_GetDragMode(m.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (m *TMemo) SetDragMode(value TDragMode) {
    Memo_SetDragMode(m.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (m *TMemo) Enabled() bool {
    return Memo_GetEnabled(m.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (m *TMemo) SetEnabled(value bool) {
    Memo_SetEnabled(m.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (m *TMemo) Font() *TFont {
    return AsFont(Memo_GetFont(m.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (m *TMemo) SetFont(value *TFont) {
    Memo_SetFont(m.instance, CheckPtr(value))
}

// CN: 获取隐藏选择。
// EN: .
func (m *TMemo) HideSelection() bool {
    return Memo_GetHideSelection(m.instance)
}

// CN: 设置隐藏选择。
// EN: .
func (m *TMemo) SetHideSelection(value bool) {
    Memo_SetHideSelection(m.instance, value)
}

func (m *TMemo) Lines() *TStrings {
    return AsStrings(Memo_GetLines(m.instance))
}

func (m *TMemo) SetLines(value IObject) {
    Memo_SetLines(m.instance, CheckPtr(value))
}

// CN: 获取最大长度。
// EN: .
func (m *TMemo) MaxLength() int32 {
    return Memo_GetMaxLength(m.instance)
}

// CN: 设置最大长度。
// EN: .
func (m *TMemo) SetMaxLength(value int32) {
    Memo_SetMaxLength(m.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (m *TMemo) ParentColor() bool {
    return Memo_GetParentColor(m.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (m *TMemo) SetParentColor(value bool) {
    Memo_SetParentColor(m.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (m *TMemo) ParentDoubleBuffered() bool {
    return Memo_GetParentDoubleBuffered(m.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (m *TMemo) SetParentDoubleBuffered(value bool) {
    Memo_SetParentDoubleBuffered(m.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (m *TMemo) ParentFont() bool {
    return Memo_GetParentFont(m.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (m *TMemo) SetParentFont(value bool) {
    Memo_SetParentFont(m.instance, value)
}

func (m *TMemo) ParentShowHint() bool {
    return Memo_GetParentShowHint(m.instance)
}

func (m *TMemo) SetParentShowHint(value bool) {
    Memo_SetParentShowHint(m.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (m *TMemo) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Memo_GetPopupMenu(m.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (m *TMemo) SetPopupMenu(value IComponent) {
    Memo_SetPopupMenu(m.instance, CheckPtr(value))
}

// CN: 获取只读。
// EN: .
func (m *TMemo) ReadOnly() bool {
    return Memo_GetReadOnly(m.instance)
}

// CN: 设置只读。
// EN: .
func (m *TMemo) SetReadOnly(value bool) {
    Memo_SetReadOnly(m.instance, value)
}

func (m *TMemo) ScrollBars() TScrollStyle {
    return Memo_GetScrollBars(m.instance)
}

func (m *TMemo) SetScrollBars(value TScrollStyle) {
    Memo_SetScrollBars(m.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (m *TMemo) ShowHint() bool {
    return Memo_GetShowHint(m.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (m *TMemo) SetShowHint(value bool) {
    Memo_SetShowHint(m.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (m *TMemo) TabOrder() TTabOrder {
    return Memo_GetTabOrder(m.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (m *TMemo) SetTabOrder(value TTabOrder) {
    Memo_SetTabOrder(m.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (m *TMemo) TabStop() bool {
    return Memo_GetTabStop(m.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (m *TMemo) SetTabStop(value bool) {
    Memo_SetTabStop(m.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (m *TMemo) Visible() bool {
    return Memo_GetVisible(m.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (m *TMemo) SetVisible(value bool) {
    Memo_SetVisible(m.instance, value)
}

func (m *TMemo) WantReturns() bool {
    return Memo_GetWantReturns(m.instance)
}

func (m *TMemo) SetWantReturns(value bool) {
    Memo_SetWantReturns(m.instance, value)
}

func (m *TMemo) WantTabs() bool {
    return Memo_GetWantTabs(m.instance)
}

func (m *TMemo) SetWantTabs(value bool) {
    Memo_SetWantTabs(m.instance, value)
}

// CN: 获取自动换行。
// EN: Get Automatic line break.
func (m *TMemo) WordWrap() bool {
    return Memo_GetWordWrap(m.instance)
}

// CN: 设置自动换行。
// EN: Set Automatic line break.
func (m *TMemo) SetWordWrap(value bool) {
    Memo_SetWordWrap(m.instance, value)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (m *TMemo) SetOnChange(fn TNotifyEvent) {
    Memo_SetOnChange(m.instance, fn)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (m *TMemo) SetOnClick(fn TNotifyEvent) {
    Memo_SetOnClick(m.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (m *TMemo) SetOnContextPopup(fn TContextPopupEvent) {
    Memo_SetOnContextPopup(m.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (m *TMemo) SetOnDblClick(fn TNotifyEvent) {
    Memo_SetOnDblClick(m.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (m *TMemo) SetOnDragDrop(fn TDragDropEvent) {
    Memo_SetOnDragDrop(m.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (m *TMemo) SetOnDragOver(fn TDragOverEvent) {
    Memo_SetOnDragOver(m.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (m *TMemo) SetOnEndDrag(fn TEndDragEvent) {
    Memo_SetOnEndDrag(m.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (m *TMemo) SetOnEnter(fn TNotifyEvent) {
    Memo_SetOnEnter(m.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (m *TMemo) SetOnExit(fn TNotifyEvent) {
    Memo_SetOnExit(m.instance, fn)
}

// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (m *TMemo) SetOnKeyDown(fn TKeyEvent) {
    Memo_SetOnKeyDown(m.instance, fn)
}

func (m *TMemo) SetOnKeyPress(fn TKeyPressEvent) {
    Memo_SetOnKeyPress(m.instance, fn)
}

// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (m *TMemo) SetOnKeyUp(fn TKeyEvent) {
    Memo_SetOnKeyUp(m.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (m *TMemo) SetOnMouseDown(fn TMouseEvent) {
    Memo_SetOnMouseDown(m.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (m *TMemo) SetOnMouseEnter(fn TNotifyEvent) {
    Memo_SetOnMouseEnter(m.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (m *TMemo) SetOnMouseLeave(fn TNotifyEvent) {
    Memo_SetOnMouseLeave(m.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (m *TMemo) SetOnMouseMove(fn TMouseMoveEvent) {
    Memo_SetOnMouseMove(m.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (m *TMemo) SetOnMouseUp(fn TMouseEvent) {
    Memo_SetOnMouseUp(m.instance, fn)
}

func (m *TMemo) CaretPos() TPoint {
    return Memo_GetCaretPos(m.instance)
}

func (m *TMemo) SetCaretPos(value TPoint) {
    Memo_SetCaretPos(m.instance, value)
}

// CN: 获取能否撤销。
// EN: .
func (m *TMemo) CanUndo() bool {
    return Memo_GetCanUndo(m.instance)
}

// CN: 获取修改。
// EN: Get modified.
func (m *TMemo) Modified() bool {
    return Memo_GetModified(m.instance)
}

// CN: 设置修改。
// EN: Set modified.
func (m *TMemo) SetModified(value bool) {
    Memo_SetModified(m.instance, value)
}

// CN: 获取选择的长度。
// EN: .
func (m *TMemo) SelLength() int32 {
    return Memo_GetSelLength(m.instance)
}

// CN: 设置选择的长度。
// EN: .
func (m *TMemo) SetSelLength(value int32) {
    Memo_SetSelLength(m.instance, value)
}

// CN: 获取选择的启始位置。
// EN: .
func (m *TMemo) SelStart() int32 {
    return Memo_GetSelStart(m.instance)
}

// CN: 设置选择的启始位置。
// EN: .
func (m *TMemo) SetSelStart(value int32) {
    Memo_SetSelStart(m.instance, value)
}

// CN: 获取选择的文本。
// EN: .
func (m *TMemo) SelText() string {
    return Memo_GetSelText(m.instance)
}

// CN: 设置选择的文本。
// EN: .
func (m *TMemo) SetSelText(value string) {
    Memo_SetSelText(m.instance, value)
}

// CN: 获取文本。
// EN: .
func (m *TMemo) Text() string {
    strLen := m.GetTextLen()
    if strLen != 0 {
        var buffStr string
        m.GetTextBuf(&buffStr, strLen + 1)
        return buffStr
    }
    return ""
}

// CN: 设置文本。
// EN: .
func (m *TMemo) SetText(value string) {
    Memo_SetText(m.instance, value)
}

// CN: 获取提示文本。
// EN: .
func (m *TMemo) TextHint() string {
    return Memo_GetTextHint(m.instance)
}

// CN: 设置提示文本。
// EN: .
func (m *TMemo) SetTextHint(value string) {
    Memo_SetTextHint(m.instance, value)
}

// CN: 获取依靠客户端总数。
// EN: .
func (m *TMemo) DockClientCount() int32 {
    return Memo_GetDockClientCount(m.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (m *TMemo) DockSite() bool {
    return Memo_GetDockSite(m.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (m *TMemo) SetDockSite(value bool) {
    Memo_SetDockSite(m.instance, value)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (m *TMemo) MouseInClient() bool {
    return Memo_GetMouseInClient(m.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (m *TMemo) VisibleDockClientCount() int32 {
    return Memo_GetVisibleDockClientCount(m.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (m *TMemo) Brush() *TBrush {
    return AsBrush(Memo_GetBrush(m.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (m *TMemo) ControlCount() int32 {
    return Memo_GetControlCount(m.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (m *TMemo) Handle() HWND {
    return Memo_GetHandle(m.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (m *TMemo) ParentWindow() HWND {
    return Memo_GetParentWindow(m.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (m *TMemo) SetParentWindow(value HWND) {
    Memo_SetParentWindow(m.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (m *TMemo) UseDockManager() bool {
    return Memo_GetUseDockManager(m.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (m *TMemo) SetUseDockManager(value bool) {
    Memo_SetUseDockManager(m.instance, value)
}

func (m *TMemo) Action() *TAction {
    return AsAction(Memo_GetAction(m.instance))
}

func (m *TMemo) SetAction(value IComponent) {
    Memo_SetAction(m.instance, CheckPtr(value))
}

func (m *TMemo) BoundsRect() TRect {
    return Memo_GetBoundsRect(m.instance)
}

func (m *TMemo) SetBoundsRect(value TRect) {
    Memo_SetBoundsRect(m.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (m *TMemo) ClientHeight() int32 {
    return Memo_GetClientHeight(m.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (m *TMemo) SetClientHeight(value int32) {
    Memo_SetClientHeight(m.instance, value)
}

func (m *TMemo) ClientOrigin() TPoint {
    return Memo_GetClientOrigin(m.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (m *TMemo) ClientRect() TRect {
    return Memo_GetClientRect(m.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (m *TMemo) ClientWidth() int32 {
    return Memo_GetClientWidth(m.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (m *TMemo) SetClientWidth(value int32) {
    Memo_SetClientWidth(m.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (m *TMemo) ControlState() TControlState {
    return Memo_GetControlState(m.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (m *TMemo) SetControlState(value TControlState) {
    Memo_SetControlState(m.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (m *TMemo) ControlStyle() TControlStyle {
    return Memo_GetControlStyle(m.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (m *TMemo) SetControlStyle(value TControlStyle) {
    Memo_SetControlStyle(m.instance, value)
}

func (m *TMemo) Floating() bool {
    return Memo_GetFloating(m.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (m *TMemo) Parent() *TWinControl {
    return AsWinControl(Memo_GetParent(m.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (m *TMemo) SetParent(value IWinControl) {
    Memo_SetParent(m.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (m *TMemo) Left() int32 {
    return Memo_GetLeft(m.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (m *TMemo) SetLeft(value int32) {
    Memo_SetLeft(m.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (m *TMemo) Top() int32 {
    return Memo_GetTop(m.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (m *TMemo) SetTop(value int32) {
    Memo_SetTop(m.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (m *TMemo) Width() int32 {
    return Memo_GetWidth(m.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (m *TMemo) SetWidth(value int32) {
    Memo_SetWidth(m.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (m *TMemo) Height() int32 {
    return Memo_GetHeight(m.instance)
}

// CN: 设置高度。
// EN: Set height.
func (m *TMemo) SetHeight(value int32) {
    Memo_SetHeight(m.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (m *TMemo) Cursor() TCursor {
    return Memo_GetCursor(m.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (m *TMemo) SetCursor(value TCursor) {
    Memo_SetCursor(m.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (m *TMemo) Hint() string {
    return Memo_GetHint(m.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (m *TMemo) SetHint(value string) {
    Memo_SetHint(m.instance, value)
}

// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (m *TMemo) Margins() *TMargins {
    return AsMargins(Memo_GetMargins(m.instance))
}

// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (m *TMemo) SetMargins(value *TMargins) {
    Memo_SetMargins(m.instance, CheckPtr(value))
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (m *TMemo) ComponentCount() int32 {
    return Memo_GetComponentCount(m.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (m *TMemo) ComponentIndex() int32 {
    return Memo_GetComponentIndex(m.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (m *TMemo) SetComponentIndex(value int32) {
    Memo_SetComponentIndex(m.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (m *TMemo) Owner() *TComponent {
    return AsComponent(Memo_GetOwner(m.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (m *TMemo) Name() string {
    return Memo_GetName(m.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (m *TMemo) SetName(value string) {
    Memo_SetName(m.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (m *TMemo) Tag() int {
    return Memo_GetTag(m.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (m *TMemo) SetTag(value int) {
    Memo_SetTag(m.instance, value)
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (m *TMemo) DockClients(Index int32) *TControl {
    return AsControl(Memo_GetDockClients(m.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (m *TMemo) Controls(Index int32) *TControl {
    return AsControl(Memo_GetControls(m.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (m *TMemo) Components(AIndex int32) *TComponent {
    return AsComponent(Memo_GetComponents(m.instance, AIndex))
}

