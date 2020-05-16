
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

type TPicture struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPicture() *TPicture {
    p := new(TPicture)
    p.instance = Picture_Create()
    p.ptr = unsafe.Pointer(p.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(p, (*TPicture).Free)
    return p
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsPicture(obj interface{}) *TPicture {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TPicture{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsPicture.
func PictureFromInst(inst uintptr) *TPicture {
    return AsPicture(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsPicture.
func PictureFromObj(obj IObject) *TPicture {
    return AsPicture(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsPicture.
func PictureFromUnsafePointer(ptr unsafe.Pointer) *TPicture {
    return AsPicture(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (p *TPicture) Free() {
    if p.instance != 0 {
        Picture_Free(p.instance)
        p.instance, p.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPicture) Instance() uintptr {
    return p.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPicture) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPicture) IsValid() bool {
    return p.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (p *TPicture) Is() TIs {
    return TIs(p.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (p *TPicture) As() TAs {
//    return TAs(p.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPictureClass() TClass {
    return Picture_StaticClassType()
}

// CN: 从文件加载。
// EN: .
func (p *TPicture) LoadFromFile(Filename string) {
    Picture_LoadFromFile(p.instance, Filename)
}

// CN: 保存至文件。
// EN: .
func (p *TPicture) SaveToFile(Filename string) {
    Picture_SaveToFile(p.instance, Filename)
}

// CN: 文件流加载。
// EN: .
func (p *TPicture) LoadFromStream(Stream IObject) {
    Picture_LoadFromStream(p.instance, CheckPtr(Stream))
}

// CN: 保存至流。
// EN: .
func (p *TPicture) SaveToStream(Stream IObject) {
    Picture_SaveToStream(p.instance, CheckPtr(Stream))
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPicture) Assign(Source IObject) {
    Picture_Assign(p.instance, CheckPtr(Source))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPicture) GetNamePath() string {
    return Picture_GetNamePath(p.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPicture) ClassType() TClass {
    return Picture_ClassType(p.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPicture) ClassName() string {
    return Picture_ClassName(p.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPicture) InstanceSize() int32 {
    return Picture_InstanceSize(p.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPicture) InheritsFrom(AClass TClass) bool {
    return Picture_InheritsFrom(p.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPicture) Equals(Obj IObject) bool {
    return Picture_Equals(p.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPicture) GetHashCode() int32 {
    return Picture_GetHashCode(p.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (p *TPicture) ToString() string {
    return Picture_ToString(p.instance)
}

func (p *TPicture) Bitmap() *TBitmap {
    return AsBitmap(Picture_GetBitmap(p.instance))
}

func (p *TPicture) SetBitmap(value *TBitmap) {
    Picture_SetBitmap(p.instance, CheckPtr(value))
}

func (p *TPicture) Graphic() *TGraphic {
    return AsGraphic(Picture_GetGraphic(p.instance))
}

func (p *TPicture) SetGraphic(value *TGraphic) {
    Picture_SetGraphic(p.instance, CheckPtr(value))
}

// CN: 获取高度。
// EN: Get height.
func (p *TPicture) Height() int32 {
    return Picture_GetHeight(p.instance)
}

// CN: 获取图标。
// EN: Get icon.
func (p *TPicture) Icon() *TIcon {
    return AsIcon(Picture_GetIcon(p.instance))
}

// CN: 设置图标。
// EN: Set icon.
func (p *TPicture) SetIcon(value *TIcon) {
    Picture_SetIcon(p.instance, CheckPtr(value))
}

// CN: 获取宽度。
// EN: Get width.
func (p *TPicture) Width() int32 {
    return Picture_GetWidth(p.instance)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (p *TPicture) SetOnChange(fn TNotifyEvent) {
    Picture_SetOnChange(p.instance, fn)
}

