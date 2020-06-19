import React, {useState} from 'react';
import {Redirect, Route, Switch, useParams} from "react-router-dom";
import {GlobalMenu} from "./GlobalMenu";
import {WorkspacesList} from "./WorkspacesList";
import {WorkspaceDetail} from "./WorkspaceDetail";
import {listProjects, refreshProject} from "../api";
import {useAsync} from "../hooks";

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
    const {project} = useParams<{ project: string }>();
    const [ts, reload] = useForceUpdate();

    const handleRefresh = () => {
        refreshProject(project).then(reload).catch(console.log);
    }

    return (
        <GlobalMenu current={project} projects={props.projects} onRefresh={handleRefresh}>
            <Switch>
                <Route exact path={`/projects/${project}/workspaces`}>
                    <WorkspacesList project={project} ts={ts}/>
                </Route>

                <Route exact path={`/projects/${project}/workspaces/:workspace`} render={(props) =>
                    <WorkspaceDetail project={project} workspace={props.match.params.workspace} ts={ts}/>
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
function useForceUpdate(): [number, () => void] {
    const [ts, setTs] = useState<number>(Date.now());
    return [ts, () => setTs(Date.now())];
}

function useProjects(): string[] {
    const [pj,] = useAsync(listProjects);
    return pj || [];
}
