/* eslint-disable */
// package: 
// file: web.proto

var web_pb = require("./web_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Web = (function () {
  function Web() {}
  Web.serviceName = "Web";
  return Web;
}());

Web.ListProjects = {
  methodName: "ListProjects",
  service: Web,
  requestStream: false,
  responseStream: false,
  requestType: web_pb.ListProjectsRequest,
  responseType: web_pb.ListProjectsResponse
};

Web.RefreshProject = {
  methodName: "RefreshProject",
  service: Web,
  requestStream: false,
  responseStream: false,
  requestType: web_pb.RefreshProjectRequest,
  responseType: web_pb.RefreshProjectResponse
};

Web.GetProject = {
  methodName: "GetProject",
  service: Web,
  requestStream: false,
  responseStream: false,
  requestType: web_pb.GetProjectRequest,
  responseType: web_pb.GetProjectResponse
};

Web.GetJob = {
  methodName: "GetJob",
  service: Web,
  requestStream: false,
  responseStream: false,
  requestType: web_pb.GetJobRequest,
  responseType: web_pb.GetJobResponse
};

Web.SubmitJob = {
  methodName: "SubmitJob",
  service: Web,
  requestStream: false,
  responseStream: false,
  requestType: web_pb.SubmitJobRequest,
  responseType: web_pb.SubmitJobResponse
};

Web.ApproveJob = {
  methodName: "ApproveJob",
  service: Web,
  requestStream: false,
  responseStream: false,
  requestType: web_pb.ApproveJobRequest,
  responseType: web_pb.ApproveJobResponse
};

exports.Web = Web;

function WebClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

WebClient.prototype.listProjects = function listProjects(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Web.ListProjects, {
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

WebClient.prototype.refreshProject = function refreshProject(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Web.RefreshProject, {
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

WebClient.prototype.getProject = function getProject(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Web.GetProject, {
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

WebClient.prototype.getJob = function getJob(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Web.GetJob, {
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

WebClient.prototype.submitJob = function submitJob(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Web.SubmitJob, {
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

WebClient.prototype.approveJob = function approveJob(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Web.ApproveJob, {
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

exports.WebClient = WebClient;

