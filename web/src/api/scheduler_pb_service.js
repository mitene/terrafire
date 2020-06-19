/* eslint-disable */
// package: 
// file: scheduler.proto

var scheduler_pb = require("./scheduler_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Scheduler = (function () {
  function Scheduler() {}
  Scheduler.serviceName = "Scheduler";
  return Scheduler;
}());

Scheduler.GetAction = {
  methodName: "GetAction",
  service: Scheduler,
  requestStream: false,
  responseStream: false,
  requestType: scheduler_pb.GetActionRequest,
  responseType: scheduler_pb.GetActionResponse
};

Scheduler.GetActionControl = {
  methodName: "GetActionControl",
  service: Scheduler,
  requestStream: false,
  responseStream: false,
  requestType: scheduler_pb.GetActionControlRequest,
  responseType: scheduler_pb.GetActionControlResponse
};

Scheduler.UpdateJobStatus = {
  methodName: "UpdateJobStatus",
  service: Scheduler,
  requestStream: false,
  responseStream: false,
  requestType: scheduler_pb.UpdateJobStatusRequest,
  responseType: scheduler_pb.UpdateJobStatusResponse
};

Scheduler.UpdateJobLog = {
  methodName: "UpdateJobLog",
  service: Scheduler,
  requestStream: false,
  responseStream: false,
  requestType: scheduler_pb.UpdateJobLogRequest,
  responseType: scheduler_pb.UpdateJobLogResponse
};

Scheduler.GetWorkspaceVersion = {
  methodName: "GetWorkspaceVersion",
  service: Scheduler,
  requestStream: false,
  responseStream: false,
  requestType: scheduler_pb.GetWorkspaceVersionRequest,
  responseType: scheduler_pb.GetWorkspaceVersionResponse
};

exports.Scheduler = Scheduler;

function SchedulerClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

SchedulerClient.prototype.getAction = function getAction(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Scheduler.GetAction, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SchedulerClient.prototype.getActionControl = function getActionControl(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Scheduler.GetActionControl, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SchedulerClient.prototype.updateJobStatus = function updateJobStatus(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Scheduler.UpdateJobStatus, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SchedulerClient.prototype.updateJobLog = function updateJobLog(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Scheduler.UpdateJobLog, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SchedulerClient.prototype.getWorkspaceVersion = function getWorkspaceVersion(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Scheduler.GetWorkspaceVersion, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.SchedulerClient = SchedulerClient;

