
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

type TComponent struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewComponent(owner IComponent) *TComponent {
    c := new(TComponent)
    c.instance = Component_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(c, (*TComponent).Free)
    return c
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsComponent(obj interface{}) *TComponent {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TComponent{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsComponent.
func ComponentFromInst(inst uintptr) *TComponent {
    return AsComponent(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsComponent.
func ComponentFromObj(obj IObject) *TComponent {
    return AsComponent(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsComponent.
func ComponentFromUnsafePointer(ptr unsafe.Pointer) *TComponent {
    return AsComponent(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (c *TComponent) Free() {
    if c.instance != 0 {
        Component_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TComponent) Instance() uintptr {
    return c.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TComponent) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TComponent) IsValid() bool {
    return c.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (c *TComponent) Is() TIs {
    return TIs(c.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (c *TComponent) As() TAs {
//    return TAs(c.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TComponentClass() TClass {
    return Component_StaticClassType()
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TComponent) FindComponent(AName string) *TComponent {
    return AsComponent(Component_FindComponent(c.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TComponent) GetNamePath() string {
    return Component_GetNamePath(c.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TComponent) HasParent() bool {
    return Component_HasParent(c.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TComponent) Assign(Source IObject) {
    Component_Assign(c.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TComponent) ClassType() TClass {
    return Component_ClassType(c.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TComponent) ClassName() string {
    return Component_ClassName(c.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TComponent) InstanceSize() int32 {
    return Component_InstanceSize(c.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TComponent) InheritsFrom(AClass TClass) bool {
    return Component_InheritsFrom(c.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TComponent) Equals(Obj IObject) bool {
    return Component_Equals(c.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TComponent) GetHashCode() int32 {
    return Component_GetHashCode(c.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (c *TComponent) ToString() string {
    return Component_ToString(c.instance)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TComponent) ComponentCount() int32 {
    return Component_GetComponentCount(c.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (c *TComponent) ComponentIndex() int32 {
    return Component_GetComponentIndex(c.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (c *TComponent) SetComponentIndex(value int32) {
    Component_SetComponentIndex(c.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TComponent) Owner() *TComponent {
    return AsComponent(Component_GetOwner(c.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (c *TComponent) Name() string {
    return Component_GetName(c.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (c *TComponent) SetName(value string) {
    Component_SetName(c.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TComponent) Tag() int {
    return Component_GetTag(c.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TComponent) SetTag(value int) {
    Component_SetTag(c.instance, value)
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TComponent) Components(AIndex int32) *TComponent {
    return AsComponent(Component_GetComponents(c.instance, AIndex))
}

