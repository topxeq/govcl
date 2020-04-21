
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

type TObject struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewObject() *TObject {
    o := new(TObject)
    o.instance = Object_Create()
    o.ptr = unsafe.Pointer(o.instance)
    return o
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsObject(obj interface{}) *TObject {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TObject{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsObject.
func ObjectFromInst(inst uintptr) *TObject {
    return AsObject(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsObject.
func ObjectFromObj(obj IObject) *TObject {
    return AsObject(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsObject.
func ObjectFromUnsafePointer(ptr unsafe.Pointer) *TObject {
    return AsObject(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (o *TObject) Free() {
    if o.instance != 0 {
        Object_Free(o.instance)
        o.instance, o.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (o *TObject) Instance() uintptr {
    return o.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (o *TObject) UnsafeAddr() unsafe.Pointer {
    return o.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (o *TObject) IsValid() bool {
    return o.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (o *TObject) Is() TIs {
    return TIs(o.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (o *TObject) As() TAs {
//    return TAs(o.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TObjectClass() TClass {
    return Object_StaticClassType()
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (o *TObject) ClassType() TClass {
    return Object_ClassType(o.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (o *TObject) ClassName() string {
    return Object_ClassName(o.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (o *TObject) InstanceSize() int32 {
    return Object_InstanceSize(o.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (o *TObject) InheritsFrom(AClass TClass) bool {
    return Object_InheritsFrom(o.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (o *TObject) Equals(Obj IObject) bool {
    return Object_Equals(o.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (o *TObject) GetHashCode() int32 {
    return Object_GetHashCode(o.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (o *TObject) ToString() string {
    return Object_ToString(o.instance)
}

