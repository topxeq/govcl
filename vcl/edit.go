
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

type TEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewEdit(owner IComponent) *TEdit {
    e := new(TEdit)
    e.instance = Edit_Create(CheckPtr(owner))
    e.ptr = unsafe.Pointer(e.instance)
    return e
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsEdit(obj interface{}) *TEdit {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TEdit{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsEdit.
func EditFromInst(inst uintptr) *TEdit {
    return AsEdit(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsEdit.
func EditFromObj(obj IObject) *TEdit {
    return AsEdit(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsEdit.
func EditFromUnsafePointer(ptr unsafe.Pointer) *TEdit {
    return AsEdit(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (e *TEdit) Free() {
    if e.instance != 0 {
        Edit_Free(e.instance)
        e.instance, e.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (e *TEdit) Instance() uintptr {
    return e.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (e *TEdit) UnsafeAddr() unsafe.Pointer {
    return e.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (e *TEdit) IsValid() bool {
    return e.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (e *TEdit) Is() TIs {
    return TIs(e.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (e *TEdit) As() TAs {
//    return TAs(e.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TEditClass() TClass {
    return Edit_StaticClassType()
}

// CN: 清除。
// EN: .
func (e *TEdit) Clear() {
    Edit_Clear(e.instance)
}

// CN: 清除选择。
// EN: .
func (e *TEdit) ClearSelection() {
    Edit_ClearSelection(e.instance)
}

// CN: 复制到粘贴板。
// EN: .
func (e *TEdit) CopyToClipboard() {
    Edit_CopyToClipboard(e.instance)
}

// CN: 剪切到粘贴板。
// EN: .
func (e *TEdit) CutToClipboard() {
    Edit_CutToClipboard(e.instance)
}

// CN: 从剪切板粘贴。
// EN: .
func (e *TEdit) PasteFromClipboard() {
    Edit_PasteFromClipboard(e.instance)
}

// CN: 撤销上一次操作。
// EN: .
func (e *TEdit) Undo() {
    Edit_Undo(e.instance)
}

// CN: 全选。
// EN: .
func (e *TEdit) SelectAll() {
    Edit_SelectAll(e.instance)
}

// CN: 是否可以获得焦点。
// EN: .
func (e *TEdit) CanFocus() bool {
    return Edit_CanFocus(e.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (e *TEdit) ContainsControl(Control IControl) bool {
    return Edit_ContainsControl(e.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (e *TEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(Edit_ControlAtPos(e.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (e *TEdit) DisableAlign() {
    Edit_DisableAlign(e.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (e *TEdit) EnableAlign() {
    Edit_EnableAlign(e.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (e *TEdit) FindChildControl(ControlName string) *TControl {
    return AsControl(Edit_FindChildControl(e.instance, ControlName))
}

func (e *TEdit) FlipChildren(AllLevels bool) {
    Edit_FlipChildren(e.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (e *TEdit) Focused() bool {
    return Edit_Focused(e.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (e *TEdit) HandleAllocated() bool {
    return Edit_HandleAllocated(e.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (e *TEdit) InsertControl(AControl IControl) {
    Edit_InsertControl(e.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (e *TEdit) Invalidate() {
    Edit_Invalidate(e.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (e *TEdit) RemoveControl(AControl IControl) {
    Edit_RemoveControl(e.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (e *TEdit) Realign() {
    Edit_Realign(e.instance)
}

// CN: 重绘。
// EN: Repaint.
func (e *TEdit) Repaint() {
    Edit_Repaint(e.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (e *TEdit) ScaleBy(M int32, D int32) {
    Edit_ScaleBy(e.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (e *TEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    Edit_ScrollBy(e.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (e *TEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Edit_SetBounds(e.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (e *TEdit) SetFocus() {
    Edit_SetFocus(e.instance)
}

// CN: 控件更新。
// EN: Update.
func (e *TEdit) Update() {
    Edit_Update(e.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (e *TEdit) BringToFront() {
    Edit_BringToFront(e.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (e *TEdit) ClientToScreen(Point TPoint) TPoint {
    return Edit_ClientToScreen(e.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (e *TEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Edit_ClientToParent(e.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (e *TEdit) Dragging() bool {
    return Edit_Dragging(e.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (e *TEdit) HasParent() bool {
    return Edit_HasParent(e.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (e *TEdit) Hide() {
    Edit_Hide(e.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (e *TEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Edit_Perform(e.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (e *TEdit) Refresh() {
    Edit_Refresh(e.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (e *TEdit) ScreenToClient(Point TPoint) TPoint {
    return Edit_ScreenToClient(e.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (e *TEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Edit_ParentToClient(e.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (e *TEdit) SendToBack() {
    Edit_SendToBack(e.instance)
}

// CN: 显示控件。
// EN: Show control.
func (e *TEdit) Show() {
    Edit_Show(e.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (e *TEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Edit_GetTextBuf(e.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (e *TEdit) GetTextLen() int32 {
    return Edit_GetTextLen(e.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (e *TEdit) SetTextBuf(Buffer string) {
    Edit_SetTextBuf(e.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (e *TEdit) FindComponent(AName string) *TComponent {
    return AsComponent(Edit_FindComponent(e.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (e *TEdit) GetNamePath() string {
    return Edit_GetNamePath(e.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (e *TEdit) Assign(Source IObject) {
    Edit_Assign(e.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (e *TEdit) ClassType() TClass {
    return Edit_ClassType(e.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (e *TEdit) ClassName() string {
    return Edit_ClassName(e.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (e *TEdit) InstanceSize() int32 {
    return Edit_InstanceSize(e.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (e *TEdit) InheritsFrom(AClass TClass) bool {
    return Edit_InheritsFrom(e.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (e *TEdit) Equals(Obj IObject) bool {
    return Edit_Equals(e.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (e *TEdit) GetHashCode() int32 {
    return Edit_GetHashCode(e.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (e *TEdit) ToString() string {
    return Edit_ToString(e.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (e *TEdit) Align() TAlign {
    return Edit_GetAlign(e.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (e *TEdit) SetAlign(value TAlign) {
    Edit_SetAlign(e.instance, value)
}

// CN: 获取文字对齐。
// EN: Get Text alignment.
func (e *TEdit) Alignment() TAlignment {
    return Edit_GetAlignment(e.instance)
}

// CN: 设置文字对齐。
// EN: Set Text alignment.
func (e *TEdit) SetAlignment(value TAlignment) {
    Edit_SetAlignment(e.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (e *TEdit) Anchors() TAnchors {
    return Edit_GetAnchors(e.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (e *TEdit) SetAnchors(value TAnchors) {
    Edit_SetAnchors(e.instance, value)
}

// CN: 获取自动选择。
// EN: .
func (e *TEdit) AutoSelect() bool {
    return Edit_GetAutoSelect(e.instance)
}

// CN: 设置自动选择。
// EN: .
func (e *TEdit) SetAutoSelect(value bool) {
    Edit_SetAutoSelect(e.instance, value)
}

// CN: 获取自动调整大小。
// EN: .
func (e *TEdit) AutoSize() bool {
    return Edit_GetAutoSize(e.instance)
}

// CN: 设置自动调整大小。
// EN: .
func (e *TEdit) SetAutoSize(value bool) {
    Edit_SetAutoSize(e.instance, value)
}

func (e *TEdit) BiDiMode() TBiDiMode {
    return Edit_GetBiDiMode(e.instance)
}

func (e *TEdit) SetBiDiMode(value TBiDiMode) {
    Edit_SetBiDiMode(e.instance, value)
}

// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (e *TEdit) BorderStyle() TBorderStyle {
    return Edit_GetBorderStyle(e.instance)
}

// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (e *TEdit) SetBorderStyle(value TBorderStyle) {
    Edit_SetBorderStyle(e.instance, value)
}

func (e *TEdit) CharCase() TEditCharCase {
    return Edit_GetCharCase(e.instance)
}

func (e *TEdit) SetCharCase(value TEditCharCase) {
    Edit_SetCharCase(e.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (e *TEdit) Color() TColor {
    return Edit_GetColor(e.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (e *TEdit) SetColor(value TColor) {
    Edit_SetColor(e.instance, value)
}

func (e *TEdit) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Edit_GetConstraints(e.instance))
}

func (e *TEdit) SetConstraints(value *TSizeConstraints) {
    Edit_SetConstraints(e.instance, CheckPtr(value))
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (e *TEdit) DoubleBuffered() bool {
    return Edit_GetDoubleBuffered(e.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (e *TEdit) SetDoubleBuffered(value bool) {
    Edit_SetDoubleBuffered(e.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (e *TEdit) DragCursor() TCursor {
    return Edit_GetDragCursor(e.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (e *TEdit) SetDragCursor(value TCursor) {
    Edit_SetDragCursor(e.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (e *TEdit) DragKind() TDragKind {
    return Edit_GetDragKind(e.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (e *TEdit) SetDragKind(value TDragKind) {
    Edit_SetDragKind(e.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (e *TEdit) DragMode() TDragMode {
    return Edit_GetDragMode(e.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (e *TEdit) SetDragMode(value TDragMode) {
    Edit_SetDragMode(e.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (e *TEdit) Enabled() bool {
    return Edit_GetEnabled(e.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (e *TEdit) SetEnabled(value bool) {
    Edit_SetEnabled(e.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (e *TEdit) Font() *TFont {
    return AsFont(Edit_GetFont(e.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (e *TEdit) SetFont(value *TFont) {
    Edit_SetFont(e.instance, CheckPtr(value))
}

// CN: 获取隐藏选择。
// EN: .
func (e *TEdit) HideSelection() bool {
    return Edit_GetHideSelection(e.instance)
}

// CN: 设置隐藏选择。
// EN: .
func (e *TEdit) SetHideSelection(value bool) {
    Edit_SetHideSelection(e.instance, value)
}

// CN: 获取最大长度。
// EN: .
func (e *TEdit) MaxLength() int32 {
    return Edit_GetMaxLength(e.instance)
}

// CN: 设置最大长度。
// EN: .
func (e *TEdit) SetMaxLength(value int32) {
    Edit_SetMaxLength(e.instance, value)
}

// CN: 获取只能输入数字。
// EN: .
func (e *TEdit) NumbersOnly() bool {
    return Edit_GetNumbersOnly(e.instance)
}

// CN: 设置只能输入数字。
// EN: .
func (e *TEdit) SetNumbersOnly(value bool) {
    Edit_SetNumbersOnly(e.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (e *TEdit) ParentColor() bool {
    return Edit_GetParentColor(e.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (e *TEdit) SetParentColor(value bool) {
    Edit_SetParentColor(e.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (e *TEdit) ParentDoubleBuffered() bool {
    return Edit_GetParentDoubleBuffered(e.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (e *TEdit) SetParentDoubleBuffered(value bool) {
    Edit_SetParentDoubleBuffered(e.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (e *TEdit) ParentFont() bool {
    return Edit_GetParentFont(e.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (e *TEdit) SetParentFont(value bool) {
    Edit_SetParentFont(e.instance, value)
}

func (e *TEdit) ParentShowHint() bool {
    return Edit_GetParentShowHint(e.instance)
}

func (e *TEdit) SetParentShowHint(value bool) {
    Edit_SetParentShowHint(e.instance, value)
}

// CN: 获取密码掩码字符。
// EN: .
func (e *TEdit) PasswordChar() uint16 {
    return Edit_GetPasswordChar(e.instance)
}

// CN: 设置密码掩码字符。
// EN: .
func (e *TEdit) SetPasswordChar(value uint16) {
    Edit_SetPasswordChar(e.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (e *TEdit) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Edit_GetPopupMenu(e.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (e *TEdit) SetPopupMenu(value IComponent) {
    Edit_SetPopupMenu(e.instance, CheckPtr(value))
}

// CN: 获取只读。
// EN: .
func (e *TEdit) ReadOnly() bool {
    return Edit_GetReadOnly(e.instance)
}

// CN: 设置只读。
// EN: .
func (e *TEdit) SetReadOnly(value bool) {
    Edit_SetReadOnly(e.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (e *TEdit) ShowHint() bool {
    return Edit_GetShowHint(e.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (e *TEdit) SetShowHint(value bool) {
    Edit_SetShowHint(e.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (e *TEdit) TabOrder() TTabOrder {
    return Edit_GetTabOrder(e.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (e *TEdit) SetTabOrder(value TTabOrder) {
    Edit_SetTabOrder(e.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (e *TEdit) TabStop() bool {
    return Edit_GetTabStop(e.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (e *TEdit) SetTabStop(value bool) {
    Edit_SetTabStop(e.instance, value)
}

// CN: 获取文本。
// EN: .
func (e *TEdit) Text() string {
    strLen := e.GetTextLen()
    if strLen != 0 {
        var buffStr string
        e.GetTextBuf(&buffStr, strLen + 1)
        return buffStr
    }
    return ""
}

// CN: 设置文本。
// EN: .
func (e *TEdit) SetText(value string) {
    Edit_SetText(e.instance, value)
}

// CN: 获取提示文本。
// EN: .
func (e *TEdit) TextHint() string {
    return Edit_GetTextHint(e.instance)
}

// CN: 设置提示文本。
// EN: .
func (e *TEdit) SetTextHint(value string) {
    Edit_SetTextHint(e.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (e *TEdit) Visible() bool {
    return Edit_GetVisible(e.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (e *TEdit) SetVisible(value bool) {
    Edit_SetVisible(e.instance, value)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (e *TEdit) SetOnChange(fn TNotifyEvent) {
    Edit_SetOnChange(e.instance, fn)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (e *TEdit) SetOnClick(fn TNotifyEvent) {
    Edit_SetOnClick(e.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (e *TEdit) SetOnContextPopup(fn TContextPopupEvent) {
    Edit_SetOnContextPopup(e.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (e *TEdit) SetOnDblClick(fn TNotifyEvent) {
    Edit_SetOnDblClick(e.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (e *TEdit) SetOnDragDrop(fn TDragDropEvent) {
    Edit_SetOnDragDrop(e.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (e *TEdit) SetOnDragOver(fn TDragOverEvent) {
    Edit_SetOnDragOver(e.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (e *TEdit) SetOnEndDrag(fn TEndDragEvent) {
    Edit_SetOnEndDrag(e.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (e *TEdit) SetOnEnter(fn TNotifyEvent) {
    Edit_SetOnEnter(e.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (e *TEdit) SetOnExit(fn TNotifyEvent) {
    Edit_SetOnExit(e.instance, fn)
}

// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (e *TEdit) SetOnKeyDown(fn TKeyEvent) {
    Edit_SetOnKeyDown(e.instance, fn)
}

func (e *TEdit) SetOnKeyPress(fn TKeyPressEvent) {
    Edit_SetOnKeyPress(e.instance, fn)
}

// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (e *TEdit) SetOnKeyUp(fn TKeyEvent) {
    Edit_SetOnKeyUp(e.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (e *TEdit) SetOnMouseDown(fn TMouseEvent) {
    Edit_SetOnMouseDown(e.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (e *TEdit) SetOnMouseEnter(fn TNotifyEvent) {
    Edit_SetOnMouseEnter(e.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (e *TEdit) SetOnMouseLeave(fn TNotifyEvent) {
    Edit_SetOnMouseLeave(e.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (e *TEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    Edit_SetOnMouseMove(e.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (e *TEdit) SetOnMouseUp(fn TMouseEvent) {
    Edit_SetOnMouseUp(e.instance, fn)
}

// CN: 获取能否撤销。
// EN: .
func (e *TEdit) CanUndo() bool {
    return Edit_GetCanUndo(e.instance)
}

// CN: 获取修改。
// EN: Get modified.
func (e *TEdit) Modified() bool {
    return Edit_GetModified(e.instance)
}

// CN: 设置修改。
// EN: Set modified.
func (e *TEdit) SetModified(value bool) {
    Edit_SetModified(e.instance, value)
}

// CN: 获取选择的长度。
// EN: .
func (e *TEdit) SelLength() int32 {
    return Edit_GetSelLength(e.instance)
}

// CN: 设置选择的长度。
// EN: .
func (e *TEdit) SetSelLength(value int32) {
    Edit_SetSelLength(e.instance, value)
}

// CN: 获取选择的启始位置。
// EN: .
func (e *TEdit) SelStart() int32 {
    return Edit_GetSelStart(e.instance)
}

// CN: 设置选择的启始位置。
// EN: .
func (e *TEdit) SetSelStart(value int32) {
    Edit_SetSelStart(e.instance, value)
}

// CN: 获取选择的文本。
// EN: .
func (e *TEdit) SelText() string {
    return Edit_GetSelText(e.instance)
}

// CN: 设置选择的文本。
// EN: .
func (e *TEdit) SetSelText(value string) {
    Edit_SetSelText(e.instance, value)
}

// CN: 获取依靠客户端总数。
// EN: .
func (e *TEdit) DockClientCount() int32 {
    return Edit_GetDockClientCount(e.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (e *TEdit) DockSite() bool {
    return Edit_GetDockSite(e.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (e *TEdit) SetDockSite(value bool) {
    Edit_SetDockSite(e.instance, value)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (e *TEdit) MouseInClient() bool {
    return Edit_GetMouseInClient(e.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (e *TEdit) VisibleDockClientCount() int32 {
    return Edit_GetVisibleDockClientCount(e.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (e *TEdit) Brush() *TBrush {
    return AsBrush(Edit_GetBrush(e.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (e *TEdit) ControlCount() int32 {
    return Edit_GetControlCount(e.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (e *TEdit) Handle() HWND {
    return Edit_GetHandle(e.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (e *TEdit) ParentWindow() HWND {
    return Edit_GetParentWindow(e.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (e *TEdit) SetParentWindow(value HWND) {
    Edit_SetParentWindow(e.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (e *TEdit) UseDockManager() bool {
    return Edit_GetUseDockManager(e.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (e *TEdit) SetUseDockManager(value bool) {
    Edit_SetUseDockManager(e.instance, value)
}

func (e *TEdit) Action() *TAction {
    return AsAction(Edit_GetAction(e.instance))
}

func (e *TEdit) SetAction(value IComponent) {
    Edit_SetAction(e.instance, CheckPtr(value))
}

func (e *TEdit) BoundsRect() TRect {
    return Edit_GetBoundsRect(e.instance)
}

func (e *TEdit) SetBoundsRect(value TRect) {
    Edit_SetBoundsRect(e.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (e *TEdit) ClientHeight() int32 {
    return Edit_GetClientHeight(e.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (e *TEdit) SetClientHeight(value int32) {
    Edit_SetClientHeight(e.instance, value)
}

func (e *TEdit) ClientOrigin() TPoint {
    return Edit_GetClientOrigin(e.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (e *TEdit) ClientRect() TRect {
    return Edit_GetClientRect(e.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (e *TEdit) ClientWidth() int32 {
    return Edit_GetClientWidth(e.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (e *TEdit) SetClientWidth(value int32) {
    Edit_SetClientWidth(e.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (e *TEdit) ControlState() TControlState {
    return Edit_GetControlState(e.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (e *TEdit) SetControlState(value TControlState) {
    Edit_SetControlState(e.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (e *TEdit) ControlStyle() TControlStyle {
    return Edit_GetControlStyle(e.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (e *TEdit) SetControlStyle(value TControlStyle) {
    Edit_SetControlStyle(e.instance, value)
}

func (e *TEdit) Floating() bool {
    return Edit_GetFloating(e.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (e *TEdit) Parent() *TWinControl {
    return AsWinControl(Edit_GetParent(e.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (e *TEdit) SetParent(value IWinControl) {
    Edit_SetParent(e.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (e *TEdit) Left() int32 {
    return Edit_GetLeft(e.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (e *TEdit) SetLeft(value int32) {
    Edit_SetLeft(e.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (e *TEdit) Top() int32 {
    return Edit_GetTop(e.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (e *TEdit) SetTop(value int32) {
    Edit_SetTop(e.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (e *TEdit) Width() int32 {
    return Edit_GetWidth(e.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (e *TEdit) SetWidth(value int32) {
    Edit_SetWidth(e.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (e *TEdit) Height() int32 {
    return Edit_GetHeight(e.instance)
}

// CN: 设置高度。
// EN: Set height.
func (e *TEdit) SetHeight(value int32) {
    Edit_SetHeight(e.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (e *TEdit) Cursor() TCursor {
    return Edit_GetCursor(e.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (e *TEdit) SetCursor(value TCursor) {
    Edit_SetCursor(e.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (e *TEdit) Hint() string {
    return Edit_GetHint(e.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (e *TEdit) SetHint(value string) {
    Edit_SetHint(e.instance, value)
}

// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (e *TEdit) Margins() *TMargins {
    return AsMargins(Edit_GetMargins(e.instance))
}

// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (e *TEdit) SetMargins(value *TMargins) {
    Edit_SetMargins(e.instance, CheckPtr(value))
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (e *TEdit) ComponentCount() int32 {
    return Edit_GetComponentCount(e.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (e *TEdit) ComponentIndex() int32 {
    return Edit_GetComponentIndex(e.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (e *TEdit) SetComponentIndex(value int32) {
    Edit_SetComponentIndex(e.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (e *TEdit) Owner() *TComponent {
    return AsComponent(Edit_GetOwner(e.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (e *TEdit) Name() string {
    return Edit_GetName(e.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (e *TEdit) SetName(value string) {
    Edit_SetName(e.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (e *TEdit) Tag() int {
    return Edit_GetTag(e.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (e *TEdit) SetTag(value int) {
    Edit_SetTag(e.instance, value)
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (e *TEdit) DockClients(Index int32) *TControl {
    return AsControl(Edit_GetDockClients(e.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (e *TEdit) Controls(Index int32) *TControl {
    return AsControl(Edit_GetControls(e.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (e *TEdit) Components(AIndex int32) *TComponent {
    return AsComponent(Edit_GetComponents(e.instance, AIndex))
}

