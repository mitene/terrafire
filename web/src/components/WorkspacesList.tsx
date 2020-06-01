import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Title from "./Title";
import {Link as RouterLink} from "react-router-dom";
import React from "react";
import * as globalStyle from "./styles"
import CardActionArea from "@material-ui/core/CardActionArea";
import {CardContent} from "@material-ui/core";
import Card from "@material-ui/core/Card";
import {Project} from "../api/common_pb"

type Props = {
    project: Project
}

export const WorkspacesList: React.FC<Props> = (props) => {
    const classes = globalStyle.useStyles();
    const pj = props.project;
    const wss = props.project.getWorkspacesList()

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={1}>
                {wss.map(ws => {
                    return (
                        <Grid item xs={12} key={ws.getName()}>
                            <Card>
                                <CardActionArea component={RouterLink}
                                                to={`/projects/${pj.getName()}/workspaces/${ws.getName()}`}>
                                    <CardContent>
                                        <Title>{ws.getName()}</Title>
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
