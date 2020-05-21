import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";
import Title from "./Title";
import Typography from "@material-ui/core/Typography";
import React, {useEffect, useState} from "react";
import * as globalStyle from "./styles"
import * as api from "./Api";

function useJob(job_id: number): [api.Job | null, () => void] {
    const [job, setJob] = useState<api.Job | null>(null);

    function reload() {
        api.getJob(job_id).then(setJob);
    }

    useEffect(reload, [job_id])

    return [job, reload];
}


export default function JobDetail(props: any) {
    const job_id = props.match.params.job_id

    const classes = globalStyle.useStyles();
    const [job] = useJob(job_id)

    return (
        job &&
        <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={2}>
                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>{job.workspace}</Title>
                        <Typography>{api.JobStatus[job.status]}</Typography>
                    </Paper>
                </Grid>

                <Grid item xs={12}>
                    <Paper className={classes.paper}>
                        <Title>Plan Result</Title>
                        <pre><code>{job.plan_result}</code></pre>
                    </Paper>
                </Grid>
            </Grid>
        </Container>
    )
}
