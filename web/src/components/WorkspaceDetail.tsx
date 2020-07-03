import {approveJob, cancelJob, getJob, listWorkspaces, submitJob} from "../api";
import {Job} from "../api/common_pb";
import {Redirect} from "react-router";
import {useAsync} from "../hooks";
import Title from "./Title";
import React, {useEffect, useState} from "react";
import {Button, Container, Grid, IconButton, Menu, MenuItem, Paper, Typography} from "@material-ui/core";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import {makeStyles} from "@material-ui/core/styles";
import {useCommonStyles} from "./styles";

type Props = {
    project: string
    workspace: string
    ts: number
}

export const WorkspaceDetail: React.FC<Props> = (props) => {
    const project = props.project;
    const workspace = props.workspace;

    const classes = {...useCommonStyles(), ...useStyles()};
    const wsExists = useWorkspaceExists(project, workspace, props.ts);
    const [j, reload] = useJob(project, workspace);
    const [menuAnchor, setMenuAnchor] = useState<null | HTMLElement>(null);

    if (!wsExists) {
        return (
            <Redirect to={`/projects/${project}/workspaces`}/>
        );
    }

    const state = getState(j);
    const handlePlan = () => submitJob(project, workspace).then(reload).catch(console.log);
    const handleApply = () => approveJob(project, workspace).then(reload).catch(console.log);
    const handleMenuClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        setMenuAnchor(event.currentTarget);
    };
    const handleMenuClose = () => {
        setMenuAnchor(null);
    };
    const handleCancel = () => {
        handleMenuClose();
        cancelJob(project, workspace).then(reload).catch(console.log);
    };

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={1}>
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>{workspace}</Title>
                        <Typography>Status: {state.status}</Typography>
                        <Grid container spacing={3}>
                            <Grid item xs>
                                <Button size="small" variant="contained" color="primary" className={classes.button}
                                        onClick={handlePlan}
                                        disabled={!state.planAvailable}>Plan</Button>
                                <Button size="small" variant="contained" color="secondary" className={classes.button}
                                        onClick={handleApply}
                                        disabled={!state.applyAvailable}>Apply</Button>

                                <IconButton
                                        aria-controls="job-menu" aria-haspopup="true" onClick={handleMenuClick}>
                                    <ExpandMoreIcon/>
                                </IconButton>
                                <Menu
                                    id="job-menu"
                                    anchorEl={menuAnchor}
                                    keepMounted
                                    open={Boolean(menuAnchor)}
                                    onClose={handleMenuClose}
                                >
                                    <MenuItem onClick={handleCancel} disabled={!state.cancelAvailable}>Cancel</MenuItem>
                                </Menu>
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

const useStyles = makeStyles({
    button: {
        margin: 5,
    },
})

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
    cancelAvailable: boolean
}

function getState(job: Job | undefined): State {
    const [status, planAvailable, applyAvailable, cancelAvailable] = (() => {
        if (!job) {
            return ["Unknown", true, false, false];
        }
        switch (job.getStatus()) {
            case Job.Status.PENDING:
                return ["Pending", false, false, true];
            case Job.Status.PLANINPROGRESS:
                return ["Plan In Progress", false, false, true];
            case Job.Status.APPLYPENDING:
                return ["Apply Pending", false, false, true];
            case Job.Status.APPLYINPROGRESS:
                return ["Apply In Progress", false, false, true];

            case Job.Status.REVIEWREQUIRED:
                return ["Review Required", true, true, false];

            case Job.Status.SUCCEEDED:
                return ["Succeeded", true, false, false];
            case Job.Status.PLANFAILED:
                return ["Plan Failed", true, false, false];
            case Job.Status.APPLYFAILED:
                return ["Apply Failed", true, false, false];
            default:
                throw new Error(`unknown status: ${job.getStatus()}`);
        }
    })() as [string, boolean, boolean, boolean];

    return {
        status: status,
        planAvailable: planAvailable,
        applyAvailable: applyAvailable,
        cancelAvailable: cancelAvailable,
    }
}