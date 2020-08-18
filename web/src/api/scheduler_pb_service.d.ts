// package: 
// file: scheduler.proto

import * as scheduler_pb from "./scheduler_pb";
import {grpc} from "@improbable-eng/grpc-web";

type SchedulerGetAction = {
  readonly methodName: string;
  readonly service: typeof Scheduler;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof scheduler_pb.GetActionRequest;
  readonly responseType: typeof scheduler_pb.GetActionResponse;
};

type SchedulerGetActionControl = {
  readonly methodName: string;
  readonly service: typeof Scheduler;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof scheduler_pb.GetActionControlRequest;
  readonly responseType: typeof scheduler_pb.GetActionControlResponse;
};

type SchedulerUpdateJobStatus = {
  readonly methodName: string;
  readonly service: typeof Scheduler;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof scheduler_pb.UpdateJobStatusRequest;
  readonly responseType: typeof scheduler_pb.UpdateJobStatusResponse;
};

type SchedulerUpdateJobLog = {
  readonly methodName: string;
  readonly service: typeof Scheduler;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof scheduler_pb.UpdateJobLogRequest;
  readonly responseType: typeof scheduler_pb.UpdateJobLogResponse;
};

type SchedulerGetWorkspaceVersion = {
  readonly methodName: string;
  readonly service: typeof Scheduler;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof scheduler_pb.GetWorkspaceVersionRequest;
  readonly responseType: typeof scheduler_pb.GetWorkspaceVersionResponse;
};

export class Scheduler {
  static readonly serviceName: string;
  static readonly GetAction: SchedulerGetAction;
  static readonly GetActionControl: SchedulerGetActionControl;
  static readonly UpdateJobStatus: SchedulerUpdateJobStatus;
  static readonly UpdateJobLog: SchedulerUpdateJobLog;
  static readonly GetWorkspaceVersion: SchedulerGetWorkspaceVersion;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class SchedulerClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getAction(
    requestMessage: scheduler_pb.GetActionRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: scheduler_pb.GetActionResponse|null) => void
  ): UnaryResponse;
  getAction(
    requestMessage: scheduler_pb.GetActionRequest,
    callback: (error: ServiceError|null, responseMessage: scheduler_pb.GetActionResponse|null) => void
  ): UnaryResponse;
  getActionControl(
    requestMessage: scheduler_pb.GetActionControlRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: scheduler_pb.GetActionControlResponse|null) => void
  ): UnaryResponse;
  getActionControl(
    requestMessage: scheduler_pb.GetActionControlRequest,
    callback: (error: ServiceError|null, responseMessage: scheduler_pb.GetActionControlResponse|null) => void
  ): UnaryResponse;
  updateJobStatus(
    requestMessage: scheduler_pb.UpdateJobStatusRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: scheduler_pb.UpdateJobStatusResponse|null) => void
  ): UnaryResponse;
  updateJobStatus(
    requestMessage: scheduler_pb.UpdateJobStatusRequest,
    callback: (error: ServiceError|null, responseMessage: scheduler_pb.UpdateJobStatusResponse|null) => void
  ): UnaryResponse;
  updateJobLog(
    requestMessage: scheduler_pb.UpdateJobLogRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: scheduler_pb.UpdateJobLogResponse|null) => void
  ): UnaryResponse;
  updateJobLog(
    requestMessage: scheduler_pb.UpdateJobLogRequest,
    callback: (error: ServiceError|null, responseMessage: scheduler_pb.UpdateJobLogResponse|null) => void
  ): UnaryResponse;
  getWorkspaceVersion(
    requestMessage: scheduler_pb.GetWorkspaceVersionRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: scheduler_pb.GetWorkspaceVersionResponse|null) => void
  ): UnaryResponse;
  getWorkspaceVersion(
    requestMessage: scheduler_pb.GetWorkspaceVersionRequest,
    callback: (error: ServiceError|null, responseMessage: scheduler_pb.GetWorkspaceVersionResponse|null) => void
  ): UnaryResponse;
}

