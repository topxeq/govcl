
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

type TMargins struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsMargins(obj interface{}) *TMargins {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TMargins{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsMargins.
func MarginsFromInst(inst uintptr) *TMargins {
    return AsMargins(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsMargins.
func MarginsFromObj(obj IObject) *TMargins {
    return AsMargins(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsMargins.
func MarginsFromUnsafePointer(ptr unsafe.Pointer) *TMargins {
    return AsMargins(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMargins) Instance() uintptr {
    return m.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMargins) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMargins) IsValid() bool {
    return m.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (m *TMargins) Is() TIs {
    return TIs(m.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (m *TMargins) As() TAs {
//    return TAs(m.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMarginsClass() TClass {
    return Margins_StaticClassType()
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (m *TMargins) SetBounds(ALeft int32, ATop int32, ARight int32, ABottom int32) {
    Margins_SetBounds(m.instance, ALeft , ATop , ARight , ABottom)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (m *TMargins) Assign(Source IObject) {
    Margins_Assign(m.instance, CheckPtr(Source))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (m *TMargins) GetNamePath() string {
    return Margins_GetNamePath(m.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMargins) ClassType() TClass {
    return Margins_ClassType(m.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMargins) ClassName() string {
    return Margins_ClassName(m.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMargins) InstanceSize() int32 {
    return Margins_InstanceSize(m.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMargins) InheritsFrom(AClass TClass) bool {
    return Margins_InheritsFrom(m.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMargins) Equals(Obj IObject) bool {
    return Margins_Equals(m.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMargins) GetHashCode() int32 {
    return Margins_GetHashCode(m.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (m *TMargins) ToString() string {
    return Margins_ToString(m.instance)
}

func (m *TMargins) ControlLeft() int32 {
    return Margins_GetControlLeft(m.instance)
}

func (m *TMargins) ControlTop() int32 {
    return Margins_GetControlTop(m.instance)
}

func (m *TMargins) ControlWidth() int32 {
    return Margins_GetControlWidth(m.instance)
}

func (m *TMargins) ControlHeight() int32 {
    return Margins_GetControlHeight(m.instance)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (m *TMargins) SetOnChange(fn TNotifyEvent) {
    Margins_SetOnChange(m.instance, fn)
}

// CN: 获取左边位置。
// EN: Get Left position.
func (m *TMargins) Left() int32 {
    return Margins_GetLeft(m.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (m *TMargins) SetLeft(value int32) {
    Margins_SetLeft(m.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (m *TMargins) Top() int32 {
    return Margins_GetTop(m.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (m *TMargins) SetTop(value int32) {
    Margins_SetTop(m.instance, value)
}

func (m *TMargins) Right() int32 {
    return Margins_GetRight(m.instance)
}

func (m *TMargins) SetRight(value int32) {
    Margins_SetRight(m.instance, value)
}

func (m *TMargins) Bottom() int32 {
    return Margins_GetBottom(m.instance)
}

func (m *TMargins) SetBottom(value int32) {
    Margins_SetBottom(m.instance, value)
}

