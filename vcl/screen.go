
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

type TScreen struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewScreen(owner IComponent) *TScreen {
    s := new(TScreen)
    s.instance = Screen_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(s, (*TScreen).Free)
    return s
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsScreen(obj interface{}) *TScreen {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TScreen{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsScreen.
func ScreenFromInst(inst uintptr) *TScreen {
    return AsScreen(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsScreen.
func ScreenFromObj(obj IObject) *TScreen {
    return AsScreen(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsScreen.
func ScreenFromUnsafePointer(ptr unsafe.Pointer) *TScreen {
    return AsScreen(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (s *TScreen) Free() {
    if s.instance != 0 {
        Screen_Free(s.instance)
        s.instance, s.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TScreen) Instance() uintptr {
    return s.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TScreen) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TScreen) IsValid() bool {
    return s.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (s *TScreen) Is() TIs {
    return TIs(s.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (s *TScreen) As() TAs {
//    return TAs(s.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TScreenClass() TClass {
    return Screen_StaticClassType()
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TScreen) FindComponent(AName string) *TComponent {
    return AsComponent(Screen_FindComponent(s.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TScreen) GetNamePath() string {
    return Screen_GetNamePath(s.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TScreen) HasParent() bool {
    return Screen_HasParent(s.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TScreen) Assign(Source IObject) {
    Screen_Assign(s.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TScreen) ClassType() TClass {
    return Screen_ClassType(s.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TScreen) ClassName() string {
    return Screen_ClassName(s.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TScreen) InstanceSize() int32 {
    return Screen_InstanceSize(s.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TScreen) InheritsFrom(AClass TClass) bool {
    return Screen_InheritsFrom(s.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TScreen) Equals(Obj IObject) bool {
    return Screen_Equals(s.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TScreen) GetHashCode() int32 {
    return Screen_GetHashCode(s.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (s *TScreen) ToString() string {
    return Screen_ToString(s.instance)
}

func (s *TScreen) ActiveForm() *TForm {
    return AsForm(Screen_GetActiveForm(s.instance))
}

func (s *TScreen) CustomFormCount() int32 {
    return Screen_GetCustomFormCount(s.instance)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TScreen) Cursor() TCursor {
    return Screen_GetCursor(s.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TScreen) SetCursor(value TCursor) {
    Screen_SetCursor(s.instance, value)
}

func (s *TScreen) FocusedForm() *TForm {
    return AsForm(Screen_GetFocusedForm(s.instance))
}

func (s *TScreen) MonitorCount() int32 {
    return Screen_GetMonitorCount(s.instance)
}

func (s *TScreen) DesktopRect() TRect {
    return Screen_GetDesktopRect(s.instance)
}

func (s *TScreen) DesktopHeight() int32 {
    return Screen_GetDesktopHeight(s.instance)
}

func (s *TScreen) DesktopLeft() int32 {
    return Screen_GetDesktopLeft(s.instance)
}

func (s *TScreen) DesktopTop() int32 {
    return Screen_GetDesktopTop(s.instance)
}

func (s *TScreen) DesktopWidth() int32 {
    return Screen_GetDesktopWidth(s.instance)
}

func (s *TScreen) WorkAreaRect() TRect {
    return Screen_GetWorkAreaRect(s.instance)
}

func (s *TScreen) WorkAreaHeight() int32 {
    return Screen_GetWorkAreaHeight(s.instance)
}

func (s *TScreen) WorkAreaLeft() int32 {
    return Screen_GetWorkAreaLeft(s.instance)
}

func (s *TScreen) WorkAreaTop() int32 {
    return Screen_GetWorkAreaTop(s.instance)
}

func (s *TScreen) WorkAreaWidth() int32 {
    return Screen_GetWorkAreaWidth(s.instance)
}

func (s *TScreen) Fonts() *TStrings {
    return AsStrings(Screen_GetFonts(s.instance))
}

func (s *TScreen) FormCount() int32 {
    return Screen_GetFormCount(s.instance)
}

// CN: 获取高度。
// EN: Get height.
func (s *TScreen) Height() int32 {
    return Screen_GetHeight(s.instance)
}

func (s *TScreen) PixelsPerInch() int32 {
    return Screen_GetPixelsPerInch(s.instance)
}

func (s *TScreen) PrimaryMonitor() *TMonitor {
    return AsMonitor(Screen_GetPrimaryMonitor(s.instance))
}

// CN: 获取宽度。
// EN: Get width.
func (s *TScreen) Width() int32 {
    return Screen_GetWidth(s.instance)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TScreen) ComponentCount() int32 {
    return Screen_GetComponentCount(s.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (s *TScreen) ComponentIndex() int32 {
    return Screen_GetComponentIndex(s.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (s *TScreen) SetComponentIndex(value int32) {
    Screen_SetComponentIndex(s.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TScreen) Owner() *TComponent {
    return AsComponent(Screen_GetOwner(s.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (s *TScreen) Name() string {
    return Screen_GetName(s.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (s *TScreen) SetName(value string) {
    Screen_SetName(s.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TScreen) Tag() int {
    return Screen_GetTag(s.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TScreen) SetTag(value int) {
    Screen_SetTag(s.instance, value)
}

func (s *TScreen) Cursors(Index int32) HICON {
    return Screen_GetCursors(s.instance, Index)
}

func (s *TScreen) SetCursors(Index int32, value HICON) {
    Screen_SetCursors(s.instance, Index, value)
}

func (s *TScreen) Monitors(Index int32) *TMonitor {
    return AsMonitor(Screen_GetMonitors(s.instance, Index))
}

func (s *TScreen) Forms(Index int32) *TForm {
    return AsForm(Screen_GetForms(s.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TScreen) Components(AIndex int32) *TComponent {
    return AsComponent(Screen_GetComponents(s.instance, AIndex))
}

