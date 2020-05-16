
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

type TClipboard struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewClipboard() *TClipboard {
    c := new(TClipboard)
    c.instance = Clipboard_Create()
    c.ptr = unsafe.Pointer(c.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(c, (*TClipboard).Free)
    return c
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsClipboard(obj interface{}) *TClipboard {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TClipboard{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsClipboard.
func ClipboardFromInst(inst uintptr) *TClipboard {
    return AsClipboard(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsClipboard.
func ClipboardFromObj(obj IObject) *TClipboard {
    return AsClipboard(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsClipboard.
func ClipboardFromUnsafePointer(ptr unsafe.Pointer) *TClipboard {
    return AsClipboard(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (c *TClipboard) Free() {
    if c.instance != 0 {
        Clipboard_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TClipboard) Instance() uintptr {
    return c.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TClipboard) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TClipboard) IsValid() bool {
    return c.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (c *TClipboard) Is() TIs {
    return TIs(c.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (c *TClipboard) As() TAs {
//    return TAs(c.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TClipboardClass() TClass {
    return Clipboard_StaticClassType()
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TClipboard) Assign(Source IObject) {
    Clipboard_Assign(c.instance, CheckPtr(Source))
}

// CN: 清除。
// EN: .
func (c *TClipboard) Clear() {
    Clipboard_Clear(c.instance)
}

// CN: 关闭。
// EN: .
func (c *TClipboard) Close() {
    Clipboard_Close(c.instance)
}

func (c *TClipboard) HasFormat(Format uint16) bool {
    return Clipboard_HasFormat(c.instance, Format)
}

func (c *TClipboard) Open() {
    Clipboard_Open(c.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TClipboard) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Clipboard_GetTextBuf(c.instance, Buffer , BufSize)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TClipboard) SetTextBuf(Buffer string) {
    Clipboard_SetTextBuf(c.instance, Buffer)
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TClipboard) GetNamePath() string {
    return Clipboard_GetNamePath(c.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TClipboard) ClassType() TClass {
    return Clipboard_ClassType(c.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TClipboard) ClassName() string {
    return Clipboard_ClassName(c.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TClipboard) InstanceSize() int32 {
    return Clipboard_InstanceSize(c.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TClipboard) InheritsFrom(AClass TClass) bool {
    return Clipboard_InheritsFrom(c.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TClipboard) Equals(Obj IObject) bool {
    return Clipboard_Equals(c.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TClipboard) GetHashCode() int32 {
    return Clipboard_GetHashCode(c.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (c *TClipboard) ToString() string {
    return Clipboard_ToString(c.instance)
}

func (c *TClipboard) AsText() string {
    return Clipboard_GetAsText(c.instance)
}

func (c *TClipboard) SetAsText(value string) {
    Clipboard_SetAsText(c.instance, value)
}

func (c *TClipboard) FormatCount() int32 {
    return Clipboard_GetFormatCount(c.instance)
}

func (c *TClipboard) Formats(Index int32) uint16 {
    return Clipboard_GetFormats(c.instance, Index)
}

