import React, {useEffect, useState} from 'react';
import {Redirect, Route, Switch, useParams} from "react-router-dom";
import {GlobalMenu} from "./GlobalMenu";
import {Project} from "../api/common_pb";
import {WorkspacesList} from "./WorkspacesList";
import {WorkspaceDetail} from "./WorkspaceDetail";
import {getProject, listProjects, refreshProject} from "../api";

export const App: React.FC = () => {
    const projects = useProjects();

    if (projects.length === 0) {
        return null;
    }

    return (
        <Switch>
            <Route path="/projects/:project">
                <AppBody projects={projects}/>
            </Route>
            <Route>
                <Redirect to={`/projects/${projects[0]}`}/>
            </Route>
        </Switch>
    );
}

type Props = {
    projects: string[]
}

const AppBody: React.FC<Props> = props => {
    const projects = props.projects;

    const {project} = useParams<{ project: string }>();
    const [pj, reload] = useProject(project);

    const handleRefresh = () => {
        refreshProject(project).then(reload).catch(console.log);
    }

    if (!pj) {
        return null;
    }

    return (
        <GlobalMenu current={project} projects={projects} onRefresh={handleRefresh}>
            <Switch>
                <Route exact path={`/projects/${project}/workspaces`}>
                    <WorkspacesList project={pj}/>
                </Route>

                <Route exact path={`/projects/${project}/workspaces/:workspace`} render={(props) =>
                    <WorkspaceDetail project={pj} workspace={props.match.params.workspace}/>
                }/>

                <Route>
                    <Redirect to={`/projects/${project}/workspaces`}/>
                </Route>
            </Switch>
        </GlobalMenu>
    );
}

/**
 * Hooks
 */


function useProjects(): string[] {
    const [pj, setPj] = useState<string[]>([]);

    useEffect(() => {
        listProjects().then(setPj).catch(console.log);
    }, [])

    return pj;
}

function useProject(project: string): [Project | undefined, () => void] {
    const [pj, setPj] = useState<Project | undefined>(undefined);

    const reload = () => {
        getProject(project).then(setPj).catch(console.log);
    }

    useEffect(reload, [project])

    return [pj, reload];
}
