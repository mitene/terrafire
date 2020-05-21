package terrafire

import "fmt"

func (r ProjectRepository) GetProject(project string) (*ProjectInfo, error) {
	pj, ok := r[project]
	if !ok {
		return nil, fmt.Errorf("project is not defined: %s", project)
	}
	return pj, nil
}

func (r ProjectRepository) GetWorkspace(project string, workspace string) (*ProjectInfo, *Workspace, error) {
	pj, err := r.GetProject(project)
	if err != nil {
		return nil, nil, err
	}

	ws, ok := pj.Manifest.Workspaces[workspace]
	if !ok {
		return pj, nil, fmt.Errorf("workspace is not defined: %s/%s", project, workspace)
	}

	return pj, ws, nil
}
