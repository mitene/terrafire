import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Title from "./Title";
import {Link as RouterLink} from "react-router-dom";
import React, {useEffect, useState} from "react";
import * as globalStyle from "./styles"
import * as api from "../Api";
import CardActionArea from "@material-ui/core/CardActionArea";
import {CardContent} from "@material-ui/core";
import Card from "@material-ui/core/Card";

function useWorkspaces(project: string): api.Workspace[] {
    const [ws, setWs] = useState<api.Workspace[]>([]);

    useEffect(() => {
        api.listWorkspaces(project).then((result) => {
            setWs(Array.from(result.values()));
        })
    }, [project])

    return ws;
}

export function WorkspacesList(props: any) {
    const project = props.match.params.project;

    const classes = globalStyle.useStyles();
    const workspaces = useWorkspaces(project);

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={1}>
                {workspaces.map(ws => {
                    return (
                        <Grid item xs={12} key={ws.name}>
                            <Card>
                                <CardActionArea component={RouterLink} to={`/projects/${project}/workspaces/${ws.name}`}>
                                    <CardContent>
                                        <Title>{ws.name}</Title>
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
