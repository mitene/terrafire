import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";
import Title from "./Title";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import React, {useEffect} from "react";

import * as globalStyle from "./styles";
import {approveJob, getJob, listWorkspaces, submitJob} from "../api";
import {Job} from "../api/common_pb";
import {Redirect} from "react-router";
import {useAsync} from "../hooks";

type Props = {
    project: string
    workspace: string
    ts: number
}

export const WorkspaceDetail: React.FC<Props> = (props) => {
    const project = props.project;
    const workspace = props.workspace;

    const classes = globalStyle.useStyles();
    const wsExists = useWorkspaceExists(project, workspace, props.ts);
    const [j, reload] = useJob(project, workspace);

    if (!wsExists) {
        return (
            <Redirect to={`/projects/${project}/workspaces`}/>
        );
    }

    const state = getState(j);
    const handlePlan = () => submitJob(project, workspace).then(reload).catch(console.log);
    const handleApply = () => approveJob(project, workspace).then(reload).catch(console.log);

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={1}>
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>{workspace}</Title>
                        <Typography>Status: {state.status}</Typography>
                        <Grid container spacing={3}>
                            <Grid item xs>
                                <Button size="small" variant="contained" color="primary"
                                        onClick={handlePlan}
                                        disabled={!state.planAvailable}>Plan</Button>
                                <Button size="small" variant="contained" color="secondary"
                                        onClick={handleApply}
                                        disabled={!state.applyAvailable}>Apply</Button>
                            </Grid>
                        </Grid>
                    </Paper>
                </Grid>
                {j && j.getApplyLog() &&
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>Apply Log</Title>
                        <pre><code>{j.getApplyLog()}</code></pre>
                    </Paper>
                </Grid>
                }
                {j && j.getPlanResult() &&
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>Plan Result</Title>
                        {j.getDestroy() && <Typography color="secondary">[DESTROY]</Typography>}
                        <Typography>Project Version: {j.getProjectVersion()}</Typography>
                        <Typography>Workspace Version: {j.getWorkspaceVersion()}</Typography>
                        <pre><code>{j.getPlanResult()}</code></pre>
                    </Paper>
                </Grid>
                }
                {j && j.getPlanLog() &&
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>Plan Log</Title>
                        <pre><code>{j.getPlanLog()}</code></pre>
                    </Paper>
                </Grid>
                }
            </Grid>
        </Container>
    );
}

/**
 * Hooks
 */

function useWorkspaceExists(project: string, workspace: string, ts: number): boolean {
    const [ws,] = useAsync(() => listWorkspaces(project).then(wss => wss.includes(workspace)), [project, ts])
    return (ws === undefined) || ws;
}

function useJob(project: string, workspace: string): [Job | undefined, () => void] {
    const [j, reload] = useAsync(() => getJob(project, workspace), [project, workspace]);

    useEffect(() => {
        const t = setInterval(reload, 5000);
        return () => {
            clearInterval(t);
        };
    }, [reload]);

    return [j, reload];
}

/**
 * Helper Functions
 */

type State = {
    status: string
    planAvailable: boolean
    applyAvailable: boolean
}

function getState(job: Job | undefined): State {
    const [status, planAvailable, applyAvailable] = (() => {
        if (!job) {
            return ["Unknown", true, false];
        }
        switch (job.getStatus()) {
            case Job.Status.PENDING:
                return ["Pending", false, false];
            case Job.Status.PLANINPROGRESS:
                return ["Plan In Progress", false, false];
            case Job.Status.APPLYPENDING:
                return ["Apply Pending", false, false];
            case Job.Status.APPLYINPROGRESS:
                return ["Apply In Progress", false, false];

            case Job.Status.REVIEWREQUIRED:
                return ["Review Required", true, true];

            case Job.Status.SUCCEEDED:
                return ["Succeeded", true, false];
            case Job.Status.PLANFAILED:
                return ["Plan Failed", true, false];
            case Job.Status.APPLYFAILED:
                return ["Apply Failed", true, false];
            default:
                throw new Error(`unknown status: ${job.getStatus()}`);
        }
    })() as [string, boolean, boolean];

    return {
        status: status,
        planAvailable: planAvailable,
        applyAvailable: applyAvailable,
    }
}