
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

type TProgressBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewProgressBar(owner IComponent) *TProgressBar {
    p := new(TProgressBar)
    p.instance = ProgressBar_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsProgressBar(obj interface{}) *TProgressBar {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TProgressBar{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsProgressBar.
func ProgressBarFromInst(inst uintptr) *TProgressBar {
    return AsProgressBar(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsProgressBar.
func ProgressBarFromObj(obj IObject) *TProgressBar {
    return AsProgressBar(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsProgressBar.
func ProgressBarFromUnsafePointer(ptr unsafe.Pointer) *TProgressBar {
    return AsProgressBar(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (p *TProgressBar) Free() {
    if p.instance != 0 {
        ProgressBar_Free(p.instance)
        p.instance, p.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TProgressBar) Instance() uintptr {
    return p.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TProgressBar) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TProgressBar) IsValid() bool {
    return p.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (p *TProgressBar) Is() TIs {
    return TIs(p.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (p *TProgressBar) As() TAs {
//    return TAs(p.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TProgressBarClass() TClass {
    return ProgressBar_StaticClassType()
}

func (p *TProgressBar) StepIt() {
    ProgressBar_StepIt(p.instance)
}

func (p *TProgressBar) StepBy(Delta int32) {
    ProgressBar_StepBy(p.instance, Delta)
}

// CN: 是否可以获得焦点。
// EN: .
func (p *TProgressBar) CanFocus() bool {
    return ProgressBar_CanFocus(p.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (p *TProgressBar) ContainsControl(Control IControl) bool {
    return ProgressBar_ContainsControl(p.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (p *TProgressBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ProgressBar_ControlAtPos(p.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (p *TProgressBar) DisableAlign() {
    ProgressBar_DisableAlign(p.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (p *TProgressBar) EnableAlign() {
    ProgressBar_EnableAlign(p.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (p *TProgressBar) FindChildControl(ControlName string) *TControl {
    return AsControl(ProgressBar_FindChildControl(p.instance, ControlName))
}

func (p *TProgressBar) FlipChildren(AllLevels bool) {
    ProgressBar_FlipChildren(p.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (p *TProgressBar) Focused() bool {
    return ProgressBar_Focused(p.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (p *TProgressBar) HandleAllocated() bool {
    return ProgressBar_HandleAllocated(p.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (p *TProgressBar) InsertControl(AControl IControl) {
    ProgressBar_InsertControl(p.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (p *TProgressBar) Invalidate() {
    ProgressBar_Invalidate(p.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (p *TProgressBar) RemoveControl(AControl IControl) {
    ProgressBar_RemoveControl(p.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (p *TProgressBar) Realign() {
    ProgressBar_Realign(p.instance)
}

// CN: 重绘。
// EN: Repaint.
func (p *TProgressBar) Repaint() {
    ProgressBar_Repaint(p.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (p *TProgressBar) ScaleBy(M int32, D int32) {
    ProgressBar_ScaleBy(p.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (p *TProgressBar) ScrollBy(DeltaX int32, DeltaY int32) {
    ProgressBar_ScrollBy(p.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (p *TProgressBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ProgressBar_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (p *TProgressBar) SetFocus() {
    ProgressBar_SetFocus(p.instance)
}

// CN: 控件更新。
// EN: Update.
func (p *TProgressBar) Update() {
    ProgressBar_Update(p.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (p *TProgressBar) BringToFront() {
    ProgressBar_BringToFront(p.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (p *TProgressBar) ClientToScreen(Point TPoint) TPoint {
    return ProgressBar_ClientToScreen(p.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (p *TProgressBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ProgressBar_ClientToParent(p.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (p *TProgressBar) Dragging() bool {
    return ProgressBar_Dragging(p.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (p *TProgressBar) HasParent() bool {
    return ProgressBar_HasParent(p.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (p *TProgressBar) Hide() {
    ProgressBar_Hide(p.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (p *TProgressBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ProgressBar_Perform(p.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (p *TProgressBar) Refresh() {
    ProgressBar_Refresh(p.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (p *TProgressBar) ScreenToClient(Point TPoint) TPoint {
    return ProgressBar_ScreenToClient(p.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (p *TProgressBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ProgressBar_ParentToClient(p.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (p *TProgressBar) SendToBack() {
    ProgressBar_SendToBack(p.instance)
}

// CN: 显示控件。
// EN: Show control.
func (p *TProgressBar) Show() {
    ProgressBar_Show(p.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (p *TProgressBar) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ProgressBar_GetTextBuf(p.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (p *TProgressBar) GetTextLen() int32 {
    return ProgressBar_GetTextLen(p.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (p *TProgressBar) SetTextBuf(Buffer string) {
    ProgressBar_SetTextBuf(p.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (p *TProgressBar) FindComponent(AName string) *TComponent {
    return AsComponent(ProgressBar_FindComponent(p.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TProgressBar) GetNamePath() string {
    return ProgressBar_GetNamePath(p.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TProgressBar) Assign(Source IObject) {
    ProgressBar_Assign(p.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TProgressBar) ClassType() TClass {
    return ProgressBar_ClassType(p.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TProgressBar) ClassName() string {
    return ProgressBar_ClassName(p.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TProgressBar) InstanceSize() int32 {
    return ProgressBar_InstanceSize(p.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TProgressBar) InheritsFrom(AClass TClass) bool {
    return ProgressBar_InheritsFrom(p.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TProgressBar) Equals(Obj IObject) bool {
    return ProgressBar_Equals(p.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TProgressBar) GetHashCode() int32 {
    return ProgressBar_GetHashCode(p.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (p *TProgressBar) ToString() string {
    return ProgressBar_ToString(p.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (p *TProgressBar) Align() TAlign {
    return ProgressBar_GetAlign(p.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (p *TProgressBar) SetAlign(value TAlign) {
    ProgressBar_SetAlign(p.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (p *TProgressBar) Anchors() TAnchors {
    return ProgressBar_GetAnchors(p.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (p *TProgressBar) SetAnchors(value TAnchors) {
    ProgressBar_SetAnchors(p.instance, value)
}

// CN: 获取边框的宽度。
// EN: .
func (p *TProgressBar) BorderWidth() int32 {
    return ProgressBar_GetBorderWidth(p.instance)
}

// CN: 设置边框的宽度。
// EN: .
func (p *TProgressBar) SetBorderWidth(value int32) {
    ProgressBar_SetBorderWidth(p.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (p *TProgressBar) DoubleBuffered() bool {
    return ProgressBar_GetDoubleBuffered(p.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (p *TProgressBar) SetDoubleBuffered(value bool) {
    ProgressBar_SetDoubleBuffered(p.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (p *TProgressBar) DragCursor() TCursor {
    return ProgressBar_GetDragCursor(p.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (p *TProgressBar) SetDragCursor(value TCursor) {
    ProgressBar_SetDragCursor(p.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (p *TProgressBar) DragKind() TDragKind {
    return ProgressBar_GetDragKind(p.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (p *TProgressBar) SetDragKind(value TDragKind) {
    ProgressBar_SetDragKind(p.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (p *TProgressBar) DragMode() TDragMode {
    return ProgressBar_GetDragMode(p.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (p *TProgressBar) SetDragMode(value TDragMode) {
    ProgressBar_SetDragMode(p.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (p *TProgressBar) Enabled() bool {
    return ProgressBar_GetEnabled(p.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (p *TProgressBar) SetEnabled(value bool) {
    ProgressBar_SetEnabled(p.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (p *TProgressBar) Hint() string {
    return ProgressBar_GetHint(p.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (p *TProgressBar) SetHint(value string) {
    ProgressBar_SetHint(p.instance, value)
}

func (p *TProgressBar) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ProgressBar_GetConstraints(p.instance))
}

func (p *TProgressBar) SetConstraints(value *TSizeConstraints) {
    ProgressBar_SetConstraints(p.instance, CheckPtr(value))
}

func (p *TProgressBar) Min() int32 {
    return ProgressBar_GetMin(p.instance)
}

func (p *TProgressBar) SetMin(value int32) {
    ProgressBar_SetMin(p.instance, value)
}

func (p *TProgressBar) Max() int32 {
    return ProgressBar_GetMax(p.instance)
}

func (p *TProgressBar) SetMax(value int32) {
    ProgressBar_SetMax(p.instance, value)
}

func (p *TProgressBar) Orientation() TProgressBarOrientation {
    return ProgressBar_GetOrientation(p.instance)
}

func (p *TProgressBar) SetOrientation(value TProgressBarOrientation) {
    ProgressBar_SetOrientation(p.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (p *TProgressBar) ParentDoubleBuffered() bool {
    return ProgressBar_GetParentDoubleBuffered(p.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (p *TProgressBar) SetParentDoubleBuffered(value bool) {
    ProgressBar_SetParentDoubleBuffered(p.instance, value)
}

func (p *TProgressBar) ParentShowHint() bool {
    return ProgressBar_GetParentShowHint(p.instance)
}

func (p *TProgressBar) SetParentShowHint(value bool) {
    ProgressBar_SetParentShowHint(p.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (p *TProgressBar) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ProgressBar_GetPopupMenu(p.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (p *TProgressBar) SetPopupMenu(value IComponent) {
    ProgressBar_SetPopupMenu(p.instance, CheckPtr(value))
}

func (p *TProgressBar) Position() int32 {
    return ProgressBar_GetPosition(p.instance)
}

func (p *TProgressBar) SetPosition(value int32) {
    ProgressBar_SetPosition(p.instance, value)
}

func (p *TProgressBar) Smooth() bool {
    return ProgressBar_GetSmooth(p.instance)
}

func (p *TProgressBar) SetSmooth(value bool) {
    ProgressBar_SetSmooth(p.instance, value)
}

func (p *TProgressBar) Style() TProgressBarStyle {
    return ProgressBar_GetStyle(p.instance)
}

func (p *TProgressBar) SetStyle(value TProgressBarStyle) {
    ProgressBar_SetStyle(p.instance, value)
}

func (p *TProgressBar) Step() int32 {
    return ProgressBar_GetStep(p.instance)
}

func (p *TProgressBar) SetStep(value int32) {
    ProgressBar_SetStep(p.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (p *TProgressBar) ShowHint() bool {
    return ProgressBar_GetShowHint(p.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (p *TProgressBar) SetShowHint(value bool) {
    ProgressBar_SetShowHint(p.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (p *TProgressBar) TabOrder() TTabOrder {
    return ProgressBar_GetTabOrder(p.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (p *TProgressBar) SetTabOrder(value TTabOrder) {
    ProgressBar_SetTabOrder(p.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (p *TProgressBar) TabStop() bool {
    return ProgressBar_GetTabStop(p.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (p *TProgressBar) SetTabStop(value bool) {
    ProgressBar_SetTabStop(p.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (p *TProgressBar) Visible() bool {
    return ProgressBar_GetVisible(p.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (p *TProgressBar) SetVisible(value bool) {
    ProgressBar_SetVisible(p.instance, value)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (p *TProgressBar) SetOnContextPopup(fn TContextPopupEvent) {
    ProgressBar_SetOnContextPopup(p.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (p *TProgressBar) SetOnDragDrop(fn TDragDropEvent) {
    ProgressBar_SetOnDragDrop(p.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (p *TProgressBar) SetOnDragOver(fn TDragOverEvent) {
    ProgressBar_SetOnDragOver(p.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (p *TProgressBar) SetOnEndDrag(fn TEndDragEvent) {
    ProgressBar_SetOnEndDrag(p.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (p *TProgressBar) SetOnEnter(fn TNotifyEvent) {
    ProgressBar_SetOnEnter(p.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (p *TProgressBar) SetOnExit(fn TNotifyEvent) {
    ProgressBar_SetOnExit(p.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (p *TProgressBar) SetOnMouseDown(fn TMouseEvent) {
    ProgressBar_SetOnMouseDown(p.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (p *TProgressBar) SetOnMouseEnter(fn TNotifyEvent) {
    ProgressBar_SetOnMouseEnter(p.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (p *TProgressBar) SetOnMouseLeave(fn TNotifyEvent) {
    ProgressBar_SetOnMouseLeave(p.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (p *TProgressBar) SetOnMouseMove(fn TMouseMoveEvent) {
    ProgressBar_SetOnMouseMove(p.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (p *TProgressBar) SetOnMouseUp(fn TMouseEvent) {
    ProgressBar_SetOnMouseUp(p.instance, fn)
}

// CN: 获取依靠客户端总数。
// EN: .
func (p *TProgressBar) DockClientCount() int32 {
    return ProgressBar_GetDockClientCount(p.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (p *TProgressBar) DockSite() bool {
    return ProgressBar_GetDockSite(p.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (p *TProgressBar) SetDockSite(value bool) {
    ProgressBar_SetDockSite(p.instance, value)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (p *TProgressBar) MouseInClient() bool {
    return ProgressBar_GetMouseInClient(p.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (p *TProgressBar) VisibleDockClientCount() int32 {
    return ProgressBar_GetVisibleDockClientCount(p.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (p *TProgressBar) Brush() *TBrush {
    return AsBrush(ProgressBar_GetBrush(p.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (p *TProgressBar) ControlCount() int32 {
    return ProgressBar_GetControlCount(p.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (p *TProgressBar) Handle() HWND {
    return ProgressBar_GetHandle(p.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (p *TProgressBar) ParentWindow() HWND {
    return ProgressBar_GetParentWindow(p.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (p *TProgressBar) SetParentWindow(value HWND) {
    ProgressBar_SetParentWindow(p.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (p *TProgressBar) UseDockManager() bool {
    return ProgressBar_GetUseDockManager(p.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (p *TProgressBar) SetUseDockManager(value bool) {
    ProgressBar_SetUseDockManager(p.instance, value)
}

func (p *TProgressBar) Action() *TAction {
    return AsAction(ProgressBar_GetAction(p.instance))
}

func (p *TProgressBar) SetAction(value IComponent) {
    ProgressBar_SetAction(p.instance, CheckPtr(value))
}

func (p *TProgressBar) BiDiMode() TBiDiMode {
    return ProgressBar_GetBiDiMode(p.instance)
}

func (p *TProgressBar) SetBiDiMode(value TBiDiMode) {
    ProgressBar_SetBiDiMode(p.instance, value)
}

func (p *TProgressBar) BoundsRect() TRect {
    return ProgressBar_GetBoundsRect(p.instance)
}

func (p *TProgressBar) SetBoundsRect(value TRect) {
    ProgressBar_SetBoundsRect(p.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (p *TProgressBar) ClientHeight() int32 {
    return ProgressBar_GetClientHeight(p.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (p *TProgressBar) SetClientHeight(value int32) {
    ProgressBar_SetClientHeight(p.instance, value)
}

func (p *TProgressBar) ClientOrigin() TPoint {
    return ProgressBar_GetClientOrigin(p.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (p *TProgressBar) ClientRect() TRect {
    return ProgressBar_GetClientRect(p.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (p *TProgressBar) ClientWidth() int32 {
    return ProgressBar_GetClientWidth(p.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (p *TProgressBar) SetClientWidth(value int32) {
    ProgressBar_SetClientWidth(p.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (p *TProgressBar) ControlState() TControlState {
    return ProgressBar_GetControlState(p.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (p *TProgressBar) SetControlState(value TControlState) {
    ProgressBar_SetControlState(p.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (p *TProgressBar) ControlStyle() TControlStyle {
    return ProgressBar_GetControlStyle(p.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (p *TProgressBar) SetControlStyle(value TControlStyle) {
    ProgressBar_SetControlStyle(p.instance, value)
}

func (p *TProgressBar) Floating() bool {
    return ProgressBar_GetFloating(p.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (p *TProgressBar) Parent() *TWinControl {
    return AsWinControl(ProgressBar_GetParent(p.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (p *TProgressBar) SetParent(value IWinControl) {
    ProgressBar_SetParent(p.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (p *TProgressBar) Left() int32 {
    return ProgressBar_GetLeft(p.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (p *TProgressBar) SetLeft(value int32) {
    ProgressBar_SetLeft(p.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (p *TProgressBar) Top() int32 {
    return ProgressBar_GetTop(p.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (p *TProgressBar) SetTop(value int32) {
    ProgressBar_SetTop(p.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (p *TProgressBar) Width() int32 {
    return ProgressBar_GetWidth(p.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (p *TProgressBar) SetWidth(value int32) {
    ProgressBar_SetWidth(p.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (p *TProgressBar) Height() int32 {
    return ProgressBar_GetHeight(p.instance)
}

// CN: 设置高度。
// EN: Set height.
func (p *TProgressBar) SetHeight(value int32) {
    ProgressBar_SetHeight(p.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (p *TProgressBar) Cursor() TCursor {
    return ProgressBar_GetCursor(p.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (p *TProgressBar) SetCursor(value TCursor) {
    ProgressBar_SetCursor(p.instance, value)
}

// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (p *TProgressBar) Margins() *TMargins {
    return AsMargins(ProgressBar_GetMargins(p.instance))
}

// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (p *TProgressBar) SetMargins(value *TMargins) {
    ProgressBar_SetMargins(p.instance, CheckPtr(value))
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TProgressBar) ComponentCount() int32 {
    return ProgressBar_GetComponentCount(p.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (p *TProgressBar) ComponentIndex() int32 {
    return ProgressBar_GetComponentIndex(p.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (p *TProgressBar) SetComponentIndex(value int32) {
    ProgressBar_SetComponentIndex(p.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TProgressBar) Owner() *TComponent {
    return AsComponent(ProgressBar_GetOwner(p.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (p *TProgressBar) Name() string {
    return ProgressBar_GetName(p.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (p *TProgressBar) SetName(value string) {
    ProgressBar_SetName(p.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TProgressBar) Tag() int {
    return ProgressBar_GetTag(p.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TProgressBar) SetTag(value int) {
    ProgressBar_SetTag(p.instance, value)
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (p *TProgressBar) DockClients(Index int32) *TControl {
    return AsControl(ProgressBar_GetDockClients(p.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (p *TProgressBar) Controls(Index int32) *TControl {
    return AsControl(ProgressBar_GetControls(p.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TProgressBar) Components(AIndex int32) *TComponent {
    return AsComponent(ProgressBar_GetComponents(p.instance, AIndex))
}

