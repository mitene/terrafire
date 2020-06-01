// package: 
// file: web.proto

import * as web_pb from "./web_pb";
import {grpc} from "@improbable-eng/grpc-web";

type WebListProjects = {
  readonly methodName: string;
  readonly service: typeof Web;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof web_pb.ListProjectsRequest;
  readonly responseType: typeof web_pb.ListProjectsResponse;
};

type WebRefreshProject = {
  readonly methodName: string;
  readonly service: typeof Web;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof web_pb.RefreshProjectRequest;
  readonly responseType: typeof web_pb.RefreshProjectResponse;
};

type WebGetProject = {
  readonly methodName: string;
  readonly service: typeof Web;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof web_pb.GetProjectRequest;
  readonly responseType: typeof web_pb.GetProjectResponse;
};

type WebGetJob = {
  readonly methodName: string;
  readonly service: typeof Web;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof web_pb.GetJobRequest;
  readonly responseType: typeof web_pb.GetJobResponse;
};

type WebSubmitJob = {
  readonly methodName: string;
  readonly service: typeof Web;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof web_pb.SubmitJobRequest;
  readonly responseType: typeof web_pb.SubmitJobResponse;
};

type WebApproveJob = {
  readonly methodName: string;
  readonly service: typeof Web;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof web_pb.ApproveJobRequest;
  readonly responseType: typeof web_pb.ApproveJobResponse;
};

export class Web {
  static readonly serviceName: string;
  static readonly ListProjects: WebListProjects;
  static readonly RefreshProject: WebRefreshProject;
  static readonly GetProject: WebGetProject;
  static readonly GetJob: WebGetJob;
  static readonly SubmitJob: WebSubmitJob;
  static readonly ApproveJob: WebApproveJob;
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

export class WebClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  listProjects(
    requestMessage: web_pb.ListProjectsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: web_pb.ListProjectsResponse|null) => void
  ): UnaryResponse;
  listProjects(
    requestMessage: web_pb.ListProjectsRequest,
    callback: (error: ServiceError|null, responseMessage: web_pb.ListProjectsResponse|null) => void
  ): UnaryResponse;
  refreshProject(
    requestMessage: web_pb.RefreshProjectRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: web_pb.RefreshProjectResponse|null) => void
  ): UnaryResponse;
  refreshProject(
    requestMessage: web_pb.RefreshProjectRequest,
    callback: (error: ServiceError|null, responseMessage: web_pb.RefreshProjectResponse|null) => void
  ): UnaryResponse;
  getProject(
    requestMessage: web_pb.GetProjectRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: web_pb.GetProjectResponse|null) => void
  ): UnaryResponse;
  getProject(
    requestMessage: web_pb.GetProjectRequest,
    callback: (error: ServiceError|null, responseMessage: web_pb.GetProjectResponse|null) => void
  ): UnaryResponse;
  getJob(
    requestMessage: web_pb.GetJobRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: web_pb.GetJobResponse|null) => void
  ): UnaryResponse;
  getJob(
    requestMessage: web_pb.GetJobRequest,
    callback: (error: ServiceError|null, responseMessage: web_pb.GetJobResponse|null) => void
  ): UnaryResponse;
  submitJob(
    requestMessage: web_pb.SubmitJobRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: web_pb.SubmitJobResponse|null) => void
  ): UnaryResponse;
  submitJob(
    requestMessage: web_pb.SubmitJobRequest,
    callback: (error: ServiceError|null, responseMessage: web_pb.SubmitJobResponse|null) => void
  ): UnaryResponse;
  approveJob(
    requestMessage: web_pb.ApproveJobRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: web_pb.ApproveJobResponse|null) => void
  ): UnaryResponse;
  approveJob(
    requestMessage: web_pb.ApproveJobRequest,
    callback: (error: ServiceError|null, responseMessage: web_pb.ApproveJobResponse|null) => void
  ): UnaryResponse;
}

