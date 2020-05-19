//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"unsafe"

	. "github.com/topxeq/govcl/vcl/api"
	. "github.com/topxeq/govcl/vcl/types"
)

type TTaskDialog struct {
	IComponent
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTaskDialog(owner IComponent) *TTaskDialog {
	t := new(TTaskDialog)
	t.instance = TaskDialog_Create(CheckPtr(owner))
	t.ptr = unsafe.Pointer(t.instance)
	// 不敢启用，因为不知道会发生什么...
	// runtime.SetFinalizer(t, (*TTaskDialog).Free)
	return t
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsTaskDialog(obj interface{}) *TTaskDialog {
	instance, ptr := getInstance(obj)
	if instance == 0 {
		return nil
	}
	return &TTaskDialog{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsTaskDialog.
func TaskDialogFromInst(inst uintptr) *TTaskDialog {
	return AsTaskDialog(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsTaskDialog.
func TaskDialogFromObj(obj IObject) *TTaskDialog {
	return AsTaskDialog(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTaskDialog.
func TaskDialogFromUnsafePointer(ptr unsafe.Pointer) *TTaskDialog {
	return AsTaskDialog(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (t *TTaskDialog) Free() {
	if t.instance != 0 {
		TaskDialog_Free(t.instance)
		t.instance, t.ptr = 0, nullptr
	}
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTaskDialog) Instance() uintptr {
	return t.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTaskDialog) UnsafeAddr() unsafe.Pointer {
	return t.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTaskDialog) IsValid() bool {
	return t.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (t *TTaskDialog) Is() TIs {
	return TIs(t.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (t *TTaskDialog) As() TAs {
//    return TAs(t.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTaskDialogClass() TClass {
	return TaskDialog_StaticClassType()
}

// CN: 执行。
// EN: .
func (t *TTaskDialog) Execute() bool {
	return TaskDialog_Execute(t.instance)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (t *TTaskDialog) FindComponent(AName string) *TComponent {
	return AsComponent(TaskDialog_FindComponent(t.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTaskDialog) GetNamePath() string {
	return TaskDialog_GetNamePath(t.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (t *TTaskDialog) HasParent() bool {
	return TaskDialog_HasParent(t.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTaskDialog) Assign(Source IObject) {
	TaskDialog_Assign(t.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTaskDialog) ClassType() TClass {
	return TaskDialog_ClassType(t.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTaskDialog) ClassName() string {
	return TaskDialog_ClassName(t.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTaskDialog) InstanceSize() int32 {
	return TaskDialog_InstanceSize(t.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTaskDialog) InheritsFrom(AClass TClass) bool {
	return TaskDialog_InheritsFrom(t.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTaskDialog) Equals(Obj IObject) bool {
	return TaskDialog_Equals(t.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTaskDialog) GetHashCode() int32 {
	return TaskDialog_GetHashCode(t.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (t *TTaskDialog) ToString() string {
	return TaskDialog_ToString(t.instance)
}

func (t *TTaskDialog) Buttons() *TTaskDialogButtons {
	return AsTaskDialogButtons(TaskDialog_GetButtons(t.instance))
}

func (t *TTaskDialog) SetButtons(value *TTaskDialogButtons) {
	TaskDialog_SetButtons(t.instance, CheckPtr(value))
}

// CN: 获取控件标题。
// EN: Get the control title.
func (t *TTaskDialog) Caption() string {
	return TaskDialog_GetCaption(t.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (t *TTaskDialog) SetCaption(value string) {
	TaskDialog_SetCaption(t.instance, value)
}

func (t *TTaskDialog) CommonButtons() TTaskDialogCommonButtons {
	return TaskDialog_GetCommonButtons(t.instance)
}

func (t *TTaskDialog) SetCommonButtons(value TTaskDialogCommonButtons) {
	TaskDialog_SetCommonButtons(t.instance, value)
}

func (t *TTaskDialog) DefaultButton() TTaskDialogCommonButton {
	return TaskDialog_GetDefaultButton(t.instance)
}

func (t *TTaskDialog) SetDefaultButton(value TTaskDialogCommonButton) {
	TaskDialog_SetDefaultButton(t.instance, value)
}

func (t *TTaskDialog) ExpandButtonCaption() string {
	return TaskDialog_GetExpandButtonCaption(t.instance)
}

func (t *TTaskDialog) SetExpandButtonCaption(value string) {
	TaskDialog_SetExpandButtonCaption(t.instance, value)
}

func (t *TTaskDialog) ExpandedText() string {
	return TaskDialog_GetExpandedText(t.instance)
}

func (t *TTaskDialog) SetExpandedText(value string) {
	TaskDialog_SetExpandedText(t.instance, value)
}

func (t *TTaskDialog) Flags() TTaskDialogFlags {
	return TaskDialog_GetFlags(t.instance)
}

func (t *TTaskDialog) SetFlags(value TTaskDialogFlags) {
	TaskDialog_SetFlags(t.instance, value)
}

func (t *TTaskDialog) FooterIcon() TTaskDialogIcon {
	return TaskDialog_GetFooterIcon(t.instance)
}

func (t *TTaskDialog) SetFooterIcon(value TTaskDialogIcon) {
	TaskDialog_SetFooterIcon(t.instance, value)
}

func (t *TTaskDialog) FooterText() string {
	return TaskDialog_GetFooterText(t.instance)
}

func (t *TTaskDialog) SetFooterText(value string) {
	TaskDialog_SetFooterText(t.instance, value)
}

func (t *TTaskDialog) MainIcon() TTaskDialogIcon {
	return TaskDialog_GetMainIcon(t.instance)
}

func (t *TTaskDialog) SetMainIcon(value TTaskDialogIcon) {
	TaskDialog_SetMainIcon(t.instance, value)
}

func (t *TTaskDialog) RadioButtons() *TTaskDialogButtons {
	return AsTaskDialogButtons(TaskDialog_GetRadioButtons(t.instance))
}

func (t *TTaskDialog) SetRadioButtons(value *TTaskDialogButtons) {
	TaskDialog_SetRadioButtons(t.instance, CheckPtr(value))
}

// CN: 获取文本。
// EN: .
func (t *TTaskDialog) Text() string {
	return TaskDialog_GetText(t.instance)
}

// CN: 设置文本。
// EN: .
func (t *TTaskDialog) SetText(value string) {
	TaskDialog_SetText(t.instance, value)
}

func (t *TTaskDialog) Title() string {
	return TaskDialog_GetTitle(t.instance)
}

func (t *TTaskDialog) SetTitle(value string) {
	TaskDialog_SetTitle(t.instance, value)
}

func (t *TTaskDialog) VerificationText() string {
	return TaskDialog_GetVerificationText(t.instance)
}

func (t *TTaskDialog) SetVerificationText(value string) {
	TaskDialog_SetVerificationText(t.instance, value)
}

func (t *TTaskDialog) SetOnButtonClicked(fn *TTaskDlgClickEvent) {
	TaskDialog_SetOnButtonClicked(t.instance, fn)
}

func (t *TTaskDialog) Button() *TTaskDialogButtonItem {
	return AsTaskDialogButtonItem(TaskDialog_GetButton(t.instance))
}

func (t *TTaskDialog) SetButton(value *TTaskDialogButtonItem) {
	TaskDialog_SetButton(t.instance, CheckPtr(value))
}

// CN: 获取模态对话框显示结果。
// EN: .
func (t *TTaskDialog) ModalResult() TModalResult {
	return TaskDialog_GetModalResult(t.instance)
}

// CN: 设置模态对话框显示结果。
// EN: .
func (t *TTaskDialog) SetModalResult(value TModalResult) {
	TaskDialog_SetModalResult(t.instance, value)
}

func (t *TTaskDialog) RadioButton() *TTaskDialogRadioButtonItem {
	return AsTaskDialogRadioButtonItem(TaskDialog_GetRadioButton(t.instance))
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTaskDialog) ComponentCount() int32 {
	return TaskDialog_GetComponentCount(t.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (t *TTaskDialog) ComponentIndex() int32 {
	return TaskDialog_GetComponentIndex(t.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (t *TTaskDialog) SetComponentIndex(value int32) {
	TaskDialog_SetComponentIndex(t.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTaskDialog) Owner() *TComponent {
	return AsComponent(TaskDialog_GetOwner(t.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTaskDialog) Name() string {
	return TaskDialog_GetName(t.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTaskDialog) SetName(value string) {
	TaskDialog_SetName(t.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTaskDialog) Tag() int {
	return TaskDialog_GetTag(t.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTaskDialog) SetTag(value int) {
	TaskDialog_SetTag(t.instance, value)
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTaskDialog) Components(AIndex int32) *TComponent {
	return AsComponent(TaskDialog_GetComponents(t.instance, AIndex))
}
