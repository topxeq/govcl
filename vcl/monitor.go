
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

type TMonitor struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMonitor() *TMonitor {
    m := new(TMonitor)
    m.instance = Monitor_Create()
    m.ptr = unsafe.Pointer(m.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(m, (*TMonitor).Free)
    return m
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsMonitor(obj interface{}) *TMonitor {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TMonitor{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsMonitor.
func MonitorFromInst(inst uintptr) *TMonitor {
    return AsMonitor(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsMonitor.
func MonitorFromObj(obj IObject) *TMonitor {
    return AsMonitor(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsMonitor.
func MonitorFromUnsafePointer(ptr unsafe.Pointer) *TMonitor {
    return AsMonitor(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (m *TMonitor) Free() {
    if m.instance != 0 {
        Monitor_Free(m.instance)
        m.instance, m.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMonitor) Instance() uintptr {
    return m.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMonitor) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMonitor) IsValid() bool {
    return m.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (m *TMonitor) Is() TIs {
    return TIs(m.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (m *TMonitor) As() TAs {
//    return TAs(m.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMonitorClass() TClass {
    return Monitor_StaticClassType()
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMonitor) ClassType() TClass {
    return Monitor_ClassType(m.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMonitor) ClassName() string {
    return Monitor_ClassName(m.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMonitor) InstanceSize() int32 {
    return Monitor_InstanceSize(m.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMonitor) InheritsFrom(AClass TClass) bool {
    return Monitor_InheritsFrom(m.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMonitor) Equals(Obj IObject) bool {
    return Monitor_Equals(m.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMonitor) GetHashCode() int32 {
    return Monitor_GetHashCode(m.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (m *TMonitor) ToString() string {
    return Monitor_ToString(m.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (m *TMonitor) Handle() HMONITOR {
    return Monitor_GetHandle(m.instance)
}

func (m *TMonitor) MonitorNum() int32 {
    return Monitor_GetMonitorNum(m.instance)
}

// CN: 获取左边位置。
// EN: Get Left position.
func (m *TMonitor) Left() int32 {
    return Monitor_GetLeft(m.instance)
}

// CN: 获取高度。
// EN: Get height.
func (m *TMonitor) Height() int32 {
    return Monitor_GetHeight(m.instance)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (m *TMonitor) Top() int32 {
    return Monitor_GetTop(m.instance)
}

// CN: 获取宽度。
// EN: Get width.
func (m *TMonitor) Width() int32 {
    return Monitor_GetWidth(m.instance)
}

func (m *TMonitor) BoundsRect() TRect {
    return Monitor_GetBoundsRect(m.instance)
}

func (m *TMonitor) WorkareaRect() TRect {
    return Monitor_GetWorkareaRect(m.instance)
}

func (m *TMonitor) Primary() bool {
    return Monitor_GetPrimary(m.instance)
}

func (m *TMonitor) PixelsPerInch() int32 {
    return Monitor_GetPixelsPerInch(m.instance)
}

