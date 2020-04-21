
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

type TListItems struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListItems(AOwner *TListView) *TListItems {
    l := new(TListItems)
    l.instance = ListItems_Create(CheckPtr(AOwner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsListItems(obj interface{}) *TListItems {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TListItems{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsListItems.
func ListItemsFromInst(inst uintptr) *TListItems {
    return AsListItems(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsListItems.
func ListItemsFromObj(obj IObject) *TListItems {
    return AsListItems(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsListItems.
func ListItemsFromUnsafePointer(ptr unsafe.Pointer) *TListItems {
    return AsListItems(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (l *TListItems) Free() {
    if l.instance != 0 {
        ListItems_Free(l.instance)
        l.instance, l.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListItems) Instance() uintptr {
    return l.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListItems) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListItems) IsValid() bool {
    return l.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (l *TListItems) Is() TIs {
    return TIs(l.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (l *TListItems) As() TAs {
//    return TAs(l.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListItemsClass() TClass {
    return ListItems_StaticClassType()
}

func (l *TListItems) Add() *TListItem {
    return AsListItem(ListItems_Add(l.instance))
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TListItems) Assign(Source IObject) {
    ListItems_Assign(l.instance, CheckPtr(Source))
}

func (l *TListItems) BeginUpdate() {
    ListItems_BeginUpdate(l.instance)
}

// CN: 清除。
// EN: .
func (l *TListItems) Clear() {
    ListItems_Clear(l.instance)
}

func (l *TListItems) Delete(Index int32) {
    ListItems_Delete(l.instance, Index)
}

func (l *TListItems) EndUpdate() {
    ListItems_EndUpdate(l.instance)
}

func (l *TListItems) IndexOf(Value *TListItem) int32 {
    return ListItems_IndexOf(l.instance, CheckPtr(Value))
}

func (l *TListItems) Insert(Index int32) *TListItem {
    return AsListItem(ListItems_Insert(l.instance, Index))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TListItems) GetNamePath() string {
    return ListItems_GetNamePath(l.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListItems) ClassType() TClass {
    return ListItems_ClassType(l.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListItems) ClassName() string {
    return ListItems_ClassName(l.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListItems) InstanceSize() int32 {
    return ListItems_InstanceSize(l.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListItems) InheritsFrom(AClass TClass) bool {
    return ListItems_InheritsFrom(l.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListItems) Equals(Obj IObject) bool {
    return ListItems_Equals(l.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListItems) GetHashCode() int32 {
    return ListItems_GetHashCode(l.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (l *TListItems) ToString() string {
    return ListItems_ToString(l.instance)
}

func (l *TListItems) Count() int32 {
    return ListItems_GetCount(l.instance)
}

func (l *TListItems) SetCount(value int32) {
    ListItems_SetCount(l.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TListItems) Owner() *TWinControl {
    return AsWinControl(ListItems_GetOwner(l.instance))
}

func (l *TListItems) Item(Index int32) *TListItem {
    return AsListItem(ListItems_GetItem(l.instance, Index))
}

func (l *TListItems) SetItem(Index int32, value *TListItem) {
    ListItems_SetItem(l.instance, Index, CheckPtr(value))
}

