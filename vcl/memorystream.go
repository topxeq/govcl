
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

type TMemoryStream struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMemoryStream() *TMemoryStream {
    m := new(TMemoryStream)
    m.instance = MemoryStream_Create()
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsMemoryStream(obj interface{}) *TMemoryStream {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TMemoryStream{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsMemoryStream.
func MemoryStreamFromInst(inst uintptr) *TMemoryStream {
    return AsMemoryStream(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsMemoryStream.
func MemoryStreamFromObj(obj IObject) *TMemoryStream {
    return AsMemoryStream(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsMemoryStream.
func MemoryStreamFromUnsafePointer(ptr unsafe.Pointer) *TMemoryStream {
    return AsMemoryStream(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (m *TMemoryStream) Free() {
    if m.instance != 0 {
        MemoryStream_Free(m.instance)
        m.instance, m.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMemoryStream) Instance() uintptr {
    return m.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMemoryStream) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMemoryStream) IsValid() bool {
    return m.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (m *TMemoryStream) Is() TIs {
    return TIs(m.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (m *TMemoryStream) As() TAs {
//    return TAs(m.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMemoryStreamClass() TClass {
    return MemoryStream_StaticClassType()
}

// CN: 清除。
// EN: .
func (m *TMemoryStream) Clear() {
    MemoryStream_Clear(m.instance)
}

// CN: 文件流加载。
// EN: .
func (m *TMemoryStream) LoadFromStream(Stream IObject) {
    MemoryStream_LoadFromStream(m.instance, CheckPtr(Stream))
}

// CN: 从文件加载。
// EN: .
func (m *TMemoryStream) LoadFromFile(FileName string) {
    MemoryStream_LoadFromFile(m.instance, FileName)
}

func (m *TMemoryStream) Seek(Offset int64, Origin TSeekOrigin) int64 {
    return MemoryStream_Seek(m.instance, Offset , Origin)
}

// CN: 保存至流。
// EN: .
func (m *TMemoryStream) SaveToStream(Stream IObject) {
    MemoryStream_SaveToStream(m.instance, CheckPtr(Stream))
}

// CN: 保存至文件。
// EN: .
func (m *TMemoryStream) SaveToFile(FileName string) {
    MemoryStream_SaveToFile(m.instance, FileName)
}

func (m *TMemoryStream) CopyFrom(Source IObject, Count int64) int64 {
    return MemoryStream_CopyFrom(m.instance, CheckPtr(Source), Count)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMemoryStream) ClassType() TClass {
    return MemoryStream_ClassType(m.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMemoryStream) ClassName() string {
    return MemoryStream_ClassName(m.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMemoryStream) InstanceSize() int32 {
    return MemoryStream_InstanceSize(m.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMemoryStream) InheritsFrom(AClass TClass) bool {
    return MemoryStream_InheritsFrom(m.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMemoryStream) Equals(Obj IObject) bool {
    return MemoryStream_Equals(m.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMemoryStream) GetHashCode() int32 {
    return MemoryStream_GetHashCode(m.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (m *TMemoryStream) ToString() string {
    return MemoryStream_ToString(m.instance)
}

func (m *TMemoryStream) Memory() uintptr {
    return MemoryStream_GetMemory(m.instance)
}

func (m *TMemoryStream) Position() int64 {
    return MemoryStream_GetPosition(m.instance)
}

func (m *TMemoryStream) SetPosition(value int64) {
    MemoryStream_SetPosition(m.instance, value)
}

func (m *TMemoryStream) Size() int64 {
    return MemoryStream_GetSize(m.instance)
}

func (m *TMemoryStream) SetSize(value int64) {
    MemoryStream_SetSize(m.instance, value)
}
