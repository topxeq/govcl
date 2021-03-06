
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

function TabSheet_Create(AOwner: TComponent): TTabSheet; extdecl;
begin
  Result :=  TTabSheet.Create(AOwner);
end;

procedure TabSheet_Free(AObj: TTabSheet); extdecl;
begin
  AObj.Free;
end;

function TabSheet_CanFocus(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

function TabSheet_ContainsControl(AObj: TTabSheet; Control: TControl): LongBool; extdecl;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function TabSheet_ControlAtPos(AObj: TTabSheet; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool): TControl; extdecl;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls);
end;

procedure TabSheet_DisableAlign(AObj: TTabSheet); extdecl;
begin
  AObj.DisableAlign;
end;

procedure TabSheet_EnableAlign(AObj: TTabSheet); extdecl;
begin
  AObj.EnableAlign;
end;

function TabSheet_FindChildControl(AObj: TTabSheet; ControlName: PChar): TControl; extdecl;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure TabSheet_FlipChildren(AObj: TTabSheet; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function TabSheet_Focused(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

function TabSheet_HandleAllocated(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure TabSheet_InsertControl(AObj: TTabSheet; AControl: TControl); extdecl;
begin
  AObj.InsertControl(AControl);
end;

procedure TabSheet_Invalidate(AObj: TTabSheet); extdecl;
begin
  AObj.Invalidate;
end;

procedure TabSheet_RemoveControl(AObj: TTabSheet; AControl: TControl); extdecl;
begin
  AObj.RemoveControl(AControl);
end;

procedure TabSheet_Realign(AObj: TTabSheet); extdecl;
begin
  AObj.Realign;
end;

procedure TabSheet_Repaint(AObj: TTabSheet); extdecl;
begin
  AObj.Repaint;
end;

procedure TabSheet_ScaleBy(AObj: TTabSheet; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure TabSheet_ScrollBy(AObj: TTabSheet; DeltaX: Integer; DeltaY: Integer); extdecl;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure TabSheet_SetBounds(AObj: TTabSheet; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure TabSheet_SetFocus(AObj: TTabSheet); extdecl;
begin
  AObj.SetFocus;
end;

procedure TabSheet_Update(AObj: TTabSheet); extdecl;
begin
  AObj.Update;
end;

procedure TabSheet_BringToFront(AObj: TTabSheet); extdecl;
begin
  AObj.BringToFront;
end;

procedure TabSheet_ClientToScreen(AObj: TTabSheet; var Point: TPoint; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure TabSheet_ClientToParent(AObj: TTabSheet; var Point: TPoint; AParent: TWinControl; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function TabSheet_Dragging(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.Dragging;
end;

function TabSheet_HasParent(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure TabSheet_Hide(AObj: TTabSheet); extdecl;
begin
  AObj.Hide;
end;

function TabSheet_Perform(AObj: TTabSheet; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure TabSheet_Refresh(AObj: TTabSheet); extdecl;
begin
  AObj.Refresh;
end;

procedure TabSheet_ScreenToClient(AObj: TTabSheet; var Point: TPoint; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure TabSheet_ParentToClient(AObj: TTabSheet; var Point: TPoint; AParent: TWinControl; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure TabSheet_SendToBack(AObj: TTabSheet); extdecl;
begin
  AObj.SendToBack;
end;

procedure TabSheet_Show(AObj: TTabSheet); extdecl;
begin
  AObj.Show;
end;

function TabSheet_GetTextBuf(AObj: TTabSheet; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function TabSheet_GetTextLen(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  {$ifdef darwin}Length(AObj.Caption){$else}AObj.GetTextLen{$endif};
end;

procedure TabSheet_SetTextBuf(AObj: TTabSheet; Buffer: PChar); extdecl;
begin
  AObj.SetTextBuf(Buffer);
end;

function TabSheet_FindComponent(AObj: TTabSheet; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function TabSheet_GetNamePath(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TabSheet_Assign(AObj: TTabSheet; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function TabSheet_ClassType(AObj: TTabSheet): TClass; extdecl;
begin
  Result :=  AObj.ClassType;
end;

function TabSheet_ClassName(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function TabSheet_InstanceSize(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.InstanceSize;
end;

function TabSheet_InheritsFrom(AObj: TTabSheet; AClass: TClass): LongBool; extdecl;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function TabSheet_Equals(AObj: TTabSheet; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TabSheet_GetHashCode(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TabSheet_ToString(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

procedure TabSheet_AnchorToNeighbour(AObj: TTabSheet; ASide: TAnchorKind; ASpace: Integer; ASibling: TControl); extdecl;
begin
  AObj.AnchorToNeighbour(ASide, ASpace, ASibling);
end;

procedure TabSheet_AnchorParallel(AObj: TTabSheet; ASide: TAnchorKind; ASpace: Integer; ASibling: TControl); extdecl;
begin
  AObj.AnchorParallel(ASide, ASpace, ASibling);
end;

procedure TabSheet_AnchorHorizontalCenterTo(AObj: TTabSheet; ASibling: TControl); extdecl;
begin
  AObj.AnchorHorizontalCenterTo(ASibling);
end;

procedure TabSheet_AnchorVerticalCenterTo(AObj: TTabSheet; ASibling: TControl); extdecl;
begin
  AObj.AnchorVerticalCenterTo(ASibling);
end;

procedure TabSheet_AnchorAsAlign(AObj: TTabSheet; ATheAlign: TAlign; ASpace: Integer); extdecl;
begin
  AObj.AnchorAsAlign(ATheAlign, ASpace);
end;

procedure TabSheet_AnchorClient(AObj: TTabSheet; ASpace: Integer); extdecl;
begin
  AObj.AnchorClient(ASpace);
end;

function TabSheet_GetPageControl(AObj: TTabSheet): TPageControl; extdecl;
begin
  Result :=  AObj.PageControl;
end;

procedure TabSheet_SetPageControl(AObj: TTabSheet; AValue: TPageControl); extdecl;
begin
  AObj.PageControl := AValue;
end;

function TabSheet_GetTabIndex(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.TabIndex;
end;

function TabSheet_GetBorderWidth(AObj: TTabSheet): TBorderWidth; extdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure TabSheet_SetBorderWidth(AObj: TTabSheet; AValue: TBorderWidth); extdecl;
begin
  AObj.BorderWidth := AValue;
end;

function TabSheet_GetCaption(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure TabSheet_SetCaption(AObj: TTabSheet; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function TabSheet_GetDoubleBuffered(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure TabSheet_SetDoubleBuffered(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function TabSheet_GetEnabled(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure TabSheet_SetEnabled(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function TabSheet_GetFont(AObj: TTabSheet): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure TabSheet_SetFont(AObj: TTabSheet; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function TabSheet_GetHeight(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure TabSheet_SetHeight(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function TabSheet_GetImageIndex(AObj: TTabSheet): TImageIndex; extdecl;
begin
  Result :=  AObj.ImageIndex;
end;

procedure TabSheet_SetImageIndex(AObj: TTabSheet; AValue: TImageIndex); extdecl;
begin
  AObj.ImageIndex := AValue;
end;

function TabSheet_GetLeft(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure TabSheet_SetLeft(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function TabSheet_GetConstraints(AObj: TTabSheet): TSizeConstraints; extdecl;
begin
  Result :=  AObj.Constraints;
end;

procedure TabSheet_SetConstraints(AObj: TTabSheet; AValue: TSizeConstraints); extdecl;
begin
  AObj.Constraints := AValue;
end;

function TabSheet_GetPageIndex(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.PageIndex;
end;

procedure TabSheet_SetPageIndex(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.PageIndex := AValue;
end;

function TabSheet_GetParentDoubleBuffered(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure TabSheet_SetParentDoubleBuffered(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function TabSheet_GetParentFont(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure TabSheet_SetParentFont(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function TabSheet_GetParentShowHint(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure TabSheet_SetParentShowHint(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function TabSheet_GetPopupMenu(AObj: TTabSheet): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure TabSheet_SetPopupMenu(AObj: TTabSheet; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function TabSheet_GetShowHint(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure TabSheet_SetShowHint(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function TabSheet_GetTabVisible(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.TabVisible;
end;

procedure TabSheet_SetTabVisible(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.TabVisible := AValue;
end;

function TabSheet_GetTop(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure TabSheet_SetTop(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function TabSheet_GetVisible(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure TabSheet_SetVisible(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function TabSheet_GetWidth(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure TabSheet_SetWidth(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

procedure TabSheet_SetOnContextPopup(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnContextPopup}
  {$DEFINE EventMethodName := OnContextPopup}
  EventMethodCode
end;

procedure TabSheet_SetOnDragDrop(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnDragDrop}
  {$DEFINE EventMethodName := OnDragDrop}
  EventMethodCode
end;

procedure TabSheet_SetOnDragOver(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnDragOver}
  {$DEFINE EventMethodName := OnDragOver}
  EventMethodCode
end;

procedure TabSheet_SetOnEndDrag(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnEndDrag}
  {$DEFINE EventMethodName := OnEndDrag}
  EventMethodCode
end;

procedure TabSheet_SetOnEnter(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnEnter}
  {$DEFINE EventMethodName := OnEnter}
  EventMethodCode
end;

procedure TabSheet_SetOnExit(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnExit}
  {$DEFINE EventMethodName := OnExit}
  EventMethodCode
end;

procedure TabSheet_SetOnHide(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnHide}
  {$DEFINE EventMethodName := OnHide}
  EventMethodCode
end;

procedure TabSheet_SetOnMouseDown(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnMouseDown}
  {$DEFINE EventMethodName := OnMouseDown}
  EventMethodCode
end;

procedure TabSheet_SetOnMouseEnter(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnMouseEnter}
  {$DEFINE EventMethodName := OnMouseEnter}
  EventMethodCode
end;

procedure TabSheet_SetOnMouseLeave(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnMouseLeave}
  {$DEFINE EventMethodName := OnMouseLeave}
  EventMethodCode
end;

procedure TabSheet_SetOnMouseMove(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnMouseMove}
  {$DEFINE EventMethodName := OnMouseMove}
  EventMethodCode
end;

procedure TabSheet_SetOnMouseUp(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnMouseUp}
  {$DEFINE EventMethodName := OnMouseUp}
  EventMethodCode
end;

procedure TabSheet_SetOnResize(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnResize}
  {$DEFINE EventMethodName := OnResize}
  EventMethodCode
end;

procedure TabSheet_SetOnShow(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnShow}
  {$DEFINE EventMethodName := OnShow}
  EventMethodCode
end;

function TabSheet_GetDockClientCount(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.DockClientCount;
end;

function TabSheet_GetDockSite(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.DockSite;
end;

procedure TabSheet_SetDockSite(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.DockSite := AValue;
end;

function TabSheet_GetMouseInClient(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.MouseInClient;
end;

function TabSheet_GetVisibleDockClientCount(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function TabSheet_GetBrush(AObj: TTabSheet): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function TabSheet_GetControlCount(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function TabSheet_GetHandle(AObj: TTabSheet): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function TabSheet_GetParentWindow(AObj: TTabSheet): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure TabSheet_SetParentWindow(AObj: TTabSheet; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function TabSheet_GetTabOrder(AObj: TTabSheet): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure TabSheet_SetTabOrder(AObj: TTabSheet; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function TabSheet_GetTabStop(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure TabSheet_SetTabStop(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function TabSheet_GetUseDockManager(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.UseDockManager;
end;

procedure TabSheet_SetUseDockManager(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.UseDockManager := AValue;
end;

function TabSheet_GetAction(AObj: TTabSheet): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure TabSheet_SetAction(AObj: TTabSheet; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function TabSheet_GetAlign(AObj: TTabSheet): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure TabSheet_SetAlign(AObj: TTabSheet; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function TabSheet_GetAnchors(AObj: TTabSheet): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure TabSheet_SetAnchors(AObj: TTabSheet; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function TabSheet_GetBiDiMode(AObj: TTabSheet): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure TabSheet_SetBiDiMode(AObj: TTabSheet; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure TabSheet_GetBoundsRect(AObj: TTabSheet; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure TabSheet_SetBoundsRect(AObj: TTabSheet; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function TabSheet_GetClientHeight(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure TabSheet_SetClientHeight(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure TabSheet_GetClientOrigin(AObj: TTabSheet; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure TabSheet_GetClientRect(AObj: TTabSheet; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function TabSheet_GetClientWidth(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure TabSheet_SetClientWidth(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function TabSheet_GetControlState(AObj: TTabSheet): TControlState; extdecl;
begin
  Result :=  AObj.ControlState;
end;

procedure TabSheet_SetControlState(AObj: TTabSheet; AValue: TControlState); extdecl;
begin
  AObj.ControlState := AValue;
end;

function TabSheet_GetControlStyle(AObj: TTabSheet): TControlStyle; extdecl;
begin
  Result :=  AObj.ControlStyle;
end;

procedure TabSheet_SetControlStyle(AObj: TTabSheet; AValue: TControlStyle); extdecl;
begin
  AObj.ControlStyle := AValue;
end;

function TabSheet_GetFloating(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.Floating;
end;

function TabSheet_GetParent(AObj: TTabSheet): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure TabSheet_SetParent(AObj: TTabSheet; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function TabSheet_GetCursor(AObj: TTabSheet): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure TabSheet_SetCursor(AObj: TTabSheet; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function TabSheet_GetHint(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure TabSheet_SetHint(AObj: TTabSheet; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function TabSheet_GetComponentCount(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function TabSheet_GetComponentIndex(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure TabSheet_SetComponentIndex(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function TabSheet_GetOwner(AObj: TTabSheet): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function TabSheet_GetName(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TabSheet_SetName(AObj: TTabSheet; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function TabSheet_GetTag(AObj: TTabSheet): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure TabSheet_SetTag(AObj: TTabSheet; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function TabSheet_GetAnchorSideLeft(AObj: TTabSheet): TAnchorSide; extdecl;
begin
  Result :=  AObj.AnchorSideLeft;
end;

procedure TabSheet_SetAnchorSideLeft(AObj: TTabSheet; AValue: TAnchorSide); extdecl;
begin
  AObj.AnchorSideLeft := AValue;
end;

function TabSheet_GetAnchorSideTop(AObj: TTabSheet): TAnchorSide; extdecl;
begin
  Result :=  AObj.AnchorSideTop;
end;

procedure TabSheet_SetAnchorSideTop(AObj: TTabSheet; AValue: TAnchorSide); extdecl;
begin
  AObj.AnchorSideTop := AValue;
end;

function TabSheet_GetAnchorSideRight(AObj: TTabSheet): TAnchorSide; extdecl;
begin
  Result :=  AObj.AnchorSideRight;
end;

procedure TabSheet_SetAnchorSideRight(AObj: TTabSheet; AValue: TAnchorSide); extdecl;
begin
  AObj.AnchorSideRight := AValue;
end;

function TabSheet_GetAnchorSideBottom(AObj: TTabSheet): TAnchorSide; extdecl;
begin
  Result :=  AObj.AnchorSideBottom;
end;

procedure TabSheet_SetAnchorSideBottom(AObj: TTabSheet; AValue: TAnchorSide); extdecl;
begin
  AObj.AnchorSideBottom := AValue;
end;

function TabSheet_GetChildSizing(AObj: TTabSheet): TControlChildSizing; extdecl;
begin
  Result :=  AObj.ChildSizing;
end;

procedure TabSheet_SetChildSizing(AObj: TTabSheet; AValue: TControlChildSizing); extdecl;
begin
  AObj.ChildSizing := AValue;
end;

function TabSheet_GetBorderSpacing(AObj: TTabSheet): TControlBorderSpacing; extdecl;
begin
  Result :=  AObj.BorderSpacing;
end;

procedure TabSheet_SetBorderSpacing(AObj: TTabSheet; AValue: TControlBorderSpacing); extdecl;
begin
  AObj.BorderSpacing := AValue;
end;

function TabSheet_GetDockClients(AObj: TTabSheet; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.DockClients[Index];
end;

function TabSheet_GetControls(AObj: TTabSheet; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function TabSheet_GetComponents(AObj: TTabSheet; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;

function TabSheet_GetAnchorSide(AObj: TTabSheet; AKind: TAnchorKind): TAnchorSide; extdecl;
begin
  Result :=  AObj.AnchorSide[AKind];
end;

function TabSheet_StaticClassType: TClass; extdecl;
begin
  Result := TTabSheet;
end;

exports
  TabSheet_Create,
  TabSheet_Free,
  TabSheet_CanFocus,
  TabSheet_ContainsControl,
  TabSheet_ControlAtPos,
  TabSheet_DisableAlign,
  TabSheet_EnableAlign,
  TabSheet_FindChildControl,
  TabSheet_FlipChildren,
  TabSheet_Focused,
  TabSheet_HandleAllocated,
  TabSheet_InsertControl,
  TabSheet_Invalidate,
  TabSheet_RemoveControl,
  TabSheet_Realign,
  TabSheet_Repaint,
  TabSheet_ScaleBy,
  TabSheet_ScrollBy,
  TabSheet_SetBounds,
  TabSheet_SetFocus,
  TabSheet_Update,
  TabSheet_BringToFront,
  TabSheet_ClientToScreen,
  TabSheet_ClientToParent,
  TabSheet_Dragging,
  TabSheet_HasParent,
  TabSheet_Hide,
  TabSheet_Perform,
  TabSheet_Refresh,
  TabSheet_ScreenToClient,
  TabSheet_ParentToClient,
  TabSheet_SendToBack,
  TabSheet_Show,
  TabSheet_GetTextBuf,
  TabSheet_GetTextLen,
  TabSheet_SetTextBuf,
  TabSheet_FindComponent,
  TabSheet_GetNamePath,
  TabSheet_Assign,
  TabSheet_ClassType,
  TabSheet_ClassName,
  TabSheet_InstanceSize,
  TabSheet_InheritsFrom,
  TabSheet_Equals,
  TabSheet_GetHashCode,
  TabSheet_ToString,
  TabSheet_AnchorToNeighbour,
  TabSheet_AnchorParallel,
  TabSheet_AnchorHorizontalCenterTo,
  TabSheet_AnchorVerticalCenterTo,
  TabSheet_AnchorAsAlign,
  TabSheet_AnchorClient,
  TabSheet_GetPageControl,
  TabSheet_SetPageControl,
  TabSheet_GetTabIndex,
  TabSheet_GetBorderWidth,
  TabSheet_SetBorderWidth,
  TabSheet_GetCaption,
  TabSheet_SetCaption,
  TabSheet_GetDoubleBuffered,
  TabSheet_SetDoubleBuffered,
  TabSheet_GetEnabled,
  TabSheet_SetEnabled,
  TabSheet_GetFont,
  TabSheet_SetFont,
  TabSheet_GetHeight,
  TabSheet_SetHeight,
  TabSheet_GetImageIndex,
  TabSheet_SetImageIndex,
  TabSheet_GetLeft,
  TabSheet_SetLeft,
  TabSheet_GetConstraints,
  TabSheet_SetConstraints,
  TabSheet_GetPageIndex,
  TabSheet_SetPageIndex,
  TabSheet_GetParentDoubleBuffered,
  TabSheet_SetParentDoubleBuffered,
  TabSheet_GetParentFont,
  TabSheet_SetParentFont,
  TabSheet_GetParentShowHint,
  TabSheet_SetParentShowHint,
  TabSheet_GetPopupMenu,
  TabSheet_SetPopupMenu,
  TabSheet_GetShowHint,
  TabSheet_SetShowHint,
  TabSheet_GetTabVisible,
  TabSheet_SetTabVisible,
  TabSheet_GetTop,
  TabSheet_SetTop,
  TabSheet_GetVisible,
  TabSheet_SetVisible,
  TabSheet_GetWidth,
  TabSheet_SetWidth,
  TabSheet_SetOnContextPopup,
  TabSheet_SetOnDragDrop,
  TabSheet_SetOnDragOver,
  TabSheet_SetOnEndDrag,
  TabSheet_SetOnEnter,
  TabSheet_SetOnExit,
  TabSheet_SetOnHide,
  TabSheet_SetOnMouseDown,
  TabSheet_SetOnMouseEnter,
  TabSheet_SetOnMouseLeave,
  TabSheet_SetOnMouseMove,
  TabSheet_SetOnMouseUp,
  TabSheet_SetOnResize,
  TabSheet_SetOnShow,
  TabSheet_GetDockClientCount,
  TabSheet_GetDockSite,
  TabSheet_SetDockSite,
  TabSheet_GetMouseInClient,
  TabSheet_GetVisibleDockClientCount,
  TabSheet_GetBrush,
  TabSheet_GetControlCount,
  TabSheet_GetHandle,
  TabSheet_GetParentWindow,
  TabSheet_SetParentWindow,
  TabSheet_GetTabOrder,
  TabSheet_SetTabOrder,
  TabSheet_GetTabStop,
  TabSheet_SetTabStop,
  TabSheet_GetUseDockManager,
  TabSheet_SetUseDockManager,
  TabSheet_GetAction,
  TabSheet_SetAction,
  TabSheet_GetAlign,
  TabSheet_SetAlign,
  TabSheet_GetAnchors,
  TabSheet_SetAnchors,
  TabSheet_GetBiDiMode,
  TabSheet_SetBiDiMode,
  TabSheet_GetBoundsRect,
  TabSheet_SetBoundsRect,
  TabSheet_GetClientHeight,
  TabSheet_SetClientHeight,
  TabSheet_GetClientOrigin,
  TabSheet_GetClientRect,
  TabSheet_GetClientWidth,
  TabSheet_SetClientWidth,
  TabSheet_GetControlState,
  TabSheet_SetControlState,
  TabSheet_GetControlStyle,
  TabSheet_SetControlStyle,
  TabSheet_GetFloating,
  TabSheet_GetParent,
  TabSheet_SetParent,
  TabSheet_GetCursor,
  TabSheet_SetCursor,
  TabSheet_GetHint,
  TabSheet_SetHint,
  TabSheet_GetComponentCount,
  TabSheet_GetComponentIndex,
  TabSheet_SetComponentIndex,
  TabSheet_GetOwner,
  TabSheet_GetName,
  TabSheet_SetName,
  TabSheet_GetTag,
  TabSheet_SetTag,
  TabSheet_GetAnchorSideLeft,
  TabSheet_SetAnchorSideLeft,
  TabSheet_GetAnchorSideTop,
  TabSheet_SetAnchorSideTop,
  TabSheet_GetAnchorSideRight,
  TabSheet_SetAnchorSideRight,
  TabSheet_GetAnchorSideBottom,
  TabSheet_SetAnchorSideBottom,
  TabSheet_GetChildSizing,
  TabSheet_SetChildSizing,
  TabSheet_GetBorderSpacing,
  TabSheet_SetBorderSpacing,
  TabSheet_GetDockClients,
  TabSheet_GetControls,
  TabSheet_GetComponents,
  TabSheet_GetAnchorSide,
  TabSheet_StaticClassType;

