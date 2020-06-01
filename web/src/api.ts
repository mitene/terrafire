import {
    ApproveJobRequest,
    GetJobRequest,
    GetProjectRequest, ListProjectsRequest,
    RefreshProjectRequest,
    SubmitJobRequest
} from "./api/web_pb";
import {Job, Project} from "./api/common_pb";
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
export function getProject(project: string): Promise<Project> {
    return new Promise<Project>((resolve, reject) => {
        const client = new WebClient("");
        const req = new GetProjectRequest();
        req.setProject(project);
        client.getProject(req, (err, resp) => {
            const pj = resp?.getProject();
            if (pj) {
                resolve(pj);
            } else {
                reject(err || "failed to get project");
            }
        })
    })
}

export function refreshProject(project: string): Promise<void> {
    return new Promise<void>((resolve, reject) => {
        const client = new WebClient("");
        const req = new RefreshProjectRequest();
        req.setProject(project);
        client.refreshProject(req, (err, resp) => {
            if (!err) {
                resolve();
            } else {
                reject(err);
            }
        });
    })
}

export function getJob(project: string, workspace: string) {
    const req = new GetJobRequest();
    req.setProject(project);
    req.setWorkspace(workspace);

    return new Promise<Job>((resolve, reject) => {
        const client = new WebClient("");
        client.getJob(req, (err, resp) => {
            const j = resp?.getJob();
            if (j) {
                resolve(j);
            } else {
                reject(err || `job not found in workspace ${project}/${workspace}`);
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
