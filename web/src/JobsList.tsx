import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Card from "@material-ui/core/Card";
import CardActionArea from "@material-ui/core/CardActionArea";
import {Link as RouterLink} from "react-router-dom";
import {CardContent} from "@material-ui/core";
import Typography from "@material-ui/core/Typography";
import React, {useEffect, useState} from "react";

import * as globalStyle from './styles';
import * as api from "./Api";

function useJobs(project: string, workspace: string): api.Job[] {
    const [jobs, setJobs] = useState<api.Job[]>([]);

    useEffect(() => {
        api.listJobs(project, workspace).then((result) => {
            setJobs(result);
        })
    }, [project, workspace])

    return jobs;
}

export default function JobsList(props: any) {
    const project = props.match.params.project;
    const workspace = props.match.params.workspace;

    const classes = globalStyle.useStyles();
    const jobs = useJobs(project, workspace);

    console.log(jobs);

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={1}>
                {jobs.map((job, i) => {
                    return (
                        <Grid item xs={12} key={i}>
                            <Card>
                                <CardActionArea component={RouterLink} to={"/job/" + i}>
                                    <CardContent>
                                        <Grid container>
                                            <Grid item xs={1}><Typography>#{i}</Typography></Grid>
                                            <Grid item xs={3}><Typography>{job.status}</Typography></Grid>
                                            <Grid item xs={3}><Typography>{job.workspace}</Typography></Grid>
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
