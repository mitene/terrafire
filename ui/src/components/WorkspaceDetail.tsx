import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";
import Title from "./Title";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import {Link as RouterLink} from "react-router-dom";
import React, {useEffect, useState} from "react";

import * as globalStyle from "./styles"
import * as api from "../Api";
import Card from "@material-ui/core/Card";
import CardActionArea from "@material-ui/core/CardActionArea";
import {CardContent} from "@material-ui/core";

function useWorkspace(project: string, workspace: string): [api.Workspace | null, () => void] {
    const [w, setW] = useState<api.Workspace | null>(null);

    function reload() {
        api.getWorkspace(project, workspace).then(setW);
    }

    useEffect(() => {
        const f = () => { api.getWorkspace(project, workspace).then(setW) };
        f();
        const t = setInterval(f, 5000);
        return () => { clearInterval(t) }
    }, [project, workspace])

    return [w, reload];
}

function useJobs(project: string, workspace: string): [api.Job[], () => void] {
    const [js, setJs] = useState<api.Job[]>([]);

    function reload() {
        api.listJobs(project, workspace).then(setJs);
    }

    useEffect(reload, [project, workspace])

    return [js, reload];
}

export function WorkspaceDetail(props: any) {
    const project = props.match.params.project;
    const workspace = props.match.params.workspace;

    const classes = globalStyle.useStyles();
    const [ws, reloadWs] = useWorkspace(project, workspace);
    const [js, reloadJs] = useJobs(project, workspace);

    const planAvailable = ws && (!ws.last_job || ![api.JobStatus.Pending, api.JobStatus.PlanInProgress, api.JobStatus.ApplyInProgress].includes(ws.last_job.status));
    const applyAvailable = ws && ws.last_job && [api.JobStatus.ReviewRequired].includes(ws.last_job.status);

    function submitJob(e: any) {
        api.submitJob(project, workspace).then(() => {
            reloadWs();
            reloadJs();
        });
    }

    function approveJob(e: any) {
        api.approveJob(project, workspace).then(() => {
            reloadWs();
            reloadJs();
        });
    }

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={1}>
                {ws &&
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>{ws.name}</Title>
                        {ws && ws.last_job &&
                        <Typography>Status: {api.JobStatus[ws.last_job.status]}</Typography>
                        }
                        <Grid container spacing={3}>
                            <Grid item xs>
                                <Button size="small" variant="contained" color="primary"
                                        onClick={submitJob}
                                        disabled={!planAvailable}>Plan</Button>
                                <Button size="small" variant="contained" color="secondary"
                                        onClick={approveJob}
                                        disabled={!applyAvailable}>Apply</Button>
                                {/*<Button size="small" variant="contained" component={RouterLink} to={"/job/" + ws.job_id} disabled={!jobAvailable}>Detail</Button>*/}
                            </Grid>
                        </Grid>
                    </Paper>
                </Grid>
                }
                {ws && ws.last_job && ws.last_job.apply_log &&
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>Apply Log</Title>
                        <pre><code>{ws.last_job.apply_log}</code></pre>
                    </Paper>
                </Grid>
                }
                {ws && ws.last_job && ws.last_job.plan_result &&
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>Plan Result</Title>
                        <pre><code>{ws.last_job.plan_result}</code></pre>
                    </Paper>
                </Grid>
                }
                {ws && ws.last_job && ws.last_job.plan_log &&
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>Plan Log</Title>
                        <pre><code>{ws.last_job.plan_log}</code></pre>
                    </Paper>
                </Grid>
                }
            </Grid>

            <Grid container spacing={1}>
                {js.map((j) => {
                    return (
                        <Grid item xs={12} key={j.id}>
                            <Card>
                                <CardActionArea component={RouterLink} to={`/jobs/${j.id}`}>
                                    <CardContent>
                                        <Grid container>
                                            <Grid item xs={1}><Typography>#{j.id}</Typography></Grid>
                                            <Grid item xs={3}><Typography>{api.JobStatus[j.status]}</Typography></Grid>
                                            <Grid item xs={3}><Typography>{j.workspace}</Typography></Grid>
                                            <Grid item xs={3}><Typography>{j.error}</Typography></Grid>
                                        </Grid>
                                    </CardContent>
                                </CardActionArea>
                            </Card>
                        </Grid>
                    )
                })}
            </Grid>
        </Container>
    );
}
