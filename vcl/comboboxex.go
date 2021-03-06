
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TComboBoxEx struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewComboBoxEx(owner IComponent) *TComboBoxEx {
    c := new(TComboBoxEx)
    c.instance = ComboBoxEx_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsComboBoxEx(obj interface{}) *TComboBoxEx {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TComboBoxEx{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsComboBoxEx.
func ComboBoxExFromInst(inst uintptr) *TComboBoxEx {
    return AsComboBoxEx(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsComboBoxEx.
func ComboBoxExFromObj(obj IObject) *TComboBoxEx {
    return AsComboBoxEx(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsComboBoxEx.
func ComboBoxExFromUnsafePointer(ptr unsafe.Pointer) *TComboBoxEx {
    return AsComboBoxEx(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (c *TComboBoxEx) Free() {
    if c.instance != 0 {
        ComboBoxEx_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TComboBoxEx) Instance() uintptr {
    return c.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TComboBoxEx) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TComboBoxEx) IsValid() bool {
    return c.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (c *TComboBoxEx) Is() TIs {
    return TIs(c.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (c *TComboBoxEx) As() TAs {
//    return TAs(c.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TComboBoxExClass() TClass {
    return ComboBoxEx_StaticClassType()
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (c *TComboBoxEx) Focused() bool {
    return ComboBoxEx_Focused(c.instance)
}

func (c *TComboBoxEx) AddItem(Item string, AObject IObject) {
    ComboBoxEx_AddItem(c.instance, Item , CheckPtr(AObject))
}

// CN: 清除。
// EN: .
func (c *TComboBoxEx) Clear() {
    ComboBoxEx_Clear(c.instance)
}

// CN: 清除选择。
// EN: .
func (c *TComboBoxEx) ClearSelection() {
    ComboBoxEx_ClearSelection(c.instance)
}

func (c *TComboBoxEx) DeleteSelected() {
    ComboBoxEx_DeleteSelected(c.instance)
}

// CN: 全选。
// EN: .
func (c *TComboBoxEx) SelectAll() {
    ComboBoxEx_SelectAll(c.instance)
}

// CN: 是否可以获得焦点。
// EN: .
func (c *TComboBoxEx) CanFocus() bool {
    return ComboBoxEx_CanFocus(c.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (c *TComboBoxEx) ContainsControl(Control IControl) bool {
    return ComboBoxEx_ContainsControl(c.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (c *TComboBoxEx) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ComboBoxEx_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (c *TComboBoxEx) DisableAlign() {
    ComboBoxEx_DisableAlign(c.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (c *TComboBoxEx) EnableAlign() {
    ComboBoxEx_EnableAlign(c.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (c *TComboBoxEx) FindChildControl(ControlName string) *TControl {
    return AsControl(ComboBoxEx_FindChildControl(c.instance, ControlName))
}

func (c *TComboBoxEx) FlipChildren(AllLevels bool) {
    ComboBoxEx_FlipChildren(c.instance, AllLevels)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TComboBoxEx) HandleAllocated() bool {
    return ComboBoxEx_HandleAllocated(c.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (c *TComboBoxEx) InsertControl(AControl IControl) {
    ComboBoxEx_InsertControl(c.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (c *TComboBoxEx) Invalidate() {
    ComboBoxEx_Invalidate(c.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (c *TComboBoxEx) RemoveControl(AControl IControl) {
    ComboBoxEx_RemoveControl(c.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (c *TComboBoxEx) Realign() {
    ComboBoxEx_Realign(c.instance)
}

// CN: 重绘。
// EN: Repaint.
func (c *TComboBoxEx) Repaint() {
    ComboBoxEx_Repaint(c.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (c *TComboBoxEx) ScaleBy(M int32, D int32) {
    ComboBoxEx_ScaleBy(c.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (c *TComboBoxEx) ScrollBy(DeltaX int32, DeltaY int32) {
    ComboBoxEx_ScrollBy(c.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (c *TComboBoxEx) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ComboBoxEx_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (c *TComboBoxEx) SetFocus() {
    ComboBoxEx_SetFocus(c.instance)
}

// CN: 控件更新。
// EN: Update.
func (c *TComboBoxEx) Update() {
    ComboBoxEx_Update(c.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (c *TComboBoxEx) BringToFront() {
    ComboBoxEx_BringToFront(c.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (c *TComboBoxEx) ClientToScreen(Point TPoint) TPoint {
    return ComboBoxEx_ClientToScreen(c.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (c *TComboBoxEx) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ComboBoxEx_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (c *TComboBoxEx) Dragging() bool {
    return ComboBoxEx_Dragging(c.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TComboBoxEx) HasParent() bool {
    return ComboBoxEx_HasParent(c.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (c *TComboBoxEx) Hide() {
    ComboBoxEx_Hide(c.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (c *TComboBoxEx) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ComboBoxEx_Perform(c.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (c *TComboBoxEx) Refresh() {
    ComboBoxEx_Refresh(c.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (c *TComboBoxEx) ScreenToClient(Point TPoint) TPoint {
    return ComboBoxEx_ScreenToClient(c.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (c *TComboBoxEx) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ComboBoxEx_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (c *TComboBoxEx) SendToBack() {
    ComboBoxEx_SendToBack(c.instance)
}

// CN: 显示控件。
// EN: Show control.
func (c *TComboBoxEx) Show() {
    ComboBoxEx_Show(c.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TComboBoxEx) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ComboBoxEx_GetTextBuf(c.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (c *TComboBoxEx) GetTextLen() int32 {
    return ComboBoxEx_GetTextLen(c.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TComboBoxEx) SetTextBuf(Buffer string) {
    ComboBoxEx_SetTextBuf(c.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TComboBoxEx) FindComponent(AName string) *TComponent {
    return AsComponent(ComboBoxEx_FindComponent(c.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TComboBoxEx) GetNamePath() string {
    return ComboBoxEx_GetNamePath(c.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TComboBoxEx) Assign(Source IObject) {
    ComboBoxEx_Assign(c.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TComboBoxEx) ClassType() TClass {
    return ComboBoxEx_ClassType(c.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TComboBoxEx) ClassName() string {
    return ComboBoxEx_ClassName(c.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TComboBoxEx) InstanceSize() int32 {
    return ComboBoxEx_InstanceSize(c.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TComboBoxEx) InheritsFrom(AClass TClass) bool {
    return ComboBoxEx_InheritsFrom(c.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TComboBoxEx) Equals(Obj IObject) bool {
    return ComboBoxEx_Equals(c.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TComboBoxEx) GetHashCode() int32 {
    return ComboBoxEx_GetHashCode(c.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (c *TComboBoxEx) ToString() string {
    return ComboBoxEx_ToString(c.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (c *TComboBoxEx) Align() TAlign {
    return ComboBoxEx_GetAlign(c.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (c *TComboBoxEx) SetAlign(value TAlign) {
    ComboBoxEx_SetAlign(c.instance, value)
}

func (c *TComboBoxEx) AutoCompleteOptions() TAutoCompleteOptions {
    return ComboBoxEx_GetAutoCompleteOptions(c.instance)
}

func (c *TComboBoxEx) SetAutoCompleteOptions(value TAutoCompleteOptions) {
    ComboBoxEx_SetAutoCompleteOptions(c.instance, value)
}

func (c *TComboBoxEx) ItemsEx() *TComboExItems {
    return AsComboExItems(ComboBoxEx_GetItemsEx(c.instance))
}

func (c *TComboBoxEx) SetItemsEx(value *TComboExItems) {
    ComboBoxEx_SetItemsEx(c.instance, CheckPtr(value))
}

func (c *TComboBoxEx) Style() TComboBoxExStyle {
    return ComboBoxEx_GetStyle(c.instance)
}

func (c *TComboBoxEx) SetStyle(value TComboBoxExStyle) {
    ComboBoxEx_SetStyle(c.instance, value)
}

func (c *TComboBoxEx) StyleEx() TComboBoxExStyles {
    return ComboBoxEx_GetStyleEx(c.instance)
}

func (c *TComboBoxEx) SetStyleEx(value TComboBoxExStyles) {
    ComboBoxEx_SetStyleEx(c.instance, value)
}

func (c *TComboBoxEx) Action() *TAction {
    return AsAction(ComboBoxEx_GetAction(c.instance))
}

func (c *TComboBoxEx) SetAction(value IComponent) {
    ComboBoxEx_SetAction(c.instance, CheckPtr(value))
}

// CN: 获取四个角位置的锚点。
// EN: .
func (c *TComboBoxEx) Anchors() TAnchors {
    return ComboBoxEx_GetAnchors(c.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (c *TComboBoxEx) SetAnchors(value TAnchors) {
    ComboBoxEx_SetAnchors(c.instance, value)
}

func (c *TComboBoxEx) BiDiMode() TBiDiMode {
    return ComboBoxEx_GetBiDiMode(c.instance)
}

func (c *TComboBoxEx) SetBiDiMode(value TBiDiMode) {
    ComboBoxEx_SetBiDiMode(c.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (c *TComboBoxEx) Color() TColor {
    return ComboBoxEx_GetColor(c.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (c *TComboBoxEx) SetColor(value TColor) {
    ComboBoxEx_SetColor(c.instance, value)
}

func (c *TComboBoxEx) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ComboBoxEx_GetConstraints(c.instance))
}

func (c *TComboBoxEx) SetConstraints(value *TSizeConstraints) {
    ComboBoxEx_SetConstraints(c.instance, CheckPtr(value))
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (c *TComboBoxEx) DoubleBuffered() bool {
    return ComboBoxEx_GetDoubleBuffered(c.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (c *TComboBoxEx) SetDoubleBuffered(value bool) {
    ComboBoxEx_SetDoubleBuffered(c.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (c *TComboBoxEx) DragCursor() TCursor {
    return ComboBoxEx_GetDragCursor(c.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (c *TComboBoxEx) SetDragCursor(value TCursor) {
    ComboBoxEx_SetDragCursor(c.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (c *TComboBoxEx) DragKind() TDragKind {
    return ComboBoxEx_GetDragKind(c.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (c *TComboBoxEx) SetDragKind(value TDragKind) {
    ComboBoxEx_SetDragKind(c.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (c *TComboBoxEx) DragMode() TDragMode {
    return ComboBoxEx_GetDragMode(c.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (c *TComboBoxEx) SetDragMode(value TDragMode) {
    ComboBoxEx_SetDragMode(c.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TComboBoxEx) Enabled() bool {
    return ComboBoxEx_GetEnabled(c.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TComboBoxEx) SetEnabled(value bool) {
    ComboBoxEx_SetEnabled(c.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (c *TComboBoxEx) Font() *TFont {
    return AsFont(ComboBoxEx_GetFont(c.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (c *TComboBoxEx) SetFont(value *TFont) {
    ComboBoxEx_SetFont(c.instance, CheckPtr(value))
}

func (c *TComboBoxEx) ItemHeight() int32 {
    return ComboBoxEx_GetItemHeight(c.instance)
}

func (c *TComboBoxEx) SetItemHeight(value int32) {
    ComboBoxEx_SetItemHeight(c.instance, value)
}

// CN: 获取最大长度。
// EN: .
func (c *TComboBoxEx) MaxLength() int32 {
    return ComboBoxEx_GetMaxLength(c.instance)
}

// CN: 设置最大长度。
// EN: .
func (c *TComboBoxEx) SetMaxLength(value int32) {
    ComboBoxEx_SetMaxLength(c.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (c *TComboBoxEx) ParentColor() bool {
    return ComboBoxEx_GetParentColor(c.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (c *TComboBoxEx) SetParentColor(value bool) {
    ComboBoxEx_SetParentColor(c.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (c *TComboBoxEx) ParentDoubleBuffered() bool {
    return ComboBoxEx_GetParentDoubleBuffered(c.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (c *TComboBoxEx) SetParentDoubleBuffered(value bool) {
    ComboBoxEx_SetParentDoubleBuffered(c.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (c *TComboBoxEx) ParentFont() bool {
    return ComboBoxEx_GetParentFont(c.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (c *TComboBoxEx) SetParentFont(value bool) {
    ComboBoxEx_SetParentFont(c.instance, value)
}

func (c *TComboBoxEx) ParentShowHint() bool {
    return ComboBoxEx_GetParentShowHint(c.instance)
}

func (c *TComboBoxEx) SetParentShowHint(value bool) {
    ComboBoxEx_SetParentShowHint(c.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (c *TComboBoxEx) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ComboBoxEx_GetPopupMenu(c.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (c *TComboBoxEx) SetPopupMenu(value IComponent) {
    ComboBoxEx_SetPopupMenu(c.instance, CheckPtr(value))
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (c *TComboBoxEx) ShowHint() bool {
    return ComboBoxEx_GetShowHint(c.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (c *TComboBoxEx) SetShowHint(value bool) {
    ComboBoxEx_SetShowHint(c.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (c *TComboBoxEx) TabOrder() TTabOrder {
    return ComboBoxEx_GetTabOrder(c.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (c *TComboBoxEx) SetTabOrder(value TTabOrder) {
    ComboBoxEx_SetTabOrder(c.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (c *TComboBoxEx) TabStop() bool {
    return ComboBoxEx_GetTabStop(c.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (c *TComboBoxEx) SetTabStop(value bool) {
    ComboBoxEx_SetTabStop(c.instance, value)
}

// CN: 获取文本。
// EN: .
func (c *TComboBoxEx) Text() string {
    strLen := c.GetTextLen()
    if strLen != 0 {
        var buffStr string
        c.GetTextBuf(&buffStr, strLen + 1)
        return buffStr
    }
    return ""
}

// CN: 设置文本。
// EN: .
func (c *TComboBoxEx) SetText(value string) {
    ComboBoxEx_SetText(c.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TComboBoxEx) Visible() bool {
    return ComboBoxEx_GetVisible(c.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TComboBoxEx) SetVisible(value bool) {
    ComboBoxEx_SetVisible(c.instance, value)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (c *TComboBoxEx) SetOnChange(fn TNotifyEvent) {
    ComboBoxEx_SetOnChange(c.instance, fn)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TComboBoxEx) SetOnClick(fn TNotifyEvent) {
    ComboBoxEx_SetOnClick(c.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (c *TComboBoxEx) SetOnContextPopup(fn TContextPopupEvent) {
    ComboBoxEx_SetOnContextPopup(c.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (c *TComboBoxEx) SetOnDblClick(fn TNotifyEvent) {
    ComboBoxEx_SetOnDblClick(c.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (c *TComboBoxEx) SetOnDragDrop(fn TDragDropEvent) {
    ComboBoxEx_SetOnDragDrop(c.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (c *TComboBoxEx) SetOnDragOver(fn TDragOverEvent) {
    ComboBoxEx_SetOnDragOver(c.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (c *TComboBoxEx) SetOnEndDock(fn TEndDragEvent) {
    ComboBoxEx_SetOnEndDock(c.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (c *TComboBoxEx) SetOnEndDrag(fn TEndDragEvent) {
    ComboBoxEx_SetOnEndDrag(c.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (c *TComboBoxEx) SetOnEnter(fn TNotifyEvent) {
    ComboBoxEx_SetOnEnter(c.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (c *TComboBoxEx) SetOnExit(fn TNotifyEvent) {
    ComboBoxEx_SetOnExit(c.instance, fn)
}

// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (c *TComboBoxEx) SetOnKeyDown(fn TKeyEvent) {
    ComboBoxEx_SetOnKeyDown(c.instance, fn)
}

func (c *TComboBoxEx) SetOnKeyPress(fn TKeyPressEvent) {
    ComboBoxEx_SetOnKeyPress(c.instance, fn)
}

// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (c *TComboBoxEx) SetOnKeyUp(fn TKeyEvent) {
    ComboBoxEx_SetOnKeyUp(c.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (c *TComboBoxEx) SetOnMouseMove(fn TMouseMoveEvent) {
    ComboBoxEx_SetOnMouseMove(c.instance, fn)
}

func (c *TComboBoxEx) SetOnSelect(fn TNotifyEvent) {
    ComboBoxEx_SetOnSelect(c.instance, fn)
}

// CN: 设置启动停靠。
// EN: .
func (c *TComboBoxEx) SetOnStartDock(fn TStartDockEvent) {
    ComboBoxEx_SetOnStartDock(c.instance, fn)
}

// CN: 获取图标索引列表对象。
// EN: .
func (c *TComboBoxEx) Images() *TImageList {
    return AsImageList(ComboBoxEx_GetImages(c.instance))
}

// CN: 设置图标索引列表对象。
// EN: .
func (c *TComboBoxEx) SetImages(value IComponent) {
    ComboBoxEx_SetImages(c.instance, CheckPtr(value))
}

func (c *TComboBoxEx) DropDownCount() int32 {
    return ComboBoxEx_GetDropDownCount(c.instance)
}

func (c *TComboBoxEx) SetDropDownCount(value int32) {
    ComboBoxEx_SetDropDownCount(c.instance, value)
}

// CN: 获取选择的文本。
// EN: .
func (c *TComboBoxEx) SelText() string {
    return ComboBoxEx_GetSelText(c.instance)
}

// CN: 设置选择的文本。
// EN: .
func (c *TComboBoxEx) SetSelText(value string) {
    ComboBoxEx_SetSelText(c.instance, value)
}

// CN: 获取画布。
// EN: .
func (c *TComboBoxEx) Canvas() *TCanvas {
    return AsCanvas(ComboBoxEx_GetCanvas(c.instance))
}

func (c *TComboBoxEx) DroppedDown() bool {
    return ComboBoxEx_GetDroppedDown(c.instance)
}

func (c *TComboBoxEx) SetDroppedDown(value bool) {
    ComboBoxEx_SetDroppedDown(c.instance, value)
}

func (c *TComboBoxEx) Items() *TStrings {
    return AsStrings(ComboBoxEx_GetItems(c.instance))
}

func (c *TComboBoxEx) SetItems(value IObject) {
    ComboBoxEx_SetItems(c.instance, CheckPtr(value))
}

// CN: 获取选择的长度。
// EN: .
func (c *TComboBoxEx) SelLength() int32 {
    return ComboBoxEx_GetSelLength(c.instance)
}

// CN: 设置选择的长度。
// EN: .
func (c *TComboBoxEx) SetSelLength(value int32) {
    ComboBoxEx_SetSelLength(c.instance, value)
}

// CN: 获取选择的启始位置。
// EN: .
func (c *TComboBoxEx) SelStart() int32 {
    return ComboBoxEx_GetSelStart(c.instance)
}

// CN: 设置选择的启始位置。
// EN: .
func (c *TComboBoxEx) SetSelStart(value int32) {
    ComboBoxEx_SetSelStart(c.instance, value)
}

func (c *TComboBoxEx) ItemIndex() int32 {
    return ComboBoxEx_GetItemIndex(c.instance)
}

func (c *TComboBoxEx) SetItemIndex(value int32) {
    ComboBoxEx_SetItemIndex(c.instance, value)
}

// CN: 获取依靠客户端总数。
// EN: .
func (c *TComboBoxEx) DockClientCount() int32 {
    return ComboBoxEx_GetDockClientCount(c.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (c *TComboBoxEx) DockSite() bool {
    return ComboBoxEx_GetDockSite(c.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (c *TComboBoxEx) SetDockSite(value bool) {
    ComboBoxEx_SetDockSite(c.instance, value)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (c *TComboBoxEx) MouseInClient() bool {
    return ComboBoxEx_GetMouseInClient(c.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (c *TComboBoxEx) VisibleDockClientCount() int32 {
    return ComboBoxEx_GetVisibleDockClientCount(c.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TComboBoxEx) Brush() *TBrush {
    return AsBrush(ComboBoxEx_GetBrush(c.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (c *TComboBoxEx) ControlCount() int32 {
    return ComboBoxEx_GetControlCount(c.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TComboBoxEx) Handle() HWND {
    return ComboBoxEx_GetHandle(c.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (c *TComboBoxEx) ParentWindow() HWND {
    return ComboBoxEx_GetParentWindow(c.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (c *TComboBoxEx) SetParentWindow(value HWND) {
    ComboBoxEx_SetParentWindow(c.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (c *TComboBoxEx) UseDockManager() bool {
    return ComboBoxEx_GetUseDockManager(c.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (c *TComboBoxEx) SetUseDockManager(value bool) {
    ComboBoxEx_SetUseDockManager(c.instance, value)
}

func (c *TComboBoxEx) BoundsRect() TRect {
    return ComboBoxEx_GetBoundsRect(c.instance)
}

func (c *TComboBoxEx) SetBoundsRect(value TRect) {
    ComboBoxEx_SetBoundsRect(c.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (c *TComboBoxEx) ClientHeight() int32 {
    return ComboBoxEx_GetClientHeight(c.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (c *TComboBoxEx) SetClientHeight(value int32) {
    ComboBoxEx_SetClientHeight(c.instance, value)
}

func (c *TComboBoxEx) ClientOrigin() TPoint {
    return ComboBoxEx_GetClientOrigin(c.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (c *TComboBoxEx) ClientRect() TRect {
    return ComboBoxEx_GetClientRect(c.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (c *TComboBoxEx) ClientWidth() int32 {
    return ComboBoxEx_GetClientWidth(c.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (c *TComboBoxEx) SetClientWidth(value int32) {
    ComboBoxEx_SetClientWidth(c.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (c *TComboBoxEx) ControlState() TControlState {
    return ComboBoxEx_GetControlState(c.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (c *TComboBoxEx) SetControlState(value TControlState) {
    ComboBoxEx_SetControlState(c.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (c *TComboBoxEx) ControlStyle() TControlStyle {
    return ComboBoxEx_GetControlStyle(c.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (c *TComboBoxEx) SetControlStyle(value TControlStyle) {
    ComboBoxEx_SetControlStyle(c.instance, value)
}

func (c *TComboBoxEx) Floating() bool {
    return ComboBoxEx_GetFloating(c.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TComboBoxEx) Parent() *TWinControl {
    return AsWinControl(ComboBoxEx_GetParent(c.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TComboBoxEx) SetParent(value IWinControl) {
    ComboBoxEx_SetParent(c.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (c *TComboBoxEx) Left() int32 {
    return ComboBoxEx_GetLeft(c.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (c *TComboBoxEx) SetLeft(value int32) {
    ComboBoxEx_SetLeft(c.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (c *TComboBoxEx) Top() int32 {
    return ComboBoxEx_GetTop(c.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (c *TComboBoxEx) SetTop(value int32) {
    ComboBoxEx_SetTop(c.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (c *TComboBoxEx) Width() int32 {
    return ComboBoxEx_GetWidth(c.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (c *TComboBoxEx) SetWidth(value int32) {
    ComboBoxEx_SetWidth(c.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (c *TComboBoxEx) Height() int32 {
    return ComboBoxEx_GetHeight(c.instance)
}

// CN: 设置高度。
// EN: Set height.
func (c *TComboBoxEx) SetHeight(value int32) {
    ComboBoxEx_SetHeight(c.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TComboBoxEx) Cursor() TCursor {
    return ComboBoxEx_GetCursor(c.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TComboBoxEx) SetCursor(value TCursor) {
    ComboBoxEx_SetCursor(c.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (c *TComboBoxEx) Hint() string {
    return ComboBoxEx_GetHint(c.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (c *TComboBoxEx) SetHint(value string) {
    ComboBoxEx_SetHint(c.instance, value)
}

// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (c *TComboBoxEx) Margins() *TMargins {
    return AsMargins(ComboBoxEx_GetMargins(c.instance))
}

// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (c *TComboBoxEx) SetMargins(value *TMargins) {
    ComboBoxEx_SetMargins(c.instance, CheckPtr(value))
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TComboBoxEx) ComponentCount() int32 {
    return ComboBoxEx_GetComponentCount(c.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (c *TComboBoxEx) ComponentIndex() int32 {
    return ComboBoxEx_GetComponentIndex(c.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (c *TComboBoxEx) SetComponentIndex(value int32) {
    ComboBoxEx_SetComponentIndex(c.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TComboBoxEx) Owner() *TComponent {
    return AsComponent(ComboBoxEx_GetOwner(c.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (c *TComboBoxEx) Name() string {
    return ComboBoxEx_GetName(c.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (c *TComboBoxEx) SetName(value string) {
    ComboBoxEx_SetName(c.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TComboBoxEx) Tag() int {
    return ComboBoxEx_GetTag(c.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TComboBoxEx) SetTag(value int) {
    ComboBoxEx_SetTag(c.instance, value)
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (c *TComboBoxEx) DockClients(Index int32) *TControl {
    return AsControl(ComboBoxEx_GetDockClients(c.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (c *TComboBoxEx) Controls(Index int32) *TControl {
    return AsControl(ComboBoxEx_GetControls(c.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TComboBoxEx) Components(AIndex int32) *TComponent {
    return AsComponent(ComboBoxEx_GetComponents(c.instance, AIndex))
}

