// package: 
// file: scheduler.proto

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";

export class GetActionRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetActionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetActionRequest): GetActionRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetActionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetActionRequest;
  static deserializeBinaryFromReader(message: GetActionRequest, reader: jspb.BinaryReader): GetActionRequest;
}

export namespace GetActionRequest {
  export type AsObject = {
  }
}

export class GetActionResponse extends jspb.Message {
  getType(): GetActionResponse.TypeMap[keyof GetActionResponse.TypeMap];
  setType(value: GetActionResponse.TypeMap[keyof GetActionResponse.TypeMap]): void;

  getProject(): string;
  setProject(value: string): void;

  getWorkspace(): string;
  setWorkspace(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetActionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetActionResponse): GetActionResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetActionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetActionResponse;
  static deserializeBinaryFromReader(message: GetActionResponse, reader: jspb.BinaryReader): GetActionResponse;
}

export namespace GetActionResponse {
  export type AsObject = {
    type: GetActionResponse.TypeMap[keyof GetActionResponse.TypeMap],
    project: string,
    workspace: string,
  }

  export interface TypeMap {
    NONE: 0;
    SUBMIT: 1;
    APPROVE: 2;
  }

  export const Type: TypeMap;
}

export class GetActionControlRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetActionControlRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetActionControlRequest): GetActionControlRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetActionControlRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetActionControlRequest;
  static deserializeBinaryFromReader(message: GetActionControlRequest, reader: jspb.BinaryReader): GetActionControlRequest;
}

export namespace GetActionControlRequest {
  export type AsObject = {
  }
}

export class GetActionControlResponse extends jspb.Message {
  getType(): GetActionControlResponse.TypeMap[keyof GetActionControlResponse.TypeMap];
  setType(value: GetActionControlResponse.TypeMap[keyof GetActionControlResponse.TypeMap]): void;

  getProject(): string;
  setProject(value: string): void;

  getWorkspace(): string;
  setWorkspace(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetActionControlResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetActionControlResponse): GetActionControlResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetActionControlResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetActionControlResponse;
  static deserializeBinaryFromReader(message: GetActionControlResponse, reader: jspb.BinaryReader): GetActionControlResponse;
}

export namespace GetActionControlResponse {
  export type AsObject = {
    type: GetActionControlResponse.TypeMap[keyof GetActionControlResponse.TypeMap],
    project: string,
    workspace: string,
  }

  export interface TypeMap {
    NONE: 0;
    CANCEL: 1;
  }

  export const Type: TypeMap;
}

export class UpdateJobStatusRequest extends jspb.Message {
  getProject(): string;
  setProject(value: string): void;

  getWorkspace(): string;
  setWorkspace(value: string): void;

  getStatus(): common_pb.Job.StatusMap[keyof common_pb.Job.StatusMap];
  setStatus(value: common_pb.Job.StatusMap[keyof common_pb.Job.StatusMap]): void;

  getResult(): string;
  setResult(value: string): void;

  getError(): string;
  setError(value: string): void;

  getProjectVersion(): string;
  setProjectVersion(value: string): void;

  getWorkspaceVersion(): string;
  setWorkspaceVersion(value: string): void;

  getDestroy(): boolean;
  setDestroy(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateJobStatusRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateJobStatusRequest): UpdateJobStatusRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateJobStatusRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateJobStatusRequest;
  static deserializeBinaryFromReader(message: UpdateJobStatusRequest, reader: jspb.BinaryReader): UpdateJobStatusRequest;
}

export namespace UpdateJobStatusRequest {
  export type AsObject = {
    project: string,
    workspace: string,
    status: common_pb.Job.StatusMap[keyof common_pb.Job.StatusMap],
    result: string,
    error: string,
    projectVersion: string,
    workspaceVersion: string,
    destroy: boolean,
  }
}

export class UpdateJobStatusResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateJobStatusResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateJobStatusResponse): UpdateJobStatusResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateJobStatusResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateJobStatusResponse;
  static deserializeBinaryFromReader(message: UpdateJobStatusResponse, reader: jspb.BinaryReader): UpdateJobStatusResponse;
}

export namespace UpdateJobStatusResponse {
  export type AsObject = {
  }
}

export class UpdateJobLogRequest extends jspb.Message {
  getProject(): string;
  setProject(value: string): void;

  getWorkspace(): string;
  setWorkspace(value: string): void;

  getPhase(): common_pb.PhaseMap[keyof common_pb.PhaseMap];
  setPhase(value: common_pb.PhaseMap[keyof common_pb.PhaseMap]): void;

  getLog(): string;
  setLog(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateJobLogRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateJobLogRequest): UpdateJobLogRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateJobLogRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateJobLogRequest;
  static deserializeBinaryFromReader(message: UpdateJobLogRequest, reader: jspb.BinaryReader): UpdateJobLogRequest;
}

export namespace UpdateJobLogRequest {
  export type AsObject = {
    project: string,
    workspace: string,
    phase: common_pb.PhaseMap[keyof common_pb.PhaseMap],
    log: string,
  }
}

export class UpdateJobLogResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateJobLogResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateJobLogResponse): UpdateJobLogResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateJobLogResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateJobLogResponse;
  static deserializeBinaryFromReader(message: UpdateJobLogResponse, reader: jspb.BinaryReader): UpdateJobLogResponse;
}

export namespace UpdateJobLogResponse {
  export type AsObject = {
  }
}

export class GetWorkspaceVersionRequest extends jspb.Message {
  getProject(): string;
  setProject(value: string): void;

  getWorkspace(): string;
  setWorkspace(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetWorkspaceVersionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetWorkspaceVersionRequest): GetWorkspaceVersionRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetWorkspaceVersionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetWorkspaceVersionRequest;
  static deserializeBinaryFromReader(message: GetWorkspaceVersionRequest, reader: jspb.BinaryReader): GetWorkspaceVersionRequest;
}

export namespace GetWorkspaceVersionRequest {
  export type AsObject = {
    project: string,
    workspace: string,
  }
}

export class GetWorkspaceVersionResponse extends jspb.Message {
  getProjectVersion(): string;
  setProjectVersion(value: string): void;

  getWorkspaceVersion(): string;
  setWorkspaceVersion(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetWorkspaceVersionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetWorkspaceVersionResponse): GetWorkspaceVersionResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetWorkspaceVersionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetWorkspaceVersionResponse;
  static deserializeBinaryFromReader(message: GetWorkspaceVersionResponse, reader: jspb.BinaryReader): GetWorkspaceVersionResponse;
}

export namespace GetWorkspaceVersionResponse {
  export type AsObject = {
    projectVersion: string,
    workspaceVersion: string,
  }
}

