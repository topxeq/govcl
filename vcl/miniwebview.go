
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

type TMiniWebview struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMiniWebview(owner IComponent) *TMiniWebview {
    m := new(TMiniWebview)
    m.instance = MiniWebview_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(m, (*TMiniWebview).Free)
    return m
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsMiniWebview(obj interface{}) *TMiniWebview {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TMiniWebview{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsMiniWebview.
func MiniWebviewFromInst(inst uintptr) *TMiniWebview {
    return AsMiniWebview(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsMiniWebview.
func MiniWebviewFromObj(obj IObject) *TMiniWebview {
    return AsMiniWebview(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsMiniWebview.
func MiniWebviewFromUnsafePointer(ptr unsafe.Pointer) *TMiniWebview {
    return AsMiniWebview(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (m *TMiniWebview) Free() {
    if m.instance != 0 {
        MiniWebview_Free(m.instance)
        m.instance, m.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMiniWebview) Instance() uintptr {
    return m.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMiniWebview) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMiniWebview) IsValid() bool {
    return m.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (m *TMiniWebview) Is() TIs {
    return TIs(m.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (m *TMiniWebview) As() TAs {
//    return TAs(m.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMiniWebviewClass() TClass {
    return MiniWebview_StaticClassType()
}

func (m *TMiniWebview) Navigate(AURL string) {
    MiniWebview_Navigate(m.instance, AURL)
}

func (m *TMiniWebview) GoBack() {
    MiniWebview_GoBack(m.instance)
}

func (m *TMiniWebview) GoForward() {
    MiniWebview_GoForward(m.instance)
}

func (m *TMiniWebview) GoHome() {
    MiniWebview_GoHome(m.instance)
}

func (m *TMiniWebview) GoSearch() {
    MiniWebview_GoSearch(m.instance)
}

// CN: 刷新控件。
// EN: Refresh control.
func (m *TMiniWebview) Refresh() {
    MiniWebview_Refresh(m.instance)
}

func (m *TMiniWebview) Stop() {
    MiniWebview_Stop(m.instance)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (m *TMiniWebview) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    MiniWebview_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

func (m *TMiniWebview) ExecuteScript(AScriptText string, AScriptType string) {
    MiniWebview_ExecuteScript(m.instance, AScriptText , AScriptType)
}

func (m *TMiniWebview) ExecuteJS(AScriptText string) {
    MiniWebview_ExecuteJS(m.instance, AScriptText)
}

func (m *TMiniWebview) LoadHTML(AStr string) {
    MiniWebview_LoadHTML(m.instance, AStr)
}

// CN: 是否可以获得焦点。
// EN: .
func (m *TMiniWebview) CanFocus() bool {
    return MiniWebview_CanFocus(m.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (m *TMiniWebview) ContainsControl(Control IControl) bool {
    return MiniWebview_ContainsControl(m.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (m *TMiniWebview) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(MiniWebview_ControlAtPos(m.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (m *TMiniWebview) DisableAlign() {
    MiniWebview_DisableAlign(m.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (m *TMiniWebview) EnableAlign() {
    MiniWebview_EnableAlign(m.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (m *TMiniWebview) FindChildControl(ControlName string) *TControl {
    return AsControl(MiniWebview_FindChildControl(m.instance, ControlName))
}

func (m *TMiniWebview) FlipChildren(AllLevels bool) {
    MiniWebview_FlipChildren(m.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (m *TMiniWebview) Focused() bool {
    return MiniWebview_Focused(m.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (m *TMiniWebview) HandleAllocated() bool {
    return MiniWebview_HandleAllocated(m.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (m *TMiniWebview) InsertControl(AControl IControl) {
    MiniWebview_InsertControl(m.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (m *TMiniWebview) Invalidate() {
    MiniWebview_Invalidate(m.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (m *TMiniWebview) RemoveControl(AControl IControl) {
    MiniWebview_RemoveControl(m.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (m *TMiniWebview) Realign() {
    MiniWebview_Realign(m.instance)
}

// CN: 重绘。
// EN: Repaint.
func (m *TMiniWebview) Repaint() {
    MiniWebview_Repaint(m.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (m *TMiniWebview) ScaleBy(M int32, D int32) {
    MiniWebview_ScaleBy(m.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (m *TMiniWebview) ScrollBy(DeltaX int32, DeltaY int32) {
    MiniWebview_ScrollBy(m.instance, DeltaX , DeltaY)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (m *TMiniWebview) SetFocus() {
    MiniWebview_SetFocus(m.instance)
}

// CN: 控件更新。
// EN: Update.
func (m *TMiniWebview) Update() {
    MiniWebview_Update(m.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (m *TMiniWebview) BringToFront() {
    MiniWebview_BringToFront(m.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (m *TMiniWebview) ClientToScreen(Point TPoint) TPoint {
    return MiniWebview_ClientToScreen(m.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (m *TMiniWebview) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return MiniWebview_ClientToParent(m.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (m *TMiniWebview) Dragging() bool {
    return MiniWebview_Dragging(m.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (m *TMiniWebview) HasParent() bool {
    return MiniWebview_HasParent(m.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (m *TMiniWebview) Hide() {
    MiniWebview_Hide(m.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (m *TMiniWebview) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return MiniWebview_Perform(m.instance, Msg , WParam , LParam)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (m *TMiniWebview) ScreenToClient(Point TPoint) TPoint {
    return MiniWebview_ScreenToClient(m.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (m *TMiniWebview) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return MiniWebview_ParentToClient(m.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (m *TMiniWebview) SendToBack() {
    MiniWebview_SendToBack(m.instance)
}

// CN: 显示控件。
// EN: Show control.
func (m *TMiniWebview) Show() {
    MiniWebview_Show(m.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (m *TMiniWebview) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return MiniWebview_GetTextBuf(m.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (m *TMiniWebview) GetTextLen() int32 {
    return MiniWebview_GetTextLen(m.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (m *TMiniWebview) SetTextBuf(Buffer string) {
    MiniWebview_SetTextBuf(m.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (m *TMiniWebview) FindComponent(AName string) *TComponent {
    return AsComponent(MiniWebview_FindComponent(m.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (m *TMiniWebview) GetNamePath() string {
    return MiniWebview_GetNamePath(m.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (m *TMiniWebview) Assign(Source IObject) {
    MiniWebview_Assign(m.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMiniWebview) ClassType() TClass {
    return MiniWebview_ClassType(m.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMiniWebview) ClassName() string {
    return MiniWebview_ClassName(m.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMiniWebview) InstanceSize() int32 {
    return MiniWebview_InstanceSize(m.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMiniWebview) InheritsFrom(AClass TClass) bool {
    return MiniWebview_InheritsFrom(m.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMiniWebview) Equals(Obj IObject) bool {
    return MiniWebview_Equals(m.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMiniWebview) GetHashCode() int32 {
    return MiniWebview_GetHashCode(m.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (m *TMiniWebview) ToString() string {
    return MiniWebview_ToString(m.instance)
}

func (m *TMiniWebview) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    MiniWebview_AnchorToNeighbour(m.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (m *TMiniWebview) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    MiniWebview_AnchorParallel(m.instance, ASide , ASpace , CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (m *TMiniWebview) AnchorHorizontalCenterTo(ASibling IControl) {
    MiniWebview_AnchorHorizontalCenterTo(m.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (m *TMiniWebview) AnchorVerticalCenterTo(ASibling IControl) {
    MiniWebview_AnchorVerticalCenterTo(m.instance, CheckPtr(ASibling))
}

func (m *TMiniWebview) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    MiniWebview_AnchorAsAlign(m.instance, ATheAlign , ASpace)
}

func (m *TMiniWebview) AnchorClient(ASpace int32) {
    MiniWebview_AnchorClient(m.instance, ASpace)
}

func (m *TMiniWebview) ReadyState() TReadyState {
    return MiniWebview_GetReadyState(m.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (m *TMiniWebview) Align() TAlign {
    return MiniWebview_GetAlign(m.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (m *TMiniWebview) SetAlign(value TAlign) {
    MiniWebview_SetAlign(m.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (m *TMiniWebview) Anchors() TAnchors {
    return MiniWebview_GetAnchors(m.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (m *TMiniWebview) SetAnchors(value TAnchors) {
    MiniWebview_SetAnchors(m.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (m *TMiniWebview) Constraints() *TSizeConstraints {
    return AsSizeConstraints(MiniWebview_GetConstraints(m.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (m *TMiniWebview) SetConstraints(value *TSizeConstraints) {
    MiniWebview_SetConstraints(m.instance, CheckPtr(value))
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (m *TMiniWebview) Enabled() bool {
    return MiniWebview_GetEnabled(m.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (m *TMiniWebview) SetEnabled(value bool) {
    MiniWebview_SetEnabled(m.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (m *TMiniWebview) Visible() bool {
    return MiniWebview_GetVisible(m.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (m *TMiniWebview) SetVisible(value bool) {
    MiniWebview_SetVisible(m.instance, value)
}

func (m *TMiniWebview) SetOnTitleChange(fn TWebTitleChangeEvent) {
    MiniWebview_SetOnTitleChange(m.instance, fn)
}

func (m *TMiniWebview) SetOnJSExternal(fn TWebJSExternalEvent) {
    MiniWebview_SetOnJSExternal(m.instance, fn)
}

// CN: 获取依靠客户端总数。
// EN: .
func (m *TMiniWebview) DockClientCount() int32 {
    return MiniWebview_GetDockClientCount(m.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (m *TMiniWebview) DockSite() bool {
    return MiniWebview_GetDockSite(m.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (m *TMiniWebview) SetDockSite(value bool) {
    MiniWebview_SetDockSite(m.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (m *TMiniWebview) DoubleBuffered() bool {
    return MiniWebview_GetDoubleBuffered(m.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (m *TMiniWebview) SetDoubleBuffered(value bool) {
    MiniWebview_SetDoubleBuffered(m.instance, value)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (m *TMiniWebview) MouseInClient() bool {
    return MiniWebview_GetMouseInClient(m.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (m *TMiniWebview) VisibleDockClientCount() int32 {
    return MiniWebview_GetVisibleDockClientCount(m.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (m *TMiniWebview) Brush() *TBrush {
    return AsBrush(MiniWebview_GetBrush(m.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (m *TMiniWebview) ControlCount() int32 {
    return MiniWebview_GetControlCount(m.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (m *TMiniWebview) Handle() HWND {
    return MiniWebview_GetHandle(m.instance)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (m *TMiniWebview) ParentDoubleBuffered() bool {
    return MiniWebview_GetParentDoubleBuffered(m.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (m *TMiniWebview) SetParentDoubleBuffered(value bool) {
    MiniWebview_SetParentDoubleBuffered(m.instance, value)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (m *TMiniWebview) ParentWindow() HWND {
    return MiniWebview_GetParentWindow(m.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (m *TMiniWebview) SetParentWindow(value HWND) {
    MiniWebview_SetParentWindow(m.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (m *TMiniWebview) TabOrder() TTabOrder {
    return MiniWebview_GetTabOrder(m.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (m *TMiniWebview) SetTabOrder(value TTabOrder) {
    MiniWebview_SetTabOrder(m.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (m *TMiniWebview) TabStop() bool {
    return MiniWebview_GetTabStop(m.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (m *TMiniWebview) SetTabStop(value bool) {
    MiniWebview_SetTabStop(m.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (m *TMiniWebview) UseDockManager() bool {
    return MiniWebview_GetUseDockManager(m.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (m *TMiniWebview) SetUseDockManager(value bool) {
    MiniWebview_SetUseDockManager(m.instance, value)
}

func (m *TMiniWebview) Action() *TAction {
    return AsAction(MiniWebview_GetAction(m.instance))
}

func (m *TMiniWebview) SetAction(value IComponent) {
    MiniWebview_SetAction(m.instance, CheckPtr(value))
}

func (m *TMiniWebview) BiDiMode() TBiDiMode {
    return MiniWebview_GetBiDiMode(m.instance)
}

func (m *TMiniWebview) SetBiDiMode(value TBiDiMode) {
    MiniWebview_SetBiDiMode(m.instance, value)
}

func (m *TMiniWebview) BoundsRect() TRect {
    return MiniWebview_GetBoundsRect(m.instance)
}

func (m *TMiniWebview) SetBoundsRect(value TRect) {
    MiniWebview_SetBoundsRect(m.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (m *TMiniWebview) ClientHeight() int32 {
    return MiniWebview_GetClientHeight(m.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (m *TMiniWebview) SetClientHeight(value int32) {
    MiniWebview_SetClientHeight(m.instance, value)
}

func (m *TMiniWebview) ClientOrigin() TPoint {
    return MiniWebview_GetClientOrigin(m.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (m *TMiniWebview) ClientRect() TRect {
    return MiniWebview_GetClientRect(m.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (m *TMiniWebview) ClientWidth() int32 {
    return MiniWebview_GetClientWidth(m.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (m *TMiniWebview) SetClientWidth(value int32) {
    MiniWebview_SetClientWidth(m.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (m *TMiniWebview) ControlState() TControlState {
    return MiniWebview_GetControlState(m.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (m *TMiniWebview) SetControlState(value TControlState) {
    MiniWebview_SetControlState(m.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (m *TMiniWebview) ControlStyle() TControlStyle {
    return MiniWebview_GetControlStyle(m.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (m *TMiniWebview) SetControlStyle(value TControlStyle) {
    MiniWebview_SetControlStyle(m.instance, value)
}

func (m *TMiniWebview) Floating() bool {
    return MiniWebview_GetFloating(m.instance)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (m *TMiniWebview) ShowHint() bool {
    return MiniWebview_GetShowHint(m.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (m *TMiniWebview) SetShowHint(value bool) {
    MiniWebview_SetShowHint(m.instance, value)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (m *TMiniWebview) Parent() *TWinControl {
    return AsWinControl(MiniWebview_GetParent(m.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (m *TMiniWebview) SetParent(value IWinControl) {
    MiniWebview_SetParent(m.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (m *TMiniWebview) Left() int32 {
    return MiniWebview_GetLeft(m.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (m *TMiniWebview) SetLeft(value int32) {
    MiniWebview_SetLeft(m.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (m *TMiniWebview) Top() int32 {
    return MiniWebview_GetTop(m.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (m *TMiniWebview) SetTop(value int32) {
    MiniWebview_SetTop(m.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (m *TMiniWebview) Width() int32 {
    return MiniWebview_GetWidth(m.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (m *TMiniWebview) SetWidth(value int32) {
    MiniWebview_SetWidth(m.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (m *TMiniWebview) Height() int32 {
    return MiniWebview_GetHeight(m.instance)
}

// CN: 设置高度。
// EN: Set height.
func (m *TMiniWebview) SetHeight(value int32) {
    MiniWebview_SetHeight(m.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (m *TMiniWebview) Cursor() TCursor {
    return MiniWebview_GetCursor(m.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (m *TMiniWebview) SetCursor(value TCursor) {
    MiniWebview_SetCursor(m.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (m *TMiniWebview) Hint() string {
    return MiniWebview_GetHint(m.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (m *TMiniWebview) SetHint(value string) {
    MiniWebview_SetHint(m.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (m *TMiniWebview) ComponentCount() int32 {
    return MiniWebview_GetComponentCount(m.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (m *TMiniWebview) ComponentIndex() int32 {
    return MiniWebview_GetComponentIndex(m.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (m *TMiniWebview) SetComponentIndex(value int32) {
    MiniWebview_SetComponentIndex(m.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (m *TMiniWebview) Owner() *TComponent {
    return AsComponent(MiniWebview_GetOwner(m.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (m *TMiniWebview) Name() string {
    return MiniWebview_GetName(m.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (m *TMiniWebview) SetName(value string) {
    MiniWebview_SetName(m.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (m *TMiniWebview) Tag() int {
    return MiniWebview_GetTag(m.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (m *TMiniWebview) SetTag(value int) {
    MiniWebview_SetTag(m.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (m *TMiniWebview) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(MiniWebview_GetAnchorSideLeft(m.instance))
}

// CN: 设置左边锚点。
// EN: .
func (m *TMiniWebview) SetAnchorSideLeft(value *TAnchorSide) {
    MiniWebview_SetAnchorSideLeft(m.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (m *TMiniWebview) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(MiniWebview_GetAnchorSideTop(m.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (m *TMiniWebview) SetAnchorSideTop(value *TAnchorSide) {
    MiniWebview_SetAnchorSideTop(m.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (m *TMiniWebview) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(MiniWebview_GetAnchorSideRight(m.instance))
}

// CN: 设置右边锚点。
// EN: .
func (m *TMiniWebview) SetAnchorSideRight(value *TAnchorSide) {
    MiniWebview_SetAnchorSideRight(m.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (m *TMiniWebview) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(MiniWebview_GetAnchorSideBottom(m.instance))
}

// CN: 设置底边锚点。
// EN: .
func (m *TMiniWebview) SetAnchorSideBottom(value *TAnchorSide) {
    MiniWebview_SetAnchorSideBottom(m.instance, CheckPtr(value))
}

func (m *TMiniWebview) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(MiniWebview_GetChildSizing(m.instance))
}

func (m *TMiniWebview) SetChildSizing(value *TControlChildSizing) {
    MiniWebview_SetChildSizing(m.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (m *TMiniWebview) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(MiniWebview_GetBorderSpacing(m.instance))
}

// CN: 设置边框间距。
// EN: .
func (m *TMiniWebview) SetBorderSpacing(value *TControlBorderSpacing) {
    MiniWebview_SetBorderSpacing(m.instance, CheckPtr(value))
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (m *TMiniWebview) DockClients(Index int32) *TControl {
    return AsControl(MiniWebview_GetDockClients(m.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (m *TMiniWebview) Controls(Index int32) *TControl {
    return AsControl(MiniWebview_GetControls(m.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (m *TMiniWebview) Components(AIndex int32) *TComponent {
    return AsComponent(MiniWebview_GetComponents(m.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (m *TMiniWebview) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(MiniWebview_GetAnchorSide(m.instance, AKind))
}

