
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

function SelectDirectoryDialog_Create(AOwner: TComponent): TSelectDirectoryDialog; extdecl;
begin
  Result :=  TSelectDirectoryDialog.Create(AOwner);
end;

procedure SelectDirectoryDialog_Free(AObj: TSelectDirectoryDialog); extdecl;
begin
  AObj.Free;
end;

function SelectDirectoryDialog_Execute(AObj: TSelectDirectoryDialog): LongBool; extdecl;
begin
  Result :=  AObj.Execute;
end;

function SelectDirectoryDialog_FindComponent(AObj: TSelectDirectoryDialog; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function SelectDirectoryDialog_GetNamePath(AObj: TSelectDirectoryDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function SelectDirectoryDialog_HasParent(AObj: TSelectDirectoryDialog): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure SelectDirectoryDialog_Assign(AObj: TSelectDirectoryDialog; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function SelectDirectoryDialog_ClassType(AObj: TSelectDirectoryDialog): TClass; extdecl;
begin
  Result :=  AObj.ClassType;
end;

function SelectDirectoryDialog_ClassName(AObj: TSelectDirectoryDialog): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function SelectDirectoryDialog_InstanceSize(AObj: TSelectDirectoryDialog): Integer; extdecl;
begin
  Result :=  AObj.InstanceSize;
end;

function SelectDirectoryDialog_InheritsFrom(AObj: TSelectDirectoryDialog; AClass: TClass): LongBool; extdecl;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function SelectDirectoryDialog_Equals(AObj: TSelectDirectoryDialog; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function SelectDirectoryDialog_GetHashCode(AObj: TSelectDirectoryDialog): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function SelectDirectoryDialog_ToString(AObj: TSelectDirectoryDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function SelectDirectoryDialog_GetFiles(AObj: TSelectDirectoryDialog): TStrings; extdecl;
begin
  Result :=  AObj.Files;
end;

function SelectDirectoryDialog_GetDefaultExt(AObj: TSelectDirectoryDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.DefaultExt);
end;

procedure SelectDirectoryDialog_SetDefaultExt(AObj: TSelectDirectoryDialog; AValue: PChar); extdecl;
begin
  AObj.DefaultExt := AValue;
end;

function SelectDirectoryDialog_GetFileName(AObj: TSelectDirectoryDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.FileName);
end;

procedure SelectDirectoryDialog_SetFileName(AObj: TSelectDirectoryDialog; AValue: PChar); extdecl;
begin
  AObj.FileName := AValue;
end;

function SelectDirectoryDialog_GetFilter(AObj: TSelectDirectoryDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Filter);
end;

procedure SelectDirectoryDialog_SetFilter(AObj: TSelectDirectoryDialog; AValue: PChar); extdecl;
begin
  AObj.Filter := AValue;
end;

function SelectDirectoryDialog_GetFilterIndex(AObj: TSelectDirectoryDialog): Integer; extdecl;
begin
  Result :=  AObj.FilterIndex;
end;

procedure SelectDirectoryDialog_SetFilterIndex(AObj: TSelectDirectoryDialog; AValue: Integer); extdecl;
begin
  AObj.FilterIndex := AValue;
end;

function SelectDirectoryDialog_GetInitialDir(AObj: TSelectDirectoryDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.InitialDir);
end;

procedure SelectDirectoryDialog_SetInitialDir(AObj: TSelectDirectoryDialog; AValue: PChar); extdecl;
begin
  AObj.InitialDir := AValue;
end;

function SelectDirectoryDialog_GetOptions(AObj: TSelectDirectoryDialog): TOpenOptions; extdecl;
begin
  Result :=  AObj.Options;
end;

procedure SelectDirectoryDialog_SetOptions(AObj: TSelectDirectoryDialog; AValue: TOpenOptions); extdecl;
begin
  AObj.Options := AValue;
end;

function SelectDirectoryDialog_GetTitle(AObj: TSelectDirectoryDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Title);
end;

procedure SelectDirectoryDialog_SetTitle(AObj: TSelectDirectoryDialog; AValue: PChar); extdecl;
begin
  AObj.Title := AValue;
end;

function SelectDirectoryDialog_GetHandle(AObj: TSelectDirectoryDialog): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure SelectDirectoryDialog_SetOnClose(AObj: TSelectDirectoryDialog; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnClose}
  {$DEFINE EventMethodName := OnClose}
  EventMethodCode
end;

procedure SelectDirectoryDialog_SetOnShow(AObj: TSelectDirectoryDialog; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnShow}
  {$DEFINE EventMethodName := OnShow}
  EventMethodCode
end;

function SelectDirectoryDialog_GetComponentCount(AObj: TSelectDirectoryDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function SelectDirectoryDialog_GetComponentIndex(AObj: TSelectDirectoryDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure SelectDirectoryDialog_SetComponentIndex(AObj: TSelectDirectoryDialog; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function SelectDirectoryDialog_GetOwner(AObj: TSelectDirectoryDialog): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function SelectDirectoryDialog_GetName(AObj: TSelectDirectoryDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure SelectDirectoryDialog_SetName(AObj: TSelectDirectoryDialog; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function SelectDirectoryDialog_GetTag(AObj: TSelectDirectoryDialog): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure SelectDirectoryDialog_SetTag(AObj: TSelectDirectoryDialog; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function SelectDirectoryDialog_GetComponents(AObj: TSelectDirectoryDialog; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;

function SelectDirectoryDialog_StaticClassType: TClass; extdecl;
begin
  Result := TSelectDirectoryDialog;
end;

exports
  SelectDirectoryDialog_Create,
  SelectDirectoryDialog_Free,
  SelectDirectoryDialog_Execute,
  SelectDirectoryDialog_FindComponent,
  SelectDirectoryDialog_GetNamePath,
  SelectDirectoryDialog_HasParent,
  SelectDirectoryDialog_Assign,
  SelectDirectoryDialog_ClassType,
  SelectDirectoryDialog_ClassName,
  SelectDirectoryDialog_InstanceSize,
  SelectDirectoryDialog_InheritsFrom,
  SelectDirectoryDialog_Equals,
  SelectDirectoryDialog_GetHashCode,
  SelectDirectoryDialog_ToString,
  SelectDirectoryDialog_GetFiles,
  SelectDirectoryDialog_GetDefaultExt,
  SelectDirectoryDialog_SetDefaultExt,
  SelectDirectoryDialog_GetFileName,
  SelectDirectoryDialog_SetFileName,
  SelectDirectoryDialog_GetFilter,
  SelectDirectoryDialog_SetFilter,
  SelectDirectoryDialog_GetFilterIndex,
  SelectDirectoryDialog_SetFilterIndex,
  SelectDirectoryDialog_GetInitialDir,
  SelectDirectoryDialog_SetInitialDir,
  SelectDirectoryDialog_GetOptions,
  SelectDirectoryDialog_SetOptions,
  SelectDirectoryDialog_GetTitle,
  SelectDirectoryDialog_SetTitle,
  SelectDirectoryDialog_GetHandle,
  SelectDirectoryDialog_SetOnClose,
  SelectDirectoryDialog_SetOnShow,
  SelectDirectoryDialog_GetComponentCount,
  SelectDirectoryDialog_GetComponentIndex,
  SelectDirectoryDialog_SetComponentIndex,
  SelectDirectoryDialog_GetOwner,
  SelectDirectoryDialog_GetName,
  SelectDirectoryDialog_SetName,
  SelectDirectoryDialog_GetTag,
  SelectDirectoryDialog_SetTag,
  SelectDirectoryDialog_GetComponents,
  SelectDirectoryDialog_StaticClassType;
