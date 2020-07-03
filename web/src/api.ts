import {
    ApproveJobRequest, CancelJobRequest,
    GetJobRequest,
    ListProjectsRequest,
    ListWorkspacesRequest,
    RefreshProjectRequest,
    SubmitJobRequest
} from "./api/web_pb";
import {Job} from "./api/common_pb";
import {WebClient} from "./api/web_pb_service";

export function listProjects(): Promise<string[]> {
    return new Promise<string[]>((resolve, reject) => {
        const client = new WebClient("");
        client.listProjects(new ListProjectsRequest(), (err, resp) => {
            if (resp) {
                resolve(resp.getProjectsList().map(e => e.getName()));
            } else {
                reject(err);
            }
        })
    })
}

export function listWorkspaces(project: string): Promise<string[]> {
    return new Promise(((resolve, reject) => {
        const client = new WebClient("");
        const req = new ListWorkspacesRequest();
        req.setProject(project);
        client.listWorkspaces(req, (err, resp) => {
            if (resp) {
                resolve(resp.getWorkspacesList().map(resp => resp.getName()));
            } else {
                reject(err || "failed to list workspaces");
            }
        })
    }))
}

export function refreshProject(project: string): Promise<void> {
    return new Promise<void>((resolve, reject) => {
        const client = new WebClient("");
        const req = new RefreshProjectRequest();
        req.setProject(project);
        client.refreshProject(req, (err) => {
            if (!err) {
                resolve();
            } else {
                reject(err);
            }
        });
    })
}

export function getJob(project: string, workspace: string): Promise<Job | undefined> {
    const req = new GetJobRequest();
    req.setProject(project);
    req.setWorkspace(workspace);

    return new Promise<Job | undefined>((resolve, reject) => {
        const client = new WebClient("");
        client.getJob(req, (err, resp) => {
            if (!err) {
                resolve(resp?.getJob());
            } else {
                reject(err);
            }
        })
    });
}

export function submitJob(project: string, workspace: string): Promise<void> {
    return new Promise<void>((resolve, reject) => {
        const req = new SubmitJobRequest();
        req.setProject(project);
        req.setWorkspace(workspace);

        const client = new WebClient("");
        client.submitJob(req, (err) => {
            if (!err) {
                resolve();
            } else {
                reject(err);
            }
        })
    })

}

export function approveJob(project: string, workspace: string): Promise<void> {
    return new Promise<void>((resolve, reject) => {
        const req = new ApproveJobRequest();
        req.setProject(project);
        req.setWorkspace(workspace);

        const client = new WebClient("");
        client.approveJob(req, (err) => {
            if (!err) {
                resolve();
            } else {
                reject(err);
            }
        })
    })
}

export function cancelJob(project: string, workspace: string): Promise<void> {
    return new Promise<void>((resolve, reject) => {
        const req = new CancelJobRequest();
        req.setProject(project);
        req.setWorkspace(workspace);

        const client = new WebClient("");
        client.cancelJob(req, (err) => {
            if (!err) {
                resolve();
            } else {
                reject(err);
            }
        })
    })
}
