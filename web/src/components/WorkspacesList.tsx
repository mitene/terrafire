import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Title from "./Title";
import {Link as RouterLink} from "react-router-dom";
import React from "react";
import * as globalStyle from "./styles"
import CardActionArea from "@material-ui/core/CardActionArea";
import {CardContent} from "@material-ui/core";
import Card from "@material-ui/core/Card";
import {listWorkspaces} from "../api";
import {useAsync} from "../hooks";

type Props = {
    project: string
    ts: number
}

export const WorkspacesList: React.FC<Props> = (props) => {
    const project = props.project;

    const classes = globalStyle.useStyles();
    const wss = useWorkspaces(project, props.ts);

    return (
        <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={1}>
                {wss.map(ws => {
                    return (
                        <Grid item xs={12} key={ws}>
                            <Card>
                                <CardActionArea component={RouterLink}
                                                to={`/projects/${project}/workspaces/${ws}`}>
                                    <CardContent>
                                        <Title>{ws}</Title>
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

/**
 * Hooks
 */

function useWorkspaces(project: string, ts: number): string[] {
    const [wss,] = useAsync(() => listWorkspaces(project), [project, ts]);
    return wss || [];
}
