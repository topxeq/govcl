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

type TSpinEdit struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewSpinEdit(owner IComponent) *TSpinEdit {
	s := new(TSpinEdit)
	s.instance = SpinEdit_Create(CheckPtr(owner))
	s.ptr = unsafe.Pointer(s.instance)
	// 不敢启用，因为不知道会发生什么...
	// runtime.SetFinalizer(s, (*TSpinEdit).Free)
	return s
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsSpinEdit(obj interface{}) *TSpinEdit {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TSpinEdit{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsSpinEdit.
func SpinEditFromInst(inst uintptr) *TSpinEdit {
	return AsSpinEdit(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsSpinEdit.
func SpinEditFromObj(obj IObject) *TSpinEdit {
	return AsSpinEdit(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsSpinEdit.
func SpinEditFromUnsafePointer(ptr unsafe.Pointer) *TSpinEdit {
	return AsSpinEdit(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (s *TSpinEdit) Free() {
	if s.instance != 0 {
		SpinEdit_Free(s.instance)
		s.instance, s.ptr = 0, nullptr
	}
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TSpinEdit) Instance() uintptr {
	return s.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TSpinEdit) UnsafeAddr() unsafe.Pointer {
	return s.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TSpinEdit) IsValid() bool {
	return s.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (s *TSpinEdit) Is() TIs {
	return TIs(s.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (s *TSpinEdit) As() TAs {
//    return TAs(s.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TSpinEditClass() TClass {
	return SpinEdit_StaticClassType()
}

// CN: 清除。
// EN: .
func (s *TSpinEdit) Clear() {
	SpinEdit_Clear(s.instance)
}

// CN: 清除选择。
// EN: .
func (s *TSpinEdit) ClearSelection() {
	SpinEdit_ClearSelection(s.instance)
}

// CN: 复制到粘贴板。
// EN: .
func (s *TSpinEdit) CopyToClipboard() {
	SpinEdit_CopyToClipboard(s.instance)
}

// CN: 剪切到粘贴板。
// EN: .
func (s *TSpinEdit) CutToClipboard() {
	SpinEdit_CutToClipboard(s.instance)
}

// CN: 从剪切板粘贴。
// EN: .
func (s *TSpinEdit) PasteFromClipboard() {
	SpinEdit_PasteFromClipboard(s.instance)
}

// CN: 撤销上一次操作。
// EN: .
func (s *TSpinEdit) Undo() {
	SpinEdit_Undo(s.instance)
}

// CN: 全选。
// EN: .
func (s *TSpinEdit) SelectAll() {
	SpinEdit_SelectAll(s.instance)
}

// CN: 是否可以获得焦点。
// EN: .
func (s *TSpinEdit) CanFocus() bool {
	return SpinEdit_CanFocus(s.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (s *TSpinEdit) ContainsControl(Control IControl) bool {
	return SpinEdit_ContainsControl(s.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (s *TSpinEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return AsControl(SpinEdit_ControlAtPos(s.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (s *TSpinEdit) DisableAlign() {
	SpinEdit_DisableAlign(s.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (s *TSpinEdit) EnableAlign() {
	SpinEdit_EnableAlign(s.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (s *TSpinEdit) FindChildControl(ControlName string) *TControl {
	return AsControl(SpinEdit_FindChildControl(s.instance, ControlName))
}

func (s *TSpinEdit) FlipChildren(AllLevels bool) {
	SpinEdit_FlipChildren(s.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (s *TSpinEdit) Focused() bool {
	return SpinEdit_Focused(s.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (s *TSpinEdit) HandleAllocated() bool {
	return SpinEdit_HandleAllocated(s.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (s *TSpinEdit) InsertControl(AControl IControl) {
	SpinEdit_InsertControl(s.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (s *TSpinEdit) Invalidate() {
	SpinEdit_Invalidate(s.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (s *TSpinEdit) RemoveControl(AControl IControl) {
	SpinEdit_RemoveControl(s.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (s *TSpinEdit) Realign() {
	SpinEdit_Realign(s.instance)
}

// CN: 重绘。
// EN: Repaint.
func (s *TSpinEdit) Repaint() {
	SpinEdit_Repaint(s.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (s *TSpinEdit) ScaleBy(M int32, D int32) {
	SpinEdit_ScaleBy(s.instance, M, D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (s *TSpinEdit) ScrollBy(DeltaX int32, DeltaY int32) {
	SpinEdit_ScrollBy(s.instance, DeltaX, DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (s *TSpinEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	SpinEdit_SetBounds(s.instance, ALeft, ATop, AWidth, AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (s *TSpinEdit) SetFocus() {
	SpinEdit_SetFocus(s.instance)
}

// CN: 控件更新。
// EN: Update.
func (s *TSpinEdit) Update() {
	SpinEdit_Update(s.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (s *TSpinEdit) BringToFront() {
	SpinEdit_BringToFront(s.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (s *TSpinEdit) ClientToScreen(Point TPoint) TPoint {
	return SpinEdit_ClientToScreen(s.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (s *TSpinEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return SpinEdit_ClientToParent(s.instance, Point, CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (s *TSpinEdit) Dragging() bool {
	return SpinEdit_Dragging(s.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TSpinEdit) HasParent() bool {
	return SpinEdit_HasParent(s.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (s *TSpinEdit) Hide() {
	SpinEdit_Hide(s.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (s *TSpinEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return SpinEdit_Perform(s.instance, Msg, WParam, LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (s *TSpinEdit) Refresh() {
	SpinEdit_Refresh(s.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (s *TSpinEdit) ScreenToClient(Point TPoint) TPoint {
	return SpinEdit_ScreenToClient(s.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (s *TSpinEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return SpinEdit_ParentToClient(s.instance, Point, CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (s *TSpinEdit) SendToBack() {
	SpinEdit_SendToBack(s.instance)
}

// CN: 显示控件。
// EN: Show control.
func (s *TSpinEdit) Show() {
	SpinEdit_Show(s.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (s *TSpinEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
	return SpinEdit_GetTextBuf(s.instance, Buffer, BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (s *TSpinEdit) GetTextLen() int32 {
	return SpinEdit_GetTextLen(s.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (s *TSpinEdit) SetTextBuf(Buffer string) {
	SpinEdit_SetTextBuf(s.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TSpinEdit) FindComponent(AName string) *TComponent {
	return AsComponent(SpinEdit_FindComponent(s.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TSpinEdit) GetNamePath() string {
	return SpinEdit_GetNamePath(s.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TSpinEdit) Assign(Source IObject) {
	SpinEdit_Assign(s.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TSpinEdit) ClassType() TClass {
	return SpinEdit_ClassType(s.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TSpinEdit) ClassName() string {
	return SpinEdit_ClassName(s.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TSpinEdit) InstanceSize() int32 {
	return SpinEdit_InstanceSize(s.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TSpinEdit) InheritsFrom(AClass TClass) bool {
	return SpinEdit_InheritsFrom(s.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TSpinEdit) Equals(Obj IObject) bool {
	return SpinEdit_Equals(s.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TSpinEdit) GetHashCode() int32 {
	return SpinEdit_GetHashCode(s.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (s *TSpinEdit) ToString() string {
	return SpinEdit_ToString(s.instance)
}

func (s *TSpinEdit) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	SpinEdit_AnchorToNeighbour(s.instance, ASide, ASpace, CheckPtr(ASibling))
}

func (s *TSpinEdit) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
	SpinEdit_AnchorParallel(s.instance, ASide, ASpace, CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (s *TSpinEdit) AnchorHorizontalCenterTo(ASibling IControl) {
	SpinEdit_AnchorHorizontalCenterTo(s.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (s *TSpinEdit) AnchorVerticalCenterTo(ASibling IControl) {
	SpinEdit_AnchorVerticalCenterTo(s.instance, CheckPtr(ASibling))
}

func (s *TSpinEdit) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
	SpinEdit_AnchorAsAlign(s.instance, ATheAlign, ASpace)
}

func (s *TSpinEdit) AnchorClient(ASpace int32) {
	SpinEdit_AnchorClient(s.instance, ASpace)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (s *TSpinEdit) Anchors() TAnchors {
	return SpinEdit_GetAnchors(s.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (s *TSpinEdit) SetAnchors(value TAnchors) {
	SpinEdit_SetAnchors(s.instance, value)
}

// CN: 获取自动选择。
// EN: .
func (s *TSpinEdit) AutoSelect() bool {
	return SpinEdit_GetAutoSelect(s.instance)
}

// CN: 设置自动选择。
// EN: .
func (s *TSpinEdit) SetAutoSelect(value bool) {
	SpinEdit_SetAutoSelect(s.instance, value)
}

// CN: 获取自动调整大小。
// EN: .
func (s *TSpinEdit) AutoSize() bool {
	return SpinEdit_GetAutoSize(s.instance)
}

// CN: 设置自动调整大小。
// EN: .
func (s *TSpinEdit) SetAutoSize(value bool) {
	SpinEdit_SetAutoSize(s.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (s *TSpinEdit) Color() TColor {
	return SpinEdit_GetColor(s.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (s *TSpinEdit) SetColor(value TColor) {
	SpinEdit_SetColor(s.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (s *TSpinEdit) Constraints() *TSizeConstraints {
	return AsSizeConstraints(SpinEdit_GetConstraints(s.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (s *TSpinEdit) SetConstraints(value *TSizeConstraints) {
	SpinEdit_SetConstraints(s.instance, CheckPtr(value))
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TSpinEdit) Enabled() bool {
	return SpinEdit_GetEnabled(s.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TSpinEdit) SetEnabled(value bool) {
	SpinEdit_SetEnabled(s.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (s *TSpinEdit) Font() *TFont {
	return AsFont(SpinEdit_GetFont(s.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (s *TSpinEdit) SetFont(value *TFont) {
	SpinEdit_SetFont(s.instance, CheckPtr(value))
}

func (s *TSpinEdit) Increment() int32 {
	return SpinEdit_GetIncrement(s.instance)
}

func (s *TSpinEdit) SetIncrement(value int32) {
	SpinEdit_SetIncrement(s.instance, value)
}

// CN: 获取最大长度。
// EN: .
func (s *TSpinEdit) MaxLength() int32 {
	return SpinEdit_GetMaxLength(s.instance)
}

// CN: 设置最大长度。
// EN: .
func (s *TSpinEdit) SetMaxLength(value int32) {
	SpinEdit_SetMaxLength(s.instance, value)
}

func (s *TSpinEdit) MaxValue() int32 {
	return SpinEdit_GetMaxValue(s.instance)
}

func (s *TSpinEdit) SetMaxValue(value int32) {
	SpinEdit_SetMaxValue(s.instance, value)
}

func (s *TSpinEdit) MinValue() int32 {
	return SpinEdit_GetMinValue(s.instance)
}

func (s *TSpinEdit) SetMinValue(value int32) {
	SpinEdit_SetMinValue(s.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (s *TSpinEdit) ParentColor() bool {
	return SpinEdit_GetParentColor(s.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (s *TSpinEdit) SetParentColor(value bool) {
	SpinEdit_SetParentColor(s.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (s *TSpinEdit) ParentFont() bool {
	return SpinEdit_GetParentFont(s.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (s *TSpinEdit) SetParentFont(value bool) {
	SpinEdit_SetParentFont(s.instance, value)
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (s *TSpinEdit) ParentShowHint() bool {
	return SpinEdit_GetParentShowHint(s.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (s *TSpinEdit) SetParentShowHint(value bool) {
	SpinEdit_SetParentShowHint(s.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (s *TSpinEdit) PopupMenu() *TPopupMenu {
	return AsPopupMenu(SpinEdit_GetPopupMenu(s.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (s *TSpinEdit) SetPopupMenu(value IComponent) {
	SpinEdit_SetPopupMenu(s.instance, CheckPtr(value))
}

// CN: 获取只读。
// EN: .
func (s *TSpinEdit) ReadOnly() bool {
	return SpinEdit_GetReadOnly(s.instance)
}

// CN: 设置只读。
// EN: .
func (s *TSpinEdit) SetReadOnly(value bool) {
	SpinEdit_SetReadOnly(s.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (s *TSpinEdit) ShowHint() bool {
	return SpinEdit_GetShowHint(s.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (s *TSpinEdit) SetShowHint(value bool) {
	SpinEdit_SetShowHint(s.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (s *TSpinEdit) TabOrder() TTabOrder {
	return SpinEdit_GetTabOrder(s.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (s *TSpinEdit) SetTabOrder(value TTabOrder) {
	SpinEdit_SetTabOrder(s.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (s *TSpinEdit) TabStop() bool {
	return SpinEdit_GetTabStop(s.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (s *TSpinEdit) SetTabStop(value bool) {
	SpinEdit_SetTabStop(s.instance, value)
}

func (s *TSpinEdit) Value() int32 {
	return SpinEdit_GetValue(s.instance)
}

func (s *TSpinEdit) SetValue(value int32) {
	SpinEdit_SetValue(s.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TSpinEdit) Visible() bool {
	return SpinEdit_GetVisible(s.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TSpinEdit) SetVisible(value bool) {
	SpinEdit_SetVisible(s.instance, value)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (s *TSpinEdit) SetOnChange(fn *TNotifyEvent) {
	SpinEdit_SetOnChange(s.instance, fn)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (s *TSpinEdit) SetOnClick(fn *TNotifyEvent) {
	SpinEdit_SetOnClick(s.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (s *TSpinEdit) SetOnEnter(fn *TNotifyEvent) {
	SpinEdit_SetOnEnter(s.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (s *TSpinEdit) SetOnExit(fn *TNotifyEvent) {
	SpinEdit_SetOnExit(s.instance, fn)
}

// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (s *TSpinEdit) SetOnKeyDown(fn *TKeyEvent) {
	SpinEdit_SetOnKeyDown(s.instance, fn)
}

func (s *TSpinEdit) SetOnKeyPress(fn *TKeyPressEvent) {
	SpinEdit_SetOnKeyPress(s.instance, fn)
}

// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (s *TSpinEdit) SetOnKeyUp(fn *TKeyEvent) {
	SpinEdit_SetOnKeyUp(s.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (s *TSpinEdit) SetOnMouseDown(fn *TMouseEvent) {
	SpinEdit_SetOnMouseDown(s.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (s *TSpinEdit) SetOnMouseMove(fn *TMouseMoveEvent) {
	SpinEdit_SetOnMouseMove(s.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (s *TSpinEdit) SetOnMouseUp(fn *TMouseEvent) {
	SpinEdit_SetOnMouseUp(s.instance, fn)
}

// CN: 获取文字对齐。
// EN: Get Text alignment.
func (s *TSpinEdit) Alignment() TAlignment {
	return SpinEdit_GetAlignment(s.instance)
}

// CN: 设置文字对齐。
// EN: Set Text alignment.
func (s *TSpinEdit) SetAlignment(value TAlignment) {
	SpinEdit_SetAlignment(s.instance, value)
}

// CN: 获取能否撤销。
// EN: .
func (s *TSpinEdit) CanUndo() bool {
	return SpinEdit_GetCanUndo(s.instance)
}

// CN: 获取修改。
// EN: Get modified.
func (s *TSpinEdit) Modified() bool {
	return SpinEdit_GetModified(s.instance)
}

// CN: 设置修改。
// EN: Set modified.
func (s *TSpinEdit) SetModified(value bool) {
	SpinEdit_SetModified(s.instance, value)
}

// CN: 获取选择的长度。
// EN: .
func (s *TSpinEdit) SelLength() int32 {
	return SpinEdit_GetSelLength(s.instance)
}

// CN: 设置选择的长度。
// EN: .
func (s *TSpinEdit) SetSelLength(value int32) {
	SpinEdit_SetSelLength(s.instance, value)
}

// CN: 获取选择的启始位置。
// EN: .
func (s *TSpinEdit) SelStart() int32 {
	return SpinEdit_GetSelStart(s.instance)
}

// CN: 设置选择的启始位置。
// EN: .
func (s *TSpinEdit) SetSelStart(value int32) {
	SpinEdit_SetSelStart(s.instance, value)
}

// CN: 获取选择的文本。
// EN: .
func (s *TSpinEdit) SelText() string {
	return SpinEdit_GetSelText(s.instance)
}

// CN: 设置选择的文本。
// EN: .
func (s *TSpinEdit) SetSelText(value string) {
	SpinEdit_SetSelText(s.instance, value)
}

// CN: 获取文本。
// EN: .
func (s *TSpinEdit) Text() string {
	strLen := s.GetTextLen()
	if strLen != 0 {
		var buffStr string
		s.GetTextBuf(&buffStr, strLen+1)
		return buffStr
	}
	return ""
}

// CN: 设置文本。
// EN: .
func (s *TSpinEdit) SetText(value string) {
	SpinEdit_SetText(s.instance, value)
}

// CN: 获取提示文本。
// EN: .
func (s *TSpinEdit) TextHint() string {
	return SpinEdit_GetTextHint(s.instance)
}

// CN: 设置提示文本。
// EN: .
func (s *TSpinEdit) SetTextHint(value string) {
	SpinEdit_SetTextHint(s.instance, value)
}

// CN: 获取依靠客户端总数。
// EN: .
func (s *TSpinEdit) DockClientCount() int32 {
	return SpinEdit_GetDockClientCount(s.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (s *TSpinEdit) DockSite() bool {
	return SpinEdit_GetDockSite(s.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (s *TSpinEdit) SetDockSite(value bool) {
	SpinEdit_SetDockSite(s.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (s *TSpinEdit) DoubleBuffered() bool {
	return SpinEdit_GetDoubleBuffered(s.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (s *TSpinEdit) SetDoubleBuffered(value bool) {
	SpinEdit_SetDoubleBuffered(s.instance, value)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (s *TSpinEdit) MouseInClient() bool {
	return SpinEdit_GetMouseInClient(s.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (s *TSpinEdit) VisibleDockClientCount() int32 {
	return SpinEdit_GetVisibleDockClientCount(s.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (s *TSpinEdit) Brush() *TBrush {
	return AsBrush(SpinEdit_GetBrush(s.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (s *TSpinEdit) ControlCount() int32 {
	return SpinEdit_GetControlCount(s.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (s *TSpinEdit) Handle() HWND {
	return SpinEdit_GetHandle(s.instance)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (s *TSpinEdit) ParentDoubleBuffered() bool {
	return SpinEdit_GetParentDoubleBuffered(s.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (s *TSpinEdit) SetParentDoubleBuffered(value bool) {
	SpinEdit_SetParentDoubleBuffered(s.instance, value)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (s *TSpinEdit) ParentWindow() HWND {
	return SpinEdit_GetParentWindow(s.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (s *TSpinEdit) SetParentWindow(value HWND) {
	SpinEdit_SetParentWindow(s.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (s *TSpinEdit) UseDockManager() bool {
	return SpinEdit_GetUseDockManager(s.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (s *TSpinEdit) SetUseDockManager(value bool) {
	SpinEdit_SetUseDockManager(s.instance, value)
}

func (s *TSpinEdit) Action() *TAction {
	return AsAction(SpinEdit_GetAction(s.instance))
}

func (s *TSpinEdit) SetAction(value IComponent) {
	SpinEdit_SetAction(s.instance, CheckPtr(value))
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (s *TSpinEdit) Align() TAlign {
	return SpinEdit_GetAlign(s.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (s *TSpinEdit) SetAlign(value TAlign) {
	SpinEdit_SetAlign(s.instance, value)
}

func (s *TSpinEdit) BiDiMode() TBiDiMode {
	return SpinEdit_GetBiDiMode(s.instance)
}

func (s *TSpinEdit) SetBiDiMode(value TBiDiMode) {
	SpinEdit_SetBiDiMode(s.instance, value)
}

func (s *TSpinEdit) BoundsRect() TRect {
	return SpinEdit_GetBoundsRect(s.instance)
}

func (s *TSpinEdit) SetBoundsRect(value TRect) {
	SpinEdit_SetBoundsRect(s.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (s *TSpinEdit) ClientHeight() int32 {
	return SpinEdit_GetClientHeight(s.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (s *TSpinEdit) SetClientHeight(value int32) {
	SpinEdit_SetClientHeight(s.instance, value)
}

func (s *TSpinEdit) ClientOrigin() TPoint {
	return SpinEdit_GetClientOrigin(s.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (s *TSpinEdit) ClientRect() TRect {
	return SpinEdit_GetClientRect(s.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (s *TSpinEdit) ClientWidth() int32 {
	return SpinEdit_GetClientWidth(s.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (s *TSpinEdit) SetClientWidth(value int32) {
	SpinEdit_SetClientWidth(s.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (s *TSpinEdit) ControlState() TControlState {
	return SpinEdit_GetControlState(s.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (s *TSpinEdit) SetControlState(value TControlState) {
	SpinEdit_SetControlState(s.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (s *TSpinEdit) ControlStyle() TControlStyle {
	return SpinEdit_GetControlStyle(s.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (s *TSpinEdit) SetControlStyle(value TControlStyle) {
	SpinEdit_SetControlStyle(s.instance, value)
}

func (s *TSpinEdit) Floating() bool {
	return SpinEdit_GetFloating(s.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TSpinEdit) Parent() *TWinControl {
	return AsWinControl(SpinEdit_GetParent(s.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TSpinEdit) SetParent(value IWinControl) {
	SpinEdit_SetParent(s.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (s *TSpinEdit) Left() int32 {
	return SpinEdit_GetLeft(s.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (s *TSpinEdit) SetLeft(value int32) {
	SpinEdit_SetLeft(s.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (s *TSpinEdit) Top() int32 {
	return SpinEdit_GetTop(s.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (s *TSpinEdit) SetTop(value int32) {
	SpinEdit_SetTop(s.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (s *TSpinEdit) Width() int32 {
	return SpinEdit_GetWidth(s.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (s *TSpinEdit) SetWidth(value int32) {
	SpinEdit_SetWidth(s.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (s *TSpinEdit) Height() int32 {
	return SpinEdit_GetHeight(s.instance)
}

// CN: 设置高度。
// EN: Set height.
func (s *TSpinEdit) SetHeight(value int32) {
	SpinEdit_SetHeight(s.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TSpinEdit) Cursor() TCursor {
	return SpinEdit_GetCursor(s.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TSpinEdit) SetCursor(value TCursor) {
	SpinEdit_SetCursor(s.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (s *TSpinEdit) Hint() string {
	return SpinEdit_GetHint(s.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (s *TSpinEdit) SetHint(value string) {
	SpinEdit_SetHint(s.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TSpinEdit) ComponentCount() int32 {
	return SpinEdit_GetComponentCount(s.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (s *TSpinEdit) ComponentIndex() int32 {
	return SpinEdit_GetComponentIndex(s.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (s *TSpinEdit) SetComponentIndex(value int32) {
	SpinEdit_SetComponentIndex(s.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TSpinEdit) Owner() *TComponent {
	return AsComponent(SpinEdit_GetOwner(s.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (s *TSpinEdit) Name() string {
	return SpinEdit_GetName(s.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (s *TSpinEdit) SetName(value string) {
	SpinEdit_SetName(s.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TSpinEdit) Tag() int {
	return SpinEdit_GetTag(s.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TSpinEdit) SetTag(value int) {
	SpinEdit_SetTag(s.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (s *TSpinEdit) AnchorSideLeft() *TAnchorSide {
	return AsAnchorSide(SpinEdit_GetAnchorSideLeft(s.instance))
}

// CN: 设置左边锚点。
// EN: .
func (s *TSpinEdit) SetAnchorSideLeft(value *TAnchorSide) {
	SpinEdit_SetAnchorSideLeft(s.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (s *TSpinEdit) AnchorSideTop() *TAnchorSide {
	return AsAnchorSide(SpinEdit_GetAnchorSideTop(s.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (s *TSpinEdit) SetAnchorSideTop(value *TAnchorSide) {
	SpinEdit_SetAnchorSideTop(s.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (s *TSpinEdit) AnchorSideRight() *TAnchorSide {
	return AsAnchorSide(SpinEdit_GetAnchorSideRight(s.instance))
}

// CN: 设置右边锚点。
// EN: .
func (s *TSpinEdit) SetAnchorSideRight(value *TAnchorSide) {
	SpinEdit_SetAnchorSideRight(s.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (s *TSpinEdit) AnchorSideBottom() *TAnchorSide {
	return AsAnchorSide(SpinEdit_GetAnchorSideBottom(s.instance))
}

// CN: 设置底边锚点。
// EN: .
func (s *TSpinEdit) SetAnchorSideBottom(value *TAnchorSide) {
	SpinEdit_SetAnchorSideBottom(s.instance, CheckPtr(value))
}

func (s *TSpinEdit) ChildSizing() *TControlChildSizing {
	return AsControlChildSizing(SpinEdit_GetChildSizing(s.instance))
}

func (s *TSpinEdit) SetChildSizing(value *TControlChildSizing) {
	SpinEdit_SetChildSizing(s.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (s *TSpinEdit) BorderSpacing() *TControlBorderSpacing {
	return AsControlBorderSpacing(SpinEdit_GetBorderSpacing(s.instance))
}

// CN: 设置边框间距。
// EN: .
func (s *TSpinEdit) SetBorderSpacing(value *TControlBorderSpacing) {
	SpinEdit_SetBorderSpacing(s.instance, CheckPtr(value))
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (s *TSpinEdit) DockClients(Index int32) *TControl {
	return AsControl(SpinEdit_GetDockClients(s.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (s *TSpinEdit) Controls(Index int32) *TControl {
	return AsControl(SpinEdit_GetControls(s.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TSpinEdit) Components(AIndex int32) *TComponent {
	return AsComponent(SpinEdit_GetComponents(s.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (s *TSpinEdit) AnchorSide(AKind TAnchorKind) *TAnchorSide {
	return AsAnchorSide(SpinEdit_GetAnchorSide(s.instance, AKind))
}
