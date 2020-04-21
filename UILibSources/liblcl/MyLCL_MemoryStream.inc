
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

function MemoryStream_Create: TMemoryStream; extdecl;
begin
  Result :=  TMemoryStream.Create;
end;

procedure MemoryStream_Free(AObj: TMemoryStream); extdecl;
begin
  AObj.Free;
end;

procedure MemoryStream_Clear(AObj: TMemoryStream); extdecl;
begin
  AObj.Clear;
end;

procedure MemoryStream_LoadFromStream(AObj: TMemoryStream; Stream: TStream); extdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure MemoryStream_LoadFromFile(AObj: TMemoryStream; FileName: PChar); extdecl;
begin
  AObj.LoadFromFile(FileName);
end;

procedure MemoryStream_Seek(AObj: TMemoryStream; var Offset: Int64; Origin: TSeekOrigin; var Result: Int64); extdecl;
begin
  Result :=  AObj.Seek(Offset, Origin);
end;

procedure MemoryStream_SaveToStream(AObj: TMemoryStream; Stream: TStream); extdecl;
begin
  AObj.SaveToStream(Stream);
end;

procedure MemoryStream_SaveToFile(AObj: TMemoryStream; FileName: PChar); extdecl;
begin
  AObj.SaveToFile(FileName);
end;

procedure MemoryStream_CopyFrom(AObj: TMemoryStream; Source: TStream; var Count: Int64; var Result: Int64); extdecl;
begin
  Result :=  AObj.CopyFrom(Source, Count);
end;

function MemoryStream_ClassType(AObj: TMemoryStream): TClass; extdecl;
begin
  Result :=  AObj.ClassType;
end;

function MemoryStream_ClassName(AObj: TMemoryStream): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function MemoryStream_InstanceSize(AObj: TMemoryStream): Integer; extdecl;
begin
  Result :=  AObj.InstanceSize;
end;

function MemoryStream_InheritsFrom(AObj: TMemoryStream; AClass: TClass): LongBool; extdecl;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function MemoryStream_Equals(AObj: TMemoryStream; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function MemoryStream_GetHashCode(AObj: TMemoryStream): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function MemoryStream_ToString(AObj: TMemoryStream): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function MemoryStream_GetMemory(AObj: TMemoryStream): Pointer; extdecl;
begin
  Result :=  AObj.Memory;
end;

procedure MemoryStream_GetPosition(AObj: TMemoryStream; var Result: Int64); extdecl;
begin
  Result :=  AObj.Position;
end;

procedure MemoryStream_SetPosition(AObj: TMemoryStream; var AValue: Int64); extdecl;
begin
  AObj.Position := AValue;
end;

procedure MemoryStream_GetSize(AObj: TMemoryStream; var Result: Int64); extdecl;
begin
  Result :=  AObj.Size;
end;

procedure MemoryStream_SetSize(AObj: TMemoryStream; var AValue: Int64); extdecl;
begin
  AObj.Size := AValue;
end;

function MemoryStream_StaticClassType: TClass; extdecl;
begin
  Result := TMemoryStream;
end;

exports
  MemoryStream_Create,
  MemoryStream_Free,
  MemoryStream_Clear,
  MemoryStream_LoadFromStream,
  MemoryStream_LoadFromFile,
  MemoryStream_Seek,
  MemoryStream_SaveToStream,
  MemoryStream_SaveToFile,
  MemoryStream_CopyFrom,
  MemoryStream_ClassType,
  MemoryStream_ClassName,
  MemoryStream_InstanceSize,
  MemoryStream_InheritsFrom,
  MemoryStream_Equals,
  MemoryStream_GetHashCode,
  MemoryStream_ToString,
  MemoryStream_GetMemory,
  MemoryStream_GetPosition,
  MemoryStream_SetPosition,
  MemoryStream_GetSize,
  MemoryStream_SetSize,
  MemoryStream_StaticClassType;
