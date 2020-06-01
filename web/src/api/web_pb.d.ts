// package: 
// file: web.proto

import * as jspb from "google-protobuf";
import * as common_pb from "./common_pb";

export class ListProjectsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListProjectsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListProjectsRequest): ListProjectsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListProjectsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListProjectsRequest;
  static deserializeBinaryFromReader(message: ListProjectsRequest, reader: jspb.BinaryReader): ListProjectsRequest;
}

export namespace ListProjectsRequest {
  export type AsObject = {
  }
}

export class ListProjectsResponse extends jspb.Message {
  clearProjectsList(): void;
  getProjectsList(): Array<ListProjectsResponse.Project>;
  setProjectsList(value: Array<ListProjectsResponse.Project>): void;
  addProjects(value?: ListProjectsResponse.Project, index?: number): ListProjectsResponse.Project;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListProjectsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListProjectsResponse): ListProjectsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListProjectsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListProjectsResponse;
  static deserializeBinaryFromReader(message: ListProjectsResponse, reader: jspb.BinaryReader): ListProjectsResponse;
}

export namespace ListProjectsResponse {
  export type AsObject = {
    projectsList: Array<ListProjectsResponse.Project.AsObject>,
  }

  export class Project extends jspb.Message {
    getName(): string;
    setName(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Project.AsObject;
    static toObject(includeInstance: boolean, msg: Project): Project.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Project, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Project;
    static deserializeBinaryFromReader(message: Project, reader: jspb.BinaryReader): Project;
  }

  export namespace Project {
    export type AsObject = {
      name: string,
    }
  }
}

export class RefreshProjectRequest extends jspb.Message {
  getProject(): string;
  setProject(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshProjectRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshProjectRequest): RefreshProjectRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RefreshProjectRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshProjectRequest;
  static deserializeBinaryFromReader(message: RefreshProjectRequest, reader: jspb.BinaryReader): RefreshProjectRequest;
}

export namespace RefreshProjectRequest {
  export type AsObject = {
    project: string,
  }
}

export class RefreshProjectResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshProjectResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshProjectResponse): RefreshProjectResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RefreshProjectResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshProjectResponse;
  static deserializeBinaryFromReader(message: RefreshProjectResponse, reader: jspb.BinaryReader): RefreshProjectResponse;
}

export namespace RefreshProjectResponse {
  export type AsObject = {
  }
}

export class GetProjectRequest extends jspb.Message {
  getProject(): string;
  setProject(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetProjectRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetProjectRequest): GetProjectRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetProjectRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetProjectRequest;
  static deserializeBinaryFromReader(message: GetProjectRequest, reader: jspb.BinaryReader): GetProjectRequest;
}

export namespace GetProjectRequest {
  export type AsObject = {
    project: string,
  }
}

export class GetProjectResponse extends jspb.Message {
  hasProject(): boolean;
  clearProject(): void;
  getProject(): common_pb.Project | undefined;
  setProject(value?: common_pb.Project): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetProjectResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetProjectResponse): GetProjectResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetProjectResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetProjectResponse;
  static deserializeBinaryFromReader(message: GetProjectResponse, reader: jspb.BinaryReader): GetProjectResponse;
}

export namespace GetProjectResponse {
  export type AsObject = {
    project?: common_pb.Project.AsObject,
  }
}

export class GetJobRequest extends jspb.Message {
  getProject(): string;
  setProject(value: string): void;

  getWorkspace(): string;
  setWorkspace(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetJobRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetJobRequest): GetJobRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetJobRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetJobRequest;
  static deserializeBinaryFromReader(message: GetJobRequest, reader: jspb.BinaryReader): GetJobRequest;
}

export namespace GetJobRequest {
  export type AsObject = {
    project: string,
    workspace: string,
  }
}

export class GetJobResponse extends jspb.Message {
  hasJob(): boolean;
  clearJob(): void;
  getJob(): common_pb.Job | undefined;
  setJob(value?: common_pb.Job): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetJobResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetJobResponse): GetJobResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetJobResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetJobResponse;
  static deserializeBinaryFromReader(message: GetJobResponse, reader: jspb.BinaryReader): GetJobResponse;
}

export namespace GetJobResponse {
  export type AsObject = {
    job?: common_pb.Job.AsObject,
  }
}

export class SubmitJobRequest extends jspb.Message {
  getProject(): string;
  setProject(value: string): void;

  getWorkspace(): string;
  setWorkspace(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SubmitJobRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SubmitJobRequest): SubmitJobRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SubmitJobRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SubmitJobRequest;
  static deserializeBinaryFromReader(message: SubmitJobRequest, reader: jspb.BinaryReader): SubmitJobRequest;
}

export namespace SubmitJobRequest {
  export type AsObject = {
    project: string,
    workspace: string,
  }
}

export class SubmitJobResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SubmitJobResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SubmitJobResponse): SubmitJobResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SubmitJobResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SubmitJobResponse;
  static deserializeBinaryFromReader(message: SubmitJobResponse, reader: jspb.BinaryReader): SubmitJobResponse;
}

export namespace SubmitJobResponse {
  export type AsObject = {
  }
}

export class ApproveJobRequest extends jspb.Message {
  getProject(): string;
  setProject(value: string): void;

  getWorkspace(): string;
  setWorkspace(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApproveJobRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ApproveJobRequest): ApproveJobRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ApproveJobRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApproveJobRequest;
  static deserializeBinaryFromReader(message: ApproveJobRequest, reader: jspb.BinaryReader): ApproveJobRequest;
}

export namespace ApproveJobRequest {
  export type AsObject = {
    project: string,
    workspace: string,
  }
}

export class ApproveJobResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApproveJobResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ApproveJobResponse): ApproveJobResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ApproveJobResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApproveJobResponse;
  static deserializeBinaryFromReader(message: ApproveJobResponse, reader: jspb.BinaryReader): ApproveJobResponse;
}

export namespace ApproveJobResponse {
  export type AsObject = {
  }
}

