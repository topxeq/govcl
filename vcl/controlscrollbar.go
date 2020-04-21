
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

type TControlScrollBar struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsControlScrollBar(obj interface{}) *TControlScrollBar {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TControlScrollBar{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsControlScrollBar.
func ControlScrollBarFromInst(inst uintptr) *TControlScrollBar {
    return AsControlScrollBar(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsControlScrollBar.
func ControlScrollBarFromObj(obj IObject) *TControlScrollBar {
    return AsControlScrollBar(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsControlScrollBar.
func ControlScrollBarFromUnsafePointer(ptr unsafe.Pointer) *TControlScrollBar {
    return AsControlScrollBar(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TControlScrollBar) Instance() uintptr {
    return c.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TControlScrollBar) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TControlScrollBar) IsValid() bool {
    return c.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (c *TControlScrollBar) Is() TIs {
    return TIs(c.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (c *TControlScrollBar) As() TAs {
//    return TAs(c.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TControlScrollBarClass() TClass {
    return ControlScrollBar_StaticClassType()
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TControlScrollBar) Assign(Source IObject) {
    ControlScrollBar_Assign(c.instance, CheckPtr(Source))
}

func (c *TControlScrollBar) IsScrollBarVisible() bool {
    return ControlScrollBar_IsScrollBarVisible(c.instance)
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TControlScrollBar) GetNamePath() string {
    return ControlScrollBar_GetNamePath(c.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TControlScrollBar) ClassType() TClass {
    return ControlScrollBar_ClassType(c.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TControlScrollBar) ClassName() string {
    return ControlScrollBar_ClassName(c.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TControlScrollBar) InstanceSize() int32 {
    return ControlScrollBar_InstanceSize(c.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TControlScrollBar) InheritsFrom(AClass TClass) bool {
    return ControlScrollBar_InheritsFrom(c.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TControlScrollBar) Equals(Obj IObject) bool {
    return ControlScrollBar_Equals(c.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TControlScrollBar) GetHashCode() int32 {
    return ControlScrollBar_GetHashCode(c.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (c *TControlScrollBar) ToString() string {
    return ControlScrollBar_ToString(c.instance)
}

func (c *TControlScrollBar) ScrollPos() int32 {
    return ControlScrollBar_GetScrollPos(c.instance)
}

func (c *TControlScrollBar) Increment() TScrollBarInc {
    return ControlScrollBar_GetIncrement(c.instance)
}

func (c *TControlScrollBar) SetIncrement(value TScrollBarInc) {
    ControlScrollBar_SetIncrement(c.instance, value)
}

func (c *TControlScrollBar) Position() int32 {
    return ControlScrollBar_GetPosition(c.instance)
}

func (c *TControlScrollBar) SetPosition(value int32) {
    ControlScrollBar_SetPosition(c.instance, value)
}

func (c *TControlScrollBar) Range() int32 {
    return ControlScrollBar_GetRange(c.instance)
}

func (c *TControlScrollBar) SetRange(value int32) {
    ControlScrollBar_SetRange(c.instance, value)
}

func (c *TControlScrollBar) Smooth() bool {
    return ControlScrollBar_GetSmooth(c.instance)
}

func (c *TControlScrollBar) SetSmooth(value bool) {
    ControlScrollBar_SetSmooth(c.instance, value)
}

func (c *TControlScrollBar) Tracking() bool {
    return ControlScrollBar_GetTracking(c.instance)
}

func (c *TControlScrollBar) SetTracking(value bool) {
    ControlScrollBar_SetTracking(c.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TControlScrollBar) Visible() bool {
    return ControlScrollBar_GetVisible(c.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TControlScrollBar) SetVisible(value bool) {
    ControlScrollBar_SetVisible(c.instance, value)
}
