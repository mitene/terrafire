import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Title from "./Title";
import React, {useEffect, useState} from "react";
import * as globalStyle from "./styles"
import * as api from "../Api";
import Card from "@material-ui/core/Card";
import CardActionArea from "@material-ui/core/CardActionArea";
import {Link as RouterLink} from "react-router-dom";
import {CardContent} from "@material-ui/core";

function useProjects(): api.Project[] {
    const [pj, setPj] = useState<api.Project[]>([]);
    useEffect(() => {
        api.listProjects().then(setPj);
    }, [])
    return pj;
}

export function ProjectsList(props: any) {
    const classes = globalStyle.useStyles();
    const project = useProjects();

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={1}>
                {project.map(pj => {
                    return (
                        <Grid item xs={12} key={pj.name}>
                            <Card>
                                <CardActionArea component={RouterLink} to={`/projects/${pj.name}/workspaces`}>
                                    <CardContent>
                                        <Title>{pj.name}</Title>
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
