
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

type TTimer struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTimer(owner IComponent) *TTimer {
    t := new(TTimer)
    t.instance = Timer_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsTimer(obj interface{}) *TTimer {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TTimer{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsTimer.
func TimerFromInst(inst uintptr) *TTimer {
    return AsTimer(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsTimer.
func TimerFromObj(obj IObject) *TTimer {
    return AsTimer(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTimer.
func TimerFromUnsafePointer(ptr unsafe.Pointer) *TTimer {
    return AsTimer(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (t *TTimer) Free() {
    if t.instance != 0 {
        Timer_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTimer) Instance() uintptr {
    return t.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTimer) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTimer) IsValid() bool {
    return t.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (t *TTimer) Is() TIs {
    return TIs(t.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (t *TTimer) As() TAs {
//    return TAs(t.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTimerClass() TClass {
    return Timer_StaticClassType()
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (t *TTimer) FindComponent(AName string) *TComponent {
    return AsComponent(Timer_FindComponent(t.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTimer) GetNamePath() string {
    return Timer_GetNamePath(t.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (t *TTimer) HasParent() bool {
    return Timer_HasParent(t.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTimer) Assign(Source IObject) {
    Timer_Assign(t.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTimer) ClassType() TClass {
    return Timer_ClassType(t.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTimer) ClassName() string {
    return Timer_ClassName(t.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTimer) InstanceSize() int32 {
    return Timer_InstanceSize(t.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTimer) InheritsFrom(AClass TClass) bool {
    return Timer_InheritsFrom(t.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTimer) Equals(Obj IObject) bool {
    return Timer_Equals(t.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTimer) GetHashCode() int32 {
    return Timer_GetHashCode(t.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (t *TTimer) ToString() string {
    return Timer_ToString(t.instance)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TTimer) Enabled() bool {
    return Timer_GetEnabled(t.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TTimer) SetEnabled(value bool) {
    Timer_SetEnabled(t.instance, value)
}

// CN: 获取时钟每次跳动间隔时间，ms。
// EN: .
func (t *TTimer) Interval() uint32 {
    return Timer_GetInterval(t.instance)
}

// CN: 设置时钟每次跳动间隔时间，ms。
// EN: .
func (t *TTimer) SetInterval(value uint32) {
    Timer_SetInterval(t.instance, value)
}

// CN: 设置时钟事件。
// EN: .
func (t *TTimer) SetOnTimer(fn TNotifyEvent) {
    Timer_SetOnTimer(t.instance, fn)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTimer) ComponentCount() int32 {
    return Timer_GetComponentCount(t.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (t *TTimer) ComponentIndex() int32 {
    return Timer_GetComponentIndex(t.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (t *TTimer) SetComponentIndex(value int32) {
    Timer_SetComponentIndex(t.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTimer) Owner() *TComponent {
    return AsComponent(Timer_GetOwner(t.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTimer) Name() string {
    return Timer_GetName(t.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTimer) SetName(value string) {
    Timer_SetName(t.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTimer) Tag() int {
    return Timer_GetTag(t.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTimer) SetTag(value int) {
    Timer_SetTag(t.instance, value)
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTimer) Components(AIndex int32) *TComponent {
    return AsComponent(Timer_GetComponents(t.instance, AIndex))
}
