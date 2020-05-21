export interface Project {
    name: string
    repo: string
    branch: string
    path: string
    commit: string | null
}

export interface ProjectInfo {
    project: Project
    commit: string
    error: string
}

export interface Workspace {
    name: string
    source: Source
    workspace: string
    vars: any
    var_files: string[]
}

export interface WorkspaceInfo {
    project: ProjectInfo
    workspace: Workspace
    last_job: Job
}

export interface Source {
    type: string
    owner: string
    repo: string
    path: string
    ref: string
}

export interface Job {
    id: number
    started_at: any
    project: string
    workspace: string
    status: number
    plan_result: string
    error: string
    plan_log: string
    apply_log: string
}

export enum JobStatus {
    Pending,
    PlanInProgress,
    ReviewRequired,
    ApplyPending,
    ApplyInProgress,
    Succeeded,
    PlanFailed,
    ApplyFailed,
    Cancelled,
}

async function api(method: string, path: string): Promise<any> {
    const data = await fetch("/api/v1" + path, {
        method: method,
        headers: {
            "Content-Type": "application/json; charset=utf-8",
        },
    });
    const json = await data.json();

    console.log({method: method, path: path, resp: json});

    return json;
}

export async function listProjects(): Promise<Project[]> {
    const data = await api("GET", "/projects");
    return Object.values(data["projects"]);
}

export async function refreshProject(project: string): Promise<any> {
    return await api("POST", `/projects/${project}/refresh`);
}

export async function listWorkspaces(project: string): Promise<Map<string, Workspace>> {
    const data = await api("GET", `/projects/${project}/workspaces`);
    return new Map(Object.entries(data["workspaces"]));
}

export async function getWorkspace(project: string, workspace: string): Promise<WorkspaceInfo> {
    return await api("GET", `/projects/${project}/workspaces/${workspace}`);
}

export async function listJobs(project: string, workspace: string): Promise<Job[]> {
    const data = await api("GET", `/projects/${project}/workspaces/${workspace}/jobs`);
    return data["jobs"];
}

export async function getJob(job_id: number): Promise<Job> {
    return await api("GET", `/jobs/${job_id}`);
}

export async function submitJob(project: string, workspace: string): Promise<Job> {
    const data = await api("POST", `/projects/${project}/workspaces/${workspace}/jobs`);
    return data["job"];
}

export async function approveJob(project: string, workspace: string): Promise<any> {
    return await api("POST", `/projects/${project}/workspaces/${workspace}/approve`);
}
