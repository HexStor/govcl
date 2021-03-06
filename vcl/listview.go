
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

type TListView struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListView(owner IComponent) *TListView {
    l := new(TListView)
    l.instance = ListView_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsListView(obj interface{}) *TListView {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TListView{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsListView.
func ListViewFromInst(inst uintptr) *TListView {
    return AsListView(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsListView.
func ListViewFromObj(obj IObject) *TListView {
    return AsListView(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsListView.
func ListViewFromUnsafePointer(ptr unsafe.Pointer) *TListView {
    return AsListView(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (l *TListView) Free() {
    if l.instance != 0 {
        ListView_Free(l.instance)
        l.instance, l.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListView) Instance() uintptr {
    return l.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListView) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListView) IsValid() bool {
    return l.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (l *TListView) Is() TIs {
    return TIs(l.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (l *TListView) As() TAs {
//    return TAs(l.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListViewClass() TClass {
    return ListView_StaticClassType()
}

func (l *TListView) AddItem(Item string, AObject IObject) {
    ListView_AddItem(l.instance, Item , CheckPtr(AObject))
}

func (l *TListView) AlphaSort() bool {
    return ListView_AlphaSort(l.instance)
}

// CN: 清除。
// EN: .
func (l *TListView) Clear() {
    ListView_Clear(l.instance)
}

// CN: 清除选择。
// EN: .
func (l *TListView) ClearSelection() {
    ListView_ClearSelection(l.instance)
}

func (l *TListView) DeleteSelected() {
    ListView_DeleteSelected(l.instance)
}

func (l *TListView) IsEditing() bool {
    return ListView_IsEditing(l.instance)
}

// CN: 全选。
// EN: .
func (l *TListView) SelectAll() {
    ListView_SelectAll(l.instance)
}

func (l *TListView) CustomSort(SortProc PFNLVCOMPARE, lParam int) bool {
    return ListView_CustomSort(l.instance, SortProc , lParam)
}

// CN: 是否可以获得焦点。
// EN: .
func (l *TListView) CanFocus() bool {
    return ListView_CanFocus(l.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (l *TListView) ContainsControl(Control IControl) bool {
    return ListView_ContainsControl(l.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (l *TListView) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ListView_ControlAtPos(l.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (l *TListView) DisableAlign() {
    ListView_DisableAlign(l.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (l *TListView) EnableAlign() {
    ListView_EnableAlign(l.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (l *TListView) FindChildControl(ControlName string) *TControl {
    return AsControl(ListView_FindChildControl(l.instance, ControlName))
}

func (l *TListView) FlipChildren(AllLevels bool) {
    ListView_FlipChildren(l.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (l *TListView) Focused() bool {
    return ListView_Focused(l.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (l *TListView) HandleAllocated() bool {
    return ListView_HandleAllocated(l.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (l *TListView) InsertControl(AControl IControl) {
    ListView_InsertControl(l.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (l *TListView) Invalidate() {
    ListView_Invalidate(l.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (l *TListView) RemoveControl(AControl IControl) {
    ListView_RemoveControl(l.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (l *TListView) Realign() {
    ListView_Realign(l.instance)
}

// CN: 重绘。
// EN: Repaint.
func (l *TListView) Repaint() {
    ListView_Repaint(l.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (l *TListView) ScaleBy(M int32, D int32) {
    ListView_ScaleBy(l.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (l *TListView) ScrollBy(DeltaX int32, DeltaY int32) {
    ListView_ScrollBy(l.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (l *TListView) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ListView_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (l *TListView) SetFocus() {
    ListView_SetFocus(l.instance)
}

// CN: 控件更新。
// EN: Update.
func (l *TListView) Update() {
    ListView_Update(l.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (l *TListView) BringToFront() {
    ListView_BringToFront(l.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (l *TListView) ClientToScreen(Point TPoint) TPoint {
    return ListView_ClientToScreen(l.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (l *TListView) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ListView_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (l *TListView) Dragging() bool {
    return ListView_Dragging(l.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (l *TListView) HasParent() bool {
    return ListView_HasParent(l.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (l *TListView) Hide() {
    ListView_Hide(l.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (l *TListView) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ListView_Perform(l.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (l *TListView) Refresh() {
    ListView_Refresh(l.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (l *TListView) ScreenToClient(Point TPoint) TPoint {
    return ListView_ScreenToClient(l.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (l *TListView) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ListView_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (l *TListView) SendToBack() {
    ListView_SendToBack(l.instance)
}

// CN: 显示控件。
// EN: Show control.
func (l *TListView) Show() {
    ListView_Show(l.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (l *TListView) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ListView_GetTextBuf(l.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (l *TListView) GetTextLen() int32 {
    return ListView_GetTextLen(l.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (l *TListView) SetTextBuf(Buffer string) {
    ListView_SetTextBuf(l.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (l *TListView) FindComponent(AName string) *TComponent {
    return AsComponent(ListView_FindComponent(l.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TListView) GetNamePath() string {
    return ListView_GetNamePath(l.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TListView) Assign(Source IObject) {
    ListView_Assign(l.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListView) ClassType() TClass {
    return ListView_ClassType(l.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListView) ClassName() string {
    return ListView_ClassName(l.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListView) InstanceSize() int32 {
    return ListView_InstanceSize(l.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListView) InheritsFrom(AClass TClass) bool {
    return ListView_InheritsFrom(l.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListView) Equals(Obj IObject) bool {
    return ListView_Equals(l.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListView) GetHashCode() int32 {
    return ListView_GetHashCode(l.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (l *TListView) ToString() string {
    return ListView_ToString(l.instance)
}

func (l *TListView) Action() *TAction {
    return AsAction(ListView_GetAction(l.instance))
}

func (l *TListView) SetAction(value IComponent) {
    ListView_SetAction(l.instance, CheckPtr(value))
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (l *TListView) Align() TAlign {
    return ListView_GetAlign(l.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (l *TListView) SetAlign(value TAlign) {
    ListView_SetAlign(l.instance, value)
}

func (l *TListView) AllocBy() int32 {
    return ListView_GetAllocBy(l.instance)
}

func (l *TListView) SetAllocBy(value int32) {
    ListView_SetAllocBy(l.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (l *TListView) Anchors() TAnchors {
    return ListView_GetAnchors(l.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (l *TListView) SetAnchors(value TAnchors) {
    ListView_SetAnchors(l.instance, value)
}

func (l *TListView) BiDiMode() TBiDiMode {
    return ListView_GetBiDiMode(l.instance)
}

func (l *TListView) SetBiDiMode(value TBiDiMode) {
    ListView_SetBiDiMode(l.instance, value)
}

// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (l *TListView) BorderStyle() TBorderStyle {
    return ListView_GetBorderStyle(l.instance)
}

// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (l *TListView) SetBorderStyle(value TBorderStyle) {
    ListView_SetBorderStyle(l.instance, value)
}

// CN: 获取边框的宽度。
// EN: .
func (l *TListView) BorderWidth() int32 {
    return ListView_GetBorderWidth(l.instance)
}

// CN: 设置边框的宽度。
// EN: .
func (l *TListView) SetBorderWidth(value int32) {
    ListView_SetBorderWidth(l.instance, value)
}

func (l *TListView) Checkboxes() bool {
    return ListView_GetCheckboxes(l.instance)
}

func (l *TListView) SetCheckboxes(value bool) {
    ListView_SetCheckboxes(l.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (l *TListView) Color() TColor {
    return ListView_GetColor(l.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (l *TListView) SetColor(value TColor) {
    ListView_SetColor(l.instance, value)
}

func (l *TListView) Columns() *TListColumns {
    return AsListColumns(ListView_GetColumns(l.instance))
}

func (l *TListView) SetColumns(value *TListColumns) {
    ListView_SetColumns(l.instance, CheckPtr(value))
}

func (l *TListView) ColumnClick() bool {
    return ListView_GetColumnClick(l.instance)
}

func (l *TListView) SetColumnClick(value bool) {
    ListView_SetColumnClick(l.instance, value)
}

func (l *TListView) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ListView_GetConstraints(l.instance))
}

func (l *TListView) SetConstraints(value *TSizeConstraints) {
    ListView_SetConstraints(l.instance, CheckPtr(value))
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (l *TListView) DoubleBuffered() bool {
    return ListView_GetDoubleBuffered(l.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (l *TListView) SetDoubleBuffered(value bool) {
    ListView_SetDoubleBuffered(l.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (l *TListView) DragCursor() TCursor {
    return ListView_GetDragCursor(l.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (l *TListView) SetDragCursor(value TCursor) {
    ListView_SetDragCursor(l.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (l *TListView) DragKind() TDragKind {
    return ListView_GetDragKind(l.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (l *TListView) SetDragKind(value TDragKind) {
    ListView_SetDragKind(l.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (l *TListView) DragMode() TDragMode {
    return ListView_GetDragMode(l.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (l *TListView) SetDragMode(value TDragMode) {
    ListView_SetDragMode(l.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (l *TListView) Enabled() bool {
    return ListView_GetEnabled(l.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (l *TListView) SetEnabled(value bool) {
    ListView_SetEnabled(l.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (l *TListView) Font() *TFont {
    return AsFont(ListView_GetFont(l.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (l *TListView) SetFont(value *TFont) {
    ListView_SetFont(l.instance, CheckPtr(value))
}

func (l *TListView) FlatScrollBars() bool {
    return ListView_GetFlatScrollBars(l.instance)
}

func (l *TListView) SetFlatScrollBars(value bool) {
    ListView_SetFlatScrollBars(l.instance, value)
}

func (l *TListView) FullDrag() bool {
    return ListView_GetFullDrag(l.instance)
}

func (l *TListView) SetFullDrag(value bool) {
    ListView_SetFullDrag(l.instance, value)
}

func (l *TListView) GridLines() bool {
    return ListView_GetGridLines(l.instance)
}

func (l *TListView) SetGridLines(value bool) {
    ListView_SetGridLines(l.instance, value)
}

// CN: 获取隐藏选择。
// EN: .
func (l *TListView) HideSelection() bool {
    return ListView_GetHideSelection(l.instance)
}

// CN: 设置隐藏选择。
// EN: .
func (l *TListView) SetHideSelection(value bool) {
    ListView_SetHideSelection(l.instance, value)
}

func (l *TListView) HotTrack() bool {
    return ListView_GetHotTrack(l.instance)
}

func (l *TListView) SetHotTrack(value bool) {
    ListView_SetHotTrack(l.instance, value)
}

func (l *TListView) IconOptions() *TIconOptions {
    return AsIconOptions(ListView_GetIconOptions(l.instance))
}

func (l *TListView) SetIconOptions(value IObject) {
    ListView_SetIconOptions(l.instance, CheckPtr(value))
}

func (l *TListView) Items() *TListItems {
    return AsListItems(ListView_GetItems(l.instance))
}

func (l *TListView) SetItems(value *TListItems) {
    ListView_SetItems(l.instance, CheckPtr(value))
}

func (l *TListView) LargeImages() *TImageList {
    return AsImageList(ListView_GetLargeImages(l.instance))
}

func (l *TListView) SetLargeImages(value IComponent) {
    ListView_SetLargeImages(l.instance, CheckPtr(value))
}

func (l *TListView) MultiSelect() bool {
    return ListView_GetMultiSelect(l.instance)
}

func (l *TListView) SetMultiSelect(value bool) {
    ListView_SetMultiSelect(l.instance, value)
}

func (l *TListView) OwnerData() bool {
    return ListView_GetOwnerData(l.instance)
}

func (l *TListView) SetOwnerData(value bool) {
    ListView_SetOwnerData(l.instance, value)
}

func (l *TListView) OwnerDraw() bool {
    return ListView_GetOwnerDraw(l.instance)
}

func (l *TListView) SetOwnerDraw(value bool) {
    ListView_SetOwnerDraw(l.instance, value)
}

// CN: 获取只读。
// EN: .
func (l *TListView) ReadOnly() bool {
    return ListView_GetReadOnly(l.instance)
}

// CN: 设置只读。
// EN: .
func (l *TListView) SetReadOnly(value bool) {
    ListView_SetReadOnly(l.instance, value)
}

func (l *TListView) RowSelect() bool {
    return ListView_GetRowSelect(l.instance)
}

func (l *TListView) SetRowSelect(value bool) {
    ListView_SetRowSelect(l.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (l *TListView) ParentColor() bool {
    return ListView_GetParentColor(l.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (l *TListView) SetParentColor(value bool) {
    ListView_SetParentColor(l.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (l *TListView) ParentDoubleBuffered() bool {
    return ListView_GetParentDoubleBuffered(l.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (l *TListView) SetParentDoubleBuffered(value bool) {
    ListView_SetParentDoubleBuffered(l.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (l *TListView) ParentFont() bool {
    return ListView_GetParentFont(l.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (l *TListView) SetParentFont(value bool) {
    ListView_SetParentFont(l.instance, value)
}

func (l *TListView) ParentShowHint() bool {
    return ListView_GetParentShowHint(l.instance)
}

func (l *TListView) SetParentShowHint(value bool) {
    ListView_SetParentShowHint(l.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (l *TListView) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ListView_GetPopupMenu(l.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (l *TListView) SetPopupMenu(value IComponent) {
    ListView_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TListView) ShowColumnHeaders() bool {
    return ListView_GetShowColumnHeaders(l.instance)
}

func (l *TListView) SetShowColumnHeaders(value bool) {
    ListView_SetShowColumnHeaders(l.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (l *TListView) ShowHint() bool {
    return ListView_GetShowHint(l.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (l *TListView) SetShowHint(value bool) {
    ListView_SetShowHint(l.instance, value)
}

func (l *TListView) SmallImages() *TImageList {
    return AsImageList(ListView_GetSmallImages(l.instance))
}

func (l *TListView) SetSmallImages(value IComponent) {
    ListView_SetSmallImages(l.instance, CheckPtr(value))
}

func (l *TListView) SortType() TSortType {
    return ListView_GetSortType(l.instance)
}

func (l *TListView) SetSortType(value TSortType) {
    ListView_SetSortType(l.instance, value)
}

func (l *TListView) StateImages() *TImageList {
    return AsImageList(ListView_GetStateImages(l.instance))
}

func (l *TListView) SetStateImages(value IComponent) {
    ListView_SetStateImages(l.instance, CheckPtr(value))
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (l *TListView) TabOrder() TTabOrder {
    return ListView_GetTabOrder(l.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (l *TListView) SetTabOrder(value TTabOrder) {
    ListView_SetTabOrder(l.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (l *TListView) TabStop() bool {
    return ListView_GetTabStop(l.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (l *TListView) SetTabStop(value bool) {
    ListView_SetTabStop(l.instance, value)
}

func (l *TListView) ViewStyle() TViewStyle {
    return ListView_GetViewStyle(l.instance)
}

func (l *TListView) SetViewStyle(value TViewStyle) {
    ListView_SetViewStyle(l.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (l *TListView) Visible() bool {
    return ListView_GetVisible(l.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (l *TListView) SetVisible(value bool) {
    ListView_SetVisible(l.instance, value)
}

func (l *TListView) SetOnAdvancedCustomDraw(fn TLVAdvancedCustomDrawEvent) {
    ListView_SetOnAdvancedCustomDraw(l.instance, fn)
}

func (l *TListView) SetOnAdvancedCustomDrawItem(fn TLVAdvancedCustomDrawItemEvent) {
    ListView_SetOnAdvancedCustomDrawItem(l.instance, fn)
}

func (l *TListView) SetOnAdvancedCustomDrawSubItem(fn TLVAdvancedCustomDrawSubItemEvent) {
    ListView_SetOnAdvancedCustomDrawSubItem(l.instance, fn)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (l *TListView) SetOnChange(fn TLVChangeEvent) {
    ListView_SetOnChange(l.instance, fn)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (l *TListView) SetOnClick(fn TNotifyEvent) {
    ListView_SetOnClick(l.instance, fn)
}

func (l *TListView) SetOnColumnClick(fn TLVColumnClickEvent) {
    ListView_SetOnColumnClick(l.instance, fn)
}

func (l *TListView) SetOnCompare(fn TLVCompareEvent) {
    ListView_SetOnCompare(l.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (l *TListView) SetOnContextPopup(fn TContextPopupEvent) {
    ListView_SetOnContextPopup(l.instance, fn)
}

func (l *TListView) SetOnCustomDraw(fn TLVCustomDrawEvent) {
    ListView_SetOnCustomDraw(l.instance, fn)
}

func (l *TListView) SetOnCustomDrawItem(fn TLVCustomDrawItemEvent) {
    ListView_SetOnCustomDrawItem(l.instance, fn)
}

func (l *TListView) SetOnCustomDrawSubItem(fn TLVCustomDrawSubItemEvent) {
    ListView_SetOnCustomDrawSubItem(l.instance, fn)
}

func (l *TListView) SetOnData(fn TLVOwnerDataEvent) {
    ListView_SetOnData(l.instance, fn)
}

func (l *TListView) SetOnDataFind(fn TLVOwnerDataFindEvent) {
    ListView_SetOnDataFind(l.instance, fn)
}

func (l *TListView) SetOnDataHint(fn TLVOwnerDataHintEvent) {
    ListView_SetOnDataHint(l.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (l *TListView) SetOnDblClick(fn TNotifyEvent) {
    ListView_SetOnDblClick(l.instance, fn)
}

func (l *TListView) SetOnDeletion(fn TLVDeletedEvent) {
    ListView_SetOnDeletion(l.instance, fn)
}

func (l *TListView) SetOnEdited(fn TLVEditedEvent) {
    ListView_SetOnEdited(l.instance, fn)
}

func (l *TListView) SetOnEditing(fn TLVEditingEvent) {
    ListView_SetOnEditing(l.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (l *TListView) SetOnEndDock(fn TEndDragEvent) {
    ListView_SetOnEndDock(l.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (l *TListView) SetOnEndDrag(fn TEndDragEvent) {
    ListView_SetOnEndDrag(l.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (l *TListView) SetOnEnter(fn TNotifyEvent) {
    ListView_SetOnEnter(l.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (l *TListView) SetOnExit(fn TNotifyEvent) {
    ListView_SetOnExit(l.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (l *TListView) SetOnDragDrop(fn TDragDropEvent) {
    ListView_SetOnDragDrop(l.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (l *TListView) SetOnDragOver(fn TDragOverEvent) {
    ListView_SetOnDragOver(l.instance, fn)
}

func (l *TListView) SetOnInsert(fn TLVDeletedEvent) {
    ListView_SetOnInsert(l.instance, fn)
}

// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (l *TListView) SetOnKeyDown(fn TKeyEvent) {
    ListView_SetOnKeyDown(l.instance, fn)
}

func (l *TListView) SetOnKeyPress(fn TKeyPressEvent) {
    ListView_SetOnKeyPress(l.instance, fn)
}

// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (l *TListView) SetOnKeyUp(fn TKeyEvent) {
    ListView_SetOnKeyUp(l.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (l *TListView) SetOnMouseDown(fn TMouseEvent) {
    ListView_SetOnMouseDown(l.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (l *TListView) SetOnMouseEnter(fn TNotifyEvent) {
    ListView_SetOnMouseEnter(l.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (l *TListView) SetOnMouseLeave(fn TNotifyEvent) {
    ListView_SetOnMouseLeave(l.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (l *TListView) SetOnMouseMove(fn TMouseMoveEvent) {
    ListView_SetOnMouseMove(l.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (l *TListView) SetOnMouseUp(fn TMouseEvent) {
    ListView_SetOnMouseUp(l.instance, fn)
}

// CN: 设置大小被改变事件。
// EN: .
func (l *TListView) SetOnResize(fn TNotifyEvent) {
    ListView_SetOnResize(l.instance, fn)
}

func (l *TListView) SetOnSelectItem(fn TLVSelectItemEvent) {
    ListView_SetOnSelectItem(l.instance, fn)
}

func (l *TListView) SetOnItemChecked(fn TLVCheckedItemEvent) {
    ListView_SetOnItemChecked(l.instance, fn)
}

// CN: 设置启动停靠。
// EN: .
func (l *TListView) SetOnStartDock(fn TStartDockEvent) {
    ListView_SetOnStartDock(l.instance, fn)
}

// CN: 获取画布。
// EN: .
func (l *TListView) Canvas() *TCanvas {
    return AsCanvas(ListView_GetCanvas(l.instance))
}

func (l *TListView) ItemFocused() *TListItem {
    return AsListItem(ListView_GetItemFocused(l.instance))
}

func (l *TListView) SetItemFocused(value *TListItem) {
    ListView_SetItemFocused(l.instance, CheckPtr(value))
}

func (l *TListView) SelCount() int32 {
    return ListView_GetSelCount(l.instance)
}

func (l *TListView) Selected() *TListItem {
    return AsListItem(ListView_GetSelected(l.instance))
}

func (l *TListView) SetSelected(value *TListItem) {
    ListView_SetSelected(l.instance, CheckPtr(value))
}

func (l *TListView) TopItem() *TListItem {
    return AsListItem(ListView_GetTopItem(l.instance))
}

func (l *TListView) VisibleRowCount() int32 {
    return ListView_GetVisibleRowCount(l.instance)
}

func (l *TListView) ItemIndex() int32 {
    return ListView_GetItemIndex(l.instance)
}

func (l *TListView) SetItemIndex(value int32) {
    ListView_SetItemIndex(l.instance, value)
}

// CN: 获取依靠客户端总数。
// EN: .
func (l *TListView) DockClientCount() int32 {
    return ListView_GetDockClientCount(l.instance)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (l *TListView) DockSite() bool {
    return ListView_GetDockSite(l.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (l *TListView) SetDockSite(value bool) {
    ListView_SetDockSite(l.instance, value)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (l *TListView) MouseInClient() bool {
    return ListView_GetMouseInClient(l.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (l *TListView) VisibleDockClientCount() int32 {
    return ListView_GetVisibleDockClientCount(l.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (l *TListView) Brush() *TBrush {
    return AsBrush(ListView_GetBrush(l.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (l *TListView) ControlCount() int32 {
    return ListView_GetControlCount(l.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (l *TListView) Handle() HWND {
    return ListView_GetHandle(l.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (l *TListView) ParentWindow() HWND {
    return ListView_GetParentWindow(l.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (l *TListView) SetParentWindow(value HWND) {
    ListView_SetParentWindow(l.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (l *TListView) UseDockManager() bool {
    return ListView_GetUseDockManager(l.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (l *TListView) SetUseDockManager(value bool) {
    ListView_SetUseDockManager(l.instance, value)
}

func (l *TListView) BoundsRect() TRect {
    return ListView_GetBoundsRect(l.instance)
}

func (l *TListView) SetBoundsRect(value TRect) {
    ListView_SetBoundsRect(l.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (l *TListView) ClientHeight() int32 {
    return ListView_GetClientHeight(l.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (l *TListView) SetClientHeight(value int32) {
    ListView_SetClientHeight(l.instance, value)
}

func (l *TListView) ClientOrigin() TPoint {
    return ListView_GetClientOrigin(l.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (l *TListView) ClientRect() TRect {
    return ListView_GetClientRect(l.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (l *TListView) ClientWidth() int32 {
    return ListView_GetClientWidth(l.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (l *TListView) SetClientWidth(value int32) {
    ListView_SetClientWidth(l.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (l *TListView) ControlState() TControlState {
    return ListView_GetControlState(l.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (l *TListView) SetControlState(value TControlState) {
    ListView_SetControlState(l.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (l *TListView) ControlStyle() TControlStyle {
    return ListView_GetControlStyle(l.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (l *TListView) SetControlStyle(value TControlStyle) {
    ListView_SetControlStyle(l.instance, value)
}

func (l *TListView) Floating() bool {
    return ListView_GetFloating(l.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (l *TListView) Parent() *TWinControl {
    return AsWinControl(ListView_GetParent(l.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (l *TListView) SetParent(value IWinControl) {
    ListView_SetParent(l.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (l *TListView) Left() int32 {
    return ListView_GetLeft(l.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (l *TListView) SetLeft(value int32) {
    ListView_SetLeft(l.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (l *TListView) Top() int32 {
    return ListView_GetTop(l.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (l *TListView) SetTop(value int32) {
    ListView_SetTop(l.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (l *TListView) Width() int32 {
    return ListView_GetWidth(l.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (l *TListView) SetWidth(value int32) {
    ListView_SetWidth(l.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (l *TListView) Height() int32 {
    return ListView_GetHeight(l.instance)
}

// CN: 设置高度。
// EN: Set height.
func (l *TListView) SetHeight(value int32) {
    ListView_SetHeight(l.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (l *TListView) Cursor() TCursor {
    return ListView_GetCursor(l.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (l *TListView) SetCursor(value TCursor) {
    ListView_SetCursor(l.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (l *TListView) Hint() string {
    return ListView_GetHint(l.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (l *TListView) SetHint(value string) {
    ListView_SetHint(l.instance, value)
}

// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (l *TListView) Margins() *TMargins {
    return AsMargins(ListView_GetMargins(l.instance))
}

// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (l *TListView) SetMargins(value *TMargins) {
    ListView_SetMargins(l.instance, CheckPtr(value))
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (l *TListView) ComponentCount() int32 {
    return ListView_GetComponentCount(l.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (l *TListView) ComponentIndex() int32 {
    return ListView_GetComponentIndex(l.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (l *TListView) SetComponentIndex(value int32) {
    ListView_SetComponentIndex(l.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TListView) Owner() *TComponent {
    return AsComponent(ListView_GetOwner(l.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (l *TListView) Name() string {
    return ListView_GetName(l.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (l *TListView) SetName(value string) {
    ListView_SetName(l.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TListView) Tag() int {
    return ListView_GetTag(l.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TListView) SetTag(value int) {
    ListView_SetTag(l.instance, value)
}

func (l *TListView) Column(Index int32) *TListColumn {
    return AsListColumn(ListView_GetColumn(l.instance, Index))
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (l *TListView) DockClients(Index int32) *TControl {
    return AsControl(ListView_GetDockClients(l.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (l *TListView) Controls(Index int32) *TControl {
    return AsControl(ListView_GetControls(l.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TListView) Components(AIndex int32) *TComponent {
    return AsComponent(ListView_GetComponents(l.instance, AIndex))
}

