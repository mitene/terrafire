// package: 
// file: common.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Project extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  clearWorkspacesList(): void;
  getWorkspacesList(): Array<Workspace>;
  setWorkspacesList(value: Array<Workspace>): void;
  addWorkspaces(value?: Workspace, index?: number): Workspace;

  getVersion(): string;
  setVersion(value: string): void;

  getRepo(): string;
  setRepo(value: string): void;

  getBranch(): string;
  setBranch(value: string): void;

  getPath(): string;
  setPath(value: string): void;

  clearEnvsList(): void;
  getEnvsList(): Array<Pair>;
  setEnvsList(value: Array<Pair>): void;
  addEnvs(value?: Pair, index?: number): Pair;

  getError(): string;
  setError(value: string): void;

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
    workspacesList: Array<Workspace.AsObject>,
    version: string,
    repo: string,
    branch: string,
    path: string,
    envsList: Array<Pair.AsObject>,
    error: string,
  }
}

export class Workspace extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  hasSource(): boolean;
  clearSource(): void;
  getSource(): Source | undefined;
  setSource(value?: Source): void;

  getWorkspace(): string;
  setWorkspace(value: string): void;

  clearVarsList(): void;
  getVarsList(): Array<Pair>;
  setVarsList(value: Array<Pair>): void;
  addVars(value?: Pair, index?: number): Pair;

  clearVarfilesList(): void;
  getVarfilesList(): Array<string>;
  setVarfilesList(value: Array<string>): void;
  addVarfiles(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Workspace.AsObject;
  static toObject(includeInstance: boolean, msg: Workspace): Workspace.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Workspace, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Workspace;
  static deserializeBinaryFromReader(message: Workspace, reader: jspb.BinaryReader): Workspace;
}

export namespace Workspace {
  export type AsObject = {
    name: string,
    source?: Source.AsObject,
    workspace: string,
    varsList: Array<Pair.AsObject>,
    varfilesList: Array<string>,
  }
}

export class Manifest extends jspb.Message {
  clearWorkspacesList(): void;
  getWorkspacesList(): Array<Workspace>;
  setWorkspacesList(value: Array<Workspace>): void;
  addWorkspaces(value?: Workspace, index?: number): Workspace;

  getVersion(): string;
  setVersion(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Manifest.AsObject;
  static toObject(includeInstance: boolean, msg: Manifest): Manifest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Manifest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Manifest;
  static deserializeBinaryFromReader(message: Manifest, reader: jspb.BinaryReader): Manifest;
}

export namespace Manifest {
  export type AsObject = {
    workspacesList: Array<Workspace.AsObject>,
    version: string,
  }
}

export class Source extends jspb.Message {
  getType(): Source.TypeMap[keyof Source.TypeMap];
  setType(value: Source.TypeMap[keyof Source.TypeMap]): void;

  getOwner(): string;
  setOwner(value: string): void;

  getRepo(): string;
  setRepo(value: string): void;

  getPath(): string;
  setPath(value: string): void;

  getRef(): string;
  setRef(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Source.AsObject;
  static toObject(includeInstance: boolean, msg: Source): Source.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Source, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Source;
  static deserializeBinaryFromReader(message: Source, reader: jspb.BinaryReader): Source;
}

export namespace Source {
  export type AsObject = {
    type: Source.TypeMap[keyof Source.TypeMap],
    owner: string,
    repo: string,
    path: string,
    ref: string,
  }

  export interface TypeMap {
    GITHUB: 0;
  }

  export const Type: TypeMap;
}

export class GitRepository extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getProtocol(): string;
  setProtocol(value: string): void;

  getHost(): string;
  setHost(value: string): void;

  getUser(): string;
  setUser(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GitRepository.AsObject;
  static toObject(includeInstance: boolean, msg: GitRepository): GitRepository.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GitRepository, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GitRepository;
  static deserializeBinaryFromReader(message: GitRepository, reader: jspb.BinaryReader): GitRepository;
}

export namespace GitRepository {
  export type AsObject = {
    name: string,
    protocol: string,
    host: string,
    user: string,
    password: string,
  }
}

export class Job extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  hasStartedAt(): boolean;
  clearStartedAt(): void;
  getStartedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStartedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getProject(): string;
  setProject(value: string): void;

  getWorkspace(): string;
  setWorkspace(value: string): void;

  getStatus(): Job.StatusMap[keyof Job.StatusMap];
  setStatus(value: Job.StatusMap[keyof Job.StatusMap]): void;

  getPlanResult(): string;
  setPlanResult(value: string): void;

  getError(): string;
  setError(value: string): void;

  getPlanLog(): string;
  setPlanLog(value: string): void;

  getApplyLog(): string;
  setApplyLog(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Job.AsObject;
  static toObject(includeInstance: boolean, msg: Job): Job.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Job, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Job;
  static deserializeBinaryFromReader(message: Job, reader: jspb.BinaryReader): Job;
}

export namespace Job {
  export type AsObject = {
    id: number,
    startedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    project: string,
    workspace: string,
    status: Job.StatusMap[keyof Job.StatusMap],
    planResult: string,
    error: string,
    planLog: string,
    applyLog: string,
  }

  export interface StatusMap {
    PENDING: 0;
    PLANINPROGRESS: 1;
    REVIEWREQUIRED: 2;
    APPLYPENDING: 3;
    APPLYINPROGRESS: 4;
    SUCCEEDED: 5;
    PLANFAILED: 6;
    APPLYFAILED: 7;
  }

  export const Status: StatusMap;
}

export class Pair extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Pair.AsObject;
  static toObject(includeInstance: boolean, msg: Pair): Pair.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Pair, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Pair;
  static deserializeBinaryFromReader(message: Pair, reader: jspb.BinaryReader): Pair;
}

export namespace Pair {
  export type AsObject = {
    key: string,
    value: string,
  }
}

export interface PhaseMap {
  PLAN: 0;
  APPLY: 1;
}

export const Phase: PhaseMap;

