// package: 
// file: scheduler.proto

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";

export class GetActionRequest extends jspb.Message {
  getTimeout(): number;
  setTimeout(value: number): void;

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
    timeout: number,
  }
}

export class GetActionResponse extends jspb.Message {
  getType(): GetActionResponse.ActionTypeMap[keyof GetActionResponse.ActionTypeMap];
  setType(value: GetActionResponse.ActionTypeMap[keyof GetActionResponse.ActionTypeMap]): void;

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
    type: GetActionResponse.ActionTypeMap[keyof GetActionResponse.ActionTypeMap],
    project: string,
    workspace: string,
  }

  export interface ActionTypeMap {
    NONE: 0;
    SUBMIT: 1;
    APPROVE: 2;
  }

  export const ActionType: ActionTypeMap;
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

